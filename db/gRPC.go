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
	_ = Json.Unmarshal([]byte(r.GetItemValues()), &values)
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
		Clause:      r.GetClause(),
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
		Clause:      r.GetClause(),
		ColumnNames: r.GetColumnNames(),
		StartRow:    int(r.GetStartRow()),
		RowCount:    int(r.GetRowCount()),
	}
	if result, err := s.gdb.GetItemsWithCount(g); err != nil {
		return nil, err
	} else {
		v := []*pb.GdbItem{}
		for _, m := range result.ItemValues {
			v = append(v, &pb.GdbItem{Items: m})
		}
		return &pb.GdbItemsWithCount{ItemValues: v, ItemCount: int32(result.ItemCount)}, nil
	}
}

func (s *server) UpdateItems(_ context.Context, r *pb.ItemsInfo) (*pb.Rows, error) {
	g := ItemsInfo{
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
	for _, itemValue := range r.GetItemValues() {
		v = append(v, ItemValue{
			ItemName: itemValue.GetItemName(),
			Value:    itemValue.GetValue(),
		})
	}
	g := BatchWriteString{
		ItemValues: v,
	}
	if result, err := s.gdb.BatchWrite(g); err != nil {
		return nil, err
	} else {
		return &pb.Rows{EffectedRows: int32(result.EffectedRows)}, nil
	}
}

// BatchWriteWithStream write data with client stream
func (s *server) BatchWriteWithStream(stream pb.Data_BatchWriteWithStreamServer) error {
	bs := []BatchWriteString{}
	for {
		b, err := stream.Recv()
		if err == io.EOF {
			eg := errgroup.Group{}
			for _, ss := range bs {
				writingString := ss
				eg.Go(func() error {
					if _, err := s.gdb.BatchWrite(writingString); err != nil {
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
			for _, itemValue := range b.GetItemValues() {
				v = append(v, ItemValue{
					ItemName: itemValue.GetItemName(),
					Value:    itemValue.GetValue(),
				})
			}
			bs = append(bs, BatchWriteString{
				ItemValues: v,
			})
		}
	}
}

func (s *server) BatchWriteHistoricalData(_ context.Context, r *pb.BatchWriteHistoricalString) (*emptypb.Empty, error) {
	values := []HistoricalItemValue{}
	v := r.GetHistoricalItemValues()
	for i := 0; i < len(v); i++ {
		values = append(values, HistoricalItemValue{
			ItemName:   v[i].ItemName,
			Values:     v[i].Values,
			TimeStamps: v[i].TimeStamps,
		})
	}
	b := BatchWriteHistoricalString{values}
	if err := s.gdb.BatchWriteHistoricalData(b); err != nil {
		return &emptypb.Empty{}, err
	} else {
		return &emptypb.Empty{}, nil
	}
}

func (s *server) BatchWriteHistoricalDataWithStream(stream pb.Data_BatchWriteHistoricalDataWithStreamServer) error {
	bs := []BatchWriteHistoricalString{}
	for {
		b, err := stream.Recv()
		if err == io.EOF {
			eg := errgroup.Group{}
			for _, ss := range bs {
				writingString := ss
				eg.Go(func() error {
					if err := s.gdb.BatchWriteHistoricalData(writingString); err != nil {
						return fmt.Errorf("writing error :" + err.Error())
					} else {
						return nil
					}
				})
			}
			if err := eg.Wait(); err != nil {
				return err
			} else {
				return stream.SendAndClose(nil)
			}
		} else if err != nil {
			return err
		} else {
			v := []HistoricalItemValue{}
			for _, itemValue := range b.GetHistoricalItemValues() {
				v = append(v, HistoricalItemValue{
					ItemName:   itemValue.ItemName,
					Values:     itemValue.Values,
					TimeStamps: itemValue.TimeStamps,
				})
			}
			bs = append(bs, BatchWriteHistoricalString{HistoricalItemValues: v})
		}
	}
}

func (s *server) GetRealTimeData(_ context.Context, r *pb.QueryRealTimeDataString) (*pb.GdbRealTimeData, error) {
	if result, err := s.gdb.GetRealTimeData(r.ItemNames...); err != nil {
		return nil, err
	} else {
		v, _ := Json.Marshal(result.RealTimeData)
		return &pb.GdbRealTimeData{RealTimeData: fmt.Sprintf("%s", v)}, nil
	}
}

func (s *server) GetHistoricalData(_ context.Context, r *pb.QueryHistoricalDataString) (*pb.GdbHistoricalData, error) {
	if result, err := s.gdb.GetHistoricalData(r.GetItemNames(), convertInt32ToInt(r.GetStartTimes()...), convertInt32ToInt(r.GetEndTimes()...), convertInt32ToInt(r.GetIntervals()...)); err != nil {
		return nil, err
	} else {
		v, _ := Json.Marshal(result.HistoricalData)
		return &pb.GdbHistoricalData{HistoricalData: fmt.Sprintf("%s", v)}, nil
	}
}

func (s *server) GetHistoricalDataWithStamp(_ context.Context, r *pb.QueryHistoricalDataWithTimeStampString) (*pb.GdbHistoricalData, error) {
	t := [][]int{}
	for _, s := range r.GetTimeStamps() {
		t = append(t, convertInt32ToInt(s.GetTimeStamp()...))
	}
	if result, err := s.gdb.GetHistoricalDataWithStamp(r.GetItemNames(), t...); err != nil {
		return nil, err
	} else {
		v, _ := Json.Marshal(result.HistoricalData)
		return &pb.GdbHistoricalData{HistoricalData: fmt.Sprintf("%s", v)}, nil
	}
}

func (s *server) GetDbInfo(_ context.Context, _ *emptypb.Empty) (*pb.GdbInfoData, error) {
	if result, err := s.gdb.getDbInfo(); err != nil {
		return nil, err
	} else {
		v, _ := Json.Marshal(result.Info)
		return &pb.GdbInfoData{Info: fmt.Sprintf("%s", v)}, nil
	}
}

func (s *server) GetDbInfoHistory(_ context.Context, r *pb.QuerySpeedHistoryDataString) (*pb.GdbHistoricalData, error) {
	if r, err := s.gdb.getDbInfoHistory(r.GetItemName(), convertInt32ToInt(r.GetStartTimes()...), convertInt32ToInt(r.GetEndTimes()...), int(r.GetInterval())); err != nil {
		return &pb.GdbHistoricalData{}, nil
	} else {
		result, _ := Json.Marshal(r)
		return &pb.GdbHistoricalData{HistoricalData: fmt.Sprintf("%s", result)}, nil
	}
}

// page handler

func (s *server) UserLogin(c context.Context, r *pb.AuthInfo) (*pb.UserToken, error) {
	g := authInfo{
		UserName: r.GetUserName(),
		PassWord: r.GetPassWord(),
	}
	md, _ := metadata.FromIncomingContext(c)
	if token, err := s.gdb.userLogin(g, md.Get("user-agent")[0]); err != nil {
		return &pb.UserToken{}, nil
	} else {
		return &pb.UserToken{Token: token.Token}, nil
	}
}

func (s *server) UserLogOut(c context.Context, _ *pb.UserName) (*emptypb.Empty, error) {
	md, _ := metadata.FromIncomingContext(c)
	au := md.Get("authorization")[0]
	userAgent := md.Get("user-agent")[0]
	userName, token, _ := parseBasicAuth(au)
	if _, err := s.gdb.userLogout(userName, token, userAgent); err != nil {
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
			UserName: &pb.UserName{Name: result.Name},
			Role:     result.Role,
		}, nil
	}
}

func (s *server) GetUsers(_ context.Context, _ *emptypb.Empty) (*pb.UserInfos, error) {
	if result, err := query(s.gdb.ItemDbPath, "select * from user_cfg"); err != nil {
		return &pb.UserInfos{}, err
	} else {
		r, _ := Json.Marshal(result)
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
	g := UserName{Name: r.GetName()}
	if result, err := s.gdb.deleteUsers(g); err != nil {
		return &pb.Rows{}, err
	} else {
		return &pb.Rows{EffectedRows: int32(result.EffectedRows)}, nil
	}
}

func (s *server) UpdateUsers(_ context.Context, r *pb.UpdatedUserInfo) (*pb.Rows, error) {
	g := updatedUserInfo{
		Id:          int(r.GetId()),
		OldUserName: r.GetOldUserName(),
		NewUserName: r.GetNewUserName(),
		Role:        r.GetRole(),
	}
	if result, err := s.gdb.updateUsers(g); err != nil {
		return &pb.Rows{}, err
	} else {
		return &pb.Rows{EffectedRows: int32(result.EffectedRows)}, nil
	}
}

func (s *server) GetLogs(_ context.Context, r *pb.QueryLogsInfo) (*pb.LogsInfo, error) {
	if result, err := s.gdb.getLogs(r.LogType, r.Condition, r.StartTime, r.EndTime); err != nil {
		return nil, err
	} else {
		logs := []*pb.LogInfo{}
		for _, item := range result.Infos {
			logs = append(logs, &pb.LogInfo{Info: item})
		}
		return &pb.LogsInfo{Infos: logs}, nil
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
	if result, err := s.gdb.AddItemsByExcel(r.GetGroupName(), "./uploadFiles/"+r.GetFileName()); err != nil {
		return &pb.Rows{}, err
	} else {
		return &pb.Rows{EffectedRows: int32(result.EffectedRows)}, nil
	}
}

func (s *server) ImportHistoryByExcel(_ context.Context, r *pb.HistoryFileInfo) (*emptypb.Empty, error) {
	if err := s.gdb.ImportHistoryByExcel("./uploadFiles/"+r.GetFileName(), r.GetItemNames(), r.GetSheetNames()...); err != nil {
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

func (s *server) AddCalcItem(_ context.Context, r *pb.AddedCalcItemInfo) (*pb.CalculationResult, error) {
	if result, err := s.gdb.testCalculation(r.GetExpression()); err != nil {
		return nil, err
	} else {
		createTime := time.Now().Format(timeFormatString)
		if _, err := updateItem(s.gdb.ItemDbPath, "insert into calc_cfg (description, expression, createTime, updatedTime, duration, status) values ('"+r.GetDescription()+"', '"+r.GetExpression()+"' , '"+createTime+"', '"+createTime+"', '"+r.GetDuration()+"', '"+r.GetFlag()+"')"); err != nil {
			return nil, err
		} else {
			return &pb.CalculationResult{Result: result.Result.(string)}, nil
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
					cs = append(cs, &pb.CalculationResult{Result: result.Result.(string)})
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
		infos := []*pb.CalcItemInfo{}
		for _, info := range result.Infos {
			infos = append(infos, &pb.CalcItemInfo{Info: info})
		}
		return &pb.CalcItemsInfo{Infos: infos}, nil
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
	if _, err := updateItem(s.gdb.ItemDbPath, "update calc_cfg set status='true' where "+strings.Join(id, " or ")); err != nil {
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
	if _, err := updateItem(s.gdb.ItemDbPath, "update calc_cfg set status='false' where "+strings.Join(id, " or ")); err != nil {
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
			userAgent := md.Get("user-agent")[0] // user agent
			if methods[len(methods)-1] == "UserLogin" {
				r := req.(*pb.AuthInfo)
				if result, err := s.gdb.userLogin(authInfo{
					UserName: r.GetUserName(),
					PassWord: r.GetPassWord(),
				}, userAgent); err != nil {
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
					if v, err := s.gdb.infoDb.Get([]byte(userName+"_token"+"_"+token+"_"+userAgent), nil); err != nil || v == nil {
						return nil, status.Errorf(codes.Unauthenticated, "invalid token")
					} else {
						if token != fmt.Sprintf("%s", v) {
							return nil, status.Errorf(codes.Unauthenticated, "invalid token")
						} else {
							// log handler
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
				userAgent := md.Get("user-agent")[0] // user agent
				var au string
				if d, ok := md["authorization"]; ok {
					au = d[0]
				} else {
					return status.Errorf(codes.Unauthenticated, "invalid token")
				}
				if userName, token, ok := parseBasicAuth(au); !ok {
					return status.Errorf(codes.Unauthenticated, "invalid token")
				} else {
					if v, err := s.gdb.infoDb.Get([]byte(userName+"_token"+"_"+token+"_"+userAgent), nil); err != nil || v == nil {
						return status.Errorf(codes.Unauthenticated, "invalid token")
					} else {
						if token != fmt.Sprintf("%s", v) {
							return status.Errorf(codes.Unauthenticated, "invalid token")
						} else {
							// log handler
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
			rpcString, _ := Json.Marshal(r)
			if v, err := handler(c, req); err != nil {
				_ = s.gdb.writeLog(Error, info.FullMethod, fmt.Sprintf("%s", rpcString), "rpc", err.Error(), remoteAddress)
				return v, err
			} else {
				if s.configs.Level == Info {
					_ = s.gdb.writeLog(Error, info.FullMethod, fmt.Sprintf("%s", rpcString), "rpc", "", remoteAddress)
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
