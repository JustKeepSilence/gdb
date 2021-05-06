/*
creatTime: 2020/11/27
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.16
*/

package db

import (
	"fmt"
	"github.com/JustKeepSilence/gdb/cmap"
	. "github.com/ahmetb/go-linq/v3"
	"github.com/dop251/goja"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"golang.org/x/sync/errgroup"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// BatchWrite write realTimeData, all items to be written should be existed in gdb.
// otherWise will fail to writing.
func (gdb *Gdb) BatchWrite(infos ...ItemValue) (Rows, error) {
	itemNames, itemValues := []string{}, []string{}
	// without timeStamp
	for _, itemValue := range infos {
		k := itemValue.ItemName
		if len(k) == 0 {
			return Rows{}, fmt.Errorf("error itemName")
		}
		v := itemValue.Value
		itemNames = append(itemNames, k)
		itemValues = append(itemValues, v)
	}
	index, ok := gdb.checkItems(itemNames...)
	// check if all items given exist
	if !ok {
		return Rows{}, itemError{"itemError: " + itemNames[index]}
	}
	currentTimeStamp := int(time.Now().Unix()) + 8*3600
	currentTimeStampString := strconv.Itoa(currentTimeStamp)
	g := errgroup.Group{}
	// write currentTimeStamp
	g.Go(func() error {
		if err := gdb.infoDb.Put([]byte(timeKey), []byte(currentTimeStampString), nil); err != nil {
			return err
		}
		return nil
	})
	// write Realtime data
	g.Go(func() error {
		batch := leveldb.Batch{}
		for i := 0; i < len(itemNames); i++ {
			batch.Put([]byte(itemNames[i]), []byte(itemValues[i]))
		}
		if err := gdb.rtDb.Write(&batch, nil); err != nil {
			return err
		}
		return nil
	})
	// write historical data
	g.Go(func() error {
		batch := leveldb.Batch{}
		for i := 0; i < len(itemNames); i++ {
			sb := strings.Builder{}
			sb.Write([]byte(itemNames[i]))
			sb.Write([]byte(currentTimeStampString))
			batch.Put([]byte(sb.String()), []byte(itemValues[i]))
		}
		if err := gdb.hisDb.Write(&batch, nil); err != nil {
			return err
		} else {
			return nil
		}
	})
	if err := g.Wait(); err != nil {
		return Rows{}, err
	} else {
		return Rows{len(itemNames)}, nil
	}
}

// BatchWriteHistoricalData write historicalData, all items to be written should be existed in gdb.
// and the timeStamp should be unix timeStamp
func (gdb *Gdb) BatchWriteHistoricalData(infos ...HistoricalItemValue) error {
	itemNames := []string{}
	for i := 0; i < len(infos); i++ {
		itemNames = append(itemNames, infos[i].ItemName)
	}
	if index, ok := gdb.checkItems(itemNames...); !ok {
		return fmt.Errorf("item " + itemNames[index] + " not existed")
	}
	g := errgroup.Group{}
	for _, item := range infos {
		info := item
		g.Go(func() error {
			itemName, values, timeStamps := info.ItemName, info.Values, info.TimeStamps
			if len(values) != len(timeStamps) {
				return fmt.Errorf("inconsistent length of values and timestamps")
			} else {
				batch := leveldb.Batch{}
				for i := 0; i < len(values); i++ {
					sb := strings.Builder{}
					sb.Write([]byte(itemName))
					sb.Write([]byte(timeStamps[i]))
					batch.Put([]byte(sb.String()), []byte(values[i]))
				}
				if err := gdb.hisDb.Write(&batch, nil); err != nil {
					return err
				} else {
					return nil
				}
			}
		})
	}
	if err := g.Wait(); err != nil {
		return err
	} else {
		return nil
	}
}

// GetRealTimeData get realTime data,that is the latest updated value of item.All items should be
// existed in gdb, otherWise will fail to getting data
func (gdb *Gdb) GetRealTimeData(itemNames ...string) (cmap.ConcurrentMap, error) {
	if index, ok := gdb.checkItems(itemNames...); !ok {
		return nil, fmt.Errorf("item " + itemNames[index] + " not existed")
	} else {
		sn, err := gdb.rtDb.GetSnapshot()
		if sn == nil || err != nil {
			return nil, snError{"snError"}
		}
		defer sn.Release() // release
		m := cmap.New()
		g := errgroup.Group{}
		for _, itemName := range itemNames {
			name := itemName
			g.Go(func() error {
				v, err := sn.Get([]byte(name), nil)
				if err != nil {
					//  key not exist
					m.Set(name, nil)
				} else {
					m.Set(name, fmt.Sprintf("%s", v)) // set values
				}
				return nil
			})
		}
		if err := g.Wait(); err != nil {
			return nil, err
		} else {
			return m, nil
		}
	}
}

// GetHistoricalData get historical data, timeStamp should be unix timeStamp, and units of
// interval is seconds.All items should be existed in gdb, otherWise will fail to getting data
func (gdb *Gdb) GetHistoricalData(itemNames []string, startTimeStamps []int, endTimeStamps []int, intervals []int) (cmap.ConcurrentMap, error) {
	if len(startTimeStamps) == len(endTimeStamps) && len(endTimeStamps) == len(intervals) {
		if index, ok := gdb.checkItems(itemNames...); !ok {
			return nil, fmt.Errorf("item " + itemNames[index] + " not existed")
		} else {
			rawData := cmap.New()
			sn, err := gdb.hisDb.GetSnapshot()
			if sn == nil || err != nil {
				return nil, snError{"snError"}
			}
			defer sn.Release() // release
			g := errgroup.Group{}
			for _, itemName := range itemNames {
				name := itemName
				g.Go(func() error {
					var values []string
					var timeStamps []int64
					for i := 0; i < len(startTimeStamps); i++ {
						s := startTimeStamps[i] // startTime
						e := endTimeStamps[i]   // endTime
						interval := intervals[i]
						if s >= e {
							// startTime > endTime, continue
							continue
						}
						startKey := strings.Builder{}
						startKey.Write([]byte(name))
						startKey.Write([]byte(strconv.Itoa(s)))
						endKey := strings.Builder{}
						endKey.Write([]byte(name))
						endKey.Write([]byte(strconv.Itoa(e)))
						it := sn.NewIterator(&util.Range{Start: []byte(startKey.String()), Limit: []byte(endKey.String())}, nil)
						count := 0
						var st int64 // start time stamp of data
						for it.Next() {
							tt, _ := strconv.ParseInt(strings.Replace(fmt.Sprintf("%s", it.Key()), name, "", -1), 10, 64) // get time stamp
							if count == 0 {
								st = tt // get start time stamp of data
								values = append(values, fmt.Sprintf("%s", it.Value()))
								timeStamps = append(timeStamps, tt)
							} else {
								if (tt-st)%int64(interval) == 0 {
									values = append(values, fmt.Sprintf("%s", it.Value()))
									timeStamps = append(timeStamps, tt)
								}
							}
							count++
						}
					}
					rawData.Set(name, []interface{}{timeStamps, values})
					return nil
				})
			}
			if err := g.Wait(); err != nil {
				return nil, err
			} else {
				return rawData, nil
			}
		}
	} else {
		return nil, fmt.Errorf("inconsistent length of startTimeStamps, endTimeStamps, intervals")
	}
}

// GetRawHistoricalData get raw(that is all  historical data, it should only be used for debugging.
// All items should be existed in gdb, otherWise will fail to getting data
func (gdb *Gdb) GetRawHistoricalData(itemNames ...string) (cmap.ConcurrentMap, error) {
	if index, ok := gdb.checkItems(itemNames...); !ok {
		return nil, fmt.Errorf("item " + itemNames[index] + " not existed")
	} else {
		rawData := cmap.New()
		g := errgroup.Group{}
		sn, err := gdb.hisDb.GetSnapshot()
		if sn == nil || err != nil {
			return nil, snError{"snError"}
		}
		defer sn.Release()
		for _, name := range itemNames {
			itemName := name
			g.Go(func() error {
				itemValues, timeStamps := []string{}, []int64{}
				it := sn.NewIterator(util.BytesPrefix([]byte(itemName)), nil)
				for it.Next() {
					t, err := strconv.ParseInt(strings.Replace(fmt.Sprintf("%s", it.Key()), itemName, "", -1), 10, 64)
					if err != nil {
						//return err
					}
					v := fmt.Sprintf("%s", it.Value())
					itemValues = append(itemValues, v)
					timeStamps = append(timeStamps, t)
				}
				rawData.Set(itemName, []interface{}{timeStamps, itemValues})
				return nil
			})
		}
		if err := g.Wait(); err != nil {
			return nil, err
		} else {
			return rawData, nil
		}
	}
}

// GetHistoricalDataWithStamp get history data according to the given time stamps
func (gdb *Gdb) GetHistoricalDataWithStamp(itemNames []string, timeStamps ...[]int) (cmap.ConcurrentMap, error) {
	if len(itemNames) != len(timeStamps) {
		return nil, fmt.Errorf("inconsistent length of itemNames and timeStamps")
	}
	if index, ok := gdb.checkItems(itemNames...); !ok {
		return nil, fmt.Errorf("item " + itemNames[index] + " not existed")
	} else {
		rawData := cmap.New()
		sn, err := gdb.hisDb.GetSnapshot()
		if sn == nil || err != nil {
			return nil, snError{"snError"}
		}
		defer sn.Release() // release
		g := errgroup.Group{}
		for index, itemName := range itemNames {
			name, i := itemName, index
			g.Go(func() error {
				var values []string
				for j := 0; j < len(timeStamps[i]); j++ {
					v, _ := sn.Get([]byte(name+fmt.Sprintf("%d", timeStamps[i][j])), nil)
					values = append(values, fmt.Sprintf("%s", v))
				}
				rawData.Set(name, []interface{}{values, timeStamps[i]})
				return nil
			})
		}
		if err := g.Wait(); err != nil {
			return nil, err
		} else {
			return rawData, nil
		}
	}
}

// GetHistoricalDataWithCondition filter condition must be correct js expression,itemName should be startedWith by item.
// eg: item["itemName1"]>10 && item["itemName2"] > 30 ....
// It should be noted that the entire judgment is based on the itemName with less historical value in the condition.
// If the longest itemName is used as the benchmark, we cannot make an accurate judgment on the AND logic in it.
// Just imagine the history of Item1 It is [3,4,5], and item2 is [10,11]. If item1 is used as the benchmark,
// we cannot determine how much other elements of item2 should be expanded, because the condition may have complicated
// logic about item1 and item2 And or logic, no matter what the number is expanded, there may be a judgment error.
// DeadZone is used to define the maximum number of continuous data allowed by itemName.eg,the deadZoneCount of item x
// is 2, that is all data in x whose number of continuous > 2 will be filtered
func (gdb *Gdb) GetHistoricalDataWithCondition(itemNames []string, startTime []int, endTime []int, intervals []int, filterCondition string, zones ...DeadZone) (cmap.ConcurrentMap, error) {
	var filterItemNames []string // filter item
	if strings.Trim(filterCondition, " ") == "true" {
		filterItemNames = itemNames
	} else {
		reg := regexp.MustCompile(`(?is)\["(.*?)"]`)
		if reg.Match([]byte(filterCondition)) {
			matchedResult := reg.FindAllStringSubmatch(filterCondition, -1)
			for _, mr := range matchedResult {
				filterItemNames = append(filterItemNames, mr[1])
			}
		} else {
			return nil, conditionError{"conditionError: invalid condition, items must be included by [] and condition can't be null "}
		}
	}
	filterHistoryData, err := gdb.getHistoricalDataWithMinLength(filterItemNames, startTime, endTime, intervals) // get history of filter item
	if err != nil {
		return nil, err
	}
	vm := goja.New()
	f := `function filterData(s){return s.filter(function(item){return ` + filterCondition + `})}`
	if _, err := vm.RunString(f); err != nil {
		return nil, conditionError{"conditionError: " + err.Error()}
	}
	filterData, ok := goja.AssertFunction(vm.Get("filterData"))
	if !ok {
		return nil, conditionError{"conditionError: fail compiling function"}
	}
	res, err := filterData(goja.Undefined(), vm.ToValue(filterHistoryData))
	if err != nil {
		return nil, conditionError{"conditionError: " + err.Error()}
	}
	filterResults := res.Export().([]interface{})
	var timeStamps []int64 // time stamp
	for _, fr := range filterResults {
		sfr := fr.(map[string]interface{})
		timeStamps = append(timeStamps, sfr["timeStamp"].(int64))
	}
	if len(zones) != 0 {
		if data, err := gdb.getHistoricalDataWithStampAndDeadZoneCount(itemNames, timeStamps, zones...); err != nil {
			return nil, err
		} else {
			return data, nil
		}
	} else {
		if data, err := gdb.getHistoricalDataWithStringTimeStamp(itemNames, timeStamps...); err != nil {
			return nil, err
		} else {
			return data, nil
		}
	}
}

func (gdb *Gdb) getHistoricalDataWithStringTimeStamp(itemNames []string, timeStamps ...int64) (cmap.ConcurrentMap, error) {
	rawData := cmap.New()
	sn, err := gdb.hisDb.GetSnapshot()
	if sn == nil || err != nil {
		return nil, snError{"snError"}
	}
	defer sn.Release() // release
	g := errgroup.Group{}
	for _, itemName := range itemNames {
		name := itemName
		g.Go(func() error {
			var values []string
			for j := 0; j < len(timeStamps); j++ {
				v, _ := sn.Get([]byte(name+fmt.Sprintf("%d", timeStamps[j])), nil)
				values = append(values, fmt.Sprintf("%s", v))
			}
			rawData.Set(name, []interface{}{timeStamps, values})
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return nil, err
	} else {
		return rawData, nil
	}
}

func (gdb *Gdb) getHistoricalDataWithStampAndDeadZoneCount(itemNames []string, timeStamps []int64, zones ...DeadZone) (cmap.ConcurrentMap, error) {
	rawData := cmap.New()
	sn, err := gdb.hisDb.GetSnapshot()
	if sn == nil || err != nil {
		return nil, snError{"snError"}
	}
	defer sn.Release() // release
	g := errgroup.Group{}
	for _, itemName := range itemNames {
		name := itemName
		g.Go(func() error {
			var values []string
			var lastValue string
			var lastValues []string
			filterIndex := From(zones).IndexOf(func(item interface{}) bool {
				return item.(DeadZone).ItemName == name
			})
			if filterIndex == -1 {
				// don't filter
				for i := 0; i < len(timeStamps); i++ {
					v, _ := sn.Get([]byte(name+fmt.Sprintf("%d", timeStamps[i])), nil)
					values = append(values, fmt.Sprintf("%s", v))
				}
				rawData.Set(name, []interface{}{timeStamps, values})
				return nil
			}
			deadZoneCounts := zones[filterIndex].DeadZoneCount
			if deadZoneCounts == 1 {
				rawData.Set(name, [][]string{{}, {}})
				return nil
			} else if deadZoneCounts < 1 {
				// don't filter
				for i := 0; i < len(timeStamps); i++ {
					v, _ := sn.Get([]byte(name+fmt.Sprintf("%d", timeStamps[i])), nil)
					values = append(values, fmt.Sprintf("%s", v))
				}
			} else {
				// filter
				for i := 0; i < len(timeStamps); i++ {
					v, _ := sn.Get([]byte(name+fmt.Sprintf("%d", timeStamps[i])), nil)
					vs := fmt.Sprintf("%s", v)
					if lastValue != vs {
						// not Repeated
						for _, lv := range lastValues {
							values = append(values, lv)
						}
						lastValue = vs
						lastValues = []string{}
						values = append(values, vs)
					} else {
						// repeated
						lastValues = append(lastValues, lastValue)
						if len(lastValues)+1 == deadZoneCounts {
							lastValues = []string{}
						}
					}
				}
			}
			rawData.Set(name, []interface{}{timeStamps, values})
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return nil, err
	} else {
		return rawData, nil
	}
}

func (gdb *Gdb) getHistoricalDataWithMinLength(itemNames []string, startTime []int, endTime []int, intervals []int) ([]map[string]interface{}, error) {
	rawData := cmap.New()
	lengthMap := make([]int, len(itemNames))
	sn, err := gdb.hisDb.GetSnapshot()
	if sn == nil || err != nil {
		return nil, snError{"snError"}
	}
	defer sn.Release() // release
	g := errgroup.Group{}
	for index, itemName := range itemNames {
		name, j := itemName, index
		g.Go(func() error {
			var values []string
			var timeStamps []int64
			for i := 0; i < len(startTime); i++ {
				s := startTime[i] // startTime
				e := endTime[i]   // endTime
				interval := intervals[i]
				if s >= e {
					// startTime > endTime, continue
					continue
				}
				startKey := strings.Builder{}
				startKey.Write([]byte(name))
				startKey.Write([]byte(strconv.Itoa(s)))
				endKey := strings.Builder{}
				endKey.Write([]byte(name))
				endKey.Write([]byte(strconv.Itoa(e)))
				it := sn.NewIterator(&util.Range{Start: []byte(startKey.String()), Limit: []byte(endKey.String())}, nil)
				count := 0
				var st int64 // start time stamp of data
				for it.Next() {
					tt, _ := strconv.ParseInt(strings.Replace(fmt.Sprintf("%s", it.Key()), name, "", -1), 10, 64) // get time stamp
					if count == 0 {
						st = tt // get start time stamp of data
						values = append(values, fmt.Sprintf("%s", it.Value()))
						timeStamps = append(timeStamps, tt)
					} else {
						if (tt-st)%int64(interval) == 0 {
							values = append(values, fmt.Sprintf("%s", it.Value()))
							timeStamps = append(timeStamps, tt)
						}
					}
					count++
				}
			}
			rawData.Set(name, []interface{}{timeStamps, values})
			lengthMap[j] = len(values)
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return nil, err
	} else {
		var result []map[string]interface{}
		minLength := From(lengthMap).Min().(int)
		for i := 0; i < minLength; i++ {
			t := map[string]interface{}{}
			for index, name := range itemNames {
				v, _ := rawData.Get(name)
				vs := v.([]interface{})
				d, err := strconv.ParseFloat(vs[1].([]string)[i], 64)
				if err != nil {
					break
				}
				t[name] = d // values
				if index == 0 {
					// first
					t["timeStamp"] = vs[0].([]int64)[i]
				} else {
					et := t["timeStamp"]
					if et != vs[0].([]int64)[i] {
						break // inconsistent timeStamp
					}
				}
			}
			result = append(result, t)
		}
		return result, nil
	}
}

// check if all itemNames given exist
func (gdb *Gdb) checkItems(itemNames ...string) (int, bool) {
	for index, itemName := range itemNames {
		_, ok := gdb.rtDbFilter.Get(itemName)
		if !ok {
			return index, false
		}
	}
	return -1, true
}

// following method is used by calc
// get unix timestamp of the given time,t should b yyyy-mm-dd hh:mm:ss
func (gdb *Gdb) getUnixTimeStamp(t string) int64 {
	t1, err := time.Parse("2006-01-02 15:04:05", t)
	if err != nil {
		return -1
	}
	return t1.Unix()
}

func (gdb *Gdb) getNowTime() string {
	return time.Now().Format(timeFormatString)
}

func (gdb *Gdb) getTime(d int) string {
	return time.Now().Add(time.Duration(d) * time.Second).Format(timeFormatString)
}

func (gdb *Gdb) getRtData(itemNames []string) ([]string, error) {
	v, err := gdb.GetRealTimeData(itemNames...)
	if err != nil {
		return nil, err
	}
	var result []string
	for _, name := range itemNames {
		rv, _ := v.Get(name)
		if rv == nil {
			return nil, fmt.Errorf("invalid itemName: " + name)
		}
		result = append(result, rv.(string))
	}
	return result, nil
}

func (gdb *Gdb) getHData(itemNames []string, timeStamps []int64) ([]string, error) {
	r, err := gdb.getHistoricalDataWithStringTimeStamp(itemNames, timeStamps...)
	var result []string
	if err != nil {
		return nil, err
	}
	for _, itemName := range itemNames {
		rv, _ := r.Get(itemName)
		if rv == nil {
			return nil, fmt.Errorf("invalid itemName: " + itemName)
		}
		result = append(result, rv.(string))
	}
	return result, nil
}

func (gdb *Gdb) getDbInfo() (cmap.ConcurrentMap, error) {
	sn, err := gdb.infoDb.GetSnapshot()
	itemNames := []string{ram, writtenItems, timeKey, speed}
	if sn == nil || err != nil {
		return nil, snError{"snError"}
	}
	defer sn.Release() // release
	m := cmap.New()
	g := errgroup.Group{}
	for _, itemName := range itemNames {
		name := itemName
		g.Go(func() error {
			v, err := sn.Get([]byte(name), nil)
			if err != nil {
				//  key not exist
				m.Set(name, nil)
			} else {
				m.Set(name, fmt.Sprintf("%s", v)) // set values
			}
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

func (gdb *Gdb) getDbInfoHistory(itemName string, startTimeStamps []int, endTimeStamps []int, interval int) (cmap.ConcurrentMap, error) {
	rawData := cmap.New()
	sn, err := gdb.infoDb.GetSnapshot()
	if sn == nil || err != nil {
		return nil, snError{"snError"}
	}
	defer sn.Release() // release
	for i := 0; i < len(startTimeStamps); i++ {
		s := startTimeStamps[i] // startTime
		e := endTimeStamps[i]   // endTime
		if s >= e {
			// startTime > endTime, continue
			continue
		}
		startKey := strings.Builder{}
		startKey.Write([]byte(itemName))
		startKey.Write([]byte(fmt.Sprintf("%d", s)))
		endKey := strings.Builder{}
		endKey.Write([]byte(itemName))
		endKey.Write([]byte(fmt.Sprintf("%d", e)))
		it := sn.NewIterator(&util.Range{Start: []byte(startKey.String()), Limit: []byte(endKey.String())}, nil)
		var values []string
		var count int32
		var timeStamps []string
		for it.Next() {
			if int(count)%interval == 0 {
				values = append(values, fmt.Sprintf("%s", it.Value()))
				timeStamps = append(timeStamps, strings.Replace(fmt.Sprintf("%s", it.Key()), itemName, "", -1))
			}
			count++
		}
		rawData.Set(itemName, [][]string{timeStamps, values})
	}
	return rawData, nil
}
