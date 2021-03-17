// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package model

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GroupClient is the client API for Group service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GroupClient interface {
	AddGroups(ctx context.Context, in *AddedGroupInfos, opts ...grpc.CallOption) (*Rows, error)
	DeleteGroups(ctx context.Context, in *GroupNamesInfo, opts ...grpc.CallOption) (*Rows, error)
	GetGroups(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GroupNamesInfo, error)
	GetGroupProperty(ctx context.Context, in *QueryGroupPropertyInfo, opts ...grpc.CallOption) (*GroupPropertyInfo, error)
	UpdateGroupNames(ctx context.Context, in *UpdatedGroupNamesInfo, opts ...grpc.CallOption) (*Rows, error)
	UpdateGroupColumnNames(ctx context.Context, in *UpdatedGroupColumnNamesInfo, opts ...grpc.CallOption) (*Cols, error)
	DeleteGroupColumns(ctx context.Context, in *DeletedGroupColumnNamesInfo, opts ...grpc.CallOption) (*Cols, error)
	AddGroupColumns(ctx context.Context, in *AddedGroupColumnsInfo, opts ...grpc.CallOption) (*Cols, error)
}

type groupClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupClient(cc grpc.ClientConnInterface) GroupClient {
	return &groupClient{cc}
}

func (c *groupClient) AddGroups(ctx context.Context, in *AddedGroupInfos, opts ...grpc.CallOption) (*Rows, error) {
	out := new(Rows)
	err := c.cc.Invoke(ctx, "/model.Group/AddGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) DeleteGroups(ctx context.Context, in *GroupNamesInfo, opts ...grpc.CallOption) (*Rows, error) {
	out := new(Rows)
	err := c.cc.Invoke(ctx, "/model.Group/DeleteGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) GetGroups(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GroupNamesInfo, error) {
	out := new(GroupNamesInfo)
	err := c.cc.Invoke(ctx, "/model.Group/GetGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) GetGroupProperty(ctx context.Context, in *QueryGroupPropertyInfo, opts ...grpc.CallOption) (*GroupPropertyInfo, error) {
	out := new(GroupPropertyInfo)
	err := c.cc.Invoke(ctx, "/model.Group/GetGroupProperty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) UpdateGroupNames(ctx context.Context, in *UpdatedGroupNamesInfo, opts ...grpc.CallOption) (*Rows, error) {
	out := new(Rows)
	err := c.cc.Invoke(ctx, "/model.Group/UpdateGroupNames", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) UpdateGroupColumnNames(ctx context.Context, in *UpdatedGroupColumnNamesInfo, opts ...grpc.CallOption) (*Cols, error) {
	out := new(Cols)
	err := c.cc.Invoke(ctx, "/model.Group/UpdateGroupColumnNames", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) DeleteGroupColumns(ctx context.Context, in *DeletedGroupColumnNamesInfo, opts ...grpc.CallOption) (*Cols, error) {
	out := new(Cols)
	err := c.cc.Invoke(ctx, "/model.Group/DeleteGroupColumns", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) AddGroupColumns(ctx context.Context, in *AddedGroupColumnsInfo, opts ...grpc.CallOption) (*Cols, error) {
	out := new(Cols)
	err := c.cc.Invoke(ctx, "/model.Group/AddGroupColumns", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupServer is the server API for Group service.
// All implementations must embed UnimplementedGroupServer
// for forward compatibility
type GroupServer interface {
	AddGroups(context.Context, *AddedGroupInfos) (*Rows, error)
	DeleteGroups(context.Context, *GroupNamesInfo) (*Rows, error)
	GetGroups(context.Context, *emptypb.Empty) (*GroupNamesInfo, error)
	GetGroupProperty(context.Context, *QueryGroupPropertyInfo) (*GroupPropertyInfo, error)
	UpdateGroupNames(context.Context, *UpdatedGroupNamesInfo) (*Rows, error)
	UpdateGroupColumnNames(context.Context, *UpdatedGroupColumnNamesInfo) (*Cols, error)
	DeleteGroupColumns(context.Context, *DeletedGroupColumnNamesInfo) (*Cols, error)
	AddGroupColumns(context.Context, *AddedGroupColumnsInfo) (*Cols, error)
	mustEmbedUnimplementedGroupServer()
}

// UnimplementedGroupServer must be embedded to have forward compatible implementations.
type UnimplementedGroupServer struct {
}

func (UnimplementedGroupServer) AddGroups(context.Context, *AddedGroupInfos) (*Rows, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddGroups not implemented")
}
func (UnimplementedGroupServer) DeleteGroups(context.Context, *GroupNamesInfo) (*Rows, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroups not implemented")
}
func (UnimplementedGroupServer) GetGroups(context.Context, *emptypb.Empty) (*GroupNamesInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroups not implemented")
}
func (UnimplementedGroupServer) GetGroupProperty(context.Context, *QueryGroupPropertyInfo) (*GroupPropertyInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupProperty not implemented")
}
func (UnimplementedGroupServer) UpdateGroupNames(context.Context, *UpdatedGroupNamesInfo) (*Rows, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroupNames not implemented")
}
func (UnimplementedGroupServer) UpdateGroupColumnNames(context.Context, *UpdatedGroupColumnNamesInfo) (*Cols, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroupColumnNames not implemented")
}
func (UnimplementedGroupServer) DeleteGroupColumns(context.Context, *DeletedGroupColumnNamesInfo) (*Cols, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroupColumns not implemented")
}
func (UnimplementedGroupServer) AddGroupColumns(context.Context, *AddedGroupColumnsInfo) (*Cols, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddGroupColumns not implemented")
}
func (UnimplementedGroupServer) mustEmbedUnimplementedGroupServer() {}

// UnsafeGroupServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GroupServer will
// result in compilation errors.
type UnsafeGroupServer interface {
	mustEmbedUnimplementedGroupServer()
}

func RegisterGroupServer(s grpc.ServiceRegistrar, srv GroupServer) {
	s.RegisterService(&Group_ServiceDesc, srv)
}

func _Group_AddGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddedGroupInfos)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).AddGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Group/AddGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).AddGroups(ctx, req.(*AddedGroupInfos))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_DeleteGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupNamesInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).DeleteGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Group/DeleteGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).DeleteGroups(ctx, req.(*GroupNamesInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_GetGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).GetGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Group/GetGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).GetGroups(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_GetGroupProperty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGroupPropertyInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).GetGroupProperty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Group/GetGroupProperty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).GetGroupProperty(ctx, req.(*QueryGroupPropertyInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_UpdateGroupNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatedGroupNamesInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).UpdateGroupNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Group/UpdateGroupNames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).UpdateGroupNames(ctx, req.(*UpdatedGroupNamesInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_UpdateGroupColumnNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatedGroupColumnNamesInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).UpdateGroupColumnNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Group/UpdateGroupColumnNames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).UpdateGroupColumnNames(ctx, req.(*UpdatedGroupColumnNamesInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_DeleteGroupColumns_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletedGroupColumnNamesInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).DeleteGroupColumns(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Group/DeleteGroupColumns",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).DeleteGroupColumns(ctx, req.(*DeletedGroupColumnNamesInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_AddGroupColumns_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddedGroupColumnsInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).AddGroupColumns(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Group/AddGroupColumns",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).AddGroupColumns(ctx, req.(*AddedGroupColumnsInfo))
	}
	return interceptor(ctx, in, info, handler)
}

// Group_ServiceDesc is the grpc.ServiceDesc for Group service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Group_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "model.Group",
	HandlerType: (*GroupServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddGroups",
			Handler:    _Group_AddGroups_Handler,
		},
		{
			MethodName: "DeleteGroups",
			Handler:    _Group_DeleteGroups_Handler,
		},
		{
			MethodName: "GetGroups",
			Handler:    _Group_GetGroups_Handler,
		},
		{
			MethodName: "GetGroupProperty",
			Handler:    _Group_GetGroupProperty_Handler,
		},
		{
			MethodName: "UpdateGroupNames",
			Handler:    _Group_UpdateGroupNames_Handler,
		},
		{
			MethodName: "UpdateGroupColumnNames",
			Handler:    _Group_UpdateGroupColumnNames_Handler,
		},
		{
			MethodName: "DeleteGroupColumns",
			Handler:    _Group_DeleteGroupColumns_Handler,
		},
		{
			MethodName: "AddGroupColumns",
			Handler:    _Group_AddGroupColumns_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gdb.proto",
}

// ItemClient is the client API for Item service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ItemClient interface {
	AddItems(ctx context.Context, in *AddedItemsInfo, opts ...grpc.CallOption) (*Rows, error)
	DeleteItems(ctx context.Context, in *DeletedItemsInfo, opts ...grpc.CallOption) (*Rows, error)
	GetItems(ctx context.Context, in *ItemsInfo, opts ...grpc.CallOption) (*GdbItems, error)
	GetItemsWithCount(ctx context.Context, in *ItemsInfo, opts ...grpc.CallOption) (*GdbItemsWithCount, error)
	UpdateItems(ctx context.Context, in *ItemsInfoWithoutRow, opts ...grpc.CallOption) (*Rows, error)
}

type itemClient struct {
	cc grpc.ClientConnInterface
}

func NewItemClient(cc grpc.ClientConnInterface) ItemClient {
	return &itemClient{cc}
}

func (c *itemClient) AddItems(ctx context.Context, in *AddedItemsInfo, opts ...grpc.CallOption) (*Rows, error) {
	out := new(Rows)
	err := c.cc.Invoke(ctx, "/model.Item/AddItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemClient) DeleteItems(ctx context.Context, in *DeletedItemsInfo, opts ...grpc.CallOption) (*Rows, error) {
	out := new(Rows)
	err := c.cc.Invoke(ctx, "/model.Item/DeleteItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemClient) GetItems(ctx context.Context, in *ItemsInfo, opts ...grpc.CallOption) (*GdbItems, error) {
	out := new(GdbItems)
	err := c.cc.Invoke(ctx, "/model.Item/GetItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemClient) GetItemsWithCount(ctx context.Context, in *ItemsInfo, opts ...grpc.CallOption) (*GdbItemsWithCount, error) {
	out := new(GdbItemsWithCount)
	err := c.cc.Invoke(ctx, "/model.Item/GetItemsWithCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemClient) UpdateItems(ctx context.Context, in *ItemsInfoWithoutRow, opts ...grpc.CallOption) (*Rows, error) {
	out := new(Rows)
	err := c.cc.Invoke(ctx, "/model.Item/UpdateItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ItemServer is the server API for Item service.
// All implementations must embed UnimplementedItemServer
// for forward compatibility
type ItemServer interface {
	AddItems(context.Context, *AddedItemsInfo) (*Rows, error)
	DeleteItems(context.Context, *DeletedItemsInfo) (*Rows, error)
	GetItems(context.Context, *ItemsInfo) (*GdbItems, error)
	GetItemsWithCount(context.Context, *ItemsInfo) (*GdbItemsWithCount, error)
	UpdateItems(context.Context, *ItemsInfoWithoutRow) (*Rows, error)
	mustEmbedUnimplementedItemServer()
}

// UnimplementedItemServer must be embedded to have forward compatible implementations.
type UnimplementedItemServer struct {
}

func (UnimplementedItemServer) AddItems(context.Context, *AddedItemsInfo) (*Rows, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddItems not implemented")
}
func (UnimplementedItemServer) DeleteItems(context.Context, *DeletedItemsInfo) (*Rows, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteItems not implemented")
}
func (UnimplementedItemServer) GetItems(context.Context, *ItemsInfo) (*GdbItems, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItems not implemented")
}
func (UnimplementedItemServer) GetItemsWithCount(context.Context, *ItemsInfo) (*GdbItemsWithCount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItemsWithCount not implemented")
}
func (UnimplementedItemServer) UpdateItems(context.Context, *ItemsInfoWithoutRow) (*Rows, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateItems not implemented")
}
func (UnimplementedItemServer) mustEmbedUnimplementedItemServer() {}

// UnsafeItemServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ItemServer will
// result in compilation errors.
type UnsafeItemServer interface {
	mustEmbedUnimplementedItemServer()
}

func RegisterItemServer(s grpc.ServiceRegistrar, srv ItemServer) {
	s.RegisterService(&Item_ServiceDesc, srv)
}

func _Item_AddItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddedItemsInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServer).AddItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Item/AddItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServer).AddItems(ctx, req.(*AddedItemsInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Item_DeleteItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletedItemsInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServer).DeleteItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Item/DeleteItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServer).DeleteItems(ctx, req.(*DeletedItemsInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Item_GetItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemsInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServer).GetItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Item/GetItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServer).GetItems(ctx, req.(*ItemsInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Item_GetItemsWithCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemsInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServer).GetItemsWithCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Item/GetItemsWithCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServer).GetItemsWithCount(ctx, req.(*ItemsInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Item_UpdateItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemsInfoWithoutRow)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServer).UpdateItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Item/UpdateItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServer).UpdateItems(ctx, req.(*ItemsInfoWithoutRow))
	}
	return interceptor(ctx, in, info, handler)
}

// Item_ServiceDesc is the grpc.ServiceDesc for Item service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Item_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "model.Item",
	HandlerType: (*ItemServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddItems",
			Handler:    _Item_AddItems_Handler,
		},
		{
			MethodName: "DeleteItems",
			Handler:    _Item_DeleteItems_Handler,
		},
		{
			MethodName: "GetItems",
			Handler:    _Item_GetItems_Handler,
		},
		{
			MethodName: "GetItemsWithCount",
			Handler:    _Item_GetItemsWithCount_Handler,
		},
		{
			MethodName: "UpdateItems",
			Handler:    _Item_UpdateItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gdb.proto",
}

// DataClient is the client API for Data service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataClient interface {
	BatchWrite(ctx context.Context, in *BatchWriteString, opts ...grpc.CallOption) (*Rows, error)
	GetRealTimeData(ctx context.Context, in *QueryRealTimeDataString, opts ...grpc.CallOption) (*GdbRealTimeData, error)
	GetHistoricalData(ctx context.Context, in *QueryHistoricalDataString, opts ...grpc.CallOption) (*GdbHistoricalData, error)
	GetHistoricalDataWithStamp(ctx context.Context, in *QueryHistoricalDataWithTimeStampString, opts ...grpc.CallOption) (*GdbHistoricalData, error)
	GetDbInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GdbInfoData, error)
}

