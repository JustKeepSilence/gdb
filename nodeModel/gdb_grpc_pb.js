// GENERATED CODE -- DO NOT EDIT!

'use strict';
var gdb_pb = require('./gdb_pb.js');
var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js');

function serialize_google_protobuf_Empty(arg) {
  if (!(arg instanceof google_protobuf_empty_pb.Empty)) {
    throw new Error('Expected argument of type google.protobuf.Empty');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_google_protobuf_Empty(buffer_arg) {
  return google_protobuf_empty_pb.Empty.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_AddedCalcItemInfo(arg) {
  if (!(arg instanceof gdb_pb.AddedCalcItemInfo)) {
    throw new Error('Expected argument of type model.AddedCalcItemInfo');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_AddedCalcItemInfo(buffer_arg) {
  return gdb_pb.AddedCalcItemInfo.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_AddedGroupColumnsInfo(arg) {
  if (!(arg instanceof gdb_pb.AddedGroupColumnsInfo)) {
    throw new Error('Expected argument of type model.AddedGroupColumnsInfo');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_AddedGroupColumnsInfo(buffer_arg) {
  return gdb_pb.AddedGroupColumnsInfo.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_AddedGroupInfos(arg) {
  if (!(arg instanceof gdb_pb.AddedGroupInfos)) {
    throw new Error('Expected argument of type model.AddedGroupInfos');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_AddedGroupInfos(buffer_arg) {
  return gdb_pb.AddedGroupInfos.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_AddedItemsInfo(arg) {
  if (!(arg instanceof gdb_pb.AddedItemsInfo)) {
    throw new Error('Expected argument of type model.AddedItemsInfo');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_AddedItemsInfo(buffer_arg) {
  return gdb_pb.AddedItemsInfo.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_AddedUserInfo(arg) {
  if (!(arg instanceof gdb_pb.AddedUserInfo)) {
    throw new Error('Expected argument of type model.AddedUserInfo');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_AddedUserInfo(buffer_arg) {
  return gdb_pb.AddedUserInfo.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_AuthInfo(arg) {
  if (!(arg instanceof gdb_pb.AuthInfo)) {
    throw new Error('Expected argument of type model.AuthInfo');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_AuthInfo(buffer_arg) {
  return gdb_pb.AuthInfo.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_BoolHItemValues(arg) {
  if (!(arg instanceof gdb_pb.BoolHItemValues)) {
    throw new Error('Expected argument of type model.BoolHItemValues');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_BoolHItemValues(buffer_arg) {
  return gdb_pb.BoolHItemValues.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_BoolItemValues(arg) {
  if (!(arg instanceof gdb_pb.BoolItemValues)) {
    throw new Error('Expected argument of type model.BoolItemValues');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_BoolItemValues(buffer_arg) {
  return gdb_pb.BoolItemValues.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_CalcId(arg) {
  if (!(arg instanceof gdb_pb.CalcId)) {
    throw new Error('Expected argument of type model.CalcId');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_CalcId(buffer_arg) {
  return gdb_pb.CalcId.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_CalcItemsInfo(arg) {
  if (!(arg instanceof gdb_pb.CalcItemsInfo)) {
    throw new Error('Expected argument of type model.CalcItemsInfo');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_CalcItemsInfo(buffer_arg) {
  return gdb_pb.CalcItemsInfo.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_CalculationResult(arg) {
  if (!(arg instanceof gdb_pb.CalculationResult)) {
    throw new Error('Expected argument of type model.CalculationResult');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_CalculationResult(buffer_arg) {
  return gdb_pb.CalculationResult.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_CalculationResults(arg) {
  if (!(arg instanceof gdb_pb.CalculationResults)) {
    throw new Error('Expected argument of type model.CalculationResults');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_CalculationResults(buffer_arg) {
  return gdb_pb.CalculationResults.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_CheckItemsInfo(arg) {
  if (!(arg instanceof gdb_pb.CheckItemsInfo)) {
    throw new Error('Expected argument of type model.CheckItemsInfo');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_CheckItemsInfo(buffer_arg) {
  return gdb_pb.CheckItemsInfo.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_CheckResult(arg) {
  if (!(arg instanceof gdb_pb.CheckResult)) {
    throw new Error('Expected argument of type model.CheckResult');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_CheckResult(buffer_arg) {
  return gdb_pb.CheckResult.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_Code(arg) {
  if (!(arg instanceof gdb_pb.Code)) {
    throw new Error('Expected argument of type model.Code');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_Code(buffer_arg) {
  return gdb_pb.Code.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_DeleteHistoricalDataString(arg) {
  if (!(arg instanceof gdb_pb.DeleteHistoricalDataString)) {
    throw new Error('Expected argument of type model.DeleteHistoricalDataString');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_DeleteHistoricalDataString(buffer_arg) {
  return gdb_pb.DeleteHistoricalDataString.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_DeletedGroupColumnNamesInfo(arg) {
  if (!(arg instanceof gdb_pb.DeletedGroupColumnNamesInfo)) {
    throw new Error('Expected argument of type model.DeletedGroupColumnNamesInfo');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_DeletedGroupColumnNamesInfo(buffer_arg) {
  return gdb_pb.DeletedGroupColumnNamesInfo.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_DeletedItemsInfo(arg) {
  if (!(arg instanceof gdb_pb.DeletedItemsInfo)) {
    throw new Error('Expected argument of type model.DeletedItemsInfo');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_DeletedItemsInfo(buffer_arg) {
  return gdb_pb.DeletedItemsInfo.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_DeletedLogInfo(arg) {
  if (!(arg instanceof gdb_pb.DeletedLogInfo)) {
    throw new Error('Expected argument of type model.DeletedLogInfo');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_DeletedLogInfo(buffer_arg) {
  return gdb_pb.DeletedLogInfo.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_FileContents(arg) {
  if (!(arg instanceof gdb_pb.FileContents)) {
    throw new Error('Expected argument of type model.FileContents');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_FileContents(buffer_arg) {
  return gdb_pb.FileContents.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_FileInfo(arg) {
  if (!(arg instanceof gdb_pb.FileInfo)) {
    throw new Error('Expected argument of type model.FileInfo');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_FileInfo(buffer_arg) {
  return gdb_pb.FileInfo.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_FileSize(arg) {
  if (!(arg instanceof gdb_pb.FileSize)) {
    throw new Error('Expected argument of type model.FileSize');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_FileSize(buffer_arg) {
  return gdb_pb.FileSize.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_FloatHItemValues(arg) {
  if (!(arg instanceof gdb_pb.FloatHItemValues)) {
    throw new Error('Expected argument of type model.FloatHItemValues');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_FloatHItemValues(buffer_arg) {
  return gdb_pb.FloatHItemValues.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_FloatItemValues(arg) {
  if (!(arg instanceof gdb_pb.FloatItemValues)) {
    throw new Error('Expected argument of type model.FloatItemValues');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_FloatItemValues(buffer_arg) {
  return gdb_pb.FloatItemValues.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_GdbHistoricalData(arg) {
  if (!(arg instanceof gdb_pb.GdbHistoricalData)) {
    throw new Error('Expected argument of type model.GdbHistoricalData');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_GdbHistoricalData(buffer_arg) {
  return gdb_pb.GdbHistoricalData.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_GdbInfoData(arg) {
  if (!(arg instanceof gdb_pb.GdbInfoData)) {
    throw new Error('Expected argument of type model.GdbInfoData');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_GdbInfoData(buffer_arg) {
  return gdb_pb.GdbInfoData.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_GdbItemsWithCount(arg) {
  if (!(arg instanceof gdb_pb.GdbItemsWithCount)) {
    throw new Error('Expected argument of type model.GdbItemsWithCount');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_GdbItemsWithCount(buffer_arg) {
  return gdb_pb.GdbItemsWithCount.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_GdbRealTimeData(arg) {
  if (!(arg instanceof gdb_pb.GdbRealTimeData)) {
    throw new Error('Expected argument of type model.GdbRealTimeData');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_GdbRealTimeData(buffer_arg) {
  return gdb_pb.GdbRealTimeData.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_GroupNamesInfo(arg) {
  if (!(arg instanceof gdb_pb.GroupNamesInfo)) {
    throw new Error('Expected argument of type model.GroupNamesInfo');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_GroupNamesInfo(buffer_arg) {
  return gdb_pb.GroupNamesInfo.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_GroupPropertyInfo(arg) {
  if (!(arg instanceof gdb_pb.GroupPropertyInfo)) {
    throw new Error('Expected argument of type model.GroupPropertyInfo');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_GroupPropertyInfo(buffer_arg) {
  return gdb_pb.GroupPropertyInfo.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_HistoryFileInfo(arg) {
  if (!(arg instanceof gdb_pb.HistoryFileInfo)) {
    throw new Error('Expected argument of type model.HistoryFileInfo');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_HistoryFileInfo(buffer_arg) {
  return gdb_pb.HistoryFileInfo.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_IntHItemValues(arg) {
  if (!(arg instanceof gdb_pb.IntHItemValues)) {
    throw new Error('Expected argument of type model.IntHItemValues');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_IntHItemValues(buffer_arg) {
  return gdb_pb.IntHItemValues.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_IntItemValues(arg) {
  if (!(arg instanceof gdb_pb.IntItemValues)) {
    throw new Error('Expected argument of type model.IntItemValues');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_IntItemValues(buffer_arg) {
  return gdb_pb.IntItemValues.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_ItemsInfo(arg) {
  if (!(arg instanceof gdb_pb.ItemsInfo)) {
    throw new Error('Expected argument of type model.ItemsInfo');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_ItemsInfo(buffer_arg) {
  return gdb_pb.ItemsInfo.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_LogsInfo(arg) {
  if (!(arg instanceof gdb_pb.LogsInfo)) {
    throw new Error('Expected argument of type model.LogsInfo');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_LogsInfo(buffer_arg) {
  return gdb_pb.LogsInfo.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_QueryCalcItemsInfo(arg) {
  if (!(arg instanceof gdb_pb.QueryCalcItemsInfo)) {
    throw new Error('Expected argument of type model.QueryCalcItemsInfo');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_QueryCalcItemsInfo(buffer_arg) {
  return gdb_pb.QueryCalcItemsInfo.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_QueryGroupPropertyInfo(arg) {
  if (!(arg instanceof gdb_pb.QueryGroupPropertyInfo)) {
    throw new Error('Expected argument of type model.QueryGroupPropertyInfo');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_QueryGroupPropertyInfo(buffer_arg) {
  return gdb_pb.QueryGroupPropertyInfo.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_QueryHistoricalDataString(arg) {
  if (!(arg instanceof gdb_pb.QueryHistoricalDataString)) {
    throw new Error('Expected argument of type model.QueryHistoricalDataString');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_QueryHistoricalDataString(buffer_arg) {
  return gdb_pb.QueryHistoricalDataString.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_QueryHistoricalDataWithConditionString(arg) {
  if (!(arg instanceof gdb_pb.QueryHistoricalDataWithConditionString)) {
    throw new Error('Expected argument of type model.QueryHistoricalDataWithConditionString');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_QueryHistoricalDataWithConditionString(buffer_arg) {
  return gdb_pb.QueryHistoricalDataWithConditionString.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_QueryHistoricalDataWithStampString(arg) {
  if (!(arg instanceof gdb_pb.QueryHistoricalDataWithStampString)) {
    throw new Error('Expected argument of type model.QueryHistoricalDataWithStampString');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_QueryHistoricalDataWithStampString(buffer_arg) {
  return gdb_pb.QueryHistoricalDataWithStampString.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_QueryLogsInfo(arg) {
  if (!(arg instanceof gdb_pb.QueryLogsInfo)) {
    throw new Error('Expected argument of type model.QueryLogsInfo');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_QueryLogsInfo(buffer_arg) {
  return gdb_pb.QueryLogsInfo.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_QueryRawHistoricalDataString(arg) {
  if (!(arg instanceof gdb_pb.QueryRawHistoricalDataString)) {
    throw new Error('Expected argument of type model.QueryRawHistoricalDataString');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_QueryRawHistoricalDataString(buffer_arg) {
  return gdb_pb.QueryRawHistoricalDataString.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_QueryRealTimeDataString(arg) {
  if (!(arg instanceof gdb_pb.QueryRealTimeDataString)) {
    throw new Error('Expected argument of type model.QueryRealTimeDataString');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_QueryRealTimeDataString(buffer_arg) {
  return gdb_pb.QueryRealTimeDataString.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_QuerySpeedHistoryDataString(arg) {
  if (!(arg instanceof gdb_pb.QuerySpeedHistoryDataString)) {
    throw new Error('Expected argument of type model.QuerySpeedHistoryDataString');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_QuerySpeedHistoryDataString(buffer_arg) {
  return gdb_pb.QuerySpeedHistoryDataString.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_Routes(arg) {
  if (!(arg instanceof gdb_pb.Routes)) {
    throw new Error('Expected argument of type model.Routes');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_Routes(buffer_arg) {
  return gdb_pb.Routes.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_RoutesInfo(arg) {
  if (!(arg instanceof gdb_pb.RoutesInfo)) {
    throw new Error('Expected argument of type model.RoutesInfo');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_RoutesInfo(buffer_arg) {
  return gdb_pb.RoutesInfo.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_StringHItemValues(arg) {
  if (!(arg instanceof gdb_pb.StringHItemValues)) {
    throw new Error('Expected argument of type model.StringHItemValues');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_StringHItemValues(buffer_arg) {
  return gdb_pb.StringHItemValues.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_StringItemValues(arg) {
  if (!(arg instanceof gdb_pb.StringItemValues)) {
    throw new Error('Expected argument of type model.StringItemValues');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_StringItemValues(buffer_arg) {
  return gdb_pb.StringItemValues.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_TestCalcItemInfo(arg) {
  if (!(arg instanceof gdb_pb.TestCalcItemInfo)) {
    throw new Error('Expected argument of type model.TestCalcItemInfo');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_TestCalcItemInfo(buffer_arg) {
  return gdb_pb.TestCalcItemInfo.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_TimeCols(arg) {
  if (!(arg instanceof gdb_pb.TimeCols)) {
    throw new Error('Expected argument of type model.TimeCols');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_TimeCols(buffer_arg) {
  return gdb_pb.TimeCols.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_TimeRows(arg) {
  if (!(arg instanceof gdb_pb.TimeRows)) {
    throw new Error('Expected argument of type model.TimeRows');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_TimeRows(buffer_arg) {
  return gdb_pb.TimeRows.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_UpdatedCalcInfo(arg) {
  if (!(arg instanceof gdb_pb.UpdatedCalcInfo)) {
    throw new Error('Expected argument of type model.UpdatedCalcInfo');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_UpdatedCalcInfo(buffer_arg) {
  return gdb_pb.UpdatedCalcInfo.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_UpdatedGroupColumnNamesInfo(arg) {
  if (!(arg instanceof gdb_pb.UpdatedGroupColumnNamesInfo)) {
    throw new Error('Expected argument of type model.UpdatedGroupColumnNamesInfo');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_UpdatedGroupColumnNamesInfo(buffer_arg) {
  return gdb_pb.UpdatedGroupColumnNamesInfo.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_UpdatedGroupNamesInfo(arg) {
  if (!(arg instanceof gdb_pb.UpdatedGroupNamesInfo)) {
    throw new Error('Expected argument of type model.UpdatedGroupNamesInfo');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_UpdatedGroupNamesInfo(buffer_arg) {
  return gdb_pb.UpdatedGroupNamesInfo.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_UpdatedItemsInfo(arg) {
  if (!(arg instanceof gdb_pb.UpdatedItemsInfo)) {
    throw new Error('Expected argument of type model.UpdatedItemsInfo');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_UpdatedItemsInfo(buffer_arg) {
  return gdb_pb.UpdatedItemsInfo.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_UpdatedUserInfo(arg) {
  if (!(arg instanceof gdb_pb.UpdatedUserInfo)) {
    throw new Error('Expected argument of type model.UpdatedUserInfo');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_UpdatedUserInfo(buffer_arg) {
  return gdb_pb.UpdatedUserInfo.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_UploadedFileInfo(arg) {
  if (!(arg instanceof gdb_pb.UploadedFileInfo)) {
    throw new Error('Expected argument of type model.UploadedFileInfo');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_UploadedFileInfo(buffer_arg) {
  return gdb_pb.UploadedFileInfo.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_UserInfo(arg) {
  if (!(arg instanceof gdb_pb.UserInfo)) {
    throw new Error('Expected argument of type model.UserInfo');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_UserInfo(buffer_arg) {
  return gdb_pb.UserInfo.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_UserInfos(arg) {
  if (!(arg instanceof gdb_pb.UserInfos)) {
    throw new Error('Expected argument of type model.UserInfos');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_UserInfos(buffer_arg) {
  return gdb_pb.UserInfos.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_UserName(arg) {
  if (!(arg instanceof gdb_pb.UserName)) {
    throw new Error('Expected argument of type model.UserName');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_UserName(buffer_arg) {
  return gdb_pb.UserName.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_model_UserToken(arg) {
  if (!(arg instanceof gdb_pb.UserToken)) {
    throw new Error('Expected argument of type model.UserToken');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_model_UserToken(buffer_arg) {
  return gdb_pb.UserToken.deserializeBinary(new Uint8Array(buffer_arg));
}


var GroupService = exports['model.Group'] = {
  addGroups: {
    path: '/model.Group/AddGroups',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.AddedGroupInfos,
    responseType: gdb_pb.TimeRows,
    requestSerialize: serialize_model_AddedGroupInfos,
    requestDeserialize: deserialize_model_AddedGroupInfos,
    responseSerialize: serialize_model_TimeRows,
    responseDeserialize: deserialize_model_TimeRows,
  },
  deleteGroups: {
    path: '/model.Group/DeleteGroups',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.GroupNamesInfo,
    responseType: gdb_pb.TimeRows,
    requestSerialize: serialize_model_GroupNamesInfo,
    requestDeserialize: deserialize_model_GroupNamesInfo,
    responseSerialize: serialize_model_TimeRows,
    responseDeserialize: deserialize_model_TimeRows,
  },
  getGroups: {
    path: '/model.Group/GetGroups',
    requestStream: false,
    responseStream: false,
    requestType: google_protobuf_empty_pb.Empty,
    responseType: gdb_pb.GroupNamesInfo,
    requestSerialize: serialize_google_protobuf_Empty,
    requestDeserialize: deserialize_google_protobuf_Empty,
    responseSerialize: serialize_model_GroupNamesInfo,
    responseDeserialize: deserialize_model_GroupNamesInfo,
  },
  getGroupProperty: {
    path: '/model.Group/GetGroupProperty',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.QueryGroupPropertyInfo,
    responseType: gdb_pb.GroupPropertyInfo,
    requestSerialize: serialize_model_QueryGroupPropertyInfo,
    requestDeserialize: deserialize_model_QueryGroupPropertyInfo,
    responseSerialize: serialize_model_GroupPropertyInfo,
    responseDeserialize: deserialize_model_GroupPropertyInfo,
  },
  updateGroupNames: {
    path: '/model.Group/UpdateGroupNames',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.UpdatedGroupNamesInfo,
    responseType: gdb_pb.TimeRows,
    requestSerialize: serialize_model_UpdatedGroupNamesInfo,
    requestDeserialize: deserialize_model_UpdatedGroupNamesInfo,
    responseSerialize: serialize_model_TimeRows,
    responseDeserialize: deserialize_model_TimeRows,
  },
  updateGroupColumnNames: {
    path: '/model.Group/UpdateGroupColumnNames',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.UpdatedGroupColumnNamesInfo,
    responseType: gdb_pb.TimeCols,
    requestSerialize: serialize_model_UpdatedGroupColumnNamesInfo,
    requestDeserialize: deserialize_model_UpdatedGroupColumnNamesInfo,
    responseSerialize: serialize_model_TimeCols,
    responseDeserialize: deserialize_model_TimeCols,
  },
  deleteGroupColumns: {
    path: '/model.Group/DeleteGroupColumns',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.DeletedGroupColumnNamesInfo,
    responseType: gdb_pb.TimeCols,
    requestSerialize: serialize_model_DeletedGroupColumnNamesInfo,
    requestDeserialize: deserialize_model_DeletedGroupColumnNamesInfo,
    responseSerialize: serialize_model_TimeCols,
    responseDeserialize: deserialize_model_TimeCols,
  },
  addGroupColumns: {
    path: '/model.Group/AddGroupColumns',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.AddedGroupColumnsInfo,
    responseType: gdb_pb.TimeCols,
    requestSerialize: serialize_model_AddedGroupColumnsInfo,
    requestDeserialize: deserialize_model_AddedGroupColumnsInfo,
    responseSerialize: serialize_model_TimeCols,
    responseDeserialize: deserialize_model_TimeCols,
  },
};

var ItemService = exports['model.Item'] = {
  addItems: {
    path: '/model.Item/AddItems',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.AddedItemsInfo,
    responseType: gdb_pb.TimeRows,
    requestSerialize: serialize_model_AddedItemsInfo,
    requestDeserialize: deserialize_model_AddedItemsInfo,
    responseSerialize: serialize_model_TimeRows,
    responseDeserialize: deserialize_model_TimeRows,
  },
  deleteItems: {
    path: '/model.Item/DeleteItems',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.DeletedItemsInfo,
    responseType: gdb_pb.TimeRows,
    requestSerialize: serialize_model_DeletedItemsInfo,
    requestDeserialize: deserialize_model_DeletedItemsInfo,
    responseSerialize: serialize_model_TimeRows,
    responseDeserialize: deserialize_model_TimeRows,
  },
  getItemsWithCount: {
    path: '/model.Item/GetItemsWithCount',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.ItemsInfo,
    responseType: gdb_pb.GdbItemsWithCount,
    requestSerialize: serialize_model_ItemsInfo,
    requestDeserialize: deserialize_model_ItemsInfo,
    responseSerialize: serialize_model_GdbItemsWithCount,
    responseDeserialize: deserialize_model_GdbItemsWithCount,
  },
  updateItems: {
    path: '/model.Item/UpdateItems',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.UpdatedItemsInfo,
    responseType: gdb_pb.TimeRows,
    requestSerialize: serialize_model_UpdatedItemsInfo,
    requestDeserialize: deserialize_model_UpdatedItemsInfo,
    responseSerialize: serialize_model_TimeRows,
    responseDeserialize: deserialize_model_TimeRows,
  },
  checkItems: {
    path: '/model.Item/CheckItems',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.CheckItemsInfo,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_model_CheckItemsInfo,
    requestDeserialize: deserialize_model_CheckItemsInfo,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
  cleanGroupItems: {
    path: '/model.Item/CleanGroupItems',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.GroupNamesInfo,
    responseType: gdb_pb.TimeRows,
    requestSerialize: serialize_model_GroupNamesInfo,
    requestDeserialize: deserialize_model_GroupNamesInfo,
    responseSerialize: serialize_model_TimeRows,
    responseDeserialize: deserialize_model_TimeRows,
  },
};

var DataService = exports['model.Data'] = {
  batchWriteFloatData: {
    path: '/model.Data/BatchWriteFloatData',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.FloatItemValues,
    responseType: gdb_pb.TimeRows,
    requestSerialize: serialize_model_FloatItemValues,
    requestDeserialize: deserialize_model_FloatItemValues,
    responseSerialize: serialize_model_TimeRows,
    responseDeserialize: deserialize_model_TimeRows,
  },
  batchWriteFloatDataWithStream: {
    path: '/model.Data/BatchWriteFloatDataWithStream',
    requestStream: true,
    responseStream: false,
    requestType: gdb_pb.FloatItemValues,
    responseType: gdb_pb.TimeRows,
    requestSerialize: serialize_model_FloatItemValues,
    requestDeserialize: deserialize_model_FloatItemValues,
    responseSerialize: serialize_model_TimeRows,
    responseDeserialize: deserialize_model_TimeRows,
  },
  batchWriteIntData: {
    path: '/model.Data/BatchWriteIntData',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.IntItemValues,
    responseType: gdb_pb.TimeRows,
    requestSerialize: serialize_model_IntItemValues,
    requestDeserialize: deserialize_model_IntItemValues,
    responseSerialize: serialize_model_TimeRows,
    responseDeserialize: deserialize_model_TimeRows,
  },
  batchWriteIntDataWithStream: {
    path: '/model.Data/BatchWriteIntDataWithStream',
    requestStream: true,
    responseStream: false,
    requestType: gdb_pb.IntItemValues,
    responseType: gdb_pb.TimeRows,
    requestSerialize: serialize_model_IntItemValues,
    requestDeserialize: deserialize_model_IntItemValues,
    responseSerialize: serialize_model_TimeRows,
    responseDeserialize: deserialize_model_TimeRows,
  },
  batchWriteStringData: {
    path: '/model.Data/BatchWriteStringData',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.StringItemValues,
    responseType: gdb_pb.TimeRows,
    requestSerialize: serialize_model_StringItemValues,
    requestDeserialize: deserialize_model_StringItemValues,
    responseSerialize: serialize_model_TimeRows,
    responseDeserialize: deserialize_model_TimeRows,
  },
  batchWriteStringDataWithStream: {
    path: '/model.Data/BatchWriteStringDataWithStream',
    requestStream: true,
    responseStream: false,
    requestType: gdb_pb.StringItemValues,
    responseType: gdb_pb.TimeRows,
    requestSerialize: serialize_model_StringItemValues,
    requestDeserialize: deserialize_model_StringItemValues,
    responseSerialize: serialize_model_TimeRows,
    responseDeserialize: deserialize_model_TimeRows,
  },
  batchWriteBoolData: {
    path: '/model.Data/BatchWriteBoolData',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.BoolItemValues,
    responseType: gdb_pb.TimeRows,
    requestSerialize: serialize_model_BoolItemValues,
    requestDeserialize: deserialize_model_BoolItemValues,
    responseSerialize: serialize_model_TimeRows,
    responseDeserialize: deserialize_model_TimeRows,
  },
  batchWriteBoolDataWithStream: {
    path: '/model.Data/BatchWriteBoolDataWithStream',
    requestStream: true,
    responseStream: false,
    requestType: gdb_pb.BoolItemValues,
    responseType: gdb_pb.TimeRows,
    requestSerialize: serialize_model_BoolItemValues,
    requestDeserialize: deserialize_model_BoolItemValues,
    responseSerialize: serialize_model_TimeRows,
    responseDeserialize: deserialize_model_TimeRows,
  },
  batchWriteFloatHistoricalData: {
    path: '/model.Data/BatchWriteFloatHistoricalData',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.FloatHItemValues,
    responseType: gdb_pb.TimeRows,
    requestSerialize: serialize_model_FloatHItemValues,
    requestDeserialize: deserialize_model_FloatHItemValues,
    responseSerialize: serialize_model_TimeRows,
    responseDeserialize: deserialize_model_TimeRows,
  },
  batchWriteFloatHistoricalDataWithStream: {
    path: '/model.Data/BatchWriteFloatHistoricalDataWithStream',
    requestStream: true,
    responseStream: false,
    requestType: gdb_pb.FloatHItemValues,
    responseType: gdb_pb.TimeRows,
    requestSerialize: serialize_model_FloatHItemValues,
    requestDeserialize: deserialize_model_FloatHItemValues,
    responseSerialize: serialize_model_TimeRows,
    responseDeserialize: deserialize_model_TimeRows,
  },
  batchWriteIntHistoricalData: {
    path: '/model.Data/BatchWriteIntHistoricalData',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.IntHItemValues,
    responseType: gdb_pb.TimeRows,
    requestSerialize: serialize_model_IntHItemValues,
    requestDeserialize: deserialize_model_IntHItemValues,
    responseSerialize: serialize_model_TimeRows,
    responseDeserialize: deserialize_model_TimeRows,
  },
  batchWriteIntHistoricalDataWithStream: {
    path: '/model.Data/BatchWriteIntHistoricalDataWithStream',
    requestStream: true,
    responseStream: false,
    requestType: gdb_pb.IntHItemValues,
    responseType: gdb_pb.TimeRows,
    requestSerialize: serialize_model_IntHItemValues,
    requestDeserialize: deserialize_model_IntHItemValues,
    responseSerialize: serialize_model_TimeRows,
    responseDeserialize: deserialize_model_TimeRows,
  },
  batchWriteStringHistoricalData: {
    path: '/model.Data/BatchWriteStringHistoricalData',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.StringHItemValues,
    responseType: gdb_pb.TimeRows,
    requestSerialize: serialize_model_StringHItemValues,
    requestDeserialize: deserialize_model_StringHItemValues,
    responseSerialize: serialize_model_TimeRows,
    responseDeserialize: deserialize_model_TimeRows,
  },
  batchWriteStringHistoricalDataWithStream: {
    path: '/model.Data/BatchWriteStringHistoricalDataWithStream',
    requestStream: true,
    responseStream: false,
    requestType: gdb_pb.StringHItemValues,
    responseType: gdb_pb.TimeRows,
    requestSerialize: serialize_model_StringHItemValues,
    requestDeserialize: deserialize_model_StringHItemValues,
    responseSerialize: serialize_model_TimeRows,
    responseDeserialize: deserialize_model_TimeRows,
  },
  batchWriteBoolHistoricalData: {
    path: '/model.Data/BatchWriteBoolHistoricalData',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.BoolHItemValues,
    responseType: gdb_pb.TimeRows,
    requestSerialize: serialize_model_BoolHItemValues,
    requestDeserialize: deserialize_model_BoolHItemValues,
    responseSerialize: serialize_model_TimeRows,
    responseDeserialize: deserialize_model_TimeRows,
  },
  batchWriteBoolHistoricalDataWithStream: {
    path: '/model.Data/BatchWriteBoolHistoricalDataWithStream',
    requestStream: true,
    responseStream: false,
    requestType: gdb_pb.BoolHItemValues,
    responseType: gdb_pb.TimeRows,
    requestSerialize: serialize_model_BoolHItemValues,
    requestDeserialize: deserialize_model_BoolHItemValues,
    responseSerialize: serialize_model_TimeRows,
    responseDeserialize: deserialize_model_TimeRows,
  },
  getRealTimeData: {
    path: '/model.Data/GetRealTimeData',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.QueryRealTimeDataString,
    responseType: gdb_pb.GdbRealTimeData,
    requestSerialize: serialize_model_QueryRealTimeDataString,
    requestDeserialize: deserialize_model_QueryRealTimeDataString,
    responseSerialize: serialize_model_GdbRealTimeData,
    responseDeserialize: deserialize_model_GdbRealTimeData,
  },
  getFloatHistoricalData: {
    path: '/model.Data/GetFloatHistoricalData',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.QueryHistoricalDataString,
    responseType: gdb_pb.GdbHistoricalData,
    requestSerialize: serialize_model_QueryHistoricalDataString,
    requestDeserialize: deserialize_model_QueryHistoricalDataString,
    responseSerialize: serialize_model_GdbHistoricalData,
    responseDeserialize: deserialize_model_GdbHistoricalData,
  },
  getIntHistoricalData: {
    path: '/model.Data/GetIntHistoricalData',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.QueryHistoricalDataString,
    responseType: gdb_pb.GdbHistoricalData,
    requestSerialize: serialize_model_QueryHistoricalDataString,
    requestDeserialize: deserialize_model_QueryHistoricalDataString,
    responseSerialize: serialize_model_GdbHistoricalData,
    responseDeserialize: deserialize_model_GdbHistoricalData,
  },
  getStringHistoricalData: {
    path: '/model.Data/GetStringHistoricalData',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.QueryHistoricalDataString,
    responseType: gdb_pb.GdbHistoricalData,
    requestSerialize: serialize_model_QueryHistoricalDataString,
    requestDeserialize: deserialize_model_QueryHistoricalDataString,
    responseSerialize: serialize_model_GdbHistoricalData,
    responseDeserialize: deserialize_model_GdbHistoricalData,
  },
  getBoolHistoricalData: {
    path: '/model.Data/GetBoolHistoricalData',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.QueryHistoricalDataString,
    responseType: gdb_pb.GdbHistoricalData,
    requestSerialize: serialize_model_QueryHistoricalDataString,
    requestDeserialize: deserialize_model_QueryHistoricalDataString,
    responseSerialize: serialize_model_GdbHistoricalData,
    responseDeserialize: deserialize_model_GdbHistoricalData,
  },
  getFloatRawHistoricalData: {
    path: '/model.Data/GetFloatRawHistoricalData',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.QueryRawHistoricalDataString,
    responseType: gdb_pb.GdbHistoricalData,
    requestSerialize: serialize_model_QueryRawHistoricalDataString,
    requestDeserialize: deserialize_model_QueryRawHistoricalDataString,
    responseSerialize: serialize_model_GdbHistoricalData,
    responseDeserialize: deserialize_model_GdbHistoricalData,
  },
  getIntRawHistoricalData: {
    path: '/model.Data/GetIntRawHistoricalData',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.QueryRawHistoricalDataString,
    responseType: gdb_pb.GdbHistoricalData,
    requestSerialize: serialize_model_QueryRawHistoricalDataString,
    requestDeserialize: deserialize_model_QueryRawHistoricalDataString,
    responseSerialize: serialize_model_GdbHistoricalData,
    responseDeserialize: deserialize_model_GdbHistoricalData,
  },
  getStringRawHistoricalData: {
    path: '/model.Data/GetStringRawHistoricalData',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.QueryRawHistoricalDataString,
    responseType: gdb_pb.GdbHistoricalData,
    requestSerialize: serialize_model_QueryRawHistoricalDataString,
    requestDeserialize: deserialize_model_QueryRawHistoricalDataString,
    responseSerialize: serialize_model_GdbHistoricalData,
    responseDeserialize: deserialize_model_GdbHistoricalData,
  },
  getBoolRawHistoricalData: {
    path: '/model.Data/GetBoolRawHistoricalData',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.QueryRawHistoricalDataString,
    responseType: gdb_pb.GdbHistoricalData,
    requestSerialize: serialize_model_QueryRawHistoricalDataString,
    requestDeserialize: deserialize_model_QueryRawHistoricalDataString,
    responseSerialize: serialize_model_GdbHistoricalData,
    responseDeserialize: deserialize_model_GdbHistoricalData,
  },
  getFloatHistoricalDataWithStamp: {
    path: '/model.Data/GetFloatHistoricalDataWithStamp',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.QueryHistoricalDataWithStampString,
    responseType: gdb_pb.GdbHistoricalData,
    requestSerialize: serialize_model_QueryHistoricalDataWithStampString,
    requestDeserialize: deserialize_model_QueryHistoricalDataWithStampString,
    responseSerialize: serialize_model_GdbHistoricalData,
    responseDeserialize: deserialize_model_GdbHistoricalData,
  },
  getIntHistoricalDataWithStamp: {
    path: '/model.Data/GetIntHistoricalDataWithStamp',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.QueryHistoricalDataWithStampString,
    responseType: gdb_pb.GdbHistoricalData,
    requestSerialize: serialize_model_QueryHistoricalDataWithStampString,
    requestDeserialize: deserialize_model_QueryHistoricalDataWithStampString,
    responseSerialize: serialize_model_GdbHistoricalData,
    responseDeserialize: deserialize_model_GdbHistoricalData,
  },
  getStringHistoricalDataWithStamp: {
    path: '/model.Data/GetStringHistoricalDataWithStamp',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.QueryHistoricalDataWithStampString,
    responseType: gdb_pb.GdbHistoricalData,
    requestSerialize: serialize_model_QueryHistoricalDataWithStampString,
    requestDeserialize: deserialize_model_QueryHistoricalDataWithStampString,
    responseSerialize: serialize_model_GdbHistoricalData,
    responseDeserialize: deserialize_model_GdbHistoricalData,
  },
  getBoolHistoricalDataWithStamp: {
    path: '/model.Data/GetBoolHistoricalDataWithStamp',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.QueryHistoricalDataWithStampString,
    responseType: gdb_pb.GdbHistoricalData,
    requestSerialize: serialize_model_QueryHistoricalDataWithStampString,
    requestDeserialize: deserialize_model_QueryHistoricalDataWithStampString,
    responseSerialize: serialize_model_GdbHistoricalData,
    responseDeserialize: deserialize_model_GdbHistoricalData,
  },
  getFloatHistoricalDataWithCondition: {
    path: '/model.Data/GetFloatHistoricalDataWithCondition',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.QueryHistoricalDataWithConditionString,
    responseType: gdb_pb.GdbHistoricalData,
    requestSerialize: serialize_model_QueryHistoricalDataWithConditionString,
    requestDeserialize: deserialize_model_QueryHistoricalDataWithConditionString,
    responseSerialize: serialize_model_GdbHistoricalData,
    responseDeserialize: deserialize_model_GdbHistoricalData,
  },
  getIntHistoricalDataWithCondition: {
    path: '/model.Data/GetIntHistoricalDataWithCondition',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.QueryHistoricalDataWithConditionString,
    responseType: gdb_pb.GdbHistoricalData,
    requestSerialize: serialize_model_QueryHistoricalDataWithConditionString,
    requestDeserialize: deserialize_model_QueryHistoricalDataWithConditionString,
    responseSerialize: serialize_model_GdbHistoricalData,
    responseDeserialize: deserialize_model_GdbHistoricalData,
  },
  getStringHistoricalDataWithCondition: {
    path: '/model.Data/GetStringHistoricalDataWithCondition',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.QueryHistoricalDataWithConditionString,
    responseType: gdb_pb.GdbHistoricalData,
    requestSerialize: serialize_model_QueryHistoricalDataWithConditionString,
    requestDeserialize: deserialize_model_QueryHistoricalDataWithConditionString,
    responseSerialize: serialize_model_GdbHistoricalData,
    responseDeserialize: deserialize_model_GdbHistoricalData,
  },
  getBoolHistoricalDataWithCondition: {
    path: '/model.Data/GetBoolHistoricalDataWithCondition',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.QueryHistoricalDataWithConditionString,
    responseType: gdb_pb.GdbHistoricalData,
    requestSerialize: serialize_model_QueryHistoricalDataWithConditionString,
    requestDeserialize: deserialize_model_QueryHistoricalDataWithConditionString,
    responseSerialize: serialize_model_GdbHistoricalData,
    responseDeserialize: deserialize_model_GdbHistoricalData,
  },
  deleteFloatHistoricalData: {
    path: '/model.Data/DeleteFloatHistoricalData',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.DeleteHistoricalDataString,
    responseType: gdb_pb.TimeRows,
    requestSerialize: serialize_model_DeleteHistoricalDataString,
    requestDeserialize: deserialize_model_DeleteHistoricalDataString,
    responseSerialize: serialize_model_TimeRows,
    responseDeserialize: deserialize_model_TimeRows,
  },
  deleteIntHistoricalData: {
    path: '/model.Data/DeleteIntHistoricalData',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.DeleteHistoricalDataString,
    responseType: gdb_pb.TimeRows,
    requestSerialize: serialize_model_DeleteHistoricalDataString,
    requestDeserialize: deserialize_model_DeleteHistoricalDataString,
    responseSerialize: serialize_model_TimeRows,
    responseDeserialize: deserialize_model_TimeRows,
  },
  deleteStringHistoricalData: {
    path: '/model.Data/DeleteStringHistoricalData',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.DeleteHistoricalDataString,
    responseType: gdb_pb.TimeRows,
    requestSerialize: serialize_model_DeleteHistoricalDataString,
    requestDeserialize: deserialize_model_DeleteHistoricalDataString,
    responseSerialize: serialize_model_TimeRows,
    responseDeserialize: deserialize_model_TimeRows,
  },
  deleteBoolHistoricalData: {
    path: '/model.Data/DeleteBoolHistoricalData',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.DeleteHistoricalDataString,
    responseType: gdb_pb.TimeRows,
    requestSerialize: serialize_model_DeleteHistoricalDataString,
    requestDeserialize: deserialize_model_DeleteHistoricalDataString,
    responseSerialize: serialize_model_TimeRows,
    responseDeserialize: deserialize_model_TimeRows,
  },
  cleanItemData: {
    path: '/model.Data/CleanItemData',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.DeletedItemsInfo,
    responseType: gdb_pb.TimeRows,
    requestSerialize: serialize_model_DeletedItemsInfo,
    requestDeserialize: deserialize_model_DeletedItemsInfo,
    responseSerialize: serialize_model_TimeRows,
    responseDeserialize: deserialize_model_TimeRows,
  },
  reLoadDb: {
    path: '/model.Data/ReLoadDb',
    requestStream: false,
    responseStream: false,
    requestType: google_protobuf_empty_pb.Empty,
    responseType: gdb_pb.TimeRows,
    requestSerialize: serialize_google_protobuf_Empty,
    requestDeserialize: deserialize_google_protobuf_Empty,
    responseSerialize: serialize_model_TimeRows,
    responseDeserialize: deserialize_model_TimeRows,
  },
};

var PageService = exports['model.Page'] = {
  userLogin: {
    path: '/model.Page/UserLogin',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.AuthInfo,
    responseType: gdb_pb.UserToken,
    requestSerialize: serialize_model_AuthInfo,
    requestDeserialize: deserialize_model_AuthInfo,
    responseSerialize: serialize_model_UserToken,
    responseDeserialize: deserialize_model_UserToken,
  },
  userLogOut: {
    path: '/model.Page/UserLogOut',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.UserName,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_model_UserName,
    requestDeserialize: deserialize_model_UserName,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
  getUserInfo: {
    path: '/model.Page/GetUserInfo',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.UserName,
    responseType: gdb_pb.UserInfo,
    requestSerialize: serialize_model_UserName,
    requestDeserialize: deserialize_model_UserName,
    responseSerialize: serialize_model_UserInfo,
    responseDeserialize: deserialize_model_UserInfo,
  },
  getUsers: {
    path: '/model.Page/GetUsers',
    requestStream: false,
    responseStream: false,
    requestType: google_protobuf_empty_pb.Empty,
    responseType: gdb_pb.UserInfos,
    requestSerialize: serialize_google_protobuf_Empty,
    requestDeserialize: deserialize_google_protobuf_Empty,
    responseSerialize: serialize_model_UserInfos,
    responseDeserialize: deserialize_model_UserInfos,
  },
  addUsers: {
    path: '/model.Page/AddUsers',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.AddedUserInfo,
    responseType: gdb_pb.TimeRows,
    requestSerialize: serialize_model_AddedUserInfo,
    requestDeserialize: deserialize_model_AddedUserInfo,
    responseSerialize: serialize_model_TimeRows,
    responseDeserialize: deserialize_model_TimeRows,
  },
  deleteUsers: {
    path: '/model.Page/DeleteUsers',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.UserName,
    responseType: gdb_pb.TimeRows,
    requestSerialize: serialize_model_UserName,
    requestDeserialize: deserialize_model_UserName,
    responseSerialize: serialize_model_TimeRows,
    responseDeserialize: deserialize_model_TimeRows,
  },
  updateUsers: {
    path: '/model.Page/UpdateUsers',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.UpdatedUserInfo,
    responseType: gdb_pb.TimeRows,
    requestSerialize: serialize_model_UpdatedUserInfo,
    requestDeserialize: deserialize_model_UpdatedUserInfo,
    responseSerialize: serialize_model_TimeRows,
    responseDeserialize: deserialize_model_TimeRows,
  },
  uploadFile: {
    path: '/model.Page/UploadFile',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.UploadedFileInfo,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_model_UploadedFileInfo,
    requestDeserialize: deserialize_model_UploadedFileInfo,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
  uploadFileWithStream: {
    path: '/model.Page/UploadFileWithStream',
    requestStream: true,
    responseStream: false,
    requestType: gdb_pb.UploadedFileInfo,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_model_UploadedFileInfo,
    requestDeserialize: deserialize_model_UploadedFileInfo,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
  addItemsByExcel: {
    path: '/model.Page/AddItemsByExcel',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.FileInfo,
    responseType: gdb_pb.TimeRows,
    requestSerialize: serialize_model_FileInfo,
    requestDeserialize: deserialize_model_FileInfo,
    responseSerialize: serialize_model_TimeRows,
    responseDeserialize: deserialize_model_TimeRows,
  },
  importHistoryByExcel: {
    path: '/model.Page/ImportHistoryByExcel',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.HistoryFileInfo,
    responseType: gdb_pb.TimeRows,
    requestSerialize: serialize_model_HistoryFileInfo,
    requestDeserialize: deserialize_model_HistoryFileInfo,
    responseSerialize: serialize_model_TimeRows,
    responseDeserialize: deserialize_model_TimeRows,
  },
  getLogs: {
    path: '/model.Page/GetLogs',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.QueryLogsInfo,
    responseType: gdb_pb.LogsInfo,
    requestSerialize: serialize_model_QueryLogsInfo,
    requestDeserialize: deserialize_model_QueryLogsInfo,
    responseSerialize: serialize_model_LogsInfo,
    responseDeserialize: deserialize_model_LogsInfo,
  },
  getJsCode: {
    path: '/model.Page/GetJsCode',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.FileInfo,
    responseType: gdb_pb.Code,
    requestSerialize: serialize_model_FileInfo,
    requestDeserialize: deserialize_model_FileInfo,
    responseSerialize: serialize_model_Code,
    responseDeserialize: deserialize_model_Code,
  },
  deleteLogs: {
    path: '/model.Page/DeleteLogs',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.DeletedLogInfo,
    responseType: gdb_pb.TimeRows,
    requestSerialize: serialize_model_DeletedLogInfo,
    requestDeserialize: deserialize_model_DeletedLogInfo,
    responseSerialize: serialize_model_TimeRows,
    responseDeserialize: deserialize_model_TimeRows,
  },
  downloadFile: {
    path: '/model.Page/DownloadFile',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.FileInfo,
    responseType: gdb_pb.FileContents,
    requestSerialize: serialize_model_FileInfo,
    requestDeserialize: deserialize_model_FileInfo,
    responseSerialize: serialize_model_FileContents,
    responseDeserialize: deserialize_model_FileContents,
  },
  getDbSize: {
    path: '/model.Page/GetDbSize',
    requestStream: false,
    responseStream: false,
    requestType: google_protobuf_empty_pb.Empty,
    responseType: gdb_pb.FileSize,
    requestSerialize: serialize_google_protobuf_Empty,
    requestDeserialize: deserialize_google_protobuf_Empty,
    responseSerialize: serialize_model_FileSize,
    responseDeserialize: deserialize_model_FileSize,
  },
  getDbInfo: {
    path: '/model.Page/GetDbInfo',
    requestStream: false,
    responseStream: false,
    requestType: google_protobuf_empty_pb.Empty,
    responseType: gdb_pb.GdbInfoData,
    requestSerialize: serialize_google_protobuf_Empty,
    requestDeserialize: deserialize_google_protobuf_Empty,
    responseSerialize: serialize_model_GdbInfoData,
    responseDeserialize: deserialize_model_GdbInfoData,
  },
  getDbInfoHistory: {
    path: '/model.Page/GetDbInfoHistory',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.QuerySpeedHistoryDataString,
    responseType: gdb_pb.GdbHistoricalData,
    requestSerialize: serialize_model_QuerySpeedHistoryDataString,
    requestDeserialize: deserialize_model_QuerySpeedHistoryDataString,
    responseSerialize: serialize_model_GdbHistoricalData,
    responseDeserialize: deserialize_model_GdbHistoricalData,
  },
  getRoutes: {
    path: '/model.Page/GetRoutes',
    requestStream: false,
    responseStream: false,
    requestType: google_protobuf_empty_pb.Empty,
    responseType: gdb_pb.Routes,
    requestSerialize: serialize_google_protobuf_Empty,
    requestDeserialize: deserialize_google_protobuf_Empty,
    responseSerialize: serialize_model_Routes,
    responseDeserialize: deserialize_model_Routes,
  },
  deleteRoutes: {
    path: '/model.Page/DeleteRoutes',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.RoutesInfo,
    responseType: gdb_pb.TimeRows,
    requestSerialize: serialize_model_RoutesInfo,
    requestDeserialize: deserialize_model_RoutesInfo,
    responseSerialize: serialize_model_TimeRows,
    responseDeserialize: deserialize_model_TimeRows,
  },
  addRoutes: {
    path: '/model.Page/AddRoutes',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.RoutesInfo,
    responseType: gdb_pb.TimeRows,
    requestSerialize: serialize_model_RoutesInfo,
    requestDeserialize: deserialize_model_RoutesInfo,
    responseSerialize: serialize_model_TimeRows,
    responseDeserialize: deserialize_model_TimeRows,
  },
  addUserRoutes: {
    path: '/model.Page/AddUserRoutes',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.RoutesInfo,
    responseType: gdb_pb.TimeRows,
    requestSerialize: serialize_model_RoutesInfo,
    requestDeserialize: deserialize_model_RoutesInfo,
    responseSerialize: serialize_model_TimeRows,
    responseDeserialize: deserialize_model_TimeRows,
  },
  deleteUserRoutes: {
    path: '/model.Page/DeleteUserRoutes',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.UserName,
    responseType: gdb_pb.TimeRows,
    requestSerialize: serialize_model_UserName,
    requestDeserialize: deserialize_model_UserName,
    responseSerialize: serialize_model_TimeRows,
    responseDeserialize: deserialize_model_TimeRows,
  },
  getAllRoutes: {
    path: '/model.Page/GetAllRoutes',
    requestStream: false,
    responseStream: false,
    requestType: google_protobuf_empty_pb.Empty,
    responseType: gdb_pb.Routes,
    requestSerialize: serialize_google_protobuf_Empty,
    requestDeserialize: deserialize_google_protobuf_Empty,
    responseSerialize: serialize_model_Routes,
    responseDeserialize: deserialize_model_Routes,
  },
  checkRoutes: {
    path: '/model.Page/CheckRoutes',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.RoutesInfo,
    responseType: gdb_pb.CheckResult,
    requestSerialize: serialize_model_RoutesInfo,
    requestDeserialize: deserialize_model_RoutesInfo,
    responseSerialize: serialize_model_CheckResult,
    responseDeserialize: deserialize_model_CheckResult,
  },
};

var CalcService = exports['model.Calc'] = {
  testCalcItem: {
    path: '/model.Calc/TestCalcItem',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.TestCalcItemInfo,
    responseType: gdb_pb.CalculationResult,
    requestSerialize: serialize_model_TestCalcItemInfo,
    requestDeserialize: deserialize_model_TestCalcItemInfo,
    responseSerialize: serialize_model_CalculationResult,
    responseDeserialize: deserialize_model_CalculationResult,
  },
  addCalcItem: {
    path: '/model.Calc/AddCalcItem',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.AddedCalcItemInfo,
    responseType: gdb_pb.CalculationResult,
    requestSerialize: serialize_model_AddedCalcItemInfo,
    requestDeserialize: deserialize_model_AddedCalcItemInfo,
    responseSerialize: serialize_model_CalculationResult,
    responseDeserialize: deserialize_model_CalculationResult,
  },
  addCalcItemWithStream: {
    path: '/model.Calc/AddCalcItemWithStream',
    requestStream: true,
    responseStream: false,
    requestType: gdb_pb.AddedCalcItemInfo,
    responseType: gdb_pb.CalculationResults,
    requestSerialize: serialize_model_AddedCalcItemInfo,
    requestDeserialize: deserialize_model_AddedCalcItemInfo,
    responseSerialize: serialize_model_CalculationResults,
    responseDeserialize: deserialize_model_CalculationResults,
  },
  getCalcItems: {
    path: '/model.Calc/GetCalcItems',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.QueryCalcItemsInfo,
    responseType: gdb_pb.CalcItemsInfo,
    requestSerialize: serialize_model_QueryCalcItemsInfo,
    requestDeserialize: deserialize_model_QueryCalcItemsInfo,
    responseSerialize: serialize_model_CalcItemsInfo,
    responseDeserialize: deserialize_model_CalcItemsInfo,
  },
  updateCalcItem: {
    path: '/model.Calc/UpdateCalcItem',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.UpdatedCalcInfo,
    responseType: gdb_pb.CalculationResult,
    requestSerialize: serialize_model_UpdatedCalcInfo,
    requestDeserialize: deserialize_model_UpdatedCalcInfo,
    responseSerialize: serialize_model_CalculationResult,
    responseDeserialize: deserialize_model_CalculationResult,
  },
  startCalcItem: {
    path: '/model.Calc/StartCalcItem',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.CalcId,
    responseType: gdb_pb.TimeRows,
    requestSerialize: serialize_model_CalcId,
    requestDeserialize: deserialize_model_CalcId,
    responseSerialize: serialize_model_TimeRows,
    responseDeserialize: deserialize_model_TimeRows,
  },
  stopCalcItem: {
    path: '/model.Calc/StopCalcItem',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.CalcId,
    responseType: gdb_pb.TimeRows,
    requestSerialize: serialize_model_CalcId,
    requestDeserialize: deserialize_model_CalcId,
    responseSerialize: serialize_model_TimeRows,
    responseDeserialize: deserialize_model_TimeRows,
  },
  deleteCalcItem: {
    path: '/model.Calc/DeleteCalcItem',
    requestStream: false,
    responseStream: false,
    requestType: gdb_pb.CalcId,
    responseType: gdb_pb.TimeRows,
    requestSerialize: serialize_model_CalcId,
    requestDeserialize: deserialize_model_CalcId,
    responseSerialize: serialize_model_TimeRows,
    responseDeserialize: deserialize_model_TimeRows,
  },
};

