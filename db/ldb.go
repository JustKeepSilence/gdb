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
	"github.com/JustKeepSilence/gdb/memap"
	pb "github.com/JustKeepSilence/gdb/model"
	. "github.com/ahmetb/go-linq/v3"
	"github.com/deckarep/golang-set"
	"github.com/dop251/goja"
	"github.com/golang/protobuf/proto"
	"github.com/shopspring/decimal"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/util"
	"golang.org/x/sync/errgroup"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

const syncBytes = "syncBytes"

// BatchWriteFloatData write float32 data to database, the dataType of written items MUST be float32, and all items MUST exist in gdb
// mapping relationship between parameters are groupName in groupNames corresponds to the itemName([]string) in itemNames at the same time
// corresponds to itemValue([]float32) in itemValues
func (gdb *Gdb) BatchWriteFloatData(groupNames []string, itemNames [][]string, itemValues [][]float32) (TimeRows, error) {
	st := time.Now()
	// check items
	if len(groupNames) == len(itemNames) && len(itemNames) == len(itemValues) {
		if names, err := gdb.checkItemDataType(groupNames, itemNames, "float32"); err != nil {
			return TimeRows{}, err
		} else {
			for i := 0; i < len(names); i++ {
				if len(itemNames[i]) != len(itemValues[i]) {
					return TimeRows{}, fmt.Errorf("inconsistent dataType")
				}
			}
			currentTimeStamp := int(time.Now().Unix()) + 8*3600
			g := errgroup.Group{}
			// write Realtime data
			g.Go(func() error {
				var keys, values [][]byte
				for i := 0; i < len(names); i++ {
					for j := 0; j < len(names[i]); j++ {
						keys = append(keys, convertStringToByte(names[i][j]))
						values = append(values, convertStringToByte(strconv.FormatFloat(float64(itemValues[i][j]), 'f', -1, 64)))
					}
				}
				if err := gdb.rtDb.BatchWrite(keys, values); err != nil {
					return err
				}
				return nil
			})
			g.Go(func() error {
				// write historical data
				for i := 0; i < len(names); i++ {
					// itemName = itemName + joiner + groupName
					for j := 0; j < len(names[i]); j++ {
						gdb.floatRmHisDb.Upsert(names[i][j], groupNames[i], memap.RmHisDbFloatItem{TimeStamp: int32(currentTimeStamp), Value: itemValues[i][j]})
					}
				}
				return nil
			})
			if err := g.Wait(); err != nil {
				return TimeRows{}, err
			} else {
				// write info
				d := time.Since(st).Milliseconds()
				counts := 0
				for i := 0; i < len(itemNames); i++ {
					counts += len(itemNames[i])
				}
				sd := fmt.Sprintf("%dms/%d", d, counts)
				infoBatch := &leveldb.Batch{}
				infoBatch.Put(convertStringToByte(speed), convertStringToByte(sd))                                             // speed
				infoBatch.Put(convertStringToByte(speed+strconv.Itoa(int(time.Now().Unix()+8*3600))), convertStringToByte(sd)) // history data of speed
				if err := gdb.floatInfoDb.Write(infoBatch, nil); err != nil {
					return TimeRows{}, err
				}
				return TimeRows{counts, d}, nil
			}
		}
	} else {
		return TimeRows{}, fmt.Errorf("inconsistent length of parameters")
	}
}

// BatchWriteIntData write int32 data to database, the dataType of written items MUST be int32 and all items MUST exist in gdb
// mapping relationship between parameters are groupName in groupNames corresponds to the itemName([]string) in itemNames at the same time
// corresponds to itemValue([]int32) in itemValues
func (gdb *Gdb) BatchWriteIntData(groupNames []string, itemNames [][]string, itemValues [][]int32) (TimeRows, error) {
	st := time.Now()
	// check items
	if len(groupNames) == len(itemNames) && len(itemNames) == len(itemValues) {
		if names, err := gdb.checkItemDataType(groupNames, itemNames, "int32"); err != nil {
			return TimeRows{}, err
		} else {
			for i := 0; i < len(names); i++ {
				if len(itemNames[i]) != len(itemValues[i]) {
					return TimeRows{}, fmt.Errorf("inconsistent dataType")
				}
			}
			currentTimeStamp := int(time.Now().Unix()) + 8*3600
			g := errgroup.Group{}
			// write Realtime data
			g.Go(func() error {
				var keys, values [][]byte
				for i := 0; i < len(names); i++ {
					for j := 0; j < len(names[i]); j++ {
						keys = append(keys, convertStringToByte(names[i][j]))
						values = append(values, convertStringToByte(strconv.Itoa(int(itemValues[i][j]))))
					}
				}
				if err := gdb.rtDb.BatchWrite(keys, values); err != nil {
					return err
				}
				return nil
			})
			g.Go(func() error {
				// write historical data
				for i := 0; i < len(names); i++ {
					// itemName = itemName + joiner + groupName
					for j := 0; j < len(names[i]); j++ {
						gdb.intRmHisDb.Upsert(names[i][j], groupNames[i], memap.RmHisDbIntItem{TimeStamp: int32(currentTimeStamp), Value: itemValues[i][j]})
					}
				}
				return nil
			})
			// write historical data
			if err := g.Wait(); err != nil {
				return TimeRows{}, err
			} else {
				// write info
				d := time.Since(st).Milliseconds()
				counts := 0
				for i := 0; i < len(itemNames); i++ {
					counts += len(itemNames[i])
				}
				sd := fmt.Sprintf("%dms/%d", d, counts)
				infoBatch := &leveldb.Batch{}
				infoBatch.Put(convertStringToByte(speed), convertStringToByte(sd))                                             // speed
				infoBatch.Put(convertStringToByte(speed+strconv.Itoa(int(time.Now().Unix()+8*3600))), convertStringToByte(sd)) // history data of speed
				if err := gdb.intInfoDb.Write(infoBatch, nil); err != nil {
					return TimeRows{}, err
				}
				return TimeRows{counts, d}, nil
			}
		}
	} else {
		return TimeRows{}, fmt.Errorf("inconsistent length of parameters")
	}
}

// BatchWriteStringData write string data to database, the dataType of written items MUST be string and all items MUST exist in gdb
// mapping relationship between parameters are groupName in groupNames corresponds to the itemName([]string) in itemNames at the same time
// corresponds to itemValue([]string) in itemValues
func (gdb *Gdb) BatchWriteStringData(groupNames []string, itemNames [][]string, itemValues [][]string) (TimeRows, error) {
	st := time.Now()
	// check items
	if len(groupNames) == len(itemNames) && len(itemNames) == len(itemValues) {
		if names, err := gdb.checkItemDataType(groupNames, itemNames, "string"); err != nil {
			return TimeRows{}, err
		} else {
			for i := 0; i < len(names); i++ {
				if len(itemNames[i]) != len(itemValues[i]) {
					return TimeRows{}, fmt.Errorf("inconsistent dataType")
				}
			}
			currentTimeStamp := int(time.Now().Unix()) + 8*3600
			g := errgroup.Group{}
			// write Realtime data
			g.Go(func() error {
				var keys, values [][]byte
				for i := 0; i < len(names); i++ {
					for j := 0; j < len(names[i]); j++ {
						keys = append(keys, convertStringToByte(names[i][j]))
						values = append(values, convertStringToByte(itemValues[i][j]))
					}
				}
				if err := gdb.rtDb.BatchWrite(keys, values); err != nil {
					return err
				}
				return nil
			})
			g.Go(func() error {
				// write historical data
				for i := 0; i < len(names); i++ {
					// itemName = itemName + joiner + groupName
					for j := 0; j < len(names[i]); j++ {
						gdb.stringRmHisDb.Upsert(names[i][j], groupNames[i], memap.RmHisDbStringItem{TimeStamp: int32(currentTimeStamp), Value: itemValues[i][j]})
					}
				}
				return nil
			})
			// write historical data
			if err := g.Wait(); err != nil {
				return TimeRows{}, err
			} else {
				// write info
				d := time.Since(st).Milliseconds()
				counts := 0
				for i := 0; i < len(itemNames); i++ {
					counts += len(itemNames[i])
				}
				sd := fmt.Sprintf("%dms/%d", d, counts)
				infoBatch := &leveldb.Batch{}
				infoBatch.Put(convertStringToByte(speed), convertStringToByte(sd))                                             // speed
				infoBatch.Put(convertStringToByte(speed+strconv.Itoa(int(time.Now().Unix()+8*3600))), convertStringToByte(sd)) // history data of speed
				if err := gdb.stringInfoDb.Write(infoBatch, nil); err != nil {
					return TimeRows{}, err
				}
				return TimeRows{counts, d}, nil
			}
		}
	} else {
		return TimeRows{}, fmt.Errorf("inconsistent length of parameters")
	}
}

// BatchWriteBoolData write bool data to database, the dataType of written items MUST be bool and all items MUST exist in gdb
// mapping relationship between parameters are groupName in groupNames corresponds to the itemName([]string) in itemNames at the same time
// corresponds to itemValue([]bool) in itemValues
func (gdb *Gdb) BatchWriteBoolData(groupNames []string, itemNames [][]string, itemValues [][]bool) (TimeRows, error) {
	st := time.Now()
	// check items
	if len(groupNames) == len(itemNames) && len(itemNames) == len(itemValues) {
		if names, err := gdb.checkItemDataType(groupNames, itemNames, "bool"); err != nil {
			return TimeRows{}, err
		} else {
			for i := 0; i < len(names); i++ {
				if len(itemNames[i]) != len(itemValues[i]) {
					return TimeRows{}, fmt.Errorf("inconsistent dataType")
				}
			}
			currentTimeStamp := int(time.Now().Unix()) + 8*3600
			g := errgroup.Group{}
			// write Realtime data
			g.Go(func() error {
				var keys, values [][]byte
				for i := 0; i < len(names); i++ {
					for j := 0; j < len(names[i]); j++ {
						keys = append(keys, convertStringToByte(names[i][j]))
						values = append(values, convertStringToByte(strconv.FormatBool(itemValues[i][j])))
					}
				}
				if err := gdb.rtDb.BatchWrite(keys, values); err != nil {
					return err
				}
				return nil
			})
			g.Go(func() error {
				// write historical data
				for i := 0; i < len(names); i++ {
					// itemName = itemName + joiner + groupName
					for j := 0; j < len(names[i]); j++ {
						gdb.boolRmHisDb.Upsert(names[i][j], groupNames[i], memap.RmHisDbBoolItem{TimeStamp: int32(currentTimeStamp), Value: itemValues[i][j]})
					}
				}
				return nil
			})
			// write historical data
			if err := g.Wait(); err != nil {
				return TimeRows{}, err
			} else {
				// write info
				d := time.Since(st).Milliseconds()
				counts := 0
				for i := 0; i < len(itemNames); i++ {
					counts += len(itemNames[i])
				}
				sd := fmt.Sprintf("%dms/%d", d, counts)
				infoBatch := &leveldb.Batch{}
				infoBatch.Put(convertStringToByte(speed), convertStringToByte(sd))                                             // speed
				infoBatch.Put(convertStringToByte(speed+strconv.Itoa(int(time.Now().Unix()+8*3600))), convertStringToByte(sd)) // history data of speed
				if err := gdb.boolInfoDb.Write(infoBatch, nil); err != nil {
					return TimeRows{}, err
				}
				return TimeRows{counts, d}, nil
			}
		}
	} else {
		return TimeRows{}, fmt.Errorf("inconsistent length of parameters")
	}
}

// BatchWriteFloatHistoricalData write float32 historicalData to database, all items to be written MUST exist in gdb,
// and the timeStamp MUST be unix timeStamp, data MUST be sorted by timestamp in ascending order
// mapping relationship between parameters are the itemName in itemNames corresponds to the groupName in groupNames at the same time corresponds to timeStamp([]int32)
// in timeStamps and itemValue ([]float32) in itemValues
func (gdb *Gdb) BatchWriteFloatHistoricalData(groupNames []string, itemNames []string, timeStamps [][]int32, itemValues [][]float32) (TimeRows, error) {
	// check items
	st := time.Now()
	if len(itemNames) == len(timeStamps) && len(timeStamps) == len(itemValues) {
		if _, namesMap, err := gdb.checkSingleItemDataTypeWithMap(groupNames, itemNames, "float32"); err != nil {
			return TimeRows{}, err
		} else {
			for i := 0; i < len(timeStamps); i++ {
				if len(timeStamps[i]) != len(itemValues[i]) {
					return TimeRows{}, fmt.Errorf("inconsistent length of parameters")
				}
			}
			count := 0
			for groupName, groupItemNames := range namesMap {
				batch := &leveldb.Batch{}
				sn, err := gdb.hisDb["float32"][groupName].GetSnapshot()
				{
					if err != nil {
						return TimeRows{}, err
					}
					for index, name := range groupItemNames {
						if err := gdb.checkTimeStampsInDb(name, int(timeStamps[index][0]), int(timeStamps[index][len(timeStamps[index])-1]), sn); err != nil {
							return TimeRows{}, err
						}
						ts, values, stt, tts := timeStamps[index], []float32{}, timeStamps[index][0], []int32{}
						for i := 0; i < len(ts); i++ {
							if ts[i]-stt < int32(gdb.hisTimeDuration.Seconds()) && i != len(ts)-1 {
								values = append(values, itemValues[index][i])
								tts = append(tts, ts[i])
							} else {
								key := name + joiner + strconv.Itoa(int(stt))
								if i == len(ts)-1 {
									// add last one
									values = append(values, itemValues[index][i])
									tts = append(tts, ts[i])
								}
								m := &pb.FloatHistoricalData{TimeStamps: tts, Values: values}
								if data, err := proto.Marshal(m); err != nil {
									return TimeRows{}, err
								} else {
									batch.Put(convertStringToByte(key), data)
									values = []float32{itemValues[index][i]}
									tts = []int32{ts[i]}
									stt = ts[i]
								}
							}
						}
						count += len(timeStamps[index])
					}
					sn.Release()
				}
				if err := gdb.hisDb["float32"][groupName].Write(batch, nil); err != nil {
					return TimeRows{}, err
				}
			}
			return TimeRows{EffectedRows: count, Times: time.Now().Sub(st).Milliseconds()}, nil
		}
	} else {
		return TimeRows{}, fmt.Errorf("inconsistent length of parameters")
	}
}

// BatchWriteIntHistoricalData write float32 historicalData to database, all items to be written MUST exist in gdb,
// and the timeStamp MUST be unix timeStamp, data MUST be sorted by timestamp in ascending order
// mapping relationship between parameters are the itemName in itemNames corresponds to the groupName in groupNames at the same time corresponds to timeStamp([]int32)
// in timeStamps and itemValue ([]int32) in itemValues
func (gdb *Gdb) BatchWriteIntHistoricalData(groupNames []string, itemNames []string, timeStamps [][]int32, itemValues [][]int32) (TimeRows, error) {
	// check items
	st := time.Now()
	if len(groupNames) == len(itemNames) && len(itemNames) == len(timeStamps) && len(timeStamps) == len(itemValues) {
		if _, namesMap, err := gdb.checkSingleItemDataTypeWithMap(groupNames, itemNames, "int32"); err != nil {
			return TimeRows{}, err
		} else {
			for i := 0; i < len(timeStamps); i++ {
				if len(timeStamps[i]) != len(itemValues[i]) {
					return TimeRows{}, fmt.Errorf("inconsistent length of parameters")
				}
			}
			count := 0
			for groupName, groupItemNames := range namesMap {
				batch := &leveldb.Batch{}
				sn, err := gdb.hisDb["int32"][groupName].GetSnapshot()
				{
					if err != nil {
						return TimeRows{}, err
					}
					for index, name := range groupItemNames {
						if err := gdb.checkTimeStampsInDb(name, int(timeStamps[index][0]), int(timeStamps[index][len(timeStamps[index])-1]), sn); err != nil {
							return TimeRows{}, err
						}
						ts, values, stt, tts := timeStamps[index], []int32{}, timeStamps[index][0], []int32{}
						for i := 0; i < len(ts); i++ {
							if ts[i]-stt < int32(gdb.hisTimeDuration.Seconds()) && i != len(ts)-1 {
								values = append(values, itemValues[index][i])
								tts = append(tts, ts[i])
							} else {
								key := name + joiner + strconv.Itoa(int(stt))
								if i == len(ts)-1 {
									// add last one
									values = append(values, itemValues[index][i])
									tts = append(tts, ts[i])
								}
								m := &pb.IntHistoricalData{TimeStamps: tts, Values: values}
								if data, err := proto.Marshal(m); err != nil {
									return TimeRows{}, err
								} else {
									batch.Put(convertStringToByte(key), data)
									values = []int32{itemValues[index][i]}
									tts = []int32{ts[i]}
									stt = ts[i]
								}
							}
						}
						count += len(ts)
					}
					sn.Release()
				}
				if err := gdb.hisDb["int32"][groupName].Write(batch, nil); err != nil {
				}
				return TimeRows{}, err
			}
			return TimeRows{EffectedRows: count, Times: time.Now().Sub(st).Milliseconds()}, nil
		}
	} else {
		return TimeRows{}, fmt.Errorf("inconsistent length of parameters")
	}
}

// BatchWriteStringHistoricalData write float32 historicalData to database, all items to be written MUST exist in gdb,
// and the timeStamp MUST be unix timeStamp, data MUST be sorted by timestamp in ascending order
// mapping relationship between parameters are the itemName in itemNames corresponds to the groupName in groupNames at the same time corresponds to timeStamp([]int32)
// in timeStamps and itemValue ([]string) in itemValues
func (gdb *Gdb) BatchWriteStringHistoricalData(groupNames []string, itemNames []string, timeStamps [][]int32, itemValues [][]string) (TimeRows, error) {
	// check items
	st := time.Now()
	if len(groupNames) == len(itemNames) && len(itemNames) == len(timeStamps) && len(timeStamps) == len(itemValues) {
		if _, namesMap, err := gdb.checkSingleItemDataTypeWithMap(groupNames, itemNames, "string"); err != nil {
			return TimeRows{}, err
		} else {
			for i := 0; i < len(timeStamps); i++ {
				if len(timeStamps[i]) != len(itemValues[i]) {
					return TimeRows{}, fmt.Errorf("inconsistent length of parameters")
				}
			}
			count := 0
			for groupName, groupItemNames := range namesMap {
				batch := &leveldb.Batch{}
				sn, err := gdb.hisDb["string"][groupName].GetSnapshot()
				{
					if err != nil {
						return TimeRows{}, err
					}
					for index, name := range groupItemNames {
						if err := gdb.checkTimeStampsInDb(name, int(timeStamps[index][0]), int(timeStamps[index][len(timeStamps[index])-1]), sn); err != nil {
							return TimeRows{}, err
						}
						ts, values, stt, tts := timeStamps[index], []string{}, timeStamps[index][0], []int32{}
						for i := 0; i < len(ts); i++ {
							if ts[i]-stt < int32(gdb.hisTimeDuration.Seconds()) && i != len(ts)-1 {
								values = append(values, itemValues[index][i])
								tts = append(tts, ts[i])
							} else {
								key := name + joiner + strconv.Itoa(int(stt))
								if i == len(ts)-1 {
									// add last one
									values = append(values, itemValues[index][i])
									tts = append(tts, ts[i])
								}
								m := &pb.StringHistoricalData{TimeStamps: tts, Values: values}
								if data, err := proto.Marshal(m); err != nil {
									return TimeRows{}, err
								} else {
									batch.Put(convertStringToByte(key), data)
									values = []string{itemValues[index][i]}
									tts = []int32{ts[i]}
									stt = ts[i]
								}
							}
						}
						count += len(ts)
					}
					sn.Release()
				}
				if err := gdb.hisDb["string"][groupName].Write(batch, nil); err != nil {
					return TimeRows{}, err
				}
			}
			return TimeRows{EffectedRows: count, Times: time.Now().Sub(st).Milliseconds()}, nil
		}
	} else {
		return TimeRows{}, fmt.Errorf("inconsistent length of parameters")
	}
}

// BatchWriteBoolHistoricalData write float32 historicalData to database, all items to be written MUST exist in gdb,
// and the timeStamp MUST be unix timeStamp, data MUST be sorted by timestamp in ascending order
// mapping relationship between parameters are the itemName in itemNames corresponds to the groupName in groupNames at the same time corresponds to timeStamp([]int32)
// in timeStamps and itemValue ([]bool) in itemValues
func (gdb *Gdb) BatchWriteBoolHistoricalData(groupNames []string, itemNames []string, timeStamps [][]int32, itemValues [][]bool) (TimeRows, error) {
	// check items
	st := time.Now()
	if len(groupNames) == len(itemNames) && len(itemNames) == len(timeStamps) && len(timeStamps) == len(itemValues) {
		if _, namesMap, err := gdb.checkSingleItemDataTypeWithMap(groupNames, itemNames, "bool"); err != nil {
			return TimeRows{}, err
		} else {
			for i := 0; i < len(timeStamps); i++ {
				if len(timeStamps[i]) != len(itemValues[i]) {
					return TimeRows{}, fmt.Errorf("inconsistent length of parameters")
				}
			}
			count := 0
			for groupName, groupItemNames := range namesMap {
				batch := &leveldb.Batch{}
				sn, err := gdb.hisDb["bool"][groupName].GetSnapshot()
				{
					if err != nil {
						return TimeRows{}, err
					}
					for index, name := range groupItemNames {
						if err := gdb.checkTimeStampsInDb(name, int(timeStamps[index][0]), int(timeStamps[index][len(timeStamps[index])-1]), sn); err != nil {
							return TimeRows{}, err
						}
						ts, values, stt, tts := timeStamps[index], []bool{}, timeStamps[index][0], []int32{}
						for i := 0; i < len(ts); i++ {
							if ts[i]-stt < int32(gdb.hisTimeDuration.Seconds()) && i != len(ts)-1 {
								values = append(values, itemValues[index][i])
								tts = append(tts, ts[i])
							} else {
								key := name + joiner + strconv.Itoa(int(stt))
								if i == len(ts)-1 {
									// add last one
									values = append(values, itemValues[index][i])
									tts = append(tts, ts[i])
								}
								m := &pb.BoolHistoricalData{TimeStamps: tts, Values: values}
								if data, err := proto.Marshal(m); err != nil {
									return TimeRows{}, err
								} else {
									batch.Put(convertStringToByte(key), data)
									values = []bool{itemValues[index][i]}
									tts = []int32{ts[i]}
									stt = ts[i]
								}
							}
						}
						count += len(ts)
					}
					sn.Release()
				}
				if err := gdb.hisDb["bool"][groupName].Write(batch, nil); err != nil {
					return TimeRows{}, err
				}
			}
			return TimeRows{EffectedRows: count, Times: time.Now().Sub(st).Milliseconds()}, nil
		}
	} else {
		return TimeRows{}, fmt.Errorf("inconsistent length of parameters")
	}
}

