/*
creatTime: 2021/2/8
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.15.3
*/

package db

// gRPC, for details of gRPC, you can see: https://github.com/grpc/grpc-go

import (
	"context"
	"fmt"
	pb "github.com/JustKeepSilence/gdb/model"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
	"time"
)

type server struct {
	pb.UnimplementedGroupServer
	pb.UnimplementedItemServer
	pb.UnimplementedDataServer
	gdb *Gdb
}

// group handler
func (s *server) AddGroups(_ context.Context, r *pb.AddedGroupInfos) (*pb.Rows, error) {
	infos := []AddedGroupInfo{}
	for _, groupInfo := range r.GetGroupInfos() {
		infos = append(infos, AddedGroupInfo{GroupName: groupInfo.GroupName, ColumnNames: groupInfo.ColumnNames})
	}
	if r, err := s.gdb.AddGroups(infos...); err != nil {
		return nil, err
	} else {
		return &pb.Rows{EffectedRows: int32(r.EffectedRows)}, nil
	}
}

func (s *server) DeleteGroups(_ context.Context, r *pb.GroupNamesInfo) (*pb.Rows, error) {
	info := GroupNamesInfo{GroupNames: r.GroupNames}
	if r, err := s.gdb.DeleteGroups(info); err != nil {
		return nil, err
	} else {
		return &pb.Rows{EffectedRows: int32(r.EffectedRows)}, nil
	}
}

func (s *server) GetGroups(_ context.Context, _ *emptypb.Empty) (*pb.GroupNamesInfo, error) {
	if groupInfos, err := s.gdb.GetGroups(); err != nil {
		return nil, err
	} else {
		return &pb.GroupNamesInfo{GroupNames: groupInfos.GroupNames}, nil
	}
}

func (s *server) GetGroupProperty(_ context.Context, r *pb.QueryGroupPropertyInfo) (*pb.GroupPropertyInfo, error) {
	if r, err := s.gdb.GetGroupProperty(r.GetGroupName(), r.GetCondition()); err != nil {
		return nil, err
	} else {
		return &pb.GroupPropertyInfo{ItemCount: r.ItemCount, ItemColumnNames: r.ItemColumnNames}, nil
	}
}

func (s *server) UpdateGroupNames(_ context.Context, r *pb.UpdatedGroupNamesInfo) (*pb.Rows, error) {
	g := []UpdatedGroupNameInfo{}
	for _, info := range r.GetInfos() {
		g = append(g, UpdatedGroupNameInfo{NewGroupName: info.GetNewGroupName(), OldGroupName: info.GetOldGroupName()})
	}
	if r, err := s.gdb.UpdateGroupNames(g...); err != nil {
		return nil, err
	} else {
		return &pb.Rows{EffectedRows: int32(r.EffectedRows)}, nil
	}
}

func (s *server) UpdateGroupColumnNames(_ context.Context, r *pb.UpdatedGroupColumnNamesInfo) (*pb.Cols, error) {
	g := UpdatedGroupColumnNamesInfo{GroupName: r.GetGroupName(), OldColumnNames: r.GetOldColumnNames(), NewColumnNames: r.GetNewColumnNames()}
	if r, err := s.gdb.UpdateGroupColumnNames(g); err != nil {
		return nil, err
	} else {
		return &pb.Cols{EffectedCols: int32(r.EffectedCols)}, nil
	}
}

func (s *server) DeleteGroupColumns(_ context.Context, r *pb.DeletedGroupColumnNamesInfo) (*pb.Cols, error) {
	g := DeletedGroupColumnNamesInfo{GroupName: r.GetGroupName(), ColumnNames: r.GetColumnNames()}
	if r, err := s.gdb.DeleteGroupColumns(g); err != nil {
		return nil, err
	} else {
		return &pb.Cols{EffectedCols: int32(r.EffectedCols)}, nil
	}
}

func (s *server) AddGroupColumns(_ context.Context, r *pb.AddedGroupColumnsInfo) (*pb.Cols, error) {
	g := AddedGroupColumnsInfo{GroupName: r.GetGroupName(), ColumnNames: r.GetColumnNames(), DefaultValues: r.GetDefaultValues()}
	if r, err := s.gdb.AddGroupColumns(g); err != nil {
		return nil, err
	} else {
		return &pb.Cols{EffectedCols: int32(r.EffectedCols)}, nil
	}
}

// item handler

func (s *server) AddItems(_ context.Context, r *pb.AddedItemsInfo) (*pb.Rows, error) {
	values := []map[string]string{}
	for _, value := range r.GetValues() {
		values = append(values, value.GetItems())
	}
	g := AddedItemsInfo{
		GroupName: r.GetGroupName(),
		GdbItems:  GdbItems{ItemValues: values},
	}
	if r, err := s.gdb.AddItems(g); err != nil {
		return nil, err
	} else {
		return &pb.Rows{EffectedRows: int32(r.EffectedRows)}, nil
	}
}

func (s *server) DeleteItems(_ context.Context, r *pb.DeletedItemsInfo) (*pb.Rows, error) {
	g := DeletedItemsInfo{
		GroupName: r.GetGroupName(),
		Condition: r.GetCondition(),
	}
	if r, err := s.gdb.DeleteItems(g); err != nil {
		return nil, err
	} else {
		return &pb.Rows{EffectedRows: int32(r.EffectedRows)}, nil
	}
}