type dataClient struct {
	cc grpc.ClientConnInterface
}

func NewDataClient(cc grpc.ClientConnInterface) DataClient {
	return &dataClient{cc}
}

func (c *dataClient) BatchWrite(ctx context.Context, in *BatchWriteString, opts ...grpc.CallOption) (*Rows, error) {
	out := new(Rows)
	err := c.cc.Invoke(ctx, "/model.Data/BatchWrite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) GetRealTimeData(ctx context.Context, in *QueryRealTimeDataString, opts ...grpc.CallOption) (*GdbRealTimeData, error) {
	out := new(GdbRealTimeData)
	err := c.cc.Invoke(ctx, "/model.Data/GetRealTimeData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) GetHistoricalData(ctx context.Context, in *QueryHistoricalDataString, opts ...grpc.CallOption) (*GdbHistoricalData, error) {
	out := new(GdbHistoricalData)
	err := c.cc.Invoke(ctx, "/model.Data/GetHistoricalData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) GetHistoricalDataWithStamp(ctx context.Context, in *QueryHistoricalDataWithTimeStampString, opts ...grpc.CallOption) (*GdbHistoricalData, error) {
	out := new(GdbHistoricalData)
	err := c.cc.Invoke(ctx, "/model.Data/GetHistoricalDataWithStamp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) GetDbInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GdbInfoData, error) {
	out := new(GdbInfoData)
	err := c.cc.Invoke(ctx, "/model.Data/GetDbInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataServer is the server API for Data service.
// All implementations must embed UnimplementedDataServer
// for forward compatibility
type DataServer interface {
	BatchWrite(context.Context, *BatchWriteString) (*Rows, error)
	GetRealTimeData(context.Context, *QueryRealTimeDataString) (*GdbRealTimeData, error)
	GetHistoricalData(context.Context, *QueryHistoricalDataString) (*GdbHistoricalData, error)
	GetHistoricalDataWithStamp(context.Context, *QueryHistoricalDataWithTimeStampString) (*GdbHistoricalData, error)
	GetDbInfo(context.Context, *emptypb.Empty) (*GdbInfoData, error)
	mustEmbedUnimplementedDataServer()
}

// UnimplementedDataServer must be embedded to have forward compatible implementations.
type UnimplementedDataServer struct {
}

func (UnimplementedDataServer) BatchWrite(context.Context, *BatchWriteString) (*Rows, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchWrite not implemented")
}
func (UnimplementedDataServer) GetRealTimeData(context.Context, *QueryRealTimeDataString) (*GdbRealTimeData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRealTimeData not implemented")
}
func (UnimplementedDataServer) GetHistoricalData(context.Context, *QueryHistoricalDataString) (*GdbHistoricalData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHistoricalData not implemented")
}
func (UnimplementedDataServer) GetHistoricalDataWithStamp(context.Context, *QueryHistoricalDataWithTimeStampString) (*GdbHistoricalData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHistoricalDataWithStamp not implemented")
}
func (UnimplementedDataServer) GetDbInfo(context.Context, *emptypb.Empty) (*GdbInfoData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDbInfo not implemented")
}
func (UnimplementedDataServer) mustEmbedUnimplementedDataServer() {}

// UnsafeDataServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataServer will
// result in compilation errors.
type UnsafeDataServer interface {
	mustEmbedUnimplementedDataServer()
}

func RegisterDataServer(s grpc.ServiceRegistrar, srv DataServer) {
	s.RegisterService(&Data_ServiceDesc, srv)
}

func _Data_BatchWrite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchWriteString)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).BatchWrite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Data/BatchWrite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).BatchWrite(ctx, req.(*BatchWriteString))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_GetRealTimeData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRealTimeDataString)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).GetRealTimeData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Data/GetRealTimeData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).GetRealTimeData(ctx, req.(*QueryRealTimeDataString))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_GetHistoricalData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryHistoricalDataString)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).GetHistoricalData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Data/GetHistoricalData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).GetHistoricalData(ctx, req.(*QueryHistoricalDataString))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_GetHistoricalDataWithStamp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryHistoricalDataWithTimeStampString)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).GetHistoricalDataWithStamp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Data/GetHistoricalDataWithStamp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).GetHistoricalDataWithStamp(ctx, req.(*QueryHistoricalDataWithTimeStampString))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_GetDbInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).GetDbInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Data/GetDbInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).GetDbInfo(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Data_ServiceDesc is the grpc.ServiceDesc for Data service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Data_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "model.Data",
	HandlerType: (*DataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BatchWrite",
			Handler:    _Data_BatchWrite_Handler,
		},
		{
			MethodName: "GetRealTimeData",
			Handler:    _Data_GetRealTimeData_Handler,
		},
		{
			MethodName: "GetHistoricalData",
			Handler:    _Data_GetHistoricalData_Handler,
		},
		{
			MethodName: "GetHistoricalDataWithStamp",
			Handler:    _Data_GetHistoricalDataWithStamp_Handler,
		},
		{
			MethodName: "GetDbInfo",
			Handler:    _Data_GetDbInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gdb.proto",
}