// GetRealTimeData get realTime data,that is the latest updated value of item.All items should be
// existed in gdb, otherWise will fail to getting data.If realTime of item not exist, the corresponding
// value will be nil
//
// The format of return value is {"itemName1": value1, "itemName2": value2}, values corresponds to the dataType of item
func (gdb *Gdb) GetRealTimeData(groupNames, itemNames []string) (GdbRealTimeData, error) {
	if len(groupNames) == len(itemNames) {
		st := time.Now()
		dataTypes := []string{}
		keys := [][]byte{}
		for index, itemName := range itemNames {
			if t, ok := gdb.rtDbFilter.Get(itemName + joiner + groupNames[index]); !ok {
				return GdbRealTimeData{}, fmt.Errorf("item " + itemName + " not existed")
			} else {
				dataTypes = append(dataTypes, t.(string))
				keys = append(keys, convertStringToByte(itemName+joiner+groupNames[index]))
			}
		}
		m := cmap.New()
		if values, err := gdb.rtDb.BatchFetch(keys); err != nil {
			return GdbRealTimeData{}, err
		} else {
			g := errgroup.Group{}
			for index, v1 := range values {
				i := index
				v, name := v1, itemNames[i]
				g.Go(func() error {
					if v == nil {
						//  realTime data not existed
						m.Set(name, nil)
					} else {
						switch dataTypes[i] {
						case "int32":
							if tv, err := strconv.ParseInt(string(v), 10, 64); err != nil {
								return err
							} else {
								m.Set(name, tv)
							}
							break
						case "float32":
							if d, err := decimal.NewFromString(string(v)); err != nil {
								return err
							} else {
								tv, _ := d.Float64()
								m.Set(name, float32(tv)) // set values
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
				return GdbRealTimeData{}, err
			} else {
				return GdbRealTimeData{RealTimeData: m, Times: time.Now().Sub(st).Milliseconds()}, nil
			}
		}
	} else {
		return GdbRealTimeData{}, fmt.Errorf("inconsistent length of parameters")
	}
}

// GetFloatHistoricalData will get float32 history data from database, the startTimes, endTimes should be unix timeStamp
// and unit of interval is seconds, interval MUST >= 1.If history data of item not exist,the corresponding
// value will be nil.All items MUST be exist in database.
//
// The returned format of values is {"itemName1": [timeStamp([]int32), values([]float32)]},
// if history of item not exist, the timeStamp and values is nil
func (gdb *Gdb) GetFloatHistoricalData(groupNames, itemNames []string, startTimes, endTimes, intervals []int32) (GdbHistoricalData, error) {
	st := time.Now()
	if len(startTimes) == len(endTimes) && len(endTimes) == len(intervals) && len(groupNames) == len(itemNames) {
		// checkItems
		if rawNamesMap, namesMap, err := gdb.checkSingleItemDataTypeWithMap(groupNames, itemNames, "float32"); err != nil {
			return GdbHistoricalData{}, err
		} else {
			if m, err := gdb.getHistoricalData("float32", rawNamesMap, namesMap, startTimes, endTimes, intervals); err != nil {
				return GdbHistoricalData{}, err
			} else {
				return GdbHistoricalData{HistoricalData: m, Times: time.Now().Sub(st).Milliseconds()}, nil
			}
		}
	} else {
		return GdbHistoricalData{}, fmt.Errorf("inconsistent length of parameters")
	}
}

// GetIntHistoricalData will get int32 history data from database, the startTimes, endTimes should be unix timeStamp
// and unit of interval is seconds, interval MUST >= 1.If history data of item not exist,the corresponding
// value will be nil.All items MUST be exist in database.
//
// The returned format of values is {"itemName1": [timeStamp([]int32), values([]int32)]},
// if history of item not exist, the timeStamp and values is nil
func (gdb *Gdb) GetIntHistoricalData(groupNames, itemNames []string, startTimes, endTimes, intervals []int32) (GdbHistoricalData, error) {
	st := time.Now()
	if len(startTimes) == len(endTimes) && len(endTimes) == len(intervals) && len(groupNames) == len(itemNames) {
		// checkItems
		if rawNamesMap, namesMap, err := gdb.checkSingleItemDataTypeWithMap(groupNames, itemNames, "int32"); err != nil {
			return GdbHistoricalData{}, err
		} else {
			if m, err := gdb.getHistoricalData("int32", rawNamesMap, namesMap, startTimes, endTimes, intervals); err != nil {
				return GdbHistoricalData{}, err
			} else {
				return GdbHistoricalData{HistoricalData: m, Times: time.Now().Sub(st).Milliseconds()}, nil
			}
		}
	} else {
		return GdbHistoricalData{}, fmt.Errorf("inconsistent length of parameters")
	}
}

// GetStringHistoricalData will get string history data from database, the startTimes, endTimes should be unix timeStamp
// and unit of interval is seconds, interval MUST >= 1.If history data of item not exist,the corresponding
// value will be nil.All items MUST be exist in database.
//
// The returned format of values is {"itemName1": [timeStamp([]int32), values([]string)]},
// if history of item not exist, the timeStamp and values is nil
func (gdb *Gdb) GetStringHistoricalData(groupNames, itemNames []string, startTimes, endTimes, intervals []int32) (GdbHistoricalData, error) {
	st := time.Now()
	if len(startTimes) == len(endTimes) && len(endTimes) == len(intervals) && len(groupNames) == len(itemNames) {
		// checkItems
		if rawNamesMap, namesMap, err := gdb.checkSingleItemDataTypeWithMap(groupNames, itemNames, "string"); err != nil {
			return GdbHistoricalData{}, err
		} else {
			if m, err := gdb.getHistoricalData("string", rawNamesMap, namesMap, startTimes, endTimes, intervals); err != nil {
				return GdbHistoricalData{}, err
			} else {
				return GdbHistoricalData{HistoricalData: m, Times: time.Now().Sub(st).Milliseconds()}, nil
			}
		}
	} else {
		return GdbHistoricalData{}, fmt.Errorf("inconsistent length of parameters")
	}
}

// GetBoolHistoricalData will get bool history data from database, the startTimes, endTimes should be unix timeStamp
// and unit of interval is seconds, interval MUST >= 1.If history data of item not exist,the corresponding
// value will be nil.All items MUST be exist in database.
//
// The returned format of values is {"itemName1": [timeStamp([]int32), values([]bool)]},
// if history of item not exist, the timeStamp and values is nil
func (gdb *Gdb) GetBoolHistoricalData(groupNames, itemNames []string, startTimes, endTimes, intervals []int32) (GdbHistoricalData, error) {
	st := time.Now()
	if len(startTimes) == len(endTimes) && len(endTimes) == len(intervals) && len(groupNames) == len(itemNames) {
		// checkItems
		if rawNamesMap, namesMap, err := gdb.checkSingleItemDataTypeWithMap(groupNames, itemNames, "bool"); err != nil {
			return GdbHistoricalData{}, err
		} else {
			if m, err := gdb.getHistoricalData("bool", rawNamesMap, namesMap, startTimes, endTimes, intervals); err != nil {
				return GdbHistoricalData{}, err
			} else {
				return GdbHistoricalData{HistoricalData: m, Times: time.Now().Sub(st).Milliseconds()}, nil
			}
		}
	} else {
		return GdbHistoricalData{}, fmt.Errorf("inconsistent length of parameters")
	}
}

// GetFloatRawHistoricalData will get all float32 history data of the given items.All items MUST be exist in database.
//
// The returned format of values is {"itemName1": [timeStamp([]int32), values([]float32)]},
// if history of item not exist, the timeStamp and values is nil
func (gdb *Gdb) GetFloatRawHistoricalData(groupNames, itemNames []string) (GdbHistoricalData, error) {
	st := time.Now()
	if len(groupNames) == len(itemNames) {
		if rawNamesMap, namesMap, err := gdb.checkSingleItemDataTypeWithMap(groupNames, itemNames, "float32"); err != nil {
			return GdbHistoricalData{}, err
		} else {
			if m, err := gdb.getRawHistoricalData("float32", rawNamesMap, namesMap); err != nil {
				return GdbHistoricalData{}, err
			} else {
				return GdbHistoricalData{HistoricalData: m, Times: time.Now().Sub(st).Milliseconds()}, nil
			}
		}
	} else {
		return GdbHistoricalData{}, fmt.Errorf("inconsistent length of parameters")
	}
}

// GetIntRawHistoricalData will get all int32 history data of the given items.All items MUST be exist in database.
//
// The returned format of values is {"itemName1": [timeStamp([]int32), values([]int32)]},
// if history of item not exist, the timeStamp and values is nil
func (gdb *Gdb) GetIntRawHistoricalData(groupNames, itemNames []string) (GdbHistoricalData, error) {
	st := time.Now()
	if len(groupNames) == len(itemNames) {
		if rawNamesMap, namesMap, err := gdb.checkSingleItemDataTypeWithMap(groupNames, itemNames, "int32"); err != nil {
			return GdbHistoricalData{}, err
		} else {
			if m, err := gdb.getRawHistoricalData("int32", rawNamesMap, namesMap); err != nil {
				return GdbHistoricalData{}, err
			} else {
				return GdbHistoricalData{HistoricalData: m, Times: time.Now().Sub(st).Milliseconds()}, nil
			}
		}
	} else {
		return GdbHistoricalData{}, fmt.Errorf("inconsistent length of parameters")
	}
}

// GetStringRawHistoricalData will get all string history data of the given items.All items MUST be exist in database.
//
// The returned format of values is {"itemName1": [timeStamp([]int32), values([]string)]},
// if history of item not exist, the timeStamp and values is nil
func (gdb *Gdb) GetStringRawHistoricalData(groupNames, itemNames []string) (GdbHistoricalData, error) {
	st := time.Now()
	if len(groupNames) == len(itemNames) {
		if rawNamesMap, namesMap, err := gdb.checkSingleItemDataTypeWithMap(groupNames, itemNames, "string"); err != nil {
			return GdbHistoricalData{}, err
		} else {
			if m, err := gdb.getRawHistoricalData("string", rawNamesMap, namesMap); err != nil {
				return GdbHistoricalData{}, err
			} else {
				return GdbHistoricalData{HistoricalData: m, Times: time.Now().Sub(st).Milliseconds()}, nil
			}
		}
	} else {
		return GdbHistoricalData{}, fmt.Errorf("inconsistent length of parameters")
	}
}

// GetBoolRawHistoricalData will get all bool history data of the given items.All items MUST be exist in database.
//
// The returned format of values is {"itemName1": [timeStamp([]int32), values([]bool)]},
// if history of item not exist, the timeStamp and values is nil
func (gdb *Gdb) GetBoolRawHistoricalData(groupNames, itemNames []string) (GdbHistoricalData, error) {
	st := time.Now()
	if len(groupNames) == len(itemNames) {
		if rawNamesMap, namesMap, err := gdb.checkSingleItemDataTypeWithMap(groupNames, itemNames, "bool"); err != nil {
			return GdbHistoricalData{}, err
		} else {
			if m, err := gdb.getRawHistoricalData("bool", rawNamesMap, namesMap); err != nil {
				return GdbHistoricalData{}, err
			} else {
				return GdbHistoricalData{HistoricalData: m, Times: time.Now().Sub(st).Milliseconds()}, nil
			}
		}
	} else {
		return GdbHistoricalData{}, fmt.Errorf("inconsistent length of parameters")
	}
}

// GetFloatHistoricalDataWithStamp get float32 history data according to the given time stamps, if history corresponding not exist
// we will not returns value
//
// The returned format of values is {"itemName1": [timeStamp([]int32), values([]float32)]},
// if all history of ts not exist, the timeStamp and value slice will be nil
func (gdb *Gdb) GetFloatHistoricalDataWithStamp(groupNames, itemNames []string, timeStamps [][]int32) (GdbHistoricalData, error) {
	st := time.Now()
	if len(groupNames) == len(itemNames) && len(itemNames) == len(timeStamps) {
		if rawNamesMap, namesMap, err := gdb.checkSingleItemDataTypeWithMap(groupNames, itemNames, "float32"); err != nil {
			return GdbHistoricalData{}, err
		} else {
			if m, err := gdb.getHistoricalDataWithTs("float32", rawNamesMap, namesMap, timeStamps...); err != nil {
				return GdbHistoricalData{}, err
			} else {
				return GdbHistoricalData{HistoricalData: m, Times: time.Now().Sub(st).Milliseconds()}, nil
			}
		}
	} else {
		return GdbHistoricalData{}, fmt.Errorf("inconsistent length of parameters")
	}
}

// GetIntHistoricalDataWithStamp get int32 history data according to the given time stamps, if history corresponding not exist
// we will not returns value
//
// The returned format of values is {"itemName1": [timeStamp([]int32), values([]int32)]},
// if all history of ts not exist, the timeStamp and value slice will be nil
func (gdb *Gdb) GetIntHistoricalDataWithStamp(groupNames, itemNames []string, timeStamps [][]int32) (GdbHistoricalData, error) {
	st := time.Now()
	if len(groupNames) == len(itemNames) && len(itemNames) == len(timeStamps) {
		if rawNamesMap, namesMap, err := gdb.checkSingleItemDataTypeWithMap(groupNames, itemNames, "int32"); err != nil {
			return GdbHistoricalData{}, err
		} else {
			if m, err := gdb.getHistoricalDataWithTs("int32", rawNamesMap, namesMap, timeStamps...); err != nil {
				return GdbHistoricalData{}, err
			} else {
				return GdbHistoricalData{HistoricalData: m, Times: time.Now().Sub(st).Milliseconds()}, nil
			}
		}
	} else {
		return GdbHistoricalData{}, fmt.Errorf("inconsistent length of parameters")
	}
}

// GetStringHistoricalDataWithStamp get string history data according to the given time stamps, if history corresponding not exist
// we will not returns value
//
// The returned format of values is {"itemName1": [timeStamp([]int32), values([]string)]},
// if all history of ts not exist, the timeStamp and value slice will be nil
func (gdb *Gdb) GetStringHistoricalDataWithStamp(groupNames, itemNames []string, timeStamps [][]int32) (GdbHistoricalData, error) {
	st := time.Now()
	if len(groupNames) == len(itemNames) && len(itemNames) == len(timeStamps) {
		if rawNamesMap, namesMap, err := gdb.checkSingleItemDataTypeWithMap(groupNames, itemNames, "string"); err != nil {
			return GdbHistoricalData{}, err
		} else {
			if m, err := gdb.getHistoricalDataWithTs("string", rawNamesMap, namesMap, timeStamps...); err != nil {
				return GdbHistoricalData{}, err
			} else {
				return GdbHistoricalData{HistoricalData: m, Times: time.Now().Sub(st).Milliseconds()}, nil
			}
		}
	} else {
		return GdbHistoricalData{}, fmt.Errorf("inconsistent length of parameters")
	}
}

// GetBoolHistoricalDataWithStamp get bool history data according to the given time stamps, if history corresponding not exist
// we will not returns value
//
// The returned format of values is {"itemName1": [timeStamp([]int32), values([]bool)]},
// if all history of ts not exist, the timeStamp and value slice will be nil
func (gdb *Gdb) GetBoolHistoricalDataWithStamp(groupNames, itemNames []string, timeStamps [][]int32) (GdbHistoricalData, error) {
	st := time.Now()
	if len(groupNames) == len(itemNames) && len(itemNames) == len(timeStamps) {
		if rawNamesMap, namesMap, err := gdb.checkSingleItemDataTypeWithMap(groupNames, itemNames, "bool"); err != nil {
			return GdbHistoricalData{}, err
		} else {
			if m, err := gdb.getHistoricalDataWithTs("bool", rawNamesMap, namesMap, timeStamps...); err != nil {
				return GdbHistoricalData{}, err
			} else {
				return GdbHistoricalData{HistoricalData: m, Times: time.Now().Sub(st).Milliseconds()}, nil
			}
		}
	} else {
		return GdbHistoricalData{}, fmt.Errorf("inconsistent length of parameters")
	}
}

// GetFloatHistoricalDataWithCondition filter condition must be correct js expression,itemName should be startedWith by item.
// eg: item["itemName1"]>10 && item["itemName2"] > 30 ....
// It should be noted that the entire judgment is based on the itemName with less historical value in the condition.
// If the longest itemName is used as the benchmark, we cannot make an accurate judgment on the AND logic in it.
// Just imagine the history of Item1 It is [3,4,5], and item2 is [10,11]. If item1 is used as the benchmark,
// we cannot determine how much other elements of item2 should be expanded, because the condition may have complicated
// logic about item1 and item2 And or logic, no matter what the number is expanded, there may be a judgment error.
// DeadZone is used to define the maximum number of continuous data allowed by itemName.eg,the deadZoneCount of item x
// is 2, that is all data in x whose number of continuous > 2 will be filtered, itemNames and itemName in filterCondition and zones
// MUST be in the same group
func (gdb *Gdb) GetFloatHistoricalDataWithCondition(groupName string, itemNames []string, startTimes, endTimes, intervals []int32, filterCondition string, zones []DeadZone) (GdbHistoricalData, error) {
	st := time.Now()
	// groupName string, itemNames []string, startTimes, endTimes, intervals []int32, filterCondition string, zones ...DeadZone
	if len(itemNames) == len(startTimes) && len(startTimes) == len(endTimes) && len(endTimes) == len(intervals) {
		if m, err := gdb.getHistoricalDataWithCondition("float32", groupName, itemNames, startTimes, endTimes, intervals, filterCondition, zones...); err != nil {
			return GdbHistoricalData{}, err
		} else {
			return GdbHistoricalData{HistoricalData: m, Times: time.Now().Sub(st).Milliseconds()}, nil
		}
	} else {
		return GdbHistoricalData{}, fmt.Errorf("inconsistent length of parameters")
	}
}

func (gdb *Gdb) GetIntHistoricalDataWithCondition(groupName string, itemNames []string, startTimes, endTimes, intervals []int32, filterCondition string, zones []DeadZone) (GdbHistoricalData, error) {
	st := time.Now()
	if len(itemNames) == len(startTimes) && len(startTimes) == len(endTimes) && len(endTimes) == len(intervals) {
		if m, err := gdb.getHistoricalDataWithCondition("int32", groupName, itemNames, startTimes, endTimes, intervals, filterCondition, zones...); err != nil {
			return GdbHistoricalData{}, err
		} else {
			return GdbHistoricalData{HistoricalData: m, Times: time.Now().Sub(st).Milliseconds()}, nil
		}
	} else {
		return GdbHistoricalData{}, fmt.Errorf("inconsistent length of parameters")
	}
}

func (gdb *Gdb) GetStringHistoricalDataWithCondition(groupName string, itemNames []string, startTimes, endTimes, intervals []int32, filterCondition string, zones []DeadZone) (GdbHistoricalData, error) {
	st := time.Now()
	if len(itemNames) == len(startTimes) && len(startTimes) == len(endTimes) && len(endTimes) == len(intervals) {
		if m, err := gdb.getHistoricalDataWithCondition("string", groupName, itemNames, startTimes, endTimes, intervals, filterCondition, zones...); err != nil {
			return GdbHistoricalData{}, err
		} else {
			return GdbHistoricalData{HistoricalData: m, Times: time.Now().Sub(st).Milliseconds()}, nil
		}
	} else {
		return GdbHistoricalData{}, fmt.Errorf("inconsistent length of parameters")
	}
}

func (gdb *Gdb) GetBoolHistoricalDataWithCondition(groupName string, itemNames []string, startTimes, endTimes, intervals []int32, filterCondition string, zones []DeadZone) (GdbHistoricalData, error) {
	st := time.Now()
	if len(itemNames) == len(startTimes) && len(startTimes) == len(endTimes) && len(endTimes) == len(intervals) {
		if m, err := gdb.getHistoricalDataWithCondition("bool", groupName, itemNames, startTimes, endTimes, intervals, filterCondition, zones...); err != nil {
			return GdbHistoricalData{}, err
		} else {
			return GdbHistoricalData{HistoricalData: m, Times: time.Now().Sub(st).Milliseconds()}, nil
		}
	} else {
		return GdbHistoricalData{}, fmt.Errorf("inconsistent length of parameters")
	}
}

// DeleteFloatHistoricalData delete history data from database according to the given groupNames, itemNames, startTimes and endTimes
// if startTime is -1, that is delete history data from nil-endTime, if endTime is -1, that is delete historical data from startTime-nil
// if both startTime and endTime is -1, that is delete all historical data
func (gdb *Gdb) DeleteFloatHistoricalData(groupNames, itemNames []string, startTimes, endTimes []int32) (TimeRows, error) {
	st := time.Now()
	if len(groupNames) == len(itemNames) && len(startTimes) == len(endTimes) {
		if _, namesMap, err := gdb.checkSingleItemDataTypeWithMap(groupNames, itemNames, "float32"); err != nil {
			return TimeRows{}, err
		} else {
			if r, err := gdb.deleteHistoricalData("float32", namesMap, startTimes, endTimes); err != nil {
				return TimeRows{}, err
			} else {
				return TimeRows{EffectedRows: r, Times: time.Now().Sub(st).Milliseconds()}, nil
			}
		}
	} else {
		return TimeRows{}, fmt.Errorf("inconsistent length of parameters")
	}
}

func (gdb *Gdb) DeleteIntHistoricalData(groupNames, itemNames []string, startTimes, endTimes []int32) (TimeRows, error) {
	st := time.Now()
	if len(groupNames) == len(itemNames) && len(startTimes) == len(endTimes) {
		if _, namesMap, err := gdb.checkSingleItemDataTypeWithMap(groupNames, itemNames, "int32"); err != nil {
			return TimeRows{}, err
		} else {
			if r, err := gdb.deleteHistoricalData("int32", namesMap, startTimes, endTimes); err != nil {
				return TimeRows{}, err
			} else {
				return TimeRows{EffectedRows: r, Times: time.Now().Sub(st).Milliseconds()}, nil
			}
		}
	} else {
		return TimeRows{}, fmt.Errorf("inconsistent length of parameters")
	}
}

func (gdb *Gdb) DeleteStringHistoricalData(groupNames, itemNames []string, startTimes, endTimes []int32) (TimeRows, error) {
	st := time.Now()
	if len(groupNames) == len(itemNames) && len(startTimes) == len(endTimes) {
		if _, namesMap, err := gdb.checkSingleItemDataTypeWithMap(groupNames, itemNames, "string"); err != nil {
			return TimeRows{}, err
		} else {
			if r, err := gdb.deleteHistoricalData("string", namesMap, startTimes, endTimes); err != nil {
				return TimeRows{}, err
			} else {
				return TimeRows{EffectedRows: r, Times: time.Now().Sub(st).Milliseconds()}, nil
			}
		}
	} else {
		return TimeRows{}, fmt.Errorf("inconsistent length of parameters")
	}
}

func (gdb *Gdb) DeleteBoolHistoricalData(groupNames, itemNames []string, startTimes, endTimes []int32) (TimeRows, error) {
	st := time.Now()
	if len(groupNames) == len(itemNames) && len(startTimes) == len(endTimes) {
		if _, namesMap, err := gdb.checkSingleItemDataTypeWithMap(groupNames, itemNames, "bool"); err != nil {
			return TimeRows{}, err
		} else {
			if r, err := gdb.deleteHistoricalData("bool", namesMap, startTimes, endTimes); err != nil {
				return TimeRows{}, err
			} else {
				return TimeRows{EffectedRows: r, Times: time.Now().Sub(st).Milliseconds()}, nil
			}
		}
	} else {
		return TimeRows{}, fmt.Errorf("inconsistent length of parameters")
	}
}

// CleanItemData delete item and history data of the given item, this may take long
// time and may block writing operation, you should use it carefully
func (gdb *Gdb) CleanItemData(itemInfo DeletedItemsInfo) (TimeRows, error) {
	st := time.Now()
	groupName := itemInfo.GroupName
	condition := itemInfo.Condition
	items, err := gdb.query("select itemName from `" + groupName + "` where " + condition)
	if len(items) == 0 {
		return TimeRows{}, fmt.Errorf("conditionError: " + condition)
	}
	if err != nil {
		return TimeRows{}, err
	}
	rows, err := gdb.updateItem("delete from `" + groupName + "` where " + condition)
	if err != nil {
		return TimeRows{}, err
	}
	dataTypes := map[string][]string{} // key is dataType, value is itemNames
	var startTimeStamp, endTimeStamp []int32
	for _, item := range items {
		itemName := item["itemName"] + joiner + groupName
		startTimeStamp = append(startTimeStamp, -1)
		endTimeStamp = append(endTimeStamp, -1)
		dataType, _ := gdb.rtDbFilter.Get(itemName)
		{
			switch dataType.(string) {
			case "float32":
				if _, ok := dataTypes["float32"]; ok {
					dataTypes["float32"] = append(dataTypes["float32"], item["itemName"])
				} else {
					dataTypes["float32"] = []string{item["itemName"]}
				}
				break
			case "int32":
				if _, ok := dataTypes["int32"]; ok {
					dataTypes["int32"] = append(dataTypes["int32"], item["itemName"])
				} else {
					dataTypes["int32"] = []string{item["itemName"]}
				}
				break
			case "string":
				if _, ok := dataTypes["string"]; ok {
					dataTypes["string"] = append(dataTypes["string"], item["itemName"])
				} else {
					dataTypes["string"] = []string{item["itemName"]}
				}
				break
			default:
				if _, ok := dataTypes["bool"]; ok {
					dataTypes["bool"] = append(dataTypes["bool"], item["itemName"])
				} else {
					dataTypes["bool"] = []string{item["itemName"]}
				}
				break
			}
		}
	}
	// delete history
	// check whether is syncing historical data
	t := time.NewTicker(time.Second)
	for {
		select {
		case <-t.C:
			if !gdb.syncStatus {
				t.Stop()
				goto next
			}
		}
	}
next:
	groupNames := strings.Split(strings.TrimRight(strings.Repeat(groupName+",", len(items)), ","), ",")
	g := errgroup.Group{}
	g.Go(func() error {
		if itemNames := dataTypes["float32"]; itemNames != nil {
			_, err := gdb.DeleteFloatHistoricalData(groupNames, dataTypes["float32"], startTimeStamp, endTimeStamp)
			return err
		}
		return nil
	})
	g.Go(func() error {
		if itemNames := dataTypes["int32"]; itemNames != nil {
			_, err := gdb.DeleteIntHistoricalData(groupNames, itemNames, startTimeStamp, endTimeStamp)
			return err
		}
		return nil
	})
	g.Go(func() error {
		if itemNames := dataTypes["string"]; itemNames != nil {
			_, err := gdb.DeleteStringHistoricalData(groupNames, itemNames, startTimeStamp, endTimeStamp)
			return err
		}
		return nil
	})
	g.Go(func() error {
		if itemNames := dataTypes["bool"]; itemNames != nil {
			_, err := gdb.DeleteBoolHistoricalData(groupNames, itemNames, startTimeStamp, endTimeStamp)
			return err
		}
		return nil
	})
	if err := g.Wait(); err != nil {
		return TimeRows{}, err
	}
	return TimeRows{EffectedRows: int(rows), Times: time.Since(st).Milliseconds()}, nil
}

// ReLoadDb delete keys in db whose value is nil, then compact db again, during this time, all write operation will failed
// so you should not write data to db during reload
func (gdb *Gdb) ReLoadDb() (TimeRows, error) {
	// check whether is syncing historical data
	t := time.NewTicker(time.Second)
	for {
		select {
		case <-t.C:
			if !gdb.syncStatus {
				t.Stop()
				goto next
			}
		}
	}
next:
	st := time.Now()
	gdb.isReloading = true // lock gdb
	defer func() {
		gdb.isReloading = false
	}()
	dataTypeMap := map[string]map[string][]string{} // dataType==>{groupName: []string{itemName}}
	for item := range gdb.rtDbFilter.IterBuffered() {
		switch item.Val.(string) {
		case "float32":
			groupName := strings.Split(item.Key, joiner)[1] // get groupName
			if _, ok := dataTypeMap["float32"]; ok {
				dataTypeMap["float32"][groupName] = append(dataTypeMap["float32"][groupName], item.Key)
			} else {
				dataTypeMap["float32"] = map[string][]string{groupName: {item.Key}}
			}
			break
		case "int32":
			groupName := strings.Split(item.Key, joiner)[1] // get groupName
			if _, ok := dataTypeMap["int32"]; ok {
				dataTypeMap["int32"][groupName] = append(dataTypeMap["int32"][groupName], item.Key)
			} else {
				dataTypeMap["int32"] = map[string][]string{groupName: {item.Key}}
			}
			break
		case "string":
			groupName := strings.Split(item.Key, joiner)[1] // get groupName
			if _, ok := dataTypeMap["string"]; ok {
				dataTypeMap["string"][groupName] = append(dataTypeMap["string"][groupName], item.Key)
			} else {
				dataTypeMap["string"] = map[string][]string{groupName: {item.Key}}
			}
			break
		default:
			groupName := strings.Split(item.Key, joiner)[1] // get groupName
			if _, ok := dataTypeMap["bool"]; ok {
				dataTypeMap["bool"][groupName] = append(dataTypeMap["bool"][groupName], item.Key)
			} else {
				dataTypeMap["bool"] = map[string][]string{groupName: {item.Key}}
			}
			break
		}
	}
	g := errgroup.Group{}
	counts := [4]int{}
	for index, dataType := range dataTypes {
		dt, i := dataType, index
		g.Go(func() error {
			for groupName, itemNames := range dataTypeMap[dt] {
				db := gdb.hisDb[dt][groupName]
				if r, err := reload(db, itemNames...); err != nil {
					return err
				} else {
					counts[i] = r
				}
			}
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return TimeRows{}, err
	}
	return TimeRows{Times: time.Since(st).Milliseconds(), EffectedRows: counts[0] + counts[1] + counts[2] + counts[3]}, nil
}

// CloseGdb will sync realTimeData and history data, then close database
func (gdb *Gdb) CloseGdb() error {
	// block util all history data in memory sync to disk
	t := time.NewTicker(time.Second)
	for {
		select {
		case <-t.C:
			if !gdb.syncStatus {
				// finish syncing
				t.Stop()
				goto next
			}
		}
	}
next:
	eg := sync.WaitGroup{}
	eg.Add(2)
	go func() {
		// syncRealTimeData
		defer eg.Done()
		_ = gdb.rtDb.Sync()
	}()
	go func() {
		// sync history data
		defer eg.Done()
		_ = gdb.innerSync(time.Now())
	}()
	eg.Wait()
	// close leveldb
	for _, dataType := range dataTypes {
		for _, db := range gdb.hisDb[dataType] {
			_ = db.Close()
		}
	}
	return nil
}

func reload(db *leveldb.DB, itemNames ...string) (int, error) {
	row := 0
	if sn, err := db.GetSnapshot(); err != nil {
		return -1, err
	} else {
		defer sn.Release()
		keys := map[string]map[string][]byte{} // keys to be deleted, key is itemName, values if {"st": startKey, "et": endKey}
		for _, itemName := range itemNames {
			it := sn.NewIterator(util.BytesPrefix(convertStringToByte(itemName)), nil)
			tk := [][]byte{}
			for it.Next() {
				if len(it.Value()) == 0 {
					tk = append(tk, it.Key())
				}
			}
			row += len(tk)
			if len(tk) != 0 {
				keys[itemName] = map[string][]byte{"st": tk[0], "et": tk[len(tk)-1]}
				// delete key in db
				for _, key := range tk {
					if err := db.Delete(key, nil); err != nil {
						return -1, err
					}
				}
			}
			it.Release()
		}
		// Compact db
		for _, key := range keys {
			if err := db.CompactRange(util.Range{
				Start: key["st"],
				Limit: key["et"],
			}); err != nil {
				return -1, err
			}
		}
		return row, nil
	}
}

func (gdb *Gdb) getHistoricalData(dataType string, rawNamesMap map[string][]string, namesMap map[string][]string, startTimeStamps, endTimeStamps, intervals []int32, flags ...bool) (cmap.ConcurrentMap, error) {
	m := cmap.New()
	var lengthMap []int
	for _, names := range namesMap {
		lengthMap = make([]int, len(names))
		break
	}
	dg := errgroup.Group{}
	switch dataType {
	case "float32":
		goto float32
	case "int32":
		goto int32
	case "string":
		goto string32
	case "bool":
		goto bool32
	default:
		return nil, fmt.Errorf("unknown dataType " + dataType)
	}
float32:
	for groupName, groupItemNames := range namesMap {
		groupName := groupName
		groupItemNames := groupItemNames       // itemName in memory, that is itemName + joiner + groupName
		rawItemNames := rawNamesMap[groupName] // origin itemName
		dg.Go(func() error {
			if sn, err := gdb.hisDb["float32"][groupName].GetSnapshot(); err != nil {
				return err
			} else {
				// get history
				defer sn.Release()
				g := errgroup.Group{}
				for i := 0; i < len(groupItemNames); i++ {
					index := i
					g.Go(func() error {
						md, ok := gdb.floatRmHisDb.Get(groupItemNames[index]) // itemName is memory is itemName + joiner + groupName, in leveldb is itemName + joiner + startTimeStamp
						var leveldbHis, memHis []float32
						var leveldbTs, memTs []int32
						for j := 0; j < len(startTimeStamps); j++ {
							if startTimeStamps[j] >= endTimeStamps[j] {
								continue
							}
							jj := j
							g1 := errgroup.Group{}
							// get history data in leveldb, historical data in the database can not be repeated
							g1.Go(func() error {
								var st int32 // start time stamp of data
								s := groupItemNames[index] + joiner + strconv.Itoa(int(startTimeStamps[jj])-int(gdb.hisTimeDuration.Seconds()))
								e := groupItemNames[index] + joiner + strconv.Itoa(int(endTimeStamps[jj])+int(gdb.hisTimeDuration.Seconds()))
								it := sn.NewIterator(&util.Range{Start: convertStringToByte(s), Limit: convertStringToByte(e)}, nil)
								defer it.Release()
								count := 0
								for it.Next() {
									d := &pb.FloatHistoricalData{}
									iv := it.Value()
									if len(iv) == 0 {
										continue
									}
									if err := proto.Unmarshal(iv, d); err != nil {
										return err
									}
									ts, values := d.GetTimeStamps(), d.GetValues()
									startIndex, endIndex := getIndex(startTimeStamps[jj], endTimeStamps[jj], ts)
									if startIndex == -1 || endIndex == -1 {
										continue
									}
									if count == 0 {
										st = ts[startIndex]
										count++
									}
									for k := startIndex; k < endIndex; k++ {
										if (ts[k]-st)%intervals[index] == 0 {
											leveldbHis = append(leveldbHis, values[k])
											leveldbTs = append(leveldbTs, ts[k])
										}
									}
								}
								return nil
							})
							if ok {
								// get history data in memory
								if len(md.TimeStamps) != 0 {
									if endTimeStamps[jj] <= md.TimeStamps[0] || startTimeStamps[jj] >= md.TimeStamps[len(md.TimeStamps)-1] {
									} else {
										g1.Go(func() error {
											var st int32 // start time stamp of data
											startIndex, endIndex := getIndex(startTimeStamps[jj], endTimeStamps[jj], md.TimeStamps)
											if startIndex == -1 || endIndex == -1 {
												return nil
											}
											st = md.TimeStamps[startIndex]
											for k := startIndex; k < endIndex; k++ {
												if (md.TimeStamps[k]-st)%intervals[index] == 0 {
													memHis = append(memHis, md.Values[k])
													memTs = append(memTs, md.TimeStamps[k])
												}
											}
											return nil
										})
									}
								}
							}
							if err := g1.Wait(); err != nil {
								return err
							}
						}
						leveldbHis = append(leveldbHis, memHis...) // history data
						leveldbTs = append(leveldbTs, memTs...)    // ts
						if len(flags) != 0 {
							if flags[0] {
								lengthMap[index] = len(leveldbTs)
							}
						}
						m.Set(rawItemNames[index], []interface{}{leveldbTs, leveldbHis})
						return nil
					})
				}
				if err := g.Wait(); err != nil {
					return err
				}
				if len(flags) != 0 {
					if flags[0] {
						m.Set("minLength", From(lengthMap).Min().(int))
					}
				}
				return nil
			}
		})
	}
	return m, dg.Wait()
int32:
	for groupName, groupItemNames := range namesMap {
		groupName := groupName
		groupItemNames := groupItemNames
		rawItemNames := rawNamesMap[groupName]
		dg.Go(func() error {
			if sn, err := gdb.hisDb["int32"][groupName].GetSnapshot(); err != nil {
				return err
			} else {
				// get history
				defer sn.Release()
				g := errgroup.Group{}
				for i := 0; i < len(groupItemNames); i++ {
					index := i
					g.Go(func() error {
						md, ok := gdb.intRmHisDb.Get(groupItemNames[index])
						var leveldbHis, memHis []int32
						var leveldbTs, memTs []int32
						for j := 0; j < len(startTimeStamps); j++ {
							if startTimeStamps[j] >= endTimeStamps[j] {
								continue
							}
							g1 := errgroup.Group{}
							jj := j
							// get history data in leveldb
							g1.Go(func() error {
								var st int32 // start time stamp of data
								s := groupItemNames[index] + joiner + strconv.Itoa(int(startTimeStamps[jj])-int(gdb.hisTimeDuration.Seconds()))
								e := groupItemNames[index] + joiner + strconv.Itoa(int(endTimeStamps[jj])+int(gdb.hisTimeDuration.Seconds()))
								it := sn.NewIterator(&util.Range{Start: convertStringToByte(s), Limit: convertStringToByte(e)}, nil)
								defer it.Release()
								count := 0
								for it.Next() {
									d := &pb.IntHistoricalData{}
									iv := it.Value()
									if len(iv) == 0 {
										continue
									}
									if err := proto.Unmarshal(iv, d); err != nil {
										return err
									}
									ts, values := d.GetTimeStamps(), d.GetValues()
									startIndex, endIndex := getIndex(startTimeStamps[jj], endTimeStamps[jj], ts)
									if startIndex == -1 || endIndex == -1 {
										continue
									}
									if count == 0 {
										st = ts[startIndex]
										count++
									}
									for k := startIndex; k < endIndex; k++ {
										if (ts[k]-st)%intervals[index] == 0 {
											leveldbHis = append(leveldbHis, values[k])
											leveldbTs = append(leveldbTs, ts[k])
										}
									}
								}
								return nil
							})
							if ok {
								// get history data in memory
								if len(md.TimeStamps) != 0 {
									if endTimeStamps[jj] <= md.TimeStamps[0] || startTimeStamps[jj] >= md.TimeStamps[len(md.TimeStamps)-1] {
									} else {
										g1.Go(func() error {
											var st int32 // start time stamp of data
											startIndex, endIndex := getIndex(startTimeStamps[jj], endTimeStamps[jj], md.TimeStamps)
											if startIndex == -1 || endIndex == -1 {
												return nil
											}
											st = md.TimeStamps[startIndex]
											for k := startIndex; k < endIndex; k++ {
												if (md.TimeStamps[k]-st)%intervals[index] == 0 {
													memHis = append(memHis, md.Values[k])
													memTs = append(memTs, md.TimeStamps[k])
												}
											}
											return nil
										})
									}
								}
							}
							if err := g1.Wait(); err != nil {
								return err
							}
						}
						leveldbHis = append(leveldbHis, memHis...) // history data
						leveldbTs = append(leveldbTs, memTs...)    // ts
						if len(flags) != 0 {
							if flags[0] {
								lengthMap[index] = len(leveldbTs)
							}
						}
						m.Set(rawItemNames[index], []interface{}{leveldbTs, leveldbHis})
						return nil
					})
				}
				if err := g.Wait(); err != nil {
					return err
				}
				if len(flags) != 0 {
					if flags[0] {
						m.Set("minLength", From(lengthMap).Min().(int))
					}
				}
				return nil
			}
		})
	}
	return m, dg.Wait()
string32:
	for groupName, groupItemNames := range namesMap {
		groupName := groupName
		groupItemNames := groupItemNames
		rawItemNames := rawNamesMap[groupName]
		dg.Go(func() error {
			if sn, err := gdb.hisDb["string"][groupName].GetSnapshot(); err != nil {
				return err
			} else {
				// get history
				defer sn.Release()
				g := errgroup.Group{}
				for i := 0; i < len(groupItemNames); i++ {
					index := i
					g.Go(func() error {
						md, ok := gdb.stringRmHisDb.Get(groupItemNames[index])
						var leveldbHis, memHis []string
						var leveldbTs, memTs []int32
						for j := 0; j < len(startTimeStamps); j++ {
							if startTimeStamps[j] >= endTimeStamps[j] {
								continue
							}
							g1 := errgroup.Group{}
							jj := j
							// get history data in leveldb
							g1.Go(func() error {
								var st int32 // start time stamp of data
								s := groupItemNames[index] + joiner + strconv.Itoa(int(startTimeStamps[jj])-int(gdb.hisTimeDuration.Seconds()))
								e := groupItemNames[index] + joiner + strconv.Itoa(int(endTimeStamps[jj])+int(gdb.hisTimeDuration.Seconds()))
								it := sn.NewIterator(&util.Range{Start: convertStringToByte(s), Limit: convertStringToByte(e)}, nil)
								defer it.Release()
								count := 0
								for it.Next() {
									d := &pb.StringHistoricalData{}
									iv := it.Value()
									if len(iv) == 0 {
										continue
									}
									if err := proto.Unmarshal(iv, d); err != nil {
										return err
									}
									ts, values := d.GetTimeStamps(), d.GetValues()
									startIndex, endIndex := getIndex(startTimeStamps[jj], endTimeStamps[jj], ts)
									if startIndex == -1 || endIndex == -1 {
										continue
									}
									if count == 0 {
										st = ts[startIndex]
										count++
									}
									for k := startIndex; k < endIndex; k++ {
										if (ts[k]-st)%intervals[index] == 0 {
											leveldbHis = append(leveldbHis, values[k])
											leveldbTs = append(leveldbTs, ts[k])
										}
									}
								}
								return nil
							})
							if ok {
								// get history data in memory
								if len(md.TimeStamps) != 0 {
									if endTimeStamps[jj] <= md.TimeStamps[0] || startTimeStamps[jj] >= md.TimeStamps[len(md.TimeStamps)-1] {
									} else {
										g1.Go(func() error {
											var st int32 // start time stamp of data
											startIndex, endIndex := getIndex(startTimeStamps[jj], endTimeStamps[jj], md.TimeStamps)
											if startIndex == -1 || endIndex == -1 {
												return nil
											}
											st = md.TimeStamps[startIndex]
											for k := startIndex; k < endIndex; k++ {
												if (md.TimeStamps[k]-st)%intervals[index] == 0 {
													memHis = append(memHis, md.Values[k])
													memTs = append(memTs, md.TimeStamps[k])
												}
											}
											return nil
										})
									}
								}
							}
							if err := g1.Wait(); err != nil {
								return err
							}
						}
						leveldbHis = append(leveldbHis, memHis...) // history data
						leveldbTs = append(leveldbTs, memTs...)    // ts
						if len(flags) != 0 {
							if flags[0] {
								lengthMap[index] = len(leveldbTs)
							}
						}
						m.Set(rawItemNames[index], []interface{}{leveldbTs, leveldbHis})
						return nil
					})
				}
				if err := g.Wait(); err != nil {
					return err
				}
				if len(flags) != 0 {
					if flags[0] {
						m.Set("minLength", From(lengthMap).Min().(int))
					}
				}
				return nil
			}
		})
	}
	return m, dg.Wait()
bool32:
	for groupName, groupItemNames := range namesMap {
		groupName := groupName
		groupItemNames := groupItemNames
		rawItemNames := rawNamesMap[groupName]
		dg.Go(func() error {
			if sn, err := gdb.hisDb["bool"][groupName].GetSnapshot(); err != nil {
				return err
			} else {
				// get history
				defer sn.Release()
				g := errgroup.Group{}
				for i := 0; i < len(groupItemNames); i++ {
					index := i
					g.Go(func() error {
						md, ok := gdb.boolRmHisDb.Get(groupItemNames[index])
						var leveldbHis, memHis []bool
						var leveldbTs, memTs []int32
						for j := 0; j < len(startTimeStamps); j++ {
							if startTimeStamps[j] >= endTimeStamps[j] {
								continue
							}
							g1 := errgroup.Group{}
							jj := j
							// get history data in leveldb
							g1.Go(func() error {
								var st int32 // start time stamp of data
								s := groupItemNames[index] + joiner + strconv.Itoa(int(startTimeStamps[jj])-int(gdb.hisTimeDuration.Seconds()))
								e := groupItemNames[index] + joiner + strconv.Itoa(int(endTimeStamps[jj])+int(gdb.hisTimeDuration.Seconds()))
								it := sn.NewIterator(&util.Range{Start: convertStringToByte(s), Limit: convertStringToByte(e)}, nil)
								defer it.Release()
								count := 0
								for it.Next() {
									d := &pb.BoolHistoricalData{}
									iv := it.Value()
									if len(iv) == 0 {
										continue
									}
									if err := proto.Unmarshal(iv, d); err != nil {
										return err
									}
									ts, values := d.GetTimeStamps(), d.GetValues()
									startIndex, endIndex := getIndex(startTimeStamps[jj], endTimeStamps[jj], ts)
									if startIndex == -1 || endIndex == -1 {
										continue
									}
									if count == 0 {
										st = ts[startIndex]
										count++
									}
									for k := startIndex; k < endIndex; k++ {
										if (ts[k]-st)%intervals[index] == 0 {
											leveldbHis = append(leveldbHis, values[k])
											leveldbTs = append(leveldbTs, ts[k])
										}
									}
								}
								return nil
							})
							if ok {
								// get history data in memory
								if len(md.TimeStamps) != 0 {
									if endTimeStamps[jj] <= md.TimeStamps[0] || startTimeStamps[jj] >= md.TimeStamps[len(md.TimeStamps)-1] {
									} else {
										g1.Go(func() error {
											var st int32 // start time stamp of data
											startIndex, endIndex := getIndex(startTimeStamps[jj], endTimeStamps[jj], md.TimeStamps)
											if startIndex == -1 || endIndex == -1 {
												return nil
											}
											st = md.TimeStamps[startIndex]
											for k := startIndex; k < endIndex; k++ {
												if (md.TimeStamps[k]-st)%intervals[index] == 0 {
													memHis = append(memHis, md.Values[k])
													memTs = append(memTs, md.TimeStamps[k])
												}
											}
											return nil
										})
									}
								}
							}
							if err := g1.Wait(); err != nil {
								return err
							}
						}
						leveldbHis = append(leveldbHis, memHis...) // history data
						leveldbTs = append(leveldbTs, memTs...)    // ts
						if len(flags) != 0 {
							if flags[0] {
								lengthMap[index] = len(leveldbTs)
							}
						}
						m.Set(rawItemNames[index], []interface{}{leveldbTs, leveldbHis})
						return nil
					})
				}
				if err := g.Wait(); err != nil {
					return err
				}
				if len(flags) != 0 {
					if flags[0] {
						m.Set("minLength", From(lengthMap).Min().(int))
					}
				}
				return nil
			}
		})
	}
	return m, dg.Wait()
}

func (gdb *Gdb) getHistoricalDataWithTs(dataType string, rawNamesMap, namesMap map[string][]string, timeStamps ...[]int32) (cmap.ConcurrentMap, error) {
	m := cmap.New()
	dg := errgroup.Group{}
	switch dataType {
	case "float32":
		goto float32
	case "int32":
		goto int32
	case "string":
		goto string32
	case "bool":
		goto bool32
	default:
		return nil, fmt.Errorf("unknown dataType " + dataType)
	}
float32:
	for groupName, groupItemNames := range namesMap {
		groupName := groupName
		groupItemNames := groupItemNames
		rawItemNames := rawNamesMap[groupName]
		dg.Go(func() error {
			if sn, err := gdb.hisDb["float32"][groupName].GetSnapshot(); err != nil {
				return err
			} else {
				defer sn.Release()
				g := errgroup.Group{}
				for i := 0; i < len(groupItemNames); i++ {
					index := i
					g.Go(func() error {
						ts := timeStamps[index]
						md, ok := gdb.floatRmHisDb.Get(groupItemNames[index])
						var leveldbHis, memHis []float32
						var leveldbTs, memTs []int32
						g1 := errgroup.Group{}
						g1.Go(func() error {
							for j := 0; j < len(ts); j++ {
								// get history data in leveldb
								s := groupItemNames[index] + joiner + strconv.Itoa(int(ts[j]-int32(gdb.hisTimeDuration.Seconds())))
								e := groupItemNames[index] + joiner + strconv.Itoa(int(ts[j]+int32(gdb.hisTimeDuration.Seconds())))
								it := sn.NewIterator(&util.Range{Start: convertStringToByte(s), Limit: convertStringToByte(e)}, nil)
								for it.Next() {
									d := &pb.FloatHistoricalData{}
									iv := it.Value()
									if len(iv) == 0 {
										continue
									}
									if err := proto.Unmarshal(iv, d); err != nil {
										return err
									}
									tts, values := d.GetTimeStamps(), d.GetValues()
									if tts[0] == ts[j] {
										leveldbHis = append(leveldbHis, values[0])
										leveldbTs = append(leveldbTs, tts[0])
										break
									}
									if k := getMiddleIndex(ts[j], tts); k == -1 {
										continue
									} else {
										leveldbHis = append(leveldbHis, values[k])
										leveldbTs = append(leveldbTs, tts[k])
										break
									}
								}
								it.Release()
							}
							return nil
						})
						if ok {
							// get history in memory
							g1.Go(func() error {
								for k := 0; k < len(ts); k++ {
									if ts[k] == md.TimeStamps[0] {
										memHis = append(memHis, md.Values[0])
										memTs = append(memTs, md.TimeStamps[0])
									} else {
										if j := getMiddleIndex(ts[k], md.TimeStamps); j == -1 {
											continue
										} else {
											memHis = append(memHis, md.Values[j])
											memTs = append(memTs, md.TimeStamps[j])
										}
									}
								}
								return nil
							})
						}
						if err := g1.Wait(); err != nil {
							return err
						}
						leveldbTs = append(leveldbTs, memTs...)
						leveldbHis = append(leveldbHis, memHis...)
						m.Set(rawItemNames[index], []interface{}{leveldbTs, leveldbHis})
						return nil
					})
				}
				if err := g.Wait(); err != nil {
					return err
				}
				return nil
			}
		})
	}
	return m, dg.Wait()
int32:
	for groupName, groupItemNames := range namesMap {
		groupName := groupName
		groupItemNames := groupItemNames
		rawItemNames := rawNamesMap[groupName]
		dg.Go(func() error {
			if sn, err := gdb.hisDb["int32"][groupName].GetSnapshot(); err != nil {
				return err
			} else {
				defer sn.Release()
				g := errgroup.Group{}
				for i := 0; i < len(groupItemNames); i++ {
					index := i
					g.Go(func() error {
						ts := timeStamps[index]
						md, ok := gdb.intRmHisDb.Get(groupItemNames[index])
						var leveldbHis, memHis []int32
						var leveldbTs, memTs []int32
						g1 := errgroup.Group{}
						g1.Go(func() error {
							for j := 0; j < len(ts); j++ {
								// get history data in leveldb
								s := groupItemNames[index] + joiner + strconv.Itoa(int(ts[j]-int32(gdb.hisTimeDuration.Seconds())))
								e := groupItemNames[index] + joiner + strconv.Itoa(int(ts[j]+int32(gdb.hisTimeDuration.Seconds())))
								it := sn.NewIterator(&util.Range{Start: convertStringToByte(s), Limit: convertStringToByte(e)}, nil)
								for it.Next() {
									d := &pb.IntHistoricalData{}
									iv := it.Value()
									if len(iv) == 0 {
										continue
									}
									if err := proto.Unmarshal(iv, d); err != nil {
										return err
									}
									tts, values := d.GetTimeStamps(), d.GetValues()
									if tts[0] == ts[j] {
										leveldbHis = append(leveldbHis, values[0])
										leveldbTs = append(leveldbTs, tts[0])
										break
									}
									if k := getMiddleIndex(ts[j], tts); k == -1 {
										continue
									} else {
										leveldbHis = append(leveldbHis, values[k])
										leveldbTs = append(leveldbTs, tts[k])
										break
									}
								}
								it.Release()
							}
							return nil
						})
						if ok {
							// get history in memory
							g1.Go(func() error {
								for k := 0; k < len(ts); k++ {
									if ts[k] == md.TimeStamps[0] {
										memHis = append(memHis, md.Values[0])
										memTs = append(memTs, md.TimeStamps[0])
									} else {
										if j := getMiddleIndex(ts[k], md.TimeStamps); j == -1 {
											continue
										} else {
											memHis = append(memHis, md.Values[j])
											memTs = append(memTs, md.TimeStamps[j])
										}
									}
								}
								return nil
							})
						}
						if err := g1.Wait(); err != nil {
							return err
						}
						leveldbTs = append(leveldbTs, memTs...)
						leveldbHis = append(leveldbHis, memHis...)
						m.Set(rawItemNames[index], []interface{}{leveldbTs, leveldbHis})
						return nil
					})
				}
				if err := g.Wait(); err != nil {
					return err
				}
				return nil
			}
		})
	}
	return m, dg.Wait()
string32:
	for groupName, groupItemNames := range namesMap {
		groupName := groupName
		groupItemNames := groupItemNames
		rawItemNames := rawNamesMap[groupName]
		dg.Go(func() error {
			if sn, err := gdb.hisDb["string"][groupName].GetSnapshot(); err != nil {
				return err
			} else {
				defer sn.Release()
				g := errgroup.Group{}
				for i := 0; i < len(groupItemNames); i++ {
					index := i
					g.Go(func() error {
						ts := timeStamps[index]
						md, ok := gdb.stringRmHisDb.Get(groupItemNames[index])
						var leveldbHis, memHis []string
						var leveldbTs, memTs []int32
						g1 := errgroup.Group{}
						g1.Go(func() error {
							for j := 0; j < len(ts); j++ {
								// get history data in leveldb
								s := groupItemNames[index] + joiner + strconv.Itoa(int(ts[j]-int32(gdb.hisTimeDuration.Seconds())))
								e := groupItemNames[index] + joiner + strconv.Itoa(int(ts[j]+int32(gdb.hisTimeDuration.Seconds())))
								it := sn.NewIterator(&util.Range{Start: convertStringToByte(s), Limit: convertStringToByte(e)}, nil)
								for it.Next() {
									d := &pb.StringHistoricalData{}
									iv := it.Value()
									if len(iv) == 0 {
										continue
									}
									if err := proto.Unmarshal(iv, d); err != nil {
										return err
									}
									tts, values := d.GetTimeStamps(), d.GetValues()
									if tts[0] == ts[j] {
										leveldbHis = append(leveldbHis, values[0])
										leveldbTs = append(leveldbTs, tts[0])
										break
									}
									if k := getMiddleIndex(ts[j], tts); k == -1 {
										continue
									} else {
										leveldbHis = append(leveldbHis, values[k])
										leveldbTs = append(leveldbTs, tts[k])
										break
									}
								}
								it.Release()
							}
							return nil
						})
						if ok {
							// get history in memory
							g1.Go(func() error {
								for k := 0; k < len(ts); k++ {
									if ts[k] == md.TimeStamps[0] {
										memHis = append(memHis, md.Values[0])
										memTs = append(memTs, md.TimeStamps[0])
									} else {
										if j := getMiddleIndex(ts[k], md.TimeStamps); j == -1 {
											continue
										} else {
											memHis = append(memHis, md.Values[j])
											memTs = append(memTs, md.TimeStamps[j])
										}
									}
								}
								return nil
							})
						}
						if err := g1.Wait(); err != nil {
							return err
						}
						leveldbTs = append(leveldbTs, memTs...)
						leveldbHis = append(leveldbHis, memHis...)
						m.Set(rawItemNames[index], []interface{}{leveldbTs, leveldbHis})
						return nil
					})
				}
				if err := g.Wait(); err != nil {
					return err
				}
				return nil
			}
		})
	}
	return m, dg.Wait()
bool32:
	for groupName, groupItemNames := range namesMap {
		groupName := groupName
		groupItemNames := groupItemNames
		rawItemNames := rawNamesMap[groupName]
		dg.Go(func() error {
			if sn, err := gdb.hisDb["bool"][groupName].GetSnapshot(); err != nil {
				return err
			} else {
				defer sn.Release()
				g := errgroup.Group{}
				for i := 0; i < len(groupItemNames); i++ {
					index := i
					g.Go(func() error {
						ts := timeStamps[index]
						md, ok := gdb.boolRmHisDb.Get(groupItemNames[index])
						var leveldbHis, memHis []bool
						var leveldbTs, memTs []int32
						g1 := errgroup.Group{}
						g1.Go(func() error {
							for j := 0; j < len(ts); j++ {
								// get history data in leveldb
								s := groupItemNames[index] + joiner + strconv.Itoa(int(ts[j]-int32(gdb.hisTimeDuration.Seconds())))
								e := groupItemNames[index] + joiner + strconv.Itoa(int(ts[j]+int32(gdb.hisTimeDuration.Seconds())))
								it := sn.NewIterator(&util.Range{Start: convertStringToByte(s), Limit: convertStringToByte(e)}, nil)
								for it.Next() {
									d := &pb.BoolHistoricalData{}
									iv := it.Value()
									if len(iv) == 0 {
										continue
									}
									if err := proto.Unmarshal(iv, d); err != nil {
										return err
									}
									tts, values := d.GetTimeStamps(), d.GetValues()
									if tts[0] == ts[j] {
										leveldbHis = append(leveldbHis, values[0])
										leveldbTs = append(leveldbTs, tts[0])
										break
									}
									if k := getMiddleIndex(ts[j], tts); k == -1 {
										continue
									} else {
										leveldbHis = append(leveldbHis, values[k])
										leveldbTs = append(leveldbTs, tts[k])
										break
									}
								}
								it.Release()
							}
							return nil
						})
						if ok {
							// get history in memory
							g1.Go(func() error {
								for k := 0; k < len(ts); k++ {
									if ts[k] == md.TimeStamps[0] {
										memHis = append(memHis, md.Values[0])
										memTs = append(memTs, md.TimeStamps[0])
									} else {
										if j := getMiddleIndex(ts[k], md.TimeStamps); j == -1 {
											continue
										} else {
											memHis = append(memHis, md.Values[j])
											memTs = append(memTs, md.TimeStamps[j])
										}
									}
								}
								return nil
							})
						}
						if err := g1.Wait(); err != nil {
							return err
						}
						leveldbTs = append(leveldbTs, memTs...)
						leveldbHis = append(leveldbHis, memHis...)
						m.Set(rawItemNames[index], []interface{}{leveldbTs, leveldbHis})
						return nil
					})
				}
				if err := g.Wait(); err != nil {
					return err
				}
				return nil
			}
		})
	}
	return m, dg.Wait()
}

func (gdb *Gdb) getRawHistoricalData(dataType string, rawNamesMap, namesMap map[string][]string) (cmap.ConcurrentMap, error) {
	m := cmap.New()
	dg := errgroup.Group{}
	switch dataType {
	case "float32":
		goto float32
	case "int32":
		goto int32
	case "string":
		goto string32
	case "bool":
		goto bool32
	default:
		return nil, fmt.Errorf("unknown dataType " + dataType)
	}
float32:
	for groupName, groupItemNames := range namesMap {
		groupName := groupName
		groupItemNames := groupItemNames
		rawItemNames := rawNamesMap[groupName]
		dg.Go(func() error {
			if sn, err := gdb.hisDb["float32"][groupName].GetSnapshot(); err != nil {
				return err
			} else {
				defer sn.Release()
				g := errgroup.Group{}
				for i := 0; i < len(groupItemNames); i++ {
					index := i
					g.Go(func() error {
						md, ok := gdb.floatRmHisDb.Get(groupItemNames[index])
						var leveldbHis, memHis []float32
						var leveldbTs, memTs []int32
						g1 := errgroup.Group{}
						g1.Go(func() error {
							it := sn.NewIterator(util.BytesPrefix(convertStringToByte(groupItemNames[index])), nil)
							defer it.Release()
							for it.Next() {
								d := &pb.FloatHistoricalData{}
								iv := it.Value()
								if len(iv) == 0 {
									continue
								}
								if err := proto.Unmarshal(iv, d); err != nil {
									return err
								}
								ts, values := d.GetTimeStamps(), d.GetValues()
								for k := 0; k < len(ts); k++ {
									leveldbTs = append(leveldbTs, ts[k])
									leveldbHis = append(leveldbHis, values[k])
								}
							}
							return nil
						})
						if ok {
							// get history in memory
							g1.Go(func() error {
								for k := 0; k < len(md.TimeStamps); k++ {
									memTs = append(memTs, md.TimeStamps[k])
									memHis = append(memHis, md.Values[k])
								}
								return nil
							})
						}
						if err := g1.Wait(); err != nil {
							return err
						}
						leveldbHis = append(leveldbHis, memHis...)
						leveldbTs = append(leveldbTs, memTs...)
						m.Set(rawItemNames[index], []interface{}{leveldbTs, leveldbHis})
						return nil
					})
				}
				if err := g.Wait(); err != nil {
					return err
				}
				return nil
			}
		})
	}
	return m, dg.Wait()
int32:
	for groupName, groupItemNames := range namesMap {
		groupName := groupName
		groupItemNames := groupItemNames
		rawItemNames := rawNamesMap[groupName]
		dg.Go(func() error {
			if sn, err := gdb.hisDb["int32"][groupName].GetSnapshot(); err != nil {
				return err
			} else {
				defer sn.Release()
				g := errgroup.Group{}
				for i := 0; i < len(groupItemNames); i++ {
					index := i
					g.Go(func() error {
						md, ok := gdb.intRmHisDb.Get(groupItemNames[index])
						var leveldbHis, memHis []int32
						var leveldbTs, memTs []int32
						g1 := errgroup.Group{}
						g1.Go(func() error {
							it := sn.NewIterator(util.BytesPrefix(convertStringToByte(groupItemNames[index])), nil)
							defer it.Release()
							for it.Next() {
								d := &pb.IntHistoricalData{}
								iv := it.Value()
								if len(iv) == 0 {
									continue
								}
								if err := proto.Unmarshal(iv, d); err != nil {
									return err
								}
								ts, values := d.GetTimeStamps(), d.GetValues()
								for k := 0; k < len(ts); k++ {
									leveldbTs = append(leveldbTs, ts[k])
									leveldbHis = append(leveldbHis, values[k])
								}
							}
							return nil
						})
						if ok {
							// get history in memory
							g1.Go(func() error {
								for k := 0; k < len(md.TimeStamps); k++ {
									memTs = append(memTs, md.TimeStamps[k])
									memHis = append(memHis, md.Values[k])
								}
								return nil
							})
						}
						if err := g1.Wait(); err != nil {
							return err
						}
						leveldbHis = append(leveldbHis, memHis...)
						leveldbTs = append(leveldbTs, memTs...)
						m.Set(rawItemNames[index], []interface{}{leveldbTs, leveldbHis})
						return nil
					})
				}
				if err := g.Wait(); err != nil {
					return err
				}
				return nil
			}
		})
	}
	return m, dg.Wait()
string32:
	for groupName, groupItemNames := range namesMap {
		groupName := groupName
		groupItemNames := groupItemNames
		rawItemNames := rawNamesMap[groupName]
		dg.Go(func() error {
			if sn, err := gdb.hisDb["string"][groupName].GetSnapshot(); err != nil {
				return err
			} else {
				defer sn.Release()
				g := errgroup.Group{}
				for i := 0; i < len(groupItemNames); i++ {
					index := i
					g.Go(func() error {
						md, ok := gdb.stringRmHisDb.Get(groupItemNames[index])
						var leveldbHis, memHis []string
						var leveldbTs, memTs []int32
						g1 := errgroup.Group{}
						g1.Go(func() error {
							it := sn.NewIterator(util.BytesPrefix(convertStringToByte(groupItemNames[index])), nil)
							defer it.Release()
							for it.Next() {
								d := &pb.StringHistoricalData{}
								iv := it.Value()
								if len(iv) == 0 {
									continue
								}
								if err := proto.Unmarshal(iv, d); err != nil {
									return err
								}
								ts, values := d.GetTimeStamps(), d.GetValues()
								for k := 0; k < len(ts); k++ {
									leveldbTs = append(leveldbTs, ts[k])
									leveldbHis = append(leveldbHis, values[k])
								}
							}
							return nil
						})
						if ok {
							// get history in memory
							g1.Go(func() error {
								for k := 0; k < len(md.TimeStamps); k++ {
									memTs = append(memTs, md.TimeStamps[k])
									memHis = append(memHis, md.Values[k])
								}
								return nil
							})
						}
						if err := g1.Wait(); err != nil {
							return err
						}
						leveldbHis = append(leveldbHis, memHis...)
						leveldbTs = append(leveldbTs, memTs...)
						m.Set(rawItemNames[index], []interface{}{leveldbTs, leveldbHis})
						return nil
					})
				}
				if err := g.Wait(); err != nil {
					return err
				}
				return nil
			}
		})
	}
	return m, dg.Wait()
bool32:
	for groupName, groupItemNames := range namesMap {
		groupName := groupName
		groupItemNames := groupItemNames
		rawItemNames := rawNamesMap[groupName]
		dg.Go(func() error {
			if sn, err := gdb.hisDb["bool"][groupName].GetSnapshot(); err != nil {
				return err
			} else {
				defer sn.Release()
				g := errgroup.Group{}
				for i := 0; i < len(groupItemNames); i++ {
					index := i
					g.Go(func() error {
						md, ok := gdb.boolRmHisDb.Get(groupItemNames[index])
						var leveldbHis, memHis []bool
						var leveldbTs, memTs []int32
						g1 := errgroup.Group{}
						g1.Go(func() error {
							it := sn.NewIterator(util.BytesPrefix(convertStringToByte(groupItemNames[index])), nil)
							defer it.Release()
							for it.Next() {
								d := &pb.BoolHistoricalData{}
								iv := it.Value()
								if len(iv) == 0 {
									continue
								}
								if err := proto.Unmarshal(iv, d); err != nil {
									return err
								}
								ts, values := d.GetTimeStamps(), d.GetValues()
								for k := 0; k < len(ts); k++ {
									leveldbTs = append(leveldbTs, ts[k])
									leveldbHis = append(leveldbHis, values[k])
								}
							}
							return nil
						})
						if ok {
							// get history in memory
							g1.Go(func() error {
								for k := 0; k < len(md.TimeStamps); k++ {
									memTs = append(memTs, md.TimeStamps[k])
									memHis = append(memHis, md.Values[k])
								}
								return nil
							})
						}
						if err := g1.Wait(); err != nil {
							return err
						}
						leveldbHis = append(leveldbHis, memHis...)
						leveldbTs = append(leveldbTs, memTs...)
						m.Set(rawItemNames[index], []interface{}{leveldbTs, leveldbHis})
						return nil
					})
				}
				if err := g.Wait(); err != nil {
					return err
				}
				return nil
			}
		})
	}
	return m, dg.Wait()
}

func (gdb *Gdb) getHistoricalDataWithMinLength(dataType string, rawNamesMap, namesMap map[string][]string, startTimes, endTimes, intervals []int32) ([]map[string]interface{}, error) {
	itemNames := []string{}
	for _, groupItemNames := range rawNamesMap {
		itemNames = append(itemNames, groupItemNames...)
	}
	m, err := gdb.getHistoricalData(dataType, rawNamesMap, namesMap, startTimes, endTimes, intervals, true)
	if err != nil {
		return nil, err
	}
	result := []map[string]interface{}{}
	floatItemValues := map[string]memap.RmHisDbFloatItems{}   // historical values
	intItemValues := map[string]memap.RmHisDbIntItems{}       // historical values
	stringItemValues := map[string]memap.RmHisDbStringItems{} // historical values
	boolItemValues := map[string]memap.RmHisDbBoolItems{}     // historical values
	l, _ := m.Get("minLength")
	minLength := l.(int)
	switch dataType {
	case "float32":
		goto float32
	case "int32":
		goto int32
	case "string":
		goto string32
	case "bool":
		goto bool32
	default:
		return nil, fmt.Errorf("unknown dataType " + dataType)
	}
float32:
	for _, itemName := range itemNames {
		v, _ := m.Get(itemName) // get history data of item
		vs := v.([]interface{}) // convert historicalData
		floatItemValues[itemName] = memap.RmHisDbFloatItems{TimeStamps: vs[0].([]int32), Values: vs[1].([]float32)}
	}
	for i := 0; i < minLength; i++ {
		t := map[string]interface{}{}
		if len(itemNames) == 1 {
			t["timeStamp"] = floatItemValues[itemNames[0]].TimeStamps[i] // ts
			t[itemNames[0]] = floatItemValues[itemNames[0]].Values[i]
		} else {
			for index, itemName := range itemNames {
				if index == 0 {
					// first
					t["timeStamp"] = floatItemValues[itemName].TimeStamps[i] // ts
					t[itemName] = floatItemValues[itemName].Values[i]        // value
				} else {
					if t["timeStamp"] != floatItemValues[itemName].TimeStamps[i] {
						// inconsistent timeStamp
						break
					} else {
						t[itemName] = floatItemValues[itemName].Values[i] // value
					}
				}
			}
		}
		result = append(result, t)
	}
	return result, nil
int32:
	for _, itemName := range itemNames {
		v, _ := m.Get(itemName) // get history data of item
		vs := v.([]interface{}) // convert historicalData
		intItemValues[itemName] = memap.RmHisDbIntItems{TimeStamps: vs[0].([]int32), Values: vs[1].([]int32)}
	}
	for i := 0; i < minLength; i++ {
		t := map[string]interface{}{}
		for index, itemName := range itemNames {
			if index == 0 {
				// first
				t["timeStamp"] = intItemValues[itemName].TimeStamps[i] // ts
				t[itemName] = floatItemValues[itemName].Values[i]      // value
			} else {
				if t["timeStamp"] != intItemValues[itemName].TimeStamps[i] {
					// inconsistent timeStamp
					break
				} else {
					t[itemName] = intItemValues[itemName].Values[i] // value
				}
			}
		}
		result = append(result, t)
	}
	return result, nil
string32:
	for _, itemName := range itemNames {
		v, _ := m.Get(itemName) // get history data of item
		vs := v.([]interface{}) // convert historicalData
		stringItemValues[itemName] = memap.RmHisDbStringItems{TimeStamps: vs[0].([]int32), Values: vs[1].([]string)}
	}
	for i := 0; i < minLength; i++ {
		t := map[string]interface{}{}
		for index, itemName := range itemNames {
			if index == 0 {
				// first
				t["timeStamp"] = stringItemValues[itemName].TimeStamps[i] // ts
				t[itemName] = floatItemValues[itemName].Values[i]         // value
			} else {
				if t["timeStamp"] != stringItemValues[itemName].TimeStamps[i] {
					// inconsistent timeStamp
					break
				} else {
					t[itemName] = stringItemValues[itemName].Values[i] // value
				}
			}
		}
		result = append(result, t)
	}
	return result, nil
bool32:
	for _, itemName := range itemNames {
		v, _ := m.Get(itemName) // get history data of item
		vs := v.([]interface{}) // convert historicalData
		boolItemValues[itemName] = memap.RmHisDbBoolItems{TimeStamps: vs[0].([]int32), Values: vs[1].([]bool)}
	}
	for i := 0; i < minLength; i++ {
		t := map[string]interface{}{}
		for index, itemName := range itemNames {
			if index == 0 {
				// first
				t["timeStamp"] = boolItemValues[itemName].TimeStamps[i] // ts
				t[itemName] = floatItemValues[itemName].Values[i]       // value
			} else {
				if t["timeStamp"] != boolItemValues[itemName].TimeStamps[i] {
					// inconsistent timeStamp
					break
				} else {
					t[itemName] = boolItemValues[itemName].Values[i] // value
				}
			}
		}
		result = append(result, t)
	}
	return result, nil
}

func (gdb *Gdb) getHistoricalDataWithStampAndDeadZoneCount(dataType string, rawItemNames []string, rawNamesMap, namesMap map[string][]string, timeStamps [][]int32, zones []DeadZone) (cmap.ConcurrentMap, error) {
	m := cmap.New()
	g := errgroup.Group{}
	dd, err := gdb.getHistoricalDataWithTs(dataType, rawNamesMap, namesMap, timeStamps...)
	if err != nil {
		return nil, err
	}
	switch dataType {
	case "float32":
		goto float32
	case "int32":
		goto int32
	case "string":
		goto string32
	case "bool":
		goto bool32
	default:
		return nil, fmt.Errorf("unknown dataType " + dataType)
	}
float32:
	for i := 0; i < len(rawItemNames); i++ {
		name := rawItemNames[i]
		g.Go(func() error {
			filterIndex := From(zones).IndexOf(func(item interface{}) bool {
				return item.(DeadZone).ItemName == name
			})
			var deadZoneCounts int32
			if filterIndex == -1 {
				// don't filter
				goto doNotFilter
			}
			deadZoneCounts = zones[filterIndex].DeadZoneCount
			if deadZoneCounts < 1 {
				goto doNotFilter
			} else {
				// get filter data
				var values, lastValues []float32
				var lastValue float32
				var tts []int32
				tLastValues := []int32{}
				data, _ := dd.Get(name)                   // get data
				ts := data.([]interface{})[0].([]int32)   // ts
				hs := data.([]interface{})[1].([]float32) // values
				for k := 0; k < len(ts); k++ {
					vs := hs[k] // value
					if k == 0 {
						// first
						for index, lv := range lastValues {
							values = append(values, lv)
							tts = append(tts, tLastValues[index])
						}
						lastValue = vs
						lastValues = []float32{}
						tLastValues = []int32{} // timeStamps
						values = append(values, vs)
						tts = append(tts, ts[k])
					} else {
						if lastValue != vs {
							// not repeated
							for index, lv := range lastValues {
								values = append(values, lv)
								tts = append(tts, tLastValues[index])
							}
							lastValue = vs
							lastValues = []float32{}
							tLastValues = []int32{}
							values = append(values, vs)
							tts = append(tts, ts[k])
						} else {
							// repeated
							lastValues = append(lastValues, lastValue)
							tLastValues = append(tLastValues, ts[k])
							if len(lastValues) == int(deadZoneCounts) {
								values = append(values, lastValues[:int(deadZoneCounts)-1]...)
								tts = append(tts, tLastValues[:int(deadZoneCounts)-1]...)
								lastValues = []float32{}
								tLastValues = []int32{}
							}
						}
					}
				}
				m.Set(name, []interface{}{tts, values})
				return nil
			}
		doNotFilter:
			v, _ := dd.Get(name)
			m.Set(name, v)
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return nil, err
	}
	return m, nil
int32:
	for i := 0; i < len(rawItemNames); i++ {
		name := rawItemNames[i]
		g.Go(func() error {
			filterIndex := From(zones).IndexOf(func(item interface{}) bool {
				return item.(DeadZone).ItemName == name
			})
			var deadZoneCounts int32
			if filterIndex == -1 {
				// don't filter
				goto doNotFilter
			}
			deadZoneCounts = zones[filterIndex].DeadZoneCount
			if deadZoneCounts < 1 {
				// don't filter
				goto doNotFilter
			} else {
				// get filter data
				var values, lastValues []int32
				var lastValue int32
				var tts []int32
				tLastValues := []int32{}
				data, _ := dd.Get(name)                 // get data
				ts := data.([]interface{})[0].([]int32) // ts
				hs := data.([]interface{})[1].([]int32) // values
				for k := 0; k < len(ts); k++ {
					vs := hs[k] // value
					if k == 0 {
						// first
						for index, lv := range lastValues {
							values = append(values, lv)
							tts = append(tts, tLastValues[index])
						}
						lastValue = vs
						lastValues = []int32{}
						tLastValues = []int32{} // timeStamps
						values = append(values, vs)
						tts = append(tts, ts[k])
					} else {
						if lastValue != vs {
							// not repeated
							for index, lv := range lastValues {
								values = append(values, lv)
								tts = append(tts, tLastValues[index])
							}
							lastValue = vs
							lastValues = []int32{}
							tLastValues = []int32{}
							values = append(values, vs)
							tts = append(tts, ts[k])
						} else {
							// repeated
							lastValues = append(lastValues, lastValue)
							tLastValues = append(tLastValues, ts[k])
							if len(lastValues) == int(deadZoneCounts) {
								values = append(values, lastValues[:int(deadZoneCounts)-1]...)
								tts = append(tts, tLastValues[:int(deadZoneCounts)-1]...)
								lastValues = []int32{}
								tLastValues = []int32{}
							}
						}
					}
				}
				m.Set(name, []interface{}{tts, values})
				return nil
			}
		doNotFilter:
			v, _ := dd.Get(name)
			m.Set(name, v)
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return nil, err
	}
	return m, nil
string32:
	for i := 0; i < len(rawItemNames); i++ {
		name := rawItemNames[i]
		g.Go(func() error {
			filterIndex := From(zones).IndexOf(func(item interface{}) bool {
				return item.(DeadZone).ItemName == name
			})
			var deadZoneCounts int32
			if filterIndex == -1 {
				// don't filter
				goto doNotFilter
			}
			deadZoneCounts = zones[filterIndex].DeadZoneCount
			if deadZoneCounts < 1 {
				// don't filter
				goto doNotFilter
			} else {
				// get filter data
				var values, lastValues []string
				var lastValue string
				var tts []int32
				tLastValues := []int32{}
				data, _ := dd.Get(name)                  // get data
				ts := data.([]interface{})[0].([]int32)  // ts
				hs := data.([]interface{})[1].([]string) // values
				for k := 0; k < len(ts); k++ {
					vs := hs[k] // value
					if k == 0 {
						// first
						for index, lv := range lastValues {
							values = append(values, lv)
							tts = append(tts, tLastValues[index])
						}
						lastValue = vs
						lastValues = []string{}
						tLastValues = []int32{} // timeStamps
						values = append(values, vs)
						tts = append(tts, ts[k])
					} else {
						if lastValue != vs {
							// not repeated
							for index, lv := range lastValues {
								values = append(values, lv)
								tts = append(tts, tLastValues[index])
							}
							lastValue = vs
							lastValues = []string{}
							tLastValues = []int32{}
							values = append(values, vs)
							tts = append(tts, ts[k])
						} else {
							// repeated
							lastValues = append(lastValues, lastValue)
							tLastValues = append(tLastValues, ts[k])
							if len(lastValues) == int(deadZoneCounts) {
								values = append(values, lastValues[:int(deadZoneCounts)-1]...)
								tts = append(tts, tLastValues[:int(deadZoneCounts)-1]...)
								lastValues = []string{}
								tLastValues = []int32{}
							}
						}
					}
				}
				m.Set(name, []interface{}{tts, values})
				return nil
			}
		doNotFilter:
			v, _ := dd.Get(name)
			m.Set(name, v)
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return nil, err
	}
	return m, nil
bool32:
	for i := 0; i < len(rawItemNames); i++ {
		name := rawItemNames[i]
		g.Go(func() error {
			filterIndex := From(zones).IndexOf(func(item interface{}) bool {
				return item.(DeadZone).ItemName == name
			})
			var deadZoneCounts int32
			if filterIndex == -1 {
				// don't filter
				goto doNotFilter
			}
			deadZoneCounts = zones[filterIndex].DeadZoneCount
			if deadZoneCounts < 1 {
				// don't filter
				goto doNotFilter
			} else {
				// get filter data
				var values, lastValues []bool
				var lastValue bool
				var tts []int32
				tLastValues := []int32{}
				data, _ := dd.Get(name)                 // get data
				ts := data.([]interface{})[0].([]int32) // ts
				hs := data.([]interface{})[1].([]bool)  // values
				for k := 0; k < len(ts); k++ {
					vs := hs[k] // value
					if k == 0 {
						// first
						for index, lv := range lastValues {
							values = append(values, lv)
							tts = append(tts, tLastValues[index])
						}
						lastValue = vs
						lastValues = []bool{}
						tLastValues = []int32{} // timeStamps
						values = append(values, vs)
						tts = append(tts, ts[k])
					} else {
						if lastValue != vs {
							// not repeated
							for index, lv := range lastValues {
								values = append(values, lv)
								tts = append(tts, tLastValues[index])
							}
							lastValue = vs
							lastValues = []bool{}
							tLastValues = []int32{}
							values = append(values, vs)
							tts = append(tts, ts[k])
						} else {
							// repeated
							lastValues = append(lastValues, lastValue)
							tLastValues = append(tLastValues, ts[k])
							if len(lastValues) == int(deadZoneCounts) {
								values = append(values, lastValues[:int(deadZoneCounts)-1]...)
								tts = append(tts, tLastValues[:int(deadZoneCounts)-1]...)
								lastValues = []bool{}
								tLastValues = []int32{}
							}
						}
					}
				}
				m.Set(name, []interface{}{tts, values})
				return nil
			}
		doNotFilter:
			v, _ := dd.Get(name)
			m.Set(name, v)
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return nil, err
	}
	return m, nil
}

func (gdb *Gdb) getHistoricalDataWithCondition(dataType, groupName string, itemNames []string, startTimes, endTimes, intervals []int32, filterCondition string, zones ...DeadZone) (cmap.ConcurrentMap, error) {
	if rawNamesMap, namesMap, err := gdb.checkSingleItemDataTypeWithMap(strings.Split(strings.TrimRight(strings.Repeat(groupName+",", len(itemNames)), ","), ","), itemNames, dataType); err != nil {
		return nil, err
	} else {
		// get filterItemNames
		filterItemNames := []string{}
		fn := mapset.NewSet()
		if strings.TrimSpace(filterCondition) == "true" {
			filterItemNames = itemNames
		} else {
			reg := regexp.MustCompile(`(?is)\["(.*?)"]`)
			if reg.Match(convertStringToByte(filterCondition)) {
				matchedResult := reg.FindAllStringSubmatch(filterCondition, -1)
				for _, mr := range matchedResult {
					fn.Add(mr[1])
					//filterItemNames = append(filterItemNames, mr[1])
				}
			} else {
				return nil, fmt.Errorf("invalid condition, items must be included by [] and condition can't be null")
			}
		}
		for item := range fn.Iter() {
			filterItemNames = append(filterItemNames, item.(string))
		}
		// check itemName in zone condition
		if len(zones) != 0 {
			for _, zone := range zones {
				if From(itemNames).IndexOf(func(item interface{}) bool {
					r := item.(string)
					return r == zone.ItemName
				}) < 0 {
					return nil, fmt.Errorf("zone error, item " + zone.ItemName + " not exist")
				}
			}
		}
		// check filterItemNames
		if rawFilterNamesMap, filterNamesMap, err := gdb.checkSingleItemDataTypeWithMap(strings.Split(strings.TrimRight(strings.Repeat(groupName+",", len(filterItemNames)), ","), ","), filterItemNames, dataType); err != nil {
			return nil, err
		} else {
			// pass check
			if filterHistoryData, err := gdb.getHistoricalDataWithMinLength(dataType, rawFilterNamesMap, filterNamesMap, startTimes, endTimes, intervals); err != nil {
				return nil, err
			} else {
				vm := goja.New()
				f := `function filterData(s){return s.filter(function(item){return ` + filterCondition + `})}`
				if _, err := vm.RunString(f); err != nil {
					return nil, fmt.Errorf("fail to compiling condition: " + err.Error())
				}
				if filterData, ok := goja.AssertFunction(vm.Get("filterData")); !ok {
					return nil, fmt.Errorf("fail to compiling function")
				} else {
					if res, err := filterData(goja.Undefined(), vm.ToValue(filterHistoryData)); err != nil {
						return nil, fmt.Errorf("fail to running condition: " + err.Error())
					} else {
						filterResults := res.Export().([]interface{})
						timeStamps := []int32{}
						tts := make([][]int32, len(itemNames))
						// get timeStamps of data
						for _, fr := range filterResults {
							sfr := fr.(map[string]interface{})
							timeStamps = append(timeStamps, sfr["timeStamp"].(int32))
						}
						for i := 0; i < len(itemNames); i++ {
							tts[i] = timeStamps
						}
						if len(zones) == 0 {
							// get data without zones
							if len(timeStamps) == 0 {
								return cmap.New(), nil
							}
							if m, err := gdb.getHistoricalDataWithTs(dataType, rawNamesMap, namesMap, tts...); err != nil {
								return nil, err
							} else {
								return m, nil
							}
						} else {
							// get history data with zones
							if m, err := gdb.getHistoricalDataWithStampAndDeadZoneCount(dataType, itemNames, rawNamesMap, namesMap, tts, zones); err != nil {
								return nil, err
							} else {
								return m, nil
							}
						}
					}
				}
			}
		}
	}
}

func (gdb *Gdb) deleteHistoricalData(dataType string, namesMap map[string][]string, startTimes, endTimes []int32) (int, error) {
	dg := errgroup.Group{}
	rows := [3]int{}
	itemNames := []string{}
	for _, groupItemNames := range namesMap {
		itemNames = append(itemNames, groupItemNames...)
	}
	switch dataType {
	case "float32":
		goto float32
	case "int32":
		goto int32
	case "string":
		goto string32
	case "bool":
		goto bool32
	default:
		return -1, fmt.Errorf("unknown dataType " + dataType)
	}
float32:
	// delete history in leveldb
	dg.Go(func() error {
		g := errgroup.Group{}
		innerRows := sync.Map{}
		for groupName, groupItemNames := range namesMap {
			groupName := groupName
			groupItemNames := groupItemNames
			g.Go(func() error {
				if sn, err := gdb.hisDb["float32"][groupName].GetSnapshot(); err != nil {
					return err
				} else {
					defer sn.Release()
					batch := &leveldb.Batch{}
					row := 0
					for index := range groupItemNames {
						index := index
						itemName := groupItemNames[index]
						for j := 0; j < len(startTimes); j++ {
							if startTimes[j] > endTimes[j] && startTimes[j] != -1 && endTimes[j] != -1 {
								continue
							}
							var s, e string
							var it iterator.Iterator
							if startTimes[j] == -1 && endTimes[j] == -1 {
								// delete all history data
								it = sn.NewIterator(util.BytesPrefix(convertStringToByte(itemName)), nil)
								goto deleteAllHistory
							} else if startTimes[j] == -1 && endTimes[j] != -1 {
								// delete history data ==> (,et)
								s = itemName
								e = itemName + joiner + strconv.Itoa(int(endTimes[j])+int(gdb.hisTimeDuration.Seconds()))
								it = sn.NewIterator(&util.Range{Start: convertStringToByte(s), Limit: convertStringToByte(e)}, nil)
								goto deleteEndHistory
							} else if startTimes[j] != -1 && endTimes[j] == -1 {
								// delete history data ==> [st,)
								s = itemName + joiner + strconv.Itoa(int(startTimes[j])-int(gdb.hisTimeDuration.Seconds()))
								e = itemName + joiner + strconv.Itoa(int(time.Now().Unix()+8*3600)+int(gdb.hisTimeDuration.Seconds()))
								it = sn.NewIterator(&util.Range{Start: convertStringToByte(s), Limit: convertStringToByte(e)}, nil)
								goto deleteStartHistory
							} else {
								// delete history data ==> (st,et]
								s := itemName + joiner + strconv.Itoa(int(startTimes[j])-int(gdb.hisTimeDuration.Seconds()))
								e := itemName + joiner + strconv.Itoa(int(endTimes[j])+int(gdb.hisTimeDuration.Seconds()))
								it = sn.NewIterator(&util.Range{Start: convertStringToByte(s), Limit: convertStringToByte(e)}, nil)
								goto deletePartHistory
							}
						deleteAllHistory:
							for it.Next() {
								batch.Put(it.Key(), nil)
							}
							it.Release()
						deleteEndHistory:
							for it.Next() {
								ts, err := strconv.ParseInt(strings.Split(string(it.Key()), joiner)[1], 10, 64) // get ts of it.Value
								{
									if err != nil {
										return err
									}
									if ts+int64(gdb.hisTimeDuration.Seconds()) <= int64(endTimes[j]) {
										// all history in it.Value needed to delete
										iv := it.Value()
										if len(iv) == 0 {
											continue
										}
										d := &pb.FloatHistoricalData{}
										if err := proto.Unmarshal(iv, d); err != nil {
											return err
										}
										row += len(d.TimeStamps)
										batch.Put(it.Key(), nil)
									} else {
										d := &pb.FloatHistoricalData{}
										iv := it.Value()
										if len(iv) == 0 {
											continue
										}
										if err := proto.Unmarshal(iv, d); err != nil {
											return err
										}
										t, v := d.GetTimeStamps(), d.GetValues()
										tts := []int32{}
										vs := []float32{}
										if k := getEndIndex(endTimes[j], t); k == -1 {
											continue
										} else {
											row += k
											if k == len(endTimes) {
												// all data to be deleted
												batch.Put(it.Key(), nil)
											} else {
												tts = append(tts, t[k:]...)
												vs = append(vs, v[k:]...)
												dn := &pb.FloatHistoricalData{TimeStamps: tts, Values: vs}
												if data, err := proto.Marshal(dn); err != nil {
													return err
												} else {
													batch.Put(it.Key(), data)
												}
											}
										}
									}
								}
							}
							it.Release()
						deleteStartHistory:
							for it.Next() {
								ts, err := strconv.ParseInt(strings.Split(string(it.Key()), joiner)[1], 10, 64)
								{
									if err != nil {
										return err
									}
									if ts >= int64(startTimes[j]) {
										iv := it.Value()
										if len(iv) == 0 {
											continue
										}
										d := &pb.FloatHistoricalData{}
										if err := proto.Unmarshal(iv, d); err != nil {
											return err
										}
										row += len(d.TimeStamps)
										batch.Put(it.Key(), nil)
									} else {
										d := &pb.FloatHistoricalData{}
										iv := it.Value()
										if len(iv) == 0 {
											continue
										}
										if err := proto.Unmarshal(iv, d); err != nil {
											return err
										}
										t, v := d.GetTimeStamps(), d.GetValues()
										tts := []int32{}
										vs := []float32{}
										if k := getStartIndex(startTimes[j], t); k == -1 {
											continue
										} else {
											row += len(t) - k
											if k == 0 {
												// all history to be deleted
												batch.Put(it.Key(), nil)
											} else {
												tts = append(tts, t[:k]...)
												vs = append(vs, v[:k]...)
												dn := &pb.FloatHistoricalData{
													TimeStamps: tts,
													Values:     vs,
												}
												if data, err := proto.Marshal(dn); err != nil {
													return err
												} else {
													batch.Put(it.Key(), data)
												}
											}
										}
									}
								}
							}
							it.Release()
						deletePartHistory:
							for it.Next() {
								d := &pb.FloatHistoricalData{}
								iv := it.Value()
								if len(iv) == 0 {
									continue
								}
								if err := proto.Unmarshal(iv, d); err != nil {
									return err
								}
								t, v := d.GetTimeStamps(), d.GetValues()
								ts := []int32{}
								vs := []float32{}
								startIndex, endIndex := getIndex(startTimes[j], endTimes[j], t)
								{
									if startIndex == -1 || endIndex == -1 {
										continue
									}
									row += endIndex - startIndex
									if endIndex-startIndex == len(t) {
										// all history to be deleted
										batch.Put(it.Key(), nil)
									} else {
										ts = append(ts, t[:startIndex]...)
										ts = append(ts, t[endIndex:]...)
										vs = append(vs, v[:startIndex]...)
										vs = append(vs, v[endIndex:]...)
										if len(ts) == 0 {
											batch.Put(it.Key(), nil)
										} else {
											dn := &pb.FloatHistoricalData{TimeStamps: ts, Values: vs}
											if data, err := proto.Marshal(dn); err != nil {
												return err
											} else {
												batch.Put(it.Key(), data)
											}
										}
									}
								}
							}
							it.Release()
						}
					}
					innerRows.Store(groupName, row)
					return gdb.hisDb["float32"][groupName].Write(batch, nil)
				}
			})
		}
		if err := g.Wait(); err != nil {
			return err
		}
		for groupName := range namesMap {
			row, _ := innerRows.Load(groupName)
			{
				rows[0] += row.(int)
			}
		}
		return nil
	})
	// delete history in memory
	dg.Go(func() error {
		for _, itemName := range itemNames {
			if md, ok := gdb.floatRmHisDb.Get(itemName); ok {
				// delete history data in memory
				if len(md.TimeStamps) != 0 {
					row := 0
					for j := 0; j < len(startTimes); j++ {
						tts := []int32{}
						vs := []float32{}
						if startTimes[j] == -1 && endTimes[j] == -1 {
							goto deleteAllHistory
						} else if startTimes[j] == -1 && endTimes[j] != -1 {
							goto deleteEndHistory
						} else if startTimes[j] != -1 && endTimes[j] == -1 {
							goto deleteStartHistory
						} else {
							goto deletePartHistory
						}
					deleteAllHistory:
						md.TimeStamps = []int32{}
						md.Values = []float32{}
						row += len(md.TimeStamps)
					deleteEndHistory:
						if k := getEndIndex(endTimes[j], md.TimeStamps); k == -1 {
							continue
						} else {
							md.TimeStamps = md.TimeStamps[k:]
							md.Values = md.Values[k:]
							row += k
						}
					deleteStartHistory:
						if k := getStartIndex(startTimes[j], md.TimeStamps); k == -1 {
							continue
						} else {
							md.TimeStamps = md.TimeStamps[:k]
							md.Values = md.Values[:k]
							row += len(md.TimeStamps) - k
						}
					deletePartHistory:
						startIndex, endIndex := getIndex(startTimes[j], endTimes[j], md.TimeStamps)
						{
							if startIndex == -1 || endIndex == -1 {
								continue
							}
							row += endIndex - startIndex
							tts = append(tts, md.TimeStamps[:startIndex]...)
							tts = append(tts, md.TimeStamps[endIndex:]...)
							vs = append(vs, md.Values[:startIndex]...)
							vs = append(vs, md.Values[endIndex:]...)
							md.TimeStamps, md.Values = tts, vs
						}
					}
					rows[1] = row
				}
				gdb.floatRmHisDb.Set(itemName, md)
			}
		}
		return nil
	})
	if err := dg.Wait(); err != nil {
		return -1, err
	} else {
		return rows[0] + rows[1] + rows[2], nil
	}
int32:
	// delete history in leveldb
	dg.Go(func() error {
		g := errgroup.Group{}
		innerRows := sync.Map{}
		for groupName, groupItemNames := range namesMap {
			groupName := groupName
			groupItemNames := groupItemNames
			g.Go(func() error {
				if sn, err := gdb.hisDb["int32"][groupName].GetSnapshot(); err != nil {
					return err
				} else {
					defer sn.Release()
					batch := &leveldb.Batch{}
					row := 0
					for index := range groupItemNames {
						index := index
						itemName := groupItemNames[index]
						for j := 0; j < len(startTimes); j++ {
							if startTimes[j] > endTimes[j] && startTimes[j] != -1 && endTimes[j] != -1 {
								continue
							}
							var s, e string
							var it iterator.Iterator
							if startTimes[j] == -1 && endTimes[j] == -1 {
								// delete all history data
								it = sn.NewIterator(util.BytesPrefix(convertStringToByte(itemName)), nil)
								goto deleteAllHistory
							} else if startTimes[j] == -1 && endTimes[j] != -1 {
								// delete history data ==> (,et)
								s = itemName
								e = itemName + joiner + strconv.Itoa(int(endTimes[j])+int(gdb.hisTimeDuration.Seconds()))
								it = sn.NewIterator(&util.Range{Start: convertStringToByte(s), Limit: convertStringToByte(e)}, nil)
								goto deleteEndHistory
							} else if startTimes[j] != -1 && endTimes[j] == -1 {
								// delete history data ==> [st,)
								s = itemName + joiner + strconv.Itoa(int(startTimes[j])-int(gdb.hisTimeDuration.Seconds()))
								e = itemName + joiner + strconv.Itoa(int(time.Now().Unix()+8*3600)+int(gdb.hisTimeDuration.Seconds()))
								it = sn.NewIterator(&util.Range{Start: convertStringToByte(s), Limit: convertStringToByte(e)}, nil)
								goto deleteStartHistory
							} else {
								// delete history data ==> (st,et]
								s := itemName + joiner + strconv.Itoa(int(startTimes[j])-int(gdb.hisTimeDuration.Seconds()))
								e := itemName + joiner + strconv.Itoa(int(endTimes[j])+int(gdb.hisTimeDuration.Seconds()))
								it = sn.NewIterator(&util.Range{Start: convertStringToByte(s), Limit: convertStringToByte(e)}, nil)
								goto deletePartHistory
							}
						deleteAllHistory:
							for it.Next() {
								batch.Put(it.Key(), nil)
							}
							it.Release()
						deleteEndHistory:
							for it.Next() {
								ts, err := strconv.ParseInt(strings.Split(string(it.Key()), joiner)[1], 10, 64) // get ts of it.Value
								{
									if err != nil {
										return err
									}
									if ts+int64(gdb.hisTimeDuration.Seconds()) <= int64(endTimes[j]) {
										// all history in it.Value needed to delete
										iv := it.Value()
										if len(iv) == 0 {
											continue
										}
										d := &pb.IntHistoricalData{}
										if err := proto.Unmarshal(iv, d); err != nil {
											return err
										}
										row += len(d.TimeStamps)
										batch.Put(it.Key(), nil)
									} else {
										d := &pb.IntHistoricalData{}
										iv := it.Value()
										if len(iv) == 0 {
											continue
										}
										if err := proto.Unmarshal(iv, d); err != nil {
											return err
										}
										t, v := d.GetTimeStamps(), d.GetValues()
										tts := []int32{}
										vs := []int32{}
										if k := getEndIndex(endTimes[j], t); k == -1 {
											continue
										} else {
											row += k
											if k == len(endTimes) {
												// all data to be deleted
												batch.Put(it.Key(), nil)
											} else {
												tts = append(tts, t[k:]...)
												vs = append(vs, v[k:]...)
												dn := &pb.IntHistoricalData{TimeStamps: tts, Values: vs}
												if data, err := proto.Marshal(dn); err != nil {
													return err
												} else {
													batch.Put(it.Key(), data)
												}
											}
										}
									}
								}
							}
							it.Release()
						deleteStartHistory:
							for it.Next() {
								ts, err := strconv.ParseInt(strings.Split(string(it.Key()), joiner)[1], 10, 64)
								{
									if err != nil {
										return err
									}
									if ts >= int64(startTimes[j]) {
										iv := it.Value()
										if len(iv) == 0 {
											continue
										}
										d := &pb.IntHistoricalData{}
										if err := proto.Unmarshal(iv, d); err != nil {
											return err
										}
										row += len(d.TimeStamps)
										batch.Put(it.Key(), nil)
									} else {
										d := &pb.IntHistoricalData{}
										iv := it.Value()
										if len(iv) == 0 {
											continue
										}
										if err := proto.Unmarshal(iv, d); err != nil {
											return err
										}
										t, v := d.GetTimeStamps(), d.GetValues()
										tts := []int32{}
										vs := []int32{}
										if k := getStartIndex(startTimes[j], t); k == -1 {
											continue
										} else {
											row += len(t) - k
											if k == 0 {
												// all history to be deleted
												batch.Put(it.Key(), nil)
											} else {
												tts = append(tts, t[:k]...)
												vs = append(vs, v[:k]...)
												dn := &pb.IntHistoricalData{
													TimeStamps: tts,
													Values:     vs,
												}
												if data, err := proto.Marshal(dn); err != nil {
													return err
												} else {
													batch.Put(it.Key(), data)
												}
											}
										}
									}
								}
							}
							it.Release()
						deletePartHistory:
							for it.Next() {
								d := &pb.IntHistoricalData{}
								iv := it.Value()
								if len(iv) == 0 {
									continue
								}
								if err := proto.Unmarshal(iv, d); err != nil {
									return err
								}
								t, v := d.GetTimeStamps(), d.GetValues()
								ts := []int32{}
								vs := []int32{}
								startIndex, endIndex := getIndex(startTimes[j], endTimes[j], t)
								{
									if startIndex == -1 || endIndex == -1 {
										continue
									}
									row += endIndex - startIndex
									if endIndex-startIndex == len(t) {
										// all history to be deleted
										batch.Put(it.Key(), nil)
									} else {
										ts = append(ts, t[:startIndex]...)
										ts = append(ts, t[endIndex:]...)
										vs = append(vs, v[:startIndex]...)
										vs = append(vs, v[endIndex:]...)
										if len(ts) == 0 {
											batch.Put(it.Key(), nil)
										} else {
											dn := &pb.IntHistoricalData{TimeStamps: ts, Values: vs}
											if data, err := proto.Marshal(dn); err != nil {
												return err
											} else {
												batch.Put(it.Key(), data)
											}
										}
									}
								}
							}
							it.Release()
						}
					}
					innerRows.Store(groupName, row)
					return gdb.hisDb["float32"][groupName].Write(batch, nil)
				}
			})
		}
		if err := g.Wait(); err != nil {
			return err
		}
		for groupName := range namesMap {
			row, _ := innerRows.Load(groupName)
			{
				rows[0] += row.(int)
			}
		}
		return nil
	})
	// delete history in memory
	dg.Go(func() error {
		for _, itemName := range itemNames {
			if md, ok := gdb.intRmHisDb.Get(itemName); ok {
				// delete history data in memory
				if len(md.TimeStamps) != 0 {
					row := 0
					for j := 0; j < len(startTimes); j++ {
						tts := []int32{}
						vs := []int32{}
						if startTimes[j] == -1 && endTimes[j] == -1 {
							goto deleteAllHistory
						} else if startTimes[j] == -1 && endTimes[j] != -1 {
							goto deleteEndHistory
						} else if startTimes[j] != -1 && endTimes[j] == -1 {
							goto deleteStartHistory
						} else {
							goto deletePartHistory
						}
					deleteAllHistory:
						md.TimeStamps = []int32{}
						md.Values = []int32{}
						row += len(md.TimeStamps)
					deleteEndHistory:
						if k := getEndIndex(endTimes[j], md.TimeStamps); k == -1 {
							continue
						} else {
							md.TimeStamps = md.TimeStamps[k:]
							md.Values = md.Values[k:]
							row += k
						}
					deleteStartHistory:
						if k := getStartIndex(startTimes[j], md.TimeStamps); k == -1 {
							continue
						} else {
							md.TimeStamps = md.TimeStamps[:k]
							md.Values = md.Values[:k]
							row += len(md.TimeStamps) - k
						}
					deletePartHistory:
						startIndex, endIndex := getIndex(startTimes[j], endTimes[j], md.TimeStamps)
						{
							if startIndex == -1 || endIndex == -1 {
								continue
							}
							row += endIndex - startIndex
							tts = append(tts, md.TimeStamps[:startIndex]...)
							tts = append(tts, md.TimeStamps[endIndex:]...)
							vs = append(vs, md.Values[:startIndex]...)
							vs = append(vs, md.Values[endIndex:]...)
							md.TimeStamps, md.Values = tts, vs
						}
					}
					rows[1] = row
				}
				gdb.intRmHisDb.Set(itemName, md)
			}
		}
		return nil
	})
	if err := dg.Wait(); err != nil {
		return -1, err
	} else {
		return rows[0] + rows[1] + rows[2], nil
	}
string32:
	// delete history in leveldb
	dg.Go(func() error {
		g := errgroup.Group{}
		innerRows := sync.Map{}
		for groupName, groupItemNames := range namesMap {
			groupName := groupName
			groupItemNames := groupItemNames
			g.Go(func() error {
				if sn, err := gdb.hisDb["string"][groupName].GetSnapshot(); err != nil {
					return err
				} else {
					defer sn.Release()
					batch := &leveldb.Batch{}
					row := 0
					for index := range groupItemNames {
						index := index
						itemName := groupItemNames[index]
						for j := 0; j < len(startTimes); j++ {
							if startTimes[j] > endTimes[j] && startTimes[j] != -1 && endTimes[j] != -1 {
								continue
							}
							var s, e string
							var it iterator.Iterator
							if startTimes[j] == -1 && endTimes[j] == -1 {
								// delete all history data
								it = sn.NewIterator(util.BytesPrefix(convertStringToByte(itemName)), nil)
								goto deleteAllHistory
							} else if startTimes[j] == -1 && endTimes[j] != -1 {
								// delete history data ==> (,et)
								s = itemName
								e = itemName + joiner + strconv.Itoa(int(endTimes[j])+int(gdb.hisTimeDuration.Seconds()))
								it = sn.NewIterator(&util.Range{Start: convertStringToByte(s), Limit: convertStringToByte(e)}, nil)
								goto deleteEndHistory
							} else if startTimes[j] != -1 && endTimes[j] == -1 {
								// delete history data ==> [st,)
								s = itemName + joiner + strconv.Itoa(int(startTimes[j])-int(gdb.hisTimeDuration.Seconds()))
								e = itemName + joiner + strconv.Itoa(int(time.Now().Unix()+8*3600)+int(gdb.hisTimeDuration.Seconds()))
								it = sn.NewIterator(&util.Range{Start: convertStringToByte(s), Limit: convertStringToByte(e)}, nil)
								goto deleteStartHistory
							} else {
								// delete history data ==> (st,et]
								s := itemName + joiner + strconv.Itoa(int(startTimes[j])-int(gdb.hisTimeDuration.Seconds()))
								e := itemName + joiner + strconv.Itoa(int(endTimes[j])+int(gdb.hisTimeDuration.Seconds()))
								it = sn.NewIterator(&util.Range{Start: convertStringToByte(s), Limit: convertStringToByte(e)}, nil)
								goto deletePartHistory
							}
						deleteAllHistory:
							for it.Next() {
								batch.Put(it.Key(), nil)
							}
							it.Release()
						deleteEndHistory:
							for it.Next() {
								ts, err := strconv.ParseInt(strings.Split(string(it.Key()), joiner)[1], 10, 64) // get ts of it.Value
								{
									if err != nil {
										return err
									}
									if ts+int64(gdb.hisTimeDuration.Seconds()) <= int64(endTimes[j]) {
										// all history in it.Value needed to delete
										iv := it.Value()
										if len(iv) == 0 {
											continue
										}
										d := &pb.StringHistoricalData{}
										if err := proto.Unmarshal(iv, d); err != nil {
											return err
										}
										row += len(d.TimeStamps)
										batch.Put(it.Key(), nil)
									} else {
										d := &pb.StringHistoricalData{}
										iv := it.Value()
										if len(iv) == 0 {
											continue
										}
										if err := proto.Unmarshal(iv, d); err != nil {
											return err
										}
										t, v := d.GetTimeStamps(), d.GetValues()
										tts := []int32{}
										vs := []string{}
										if k := getEndIndex(endTimes[j], t); k == -1 {
											continue
										} else {
											row += k
											if k == len(endTimes) {
												// all data to be deleted
												batch.Put(it.Key(), nil)
											} else {
												tts = append(tts, t[k:]...)
												vs = append(vs, v[k:]...)
												dn := &pb.StringHistoricalData{TimeStamps: tts, Values: vs}
												if data, err := proto.Marshal(dn); err != nil {
													return err
												} else {
													batch.Put(it.Key(), data)
												}
											}
										}
									}
								}
							}
							it.Release()
						deleteStartHistory:
							for it.Next() {
								ts, err := strconv.ParseInt(strings.Split(string(it.Key()), joiner)[1], 10, 64)
								{
									if err != nil {
										return err
									}
									if ts >= int64(startTimes[j]) {
										iv := it.Value()
										if len(iv) == 0 {
											continue
										}
										d := &pb.StringHistoricalData{}
										if err := proto.Unmarshal(iv, d); err != nil {
											return err
										}
										row += len(d.TimeStamps)
										batch.Put(it.Key(), nil)
									} else {
										d := &pb.StringHistoricalData{}
										iv := it.Value()
										if len(iv) == 0 {
											continue
										}
										if err := proto.Unmarshal(iv, d); err != nil {
											return err
										}
										t, v := d.GetTimeStamps(), d.GetValues()
										tts := []int32{}
										vs := []string{}
										if k := getStartIndex(startTimes[j], t); k == -1 {
											continue
										} else {
											row += len(t) - k
											if k == 0 {
												// all history to be deleted
												batch.Put(it.Key(), nil)
											} else {
												tts = append(tts, t[:k]...)
												vs = append(vs, v[:k]...)
												dn := &pb.StringHistoricalData{
													TimeStamps: tts,
													Values:     vs,
												}
												if data, err := proto.Marshal(dn); err != nil {
													return err
												} else {
													batch.Put(it.Key(), data)
												}
											}
										}
									}
								}
							}
							it.Release()
						deletePartHistory:
							for it.Next() {
								d := &pb.StringHistoricalData{}
								iv := it.Value()
								if len(iv) == 0 {
									continue
								}
								if err := proto.Unmarshal(iv, d); err != nil {
									return err
								}
								t, v := d.GetTimeStamps(), d.GetValues()
								ts := []int32{}
								vs := []string{}
								startIndex, endIndex := getIndex(startTimes[j], endTimes[j], t)
								{
									if startIndex == -1 || endIndex == -1 {
										continue
									}
									row += endIndex - startIndex
									if endIndex-startIndex == len(t) {
										// all history to be deleted
										batch.Put(it.Key(), nil)
									} else {
										ts = append(ts, t[:startIndex]...)
										ts = append(ts, t[endIndex:]...)
										vs = append(vs, v[:startIndex]...)
										vs = append(vs, v[endIndex:]...)
										if len(ts) == 0 {
											batch.Put(it.Key(), nil)
										} else {
											dn := &pb.StringHistoricalData{TimeStamps: ts, Values: vs}
											if data, err := proto.Marshal(dn); err != nil {
												return err
											} else {
												batch.Put(it.Key(), data)
											}
										}
									}
								}
							}
							it.Release()
						}
					}
					innerRows.Store(groupName, row)
					return gdb.hisDb["float32"][groupName].Write(batch, nil)
				}
			})
		}
		if err := g.Wait(); err != nil {
			return err
		}
		for groupName := range namesMap {
			row, _ := innerRows.Load(groupName)
			{
				rows[0] += row.(int)
			}
		}
		return nil
	})
	// delete history in memory
	dg.Go(func() error {
		for _, itemName := range itemNames {
			if md, ok := gdb.stringRmHisDb.Get(itemName); ok {
				// delete history data in memory
				if len(md.TimeStamps) != 0 {
					row := 0
					for j := 0; j < len(startTimes); j++ {
						tts := []int32{}
						vs := []string{}
						if startTimes[j] == -1 && endTimes[j] == -1 {
							goto deleteAllHistory
						} else if startTimes[j] == -1 && endTimes[j] != -1 {
							goto deleteEndHistory
						} else if startTimes[j] != -1 && endTimes[j] == -1 {
							goto deleteStartHistory
						} else {
							goto deletePartHistory
						}
					deleteAllHistory:
						md.TimeStamps = []int32{}
						md.Values = []string{}
						row += len(md.TimeStamps)
					deleteEndHistory:
						if k := getEndIndex(endTimes[j], md.TimeStamps); k == -1 {
							continue
						} else {
							md.TimeStamps = md.TimeStamps[k:]
							md.Values = md.Values[k:]
							row += k
						}
					deleteStartHistory:
						if k := getStartIndex(startTimes[j], md.TimeStamps); k == -1 {
							continue
						} else {
							md.TimeStamps = md.TimeStamps[:k]
							md.Values = md.Values[:k]
							row += len(md.TimeStamps) - k
						}
					deletePartHistory:
						startIndex, endIndex := getIndex(startTimes[j], endTimes[j], md.TimeStamps)
						{
							if startIndex == -1 || endIndex == -1 {
								continue
							}
							row += endIndex - startIndex
							tts = append(tts, md.TimeStamps[:startIndex]...)
							tts = append(tts, md.TimeStamps[endIndex:]...)
							vs = append(vs, md.Values[:startIndex]...)
							vs = append(vs, md.Values[endIndex:]...)
							md.TimeStamps, md.Values = tts, vs
						}
					}
					rows[1] = row
				}
				gdb.stringRmHisDb.Set(itemName, md)
			}
		}
		return nil
	})
	if err := dg.Wait(); err != nil {
		return -1, err
	} else {
		return rows[0] + rows[1] + rows[2], nil
	}
bool32:
	// delete history in leveldb
	dg.Go(func() error {
		g := errgroup.Group{}
		innerRows := sync.Map{}
		for groupName, groupItemNames := range namesMap {
			groupName := groupName
			groupItemNames := groupItemNames
			g.Go(func() error {
				if sn, err := gdb.hisDb["bool"][groupName].GetSnapshot(); err != nil {
					return err
				} else {
					defer sn.Release()
					batch := &leveldb.Batch{}
					row := 0
					for index := range groupItemNames {
						index := index
						itemName := groupItemNames[index]
						for j := 0; j < len(startTimes); j++ {
							if startTimes[j] > endTimes[j] && startTimes[j] != -1 && endTimes[j] != -1 {
								continue
							}
							var s, e string
							var it iterator.Iterator
							if startTimes[j] == -1 && endTimes[j] == -1 {
								// delete all history data
								it = sn.NewIterator(util.BytesPrefix(convertStringToByte(itemName)), nil)
								goto deleteAllHistory
							} else if startTimes[j] == -1 && endTimes[j] != -1 {
								// delete history data ==> (,et)
								s = itemName
								e = itemName + joiner + strconv.Itoa(int(endTimes[j])+int(gdb.hisTimeDuration.Seconds()))
								it = sn.NewIterator(&util.Range{Start: convertStringToByte(s), Limit: convertStringToByte(e)}, nil)
								goto deleteEndHistory
							} else if startTimes[j] != -1 && endTimes[j] == -1 {
								// delete history data ==> [st,)
								s = itemName + joiner + strconv.Itoa(int(startTimes[j])-int(gdb.hisTimeDuration.Seconds()))
								e = itemName + joiner + strconv.Itoa(int(time.Now().Unix()+8*3600)+int(gdb.hisTimeDuration.Seconds()))
								it = sn.NewIterator(&util.Range{Start: convertStringToByte(s), Limit: convertStringToByte(e)}, nil)
								goto deleteStartHistory
							} else {
								// delete history data ==> (st,et]
								s := itemName + joiner + strconv.Itoa(int(startTimes[j])-int(gdb.hisTimeDuration.Seconds()))
								e := itemName + joiner + strconv.Itoa(int(endTimes[j])+int(gdb.hisTimeDuration.Seconds()))
								it = sn.NewIterator(&util.Range{Start: convertStringToByte(s), Limit: convertStringToByte(e)}, nil)
								goto deletePartHistory
							}
						deleteAllHistory:
							for it.Next() {
								batch.Put(it.Key(), nil)
							}
							it.Release()
						deleteEndHistory:
							for it.Next() {
								ts, err := strconv.ParseInt(strings.Split(string(it.Key()), joiner)[1], 10, 64) // get ts of it.Value
								{
									if err != nil {
										return err
									}
									if ts+int64(gdb.hisTimeDuration.Seconds()) <= int64(endTimes[j]) {
										// all history in it.Value needed to delete
										iv := it.Value()
										if len(iv) == 0 {
											continue
										}
										d := &pb.BoolHistoricalData{}
										if err := proto.Unmarshal(iv, d); err != nil {
											return err
										}
										row += len(d.TimeStamps)
										batch.Put(it.Key(), nil)
									} else {
										d := &pb.BoolHistoricalData{}
										iv := it.Value()
										if len(iv) == 0 {
											continue
										}
										if err := proto.Unmarshal(iv, d); err != nil {
											return err
										}
										t, v := d.GetTimeStamps(), d.GetValues()
										tts := []int32{}
										vs := []bool{}
										if k := getEndIndex(endTimes[j], t); k == -1 {
											continue
										} else {
											row += k
											if k == len(endTimes) {
												// all data to be deleted
												batch.Put(it.Key(), nil)
											} else {
												tts = append(tts, t[k:]...)
												vs = append(vs, v[k:]...)
												dn := &pb.BoolHistoricalData{TimeStamps: tts, Values: vs}
												if data, err := proto.Marshal(dn); err != nil {
													return err
												} else {
													batch.Put(it.Key(), data)
												}
											}
										}
									}
								}
							}
							it.Release()
						deleteStartHistory:
							for it.Next() {
								ts, err := strconv.ParseInt(strings.Split(string(it.Key()), joiner)[1], 10, 64)
								{
									if err != nil {
										return err
									}
									if ts >= int64(startTimes[j]) {
										iv := it.Value()
										if len(iv) == 0 {
											continue
										}
										d := &pb.BoolHistoricalData{}
										if err := proto.Unmarshal(iv, d); err != nil {
											return err
										}
										row += len(d.TimeStamps)
										batch.Put(it.Key(), nil)
									} else {
										d := &pb.BoolHistoricalData{}
										iv := it.Value()
										if len(iv) == 0 {
											continue
										}
										if err := proto.Unmarshal(iv, d); err != nil {
											return err
										}
										t, v := d.GetTimeStamps(), d.GetValues()
										tts := []int32{}
										vs := []bool{}
										if k := getStartIndex(startTimes[j], t); k == -1 {
											continue
										} else {
											row += len(t) - k
											if k == 0 {
												// all history to be deleted
												batch.Put(it.Key(), nil)
											} else {
												tts = append(tts, t[:k]...)
												vs = append(vs, v[:k]...)
												dn := &pb.BoolHistoricalData{
													TimeStamps: tts,
													Values:     vs,
												}
												if data, err := proto.Marshal(dn); err != nil {
													return err
												} else {
													batch.Put(it.Key(), data)
												}
											}
										}
									}
								}
							}
							it.Release()
						deletePartHistory:
							for it.Next() {
								d := &pb.BoolHistoricalData{}
								iv := it.Value()
								if len(iv) == 0 {
									continue
								}
								if err := proto.Unmarshal(iv, d); err != nil {
									return err
								}
								t, v := d.GetTimeStamps(), d.GetValues()
								ts := []int32{}
								vs := []bool{}
								startIndex, endIndex := getIndex(startTimes[j], endTimes[j], t)
								{
									if startIndex == -1 || endIndex == -1 {
										continue
									}
									row += endIndex - startIndex
									if endIndex-startIndex == len(t) {
										// all history to be deleted
										batch.Put(it.Key(), nil)
									} else {
										ts = append(ts, t[:startIndex]...)
										ts = append(ts, t[endIndex:]...)
										vs = append(vs, v[:startIndex]...)
										vs = append(vs, v[endIndex:]...)
										if len(ts) == 0 {
											batch.Put(it.Key(), nil)
										} else {
											dn := &pb.BoolHistoricalData{TimeStamps: ts, Values: vs}
											if data, err := proto.Marshal(dn); err != nil {
												return err
											} else {
												batch.Put(it.Key(), data)
											}
										}
									}
								}
							}
							it.Release()
						}
					}
					innerRows.Store(groupName, row)
					return gdb.hisDb["bool"][groupName].Write(batch, nil)
				}
			})
		}
		if err := g.Wait(); err != nil {
			return err
		}
		for groupName := range namesMap {
			row, _ := innerRows.Load(groupName)
			{
				rows[0] += row.(int)
			}
		}
		return nil
	})
	// delete history in memory
	dg.Go(func() error {
		for _, itemName := range itemNames {
			if md, ok := gdb.boolRmHisDb.Get(itemName); ok {
				// delete history data in memory
				if len(md.TimeStamps) != 0 {
					row := 0
					for j := 0; j < len(startTimes); j++ {
						tts := []int32{}
						vs := []bool{}
						if startTimes[j] == -1 && endTimes[j] == -1 {
							goto deleteAllHistory
						} else if startTimes[j] == -1 && endTimes[j] != -1 {
							goto deleteEndHistory
						} else if startTimes[j] != -1 && endTimes[j] == -1 {
							goto deleteStartHistory
						} else {
							goto deletePartHistory
						}
					deleteAllHistory:
						md.TimeStamps = []int32{}
						md.Values = []bool{}
						row += len(md.TimeStamps)
					deleteEndHistory:
						if k := getEndIndex(endTimes[j], md.TimeStamps); k == -1 {
							continue
						} else {
							md.TimeStamps = md.TimeStamps[k:]
							md.Values = md.Values[k:]
							row += k
						}
					deleteStartHistory:
						if k := getStartIndex(startTimes[j], md.TimeStamps); k == -1 {
							continue
						} else {
							md.TimeStamps = md.TimeStamps[:k]
							md.Values = md.Values[:k]
							row += len(md.TimeStamps) - k
						}
					deletePartHistory:
						startIndex, endIndex := getIndex(startTimes[j], endTimes[j], md.TimeStamps)
						{
							if startIndex == -1 || endIndex == -1 {
								continue
							}
							row += endIndex - startIndex
							tts = append(tts, md.TimeStamps[:startIndex]...)
							tts = append(tts, md.TimeStamps[endIndex:]...)
							vs = append(vs, md.Values[:startIndex]...)
							vs = append(vs, md.Values[endIndex:]...)
							md.TimeStamps, md.Values = tts, vs
						}
					}
					rows[1] = row
				}
				gdb.boolRmHisDb.Set(itemName, md)
			}
		}
		return nil
	})
	if err := dg.Wait(); err != nil {
		return -1, err
	} else {
		return rows[0] + rows[1] + rows[2], nil
	}
}

func (gdb *Gdb) getNowTime() string {
	return time.Now().Format(timeFormatString)
}

func (gdb *Gdb) console(v interface{}) {
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
	if v, err := gdb.GetRealTimeData(itemNames, groupNames); err != nil {
		return "", err
	} else {
		if r, err := json.Marshal(v); err != nil {
			return "", err
		} else {
			return string(r), err
		}
	}
}

func (gdb *Gdb) getFloatHData(groupNames, itemNames []string, startTimes, endTimes, intervals []int32) (string, error) {
	if result, err := gdb.GetFloatHistoricalData(groupNames, itemNames, startTimes, endTimes, intervals); err != nil {
		return "", err
	} else {
		if r, err := json.Marshal(result); err != nil {
			return "", err
		} else {
			return string(r), nil
		}
	}
}

func (gdb *Gdb) getIntHData(groupNames, itemNames []string, startTimes, endTimes, intervals []int32) (string, error) {
	if result, err := gdb.GetIntHistoricalData(groupNames, itemNames, startTimes, endTimes, intervals); err != nil {
		return "", err
	} else {
		if r, err := json.Marshal(result); err != nil {
			return "", err
		} else {
			return string(r), nil
		}
	}
}

func (gdb *Gdb) getStringHData(groupNames, itemNames []string, startTimes, endTimes, intervals []int32) (string, error) {
	if result, err := gdb.GetStringHistoricalData(groupNames, itemNames, startTimes, endTimes, intervals); err != nil {
		return "", err
	} else {
		if r, err := json.Marshal(result); err != nil {
			return "", err
		} else {
			return string(r), nil
		}
	}
}

func (gdb *Gdb) getBoolHData(groupNames, itemNames []string, startTimes, endTimes, intervals []int32) (string, error) {
	if result, err := gdb.GetFloatHistoricalData(groupNames, itemNames, startTimes, endTimes, intervals); err != nil {
		return "", err
	} else {
		if r, err := json.Marshal(result); err != nil {
			return "", err
		} else {
			return string(r), nil
		}
	}
}

func (gdb *Gdb) getFloatHDataWithTs(groupNames, itemNames []string, timeStamps ...[]int32) (string, error) {
	if result, err := gdb.GetFloatHistoricalDataWithStamp(groupNames, itemNames, timeStamps); err != nil {
		return "", err
	} else {
		if r, err := json.Marshal(result); err != nil {
			return "", err
		} else {
			return string(r), nil
		}
	}
}

func (gdb *Gdb) getIntHDataWithTs(groupNames, itemNames []string, timeStamps ...[]int32) (string, error) {
	if result, err := gdb.GetIntHistoricalDataWithStamp(groupNames, itemNames, timeStamps); err != nil {
		return "", err
	} else {
		if r, err := json.Marshal(result); err != nil {
			return "", err
		} else {
			return string(r), nil
		}
	}
}

func (gdb *Gdb) getStringHDataWithTs(groupNames, itemNames []string, timeStamps ...[]int32) (string, error) {
	if result, err := gdb.GetStringHistoricalDataWithStamp(groupNames, itemNames, timeStamps); err != nil {
		return "", err
	} else {
		if r, err := json.Marshal(result); err != nil {
			return "", err
		} else {
			return string(r), nil
		}
	}
}

func (gdb *Gdb) getBoolHDataWithTs(groupNames, itemNames []string, timeStamps ...[]int32) (string, error) {
	if result, err := gdb.GetBoolHistoricalDataWithStamp(groupNames, itemNames, timeStamps); err != nil {
		return "", err
	} else {
		if r, err := json.Marshal(result); err != nil {
			return "", err
		} else {
			return string(r), nil
		}
	}
}

// returned itemName = itemName + joiner + groupName
func (gdb *Gdb) checkItemDataType(groupNames []string, itemNames [][]string, dataType string) ([][]string, error) {
	names := [][]string{}
	for i := 0; i < len(groupNames); i++ {
		name := []string{}
		for j := 0; j < len(itemNames[i]); j++ {
			itemName := itemNames[i][j] + joiner + groupNames[i]
			if dt, ok := gdb.rtDbFilter.Get(itemName); ok {
				if dt.(string) != dataType {
					return nil, fmt.Errorf("inconsistent dataType")
				}
				name = append(name, itemName)
			} else {
				return nil, fmt.Errorf("item " + itemNames[i][j] + "not exist")
			}
		}
		names = append(names, name)
	}
	return names, nil
}

// returned itemName = itemName + joiner + groupName
func (gdb *Gdb) checkSingleItemDataTypeWithMap(groupNames []string, itemNames []string, dataType string) (map[string][]string, map[string][]string, error) {
	namesMap := map[string][]string{}
	rawNamesMap := map[string][]string{}
	for i := 0; i < len(groupNames); i++ {
		itemName := itemNames[i] + joiner + groupNames[i]
		if dt, ok := gdb.rtDbFilter.Get(itemName); ok {
			if dt.(string) != dataType {
				return nil, nil, fmt.Errorf("inconsistent dataType")
			} else {
				if _, ok := namesMap[groupNames[i]]; ok {
					namesMap[groupNames[i]] = append(namesMap[groupNames[i]], itemName)
					rawNamesMap[groupNames[i]] = append(rawNamesMap[groupNames[i]], itemNames[i])
				} else {
					namesMap[groupNames[i]] = []string{itemName}
					rawNamesMap[groupNames[i]] = []string{itemNames[i]}
				}
			}
		} else {
			return nil, nil, fmt.Errorf("item " + itemNames[i] + "not exist")
		}
	}
	return rawNamesMap, namesMap, nil
}

// returned itemName = itemName + joiner + groupName
func (gdb *Gdb) checkItemDataTypeWithMap(groupNames []string, itemNames [][]string, dataType string) (map[string][]string, error) {
	namesMap := map[string][]string{}
	for i := 0; i < len(groupNames); i++ {
		names := []string{}
		for j := 0; j < len(itemNames[i]); j++ {
			itemName := itemNames[i][j] + joiner + groupNames[i]
			names = append(names, itemName)
			if dt, ok := gdb.rtDbFilter.Get(itemName); ok {
				if dt.(string) != dataType {
					return nil, fmt.Errorf("inconsistent dataType")
				}
			} else {
				return nil, fmt.Errorf("item " + itemNames[i][j] + "not exist")
			}
		}
		namesMap[groupNames[i]] = names
	}
	return namesMap, nil
}

// ensure that historical data in the database will not be repeated
func (gdb *Gdb) checkTimeStampsInDb(itemName string, startTime, endTime int, sn *leveldb.Snapshot) error {
	s := itemName + joiner + strconv.Itoa(startTime)
	e := itemName + joiner + strconv.Itoa(endTime)
	it := sn.NewIterator(&util.Range{Start: convertStringToByte(s), Limit: convertStringToByte(e)}, nil)
	defer it.Release()
	if !it.First() {
		return nil
	}
	if it.Value() != nil {
		st, _ := strconv.ParseInt(strings.Replace(string(it.Key()), itemName+joiner, "", -1), 10, 64)
		it.Last()
		et, _ := strconv.ParseInt(strings.Replace(string(it.Key()), itemName+joiner, "", -1), 10, 64)
		return fmt.Errorf("time from "+time.Unix(st, 0).Format(timeFormatString), " to "+time.Unix(et, 0).Format(timeFormatString)+" has history data")
	}
	for it.Next() {
		if len(it.Value()) == 0 {
			continue
		} else {
			st, _ := strconv.ParseInt(strings.Replace(string(it.Key()), itemName+joiner, "", -1), 10, 64)
			it.Last()
			et, _ := strconv.ParseInt(strings.Replace(string(it.Key()), itemName+joiner, "", -1), 10, 64)
			return fmt.Errorf("time from "+time.Unix(st, 0).Format(timeFormatString), " to "+time.Unix(et, 0).Format(timeFormatString)+" has history data")
		}
	}
	return nil
}

func (gdb *Gdb) writeFloatRtData(infos []map[string]interface{}) (TimeRows, error) {
	var groupNames []string
	var itemNames [][]string
	var itemValues [][]float32
	for _, info := range infos {
		groupNames = append(groupNames, info["groupName"].(string))
		itemNames = append(itemNames, info["itemName"].([]string))
		itemValues = append(itemValues, info["itemValue"].([]float32))
	}
	return gdb.BatchWriteFloatData(groupNames, itemNames, itemValues)
}

func (gdb *Gdb) writeIntRtData(infos []map[string]interface{}) (TimeRows, error) {
	var groupNames []string
	var itemNames [][]string
	var itemValues [][]int32
	for _, info := range infos {
		groupNames = append(groupNames, info["groupName"].(string))
		itemNames = append(itemNames, info["itemName"].([]string))
		itemValues = append(itemValues, info["itemValue"].([]int32))
	}
	return gdb.BatchWriteIntData(groupNames, itemNames, itemValues)
}

func (gdb *Gdb) writeStringRtData(infos []map[string]interface{}) (TimeRows, error) {
	var groupNames []string
	var itemNames [][]string
	var itemValues [][]string
	for _, info := range infos {
		groupNames = append(groupNames, info["groupName"].(string))
		itemNames = append(itemNames, info["itemName"].([]string))
		itemValues = append(itemValues, info["itemValue"].([]string))
	}
	return gdb.BatchWriteStringData(groupNames, itemNames, itemValues)
}

func (gdb *Gdb) writeBoolRtData(infos []map[string]interface{}) (TimeRows, error) {
	var groupNames []string
	var itemNames [][]string
	var itemValues [][]bool
	for _, info := range infos {
		groupNames = append(groupNames, info["groupName"].(string))
		itemNames = append(itemNames, info["itemName"].([]string))
		itemValues = append(itemValues, info["itemValue"].([]bool))
	}
	return gdb.BatchWriteBoolData(groupNames, itemNames, itemValues)
}

func (gdb *Gdb) getDbInfo() (cmap.ConcurrentMap, error) {
	infoNames := []string{"systemInfo", "floatInfo", "intInfo", "stringInfo", "boolInfo"}
	m := cmap.New()
	g := errgroup.Group{}
	for _, name := range infoNames {
		infoName := name
		g.Go(func() error {
			var sn *leveldb.Snapshot
			infos := map[string]string{}
			var err error
			switch infoName {
			case "systemInfo":
				sn, err = gdb.systemInfoDb.GetSnapshot()
				if err != nil {
					return err
				}
				defer sn.Release()
				goto system
			case "floatInfo":
				sn, err = gdb.floatInfoDb.GetSnapshot()
				if err != nil {
					return err
				}
				defer sn.Release()
				goto other
			case "intInfo":
				sn, err = gdb.intInfoDb.GetSnapshot()
				if err != nil {
					return err
				}
				defer sn.Release()
				goto other
			case "stringInfo":
				sn, err = gdb.stringInfoDb.GetSnapshot()
				if err != nil {
					return err
				}
				defer sn.Release()
				goto other
			case "boolInfo":
				sn, err = gdb.boolInfoDb.GetSnapshot()
				if err != nil {
					return err
				}
				defer sn.Release()
				goto other
			}
		system:
			if r, err := sn.Get(convertStringToByte(ram), nil); err != nil {
				infos[ram] = ""
			} else {
				infos[ram] = string(r) // ram
			}
			if c, err := sn.Get(convertStringToByte(cpu), nil); err != nil {
				infos[cpu] = ""
			} else {
				infos[cpu] = string(c) // cpu
			}
			if fs, err := sn.Get(convertStringToByte(fileSize), nil); err != nil {
				infos[fileSize] = ""
			} else {
				infos[fileSize] = string(fs)
			}
			if st, err := sn.Get(convertStringToByte(timeKey), nil); err != nil {
				infos[timeKey] = ""
			} else {
				infos[timeKey] = string(st)
			}
			m.Set(infoName, infos)
			return nil
		other:
			if w, err := sn.Get(convertStringToByte(syncBytes), nil); err != nil {
				infos[syncBytes] = ""
			} else {
				infos[syncBytes] = string(w)
			}
			if s, err := sn.Get(convertStringToByte(speed), nil); err != nil {
				infos[speed] = ""
			} else {
				infos[speed] = string(s)
			}
			if st, err := sn.Get(convertStringToByte(syncTime), nil); err != nil {
				infos[syncTime] = ""
			} else {
				infos[syncTime] = string(st)
			}
			if stc, err := sn.Get(convertStringToByte(syncConsumeTime), nil); err != nil {
				infos[syncConsumeTime] = ""
			} else {
				infos[syncConsumeTime] = string(stc)
			}
			m.Set(infoName, infos)
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return nil, err
	}
	return m, nil
}

func (gdb *Gdb) getDbInfoHistory(infoType, itemName string, startTimeStamps, endTimeStamps, intervals []int32) (cmap.ConcurrentMap, error) {
	if len(startTimeStamps) == len(endTimeStamps) && len(startTimeStamps) == len(intervals) {
		rawData := cmap.New()
		var sn *leveldb.Snapshot
		var err error
		switch infoType {
		case "systemInfo":
			sn, err = gdb.systemInfoDb.GetSnapshot()
			if err != nil {
				return nil, err
			}
			break
		case "floatInfo":
			sn, err = gdb.floatInfoDb.GetSnapshot()
			if err != nil {
				return nil, err
			}
			break
		case "intInfo":
			sn, err = gdb.intInfoDb.GetSnapshot()
			if err != nil {
				return nil, err
			}
			break
		case "stringInfo":
			sn, err = gdb.stringInfoDb.GetSnapshot()
			if err != nil {
				return nil, err
			}
			break
		case "boolInfo":
			sn, err = gdb.boolInfoDb.GetSnapshot()
			if err != nil {
				return nil, err
			}
			break
		default:
			return nil, fmt.Errorf("incorrect infoType:" + infoType)
		}
		defer sn.Release()
		for i := 0; i < len(startTimeStamps); i++ {
			s := startTimeStamps[i] // startTime
			e := endTimeStamps[i]   // endTime
			interval := intervals[i]
			if s >= e {
				// startTime > endTime, continue
				continue
			}
			startKey := strings.Builder{}
			startKey.Write(convertStringToByte(itemName))
			startKey.Write(convertStringToByte(fmt.Sprintf("%d", s)))
			endKey := strings.Builder{}
			endKey.Write(convertStringToByte(itemName))
			endKey.Write(convertStringToByte(fmt.Sprintf("%d", e)))
			it := sn.NewIterator(&util.Range{Start: convertStringToByte(startKey.String()), Limit: convertStringToByte(endKey.String())}, nil)
			var values []string
			var count int32
			var timeStamps []string
			for it.Next() {
				if count%interval == 0 {
					values = append(values, fmt.Sprintf("%s", it.Value()))
					timeStamps = append(timeStamps, strings.Replace(fmt.Sprintf("%s", it.Key()), itemName, "", -1))
				}
				count++
			}
			it.Release()
			rawData.Set(itemName, [][]string{timeStamps, values})
		}
		return rawData, nil
	} else {
		return nil, fmt.Errorf("inconsistent length of starTimes, endTimes and intervals")
	}
}

func (gdb *Gdb) getDbSize() ([]map[string]interface{}, error) {
	r := []map[string]interface{}{}
	for _, groupName := range gdb.groupNames {
		t := map[string]interface{}{"groupName": groupName}
		var total float64
		for _, dataType := range dataTypes {
			if size, err := dirSize(gdb.dbPath + gdb.delimiter + "historicalData" + gdb.delimiter + dataType + gdb.delimiter + groupName); err != nil {
				return nil, err
			} else {
				t[dataType] = size
				total += size
			}
		}
		tt, _ := decimal.NewFromFloat(total).Round(2).Float64()
		t["total"] = tt
		r = append(r, t)
	}
	return r, nil
}

func (gdb *Gdb) writeFloatMemDataToLevelDb(key string, value memap.RmHisDbFloatItems, batch *leveldb.Batch) (int, error) {
	count := 0
	if len(value.TimeStamps) != 0 {
		k := key + joiner + strconv.Itoa(int(value.TimeStamps[0]))
		m := &pb.FloatHistoricalData{TimeStamps: value.TimeStamps, Values: value.Values}
		if data, err := proto.Marshal(m); err != nil {
			return -1, err
		} else {
			batch.Put(convertStringToByte(k), data)
			count += len(data)
		}
	} else {
		return -1, nil
	}
	return count, nil
}

func (gdb *Gdb) writeIntMemDataToLevelDb(key string, value memap.RmHisDbIntItems, batch *leveldb.Batch) (int, error) {
	count := 0
	if len(value.TimeStamps) != 0 {
		k := key + joiner + strconv.Itoa(int(value.TimeStamps[0])) + joiner
		m := &pb.IntHistoricalData{TimeStamps: value.TimeStamps, Values: value.Values}
		if data, err := proto.Marshal(m); err != nil {
			return -1, err
		} else {
			batch.Put(convertStringToByte(k), data)
			count += len(data)
		}
	} else {
		return -1, nil
	}
	return count, nil
}

func (gdb *Gdb) writeStringMemDataToLevelDb(key string, value memap.RmHisDbStringItems, batch *leveldb.Batch) (int, error) {
	count := 0
	if len(value.TimeStamps) != 0 {
		k := key + joiner + strconv.Itoa(int(value.TimeStamps[0]))
		m := &pb.StringHistoricalData{TimeStamps: value.TimeStamps, Values: value.Values}
		if data, err := proto.Marshal(m); err != nil {
			return -1, err
		} else {
			batch.Put(convertStringToByte(k), data)
			count += len(data)
		}
	} else {
		return -1, nil
	}
	return count, nil
}

func (gdb *Gdb) writeBoolMemDataToLevelDb(key string, value memap.RmHisDbBoolItems, batch *leveldb.Batch) (int, error) {
	count := 0
	if len(value.TimeStamps) != 0 {
		k := key + joiner + strconv.Itoa(int(value.TimeStamps[0]))
		m := &pb.BoolHistoricalData{TimeStamps: value.TimeStamps, Values: value.Values}
		if data, err := proto.Marshal(m); err != nil {
			return -1, err
		} else {
			batch.Put(convertStringToByte(k), data)
			count += len(data)
		}
	} else {
		return -1, nil
	}
	return count, nil
}

// syncRtData sync realTime data in memory to other database or file system
// you should call in your own program in another goroutine
func (gdb *Gdb) syncRtData() error {
	t := time.NewTicker(gdb.rtTimeDuration)
	for {
		select {
		case <-t.C:
			if err := gdb.rtDb.Sync(); err != nil {
				fmt.Println("fatal error occurred while synchronizing realTime data:" + err.Error())
				time.Sleep(time.Minute)
				os.Exit(-1)
				return err
			}
		}
	}
}

// syncHisData sync history data in memory to other database or file system
// you should call in your own program in another goroutine
func (gdb *Gdb) syncHisData() error {
	t := time.NewTicker(gdb.hisTimeDuration)
	for {
		select {
		case st := <-t.C:
			if !gdb.isReloading {
				if err := gdb.innerSync(st); err != nil {
					fmt.Println("fatal error occurred while synchronizing history data:" + err.Error())
					time.Sleep(time.Minute)
					os.Exit(-1)
					return err
				}
			}
		}
	}
}

func (gdb *Gdb) innerSync(st time.Time) error {
	gdb.mu.RLock()
	gdb.syncStatus = true
	g := errgroup.Group{}
	g.Go(func() error {
		// float
		batchMap := map[string]*leveldb.Batch{}
		for i := 0; i < len(gdb.groupNames); i++ {
			batchMap[gdb.groupNames[i]] = &leveldb.Batch{}
		}
		count := 0
		for item := range gdb.floatRmHisDb.IterBuffered() {
			// key = itemName + joiner + groupName
			if c, err := gdb.writeFloatMemDataToLevelDb(item.Key, item.Val, batchMap[item.Val.GroupName]); err != nil {
				return err
			} else {
				count += c
				gdb.floatRmHisDb.Remove(item.Key)
			}
		}
		g1 := errgroup.Group{}
		for _, name := range gdb.groupNames {
			groupName := name
			g1.Go(func() error {
				if batchMap[groupName].Len() == 0 {
					return nil
				}
				return gdb.hisDb["float32"][groupName].Write(batchMap[groupName], nil)
			})
		}
		if err := g1.Wait(); err != nil {
			return err
		}
		d := time.Since(st).Milliseconds()
		sd := fmt.Sprintf("%dms/%d", d, count)
		infoBatch := &leveldb.Batch{}
		infoBatch.Put(convertStringToByte(syncTime), convertStringToByte(st.Format(timeFormatString)))                   // syncTime
		infoBatch.Put(convertStringToByte(syncBytes), convertStringToByte(strconv.Itoa(count)))                          // syncBytes
		infoBatch.Put(convertStringToByte(syncConsumeTime), convertStringToByte(strconv.Itoa(int(d))))                   // consume time
		infoBatch.Put(convertStringToByte(syncConsumeTime+strconv.Itoa(int(st.Unix()+8*3600))), convertStringToByte(sd)) // history data of consuming time
		return gdb.floatInfoDb.Write(infoBatch, nil)
	})
	g.Go(func() error {
		// int
		batchMap := map[string]*leveldb.Batch{}
		for i := 0; i < len(gdb.groupNames); i++ {
			batchMap[gdb.groupNames[i]] = &leveldb.Batch{}
		}
		count := 0
		for item := range gdb.intRmHisDb.IterBuffered() {
			if c, err := gdb.writeIntMemDataToLevelDb(item.Key, item.Val, batchMap[item.Val.GroupName]); err != nil {
				return err
			} else {
				count += c
				gdb.intRmHisDb.Remove(item.Key)
			}
		}
		g1 := errgroup.Group{}
		for _, name := range gdb.groupNames {
			groupName := name
			g1.Go(func() error {
				if batchMap[groupName].Len() == 0 {
					return nil
				}
				return gdb.hisDb["int32"][groupName].Write(batchMap[groupName], nil)
			})
		}
		if err := g1.Wait(); err != nil {
			return err
		}
		d := time.Since(st).Milliseconds()
		sd := fmt.Sprintf("%dms/%d", d, count)
		infoBatch := &leveldb.Batch{}
		infoBatch.Put(convertStringToByte(syncTime), convertStringToByte(st.Format(timeFormatString)))                   // syncTime
		infoBatch.Put(convertStringToByte(syncBytes), convertStringToByte(strconv.Itoa(count)))                          // syncBytes
		infoBatch.Put(convertStringToByte(syncConsumeTime), convertStringToByte(strconv.Itoa(int(d))))                   // consume time
		infoBatch.Put(convertStringToByte(syncConsumeTime+strconv.Itoa(int(st.Unix()+8*3600))), convertStringToByte(sd)) // history data of consuming time
		return gdb.intInfoDb.Write(infoBatch, nil)
	})
	g.Go(func() error {
		// string
		batchMap := map[string]*leveldb.Batch{}
		for i := 0; i < len(gdb.groupNames); i++ {
			batchMap[gdb.groupNames[i]] = &leveldb.Batch{}
		}
		count := 0
		for item := range gdb.stringRmHisDb.IterBuffered() {
			if c, err := gdb.writeStringMemDataToLevelDb(item.Key, item.Val, batchMap[item.Val.GroupName]); err != nil {
				return err
			} else {
				count += c
				gdb.stringRmHisDb.Remove(item.Key)
			}
		}
		g1 := errgroup.Group{}
		for _, name := range gdb.groupNames {
			groupName := name
			g1.Go(func() error {
				if batchMap[groupName].Len() == 0 {
					return nil
				}
				return gdb.hisDb["string"][groupName].Write(batchMap[groupName], nil)
			})
		}
		if err := g1.Wait(); err != nil {
			return err
		}
		d := time.Since(st).Milliseconds()
		sd := fmt.Sprintf("%dms/%d", d, count)
		infoBatch := &leveldb.Batch{}
		infoBatch.Put(convertStringToByte(syncTime), convertStringToByte(st.Format(timeFormatString)))                   // syncTime
		infoBatch.Put(convertStringToByte(syncBytes), convertStringToByte(strconv.Itoa(count)))                          // syncBytes
		infoBatch.Put(convertStringToByte(syncConsumeTime), convertStringToByte(strconv.Itoa(int(d))))                   // consume time
		infoBatch.Put(convertStringToByte(syncConsumeTime+strconv.Itoa(int(st.Unix()+8*3600))), convertStringToByte(sd)) // history data of consuming time
		return gdb.stringInfoDb.Write(infoBatch, nil)
	})
	g.Go(func() error {
		// bool
		batchMap := map[string]*leveldb.Batch{}
		for i := 0; i < len(gdb.groupNames); i++ {
			batchMap[gdb.groupNames[i]] = &leveldb.Batch{}
		}
		count := 0
		for item := range gdb.boolRmHisDb.IterBuffered() {
			if c, err := gdb.writeBoolMemDataToLevelDb(item.Key, item.Val, batchMap[item.Val.GroupName]); err != nil {
				return err
			} else {
				count += c
				gdb.boolRmHisDb.Remove(item.Key)
			}
		}
		g1 := errgroup.Group{}
		for _, name := range gdb.groupNames {
			groupName := name
			g1.Go(func() error {
				if batchMap[groupName].Len() == 0 {
					return nil
				}
				return gdb.hisDb["bool"][groupName].Write(batchMap[groupName], nil)
			})
		}
		if err := g1.Wait(); err != nil {
			return err
		}
		d := time.Since(st).Milliseconds()
		sd := fmt.Sprintf("%dms/%d", d, count)
		infoBatch := &leveldb.Batch{}
		infoBatch.Put(convertStringToByte(syncTime), convertStringToByte(st.Format(timeFormatString)))                   // syncTime
		infoBatch.Put(convertStringToByte(syncBytes), convertStringToByte(strconv.Itoa(count)))                          // syncBytes
		infoBatch.Put(convertStringToByte(syncConsumeTime), convertStringToByte(strconv.Itoa(int(d))))                   // consume time
		infoBatch.Put(convertStringToByte(syncConsumeTime+strconv.Itoa(int(st.Unix()+8*3600))), convertStringToByte(sd)) // history data of consuming time
		return gdb.boolInfoDb.Write(infoBatch, nil)
	})
	if err := g.Wait(); err != nil {
		gdb.syncStatus = false
		gdb.mu.RUnlock()
		fmt.Println(st.Format("2006-01-02 15:04:05"), ": fail in syncing history data: "+err.Error())
		return err
	}
	gdb.syncStatus = false
	gdb.mu.RUnlock()
	return nil
}
