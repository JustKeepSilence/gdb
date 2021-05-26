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

// AddGroups group handler
func (s *server) AddGroups(_ context.Context, r *pb.AddedGroupInfos) (*pb.Rows, error) {
	infos := []AddedGroupInfo{}
	for _, groupInfo := range r.GetGroupInfos() {
		infos = append(infos, AddedGroupInfo{GroupName: groupInfo.GroupName, ColumnNames: groupInfo.ColumnNames})
	}
	if result, err := s.gdb.AddGroups(infos...); err != nil {
		return nil, err
	} else {
		return &pb.Rows{EffectedRows: int32(result.EffectedRows)}, nil
	}
}

func (s *server) DeleteGroups(_ context.Context, r *pb.GroupNamesInfo) (*pb.Rows, error) {
	info := GroupNamesInfo{GroupNames: r.GroupNames}
	if result, err := s.gdb.DeleteGroups(info); err != nil {
		return nil, err
	} else {
		return &pb.Rows{EffectedRows: int32(result.EffectedRows)}, nil
	}
}

func (s *server) GetGroups(_ context.Context, _ *emptypb.Empty) (*pb.GroupNamesInfo, error) {
	if result, err := s.gdb.GetGroups(); err != nil {
		return nil, err
	} else {
		return &pb.GroupNamesInfo{GroupNames: result.GroupNames}, nil
	}
}

func (s *server) GetGroupProperty(_ context.Context, r *pb.QueryGroupPropertyInfo) (*pb.GroupPropertyInfo, error) {
	if result, err := s.gdb.GetGroupProperty(r.GetGroupName(), r.GetCondition()); err != nil {
		return nil, err
	} else {
		return &pb.GroupPropertyInfo{ItemCount: result.ItemCount, ItemColumnNames: result.ItemColumnNames}, nil
	}
}

func (s *server) UpdateGroupNames(_ context.Context, r *pb.UpdatedGroupNamesInfo) (*pb.Rows, error) {
	g := []UpdatedGroupNameInfo{}
	for _, info := range r.GetInfos() {
		g = append(g, UpdatedGroupNameInfo{NewGroupName: info.GetNewGroupName(), OldGroupName: info.GetOldGroupName()})
	}
	if result, err := s.gdb.UpdateGroupNames(g...); err != nil {
		return nil, err
	} else {
		return &pb.Rows{EffectedRows: int32(result.EffectedRows)}, nil
	}
}

func (s *server) UpdateGroupColumnNames(_ context.Context, r *pb.UpdatedGroupColumnNamesInfo) (*pb.Cols, error) {
	g := UpdatedGroupColumnNamesInfo{GroupName: r.GetGroupName(), OldColumnNames: r.GetOldColumnNames(), NewColumnNames: r.GetNewColumnNames()}
	if result, err := s.gdb.UpdateGroupColumnNames(g); err != nil {
		return nil, err
	} else {
		return &pb.Cols{EffectedCols: int32(result.EffectedCols)}, nil
	}
}

func (s *server) DeleteGroupColumns(_ context.Context, r *pb.DeletedGroupColumnNamesInfo) (*pb.Cols, error) {
	g := DeletedGroupColumnNamesInfo{GroupName: r.GetGroupName(), ColumnNames: r.GetColumnNames()}
	if result, err := s.gdb.DeleteGroupColumns(g); err != nil {
		return nil, err
	} else {
		return &pb.Cols{EffectedCols: int32(result.EffectedCols)}, nil
	}
}

func (s *server) AddGroupColumns(_ context.Context, r *pb.AddedGroupColumnsInfo) (*pb.Cols, error) {
	g := AddedGroupColumnsInfo{GroupName: r.GetGroupName(), ColumnNames: r.GetColumnNames(), DefaultValues: r.GetDefaultValues()}
	if result, err := s.gdb.AddGroupColumns(g); err != nil {
		return nil, err
	} else {
		return &pb.Cols{EffectedCols: int32(result.EffectedCols)}, nil
	}
}

func (s *server) CleanGroupItems(_ context.Context, r *pb.GroupNamesInfo) (*pb.Rows, error) {
	if result, err := s.gdb.CleanGroupItems(r.GroupNames...); err != nil {
		return nil, err
	} else {
		return &pb.Rows{EffectedRows: int32(result.EffectedRows)}, nil
	}
}

// item handler

