// +build gdbClient

/*
creatTime: 2021/2/8
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.16
*/

package db

// gRPC, for details of gRPC, you can see: https://github.com/grpc/grpc-go

import (
	"context"
	"encoding/base64"
	"fmt"
	pb "github.com/JustKeepSilence/gdb/model"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"io/ioutil"
	"reflect"
	"strings"
	"time"
)

type server struct {
	pb.UnimplementedGroupServer
	pb.UnimplementedItemServer
	pb.UnimplementedDataServer
	pb.UnimplementedPageServer
	pb.UnimplementedCalcServer
	gdb     *Gdb
	configs Config
}

// group handler

func (s *server) AddGroups(_ context.Context, r *pb.AddedGroupInfos) (*pb.TimeRows, error) {
	st := time.Now()
	infos := []AddedGroupInfo{}
	for _, groupInfo := range r.GetGroupInfos() {
		infos = append(infos, AddedGroupInfo{GroupName: groupInfo.GroupName, ColumnNames: groupInfo.ColumnNames})
	}
	if result, err := s.gdb.AddGroups(infos...); err != nil {
		return &pb.TimeRows{}, err
	} else {
		return &pb.TimeRows{EffectedRows: int32(result.EffectedRows), Times: time.Since(st).Milliseconds()}, nil
	}
}

func (s *server) DeleteGroups(_ context.Context, r *pb.GroupNamesInfo) (*pb.TimeRows, error) {
	st := time.Now()
	info := GroupNamesInfo{GroupNames: r.GroupNames}
	if result, err := s.gdb.DeleteGroups(info); err != nil {
		return &pb.TimeRows{}, err
	} else {
		return &pb.TimeRows{EffectedRows: int32(result.EffectedRows), Times: time.Since(st).Milliseconds()}, nil
	}
}

func (s *server) GetGroups(_ context.Context, _ *emptypb.Empty) (*pb.GroupNamesInfo, error) {
	if result, err := s.gdb.GetGroups(); err != nil {
		return &pb.GroupNamesInfo{}, err
	} else {
		return &pb.GroupNamesInfo{GroupNames: result.GroupNames}, nil
	}
}

func (s *server) GetGroupProperty(_ context.Context, r *pb.QueryGroupPropertyInfo) (*pb.GroupPropertyInfo, error) {
	if result, err := s.gdb.GetGroupProperty(r.GetGroupName(), r.GetCondition()); err != nil {
		return &pb.GroupPropertyInfo{}, err
	} else {
		return &pb.GroupPropertyInfo{ItemCount: result.ItemCount, ItemColumnNames: result.ItemColumnNames}, nil
	}
}

func (s *server) UpdateGroupNames(_ context.Context, r *pb.UpdatedGroupNamesInfo) (*pb.TimeRows, error) {
	st := time.Now()
	g := []UpdatedGroupNameInfo{}
	for _, info := range r.GetInfos() {
		g = append(g, UpdatedGroupNameInfo{NewGroupName: info.GetNewGroupName(), OldGroupName: info.GetOldGroupName()})
	}
	if result, err := s.gdb.UpdateGroupNames(g...); err != nil {
		return &pb.TimeRows{}, err
	} else {
		return &pb.TimeRows{EffectedRows: int32(result.EffectedRows), Times: time.Since(st).Milliseconds()}, nil
	}
}

func (s *server) UpdateGroupColumnNames(_ context.Context, r *pb.UpdatedGroupColumnNamesInfo) (*pb.TimeCols, error) {
	st := time.Now()
	g := UpdatedGroupColumnNamesInfo{GroupName: r.GetGroupName(), OldColumnNames: r.GetOldColumnNames(), NewColumnNames: r.GetNewColumnNames()}
	if result, err := s.gdb.UpdateGroupColumnNames(g); err != nil {
		return &pb.TimeCols{}, err
	} else {
		return &pb.TimeCols{EffectedCols: int32(result.EffectedCols), Times: time.Since(st).Milliseconds()}, nil
	}
}

func (s *server) DeleteGroupColumns(_ context.Context, r *pb.DeletedGroupColumnNamesInfo) (*pb.TimeCols, error) {
	st := time.Now()
	g := DeletedGroupColumnNamesInfo{GroupName: r.GetGroupName(), ColumnNames: r.GetColumnNames()}
	if result, err := s.gdb.DeleteGroupColumns(g); err != nil {
		return &pb.TimeCols{}, err
	} else {
		return &pb.TimeCols{EffectedCols: int32(result.EffectedCols), Times: time.Since(st).Milliseconds()}, nil
	}
}

func (s *server) AddGroupColumns(_ context.Context, r *pb.AddedGroupColumnsInfo) (*pb.TimeCols, error) {
	st := time.Now()
	g := AddedGroupColumnsInfo{GroupName: r.GetGroupName(), ColumnNames: r.GetColumnNames(), DefaultValues: r.GetDefaultValues()}
	if result, err := s.gdb.AddGroupColumns(g); err != nil {
		return &pb.TimeCols{}, err
	} else {
		return &pb.TimeCols{EffectedCols: int32(result.EffectedCols), Times: time.Since(st).Milliseconds()}, nil
	}
}

// item handler

func (s *server) AddItems(_ context.Context, r *pb.AddedItemsInfo) (*pb.TimeRows, error) {
	st := time.Now()
	values := []map[string]string{}
	_ = json.Unmarshal(convertStringToByte(r.GetItemValues()), &values)
	g := AddedItemsInfo{
		GroupName:  r.GetGroupName(),
		ItemValues: values,
	}
	if result, err := s.gdb.AddItems(g); err != nil {
		return &pb.TimeRows{}, err
	} else {
		return &pb.TimeRows{EffectedRows: int32(result.EffectedRows), Times: time.Since(st).Milliseconds()}, nil
	}
}

func (s *server) DeleteItems(_ context.Context, r *pb.DeletedItemsInfo) (*pb.TimeRows, error) {
	st := time.Now()
	g := DeletedItemsInfo{
		GroupName: r.GetGroupName(),
		Condition: r.GetCondition(),
	}
	if result, err := s.gdb.DeleteItems(g); err != nil {
		return &pb.TimeRows{}, err
	} else {
		return &pb.TimeRows{EffectedRows: int32(result.EffectedRows), Times: time.Since(st).Milliseconds()}, nil
	}
}

func (s *server) GetItemsWithCount(_ context.Context, r *pb.ItemsInfo) (*pb.GdbItemsWithCount, error) {
	g := ItemsInfo{
		GroupName:   r.GetGroupName(),
		Condition:   r.GetCondition(),
		ColumnNames: r.GetColumnNames(),
		StartRow:    r.GetStartRow(),
		RowCount:    r.GetRowCount(),
	}
	if result, err := s.gdb.getItemsWithCount(g); err != nil {
		return &pb.GdbItemsWithCount{}, err
	} else {
		v := []*pb.GdbItems{}
		for i := 0; i < len(result.ItemValues); i++ {
			v = append(v, &pb.GdbItems{Items: result.ItemValues[i]})
		}
		return &pb.GdbItemsWithCount{ItemValues: v, ItemCount: result.ItemCount}, nil
	}
}

func (s *server) UpdateItems(_ context.Context, r *pb.UpdatedItemsInfo) (*pb.TimeRows, error) {
	st := time.Now()
	g := UpdatedItemsInfo{
		GroupName: r.GetGroupName(),
		Condition: r.GetCondition(),
		Clause:    r.GetClause(),
	}
	if result, err := s.gdb.UpdateItems(g); err != nil {
		return &pb.TimeRows{}, err
	} else {
		return &pb.TimeRows{EffectedRows: int32(result.EffectedRows), Times: time.Since(st).Milliseconds()}, nil
	}
}

