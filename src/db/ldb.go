/*
creatTime: 2020/11/27
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.15.3
*/

package db

import (
	"errors"
	"fmt"
	. "github.com/ahmetb/go-linq"
	"github.com/dop251/goja"
	cmap "github.com/orcaman"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
	"utils"
)

/* Batch write real-time data
   kv : Key-value pair to be written
   withTimeStamp: Whether there is a time stamp when writing, if not, the default is the current time Unix timestamp
   if it is true, then v should be like [[ItemNames...],[Values...],[UnixTimeStamp... ]]
*/
func (ldb *LevelDb) BatchWrite(kv [][]string, withTimeStamp bool) (Rows, error) {
	itemNames := kv[0]  // itemNames
	itemValues := kv[1] // itemValues
	index, ok := ldb.checkItems(itemNames...)
	// check if all items given exist
	if !ok {
		return Rows{}, itemError{"itemError: " + itemNames[index]}
	}
	currentTimeStamp := int(time.Now().Unix()) // unix timestamp
	currentTimeStampString := strconv.Itoa(currentTimeStamp)
	wg := sync.WaitGroup{}
	messageChan := make(chan error, 2) // message chan
	wg.Add(3)
	// write currentTimeStamp
	go func() {
		defer wg.Done()
		_ = ldb.InfoDb.Put([]byte(TimeKey), []byte(currentTimeStampString), nil)
	}()
	// write Realtime data
	go func() {
		defer wg.Done()
		batch := leveldb.Batch{}
		for i := 0; i < len(itemNames); i++ {
			batch.Put([]byte(itemNames[i]), []byte(itemValues[i]))
		}
		if err := ldb.RtDb.Write(&batch, nil); err != nil {
			messageChan <- err
		} else {
			messageChan <- nil
		}
		return
	}()
	if withTimeStamp {
		// write historical data with timestamp
		timeStamps := kv[2]
		go func() {
			defer wg.Done()
			batch := leveldb.Batch{}
			for i := 0; i < len(itemNames); i++ {
				sb := strings.Builder{}
				sb.Write([]byte(itemNames[i]))  // itemName
				sb.Write([]byte(timeStamps[i])) // time stamp
				batch.Put([]byte(sb.String()), []byte(itemValues[i]))
			}
			if err := ldb.HisDb.Write(&batch, nil); err != nil {
				messageChan <- err
			} else {
				messageChan <- nil
			}
			return
		}()

	} else {
		// write historical data without timestamp
		go func() {
			defer wg.Done()
			batch := leveldb.Batch{}
			for i := 0; i < len(itemNames); i++ {
				sb := strings.Builder{}
				sb.Write([]byte(itemNames[i]))
				sb.Write([]byte(currentTimeStampString))
				batch.Put([]byte(sb.String()), []byte(itemValues[i]))
			}
			if err := ldb.HisDb.Write(&batch, nil); err != nil {
				messageChan <- err
			} else {
				messageChan <- nil
			}
			return
		}()
	}
	// monitor
	go func() {
		wg.Wait()
		close(messageChan)
	}()
	errorFlag := false
	errorMsg := ""
	for msg := range messageChan {
		if msg != nil {
			errorFlag = true
			errorMsg += msg.Error()
		}
	}
	if errorFlag {
		return Rows{}, errors.New(errorMsg)
	}
	return Rows{len(itemNames)}, nil
}

//  get realTime data
func (ldb *LevelDb) GetRealTimeData(itemNames ...string) (cmap.ConcurrentMap, error) {
	sn, err := ldb.RtDb.GetSnapshot()
	if sn == nil || err != nil {
		return nil, snError{"snError"}
	}
	defer sn.Release() // release
	m := cmap.New()
	wg := sync.WaitGroup{}
	for _, itemName := range itemNames {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			v, err := sn.Get([]byte(name), nil)
			if err != nil {
				//  key not exist
				m.Set(name, nil)
			} else {
				m.Set(name, fmt.Sprintf("%s", v)) // set values
			}
		}(itemName)
	}
	wg.Wait()
	return m, nil
}