func (s *server) AddItems(_ context.Context, r *pb.AddedItemsInfo) (*pb.Rows, error) {
	values := []map[string]string{}
	_ = json.Unmarshal([]byte(r.GetItemValues()), &values)
	g := AddedItemsInfo{
		GroupName:  r.GetGroupName(),
		ItemValues: values,
	}
	if result, err := s.gdb.AddItems(g); err != nil {
		return nil, err
	} else {
		return &pb.Rows{EffectedRows: int32(result.EffectedRows)}, nil
	}
}

func (s *server) DeleteItems(_ context.Context, r *pb.DeletedItemsInfo) (*pb.Rows, error) {
	g := DeletedItemsInfo{
		GroupName: r.GetGroupName(),
		Condition: r.GetCondition(),
	}
	if result, err := s.gdb.DeleteItems(g); err != nil {
		return nil, err
	} else {
		return &pb.Rows{EffectedRows: int32(result.EffectedRows)}, nil
	}
}

func (s *server) GetItems(_ context.Context, r *pb.ItemsInfo) (*pb.GdbItems, error) {
	g := ItemsInfo{
		GroupName:   r.GetGroupName(),
		Condition:   r.GetCondition(),
		ColumnNames: r.GetColumnNames(),
		StartRow:    int(r.GetStartRow()),
		RowCount:    int(r.GetRowCount()),
	}
	if result, err := s.gdb.GetItems(g); err != nil {
		return nil, err
	} else {
		v := []*pb.GdbItem{}
		for _, m := range result.ItemValues {
			v = append(v, &pb.GdbItem{Items: m})
		}
		return &pb.GdbItems{ItemValues: v}, nil
	}
}

func (s *server) GetItemsWithCount(_ context.Context, r *pb.ItemsInfo) (*pb.GdbItemsWithCount, error) {
	g := ItemsInfo{
		GroupName:   r.GetGroupName(),
		Condition:   r.GetCondition(),
		ColumnNames: r.GetColumnNames(),
		StartRow:    int(r.GetStartRow()),
		RowCount:    int(r.GetRowCount()),
	}
	if result, err := s.gdb.getItemsWithCount(g); err != nil {
		return nil, err
	} else {
		v := []*pb.GdbItem{}
		for _, m := range result.ItemValues {
			v = append(v, &pb.GdbItem{Items: m})
		}
		return &pb.GdbItemsWithCount{ItemValues: v, ItemCount: int32(result.ItemCount)}, nil
	}
}

func (s *server) UpdateItems(_ context.Context, r *pb.UpdatedItemsInfo) (*pb.Rows, error) {
	g := UpdatedItemsInfo{
		GroupName: r.GetGroupName(),
		Condition: r.GetCondition(),
		Clause:    r.GetClause(),
	}
	if result, err := s.gdb.UpdateItems(g); err != nil {
		return nil, err
	} else {
		return &pb.Rows{EffectedRows: int32(result.EffectedRows)}, nil
	}
}

func (s *server) CheckItems(_ context.Context, r *pb.CheckItemsInfo) (*emptypb.Empty, error) {
	if err := s.gdb.CheckItems(r.GetGroupName(), r.GetItemNames()...); err != nil {
		return &emptypb.Empty{}, err
	} else {
		return &emptypb.Empty{}, nil
	}
}

// data handler

func (s *server) BatchWrite(_ context.Context, r *pb.BatchWriteString) (*pb.Rows, error) {
	v := []ItemValue{}
	if err := json.Unmarshal([]byte(r.GetItemValues()), &v); err != nil {
		return &pb.Rows{}, err
	} else {
		if result, err := s.gdb.BatchWrite(v...); err != nil {
			return nil, err
		} else {
			return &pb.Rows{EffectedRows: int32(result.EffectedRows)}, nil
		}
	}
}

// BatchWriteWithStream write data with client stream
func (s *server) BatchWriteWithStream(stream pb.Data_BatchWriteWithStreamServer) error {
	bs := []batchWriteString{}
	for {
		b, err := stream.Recv()
		if err == io.EOF {
			eg := errgroup.Group{}
			for _, ss := range bs {
				writingString := ss
				eg.Go(func() error {
					if _, err := s.gdb.BatchWrite(writingString.ItemValues...); err != nil {
						return fmt.Errorf("writing error :" + err.Error())
					} else {
						return nil
					}
				})
			}
			if err := eg.Wait(); err != nil {
				return err
			} else {
				return stream.SendAndClose(&pb.Rows{EffectedRows: int32(len(bs))})
			}
		} else if err != nil {
			return err
		} else {
			v := []ItemValue{}
			if err := json.Unmarshal([]byte(b.GetItemValues()), &v); err != nil {
				return err
			} else {
				bs = append(bs, batchWriteString{
					ItemValues: v,
				})
			}
		}
	}
}