func (s *server) CheckItems(_ context.Context, r *pb.CheckItemsInfo) (*emptypb.Empty, error) {
	if err := s.gdb.CheckItems(r.GetGroupName(), r.GetItemNames()...); err != nil {
		return &emptypb.Empty{}, err
	} else {
		return &emptypb.Empty{}, nil
	}
}

func (s *server) CleanGroupItems(_ context.Context, r *pb.GroupNamesInfo) (*pb.TimeRows, error) {
	st := time.Now()
	if result, err := s.gdb.CleanGroupItems(r.GroupNames...); err != nil {
		return &pb.TimeRows{}, err
	} else {
		return &pb.TimeRows{EffectedRows: int32(result.EffectedRows), Times: time.Since(st).Milliseconds()}, nil
	}
}

// data handler

func (s *server) BatchWriteFloatData(_ context.Context, r *pb.FloatItemValues) (*pb.TimeRows, error) {
	var itemNames [][]string
	var itemValues [][]float32
	for i := 0; i < len(r.ItemNames); i++ {
		itemNames = append(itemNames, r.GetItemNames()[i].ItemName)
		itemValues = append(itemValues, r.ItemValues[i].ItemValue)
	}
	if result, err := s.gdb.BatchWriteFloatData(r.GroupNames, itemNames, itemValues); err != nil {
		return &pb.TimeRows{}, err
	} else {
		return &pb.TimeRows{EffectedRows: int32(result.EffectedRows), Times: result.Times}, nil
	}
}

// BatchWriteFloatDataWithStream write data with client stream
func (s *server) BatchWriteFloatDataWithStream(stream pb.Data_BatchWriteFloatDataWithStreamServer) error {
	bs := []floatItemValues{}
	for {
		b, err := stream.Recv()
		if err == io.EOF {
			st := time.Now()
			eg := errgroup.Group{}
			for _, ss := range bs {
				writingString := ss
				eg.Go(func() error {
					if _, err := s.gdb.BatchWriteFloatData(writingString.GroupNames, writingString.ItemNames, writingString.ItemValues); err != nil {
						return fmt.Errorf("writing error :" + err.Error())
					} else {
						return nil
					}
				})
			}
			if err := eg.Wait(); err != nil {
				return err
			} else {
				return stream.SendAndClose(&pb.TimeRows{EffectedRows: int32(len(bs)), Times: time.Since(st).Milliseconds()})
			}
		} else if err != nil {
			return err
		} else {
			var itemNames [][]string
			var itemValues [][]float32
			for i := 0; i < len(b.ItemNames); i++ {
				itemNames = append(itemNames, b.GetItemNames()[i].ItemName)
				itemValues = append(itemValues, b.ItemValues[i].ItemValue)
			}
			bs = append(bs, floatItemValues{
				GroupNames: b.GroupNames,
				ItemNames:  itemNames,
				ItemValues: itemValues,
			})
		}
	}
}

func (s *server) BatchWriteIntData(_ context.Context, r *pb.IntItemValues) (*pb.TimeRows, error) {
	var itemNames [][]string
	var itemValues [][]int32
	for i := 0; i < len(r.ItemNames); i++ {
		itemNames = append(itemNames, r.GetItemNames()[i].ItemName)
		itemValues = append(itemValues, r.ItemValues[i].ItemValue)
	}
	if result, err := s.gdb.BatchWriteIntData(r.GroupNames, itemNames, itemValues); err != nil {
		return &pb.TimeRows{}, err
	} else {
		return &pb.TimeRows{EffectedRows: int32(result.EffectedRows), Times: result.Times}, nil
	}
}

// BatchWriteIntDataWithStream write data with client stream
func (s *server) BatchWriteIntDataWithStream(stream pb.Data_BatchWriteIntDataWithStreamServer) error {
	bs := []intItemValues{}
	for {
		b, err := stream.Recv()
		if err == io.EOF {
			st := time.Now()
			eg := errgroup.Group{}
			for _, ss := range bs {
				writingString := ss
				eg.Go(func() error {
					if _, err := s.gdb.BatchWriteIntData(writingString.GroupNames, writingString.ItemNames, writingString.ItemValues); err != nil {
						return fmt.Errorf("writing error :" + err.Error())
					} else {
						return nil
					}
				})
			}
			if err := eg.Wait(); err != nil {
				return err
			} else {
				return stream.SendAndClose(&pb.TimeRows{EffectedRows: int32(len(bs)), Times: time.Since(st).Milliseconds()})
			}
		} else if err != nil {
			return err
		} else {
			var itemNames [][]string
			var itemValues [][]int32
			for i := 0; i < len(b.ItemNames); i++ {
				itemNames = append(itemNames, b.GetItemNames()[i].ItemName)
				itemValues = append(itemValues, b.ItemValues[i].ItemValue)
			}
			bs = append(bs, intItemValues{
				GroupNames: b.GroupNames,
				ItemNames:  itemNames,
				ItemValues: itemValues,
			})
		}
	}
}

func (s *server) BatchWriteStringData(_ context.Context, r *pb.StringItemValues) (*pb.TimeRows, error) {
	var itemNames [][]string
	var itemValues [][]string
	for i := 0; i < len(r.ItemNames); i++ {
		itemNames = append(itemNames, r.GetItemNames()[i].ItemName)
		itemValues = append(itemValues, r.ItemValues[i].ItemValue)
	}
	if result, err := s.gdb.BatchWriteStringData(r.GroupNames, itemNames, itemValues); err != nil {
		return &pb.TimeRows{}, err
	} else {
		return &pb.TimeRows{EffectedRows: int32(result.EffectedRows), Times: result.Times}, nil
	}
}

// BatchWriteStringDataWithStream write data with client stream
func (s *server) BatchWriteStringDataWithStream(stream pb.Data_BatchWriteStringDataWithStreamServer) error {
	bs := []stringItemValues{}
	for {
		b, err := stream.Recv()
		if err == io.EOF {
			st := time.Now()
			eg := errgroup.Group{}
			for _, ss := range bs {
				writingString := ss
				eg.Go(func() error {
					if _, err := s.gdb.BatchWriteStringData(writingString.GroupNames, writingString.ItemNames, writingString.ItemValues); err != nil {
						return fmt.Errorf("writing error :" + err.Error())
					} else {
						return nil
					}
				})
			}
			if err := eg.Wait(); err != nil {
				return err
			} else {
				return stream.SendAndClose(&pb.TimeRows{EffectedRows: int32(len(bs)), Times: time.Since(st).Milliseconds()})
			}
		} else if err != nil {
			return err
		} else {
			var itemNames [][]string
			var itemValues [][]string
			for i := 0; i < len(b.ItemNames); i++ {
				itemNames = append(itemNames, b.GetItemNames()[i].ItemName)
				itemValues = append(itemValues, b.ItemValues[i].ItemValue)
			}
			bs = append(bs, stringItemValues{
				GroupNames: b.GroupNames,
				ItemNames:  itemNames,
				ItemValues: itemValues,
			})
		}
	}
}

func (s *server) BatchWriteBoolData(_ context.Context, r *pb.BoolItemValues) (*pb.TimeRows, error) {
	var itemNames [][]string
	var itemValues [][]bool
	for i := 0; i < len(r.ItemNames); i++ {
		itemNames = append(itemNames, r.GetItemNames()[i].ItemName)
		itemValues = append(itemValues, r.ItemValues[i].ItemValue)
	}
	if result, err := s.gdb.BatchWriteBoolData(r.GroupNames, itemNames, itemValues); err != nil {
		return &pb.TimeRows{}, err
	} else {
		return &pb.TimeRows{EffectedRows: int32(result.EffectedRows), Times: result.Times}, nil
	}
}