func (s *server) GetItems(_ context.Context, r *pb.ItemsInfo) (*pb.GdbItems, error) {
	g := ItemsInfo{
		ItemsInfoWithoutRow: ItemsInfoWithoutRow{GroupName: r.GetInfos().GroupName, Condition: r.GetInfos().Condition, Clause: r.GetInfos().Clause},
		ColumnNames:         r.GetColumnNames(),
		StartRow:            int(r.GetStartRow()),
		RowCount:            int(r.GetRowCount()),
	}
	if r, err := s.gdb.GetItems(g); err != nil {
		return nil, err
	} else {
		v := []*pb.GdbItem{}
		for _, m := range r.ItemValues {
			v = append(v, &pb.GdbItem{Items: m})
		}
		return &pb.GdbItems{ItemValues: v}, nil
	}
}

func (s *server) GetItemsWithCount(_ context.Context, r *pb.ItemsInfo) (*pb.GdbItemsWithCount, error) {
	g := ItemsInfo{
		ItemsInfoWithoutRow: ItemsInfoWithoutRow{GroupName: r.GetInfos().GroupName, Condition: r.GetInfos().Condition, Clause: r.GetInfos().Clause},
		ColumnNames:         r.GetColumnNames(),
		StartRow:            int(r.GetStartRow()),
		RowCount:            int(r.GetRowCount()),
	}
	if r, err := s.gdb.GetItemsWithCount(g); err != nil {
		return nil, err
	} else {
		v := []*pb.GdbItem{}
		for _, m := range r.ItemValues {
			v = append(v, &pb.GdbItem{Items: m})
		}
		return &pb.GdbItemsWithCount{ItemValues: v, ItemCount: int32(r.ItemCount)}, nil
	}
}

func (s *server) UpdateItems(_ context.Context, r *pb.ItemsInfoWithoutRow) (*pb.Rows, error) {
	g := ItemsInfoWithoutRow{
		GroupName: r.GetGroupName(),
		Condition: r.GetCondition(),
		Clause:    r.GetClause(),
	}
	if r, err := s.gdb.UpdateItems(g); err != nil {
		return nil, err
	} else {
		return &pb.Rows{EffectedRows: int32(r.EffectedRows)}, nil
	}
}

// data handler

func (s *server) BatchWrite(_ context.Context, r *pb.BatchWriteString) (*pb.Rows, error) {
	v := []ItemValue{}
	for _, itemValue := range r.GetItemValues() {
		v = append(v, ItemValue{
			ItemName:  itemValue.GetItemName(),
			Value:     itemValue.GetValue(),
			TimeStamp: itemValue.GetTimeStamp(),
		})
	}
	g := BatchWriteString{
		GroupName:     r.GetGroupName(),
		ItemValues:    v,
		WithTimeStamp: r.GetWithTimeStamp(),
	}
	if r, err := s.gdb.BatchWrite(g); err != nil {
		return nil, err
	} else {
		return &pb.Rows{EffectedRows: int32(r.EffectedRows)}, nil
	}
}

// write data with client stream
func (s *server) BatchWriteWithStream(_ context.Context) {

}

func (s *server) GetRealTimeData(_ context.Context, r *pb.QueryRealTimeDataString) (*pb.GdbRealTimeData, error) {
	if r, err := s.gdb.GetRealTimeData(r.ItemNames...); err != nil {
		return nil, err
	} else {
		v, _ := Json.Marshal(r.RealTimeData)
		return &pb.GdbRealTimeData{RealTimeData: fmt.Sprintf("%s", v)}, nil
	}
}

func (s *server) GetHistoricalData(_ context.Context, r *pb.QueryHistoricalDataString) (*pb.GdbHistoricalData, error) {
	if r, err := s.gdb.GetHistoricalData(r.GetItemNames(), r.GetStartTimes(), r.GetEndTimes(), r.GetIntervals()); err != nil {
		return nil, err
	} else {
		v, _ := Json.Marshal(r.HistoricalData)
		return &pb.GdbHistoricalData{HistoricalData: fmt.Sprintf("%s", v)}, nil
	}
}

func (s *server) GetHistoricalDataWithStamp(_ context.Context, r *pb.QueryHistoricalDataWithTimeStampString) (*pb.GdbHistoricalData, error) {
	t := [][]int32{}
	for _, s := range r.GetTimeStamps() {
		t = append(t, s.GetTimeStamp())
	}
	if r, err := s.gdb.GetHistoricalDataWithStamp(r.GetItemNames(), t...); err != nil {
		return nil, err
	} else {
		v, _ := Json.Marshal(r.HistoricalData)
		return &pb.GdbHistoricalData{HistoricalData: fmt.Sprintf("%s", v)}, nil
	}
}

func (s *server) GetDbInfo(_ context.Context, _ *emptypb.Empty) (*pb.GdbInfoData, error) {
	if r, err := s.gdb.getDbInfo(); err != nil {
		return nil, err
	} else {
		v, _ := Json.Marshal(r.Info)
		return &pb.GdbInfoData{Info: fmt.Sprintf("%s", v)}, nil
	}
}

func InitialDbRPCServer(port, dbPath, itemDbPath string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Print(err)
	}
	s := grpc.NewServer()
	if g, err := NewGdb(dbPath, itemDbPath); err != nil {
		log.Println("fail in initialing gdb: " + err.Error())
		time.Sleep(time.Second * 60)
	} else {
		pb.RegisterGroupServer(s, &server{gdb: g})
		fmt.Println("launch gRPC successfully!")
		if err := s.Serve(lis); err != nil {
			log.Print(err)
		}
	}
}