func (s *server) BatchWriteHistoricalData(_ context.Context, r *pb.BatchWriteHistoricalString) (*emptypb.Empty, error) {
	values := []HistoricalItemValue{}
	if err := json.Unmarshal([]byte(r.GetHistoricalItemValues()), &values); err != nil {
		return &emptypb.Empty{}, err
	} else {
		if err := s.gdb.BatchWriteHistoricalData(values...); err != nil {
			return &emptypb.Empty{}, err
		} else {
			return &emptypb.Empty{}, nil
		}
	}
}

func (s *server) BatchWriteHistoricalDataWithStream(stream pb.Data_BatchWriteHistoricalDataWithStreamServer) error {
	bs := []batchWriteHistoricalString{}
	for {
		b, err := stream.Recv()
		if err == io.EOF {
			eg := errgroup.Group{}
			for _, ss := range bs {
				writingString := ss
				eg.Go(func() error {
					if err := s.gdb.BatchWriteHistoricalData(writingString.HistoricalItemValues...); err != nil {
						return fmt.Errorf("writing error :" + err.Error())
					} else {
						return nil
					}
				})
			}
			if err := eg.Wait(); err != nil {
				return err
			} else {
				return stream.SendAndClose(&emptypb.Empty{})
			}
		} else if err != nil {
			return err
		} else {
			v := []HistoricalItemValue{}
			if err := json.Unmarshal([]byte(b.GetHistoricalItemValues()), &v); err != nil {
				return err
			} else {
				bs = append(bs, batchWriteHistoricalString{HistoricalItemValues: v})
			}
		}
	}
}

func (s *server) GetRealTimeData(_ context.Context, r *pb.QueryRealTimeDataString) (*pb.GdbRealTimeData, error) {
	if result, err := s.gdb.GetRealTimeData(r.GetGroupNames(), r.ItemNames...); err != nil {
		return nil, err
	} else {
		v, _ := json.Marshal(result)
		return &pb.GdbRealTimeData{RealTimeData: fmt.Sprintf("%s", v)}, nil
	}
}

func (s *server) GetHistoricalData(_ context.Context, r *pb.QueryHistoricalDataString) (*pb.GdbHistoricalData, error) {
	if result, err := s.gdb.GetHistoricalData(r.GetGroupNames(), r.GetItemNames(), convertInt32ToInt(r.GetStartTimes()...), convertInt32ToInt(r.GetEndTimes()...), convertInt32ToInt(r.GetIntervals()...)); err != nil {
		return nil, err
	} else {
		v, _ := json.Marshal(result)
		return &pb.GdbHistoricalData{HistoricalData: fmt.Sprintf("%s", v)}, nil
	}
}

func (s *server) GetHistoricalDataWithStamp(_ context.Context, r *pb.QueryHistoricalDataWithTimeStampString) (*pb.GdbHistoricalData, error) {
	t := [][]int{}
	for _, s := range r.GetTimeStamps() {
		t = append(t, convertInt32ToInt(s.GetTimeStamp()...))
	}
	if result, err := s.gdb.GetHistoricalDataWithStamp(r.GetGroupNames(), r.GetItemNames(), t...); err != nil {
		return nil, err
	} else {
		v, _ := json.Marshal(result)
		return &pb.GdbHistoricalData{HistoricalData: fmt.Sprintf("%s", v)}, nil
	}
}

func (s *server) GetHistoricalDataWithCondition(_ context.Context, r *pb.QueryHistoricalDataWithConditionString) (*pb.GdbHistoricalData, error) {
	dz := []DeadZone{}
	for _, zone := range r.GetDeadZones() {
		dz = append(dz, DeadZone{
			ItemName:      zone.ItemName,
			DeadZoneCount: int(zone.DeadZoneCount),
		})
	}
	if result, err := s.gdb.GetHistoricalDataWithCondition(r.GetGroupNames(), r.GetItemNames(), convertInt32ToInt(r.GetStartTimes()...),
		convertInt32ToInt(r.GetEndTimes()...), convertInt32ToInt(r.GetIntervals()...), r.GetFilterCondition(), dz...); err != nil {
		return &pb.GdbHistoricalData{}, nil
	} else {
		v, _ := json.Marshal(result)
		return &pb.GdbHistoricalData{HistoricalData: string(v)}, nil
	}
}