// BatchWriteBoolDataWithStream write data with client stream
func (s *server) BatchWriteBoolDataWithStream(stream pb.Data_BatchWriteBoolDataWithStreamServer) error {
	bs := []boolItemValues{}
	for {
		b, err := stream.Recv()
		if err == io.EOF {
			st := time.Now()
			eg := errgroup.Group{}
			for _, ss := range bs {
				writingString := ss
				eg.Go(func() error {
					if _, err := s.gdb.BatchWriteBoolData(writingString.GroupNames, writingString.ItemNames, writingString.ItemValues); err != nil {
						return fmt.Errorf("writing error :" + err.Error())
					} else {
						return nil
					}
				})
			}
			if err := eg.Wait(); err != nil {
				return err
			} else {
				return stream.SendAndClose(&pb.TimeRows{EffectedRows: int32(len(bs)), Times: time.Since(st).Milliseconds()})
			}
		} else if err != nil {
			return err
		} else {
			var itemNames [][]string
			var itemValues [][]bool
			for i := 0; i < len(b.ItemNames); i++ {
				itemNames = append(itemNames, b.GetItemNames()[i].ItemName)
				itemValues = append(itemValues, b.ItemValues[i].ItemValue)
			}
			bs = append(bs, boolItemValues{
				GroupNames: b.GroupNames,
				ItemNames:  itemNames,
				ItemValues: itemValues,
			})
		}
	}
}

func (s *server) BatchWriteFloatHistoricalData(_ context.Context, r *pb.FloatHItemValues) (*pb.TimeRows, error) {
	timeStamps := [][]int32{}
	itemValues := [][]float32{}
	for i := 0; i < len(r.TimeStamps); i++ {
		timeStamps = append(timeStamps, r.TimeStamps[i].TimeStamps)
		itemValues = append(itemValues, r.ItemValues[i].ItemValue)
	}
	if result, err := s.gdb.BatchWriteFloatHistoricalData(r.GroupNames, r.ItemNames, timeStamps, itemValues); err != nil {
		return &pb.TimeRows{}, err
	} else {
		return &pb.TimeRows{EffectedRows: int32(result.EffectedRows), Times: result.Times}, nil
	}
}

func (s *server) BatchWriteFloatHistoricalDataWithStream(stream pb.Data_BatchWriteFloatHistoricalDataWithStreamServer) error {
	bs := []floatHItemValues{}
	st := time.Now()
	for {
		b, err := stream.Recv()
		if err == io.EOF {
			eg := errgroup.Group{}
			for _, ss := range bs {
				ws := ss
				eg.Go(func() error {
					if _, err := s.gdb.BatchWriteFloatHistoricalData(ws.GroupNames, ws.ItemNames, ws.TimeStamps, ws.ItemValues); err != nil {
						return err
					}
					return nil
				})
				if err := eg.Wait(); err != nil {
					return err
				} else {
					return stream.SendAndClose(&pb.TimeRows{Times: time.Since(st).Milliseconds()})
				}
			}
		} else if err != nil {
			return err
		} else {
			timeStamps := [][]int32{}
			itemValues := [][]float32{}
			for i := 0; i < len(b.TimeStamps); i++ {
				timeStamps = append(timeStamps, b.TimeStamps[i].TimeStamps)
				itemValues = append(itemValues, b.ItemValues[i].ItemValue)
			}
			bs = append(bs, floatHItemValues{
				GroupNames: b.GroupNames,
				ItemNames:  b.ItemNames,
				ItemValues: itemValues,
				TimeStamps: timeStamps,
			})
		}
	}
}

func (s *server) BatchWriteIntHistoricalData(_ context.Context, r *pb.IntHItemValues) (*pb.TimeRows, error) {
	timeStamps := [][]int32{}
	itemValues := [][]int32{}
	for i := 0; i < len(r.TimeStamps); i++ {
		timeStamps = append(timeStamps, r.TimeStamps[i].TimeStamps)
		itemValues = append(itemValues, r.ItemValues[i].ItemValue)
	}
	if result, err := s.gdb.BatchWriteIntHistoricalData(r.GroupNames, r.ItemNames, timeStamps, itemValues); err != nil {
		return &pb.TimeRows{}, err
	} else {
		return &pb.TimeRows{EffectedRows: int32(result.EffectedRows), Times: result.Times}, nil
	}
}

func (s *server) BatchWriteIntHistoricalDataWithStream(stream pb.Data_BatchWriteIntHistoricalDataWithStreamServer) error {
	bs := []intHItemValues{}
	st := time.Now()
	for {
		b, err := stream.Recv()
		if err == io.EOF {
			eg := errgroup.Group{}
			for _, ss := range bs {
				ws := ss
				eg.Go(func() error {
					if _, err := s.gdb.BatchWriteIntHistoricalData(ws.GroupNames, ws.ItemNames, ws.TimeStamps, ws.ItemValues); err != nil {
						return err
					}
					return nil
				})
				if err := eg.Wait(); err != nil {
					return err
				} else {
					return stream.SendAndClose(&pb.TimeRows{Times: time.Since(st).Milliseconds()})
				}
			}
		} else if err != nil {
			return err
		} else {
			timeStamps := [][]int32{}
			itemValues := [][]int32{}
			for i := 0; i < len(b.TimeStamps); i++ {
				timeStamps = append(timeStamps, b.TimeStamps[i].TimeStamps)
				itemValues = append(itemValues, b.ItemValues[i].ItemValue)
			}
			bs = append(bs, intHItemValues{
				GroupNames: b.GroupNames,
				ItemNames:  b.ItemNames,
				ItemValues: itemValues,
				TimeStamps: timeStamps,
			})
		}
	}
}

func (s *server) BatchWriteStringHistoricalData(_ context.Context, r *pb.StringHItemValues) (*pb.TimeRows, error) {
	timeStamps := [][]int32{}
	itemValues := [][]string{}
	for i := 0; i < len(r.TimeStamps); i++ {
		timeStamps = append(timeStamps, r.TimeStamps[i].TimeStamps)
		itemValues = append(itemValues, r.ItemValues[i].ItemValue)
	}
	if result, err := s.gdb.BatchWriteStringHistoricalData(r.GroupNames, r.ItemNames, timeStamps, itemValues); err != nil {
		return &pb.TimeRows{}, err
	} else {
		return &pb.TimeRows{EffectedRows: int32(result.EffectedRows), Times: result.Times}, nil
	}
}

func (s *server) BatchWriteStringHistoricalDataWithStream(stream pb.Data_BatchWriteStringHistoricalDataWithStreamServer) error {
	bs := []stringHItemValues{}
	st := time.Now()
	for {
		b, err := stream.Recv()
		if err == io.EOF {
			eg := errgroup.Group{}
			for _, ss := range bs {
				ws := ss
				eg.Go(func() error {
					if _, err := s.gdb.BatchWriteStringHistoricalData(ws.GroupNames, ws.ItemNames, ws.TimeStamps, ws.ItemValues); err != nil {
						return err
					}
					return nil
				})
				if err := eg.Wait(); err != nil {
					return err
				} else {
					return stream.SendAndClose(&pb.TimeRows{Times: time.Since(st).Milliseconds()})
				}
			}
		} else if err != nil {
			return err
		} else {
			timeStamps := [][]int32{}
			itemValues := [][]string{}
			for i := 0; i < len(b.TimeStamps); i++ {
				timeStamps = append(timeStamps, b.TimeStamps[i].TimeStamps)
				itemValues = append(itemValues, b.ItemValues[i].ItemValue)
			}
			bs = append(bs, stringHItemValues{
				GroupNames: b.GroupNames,
				ItemNames:  b.ItemNames,
				ItemValues: itemValues,
				TimeStamps: timeStamps,
			})
		}
	}
}