// get historical data
func (ldb *LevelDb) GetHistoricalData(itemNames []string, startTimeStamps []int, endTimeStamps []int, intervals []int) (cmap.ConcurrentMap, error) {
	rawData := cmap.New()
	wg := sync.WaitGroup{}
	sn, err := ldb.HisDb.GetSnapshot()
	if sn == nil || err != nil {
		return nil, snError{"snError"}
	}
	defer sn.Release() // release
	latestTimeStampString, _ := ldb.InfoDb.Get([]byte(TimeKey), nil)
	latestTimeStamp, _ := strconv.ParseInt(fmt.Sprintf("%s", latestTimeStampString), 10, 0)
	for _, itemName := range itemNames {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
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
				if e > int(latestTimeStamp) {
					// startTime to currentTimeStamp
					endKey.Write([]byte(strconv.Itoa(int(latestTimeStamp))))
				} else {
					// startTime to endTime
					endKey.Write([]byte(strconv.Itoa(e)))
				}
				it := sn.NewIterator(&util.Range{Start: []byte(startKey.String()), Limit: []byte(endKey.String())}, nil)
				values := []string{}
				count := 0
				timeStamps := []string{}
				for it.Next() {
					if count%interval == 0 {
						values = append(values, fmt.Sprintf("%s", it.Value()))
						timeStamps = append(timeStamps, strings.Replace(fmt.Sprintf("%s", it.Key()), name, "", -1))
					}
					count++
				}
				rawData.Set(name, [][]string{timeStamps, values})
			}
		}(itemName)
	}
	wg.Wait()
	return rawData, nil
}

// get history data according to the given time stamps
func (ldb *LevelDb) GetHistoricalDataWithStamp(itemNames []string, timeStamps ...string) (cmap.ConcurrentMap, error) {
	rawData := cmap.New()
	wg := sync.WaitGroup{}
	sn, err := ldb.HisDb.GetSnapshot()
	if sn == nil || err != nil {
		return nil, snError{"snError"}
	}
	defer sn.Release() // release
	for _, itemName := range itemNames {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			values := []string{}
			for i := 0; i < len(timeStamps); i++ {
				v, _ := sn.Get([]byte(name+timeStamps[i]), nil)
				values = append(values, fmt.Sprintf("%s", v))
			}
			rawData.Set(name, [][]string{timeStamps, values})
		}(itemName)
	}
	wg.Wait()
	return rawData, nil
}

// filter condition must be correct js expression,itemName should be startedWith by item.
// eg: item["itemName1"]>10 && item["itemName2"] > 30 ....
// It should be noted that the entire judgment is based on the itemName with less historical value in the condition.
//If the longest itemName is used as the benchmark, we cannot make an accurate judgment on the AND logic in it.
//Just imagine the history of Item1 It is [3,4,5], and item2 is [10,11]. If item1 is used as the benchmark,
//we cannot determine how much other elements of item2 should be expanded, because the condition may have complicated
//logic about item1 and item2 And or logic, no matter what the number is expanded, there may be a judgment error.
func (ldb *LevelDb) GetHistoricalDataWithCondition(itemNames []string, startTime []int, endTime []int, intervals []int, filterCondition string, zones ...DeadZone) (cmap.ConcurrentMap, error) {
	filterItemNames := []string{} // filter item
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
	filterHistoryData, err := ldb.getHistoricalDataWithMinLength(filterItemNames, startTime, endTime, intervals) // get history of filter item
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
	timeStamps := []string{} // time stamp
	for _, fr := range filterResults {
		sfr := fr.(map[string]interface{})
		timeStamps = append(timeStamps, sfr["timeStamp"].(string))
	}
	if zones != nil && len(zones) != 0 {
		if data, err := ldb.getHistoricalDataWithStampAndDeadZoneCount(itemNames, timeStamps, zones...); err != nil {
			return nil, err
		} else {
			return data, nil
		}
	} else {
		if data, err := ldb.GetHistoricalDataWithStamp(itemNames, timeStamps...); err != nil {
			return nil, err
		} else {
			return data, nil
		}
	}
}