func (s *server) GetRawData(_ context.Context, r *pb.QueryRealTimeDataString) (*pb.GdbHistoricalData, error) {
	if result, err := s.gdb.GetRawHistoricalData(r.GetGroupNames(), r.GetItemNames()...); err != nil {
		return &pb.GdbHistoricalData{}, nil
	} else {
		v, _ := json.Marshal(result)
		return &pb.GdbHistoricalData{HistoricalData: string(v)}, nil
	}
}

func (s *server) GetDbInfo(_ context.Context, _ *emptypb.Empty) (*pb.GdbInfoData, error) {
	if result, err := s.gdb.getDbInfo(); err != nil {
		return nil, err
	} else {
		v, _ := json.Marshal(result)
		return &pb.GdbInfoData{Info: fmt.Sprintf("%s", v)}, nil
	}
}

func (s *server) GetDbInfoHistory(_ context.Context, r *pb.QuerySpeedHistoryDataString) (*pb.GdbHistoricalData, error) {
	if r, err := s.gdb.getDbInfoHistory(r.GetItemName(), convertInt32ToInt(r.GetStartTimes()...), convertInt32ToInt(r.GetEndTimes()...), convertInt32ToInt(r.GetIntervals()...)); err != nil {
		return &pb.GdbHistoricalData{}, nil
	} else {
		result, _ := json.Marshal(r)
		return &pb.GdbHistoricalData{HistoricalData: fmt.Sprintf("%s", result)}, nil
	}
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
	if result, err := query(s.gdb.ItemDbPath, "select * from user_cfg"); err != nil {
		return &pb.UserInfos{}, err
	} else {
		r, _ := json.Marshal(result)
		return &pb.UserInfos{UserInfos: fmt.Sprintf("%s", r)}, nil
	}
}

func (s *server) AddUsers(_ context.Context, r *pb.AddUserInfo) (*pb.Rows, error) {
	g := addedUserInfo{
		Name:     r.GetName(),
		Role:     r.GetRole(),
		PassWord: r.GetPassWord(),
	}
	if result, err := s.gdb.addUsers(g); err != nil {
		return &pb.Rows{}, err
	} else {
		return &pb.Rows{EffectedRows: int32(result.EffectedRows)}, nil
	}
}

func (s *server) DeleteUsers(_ context.Context, r *pb.UserName) (*pb.Rows, error) {
	g := userName{Name: r.GetName()}
	if result, err := s.gdb.deleteUsers(g); err != nil {
		return &pb.Rows{}, err
	} else {
		return &pb.Rows{EffectedRows: int32(result.EffectedRows)}, nil
	}
}

func (s *server) UpdateUsers(_ context.Context, r *pb.UpdatedUserInfo) (*pb.Rows, error) {
	g := updatedUserInfo{
		UserName:    r.GetUserName(),
		NewUserName: r.GetNewUserName(),
		NewPassWord: r.GetNewPassWord(),
		NewRole:     r.GetNewRole(),
	}
	if result, err := s.gdb.updateUsers(g); err != nil {
		return &pb.Rows{}, err
	} else {
		return &pb.Rows{EffectedRows: int32(result.EffectedRows)}, nil
	}
}

func (s *server) GetLogs(_ context.Context, r *pb.QueryLogsInfo) (*pb.LogsInfo, error) {
	g := queryLogsInfo{
		Level:     r.GetLevel(),
		StartTime: r.GetStartTime(),
		EndTime:   r.GetEndTime(),
		StartRow:  int(r.GetStartRow()),
		RowCount:  int(r.GetRowCount()),
		Name:      r.GetName(),
	}
	if result, err := s.gdb.getLogs(g); err != nil {
		return nil, err
	} else {
		r, _ := json.Marshal(result.Infos)
		return &pb.LogsInfo{Infos: string(r), Count: int32(result.Count)}, nil
	}
}