func (s *server) BatchWriteBoolHistoricalData(_ context.Context, r *pb.BoolHItemValues) (*pb.TimeRows, error) {
	timeStamps := [][]int32{}
	itemValues := [][]bool{}
	for i := 0; i < len(r.TimeStamps); i++ {
		timeStamps = append(timeStamps, r.TimeStamps[i].TimeStamps)
		itemValues = append(itemValues, r.ItemValues[i].ItemValue)
	}
	if result, err := s.gdb.BatchWriteBoolHistoricalData(r.GroupNames, r.ItemNames, timeStamps, itemValues); err != nil {
		return &pb.TimeRows{}, err
	} else {
		return &pb.TimeRows{EffectedRows: int32(result.EffectedRows), Times: result.Times}, nil
	}
}

func (s *server) BatchWriteBoolHistoricalDataWithStream(stream pb.Data_BatchWriteBoolHistoricalDataWithStreamServer) error {
	bs := []boolHItemValues{}
	st := time.Now()
	for {
		b, err := stream.Recv()
		if err == io.EOF {
			eg := errgroup.Group{}
			for _, ss := range bs {
				ws := ss
				eg.Go(func() error {
					if _, err := s.gdb.BatchWriteBoolHistoricalData(ws.GroupNames, ws.ItemNames, ws.TimeStamps, ws.ItemValues); err != nil {
						return err
					}
					return nil
				})
				if err := eg.Wait(); err != nil {
					return err
				} else {
					return stream.SendAndClose(&pb.TimeRows{Times: time.Since(st).Milliseconds()})
				}
			}
		} else if err != nil {
			return err
		} else {
			timeStamps := [][]int32{}
			itemValues := [][]bool{}
			for i := 0; i < len(b.TimeStamps); i++ {
				timeStamps = append(timeStamps, b.TimeStamps[i].TimeStamps)
				itemValues = append(itemValues, b.ItemValues[i].ItemValue)
			}
			bs = append(bs, boolHItemValues{
				GroupNames: b.GroupNames,
				ItemNames:  b.ItemNames,
				ItemValues: itemValues,
				TimeStamps: timeStamps,
			})
		}
	}
}

func (s *server) GetRealTimeData(_ context.Context, r *pb.QueryRealTimeDataString) (*pb.GdbRealTimeData, error) {
	if result, err := s.gdb.GetRealTimeData(r.GetGroupNames(), r.ItemNames); err != nil {
		return &pb.GdbRealTimeData{}, err
	} else {
		v, _ := json.Marshal(result.RealTimeData)
		return &pb.GdbRealTimeData{RealTimeData: string(v), Times: result.Times}, nil
	}
}

func (s *server) GetFloatHistoricalData(_ context.Context, r *pb.QueryHistoricalDataString) (*pb.GdbHistoricalData, error) {
	if result, err := s.gdb.GetFloatHistoricalData(r.GetGroupNames(), r.GetItemNames(), r.GetStartTimes(), r.GetEndTimes(), r.GetIntervals()); err != nil {
		return &pb.GdbHistoricalData{}, err
	} else {
		v, _ := json.Marshal(result.HistoricalData)
		return &pb.GdbHistoricalData{HistoricalData: string(v), Times: result.Times}, nil
	}
}

func (s *server) GetIntHistoricalData(_ context.Context, r *pb.QueryHistoricalDataString) (*pb.GdbHistoricalData, error) {
	if result, err := s.gdb.GetIntHistoricalData(r.GetGroupNames(), r.GetItemNames(), r.GetStartTimes(), r.GetEndTimes(), r.GetIntervals()); err != nil {
		return &pb.GdbHistoricalData{}, err
	} else {
		v, _ := json.Marshal(result.HistoricalData)
		return &pb.GdbHistoricalData{HistoricalData: string(v), Times: result.Times}, nil
	}
}

func (s *server) GetStringHistoricalData(_ context.Context, r *pb.QueryHistoricalDataString) (*pb.GdbHistoricalData, error) {
	if result, err := s.gdb.GetStringHistoricalData(r.GetGroupNames(), r.GetItemNames(), r.GetStartTimes(), r.GetEndTimes(), r.GetIntervals()); err != nil {
		return &pb.GdbHistoricalData{}, err
	} else {
		v, _ := json.Marshal(result.HistoricalData)
		return &pb.GdbHistoricalData{HistoricalData: string(v), Times: result.Times}, nil
	}
}

func (s *server) GetBoolHistoricalData(_ context.Context, r *pb.QueryHistoricalDataString) (*pb.GdbHistoricalData, error) {
	if result, err := s.gdb.GetBoolHistoricalData(r.GetGroupNames(), r.GetItemNames(), r.GetStartTimes(), r.GetEndTimes(), r.GetIntervals()); err != nil {
		return &pb.GdbHistoricalData{}, err
	} else {
		v, _ := json.Marshal(result.HistoricalData)
		return &pb.GdbHistoricalData{HistoricalData: string(v), Times: result.Times}, nil
	}
}

func (s *server) GetFloatRawHistoricalData(_ context.Context, r *pb.QueryRawHistoricalDataString) (*pb.GdbHistoricalData, error) {
	if result, err := s.gdb.GetFloatRawHistoricalData(r.GroupNames, r.ItemNames); err != nil {
		return &pb.GdbHistoricalData{}, err
	} else {
		v, _ := json.Marshal(result.HistoricalData)
		return &pb.GdbHistoricalData{HistoricalData: string(v), Times: result.Times}, nil
	}
}

func (s *server) GetIntRawHistoricalData(_ context.Context, r *pb.QueryRawHistoricalDataString) (*pb.GdbHistoricalData, error) {
	if result, err := s.gdb.GetIntRawHistoricalData(r.GroupNames, r.ItemNames); err != nil {
		return &pb.GdbHistoricalData{}, err
	} else {
		v, _ := json.Marshal(result.HistoricalData)
		return &pb.GdbHistoricalData{HistoricalData: string(v), Times: result.Times}, nil
	}
}

func (s *server) GetStringRawHistoricalData(_ context.Context, r *pb.QueryRawHistoricalDataString) (*pb.GdbHistoricalData, error) {
	if result, err := s.gdb.GetStringRawHistoricalData(r.GroupNames, r.ItemNames); err != nil {
		return &pb.GdbHistoricalData{}, err
	} else {
		v, _ := json.Marshal(result.HistoricalData)
		return &pb.GdbHistoricalData{HistoricalData: string(v), Times: result.Times}, nil
	}
}

func (s *server) GetBoolRawHistoricalData(_ context.Context, r *pb.QueryRawHistoricalDataString) (*pb.GdbHistoricalData, error) {
	if result, err := s.gdb.GetBoolRawHistoricalData(r.GroupNames, r.ItemNames); err != nil {
		return &pb.GdbHistoricalData{}, err
	} else {
		v, _ := json.Marshal(result.HistoricalData)
		return &pb.GdbHistoricalData{HistoricalData: string(v), Times: result.Times}, nil
	}
}

func (s *server) GetFloatHistoricalDataWithStamp(_ context.Context, r *pb.QueryHistoricalDataWithStampString) (*pb.GdbHistoricalData, error) {
	var groupNames, itemNames []string
	timeStamps := [][]int32{}
	values := r.QueryString
	{
		for i := 0; i < len(values); i++ {
			groupNames = append(groupNames, values[i].GroupName)
			itemNames = append(itemNames, values[i].ItemName)
			timeStamps = append(timeStamps, values[i].TimeStamps)
		}
	}
	if result, err := s.gdb.GetFloatHistoricalDataWithStamp(groupNames, itemNames, timeStamps); err != nil {
		return &pb.GdbHistoricalData{}, err
	} else {
		v, _ := json.Marshal(result.HistoricalData)
		return &pb.GdbHistoricalData{HistoricalData: string(v), Times: result.Times}, nil
	}
}

