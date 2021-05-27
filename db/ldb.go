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
	"math"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// BatchWrite write realTimeData, all items to be written should be existed in gdb.
// otherWise will fail to writing.
func (gdb *Gdb) BatchWrite(infos ...ItemValue) (Rows, error) {
	itemNames, itemValues, groupNames := []string{}, []string{}, []string{}
	// without timeStamp
	for _, itemValue := range infos {
		k := strings.Trim(itemValue.ItemName, " ")
		if len(k) == 0 {
			return Rows{}, fmt.Errorf("itemName can't be empty string")
		}
		if t, ok := gdb.rtDbFilter.Get(k + joiner + itemValue.GroupName); !ok {
			return Rows{}, itemError{"itemError: " + k}
		} else {
			v := itemValue.Value
			vt := reflect.TypeOf(v).String()
			if vt == t.(string) || t == "int64" {
				// check whether type of value and item is consistent
				itemNames = append(itemNames, k)
				groupNames = append(groupNames, itemValue.GroupName)
				switch t {
				case "int64":
					itemValues = append(itemValues, fmt.Sprintf("%.f", v.(float64)))
					break
				case "float64":
					if math.IsNaN(v.(float64)) {
						return Rows{}, fmt.Errorf("you can't write NaN to database")
					}
					itemValues = append(itemValues, fmt.Sprintf("%f", v.(float64)))
					break
				case "bool":
					itemValues = append(itemValues, strconv.FormatBool(v.(bool)))
					break
				default:
					itemValues = append(itemValues, v.(string))
					break
				}
			} else {
				return Rows{}, fmt.Errorf("inconsistent type of value and item")
			}
		}
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
			// itemName = itemName + joiner + groupName
			batch.Put([]byte(itemNames[i]+joiner+groupNames[i]), []byte(itemValues[i]))
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
			// itemName = itemName + joiner + groupName + timeStamp
			sb.Write([]byte(itemNames[i] + joiner + groupNames[i]))
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
	itemNames, itemValues, timeStamps, groupNames := []string{}, [][]string{}, [][]string{}, []string{}
	for i := 0; i < len(infos); i++ {
		itemName := infos[i].ItemName
		if t, ok := gdb.rtDbFilter.Get(itemName + joiner + infos[i].GroupName); !ok {
			return fmt.Errorf("item " + itemName + " not exsited")
		} else {
			itemNames = append(itemNames, itemName)
			groupNames = append(groupNames, infos[i].GroupName)
			if len(infos[i].TimeStamps) != len(infos[i].Values) {
				return fmt.Errorf("inconsistent length of values and timestamps")
			}
			values, ts := []string{}, []string{}
			for index := 0; index < len(infos[i].Values); index++ {
				v := infos[i].Values[index]
				vt := reflect.TypeOf(v).String()
				if vt == t.(string) || t == "int64" {
					// check whether type of value and item is consistent
					switch t {
					case "int64":
						// do not why reflect.TypeOf in go will judge int64 as float 64 after json serialization
						// so we convert int64 to float64 with %.f which will guarantee the strconv.ParseInt function
						// will work normally as well
						values = append(values, fmt.Sprintf("%.f", v.(float64)))
						break
					case "float64":
						if math.IsNaN(v.(float64)) {
							return fmt.Errorf("you can't write NaN to database")
						}
						values = append(values, fmt.Sprintf("%f", v.(float64)))
						break
					case "bool":
						values = append(values, strconv.FormatBool(v.(bool)))
						break
					default:
						values = append(values, v.(string))
						break
					}
					ts = append(ts, fmt.Sprintf("%d", infos[i].TimeStamps[index]))
					//return nil
				} else {
					return fmt.Errorf("inconsistent type of value and item")
				}
			}
			itemValues = append(itemValues, values)
			timeStamps = append(timeStamps, ts)
		}
	}
	g := errgroup.Group{}
	for j := 0; j < len(itemNames); j++ {
		index := j
		itemName, values, ts, groupName := itemNames[index], itemValues[index], timeStamps[index], groupNames[index]
		g.Go(func() error {
			batch := leveldb.Batch{}
			for i := 0; i < len(values); i++ {
				sb := strings.Builder{}
				sb.Write([]byte(itemName + joiner + groupName))
				sb.Write([]byte(ts[i]))
				batch.Put([]byte(sb.String()), []byte(values[i]))
			}
			if err := gdb.hisDb.Write(&batch, nil); err != nil {
				return err
			} else {
				return nil
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
func (gdb *Gdb) GetRealTimeData(groupNames []string, itemNames ...string) (cmap.ConcurrentMap, error) {
	dataTypes := []string{}
	for index, itemName := range itemNames {
		if t, ok := gdb.rtDbFilter.Get(itemName + joiner + groupNames[index]); !ok {
			return nil, fmt.Errorf("item " + itemName + " not existed")
		} else {
			dataTypes = append(dataTypes, t.(string))
		}
	}
	sn, err := gdb.rtDb.GetSnapshot()
	if sn == nil || err != nil {
		return nil, snError{"snError"}
	}
	defer sn.Release() // release
	m := cmap.New()
	g := errgroup.Group{}
	for index, itemName := range itemNames {
		name, i := itemName, index
		g.Go(func() error {
			v, err := sn.Get([]byte(name+joiner+groupNames[i]), nil)
			if err != nil || v == nil {
				//  realTime data not existed
				m.Set(name, nil)
			} else {
				switch dataTypes[i] {
				case "int64":
					if tv, err := strconv.ParseInt(string(v), 10, 64); err != nil {
						return err
					} else {
						m.Set(name, tv)
					}
					break
				case "float64":
					if tv, err := strconv.ParseFloat(string(v), 64); err != nil {
						return err
					} else {
						m.Set(name, tv) // set values
					}
					break
				case "bool":
					if tv, err := strconv.ParseBool(string(v)); err != nil {
						return err
					} else {
						m.Set(name, tv)
					}
					break
				default:
					m.Set(name, string(v))
					break
				}
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

// GetHistoricalData get historical data, timeStamp should be unix timeStamp, and units of
// interval is seconds.All items should be existed in gdb, otherWise will fail to getting data
func (gdb *Gdb) GetHistoricalData(groupNames, itemNames []string, startTimeStamps, endTimeStamps, intervals []int) (cmap.ConcurrentMap, error) {
	if len(startTimeStamps) == len(endTimeStamps) && len(endTimeStamps) == len(intervals) {
		dataTypes := []string{}
		for index, itemName := range itemNames {
			if t, ok := gdb.rtDbFilter.Get(itemName + joiner + groupNames[index]); !ok {
				return nil, fmt.Errorf("item " + itemName + " not existed")
			} else {
				dataTypes = append(dataTypes, t.(string))
			}
		}
		rawData := cmap.New()
		sn, err := gdb.hisDb.GetSnapshot()
		if sn == nil || err != nil {
			return nil, snError{"snError"}
		}
		defer sn.Release() // release
		g := errgroup.Group{}
		for j, itemName := range itemNames {
			name, index := itemName, j
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
					startKey.Write([]byte(name + joiner + groupNames[index]))
					startKey.Write([]byte(strconv.Itoa(s)))
					endKey := strings.Builder{}
					endKey.Write([]byte(name + joiner + groupNames[index]))
					endKey.Write([]byte(strconv.Itoa(e)))
					it := sn.NewIterator(&util.Range{Start: []byte(startKey.String()), Limit: []byte(endKey.String())}, nil)
					count := 0
					var st int64 // start time stamp of data
					for it.Next() {
						tt, _ := strconv.ParseInt(strings.Replace(fmt.Sprintf("%s", it.Key()), name+joiner+groupNames[index], "", -1), 10, 64) // get time stamp
						if count == 0 {
							st = tt // get start time stamp of data
							values = append(values, string(it.Value()))
							timeStamps = append(timeStamps, tt)
						} else {
							if (tt-st)%int64(interval) == 0 {
								values = append(values, string(it.Value()))
								timeStamps = append(timeStamps, tt)
							}
						}
						count++
					}
				}
				// values is nil === not historical data , and return nil
				if v, err := convertValues(dataTypes[index], values...); err != nil {
					return err
				} else {
					rawData.Set(name, []interface{}{timeStamps, v})
				}
				return nil
			})
		}
		if err := g.Wait(); err != nil {
			return nil, err
		} else {
			return rawData, nil
		}
	} else {
		return nil, fmt.Errorf("inconsistent length of startTimeStamps, endTimeStamps, intervals")
	}
}

// GetRawHistoricalData get raw(that is all  historical data, it should only be used for debugging.
// All items should be existed in gdb, otherWise will fail to getting data
func (gdb *Gdb) GetRawHistoricalData(groupNames []string, itemNames ...string) (cmap.ConcurrentMap, error) {
	dataTypes := []string{}
	for index, itemName := range itemNames {
		if t, ok := gdb.rtDbFilter.Get(itemName + joiner + groupNames[index]); !ok {
			return nil, fmt.Errorf("item " + itemName + " not existed")
		} else {
			dataTypes = append(dataTypes, t.(string))
		}
	}
	rawData := cmap.New()
	g := errgroup.Group{}
	sn, err := gdb.hisDb.GetSnapshot()
	if sn == nil || err != nil {
		return nil, snError{"snError"}
	}
	defer sn.Release()
	for i, name := range itemNames {
		itemName, index := name, i
		g.Go(func() error {
			values, timeStamps := []string{}, []int64{}
			it := sn.NewIterator(util.BytesPrefix([]byte(itemName+joiner+groupNames[index])), nil)
			for it.Next() {
				t, err := strconv.ParseInt(strings.Replace(fmt.Sprintf("%s", it.Key()),
					itemName+joiner+groupNames[index], "", -1), 10, 64)
				if err != nil {
					return err
				}
				values = append(values, string(it.Value()))
				timeStamps = append(timeStamps, t)
			}
			if v, err := convertValues(dataTypes[index], values...); err != nil {
				return err
			} else {
				rawData.Set(itemName, []interface{}{timeStamps, v})
			}
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return nil, err
	} else {
		return rawData, nil
	}
}

// GetHistoricalDataWithStamp get history data according to the given time stamps
func (gdb *Gdb) GetHistoricalDataWithStamp(groupNames, itemNames []string, timeStamps ...[]int) (cmap.ConcurrentMap, error) {
	dataTypes := []string{}
	for index, itemName := range itemNames {
		if t, ok := gdb.rtDbFilter.Get(itemName + joiner + groupNames[index]); !ok {
			return nil, fmt.Errorf("item " + itemName + " not existed")
		} else {
			dataTypes = append(dataTypes, t.(string))
		}
	}
	if len(itemNames) != len(timeStamps) {
		return nil, fmt.Errorf("inconsistent length of itemNames and timeStamps")
	}
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
				v, _ := sn.Get([]byte(name+joiner+groupNames[i]+fmt.Sprintf("%d", timeStamps[i][j])), nil)
				if v != nil {
					values = append(values, string(v))
				} else {
					// v is nil === not historical data with given ts
					values = append(values, "nil")
				}
			}
			if v, err := convertValues(dataTypes[i], values...); err != nil {
				return err
			} else {
				rawData.Set(name, []interface{}{timeStamps[i], v})
				return nil
			}
		})
	}
	if err := g.Wait(); err != nil {
		return nil, err
	} else {
		return rawData, nil
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
func (gdb *Gdb) GetHistoricalDataWithCondition(groupNames, itemNames []string, startTime []int, endTime []int, intervals []int, filterCondition string, zones ...DeadZone) (cmap.ConcurrentMap, error) {
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
	dataTypes := []string{}
	for index, itemName := range itemNames {
		if t, ok := gdb.rtDbFilter.Get(itemName + joiner + groupNames[index]); !ok {
			return nil, fmt.Errorf("item " + itemName + " not existed")
		} else {
			dataTypes = append(dataTypes, t.(string))
		}
	}
	filterHistoryData, err := gdb.getHistoricalDataWithMinLength(groupNames, filterItemNames, startTime, endTime, intervals, dataTypes...) // get history of filter item
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
		if data, err := gdb.getHistoricalDataWithStampAndDeadZoneCount(groupNames, itemNames, timeStamps, zones, dataTypes...); err != nil {
			return nil, err
		} else {
			return data, nil
		}
	} else {
		if data, err := gdb.getHistoricalDataWithStringTimeStamp(groupNames, itemNames, timeStamps, dataTypes...); err != nil {
			return nil, err
		} else {
			return data, nil
		}
	}
}

func (gdb *Gdb) getHistoricalDataWithStringTimeStamp(groupNames, itemNames []string, timeStamps []int64, dataTypes ...string) (cmap.ConcurrentMap, error) {
	rawData := cmap.New()
	sn, err := gdb.hisDb.GetSnapshot()
	if sn == nil || err != nil {
		return nil, snError{"snError"}
	}
	defer sn.Release() // release
	g := errgroup.Group{}
	for i, itemName := range itemNames {
		name, index := itemName, i
		g.Go(func() error {
			var values []string
			for j := 0; j < len(timeStamps); j++ {
				v, _ := sn.Get([]byte(name+joiner+groupNames[index]+fmt.Sprintf("%d", timeStamps[j])), nil)
				values = append(values, fmt.Sprintf("%s", v))
			}
			if v, err := convertValues(dataTypes[index], values...); err != nil {
				return err
			} else {
				rawData.Set(name, []interface{}{timeStamps, v})
				return nil
			}

		})
	}
	if err := g.Wait(); err != nil {
		return nil, err
	} else {
		return rawData, nil
	}
}

func (gdb *Gdb) getHistoricalDataWithStampAndDeadZoneCount(groupNames, itemNames []string, timeStamps []int64, zones []DeadZone, dataTypes ...string) (cmap.ConcurrentMap, error) {
	rawData := cmap.New()
	sn, err := gdb.hisDb.GetSnapshot()
	if sn == nil || err != nil {
		return nil, snError{"snError"}
	}
	defer sn.Release() // release
	g := errgroup.Group{}
	for j, itemName := range itemNames {
		name, index := itemName, j
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
					v, _ := sn.Get([]byte(name+joiner+groupNames[index]+fmt.Sprintf("%d", timeStamps[i])), nil)
					values = append(values, fmt.Sprintf("%s", v))
				}
				if v, err := convertValues(dataTypes[index], values...); err != nil {
					return err
				} else {
					rawData.Set(name, []interface{}{timeStamps, v})
					return nil
				}
			}
			deadZoneCounts := zones[filterIndex].DeadZoneCount
			if deadZoneCounts == 1 {
				rawData.Set(name, [][]string{{}, {}})
				return nil
			} else if deadZoneCounts < 1 {
				// don't filter
				for i := 0; i < len(timeStamps); i++ {
					v, _ := sn.Get([]byte(name+joiner+groupNames[index]+fmt.Sprintf("%d", timeStamps[i])), nil)
					values = append(values, fmt.Sprintf("%s", v))
				}
			} else {
				// filter
				for i := 0; i < len(timeStamps); i++ {
					v, _ := sn.Get([]byte(name+joiner+groupNames[index]+fmt.Sprintf("%d", timeStamps[i])), nil)
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
			if v, err := convertValues(dataTypes[index], values...); err != nil {
				return err
			} else {
				rawData.Set(name, []interface{}{timeStamps, v})
				return nil
			}
		})
	}
	if err := g.Wait(); err != nil {
		return nil, err
	} else {
		return rawData, nil
	}
}

func (gdb *Gdb) getHistoricalDataWithMinLength(groupNames, itemNames []string, startTime []int, endTime []int, intervals []int, dataTypes ...string) ([]map[string]interface{}, error) {
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
				startKey.Write([]byte(name + joiner + groupNames[j]))
				startKey.Write([]byte(strconv.Itoa(s)))
				endKey := strings.Builder{}
				endKey.Write([]byte(name + joiner + groupNames[j]))
				endKey.Write([]byte(strconv.Itoa(e)))
				it := sn.NewIterator(&util.Range{Start: []byte(startKey.String()), Limit: []byte(endKey.String())}, nil)
				count := 0
				var st int64 // start time stamp of data
				for it.Next() {
					tt, _ := strconv.ParseInt(strings.Replace(fmt.Sprintf("%s", it.Key()), name+joiner+groupNames[j], "", -1), 10, 64) // get time stamp
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
			if v, err := convertValues(dataTypes[j], values...); err != nil {
				return err
			} else {
				rawData.Set(name, []interface{}{timeStamps, v})
				lengthMap[j] = len(values)
				return nil
			}
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
				d := vs[1].([]interface{})[i]
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

func (gdb *Gdb) getNowTime() string {
	return time.Now().Format(timeFormatString)
}

func (gdb *Gdb) testItemValue(v interface{}) {
	if r, err := json.Marshal(v); err != nil {
		fmt.Println(v)
		fmt.Println(err)
	} else {
		fmt.Println(string(r))
	}
}

// following method is used by calc
// get unix timestamp of the given time,t should b yyyy-mm-dd hh:mm:ss
func (gdb *Gdb) getUnixTimeStamp(t string, d int) (int64, error) {
	if st, err := time.Parse(timeFormatString, t); err != nil {
		return -1, err
	} else {
		return st.Add(time.Duration(d) * time.Second).Unix(), nil
	}
}

func (gdb *Gdb) getRtData(itemNames, groupNames []string) (string, error) {
	if v, err := gdb.GetRealTimeData(groupNames, itemNames...); err != nil {
		return "", err
	} else {
		if r, err := json.Marshal(v); err != nil {
			return "", err
		} else {
			return string(r), err
		}
	}
}

func (gdb *Gdb) getHData(groupNames, itemNames []string, startTimeStamps, endTimeStamps, intervals []int) (string, error) {
	if result, err := gdb.GetHistoricalData(groupNames, itemNames, startTimeStamps, endTimeStamps, intervals); err != nil {
		return "", err
	} else {
		if r, err := json.Marshal(result); err != nil {
			return "", err
		} else {
			return string(r), nil
		}
	}
}

func (gdb *Gdb) getHDataWithTs(groupNames, itemNames []string, timeStamps ...[]int) (string, error) {
	if result, err := gdb.GetHistoricalDataWithStamp(groupNames, itemNames, timeStamps...); err != nil {
		return "", err
	} else {
		if r, err := json.Marshal(result); err != nil {
			return "", err
		} else {
			return string(r), nil
		}
	}
}

func (gdb *Gdb) writeRtData(infos []map[string]interface{}) (Rows, error) {
	items := []ItemValue{}
	for _, info := range infos {
		items = append(items, ItemValue{
			GroupName: info["groupName"].(string),
			ItemName:  info["itemName"].(string),
			Value:     info["value"],
		})
	}
	rows, err := gdb.BatchWrite(items...)
	return rows, err
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

func (gdb *Gdb) getDbInfoHistory(itemName string, startTimeStamps []int, endTimeStamps []int, intervals []int) (cmap.ConcurrentMap, error) {
	if len(startTimeStamps) == len(endTimeStamps) && len(startTimeStamps) == len(intervals) {
		rawData := cmap.New()
		sn, err := gdb.infoDb.GetSnapshot()
		if sn == nil || err != nil {
			return nil, snError{"snError"}
		}
		defer sn.Release() // release
		for i := 0; i < len(startTimeStamps); i++ {
			s := startTimeStamps[i] // startTime
			e := endTimeStamps[i]   // endTime
			interval := intervals[i]
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
	} else {
		return nil, fmt.Errorf("inconsistent length of starTimes, endTimes and intervals")
	}
}

func convertValues(t string, values ...string) ([]interface{}, error) {
	if values == nil {
		return nil, nil
	}
	r := make([]interface{}, len(values))
	g := errgroup.Group{}
	for i := 0; i < len(values); i++ {
		index := i
		if values[index] == "nil" {
			r[index] = nil
		} else {
			g.Go(func() error {
				switch t {
				case "int64":
					if result, err := strconv.ParseInt(values[index], 10, 64); err != nil {
						return err
					} else {
						r[index] = result
					}
					break
				case "float64":
					if result, err := strconv.ParseFloat(values[index], 64); err != nil {
						return err
					} else {
						r[index] = result
					}
					break
				case "bool":
					if result, err := strconv.ParseBool(values[index]); err != nil {
						return err
					} else {
						r[index] = result
					}
					break
				default:
					r[index] = values[index]
					break
				}
				return nil
			})
		}
	}
	if err := g.Wait(); err != nil {
		return r, err
	} else {
		return r, nil
	}
}