func (s *server) DeleteLogs(_ context.Context, r *pb.DeletedLogInfo) (*pb.Rows, error) {
	g := deletedLogInfo{
		Id:                r.GetId(),
		StartTime:         r.GetStartTime(),
		EndTime:           r.GetEndTime(),
		UserNameCondition: r.GetUserNameCondition(),
	}
	if result, err := s.gdb.deleteLogs(g); err != nil {
		return &pb.Rows{}, err
	} else {
		return &pb.Rows{EffectedRows: int32(result.EffectedRows)}, nil
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

func (s *server) AddItemsByExcel(_ context.Context, r *pb.FileInfo) (*pb.Rows, error) {
	if result, err := s.gdb.addItemsByExcel(r.GetGroupName(), "./uploadFiles/"+r.GetFileName()); err != nil {
		return &pb.Rows{}, err
	} else {
		return &pb.Rows{EffectedRows: int32(result.EffectedRows)}, nil
	}
}

func (s *server) ImportHistoryByExcel(_ context.Context, r *pb.HistoryFileInfo) (*emptypb.Empty, error) {
	if err := s.gdb.importHistoryByExcel("./uploadFiles/"+r.GetFileName(), r.GetGroupName(), r.GetItemNames(), r.GetSheetNames()...); err != nil {
		return &emptypb.Empty{}, err
	} else {
		return &emptypb.Empty{}, nil
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

func (s *server) TestCalcItem(_ context.Context, r *pb.TestCalcItemInfo) (*pb.TestResult, error) {
	if result, err := s.gdb.testCalculation(r.GetExpression()); err != nil {
		return &pb.TestResult{}, err
	} else {
		r, _ := json.Marshal(result.Result)
		return &pb.TestResult{Result: string(r)}, nil
	}
}

func (s *server) AddCalcItem(_ context.Context, r *pb.AddedCalcItemInfo) (*pb.CalculationResult, error) {
	if result, err := s.gdb.testCalculation(r.GetExpression()); err != nil {
		return nil, err
	} else {
		createTime := time.Now().Format(timeFormatString)
		if _, err := updateItem(s.gdb.ItemDbPath, "insert into calc_cfg (description, expression, createTime, updatedTime, duration, status) values ('"+r.GetDescription()+"', '"+r.GetExpression()+"' , '"+createTime+"', '"+createTime+"', '"+r.GetDuration()+"', '"+r.GetFlag()+"')"); err != nil {
			return nil, err
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
			_ = updateItems(s.gdb.ItemDbPath, ss...)
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
		return nil, err
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
			return nil, err
		} else {
			return &pb.CalculationResult{
				Result: result.Result.(string),
			}, nil
		}
	}
}

func (s *server) StartCalcItem(_ context.Context, r *pb.CalcId) (*pb.Rows, error) {
	id := []string{}
	for _, item := range r.GetId() {
		id = append(id, "id = '"+item+"'")
	}
	if _, err := updateItem(s.gdb.ItemDbPath, "update calc_cfg set status='true', updatedTime='"+time.Now().Format(timeFormatString)+"' where "+strings.Join(id, " or ")); err != nil {
		return nil, err
	} else {
		return &pb.Rows{EffectedRows: 1}, nil
	}
}

func (s *server) StopCalcItem(_ context.Context, r *pb.CalcId) (*pb.Rows, error) {
	id := []string{}
	for _, item := range r.GetId() {
		id = append(id, "id = '"+item+"'")
	}
	if _, err := updateItem(s.gdb.ItemDbPath, "update calc_cfg set status='false', updatedTime='"+time.Now().Format(timeFormatString)+"' where "+strings.Join(id, " or ")); err != nil {
		return nil, err
	} else {
		return &pb.Rows{EffectedRows: 1}, nil
	}
}

func (s *server) DeleteCalcItem(_ context.Context, r *pb.CalcId) (*pb.Rows, error) {
	id := []string{}
	for _, item := range r.GetId() {
		id = append(id, "id = '"+item+"'")
	}
	if _, err := updateItem(s.gdb.ItemDbPath, "delete from calc_cfg where "+strings.Join(id, " or ")); err != nil {
		return nil, err
	} else {
		return &pb.Rows{EffectedRows: 1}, nil
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
					if r, err := query(s.gdb.ItemDbPath, "select token from user_cfg where userName='"+userName+"'"); err != nil || len(r) == 0 {
						return nil, status.Errorf(codes.Unauthenticated, "invalid token")
					} else {
						if token != r[0]["token"] {
							return nil, status.Errorf(codes.Unauthenticated, "invalid token")
						} else {
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
					if r, err := query(s.gdb.ItemDbPath, "select token from user_cfg where userName='"+userName+"'"); err != nil || len(r) == 0 {
						return status.Errorf(codes.Unauthenticated, "invalid token")
					} else {
						if token != r[0]["token"] {
							return status.Errorf(codes.Unauthenticated, "invalid token")
						} else {
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

func convertInt32ToInt(items ...int32) []int {
	var r []int
	for _, item := range items {
		r = append(r, int(item))
	}
	return r
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