func (s *server) GetIntHistoricalDataWithStamp(_ context.Context, r *pb.QueryHistoricalDataWithStampString) (*pb.GdbHistoricalData, error) {
	var groupNames, itemNames []string
	timeStamps := [][]int32{}
	values := r.QueryString
	{
		for i := 0; i < len(values); i++ {
			groupNames = append(groupNames, values[i].GroupName)
			itemNames = append(itemNames, values[i].ItemName)
			timeStamps = append(timeStamps, values[i].TimeStamps)
		}
	}
	if result, err := s.gdb.GetIntHistoricalDataWithStamp(groupNames, itemNames, timeStamps); err != nil {
		return &pb.GdbHistoricalData{}, err
	} else {
		v, _ := json.Marshal(result.HistoricalData)
		return &pb.GdbHistoricalData{HistoricalData: string(v), Times: result.Times}, nil
	}
}

func (s *server) GetStringHistoricalDataWithStamp(_ context.Context, r *pb.QueryHistoricalDataWithStampString) (*pb.GdbHistoricalData, error) {
	var groupNames, itemNames []string
	timeStamps := [][]int32{}
	values := r.QueryString
	{
		for i := 0; i < len(values); i++ {
			groupNames = append(groupNames, values[i].GroupName)
			itemNames = append(itemNames, values[i].ItemName)
			timeStamps = append(timeStamps, values[i].TimeStamps)
		}
	}
	if result, err := s.gdb.GetStringHistoricalDataWithStamp(groupNames, itemNames, timeStamps); err != nil {
		return &pb.GdbHistoricalData{}, err
	} else {
		v, _ := json.Marshal(result.HistoricalData)
		return &pb.GdbHistoricalData{HistoricalData: string(v), Times: result.Times}, nil
	}
}

func (s *server) GetBoolHistoricalDataWithStamp(_ context.Context, r *pb.QueryHistoricalDataWithStampString) (*pb.GdbHistoricalData, error) {
	var groupNames, itemNames []string
	timeStamps := [][]int32{}
	values := r.QueryString
	{
		for i := 0; i < len(values); i++ {
			groupNames = append(groupNames, values[i].GroupName)
			itemNames = append(itemNames, values[i].ItemName)
			timeStamps = append(timeStamps, values[i].TimeStamps)
		}
	}
	if result, err := s.gdb.GetBoolHistoricalDataWithStamp(groupNames, itemNames, timeStamps); err != nil {
		return &pb.GdbHistoricalData{}, err
	} else {
		v, _ := json.Marshal(result.HistoricalData)
		return &pb.GdbHistoricalData{HistoricalData: string(v), Times: result.Times}, nil
	}
}

func (s *server) GetFloatHistoricalDataWithCondition(_ context.Context, r *pb.QueryHistoricalDataWithConditionString) (*pb.GdbHistoricalData, error) {
	dz := []DeadZone{}
	for _, zone := range r.GetDeadZones() {
		dz = append(dz, DeadZone{
			ItemName:      zone.ItemName,
			DeadZoneCount: zone.DeadZoneCount,
		})
	}
	if result, err := s.gdb.GetFloatHistoricalDataWithCondition(r.GetGroupName(), r.GetItemNames(), r.GetStartTime(),
		r.GetEndTime(), r.GetInterval(), r.GetFilterCondition(), dz); err != nil {
		return &pb.GdbHistoricalData{}, err
	} else {
		v, _ := json.Marshal(result.HistoricalData)
		return &pb.GdbHistoricalData{HistoricalData: string(v), Times: result.Times}, nil
	}
}

func (s *server) GetIntHistoricalDataWithCondition(_ context.Context, r *pb.QueryHistoricalDataWithConditionString) (*pb.GdbHistoricalData, error) {
	dz := []DeadZone{}
	for _, zone := range r.GetDeadZones() {
		dz = append(dz, DeadZone{
			ItemName:      zone.ItemName,
			DeadZoneCount: zone.DeadZoneCount,
		})
	}
	if result, err := s.gdb.GetIntHistoricalDataWithCondition(r.GetGroupName(), r.GetItemNames(), r.GetStartTime(),
		r.GetEndTime(), r.GetInterval(), r.GetFilterCondition(), dz); err != nil {
		return &pb.GdbHistoricalData{}, err
	} else {
		v, _ := json.Marshal(result.HistoricalData)
		return &pb.GdbHistoricalData{HistoricalData: string(v), Times: result.Times}, nil
	}
}

func (s *server) GetStringHistoricalDataWithCondition(_ context.Context, r *pb.QueryHistoricalDataWithConditionString) (*pb.GdbHistoricalData, error) {
	dz := []DeadZone{}
	for _, zone := range r.GetDeadZones() {
		dz = append(dz, DeadZone{
			ItemName:      zone.ItemName,
			DeadZoneCount: zone.DeadZoneCount,
		})
	}
	if result, err := s.gdb.GetStringHistoricalDataWithCondition(r.GetGroupName(), r.GetItemNames(), r.GetStartTime(),
		r.GetEndTime(), r.GetInterval(), r.GetFilterCondition(), dz); err != nil {
		return &pb.GdbHistoricalData{}, err
	} else {
		v, _ := json.Marshal(result.HistoricalData)
		return &pb.GdbHistoricalData{HistoricalData: string(v), Times: result.Times}, nil
	}
}

func (s *server) GetBoolHistoricalDataWithCondition(_ context.Context, r *pb.QueryHistoricalDataWithConditionString) (*pb.GdbHistoricalData, error) {
	dz := []DeadZone{}
	for _, zone := range r.GetDeadZones() {
		dz = append(dz, DeadZone{
			ItemName:      zone.ItemName,
			DeadZoneCount: zone.DeadZoneCount,
		})
	}
	if result, err := s.gdb.GetBoolHistoricalDataWithCondition(r.GetGroupName(), r.GetItemNames(), r.GetStartTime(),
		r.GetEndTime(), r.GetInterval(), r.GetFilterCondition(), dz); err != nil {
		return &pb.GdbHistoricalData{}, err
	} else {
		v, _ := json.Marshal(result.HistoricalData)
		return &pb.GdbHistoricalData{HistoricalData: string(v), Times: result.Times}, nil
	}
}

func (s *server) DeleteFloatHistoricalData(_ context.Context, r *pb.DeleteHistoricalDataString) (*pb.TimeRows, error) {
	if result, err := s.gdb.DeleteFloatHistoricalData(r.GetGroupNames(), r.GetItemNames(), r.GetStartTimes(), r.GetEndTimes()); err != nil {
		return &pb.TimeRows{}, err
	} else {
		return &pb.TimeRows{EffectedRows: int32(result.EffectedRows), Times: result.Times}, nil
	}
}

func (s *server) DeleteIntHistoricalData(_ context.Context, r *pb.DeleteHistoricalDataString) (*pb.TimeRows, error) {
	if result, err := s.gdb.DeleteIntHistoricalData(r.GetGroupNames(), r.GetItemNames(), r.GetStartTimes(), r.GetEndTimes()); err != nil {
		return &pb.TimeRows{}, err
	} else {
		return &pb.TimeRows{EffectedRows: int32(result.EffectedRows), Times: result.Times}, nil
	}
}

func (s *server) DeleteStringHistoricalData(_ context.Context, r *pb.DeleteHistoricalDataString) (*pb.TimeRows, error) {
	if result, err := s.gdb.DeleteStringHistoricalData(r.GetGroupNames(), r.GetItemNames(), r.GetStartTimes(), r.GetEndTimes()); err != nil {
		return &pb.TimeRows{}, err
	} else {
		return &pb.TimeRows{EffectedRows: int32(result.EffectedRows), Times: result.Times}, nil
	}
}