func (ldb *LevelDb) getHistoricalDataWithStampAndDeadZoneCount(itemNames []string, timeStamps []string, zones ...DeadZone) (cmap.ConcurrentMap, error) {
	rawData := cmap.New()
	wg := sync.WaitGroup{}
	sn, err := ldb.HisDb.GetSnapshot()
	if sn == nil || err != nil {
		return nil, snError{"snError"}
	}
	defer sn.Release() // release
	for index, itemName := range itemNames {
		wg.Add(1)
		go func(name string, j int) {
			defer wg.Done()
			values := []string{}
			var lastValue string
			lastValues := []string{}
			filterIndex := From(zones).IndexOf(func(item interface{}) bool {
				return item.(DeadZone).ItemName == name
			})
			if filterIndex == -1 {
				// don't filter
				for i := 0; i < len(timeStamps); i++ {
					v, _ := sn.Get([]byte(name+timeStamps[i]), nil)
					values = append(values, fmt.Sprintf("%s", v))
				}
				rawData.Set(name, [][]string{timeStamps, values})
				return
			}
			deadZoneCounts := zones[filterIndex].DeadZoneCount
			if deadZoneCounts == 1 {
				rawData.Set(name, [][]string{{}, {}})
				return
			} else if deadZoneCounts < 1 {
				// don't filter
				for i := 0; i < len(timeStamps); i++ {
					v, _ := sn.Get([]byte(name+timeStamps[i]), nil)
					values = append(values, fmt.Sprintf("%s", v))
				}
			} else {
				// filter
				for i := 0; i < len(timeStamps); i++ {
					v, _ := sn.Get([]byte(name+timeStamps[i]), nil)
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
			rawData.Set(name, [][]string{timeStamps, values})
			return
		}(itemName, index)
	}
	wg.Wait()
	return rawData, nil
}

func (ldb *LevelDb) getHistoricalDataWithMinLength(itemNames []string, startTime []int, endTime []int, intervals []int) ([]map[string]interface{}, error) {
	rawData := cmap.New()
	lengthMap := make([]int, len(itemNames))
	wg := sync.WaitGroup{}
	sn, err := ldb.HisDb.GetSnapshot()
	if sn == nil || err != nil {
		return nil, snError{"snError"}
	}
	defer sn.Release() // release
	latestTimeStampString, _ := ldb.InfoDb.Get([]byte(TimeKey), nil)
	latestTimeStamp, _ := strconv.ParseInt(fmt.Sprintf("%s", latestTimeStampString), 10, 0)
	for index, itemName := range itemNames {
		wg.Add(1)
		go func(name string, j int) {
			defer wg.Done()
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
				if e > int(latestTimeStamp) {
					// startTime to currentTimeStamp
					endKey.Write([]byte(strconv.Itoa(int(latestTimeStamp))))
				} else {
					// startTime to endTime
					endKey.Write([]byte(strconv.Itoa(e)))
				}
				it := sn.NewIterator(&util.Range{Start: []byte(startKey.String()), Limit: []byte(endKey.String())}, nil)
				values := []string{}
				count := 0
				timeStamps := []string{}
				for it.Next() {
					if count%interval == 0 {
						values = append(values, fmt.Sprintf("%s", it.Value()))
						timeStamps = append(timeStamps, strings.Replace(fmt.Sprintf("%s", it.Key()), name, "", -1))
					}
					count++
				}
				rawData.Set(name, [][]string{timeStamps, values})
				lengthMap[j] = len(values)
			}
		}(itemName, index)
	}
	wg.Wait()
	result := []map[string]interface{}{}
	minLength := From(lengthMap).Min().(int)
	for i := 0; i < minLength; i++ {
		t := map[string]interface{}{}
		for index, name := range itemNames {
			v, _ := rawData.Get(name)
			vs := v.([][]string)
			d, err := strconv.ParseFloat(vs[1][i], 64)
			if err != nil {
				break
			}
			t[name] = d // values
			if index == 0 {
				// first
				t["timeStamp"] = vs[0][i]
			} else {
				et := t["timeStamp"]
				if et != vs[0][i] {
					break // inconsistent timeStamp
				}
			}
		}
		result = append(result, t)
	}
	return result, nil
}

// check if all itemNames given exist
func (ldb *LevelDb) checkItems(itemNames ...string) (int, bool) {
	for index, itemName := range itemNames {
		_, ok := ldb.RtDbFilter.Get(itemName)
		if !ok {
			return index, false
		}
	}
	return -1, true
}

// following method is used by calc
// get unix timestamp of the given time,t should b yyyy-mm-dd hh:mm:ss
func (ldb *LevelDb) getUnixTimeStamp(t string) int64 {
	t1, err := time.Parse("2006-01-02 15:04:05", t)
	if err != nil {
		return -1
	}
	return t1.Unix()
}

func (ldb *LevelDb) getNowTime() string {
	return time.Now().Format(utils.TimeFormatString)
}

func (ldb *LevelDb) getTime(d int) string {
	return time.Now().Add(time.Duration(d) * time.Second).Format(utils.TimeFormatString)
}

func (ldb *LevelDb) getRtData(itemNames []string) ([]string, error) {
	v, err := ldb.GetRealTimeData(itemNames...)
	if err != nil {
		return nil, err
	}
	result := []string{}
	for _, name := range itemNames {
		rv, _ := v.Get(name)
		if rv == nil {
			return nil, fmt.Errorf("invalid itemName: " + name)
		}
		result = append(result, rv.(string))
	}
	return result, nil
}

func (ldb *LevelDb) getHData(itemNames []string, timeStamps []string) ([]string, error) {
	r, err := ldb.GetHistoricalDataWithStamp(itemNames, timeStamps...)
	result := []string{}
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

func (ldb *LevelDb) getDbInfo() (cmap.ConcurrentMap, error) {
	sn, err := ldb.InfoDb.GetSnapshot()
	itemNames := []string{Ram, WrittenItems, TimeKey, Speed}
	if sn == nil || err != nil {
		return nil, snError{"snError"}
	}
	defer sn.Release() // release
	m := cmap.New()
	wg := sync.WaitGroup{}
	for _, itemName := range itemNames {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			v, err := sn.Get([]byte(name), nil)
			if err != nil {
				//  key not exist
				m.Set(name, nil)
			} else {
				m.Set(name, fmt.Sprintf("%s", v)) // set values
			}
		}(itemName)
	}
	wg.Wait()
	return m, nil
}

func (ldb *LevelDb) getDbSpeedHistory(itemName string, startTimeStamps []int, endTimeStamps []int, interval int) (cmap.ConcurrentMap, error) {
	rawData := cmap.New()
	sn, err := ldb.InfoDb.GetSnapshot()
	if sn == nil || err != nil {
		return nil, snError{"snError"}
	}
	defer sn.Release() // release
	latestTimeStampString, _ := sn.Get([]byte(TimeKey), nil)
	latestTimeStamp, _ := strconv.ParseInt(fmt.Sprintf("%s", latestTimeStampString), 10, 0)
	for i := 0; i < len(startTimeStamps); i++ {
		s := startTimeStamps[i] // startTime
		e := endTimeStamps[i]   // endTime
		if s >= e {
			// startTime > endTime, continue
			continue
		}
		startKey := strings.Builder{}
		startKey.Write([]byte(itemName))
		startKey.Write([]byte(strconv.Itoa(s)))
		endKey := strings.Builder{}
		endKey.Write([]byte(itemName))
		if e > int(latestTimeStamp) {
			// startTime to currentTimeStamp
			endKey.Write([]byte(strconv.Itoa(int(latestTimeStamp))))
		} else {
			// startTime to endTime
			endKey.Write([]byte(strconv.Itoa(e)))
		}
		it := sn.NewIterator(&util.Range{Start: []byte(startKey.String()), Limit: []byte(endKey.String())}, nil)
		values := []string{}
		count := 0
		timeStamps := []string{}
		for it.Next() {
			if count%interval == 0 {
				values = append(values, fmt.Sprintf("%s", it.Value()))
				timeStamps = append(timeStamps, strings.Replace(fmt.Sprintf("%s", it.Key()), itemName, "", -1))
			}
			count++
		}
		rawData.Set(itemName, [][]string{timeStamps, values})
	}
	return rawData, nil
}