func (s *server) DeleteBoolHistoricalData(_ context.Context, r *pb.DeleteHistoricalDataString) (*pb.TimeRows, error) {
	if result, err := s.gdb.DeleteBoolHistoricalData(r.GetGroupNames(), r.GetItemNames(), r.GetStartTimes(), r.GetEndTimes()); err != nil {
		return &pb.TimeRows{}, err
	} else {
		return &pb.TimeRows{EffectedRows: int32(result.EffectedRows), Times: result.Times}, nil
	}
}

func (s *server) CleanItemData(_ context.Context, r *pb.DeletedItemsInfo) (*pb.TimeRows, error) {
	if result, err := s.gdb.CleanItemData(DeletedItemsInfo{
		GroupName: r.GroupName,
		Condition: r.Condition,
	}); err != nil {
		return &pb.TimeRows{}, nil
	} else {
		return &pb.TimeRows{EffectedRows: int32(result.EffectedRows), Times: result.Times}, nil
	}

}

func (s *server) ReLoadDb(_ context.Context, _ *emptypb.Empty) (*pb.TimeRows, error) {
	if result, err := s.gdb.ReLoadDb(); err != nil {
		return &pb.TimeRows{}, err
	} else {
		return &pb.TimeRows{Times: result.Times, EffectedRows: int32(result.EffectedRows)}, nil
	}
}

func (s *server) GetDbSize(_ context.Context, _ *emptypb.Empty) (*pb.FileSize, error) {
	if result, err := s.gdb.getDbSize(); err != nil {
		return &pb.FileSize{}, err
	} else {
		r, _ := json.Marshal(result)
		return &pb.FileSize{FileSize: string(r)}, nil
	}
}

func (s *server) GetDbInfo(_ context.Context, _ *emptypb.Empty) (*pb.GdbInfoData, error) {
	if result, err := s.gdb.getDbInfo(); err != nil {
		return &pb.GdbInfoData{}, err
	} else {
		v, _ := json.Marshal(result)
		return &pb.GdbInfoData{Info: fmt.Sprintf("%s", v)}, nil
	}
}

func (s *server) GetDbInfoHistory(_ context.Context, r *pb.QuerySpeedHistoryDataString) (*pb.GdbHistoricalData, error) {
	if r, err := s.gdb.getDbInfoHistory(r.GetInfoType(), r.GetItemName(), r.GetStartTimes(), r.GetEndTimes(), r.GetIntervals()); err != nil {
		return &pb.GdbHistoricalData{}, err
	} else {
		result, _ := json.Marshal(r)
		return &pb.GdbHistoricalData{HistoricalData: fmt.Sprintf("%s", result)}, nil
	}
}

func (s *server) GetRoutes(_ context.Context, _ *emptypb.Empty) (*pb.Routes, error) {
	if r, err := s.gdb.getRoutes(); err != nil {
		return &pb.Routes{}, err
	} else {
		result, _ := json.Marshal(r)
		return &pb.Routes{Routes: string(result)}, nil
	}
}

func (s *server) DeleteRoutes(_ context.Context, r *pb.RoutesInfo) (*pb.TimeRows, error) {
	if err := s.gdb.deleteRoutes(r.GetName(), r.GetRoutes()...); err != nil {
		return &pb.TimeRows{}, err
	} else {
		return &pb.TimeRows{EffectedRows: int32(len(r.GetRoutes()))}, nil
	}
}

func (s *server) AddRoutes(_ context.Context, r *pb.RoutesInfo) (*pb.TimeRows, error) {
	if err := s.gdb.addRoutes(r.GetName(), r.GetRoutes()...); err != nil {
		return &pb.TimeRows{}, err
	} else {
		return &pb.TimeRows{EffectedRows: 1}, nil
	}
}

func (s *server) AddUserRoutes(_ context.Context, r *pb.RoutesInfo) (*pb.TimeRows, error) {
	if err := s.gdb.addUserRoutes(r.GetName(), r.GetRoutes()...); err != nil {
		return &pb.TimeRows{}, err
	} else {
		return &pb.TimeRows{EffectedRows: 1}, nil
	}
}

func (s *server) DeleteUserRoutes(_ context.Context, r *pb.UserName) (*pb.TimeRows, error) {
	if _, err := s.gdb.updateItem("delete from route_cfg where userName='" + r.GetName() + "'"); err != nil {
		return &pb.TimeRows{}, err
	} else {
		_ = s.gdb.e.LoadPolicy()
		return &pb.TimeRows{EffectedRows: 1}, err
	}
}

func (s *server) GetAllRoutes(_ context.Context, _ *emptypb.Empty) (*pb.Routes, error) {
	r, _ := json.Marshal([][]string{superUserRoutes, commonUserRoutes, visitorUserRoutes})
	return &pb.Routes{Routes: string(r)}, nil
}

func (s *server) CheckRoutes(_ context.Context, r *pb.RoutesInfo) (*pb.CheckResult, error) {
	result, _ := s.gdb.checkRoutes(r.GetName(), r.GetRoutes()...)
	return &pb.CheckResult{Result: result}, nil
}

// page handler

func (s *server) UserLogin(_ context.Context, r *pb.AuthInfo) (*pb.UserToken, error) {
	g := authInfo{
		UserName: r.GetUserName(),
		PassWord: r.GetPassWord(),
	}
	if token, err := s.gdb.userLogin(g); err != nil {
		return &pb.UserToken{}, nil
	} else {
		return &pb.UserToken{Token: token.Token}, nil
	}
}

func (s *server) UserLogOut(_ context.Context, r *pb.UserName) (*emptypb.Empty, error) {
	if _, err := s.gdb.userLogout(r.GetName()); err != nil {
		return &emptypb.Empty{}, err
	} else {
		return &emptypb.Empty{}, nil
	}
}

func (s *server) GetUserInfo(_ context.Context, r *pb.UserName) (*pb.UserInfo, error) {
	if result, err := s.gdb.getUserInfo(r.GetName()); err != nil {
		return nil, err
	} else {
		return &pb.UserInfo{
			UserName: result.UserName,
			Role:     result.Role,
		}, nil
	}
}

func (s *server) GetUsers(_ context.Context, _ *emptypb.Empty) (*pb.UserInfos, error) {
	if result, err := s.gdb.query("select * from user_cfg"); err != nil {
		return &pb.UserInfos{}, err
	} else {
		r, _ := json.Marshal(result)
		return &pb.UserInfos{UserInfos: fmt.Sprintf("%s", r)}, nil
	}
}

func (s *server) AddUsers(_ context.Context, r *pb.AddedUserInfo) (*pb.TimeRows, error) {
	g := addedUserInfo{
		Name:     r.GetName(),
		Role:     r.GetRole(),
		PassWord: r.GetPassWord(),
	}
	if result, err := s.gdb.addUsers(g); err != nil {
		return &pb.TimeRows{}, err
	} else {
		return &pb.TimeRows{EffectedRows: int32(result.EffectedRows)}, nil
	}
}

func (s *server) DeleteUsers(_ context.Context, r *pb.UserName) (*pb.TimeRows, error) {
	g := userName{Name: r.GetName()}
	if result, err := s.gdb.deleteUsers(g); err != nil {
		return &pb.TimeRows{}, err
	} else {
		return &pb.TimeRows{EffectedRows: int32(result.EffectedRows)}, nil
	}
}

func (s *server) UpdateUsers(_ context.Context, r *pb.UpdatedUserInfo) (*pb.TimeRows, error) {
	g := updatedUserInfo{
		UserName:    r.GetUserName(),
		NewUserName: r.GetNewUserName(),
		NewPassWord: r.GetNewPassWord(),
		NewRole:     r.GetNewRole(),
	}
	if result, err := s.gdb.updateUsers(g); err != nil {
		return &pb.TimeRows{}, err
	} else {
		return &pb.TimeRows{EffectedRows: int32(result.EffectedRows)}, nil
	}
}

func (s *server) GetLogs(_ context.Context, r *pb.QueryLogsInfo) (*pb.LogsInfo, error) {
	g := queryLogsInfo{
		Level:     r.GetLevel(),
		StartTime: r.GetStartTime(),
		EndTime:   r.GetEndTime(),
		StartRow:  r.GetStartRow(),
		RowCount:  r.GetRowCount(),
		Name:      r.GetName(),
	}
	if result, err := s.gdb.getLogs(g); err != nil {
		return &pb.LogsInfo{}, err
	} else {
		r, _ := json.Marshal(result.Infos)
		return &pb.LogsInfo{Infos: string(r), Count: int32(result.Count)}, nil
	}
}

func (s *server) GetJsCode(_ context.Context, r *pb.FileInfo) (*pb.Code, error) {
	if result, err := getJsCode(r.GetFileName()); err != nil {
		return &pb.Code{}, err
	} else {
		return &pb.Code{Code: result}, nil
	}
}

func (s *server) DeleteLogs(_ context.Context, r *pb.DeletedLogInfo) (*pb.TimeRows, error) {
	g := deletedLogInfo{
		Id:                r.GetId(),
		StartTime:         r.GetStartTime(),
		EndTime:           r.GetEndTime(),
		UserNameCondition: r.GetUserNameCondition(),
	}
	if result, err := s.gdb.deleteLogs(g); err != nil {
		return &pb.TimeRows{}, err
	} else {
		return &pb.TimeRows{EffectedRows: int32(result.EffectedRows)}, nil
	}

}

func (s *server) UploadFile(_ context.Context, r *pb.UploadedFileInfo) (*emptypb.Empty, error) {
	fileName, fileContents := r.GetFileName(), r.GetFile()
	contents := []uint8{}
	for _, c := range fileContents {
		contents = append(contents, uint8(c))
	}
	if err := ioutil.WriteFile("./uploadFiles/"+fileName, contents, 0644); err != nil {
		return &emptypb.Empty{}, err
	} else {
		return &emptypb.Empty{}, nil
	}
}

func (s *server) UploadFileWithStream(stream pb.Page_UploadFileWithStreamServer) error {
	contents := []uint8{}
	var fileName string
	for {
		b, err := stream.Recv()
		if err == io.EOF {
			if err1 := ioutil.WriteFile("./uploadFiles/"+fileName, contents, 0644); err1 != nil {
				return err1
			} else {
				return stream.SendAndClose(&emptypb.Empty{})
			}
		} else if err != nil {
			return err
		} else {
			fileName = b.GetFileName()
			for _, c := range b.GetFile() {
				contents = append(contents, uint8(c))
			}
		}
	}
}

func (s *server) AddItemsByExcel(_ context.Context, r *pb.FileInfo) (*pb.TimeRows, error) {
	if result, err := s.gdb.addItemsByExcel(r.GetGroupName(), "./uploadFiles/"+r.GetFileName()); err != nil {
		return &pb.TimeRows{}, err
	} else {
		return &pb.TimeRows{EffectedRows: int32(result.EffectedRows)}, nil
	}
}

func (s *server) ImportHistoryByExcel(_ context.Context, r *pb.HistoryFileInfo) (*pb.TimeRows, error) {
	if result, err := s.gdb.importHistoryByExcel("./uploadFiles/"+r.GetFileName(), r.GetGroupName(), r.GetItemNames(), r.GetSheetNames()...); err != nil {
		return &pb.TimeRows{}, err
	} else {
		return &pb.TimeRows{EffectedRows: int32(result.EffectedRows), Times: result.Times}, nil
	}
}

func (s *server) DownloadFile(_ context.Context, r *pb.FileInfo) (*pb.FileContents, error) {
	fileName := r.GetFileName()
	contents := []int32{}
	if fileContent, err := dFiles.ReadFile("templateFiles/" + fileName); err != nil {
		return &pb.FileContents{}, err
	} else {
		for _, c := range fileContent {
			contents = append(contents, int32(c))
		}
		return &pb.FileContents{Contents: contents}, nil
	}
}

// calc handler

func (s *server) TestCalcItem(_ context.Context, r *pb.TestCalcItemInfo) (*pb.CalculationResult, error) {
	if result, err := s.gdb.testCalculation(r.GetExpression()); err != nil {
		return &pb.CalculationResult{}, err
	} else {
		r, _ := json.Marshal(result.Result)
		return &pb.CalculationResult{Result: string(r)}, nil
	}
}

func (s *server) AddCalcItem(_ context.Context, r *pb.AddedCalcItemInfo) (*pb.CalculationResult, error) {
	if result, err := s.gdb.testCalculation(r.GetExpression()); err != nil {
		return nil, err
	} else {
		createTime := time.Now().Format(timeFormatString)
		if _, err := s.gdb.updateItem("insert into calc_cfg (description, expression, createTime, updatedTime, duration, status) values ('" + r.GetDescription() + "', '" + r.GetExpression() + "' , '" + createTime + "', '" + createTime + "', '" + r.GetDuration() + "', '" + r.GetFlag() + "')"); err != nil {
			return &pb.CalculationResult{}, err
		} else {
			r, _ := json.Marshal(result.Result)
			return &pb.CalculationResult{Result: string(r)}, nil
		}
	}
}

func (s *server) AddCalcItemWithStream(stream pb.Calc_AddCalcItemWithStreamServer) error {
	as := []addedCalcItemInfo{}
	ss := []string{}
	cs := []*pb.CalculationResult{}
	createTime := time.Now().Format(timeFormatString)
	for {
		c, err := stream.Recv()
		if err == io.EOF {
			for _, item := range as {
				if result, err := s.gdb.testCalculation(item.Expression); err != nil {
					return err
				} else {
					ss = append(ss, "insert into calc_cfg (description, expression, createTime, updatedTime, duration, status) values ('"+item.Description+"', '"+item.Expression+"' , '"+createTime+"', '"+createTime+"', '"+item.Duration+"', '"+item.Flag+"')")
					r, _ := json.Marshal(result.Result)
					cs = append(cs, &pb.CalculationResult{Result: string(r)})
				}
			}
			_ = s.gdb.updateItems(ss...)
			return stream.SendAndClose(&pb.CalculationResults{Results: cs})
		} else if err != nil {
			return err
		} else {
			as = append(as, addedCalcItemInfo{
				Expression:  c.GetExpression(),
				Flag:        c.GetFlag(),
				Duration:    c.GetDuration(),
				Description: c.GetDescription(),
			})
		}
	}
}

func (s *server) GetCalcItems(_ context.Context, r *pb.QueryCalcItemsInfo) (*pb.CalcItemsInfo, error) {
	if result, err := s.gdb.getCalculationItem(r.GetCondition()); err != nil {
		return &pb.CalcItemsInfo{}, err
	} else {
		r, _ := json.Marshal(result.Infos)
		return &pb.CalcItemsInfo{Infos: string(r)}, nil
	}
}

func (s *server) UpdateCalcItem(_ context.Context, r *pb.UpdatedCalcInfo) (*pb.CalculationResult, error) {
	if result, err := s.gdb.testCalculation(r.GetExpression()); err != nil {
		return nil, err
	} else {
		if _, err := s.gdb.updateCalculationItem(updatedCalcInfo{
			Id:          r.GetId(),
			Description: r.GetDescription(),
			Expression:  r.GetExpression(),
			Duration:    r.GetDuration(),
		}); err != nil {
			return &pb.CalculationResult{}, err
		} else {
			r, _ := json.Marshal(result.Result)
			return &pb.CalculationResult{
				Result: string(r),
			}, nil
		}
	}
}

func (s *server) StartCalcItem(_ context.Context, r *pb.CalcId) (*pb.TimeRows, error) {
	id := []string{}
	for _, item := range r.GetId() {
		id = append(id, "id = '"+item+"'")
	}
	if _, err := s.gdb.updateItem("update calc_cfg set status='true', updatedTime='" + time.Now().Format(timeFormatString) + "' where " + strings.Join(id, " or ")); err != nil {
		return &pb.TimeRows{}, err
	} else {
		return &pb.TimeRows{EffectedRows: 1}, nil
	}
}

func (s *server) StopCalcItem(_ context.Context, r *pb.CalcId) (*pb.TimeRows, error) {
	id := []string{}
	for _, item := range r.GetId() {
		id = append(id, "id = '"+item+"'")
	}
	if _, err := s.gdb.updateItem("update calc_cfg set status='false', updatedTime='" + time.Now().Format(timeFormatString) + "' where " + strings.Join(id, " or ")); err != nil {
		return &pb.TimeRows{}, err
	} else {
		return &pb.TimeRows{EffectedRows: 1}, nil
	}
}

func (s *server) DeleteCalcItem(_ context.Context, r *pb.CalcId) (*pb.TimeRows, error) {
	id := []string{}
	for _, item := range r.GetId() {
		id = append(id, "id = '"+item+"'")
	}
	if _, err := s.gdb.updateItem("delete from calc_cfg where " + strings.Join(id, " or ")); err != nil {
		return &pb.TimeRows{}, err
	} else {
		return &pb.TimeRows{EffectedRows: 1}, nil
	}
}

// use interceptor token authorization and userLogin

func (s *server) authInterceptor(c context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	if s.configs.Authorization {
		methods := strings.Split(info.FullMethod, "/")
		if md, ok := metadata.FromIncomingContext(c); !ok {
			return nil, status.Errorf(codes.Unauthenticated, "invalid token")
		} else {
			if methods[len(methods)-1] == "UserLogin" {
				r := req.(*pb.AuthInfo)
				if result, err := s.gdb.userLogin(authInfo{
					UserName: r.GetUserName(),
					PassWord: r.GetPassWord(),
				}); err != nil {
					return nil, status.Errorf(codes.Unauthenticated, "invalid token")
				} else {
					return &pb.UserToken{Token: result.Token}, nil
				}
			} else {
				var au string
				if d, ok := md["authorization"]; ok {
					au = d[0]
				} else {
					return nil, status.Errorf(codes.Unauthenticated, "invalid token")
				}
				if userName, token, ok := parseBasicAuth(au); !ok {
					return nil, status.Errorf(codes.Unauthenticated, "invalid token")
				} else {
					if r, err := s.gdb.query("select token from user_cfg where userName='" + userName + "'"); err != nil || len(r) == 0 {
						return nil, status.Errorf(codes.Unauthenticated, "invalid token")
					} else {
						if token != r[0]["token"] {
							return nil, status.Errorf(codes.Unauthenticated, "invalid token")
						} else {
							sub, obj, act := userName, methods[len(methods)-1], "POST" // replace gRCP with POST
							if ok, _ := s.gdb.e.Enforce(sub, obj, act); !ok {
								return nil, status.Errorf(codes.Unauthenticated, "invalid token")
							}
							return handler(c, req)
						}
					}
				}
			}
		}
	} else {
		return handler(c, req)
	}
}

func (s *server) authWithServerStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	if s.configs.Authorization {
		if !info.IsClientStream {
			return status.Errorf(codes.Unknown, "unknown service type")
		} else {
			if md, ok := metadata.FromIncomingContext(ss.Context()); !ok {
				return status.Errorf(codes.Unauthenticated, "invalid token")
			} else {
				var au string
				if d, ok := md["authorization"]; ok {
					au = d[0]
				} else {
					return status.Errorf(codes.Unauthenticated, "invalid token")
				}
				if userName, token, ok := parseBasicAuth(au); !ok {
					return status.Errorf(codes.Unauthenticated, "invalid token")
				} else {
					if r, err := s.gdb.query("select token from user_cfg where userName='" + userName + "'"); err != nil || len(r) == 0 {
						return status.Errorf(codes.Unauthenticated, "invalid token")
					} else {
						if token != r[0]["token"] {
							return status.Errorf(codes.Unauthenticated, "invalid token")
						} else {
							sub, obj, act := userName, "StreamWrite", "POST" // for stream gRPC, route permission is StreamWrite
							if ok, _ := s.gdb.e.Enforce(sub, obj, act); !ok {
								return status.Errorf(codes.Unauthenticated, "invalid token")
							}
							return handler(srv, ss)
						}
					}
				}
			}
		}
	} else {
		return handler(srv, ss)
	}
}

func (s *server) logInterceptor(c context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	if s.configs.LogWriting {
		if md, ok := metadata.FromIncomingContext(c); !ok {
			return nil, status.Errorf(codes.Unauthenticated, "invalid token")
		} else {
			remoteAddress := md.Get(":authority")[0] // address
			v := reflect.ValueOf(req)
			t := reflect.TypeOf(req)
			methodNames := []string{}
			r := map[string]interface{}{} // grpc request data
			for i := 0; i < t.NumMethod(); i++ {
				if strings.HasPrefix(t.Method(i).Name, "Get") {
					methodNames = append(methodNames, t.Method(i).Name) // get all grpc field functions
				}
			}
			for _, name := range methodNames {
				m := v.MethodByName(name) // get method
				p := make([]reflect.Value, m.Type().NumIn())
				result := m.Call(p)[0].Interface() // call function and get result
				r[strings.Replace(name, "Get", "", -1)] = result
			}
			rpcString, _ := json.Marshal(r)
			logMessage := logMessage{
				RequestUrl:    info.FullMethod,
				RequestMethod: "gRPC",
				UserAgent:     md.Get("user-agent")[0],
				RequestBody:   fmt.Sprintf("%s", rpcString),
				RemoteAddress: remoteAddress,
				Message:       "",
			}
			var userName string
			var au string
			if d, ok := md["authorization"]; ok {
				au = d[0]
				if name, _, ok := parseBasicAuth(au); ok {
					userName = name
				}
			}
			if v, err := handler(c, req); err != nil {
				logMessage.Message = strings.Replace(err.Error(), "'", `"`, -1)
				m, _ := json.Marshal(logMessage)
				_ = s.gdb.writeLog("Error", fmt.Sprintf("%s", m), userName)
				return v, err
			} else {
				// no errors
				if s.configs.Level == "Info" {
					m, _ := json.Marshal(logMessage)
					_ = s.gdb.writeLog("Info", fmt.Sprintf("%s", m), userName)
				}
				return v, nil
			}
		}
	} else {
		return handler(c, req)
	}
}

func (s *server) panicInterceptor(c context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic: %v\n", r)
		}
	}()
	return handler(c, req)
}

func (s *server) panicWithServerStreamInterceptor(srv interface{}, ss grpc.ServerStream, _ *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic: %v\n", r)
		}
	}()
	return handler(srv, ss)
}

func parseBasicAuth(auth string) (username, password string, ok bool) {
	const prefix = "Basic "
	// Case insensitive prefix match. See Issue 22736.
	if len(auth) < len(prefix) || !strings.EqualFold(auth[:len(prefix)], prefix) {
		return
	}
	c, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
	if err != nil {
		return
	}
	cs := string(c)
	s := strings.IndexByte(cs, ':')
	if s < 0 {
		return
	}
	return cs[:s], cs[s+1:], true
}
