// source: gdb.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js');
goog.object.extend(proto, google_protobuf_empty_pb);
goog.exportSymbol('proto.model.AddedCalcItemInfo', null, global);
goog.exportSymbol('proto.model.AddedGroupColumnsInfo', null, global);
goog.exportSymbol('proto.model.AddedGroupInfo', null, global);
goog.exportSymbol('proto.model.AddedGroupInfos', null, global);
goog.exportSymbol('proto.model.AddedItemsInfo', null, global);
goog.exportSymbol('proto.model.AddedUserInfo', null, global);
goog.exportSymbol('proto.model.AuthInfo', null, global);
goog.exportSymbol('proto.model.BoolHItemValue', null, global);
goog.exportSymbol('proto.model.BoolHItemValues', null, global);
goog.exportSymbol('proto.model.BoolHistoricalData', null, global);
goog.exportSymbol('proto.model.BoolItemValue', null, global);
goog.exportSymbol('proto.model.BoolItemValues', null, global);
goog.exportSymbol('proto.model.CalcId', null, global);
goog.exportSymbol('proto.model.CalcItemsInfo', null, global);
goog.exportSymbol('proto.model.CalculationResult', null, global);
goog.exportSymbol('proto.model.CalculationResults', null, global);
goog.exportSymbol('proto.model.CheckItemsInfo', null, global);
goog.exportSymbol('proto.model.CheckResult', null, global);
goog.exportSymbol('proto.model.Code', null, global);
goog.exportSymbol('proto.model.DeadZone', null, global);
goog.exportSymbol('proto.model.DeleteHistoricalDataString', null, global);
goog.exportSymbol('proto.model.DeletedGroupColumnNamesInfo', null, global);
goog.exportSymbol('proto.model.DeletedItemsInfo', null, global);
goog.exportSymbol('proto.model.DeletedLogInfo', null, global);
goog.exportSymbol('proto.model.FileContents', null, global);
goog.exportSymbol('proto.model.FileInfo', null, global);
goog.exportSymbol('proto.model.FileSize', null, global);
goog.exportSymbol('proto.model.FloatHItemValue', null, global);
goog.exportSymbol('proto.model.FloatHItemValues', null, global);
goog.exportSymbol('proto.model.FloatHistoricalData', null, global);
goog.exportSymbol('proto.model.FloatItemValue', null, global);
goog.exportSymbol('proto.model.FloatItemValues', null, global);
goog.exportSymbol('proto.model.GdbHistoricalData', null, global);
goog.exportSymbol('proto.model.GdbInfoData', null, global);
goog.exportSymbol('proto.model.GdbItems', null, global);
goog.exportSymbol('proto.model.GdbItemsWithCount', null, global);
goog.exportSymbol('proto.model.GdbRealTimeData', null, global);
goog.exportSymbol('proto.model.GroupNamesInfo', null, global);
goog.exportSymbol('proto.model.GroupPropertyInfo', null, global);
goog.exportSymbol('proto.model.HistoryFileInfo', null, global);
goog.exportSymbol('proto.model.IntHItemValue', null, global);
goog.exportSymbol('proto.model.IntHItemValues', null, global);
goog.exportSymbol('proto.model.IntHistoricalData', null, global);
goog.exportSymbol('proto.model.IntItemValue', null, global);
goog.exportSymbol('proto.model.IntItemValues', null, global);
goog.exportSymbol('proto.model.ItemName', null, global);
goog.exportSymbol('proto.model.ItemsInfo', null, global);
goog.exportSymbol('proto.model.LogsInfo', null, global);
goog.exportSymbol('proto.model.QueryCalcItemsInfo', null, global);
goog.exportSymbol('proto.model.QueryGroupPropertyInfo', null, global);
goog.exportSymbol('proto.model.QueryHistoricalDataString', null, global);
goog.exportSymbol('proto.model.QueryHistoricalDataWithConditionString', null, global);
goog.exportSymbol('proto.model.QueryHistoricalDataWithStamp', null, global);
goog.exportSymbol('proto.model.QueryHistoricalDataWithStampString', null, global);
goog.exportSymbol('proto.model.QueryLogsInfo', null, global);
goog.exportSymbol('proto.model.QueryRawHistoricalDataString', null, global);
goog.exportSymbol('proto.model.QueryRealTimeDataString', null, global);
goog.exportSymbol('proto.model.QuerySpeedHistoryDataString', null, global);
goog.exportSymbol('proto.model.Routes', null, global);
goog.exportSymbol('proto.model.RoutesInfo', null, global);
goog.exportSymbol('proto.model.StringHItemValue', null, global);
goog.exportSymbol('proto.model.StringHItemValues', null, global);
goog.exportSymbol('proto.model.StringHistoricalData', null, global);
goog.exportSymbol('proto.model.StringItemValue', null, global);
goog.exportSymbol('proto.model.StringItemValues', null, global);
goog.exportSymbol('proto.model.TestCalcItemInfo', null, global);
goog.exportSymbol('proto.model.TimeCols', null, global);
goog.exportSymbol('proto.model.TimeRows', null, global);
goog.exportSymbol('proto.model.UpdatedCalcInfo', null, global);
goog.exportSymbol('proto.model.UpdatedGroupColumnNamesInfo', null, global);
goog.exportSymbol('proto.model.UpdatedGroupNameInfo', null, global);
goog.exportSymbol('proto.model.UpdatedGroupNamesInfo', null, global);
goog.exportSymbol('proto.model.UpdatedItemsInfo', null, global);
goog.exportSymbol('proto.model.UpdatedUserInfo', null, global);
goog.exportSymbol('proto.model.UploadedFileInfo', null, global);
goog.exportSymbol('proto.model.UserInfo', null, global);
goog.exportSymbol('proto.model.UserInfos', null, global);
goog.exportSymbol('proto.model.UserName', null, global);
goog.exportSymbol('proto.model.UserToken', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.TimeRows = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.model.TimeRows, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.TimeRows.displayName = 'proto.model.TimeRows';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.TimeCols = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.model.TimeCols, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.TimeCols.displayName = 'proto.model.TimeCols';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.AddedGroupInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.AddedGroupInfo.repeatedFields_, null);
};
goog.inherits(proto.model.AddedGroupInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.AddedGroupInfo.displayName = 'proto.model.AddedGroupInfo';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.AddedGroupInfos = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.AddedGroupInfos.repeatedFields_, null);
};
goog.inherits(proto.model.AddedGroupInfos, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.AddedGroupInfos.displayName = 'proto.model.AddedGroupInfos';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.GroupNamesInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.GroupNamesInfo.repeatedFields_, null);
};
goog.inherits(proto.model.GroupNamesInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.GroupNamesInfo.displayName = 'proto.model.GroupNamesInfo';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.QueryGroupPropertyInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.model.QueryGroupPropertyInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.QueryGroupPropertyInfo.displayName = 'proto.model.QueryGroupPropertyInfo';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.UpdatedGroupNameInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.model.UpdatedGroupNameInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.UpdatedGroupNameInfo.displayName = 'proto.model.UpdatedGroupNameInfo';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.UpdatedGroupNamesInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.UpdatedGroupNamesInfo.repeatedFields_, null);
};
goog.inherits(proto.model.UpdatedGroupNamesInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.UpdatedGroupNamesInfo.displayName = 'proto.model.UpdatedGroupNamesInfo';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.AddedGroupColumnsInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.AddedGroupColumnsInfo.repeatedFields_, null);
};
goog.inherits(proto.model.AddedGroupColumnsInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.AddedGroupColumnsInfo.displayName = 'proto.model.AddedGroupColumnsInfo';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.DeletedGroupColumnNamesInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.DeletedGroupColumnNamesInfo.repeatedFields_, null);
};
goog.inherits(proto.model.DeletedGroupColumnNamesInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.DeletedGroupColumnNamesInfo.displayName = 'proto.model.DeletedGroupColumnNamesInfo';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.GroupPropertyInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.GroupPropertyInfo.repeatedFields_, null);
};
goog.inherits(proto.model.GroupPropertyInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.GroupPropertyInfo.displayName = 'proto.model.GroupPropertyInfo';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.UpdatedGroupColumnNamesInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.UpdatedGroupColumnNamesInfo.repeatedFields_, null);
};
goog.inherits(proto.model.UpdatedGroupColumnNamesInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.UpdatedGroupColumnNamesInfo.displayName = 'proto.model.UpdatedGroupColumnNamesInfo';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.AddedItemsInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.model.AddedItemsInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.AddedItemsInfo.displayName = 'proto.model.AddedItemsInfo';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.DeletedItemsInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.model.DeletedItemsInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.DeletedItemsInfo.displayName = 'proto.model.DeletedItemsInfo';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.ItemsInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.model.ItemsInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.ItemsInfo.displayName = 'proto.model.ItemsInfo';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.GdbItemsWithCount = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.GdbItemsWithCount.repeatedFields_, null);
};
goog.inherits(proto.model.GdbItemsWithCount, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.GdbItemsWithCount.displayName = 'proto.model.GdbItemsWithCount';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.GdbItems = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.model.GdbItems, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.GdbItems.displayName = 'proto.model.GdbItems';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.CheckItemsInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.CheckItemsInfo.repeatedFields_, null);
};
goog.inherits(proto.model.CheckItemsInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.CheckItemsInfo.displayName = 'proto.model.CheckItemsInfo';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.UpdatedItemsInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.model.UpdatedItemsInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.UpdatedItemsInfo.displayName = 'proto.model.UpdatedItemsInfo';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.ItemName = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.ItemName.repeatedFields_, null);
};
goog.inherits(proto.model.ItemName, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.ItemName.displayName = 'proto.model.ItemName';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.FloatItemValue = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.FloatItemValue.repeatedFields_, null);
};
goog.inherits(proto.model.FloatItemValue, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.FloatItemValue.displayName = 'proto.model.FloatItemValue';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.IntItemValue = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.IntItemValue.repeatedFields_, null);
};
goog.inherits(proto.model.IntItemValue, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.IntItemValue.displayName = 'proto.model.IntItemValue';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.StringItemValue = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.StringItemValue.repeatedFields_, null);
};
goog.inherits(proto.model.StringItemValue, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.StringItemValue.displayName = 'proto.model.StringItemValue';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.BoolItemValue = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.BoolItemValue.repeatedFields_, null);
};
goog.inherits(proto.model.BoolItemValue, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.BoolItemValue.displayName = 'proto.model.BoolItemValue';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.FloatItemValues = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.FloatItemValues.repeatedFields_, null);
};
goog.inherits(proto.model.FloatItemValues, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.FloatItemValues.displayName = 'proto.model.FloatItemValues';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.IntItemValues = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.IntItemValues.repeatedFields_, null);
};
goog.inherits(proto.model.IntItemValues, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.IntItemValues.displayName = 'proto.model.IntItemValues';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.BoolItemValues = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.BoolItemValues.repeatedFields_, null);
};
goog.inherits(proto.model.BoolItemValues, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.BoolItemValues.displayName = 'proto.model.BoolItemValues';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.StringItemValues = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.StringItemValues.repeatedFields_, null);
};
goog.inherits(proto.model.StringItemValues, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.StringItemValues.displayName = 'proto.model.StringItemValues';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.FloatHItemValues = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.FloatHItemValues.repeatedFields_, null);
};
goog.inherits(proto.model.FloatHItemValues, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.FloatHItemValues.displayName = 'proto.model.FloatHItemValues';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.FloatHItemValue = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.FloatHItemValue.repeatedFields_, null);
};
goog.inherits(proto.model.FloatHItemValue, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.FloatHItemValue.displayName = 'proto.model.FloatHItemValue';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.IntHItemValues = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.IntHItemValues.repeatedFields_, null);
};
goog.inherits(proto.model.IntHItemValues, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.IntHItemValues.displayName = 'proto.model.IntHItemValues';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.IntHItemValue = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.IntHItemValue.repeatedFields_, null);
};
goog.inherits(proto.model.IntHItemValue, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.IntHItemValue.displayName = 'proto.model.IntHItemValue';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.StringHItemValues = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.StringHItemValues.repeatedFields_, null);
};
goog.inherits(proto.model.StringHItemValues, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.StringHItemValues.displayName = 'proto.model.StringHItemValues';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.StringHItemValue = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.StringHItemValue.repeatedFields_, null);
};
goog.inherits(proto.model.StringHItemValue, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.StringHItemValue.displayName = 'proto.model.StringHItemValue';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.BoolHItemValues = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.BoolHItemValues.repeatedFields_, null);
};
goog.inherits(proto.model.BoolHItemValues, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.BoolHItemValues.displayName = 'proto.model.BoolHItemValues';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.BoolHItemValue = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.BoolHItemValue.repeatedFields_, null);
};
goog.inherits(proto.model.BoolHItemValue, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.BoolHItemValue.displayName = 'proto.model.BoolHItemValue';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.QueryRealTimeDataString = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.QueryRealTimeDataString.repeatedFields_, null);
};
goog.inherits(proto.model.QueryRealTimeDataString, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.QueryRealTimeDataString.displayName = 'proto.model.QueryRealTimeDataString';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.QueryHistoricalDataString = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.QueryHistoricalDataString.repeatedFields_, null);
};
goog.inherits(proto.model.QueryHistoricalDataString, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.QueryHistoricalDataString.displayName = 'proto.model.QueryHistoricalDataString';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.QueryRawHistoricalDataString = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.QueryRawHistoricalDataString.repeatedFields_, null);
};
goog.inherits(proto.model.QueryRawHistoricalDataString, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.QueryRawHistoricalDataString.displayName = 'proto.model.QueryRawHistoricalDataString';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.QueryHistoricalDataWithStamp = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.QueryHistoricalDataWithStamp.repeatedFields_, null);
};
goog.inherits(proto.model.QueryHistoricalDataWithStamp, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.QueryHistoricalDataWithStamp.displayName = 'proto.model.QueryHistoricalDataWithStamp';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.QueryHistoricalDataWithStampString = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.QueryHistoricalDataWithStampString.repeatedFields_, null);
};
goog.inherits(proto.model.QueryHistoricalDataWithStampString, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.QueryHistoricalDataWithStampString.displayName = 'proto.model.QueryHistoricalDataWithStampString';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.DeadZone = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.model.DeadZone, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.DeadZone.displayName = 'proto.model.DeadZone';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.QueryHistoricalDataWithConditionString = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.QueryHistoricalDataWithConditionString.repeatedFields_, null);
};
goog.inherits(proto.model.QueryHistoricalDataWithConditionString, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.QueryHistoricalDataWithConditionString.displayName = 'proto.model.QueryHistoricalDataWithConditionString';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.DeleteHistoricalDataString = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.DeleteHistoricalDataString.repeatedFields_, null);
};
goog.inherits(proto.model.DeleteHistoricalDataString, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.DeleteHistoricalDataString.displayName = 'proto.model.DeleteHistoricalDataString';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.GdbRealTimeData = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.model.GdbRealTimeData, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.GdbRealTimeData.displayName = 'proto.model.GdbRealTimeData';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.GdbHistoricalData = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.model.GdbHistoricalData, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.GdbHistoricalData.displayName = 'proto.model.GdbHistoricalData';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.AuthInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.model.AuthInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.AuthInfo.displayName = 'proto.model.AuthInfo';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.UserName = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.model.UserName, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.UserName.displayName = 'proto.model.UserName';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.AddedUserInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.model.AddedUserInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.AddedUserInfo.displayName = 'proto.model.AddedUserInfo';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.UpdatedUserInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.model.UpdatedUserInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.UpdatedUserInfo.displayName = 'proto.model.UpdatedUserInfo';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.FileInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.model.FileInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.FileInfo.displayName = 'proto.model.FileInfo';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.HistoryFileInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.HistoryFileInfo.repeatedFields_, null);
};
goog.inherits(proto.model.HistoryFileInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.HistoryFileInfo.displayName = 'proto.model.HistoryFileInfo';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.QueryLogsInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.model.QueryLogsInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.QueryLogsInfo.displayName = 'proto.model.QueryLogsInfo';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.DeletedLogInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.model.DeletedLogInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.DeletedLogInfo.displayName = 'proto.model.DeletedLogInfo';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.QuerySpeedHistoryDataString = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.QuerySpeedHistoryDataString.repeatedFields_, null);
};
goog.inherits(proto.model.QuerySpeedHistoryDataString, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.QuerySpeedHistoryDataString.displayName = 'proto.model.QuerySpeedHistoryDataString';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.RoutesInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.RoutesInfo.repeatedFields_, null);
};
goog.inherits(proto.model.RoutesInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.RoutesInfo.displayName = 'proto.model.RoutesInfo';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.GdbInfoData = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.model.GdbInfoData, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.GdbInfoData.displayName = 'proto.model.GdbInfoData';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.UserToken = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.model.UserToken, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.UserToken.displayName = 'proto.model.UserToken';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.UserInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.UserInfo.repeatedFields_, null);
};
goog.inherits(proto.model.UserInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.UserInfo.displayName = 'proto.model.UserInfo';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.LogsInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.model.LogsInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.LogsInfo.displayName = 'proto.model.LogsInfo';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.UploadedFileInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.UploadedFileInfo.repeatedFields_, null);
};
goog.inherits(proto.model.UploadedFileInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.UploadedFileInfo.displayName = 'proto.model.UploadedFileInfo';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.FileContents = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.FileContents.repeatedFields_, null);
};
goog.inherits(proto.model.FileContents, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.FileContents.displayName = 'proto.model.FileContents';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.FileSize = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.model.FileSize, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.FileSize.displayName = 'proto.model.FileSize';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.UserInfos = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.model.UserInfos, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.UserInfos.displayName = 'proto.model.UserInfos';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.Routes = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.model.Routes, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.Routes.displayName = 'proto.model.Routes';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.CheckResult = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.CheckResult.repeatedFields_, null);
};
goog.inherits(proto.model.CheckResult, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.CheckResult.displayName = 'proto.model.CheckResult';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.Code = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.model.Code, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.Code.displayName = 'proto.model.Code';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.AddedCalcItemInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.model.AddedCalcItemInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.AddedCalcItemInfo.displayName = 'proto.model.AddedCalcItemInfo';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.QueryCalcItemsInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.model.QueryCalcItemsInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.QueryCalcItemsInfo.displayName = 'proto.model.QueryCalcItemsInfo';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.UpdatedCalcInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.model.UpdatedCalcInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.UpdatedCalcInfo.displayName = 'proto.model.UpdatedCalcInfo';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.CalcId = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.CalcId.repeatedFields_, null);
};
goog.inherits(proto.model.CalcId, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.CalcId.displayName = 'proto.model.CalcId';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.CalculationResult = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.model.CalculationResult, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.CalculationResult.displayName = 'proto.model.CalculationResult';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.CalcItemsInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.model.CalcItemsInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.CalcItemsInfo.displayName = 'proto.model.CalcItemsInfo';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.CalculationResults = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.CalculationResults.repeatedFields_, null);
};
goog.inherits(proto.model.CalculationResults, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.CalculationResults.displayName = 'proto.model.CalculationResults';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.TestCalcItemInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.model.TestCalcItemInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.TestCalcItemInfo.displayName = 'proto.model.TestCalcItemInfo';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.FloatHistoricalData = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.FloatHistoricalData.repeatedFields_, null);
};
goog.inherits(proto.model.FloatHistoricalData, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.FloatHistoricalData.displayName = 'proto.model.FloatHistoricalData';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.IntHistoricalData = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.IntHistoricalData.repeatedFields_, null);
};
goog.inherits(proto.model.IntHistoricalData, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.IntHistoricalData.displayName = 'proto.model.IntHistoricalData';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.StringHistoricalData = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.StringHistoricalData.repeatedFields_, null);
};
goog.inherits(proto.model.StringHistoricalData, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.StringHistoricalData.displayName = 'proto.model.StringHistoricalData';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.model.BoolHistoricalData = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.model.BoolHistoricalData.repeatedFields_, null);
};
goog.inherits(proto.model.BoolHistoricalData, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.model.BoolHistoricalData.displayName = 'proto.model.BoolHistoricalData';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.TimeRows.prototype.toObject = function(opt_includeInstance) {
  return proto.model.TimeRows.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.TimeRows} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.TimeRows.toObject = function(includeInstance, msg) {
  var f, obj = {
    effectedrows: jspb.Message.getFieldWithDefault(msg, 1, 0),
    times: jspb.Message.getFieldWithDefault(msg, 2, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.TimeRows}
 */
proto.model.TimeRows.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.TimeRows;
  return proto.model.TimeRows.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.TimeRows} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.TimeRows}
 */
proto.model.TimeRows.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setEffectedrows(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setTimes(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.TimeRows.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.TimeRows.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.TimeRows} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.TimeRows.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getEffectedrows();
  if (f !== 0) {
    writer.writeInt32(
      1,
      f
    );
  }
  f = message.getTimes();
  if (f !== 0) {
    writer.writeInt64(
      2,
      f
    );
  }
};


/**
 * optional int32 effectedRows = 1;
 * @return {number}
 */
proto.model.TimeRows.prototype.getEffectedrows = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.model.TimeRows} returns this
 */
proto.model.TimeRows.prototype.setEffectedrows = function(value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional int64 times = 2;
 * @return {number}
 */
proto.model.TimeRows.prototype.getTimes = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.model.TimeRows} returns this
 */
proto.model.TimeRows.prototype.setTimes = function(value) {
  return jspb.Message.setProto3IntField(this, 2, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.TimeCols.prototype.toObject = function(opt_includeInstance) {
  return proto.model.TimeCols.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.TimeCols} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.TimeCols.toObject = function(includeInstance, msg) {
  var f, obj = {
    effectedcols: jspb.Message.getFieldWithDefault(msg, 1, 0),
    times: jspb.Message.getFieldWithDefault(msg, 2, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.TimeCols}
 */
proto.model.TimeCols.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.TimeCols;
  return proto.model.TimeCols.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.TimeCols} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.TimeCols}
 */
proto.model.TimeCols.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setEffectedcols(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setTimes(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.TimeCols.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.TimeCols.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.TimeCols} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.TimeCols.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getEffectedcols();
  if (f !== 0) {
    writer.writeInt32(
      1,
      f
    );
  }
  f = message.getTimes();
  if (f !== 0) {
    writer.writeInt64(
      2,
      f
    );
  }
};


/**
 * optional int32 effectedCols = 1;
 * @return {number}
 */
proto.model.TimeCols.prototype.getEffectedcols = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.model.TimeCols} returns this
 */
proto.model.TimeCols.prototype.setEffectedcols = function(value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional int64 times = 2;
 * @return {number}
 */
proto.model.TimeCols.prototype.getTimes = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.model.TimeCols} returns this
 */
proto.model.TimeCols.prototype.setTimes = function(value) {
  return jspb.Message.setProto3IntField(this, 2, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.AddedGroupInfo.repeatedFields_ = [2];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.AddedGroupInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.model.AddedGroupInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.AddedGroupInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.AddedGroupInfo.toObject = function(includeInstance, msg) {
  var f, obj = {
    groupname: jspb.Message.getFieldWithDefault(msg, 1, ""),
    columnnamesList: (f = jspb.Message.getRepeatedField(msg, 2)) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.AddedGroupInfo}
 */
proto.model.AddedGroupInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.AddedGroupInfo;
  return proto.model.AddedGroupInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.AddedGroupInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.AddedGroupInfo}
 */
proto.model.AddedGroupInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setGroupname(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.addColumnnames(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.AddedGroupInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.AddedGroupInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.AddedGroupInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.AddedGroupInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getGroupname();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getColumnnamesList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      2,
      f
    );
  }
};


/**
 * optional string groupName = 1;
 * @return {string}
 */
proto.model.AddedGroupInfo.prototype.getGroupname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.AddedGroupInfo} returns this
 */
proto.model.AddedGroupInfo.prototype.setGroupname = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * repeated string columnNames = 2;
 * @return {!Array<string>}
 */
proto.model.AddedGroupInfo.prototype.getColumnnamesList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 2));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.model.AddedGroupInfo} returns this
 */
proto.model.AddedGroupInfo.prototype.setColumnnamesList = function(value) {
  return jspb.Message.setField(this, 2, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.model.AddedGroupInfo} returns this
 */
proto.model.AddedGroupInfo.prototype.addColumnnames = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 2, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.AddedGroupInfo} returns this
 */
proto.model.AddedGroupInfo.prototype.clearColumnnamesList = function() {
  return this.setColumnnamesList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.AddedGroupInfos.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.AddedGroupInfos.prototype.toObject = function(opt_includeInstance) {
  return proto.model.AddedGroupInfos.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.AddedGroupInfos} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.AddedGroupInfos.toObject = function(includeInstance, msg) {
  var f, obj = {
    groupinfosList: jspb.Message.toObjectList(msg.getGroupinfosList(),
    proto.model.AddedGroupInfo.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.AddedGroupInfos}
 */
proto.model.AddedGroupInfos.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.AddedGroupInfos;
  return proto.model.AddedGroupInfos.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.AddedGroupInfos} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.AddedGroupInfos}
 */
proto.model.AddedGroupInfos.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.model.AddedGroupInfo;
      reader.readMessage(value,proto.model.AddedGroupInfo.deserializeBinaryFromReader);
      msg.addGroupinfos(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.AddedGroupInfos.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.AddedGroupInfos.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.AddedGroupInfos} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.AddedGroupInfos.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getGroupinfosList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.model.AddedGroupInfo.serializeBinaryToWriter
    );
  }
};


/**
 * repeated AddedGroupInfo groupInfos = 1;
 * @return {!Array<!proto.model.AddedGroupInfo>}
 */
proto.model.AddedGroupInfos.prototype.getGroupinfosList = function() {
  return /** @type{!Array<!proto.model.AddedGroupInfo>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.model.AddedGroupInfo, 1));
};


/**
 * @param {!Array<!proto.model.AddedGroupInfo>} value
 * @return {!proto.model.AddedGroupInfos} returns this
*/
proto.model.AddedGroupInfos.prototype.setGroupinfosList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.model.AddedGroupInfo=} opt_value
 * @param {number=} opt_index
 * @return {!proto.model.AddedGroupInfo}
 */
proto.model.AddedGroupInfos.prototype.addGroupinfos = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.model.AddedGroupInfo, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.AddedGroupInfos} returns this
 */
proto.model.AddedGroupInfos.prototype.clearGroupinfosList = function() {
  return this.setGroupinfosList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.GroupNamesInfo.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.GroupNamesInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.model.GroupNamesInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.GroupNamesInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.GroupNamesInfo.toObject = function(includeInstance, msg) {
  var f, obj = {
    groupnamesList: (f = jspb.Message.getRepeatedField(msg, 1)) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.GroupNamesInfo}
 */
proto.model.GroupNamesInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.GroupNamesInfo;
  return proto.model.GroupNamesInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.GroupNamesInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.GroupNamesInfo}
 */
proto.model.GroupNamesInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.addGroupnames(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.GroupNamesInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.GroupNamesInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.GroupNamesInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.GroupNamesInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getGroupnamesList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      1,
      f
    );
  }
};


/**
 * repeated string groupNames = 1;
 * @return {!Array<string>}
 */
proto.model.GroupNamesInfo.prototype.getGroupnamesList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 1));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.model.GroupNamesInfo} returns this
 */
proto.model.GroupNamesInfo.prototype.setGroupnamesList = function(value) {
  return jspb.Message.setField(this, 1, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.model.GroupNamesInfo} returns this
 */
proto.model.GroupNamesInfo.prototype.addGroupnames = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 1, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.GroupNamesInfo} returns this
 */
proto.model.GroupNamesInfo.prototype.clearGroupnamesList = function() {
  return this.setGroupnamesList([]);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.QueryGroupPropertyInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.model.QueryGroupPropertyInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.QueryGroupPropertyInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.QueryGroupPropertyInfo.toObject = function(includeInstance, msg) {
  var f, obj = {
    groupname: jspb.Message.getFieldWithDefault(msg, 1, ""),
    condition: jspb.Message.getFieldWithDefault(msg, 2, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.QueryGroupPropertyInfo}
 */
proto.model.QueryGroupPropertyInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.QueryGroupPropertyInfo;
  return proto.model.QueryGroupPropertyInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.QueryGroupPropertyInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.QueryGroupPropertyInfo}
 */
proto.model.QueryGroupPropertyInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setGroupname(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setCondition(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.QueryGroupPropertyInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.QueryGroupPropertyInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.QueryGroupPropertyInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.QueryGroupPropertyInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getGroupname();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getCondition();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * optional string groupName = 1;
 * @return {string}
 */
proto.model.QueryGroupPropertyInfo.prototype.getGroupname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.QueryGroupPropertyInfo} returns this
 */
proto.model.QueryGroupPropertyInfo.prototype.setGroupname = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string condition = 2;
 * @return {string}
 */
proto.model.QueryGroupPropertyInfo.prototype.getCondition = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.QueryGroupPropertyInfo} returns this
 */
proto.model.QueryGroupPropertyInfo.prototype.setCondition = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.UpdatedGroupNameInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.model.UpdatedGroupNameInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.UpdatedGroupNameInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.UpdatedGroupNameInfo.toObject = function(includeInstance, msg) {
  var f, obj = {
    oldgroupname: jspb.Message.getFieldWithDefault(msg, 1, ""),
    newgroupname: jspb.Message.getFieldWithDefault(msg, 2, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.UpdatedGroupNameInfo}
 */
proto.model.UpdatedGroupNameInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.UpdatedGroupNameInfo;
  return proto.model.UpdatedGroupNameInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.UpdatedGroupNameInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.UpdatedGroupNameInfo}
 */
proto.model.UpdatedGroupNameInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setOldgroupname(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setNewgroupname(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.UpdatedGroupNameInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.UpdatedGroupNameInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.UpdatedGroupNameInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.UpdatedGroupNameInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getOldgroupname();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getNewgroupname();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * optional string oldGroupName = 1;
 * @return {string}
 */
proto.model.UpdatedGroupNameInfo.prototype.getOldgroupname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.UpdatedGroupNameInfo} returns this
 */
proto.model.UpdatedGroupNameInfo.prototype.setOldgroupname = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string newGroupName = 2;
 * @return {string}
 */
proto.model.UpdatedGroupNameInfo.prototype.getNewgroupname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.UpdatedGroupNameInfo} returns this
 */
proto.model.UpdatedGroupNameInfo.prototype.setNewgroupname = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.UpdatedGroupNamesInfo.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.UpdatedGroupNamesInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.model.UpdatedGroupNamesInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.UpdatedGroupNamesInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.UpdatedGroupNamesInfo.toObject = function(includeInstance, msg) {
  var f, obj = {
    infosList: jspb.Message.toObjectList(msg.getInfosList(),
    proto.model.UpdatedGroupNameInfo.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.UpdatedGroupNamesInfo}
 */
proto.model.UpdatedGroupNamesInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.UpdatedGroupNamesInfo;
  return proto.model.UpdatedGroupNamesInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.UpdatedGroupNamesInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.UpdatedGroupNamesInfo}
 */
proto.model.UpdatedGroupNamesInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.model.UpdatedGroupNameInfo;
      reader.readMessage(value,proto.model.UpdatedGroupNameInfo.deserializeBinaryFromReader);
      msg.addInfos(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.UpdatedGroupNamesInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.UpdatedGroupNamesInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.UpdatedGroupNamesInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.UpdatedGroupNamesInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getInfosList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.model.UpdatedGroupNameInfo.serializeBinaryToWriter
    );
  }
};


/**
 * repeated UpdatedGroupNameInfo infos = 1;
 * @return {!Array<!proto.model.UpdatedGroupNameInfo>}
 */
proto.model.UpdatedGroupNamesInfo.prototype.getInfosList = function() {
  return /** @type{!Array<!proto.model.UpdatedGroupNameInfo>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.model.UpdatedGroupNameInfo, 1));
};


/**
 * @param {!Array<!proto.model.UpdatedGroupNameInfo>} value
 * @return {!proto.model.UpdatedGroupNamesInfo} returns this
*/
proto.model.UpdatedGroupNamesInfo.prototype.setInfosList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.model.UpdatedGroupNameInfo=} opt_value
 * @param {number=} opt_index
 * @return {!proto.model.UpdatedGroupNameInfo}
 */
proto.model.UpdatedGroupNamesInfo.prototype.addInfos = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.model.UpdatedGroupNameInfo, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.UpdatedGroupNamesInfo} returns this
 */
proto.model.UpdatedGroupNamesInfo.prototype.clearInfosList = function() {
  return this.setInfosList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.AddedGroupColumnsInfo.repeatedFields_ = [2,3];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.AddedGroupColumnsInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.model.AddedGroupColumnsInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.AddedGroupColumnsInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.AddedGroupColumnsInfo.toObject = function(includeInstance, msg) {
  var f, obj = {
    groupname: jspb.Message.getFieldWithDefault(msg, 1, ""),
    columnnamesList: (f = jspb.Message.getRepeatedField(msg, 2)) == null ? undefined : f,
    defaultvaluesList: (f = jspb.Message.getRepeatedField(msg, 3)) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.AddedGroupColumnsInfo}
 */
proto.model.AddedGroupColumnsInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.AddedGroupColumnsInfo;
  return proto.model.AddedGroupColumnsInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.AddedGroupColumnsInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.AddedGroupColumnsInfo}
 */
proto.model.AddedGroupColumnsInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setGroupname(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.addColumnnames(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.addDefaultvalues(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.AddedGroupColumnsInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.AddedGroupColumnsInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.AddedGroupColumnsInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.AddedGroupColumnsInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getGroupname();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getColumnnamesList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      2,
      f
    );
  }
  f = message.getDefaultvaluesList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      3,
      f
    );
  }
};


/**
 * optional string groupName = 1;
 * @return {string}
 */
proto.model.AddedGroupColumnsInfo.prototype.getGroupname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.AddedGroupColumnsInfo} returns this
 */
proto.model.AddedGroupColumnsInfo.prototype.setGroupname = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * repeated string columnNames = 2;
 * @return {!Array<string>}
 */
proto.model.AddedGroupColumnsInfo.prototype.getColumnnamesList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 2));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.model.AddedGroupColumnsInfo} returns this
 */
proto.model.AddedGroupColumnsInfo.prototype.setColumnnamesList = function(value) {
  return jspb.Message.setField(this, 2, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.model.AddedGroupColumnsInfo} returns this
 */
proto.model.AddedGroupColumnsInfo.prototype.addColumnnames = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 2, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.AddedGroupColumnsInfo} returns this
 */
proto.model.AddedGroupColumnsInfo.prototype.clearColumnnamesList = function() {
  return this.setColumnnamesList([]);
};


/**
 * repeated string defaultValues = 3;
 * @return {!Array<string>}
 */
proto.model.AddedGroupColumnsInfo.prototype.getDefaultvaluesList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 3));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.model.AddedGroupColumnsInfo} returns this
 */
proto.model.AddedGroupColumnsInfo.prototype.setDefaultvaluesList = function(value) {
  return jspb.Message.setField(this, 3, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.model.AddedGroupColumnsInfo} returns this
 */
proto.model.AddedGroupColumnsInfo.prototype.addDefaultvalues = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 3, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.AddedGroupColumnsInfo} returns this
 */
proto.model.AddedGroupColumnsInfo.prototype.clearDefaultvaluesList = function() {
  return this.setDefaultvaluesList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.DeletedGroupColumnNamesInfo.repeatedFields_ = [2];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.DeletedGroupColumnNamesInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.model.DeletedGroupColumnNamesInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.DeletedGroupColumnNamesInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.DeletedGroupColumnNamesInfo.toObject = function(includeInstance, msg) {
  var f, obj = {
    groupname: jspb.Message.getFieldWithDefault(msg, 1, ""),
    columnnamesList: (f = jspb.Message.getRepeatedField(msg, 2)) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.DeletedGroupColumnNamesInfo}
 */
proto.model.DeletedGroupColumnNamesInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.DeletedGroupColumnNamesInfo;
  return proto.model.DeletedGroupColumnNamesInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.DeletedGroupColumnNamesInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.DeletedGroupColumnNamesInfo}
 */
proto.model.DeletedGroupColumnNamesInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setGroupname(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.addColumnnames(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.DeletedGroupColumnNamesInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.DeletedGroupColumnNamesInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.DeletedGroupColumnNamesInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.DeletedGroupColumnNamesInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getGroupname();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getColumnnamesList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      2,
      f
    );
  }
};


/**
 * optional string groupName = 1;
 * @return {string}
 */
proto.model.DeletedGroupColumnNamesInfo.prototype.getGroupname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.DeletedGroupColumnNamesInfo} returns this
 */
proto.model.DeletedGroupColumnNamesInfo.prototype.setGroupname = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * repeated string columnNames = 2;
 * @return {!Array<string>}
 */
proto.model.DeletedGroupColumnNamesInfo.prototype.getColumnnamesList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 2));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.model.DeletedGroupColumnNamesInfo} returns this
 */
proto.model.DeletedGroupColumnNamesInfo.prototype.setColumnnamesList = function(value) {
  return jspb.Message.setField(this, 2, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.model.DeletedGroupColumnNamesInfo} returns this
 */
proto.model.DeletedGroupColumnNamesInfo.prototype.addColumnnames = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 2, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.DeletedGroupColumnNamesInfo} returns this
 */
proto.model.DeletedGroupColumnNamesInfo.prototype.clearColumnnamesList = function() {
  return this.setColumnnamesList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.GroupPropertyInfo.repeatedFields_ = [2];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.GroupPropertyInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.model.GroupPropertyInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.GroupPropertyInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.GroupPropertyInfo.toObject = function(includeInstance, msg) {
  var f, obj = {
    itemcount: jspb.Message.getFieldWithDefault(msg, 1, ""),
    itemcolumnnamesList: (f = jspb.Message.getRepeatedField(msg, 2)) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.GroupPropertyInfo}
 */
proto.model.GroupPropertyInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.GroupPropertyInfo;
  return proto.model.GroupPropertyInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.GroupPropertyInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.GroupPropertyInfo}
 */
proto.model.GroupPropertyInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setItemcount(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.addItemcolumnnames(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.GroupPropertyInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.GroupPropertyInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.GroupPropertyInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.GroupPropertyInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getItemcount();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getItemcolumnnamesList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      2,
      f
    );
  }
};


/**
 * optional string itemCount = 1;
 * @return {string}
 */
proto.model.GroupPropertyInfo.prototype.getItemcount = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.GroupPropertyInfo} returns this
 */
proto.model.GroupPropertyInfo.prototype.setItemcount = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * repeated string itemColumnNames = 2;
 * @return {!Array<string>}
 */
proto.model.GroupPropertyInfo.prototype.getItemcolumnnamesList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 2));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.model.GroupPropertyInfo} returns this
 */
proto.model.GroupPropertyInfo.prototype.setItemcolumnnamesList = function(value) {
  return jspb.Message.setField(this, 2, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.model.GroupPropertyInfo} returns this
 */
proto.model.GroupPropertyInfo.prototype.addItemcolumnnames = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 2, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.GroupPropertyInfo} returns this
 */
proto.model.GroupPropertyInfo.prototype.clearItemcolumnnamesList = function() {
  return this.setItemcolumnnamesList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.UpdatedGroupColumnNamesInfo.repeatedFields_ = [2,3];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.UpdatedGroupColumnNamesInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.model.UpdatedGroupColumnNamesInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.UpdatedGroupColumnNamesInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.UpdatedGroupColumnNamesInfo.toObject = function(includeInstance, msg) {
  var f, obj = {
    groupname: jspb.Message.getFieldWithDefault(msg, 1, ""),
    oldcolumnnamesList: (f = jspb.Message.getRepeatedField(msg, 2)) == null ? undefined : f,
    newcolumnnamesList: (f = jspb.Message.getRepeatedField(msg, 3)) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.UpdatedGroupColumnNamesInfo}
 */
proto.model.UpdatedGroupColumnNamesInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.UpdatedGroupColumnNamesInfo;
  return proto.model.UpdatedGroupColumnNamesInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.UpdatedGroupColumnNamesInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.UpdatedGroupColumnNamesInfo}
 */
proto.model.UpdatedGroupColumnNamesInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setGroupname(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.addOldcolumnnames(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.addNewcolumnnames(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.UpdatedGroupColumnNamesInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.UpdatedGroupColumnNamesInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.UpdatedGroupColumnNamesInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.UpdatedGroupColumnNamesInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getGroupname();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getOldcolumnnamesList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      2,
      f
    );
  }
  f = message.getNewcolumnnamesList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      3,
      f
    );
  }
};


/**
 * optional string groupName = 1;
 * @return {string}
 */
proto.model.UpdatedGroupColumnNamesInfo.prototype.getGroupname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.UpdatedGroupColumnNamesInfo} returns this
 */
proto.model.UpdatedGroupColumnNamesInfo.prototype.setGroupname = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * repeated string oldColumnNames = 2;
 * @return {!Array<string>}
 */
proto.model.UpdatedGroupColumnNamesInfo.prototype.getOldcolumnnamesList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 2));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.model.UpdatedGroupColumnNamesInfo} returns this
 */
proto.model.UpdatedGroupColumnNamesInfo.prototype.setOldcolumnnamesList = function(value) {
  return jspb.Message.setField(this, 2, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.model.UpdatedGroupColumnNamesInfo} returns this
 */
proto.model.UpdatedGroupColumnNamesInfo.prototype.addOldcolumnnames = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 2, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.UpdatedGroupColumnNamesInfo} returns this
 */
proto.model.UpdatedGroupColumnNamesInfo.prototype.clearOldcolumnnamesList = function() {
  return this.setOldcolumnnamesList([]);
};


/**
 * repeated string newColumnNames = 3;
 * @return {!Array<string>}
 */
proto.model.UpdatedGroupColumnNamesInfo.prototype.getNewcolumnnamesList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 3));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.model.UpdatedGroupColumnNamesInfo} returns this
 */
proto.model.UpdatedGroupColumnNamesInfo.prototype.setNewcolumnnamesList = function(value) {
  return jspb.Message.setField(this, 3, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.model.UpdatedGroupColumnNamesInfo} returns this
 */
proto.model.UpdatedGroupColumnNamesInfo.prototype.addNewcolumnnames = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 3, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.UpdatedGroupColumnNamesInfo} returns this
 */
proto.model.UpdatedGroupColumnNamesInfo.prototype.clearNewcolumnnamesList = function() {
  return this.setNewcolumnnamesList([]);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.AddedItemsInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.model.AddedItemsInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.AddedItemsInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.AddedItemsInfo.toObject = function(includeInstance, msg) {
  var f, obj = {
    groupname: jspb.Message.getFieldWithDefault(msg, 1, ""),
    itemvalues: jspb.Message.getFieldWithDefault(msg, 2, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.AddedItemsInfo}
 */
proto.model.AddedItemsInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.AddedItemsInfo;
  return proto.model.AddedItemsInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.AddedItemsInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.AddedItemsInfo}
 */
proto.model.AddedItemsInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setGroupname(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setItemvalues(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.AddedItemsInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.AddedItemsInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.AddedItemsInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.AddedItemsInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getGroupname();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getItemvalues();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * optional string groupName = 1;
 * @return {string}
 */
proto.model.AddedItemsInfo.prototype.getGroupname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.AddedItemsInfo} returns this
 */
proto.model.AddedItemsInfo.prototype.setGroupname = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string itemValues = 2;
 * @return {string}
 */
proto.model.AddedItemsInfo.prototype.getItemvalues = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.AddedItemsInfo} returns this
 */
proto.model.AddedItemsInfo.prototype.setItemvalues = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.DeletedItemsInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.model.DeletedItemsInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.DeletedItemsInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.DeletedItemsInfo.toObject = function(includeInstance, msg) {
  var f, obj = {
    groupname: jspb.Message.getFieldWithDefault(msg, 1, ""),
    condition: jspb.Message.getFieldWithDefault(msg, 2, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.DeletedItemsInfo}
 */
proto.model.DeletedItemsInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.DeletedItemsInfo;
  return proto.model.DeletedItemsInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.DeletedItemsInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.DeletedItemsInfo}
 */
proto.model.DeletedItemsInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setGroupname(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setCondition(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.DeletedItemsInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.DeletedItemsInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.DeletedItemsInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.DeletedItemsInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getGroupname();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getCondition();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * optional string groupName = 1;
 * @return {string}
 */
proto.model.DeletedItemsInfo.prototype.getGroupname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.DeletedItemsInfo} returns this
 */
proto.model.DeletedItemsInfo.prototype.setGroupname = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string condition = 2;
 * @return {string}
 */
proto.model.DeletedItemsInfo.prototype.getCondition = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.DeletedItemsInfo} returns this
 */
proto.model.DeletedItemsInfo.prototype.setCondition = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.ItemsInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.model.ItemsInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.ItemsInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.ItemsInfo.toObject = function(includeInstance, msg) {
  var f, obj = {
    groupname: jspb.Message.getFieldWithDefault(msg, 1, ""),
    condition: jspb.Message.getFieldWithDefault(msg, 5, ""),
    columnnames: jspb.Message.getFieldWithDefault(msg, 2, ""),
    startrow: jspb.Message.getFieldWithDefault(msg, 3, 0),
    rowcount: jspb.Message.getFieldWithDefault(msg, 4, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.ItemsInfo}
 */
proto.model.ItemsInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.ItemsInfo;
  return proto.model.ItemsInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.ItemsInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.ItemsInfo}
 */
proto.model.ItemsInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setGroupname(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setCondition(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setColumnnames(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setStartrow(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setRowcount(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.ItemsInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.ItemsInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.ItemsInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.ItemsInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getGroupname();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getCondition();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getColumnnames();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getStartrow();
  if (f !== 0) {
    writer.writeInt32(
      3,
      f
    );
  }
  f = message.getRowcount();
  if (f !== 0) {
    writer.writeInt32(
      4,
      f
    );
  }
};


/**
 * optional string groupName = 1;
 * @return {string}
 */
proto.model.ItemsInfo.prototype.getGroupname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.ItemsInfo} returns this
 */
proto.model.ItemsInfo.prototype.setGroupname = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string condition = 5;
 * @return {string}
 */
proto.model.ItemsInfo.prototype.getCondition = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.ItemsInfo} returns this
 */
proto.model.ItemsInfo.prototype.setCondition = function(value) {
  return jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional string columnNames = 2;
 * @return {string}
 */
proto.model.ItemsInfo.prototype.getColumnnames = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.ItemsInfo} returns this
 */
proto.model.ItemsInfo.prototype.setColumnnames = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional int32 startRow = 3;
 * @return {number}
 */
proto.model.ItemsInfo.prototype.getStartrow = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.model.ItemsInfo} returns this
 */
proto.model.ItemsInfo.prototype.setStartrow = function(value) {
  return jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional int32 rowCount = 4;
 * @return {number}
 */
proto.model.ItemsInfo.prototype.getRowcount = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/**
 * @param {number} value
 * @return {!proto.model.ItemsInfo} returns this
 */
proto.model.ItemsInfo.prototype.setRowcount = function(value) {
  return jspb.Message.setProto3IntField(this, 4, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.GdbItemsWithCount.repeatedFields_ = [2];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.GdbItemsWithCount.prototype.toObject = function(opt_includeInstance) {
  return proto.model.GdbItemsWithCount.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.GdbItemsWithCount} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.GdbItemsWithCount.toObject = function(includeInstance, msg) {
  var f, obj = {
    itemcount: jspb.Message.getFieldWithDefault(msg, 1, 0),
    itemvaluesList: jspb.Message.toObjectList(msg.getItemvaluesList(),
    proto.model.GdbItems.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.GdbItemsWithCount}
 */
proto.model.GdbItemsWithCount.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.GdbItemsWithCount;
  return proto.model.GdbItemsWithCount.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.GdbItemsWithCount} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.GdbItemsWithCount}
 */
proto.model.GdbItemsWithCount.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setItemcount(value);
      break;
    case 2:
      var value = new proto.model.GdbItems;
      reader.readMessage(value,proto.model.GdbItems.deserializeBinaryFromReader);
      msg.addItemvalues(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.GdbItemsWithCount.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.GdbItemsWithCount.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.GdbItemsWithCount} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.GdbItemsWithCount.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getItemcount();
  if (f !== 0) {
    writer.writeInt32(
      1,
      f
    );
  }
  f = message.getItemvaluesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      2,
      f,
      proto.model.GdbItems.serializeBinaryToWriter
    );
  }
};


/**
 * optional int32 itemCount = 1;
 * @return {number}
 */
proto.model.GdbItemsWithCount.prototype.getItemcount = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.model.GdbItemsWithCount} returns this
 */
proto.model.GdbItemsWithCount.prototype.setItemcount = function(value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * repeated GdbItems itemValues = 2;
 * @return {!Array<!proto.model.GdbItems>}
 */
proto.model.GdbItemsWithCount.prototype.getItemvaluesList = function() {
  return /** @type{!Array<!proto.model.GdbItems>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.model.GdbItems, 2));
};


/**
 * @param {!Array<!proto.model.GdbItems>} value
 * @return {!proto.model.GdbItemsWithCount} returns this
*/
proto.model.GdbItemsWithCount.prototype.setItemvaluesList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 2, value);
};


/**
 * @param {!proto.model.GdbItems=} opt_value
 * @param {number=} opt_index
 * @return {!proto.model.GdbItems}
 */
proto.model.GdbItemsWithCount.prototype.addItemvalues = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 2, opt_value, proto.model.GdbItems, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.GdbItemsWithCount} returns this
 */
proto.model.GdbItemsWithCount.prototype.clearItemvaluesList = function() {
  return this.setItemvaluesList([]);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.GdbItems.prototype.toObject = function(opt_includeInstance) {
  return proto.model.GdbItems.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.GdbItems} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.GdbItems.toObject = function(includeInstance, msg) {
  var f, obj = {
    itemsMap: (f = msg.getItemsMap()) ? f.toObject(includeInstance, undefined) : []
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.GdbItems}
 */
proto.model.GdbItems.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.GdbItems;
  return proto.model.GdbItems.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.GdbItems} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.GdbItems}
 */
proto.model.GdbItems.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = msg.getItemsMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readString, null, "", "");
         });
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.GdbItems.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.GdbItems.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.GdbItems} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.GdbItems.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getItemsMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(1, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeString);
  }
};


/**
 * map<string, string> items = 1;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,string>}
 */
proto.model.GdbItems.prototype.getItemsMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,string>} */ (
      jspb.Message.getMapField(this, 1, opt_noLazyCreate,
      null));
};


/**
 * Clears values from the map. The map will be non-null.
 * @return {!proto.model.GdbItems} returns this
 */
proto.model.GdbItems.prototype.clearItemsMap = function() {
  this.getItemsMap().clear();
  return this;};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.CheckItemsInfo.repeatedFields_ = [2];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.CheckItemsInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.model.CheckItemsInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.CheckItemsInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.CheckItemsInfo.toObject = function(includeInstance, msg) {
  var f, obj = {
    groupname: jspb.Message.getFieldWithDefault(msg, 1, ""),
    itemnamesList: (f = jspb.Message.getRepeatedField(msg, 2)) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.CheckItemsInfo}
 */
proto.model.CheckItemsInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.CheckItemsInfo;
  return proto.model.CheckItemsInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.CheckItemsInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.CheckItemsInfo}
 */
proto.model.CheckItemsInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setGroupname(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.addItemnames(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.CheckItemsInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.CheckItemsInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.CheckItemsInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.CheckItemsInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getGroupname();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getItemnamesList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      2,
      f
    );
  }
};


/**
 * optional string groupName = 1;
 * @return {string}
 */
proto.model.CheckItemsInfo.prototype.getGroupname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.CheckItemsInfo} returns this
 */
proto.model.CheckItemsInfo.prototype.setGroupname = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * repeated string itemNames = 2;
 * @return {!Array<string>}
 */
proto.model.CheckItemsInfo.prototype.getItemnamesList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 2));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.model.CheckItemsInfo} returns this
 */
proto.model.CheckItemsInfo.prototype.setItemnamesList = function(value) {
  return jspb.Message.setField(this, 2, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.model.CheckItemsInfo} returns this
 */
proto.model.CheckItemsInfo.prototype.addItemnames = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 2, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.CheckItemsInfo} returns this
 */
proto.model.CheckItemsInfo.prototype.clearItemnamesList = function() {
  return this.setItemnamesList([]);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.UpdatedItemsInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.model.UpdatedItemsInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.UpdatedItemsInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.UpdatedItemsInfo.toObject = function(includeInstance, msg) {
  var f, obj = {
    groupname: jspb.Message.getFieldWithDefault(msg, 1, ""),
    condition: jspb.Message.getFieldWithDefault(msg, 2, ""),
    clause: jspb.Message.getFieldWithDefault(msg, 3, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.UpdatedItemsInfo}
 */
proto.model.UpdatedItemsInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.UpdatedItemsInfo;
  return proto.model.UpdatedItemsInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.UpdatedItemsInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.UpdatedItemsInfo}
 */
proto.model.UpdatedItemsInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setGroupname(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setCondition(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setClause(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.UpdatedItemsInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.UpdatedItemsInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.UpdatedItemsInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.UpdatedItemsInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getGroupname();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getCondition();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getClause();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
};


/**
 * optional string groupName = 1;
 * @return {string}
 */
proto.model.UpdatedItemsInfo.prototype.getGroupname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.UpdatedItemsInfo} returns this
 */
proto.model.UpdatedItemsInfo.prototype.setGroupname = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string condition = 2;
 * @return {string}
 */
proto.model.UpdatedItemsInfo.prototype.getCondition = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.UpdatedItemsInfo} returns this
 */
proto.model.UpdatedItemsInfo.prototype.setCondition = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string clause = 3;
 * @return {string}
 */
proto.model.UpdatedItemsInfo.prototype.getClause = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.UpdatedItemsInfo} returns this
 */
proto.model.UpdatedItemsInfo.prototype.setClause = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.ItemName.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.ItemName.prototype.toObject = function(opt_includeInstance) {
  return proto.model.ItemName.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.ItemName} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.ItemName.toObject = function(includeInstance, msg) {
  var f, obj = {
    itemnameList: (f = jspb.Message.getRepeatedField(msg, 1)) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.ItemName}
 */
proto.model.ItemName.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.ItemName;
  return proto.model.ItemName.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.ItemName} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.ItemName}
 */
proto.model.ItemName.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.addItemname(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.ItemName.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.ItemName.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.ItemName} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.ItemName.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getItemnameList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      1,
      f
    );
  }
};


/**
 * repeated string itemName = 1;
 * @return {!Array<string>}
 */
proto.model.ItemName.prototype.getItemnameList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 1));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.model.ItemName} returns this
 */
proto.model.ItemName.prototype.setItemnameList = function(value) {
  return jspb.Message.setField(this, 1, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.model.ItemName} returns this
 */
proto.model.ItemName.prototype.addItemname = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 1, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.ItemName} returns this
 */
proto.model.ItemName.prototype.clearItemnameList = function() {
  return this.setItemnameList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.FloatItemValue.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.FloatItemValue.prototype.toObject = function(opt_includeInstance) {
  return proto.model.FloatItemValue.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.FloatItemValue} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.FloatItemValue.toObject = function(includeInstance, msg) {
  var f, obj = {
    itemvalueList: (f = jspb.Message.getRepeatedFloatingPointField(msg, 1)) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.FloatItemValue}
 */
proto.model.FloatItemValue.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.FloatItemValue;
  return proto.model.FloatItemValue.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.FloatItemValue} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.FloatItemValue}
 */
proto.model.FloatItemValue.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedFloat() : [reader.readFloat()]);
      for (var i = 0; i < values.length; i++) {
        msg.addItemvalue(values[i]);
      }
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.FloatItemValue.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.FloatItemValue.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.FloatItemValue} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.FloatItemValue.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getItemvalueList();
  if (f.length > 0) {
    writer.writePackedFloat(
      1,
      f
    );
  }
};


/**
 * repeated float itemValue = 1;
 * @return {!Array<number>}
 */
proto.model.FloatItemValue.prototype.getItemvalueList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedFloatingPointField(this, 1));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.model.FloatItemValue} returns this
 */
proto.model.FloatItemValue.prototype.setItemvalueList = function(value) {
  return jspb.Message.setField(this, 1, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.model.FloatItemValue} returns this
 */
proto.model.FloatItemValue.prototype.addItemvalue = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 1, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.FloatItemValue} returns this
 */
proto.model.FloatItemValue.prototype.clearItemvalueList = function() {
  return this.setItemvalueList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.IntItemValue.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.IntItemValue.prototype.toObject = function(opt_includeInstance) {
  return proto.model.IntItemValue.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.IntItemValue} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.IntItemValue.toObject = function(includeInstance, msg) {
  var f, obj = {
    itemvalueList: (f = jspb.Message.getRepeatedField(msg, 1)) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.IntItemValue}
 */
proto.model.IntItemValue.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.IntItemValue;
  return proto.model.IntItemValue.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.IntItemValue} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.IntItemValue}
 */
proto.model.IntItemValue.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedInt32() : [reader.readInt32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addItemvalue(values[i]);
      }
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.IntItemValue.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.IntItemValue.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.IntItemValue} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.IntItemValue.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getItemvalueList();
  if (f.length > 0) {
    writer.writePackedInt32(
      1,
      f
    );
  }
};


/**
 * repeated int32 itemValue = 1;
 * @return {!Array<number>}
 */
proto.model.IntItemValue.prototype.getItemvalueList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 1));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.model.IntItemValue} returns this
 */
proto.model.IntItemValue.prototype.setItemvalueList = function(value) {
  return jspb.Message.setField(this, 1, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.model.IntItemValue} returns this
 */
proto.model.IntItemValue.prototype.addItemvalue = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 1, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.IntItemValue} returns this
 */
proto.model.IntItemValue.prototype.clearItemvalueList = function() {
  return this.setItemvalueList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.StringItemValue.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.StringItemValue.prototype.toObject = function(opt_includeInstance) {
  return proto.model.StringItemValue.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.StringItemValue} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.StringItemValue.toObject = function(includeInstance, msg) {
  var f, obj = {
    itemvalueList: (f = jspb.Message.getRepeatedField(msg, 1)) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.StringItemValue}
 */
proto.model.StringItemValue.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.StringItemValue;
  return proto.model.StringItemValue.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.StringItemValue} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.StringItemValue}
 */
proto.model.StringItemValue.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.addItemvalue(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.StringItemValue.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.StringItemValue.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.StringItemValue} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.StringItemValue.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getItemvalueList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      1,
      f
    );
  }
};


/**
 * repeated string itemValue = 1;
 * @return {!Array<string>}
 */
proto.model.StringItemValue.prototype.getItemvalueList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 1));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.model.StringItemValue} returns this
 */
proto.model.StringItemValue.prototype.setItemvalueList = function(value) {
  return jspb.Message.setField(this, 1, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.model.StringItemValue} returns this
 */
proto.model.StringItemValue.prototype.addItemvalue = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 1, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.StringItemValue} returns this
 */
proto.model.StringItemValue.prototype.clearItemvalueList = function() {
  return this.setItemvalueList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.BoolItemValue.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.BoolItemValue.prototype.toObject = function(opt_includeInstance) {
  return proto.model.BoolItemValue.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.BoolItemValue} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.BoolItemValue.toObject = function(includeInstance, msg) {
  var f, obj = {
    itemvalueList: (f = jspb.Message.getRepeatedBooleanField(msg, 1)) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.BoolItemValue}
 */
proto.model.BoolItemValue.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.BoolItemValue;
  return proto.model.BoolItemValue.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.BoolItemValue} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.BoolItemValue}
 */
proto.model.BoolItemValue.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var values = /** @type {!Array<boolean>} */ (reader.isDelimited() ? reader.readPackedBool() : [reader.readBool()]);
      for (var i = 0; i < values.length; i++) {
        msg.addItemvalue(values[i]);
      }
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.BoolItemValue.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.BoolItemValue.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.BoolItemValue} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.BoolItemValue.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getItemvalueList();
  if (f.length > 0) {
    writer.writePackedBool(
      1,
      f
    );
  }
};


/**
 * repeated bool itemValue = 1;
 * @return {!Array<boolean>}
 */
proto.model.BoolItemValue.prototype.getItemvalueList = function() {
  return /** @type {!Array<boolean>} */ (jspb.Message.getRepeatedBooleanField(this, 1));
};


/**
 * @param {!Array<boolean>} value
 * @return {!proto.model.BoolItemValue} returns this
 */
proto.model.BoolItemValue.prototype.setItemvalueList = function(value) {
  return jspb.Message.setField(this, 1, value || []);
};


/**
 * @param {boolean} value
 * @param {number=} opt_index
 * @return {!proto.model.BoolItemValue} returns this
 */
proto.model.BoolItemValue.prototype.addItemvalue = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 1, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.BoolItemValue} returns this
 */
proto.model.BoolItemValue.prototype.clearItemvalueList = function() {
  return this.setItemvalueList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.FloatItemValues.repeatedFields_ = [1,2,3];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.FloatItemValues.prototype.toObject = function(opt_includeInstance) {
  return proto.model.FloatItemValues.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.FloatItemValues} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.FloatItemValues.toObject = function(includeInstance, msg) {
  var f, obj = {
    groupnamesList: (f = jspb.Message.getRepeatedField(msg, 1)) == null ? undefined : f,
    itemnamesList: jspb.Message.toObjectList(msg.getItemnamesList(),
    proto.model.ItemName.toObject, includeInstance),
    itemvaluesList: jspb.Message.toObjectList(msg.getItemvaluesList(),
    proto.model.FloatItemValue.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.FloatItemValues}
 */
proto.model.FloatItemValues.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.FloatItemValues;
  return proto.model.FloatItemValues.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.FloatItemValues} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.FloatItemValues}
 */
proto.model.FloatItemValues.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.addGroupnames(value);
      break;
    case 2:
      var value = new proto.model.ItemName;
      reader.readMessage(value,proto.model.ItemName.deserializeBinaryFromReader);
      msg.addItemnames(value);
      break;
    case 3:
      var value = new proto.model.FloatItemValue;
      reader.readMessage(value,proto.model.FloatItemValue.deserializeBinaryFromReader);
      msg.addItemvalues(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.FloatItemValues.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.FloatItemValues.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.FloatItemValues} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.FloatItemValues.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getGroupnamesList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      1,
      f
    );
  }
  f = message.getItemnamesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      2,
      f,
      proto.model.ItemName.serializeBinaryToWriter
    );
  }
  f = message.getItemvaluesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      3,
      f,
      proto.model.FloatItemValue.serializeBinaryToWriter
    );
  }
};


/**
 * repeated string groupNames = 1;
 * @return {!Array<string>}
 */
proto.model.FloatItemValues.prototype.getGroupnamesList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 1));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.model.FloatItemValues} returns this
 */
proto.model.FloatItemValues.prototype.setGroupnamesList = function(value) {
  return jspb.Message.setField(this, 1, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.model.FloatItemValues} returns this
 */
proto.model.FloatItemValues.prototype.addGroupnames = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 1, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.FloatItemValues} returns this
 */
proto.model.FloatItemValues.prototype.clearGroupnamesList = function() {
  return this.setGroupnamesList([]);
};


/**
 * repeated ItemName itemNames = 2;
 * @return {!Array<!proto.model.ItemName>}
 */
proto.model.FloatItemValues.prototype.getItemnamesList = function() {
  return /** @type{!Array<!proto.model.ItemName>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.model.ItemName, 2));
};


/**
 * @param {!Array<!proto.model.ItemName>} value
 * @return {!proto.model.FloatItemValues} returns this
*/
proto.model.FloatItemValues.prototype.setItemnamesList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 2, value);
};


/**
 * @param {!proto.model.ItemName=} opt_value
 * @param {number=} opt_index
 * @return {!proto.model.ItemName}
 */
proto.model.FloatItemValues.prototype.addItemnames = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 2, opt_value, proto.model.ItemName, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.FloatItemValues} returns this
 */
proto.model.FloatItemValues.prototype.clearItemnamesList = function() {
  return this.setItemnamesList([]);
};


/**
 * repeated FloatItemValue itemValues = 3;
 * @return {!Array<!proto.model.FloatItemValue>}
 */
proto.model.FloatItemValues.prototype.getItemvaluesList = function() {
  return /** @type{!Array<!proto.model.FloatItemValue>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.model.FloatItemValue, 3));
};


/**
 * @param {!Array<!proto.model.FloatItemValue>} value
 * @return {!proto.model.FloatItemValues} returns this
*/
proto.model.FloatItemValues.prototype.setItemvaluesList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 3, value);
};


/**
 * @param {!proto.model.FloatItemValue=} opt_value
 * @param {number=} opt_index
 * @return {!proto.model.FloatItemValue}
 */
proto.model.FloatItemValues.prototype.addItemvalues = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 3, opt_value, proto.model.FloatItemValue, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.FloatItemValues} returns this
 */
proto.model.FloatItemValues.prototype.clearItemvaluesList = function() {
  return this.setItemvaluesList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.IntItemValues.repeatedFields_ = [1,2,3];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.IntItemValues.prototype.toObject = function(opt_includeInstance) {
  return proto.model.IntItemValues.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.IntItemValues} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.IntItemValues.toObject = function(includeInstance, msg) {
  var f, obj = {
    groupnamesList: (f = jspb.Message.getRepeatedField(msg, 1)) == null ? undefined : f,
    itemnamesList: jspb.Message.toObjectList(msg.getItemnamesList(),
    proto.model.ItemName.toObject, includeInstance),
    itemvaluesList: jspb.Message.toObjectList(msg.getItemvaluesList(),
    proto.model.IntItemValue.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.IntItemValues}
 */
proto.model.IntItemValues.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.IntItemValues;
  return proto.model.IntItemValues.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.IntItemValues} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.IntItemValues}
 */
proto.model.IntItemValues.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.addGroupnames(value);
      break;
    case 2:
      var value = new proto.model.ItemName;
      reader.readMessage(value,proto.model.ItemName.deserializeBinaryFromReader);
      msg.addItemnames(value);
      break;
    case 3:
      var value = new proto.model.IntItemValue;
      reader.readMessage(value,proto.model.IntItemValue.deserializeBinaryFromReader);
      msg.addItemvalues(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.IntItemValues.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.IntItemValues.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.IntItemValues} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.IntItemValues.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getGroupnamesList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      1,
      f
    );
  }
  f = message.getItemnamesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      2,
      f,
      proto.model.ItemName.serializeBinaryToWriter
    );
  }
  f = message.getItemvaluesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      3,
      f,
      proto.model.IntItemValue.serializeBinaryToWriter
    );
  }
};


/**
 * repeated string groupNames = 1;
 * @return {!Array<string>}
 */
proto.model.IntItemValues.prototype.getGroupnamesList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 1));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.model.IntItemValues} returns this
 */
proto.model.IntItemValues.prototype.setGroupnamesList = function(value) {
  return jspb.Message.setField(this, 1, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.model.IntItemValues} returns this
 */
proto.model.IntItemValues.prototype.addGroupnames = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 1, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.IntItemValues} returns this
 */
proto.model.IntItemValues.prototype.clearGroupnamesList = function() {
  return this.setGroupnamesList([]);
};


/**
 * repeated ItemName itemNames = 2;
 * @return {!Array<!proto.model.ItemName>}
 */
proto.model.IntItemValues.prototype.getItemnamesList = function() {
  return /** @type{!Array<!proto.model.ItemName>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.model.ItemName, 2));
};


/**
 * @param {!Array<!proto.model.ItemName>} value
 * @return {!proto.model.IntItemValues} returns this
*/
proto.model.IntItemValues.prototype.setItemnamesList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 2, value);
};


/**
 * @param {!proto.model.ItemName=} opt_value
 * @param {number=} opt_index
 * @return {!proto.model.ItemName}
 */
proto.model.IntItemValues.prototype.addItemnames = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 2, opt_value, proto.model.ItemName, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.IntItemValues} returns this
 */
proto.model.IntItemValues.prototype.clearItemnamesList = function() {
  return this.setItemnamesList([]);
};


/**
 * repeated IntItemValue itemValues = 3;
 * @return {!Array<!proto.model.IntItemValue>}
 */
proto.model.IntItemValues.prototype.getItemvaluesList = function() {
  return /** @type{!Array<!proto.model.IntItemValue>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.model.IntItemValue, 3));
};


/**
 * @param {!Array<!proto.model.IntItemValue>} value
 * @return {!proto.model.IntItemValues} returns this
*/
proto.model.IntItemValues.prototype.setItemvaluesList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 3, value);
};


/**
 * @param {!proto.model.IntItemValue=} opt_value
 * @param {number=} opt_index
 * @return {!proto.model.IntItemValue}
 */
proto.model.IntItemValues.prototype.addItemvalues = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 3, opt_value, proto.model.IntItemValue, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.IntItemValues} returns this
 */
proto.model.IntItemValues.prototype.clearItemvaluesList = function() {
  return this.setItemvaluesList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.BoolItemValues.repeatedFields_ = [1,2,3];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.BoolItemValues.prototype.toObject = function(opt_includeInstance) {
  return proto.model.BoolItemValues.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.BoolItemValues} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.BoolItemValues.toObject = function(includeInstance, msg) {
  var f, obj = {
    groupnamesList: (f = jspb.Message.getRepeatedField(msg, 1)) == null ? undefined : f,
    itemnamesList: jspb.Message.toObjectList(msg.getItemnamesList(),
    proto.model.ItemName.toObject, includeInstance),
    itemvaluesList: jspb.Message.toObjectList(msg.getItemvaluesList(),
    proto.model.BoolItemValue.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.BoolItemValues}
 */
proto.model.BoolItemValues.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.BoolItemValues;
  return proto.model.BoolItemValues.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.BoolItemValues} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.BoolItemValues}
 */
proto.model.BoolItemValues.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.addGroupnames(value);
      break;
    case 2:
      var value = new proto.model.ItemName;
      reader.readMessage(value,proto.model.ItemName.deserializeBinaryFromReader);
      msg.addItemnames(value);
      break;
    case 3:
      var value = new proto.model.BoolItemValue;
      reader.readMessage(value,proto.model.BoolItemValue.deserializeBinaryFromReader);
      msg.addItemvalues(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.BoolItemValues.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.BoolItemValues.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.BoolItemValues} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.BoolItemValues.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getGroupnamesList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      1,
      f
    );
  }
  f = message.getItemnamesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      2,
      f,
      proto.model.ItemName.serializeBinaryToWriter
    );
  }
  f = message.getItemvaluesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      3,
      f,
      proto.model.BoolItemValue.serializeBinaryToWriter
    );
  }
};


/**
 * repeated string groupNames = 1;
 * @return {!Array<string>}
 */
proto.model.BoolItemValues.prototype.getGroupnamesList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 1));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.model.BoolItemValues} returns this
 */
proto.model.BoolItemValues.prototype.setGroupnamesList = function(value) {
  return jspb.Message.setField(this, 1, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.model.BoolItemValues} returns this
 */
proto.model.BoolItemValues.prototype.addGroupnames = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 1, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.BoolItemValues} returns this
 */
proto.model.BoolItemValues.prototype.clearGroupnamesList = function() {
  return this.setGroupnamesList([]);
};


/**
 * repeated ItemName itemNames = 2;
 * @return {!Array<!proto.model.ItemName>}
 */
proto.model.BoolItemValues.prototype.getItemnamesList = function() {
  return /** @type{!Array<!proto.model.ItemName>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.model.ItemName, 2));
};


/**
 * @param {!Array<!proto.model.ItemName>} value
 * @return {!proto.model.BoolItemValues} returns this
*/
proto.model.BoolItemValues.prototype.setItemnamesList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 2, value);
};


/**
 * @param {!proto.model.ItemName=} opt_value
 * @param {number=} opt_index
 * @return {!proto.model.ItemName}
 */
proto.model.BoolItemValues.prototype.addItemnames = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 2, opt_value, proto.model.ItemName, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.BoolItemValues} returns this
 */
proto.model.BoolItemValues.prototype.clearItemnamesList = function() {
  return this.setItemnamesList([]);
};


/**
 * repeated BoolItemValue itemValues = 3;
 * @return {!Array<!proto.model.BoolItemValue>}
 */
proto.model.BoolItemValues.prototype.getItemvaluesList = function() {
  return /** @type{!Array<!proto.model.BoolItemValue>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.model.BoolItemValue, 3));
};


/**
 * @param {!Array<!proto.model.BoolItemValue>} value
 * @return {!proto.model.BoolItemValues} returns this
*/
proto.model.BoolItemValues.prototype.setItemvaluesList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 3, value);
};


/**
 * @param {!proto.model.BoolItemValue=} opt_value
 * @param {number=} opt_index
 * @return {!proto.model.BoolItemValue}
 */
proto.model.BoolItemValues.prototype.addItemvalues = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 3, opt_value, proto.model.BoolItemValue, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.BoolItemValues} returns this
 */
proto.model.BoolItemValues.prototype.clearItemvaluesList = function() {
  return this.setItemvaluesList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.StringItemValues.repeatedFields_ = [1,2,3];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.StringItemValues.prototype.toObject = function(opt_includeInstance) {
  return proto.model.StringItemValues.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.StringItemValues} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.StringItemValues.toObject = function(includeInstance, msg) {
  var f, obj = {
    groupnamesList: (f = jspb.Message.getRepeatedField(msg, 1)) == null ? undefined : f,
    itemnamesList: jspb.Message.toObjectList(msg.getItemnamesList(),
    proto.model.ItemName.toObject, includeInstance),
    itemvaluesList: jspb.Message.toObjectList(msg.getItemvaluesList(),
    proto.model.StringItemValue.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.StringItemValues}
 */
proto.model.StringItemValues.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.StringItemValues;
  return proto.model.StringItemValues.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.StringItemValues} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.StringItemValues}
 */
proto.model.StringItemValues.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.addGroupnames(value);
      break;
    case 2:
      var value = new proto.model.ItemName;
      reader.readMessage(value,proto.model.ItemName.deserializeBinaryFromReader);
      msg.addItemnames(value);
      break;
    case 3:
      var value = new proto.model.StringItemValue;
      reader.readMessage(value,proto.model.StringItemValue.deserializeBinaryFromReader);
      msg.addItemvalues(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.StringItemValues.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.StringItemValues.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.StringItemValues} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.StringItemValues.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getGroupnamesList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      1,
      f
    );
  }
  f = message.getItemnamesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      2,
      f,
      proto.model.ItemName.serializeBinaryToWriter
    );
  }
  f = message.getItemvaluesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      3,
      f,
      proto.model.StringItemValue.serializeBinaryToWriter
    );
  }
};


/**
 * repeated string groupNames = 1;
 * @return {!Array<string>}
 */
proto.model.StringItemValues.prototype.getGroupnamesList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 1));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.model.StringItemValues} returns this
 */
proto.model.StringItemValues.prototype.setGroupnamesList = function(value) {
  return jspb.Message.setField(this, 1, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.model.StringItemValues} returns this
 */
proto.model.StringItemValues.prototype.addGroupnames = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 1, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.StringItemValues} returns this
 */
proto.model.StringItemValues.prototype.clearGroupnamesList = function() {
  return this.setGroupnamesList([]);
};


/**
 * repeated ItemName itemNames = 2;
 * @return {!Array<!proto.model.ItemName>}
 */
proto.model.StringItemValues.prototype.getItemnamesList = function() {
  return /** @type{!Array<!proto.model.ItemName>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.model.ItemName, 2));
};


/**
 * @param {!Array<!proto.model.ItemName>} value
 * @return {!proto.model.StringItemValues} returns this
*/
proto.model.StringItemValues.prototype.setItemnamesList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 2, value);
};


/**
 * @param {!proto.model.ItemName=} opt_value
 * @param {number=} opt_index
 * @return {!proto.model.ItemName}
 */
proto.model.StringItemValues.prototype.addItemnames = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 2, opt_value, proto.model.ItemName, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.StringItemValues} returns this
 */
proto.model.StringItemValues.prototype.clearItemnamesList = function() {
  return this.setItemnamesList([]);
};


/**
 * repeated StringItemValue itemValues = 3;
 * @return {!Array<!proto.model.StringItemValue>}
 */
proto.model.StringItemValues.prototype.getItemvaluesList = function() {
  return /** @type{!Array<!proto.model.StringItemValue>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.model.StringItemValue, 3));
};


/**
 * @param {!Array<!proto.model.StringItemValue>} value
 * @return {!proto.model.StringItemValues} returns this
*/
proto.model.StringItemValues.prototype.setItemvaluesList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 3, value);
};


/**
 * @param {!proto.model.StringItemValue=} opt_value
 * @param {number=} opt_index
 * @return {!proto.model.StringItemValue}
 */
proto.model.StringItemValues.prototype.addItemvalues = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 3, opt_value, proto.model.StringItemValue, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.StringItemValues} returns this
 */
proto.model.StringItemValues.prototype.clearItemvaluesList = function() {
  return this.setItemvaluesList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.FloatHItemValues.repeatedFields_ = [3];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.FloatHItemValues.prototype.toObject = function(opt_includeInstance) {
  return proto.model.FloatHItemValues.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.FloatHItemValues} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.FloatHItemValues.toObject = function(includeInstance, msg) {
  var f, obj = {
    valuesList: jspb.Message.toObjectList(msg.getValuesList(),
    proto.model.FloatHItemValue.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.FloatHItemValues}
 */
proto.model.FloatHItemValues.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.FloatHItemValues;
  return proto.model.FloatHItemValues.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.FloatHItemValues} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.FloatHItemValues}
 */
proto.model.FloatHItemValues.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 3:
      var value = new proto.model.FloatHItemValue;
      reader.readMessage(value,proto.model.FloatHItemValue.deserializeBinaryFromReader);
      msg.addValues(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.FloatHItemValues.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.FloatHItemValues.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.FloatHItemValues} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.FloatHItemValues.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getValuesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      3,
      f,
      proto.model.FloatHItemValue.serializeBinaryToWriter
    );
  }
};


/**
 * repeated FloatHItemValue values = 3;
 * @return {!Array<!proto.model.FloatHItemValue>}
 */
proto.model.FloatHItemValues.prototype.getValuesList = function() {
  return /** @type{!Array<!proto.model.FloatHItemValue>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.model.FloatHItemValue, 3));
};


/**
 * @param {!Array<!proto.model.FloatHItemValue>} value
 * @return {!proto.model.FloatHItemValues} returns this
*/
proto.model.FloatHItemValues.prototype.setValuesList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 3, value);
};


/**
 * @param {!proto.model.FloatHItemValue=} opt_value
 * @param {number=} opt_index
 * @return {!proto.model.FloatHItemValue}
 */
proto.model.FloatHItemValues.prototype.addValues = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 3, opt_value, proto.model.FloatHItemValue, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.FloatHItemValues} returns this
 */
proto.model.FloatHItemValues.prototype.clearValuesList = function() {
  return this.setValuesList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.FloatHItemValue.repeatedFields_ = [2,3,4];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.FloatHItemValue.prototype.toObject = function(opt_includeInstance) {
  return proto.model.FloatHItemValue.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.FloatHItemValue} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.FloatHItemValue.toObject = function(includeInstance, msg) {
  var f, obj = {
    groupname: jspb.Message.getFieldWithDefault(msg, 1, ""),
    itemnameList: (f = jspb.Message.getRepeatedField(msg, 2)) == null ? undefined : f,
    itemvaluesList: (f = jspb.Message.getRepeatedFloatingPointField(msg, 3)) == null ? undefined : f,
    timestampsList: (f = jspb.Message.getRepeatedField(msg, 4)) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.FloatHItemValue}
 */
proto.model.FloatHItemValue.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.FloatHItemValue;
  return proto.model.FloatHItemValue.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.FloatHItemValue} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.FloatHItemValue}
 */
proto.model.FloatHItemValue.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setGroupname(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.addItemname(value);
      break;
    case 3:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedFloat() : [reader.readFloat()]);
      for (var i = 0; i < values.length; i++) {
        msg.addItemvalues(values[i]);
      }
      break;
    case 4:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedInt32() : [reader.readInt32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addTimestamps(values[i]);
      }
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.FloatHItemValue.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.FloatHItemValue.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.FloatHItemValue} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.FloatHItemValue.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getGroupname();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getItemnameList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      2,
      f
    );
  }
  f = message.getItemvaluesList();
  if (f.length > 0) {
    writer.writePackedFloat(
      3,
      f
    );
  }
  f = message.getTimestampsList();
  if (f.length > 0) {
    writer.writePackedInt32(
      4,
      f
    );
  }
};


/**
 * optional string groupName = 1;
 * @return {string}
 */
proto.model.FloatHItemValue.prototype.getGroupname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.FloatHItemValue} returns this
 */
proto.model.FloatHItemValue.prototype.setGroupname = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * repeated string itemName = 2;
 * @return {!Array<string>}
 */
proto.model.FloatHItemValue.prototype.getItemnameList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 2));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.model.FloatHItemValue} returns this
 */
proto.model.FloatHItemValue.prototype.setItemnameList = function(value) {
  return jspb.Message.setField(this, 2, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.model.FloatHItemValue} returns this
 */
proto.model.FloatHItemValue.prototype.addItemname = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 2, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.FloatHItemValue} returns this
 */
proto.model.FloatHItemValue.prototype.clearItemnameList = function() {
  return this.setItemnameList([]);
};


/**
 * repeated float itemValues = 3;
 * @return {!Array<number>}
 */
proto.model.FloatHItemValue.prototype.getItemvaluesList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedFloatingPointField(this, 3));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.model.FloatHItemValue} returns this
 */
proto.model.FloatHItemValue.prototype.setItemvaluesList = function(value) {
  return jspb.Message.setField(this, 3, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.model.FloatHItemValue} returns this
 */
proto.model.FloatHItemValue.prototype.addItemvalues = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 3, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.FloatHItemValue} returns this
 */
proto.model.FloatHItemValue.prototype.clearItemvaluesList = function() {
  return this.setItemvaluesList([]);
};


/**
 * repeated int32 timeStamps = 4;
 * @return {!Array<number>}
 */
proto.model.FloatHItemValue.prototype.getTimestampsList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 4));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.model.FloatHItemValue} returns this
 */
proto.model.FloatHItemValue.prototype.setTimestampsList = function(value) {
  return jspb.Message.setField(this, 4, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.model.FloatHItemValue} returns this
 */
proto.model.FloatHItemValue.prototype.addTimestamps = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 4, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.FloatHItemValue} returns this
 */
proto.model.FloatHItemValue.prototype.clearTimestampsList = function() {
  return this.setTimestampsList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.IntHItemValues.repeatedFields_ = [3];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.IntHItemValues.prototype.toObject = function(opt_includeInstance) {
  return proto.model.IntHItemValues.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.IntHItemValues} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.IntHItemValues.toObject = function(includeInstance, msg) {
  var f, obj = {
    valuesList: jspb.Message.toObjectList(msg.getValuesList(),
    proto.model.IntHItemValue.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.IntHItemValues}
 */
proto.model.IntHItemValues.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.IntHItemValues;
  return proto.model.IntHItemValues.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.IntHItemValues} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.IntHItemValues}
 */
proto.model.IntHItemValues.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 3:
      var value = new proto.model.IntHItemValue;
      reader.readMessage(value,proto.model.IntHItemValue.deserializeBinaryFromReader);
      msg.addValues(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.IntHItemValues.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.IntHItemValues.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.IntHItemValues} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.IntHItemValues.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getValuesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      3,
      f,
      proto.model.IntHItemValue.serializeBinaryToWriter
    );
  }
};


/**
 * repeated IntHItemValue values = 3;
 * @return {!Array<!proto.model.IntHItemValue>}
 */
proto.model.IntHItemValues.prototype.getValuesList = function() {
  return /** @type{!Array<!proto.model.IntHItemValue>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.model.IntHItemValue, 3));
};


/**
 * @param {!Array<!proto.model.IntHItemValue>} value
 * @return {!proto.model.IntHItemValues} returns this
*/
proto.model.IntHItemValues.prototype.setValuesList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 3, value);
};


/**
 * @param {!proto.model.IntHItemValue=} opt_value
 * @param {number=} opt_index
 * @return {!proto.model.IntHItemValue}
 */
proto.model.IntHItemValues.prototype.addValues = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 3, opt_value, proto.model.IntHItemValue, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.IntHItemValues} returns this
 */
proto.model.IntHItemValues.prototype.clearValuesList = function() {
  return this.setValuesList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.IntHItemValue.repeatedFields_ = [2,3,4];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.IntHItemValue.prototype.toObject = function(opt_includeInstance) {
  return proto.model.IntHItemValue.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.IntHItemValue} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.IntHItemValue.toObject = function(includeInstance, msg) {
  var f, obj = {
    groupname: jspb.Message.getFieldWithDefault(msg, 1, ""),
    itemnameList: (f = jspb.Message.getRepeatedField(msg, 2)) == null ? undefined : f,
    itemvaluesList: (f = jspb.Message.getRepeatedField(msg, 3)) == null ? undefined : f,
    timestampsList: (f = jspb.Message.getRepeatedField(msg, 4)) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.IntHItemValue}
 */
proto.model.IntHItemValue.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.IntHItemValue;
  return proto.model.IntHItemValue.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.IntHItemValue} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.IntHItemValue}
 */
proto.model.IntHItemValue.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setGroupname(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.addItemname(value);
      break;
    case 3:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedInt32() : [reader.readInt32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addItemvalues(values[i]);
      }
      break;
    case 4:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedInt32() : [reader.readInt32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addTimestamps(values[i]);
      }
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.IntHItemValue.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.IntHItemValue.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.IntHItemValue} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.IntHItemValue.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getGroupname();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getItemnameList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      2,
      f
    );
  }
  f = message.getItemvaluesList();
  if (f.length > 0) {
    writer.writePackedInt32(
      3,
      f
    );
  }
  f = message.getTimestampsList();
  if (f.length > 0) {
    writer.writePackedInt32(
      4,
      f
    );
  }
};


/**
 * optional string groupName = 1;
 * @return {string}
 */
proto.model.IntHItemValue.prototype.getGroupname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.IntHItemValue} returns this
 */
proto.model.IntHItemValue.prototype.setGroupname = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * repeated string itemName = 2;
 * @return {!Array<string>}
 */
proto.model.IntHItemValue.prototype.getItemnameList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 2));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.model.IntHItemValue} returns this
 */
proto.model.IntHItemValue.prototype.setItemnameList = function(value) {
  return jspb.Message.setField(this, 2, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.model.IntHItemValue} returns this
 */
proto.model.IntHItemValue.prototype.addItemname = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 2, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.IntHItemValue} returns this
 */
proto.model.IntHItemValue.prototype.clearItemnameList = function() {
  return this.setItemnameList([]);
};


/**
 * repeated int32 itemValues = 3;
 * @return {!Array<number>}
 */
proto.model.IntHItemValue.prototype.getItemvaluesList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 3));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.model.IntHItemValue} returns this
 */
proto.model.IntHItemValue.prototype.setItemvaluesList = function(value) {
  return jspb.Message.setField(this, 3, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.model.IntHItemValue} returns this
 */
proto.model.IntHItemValue.prototype.addItemvalues = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 3, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.IntHItemValue} returns this
 */
proto.model.IntHItemValue.prototype.clearItemvaluesList = function() {
  return this.setItemvaluesList([]);
};


/**
 * repeated int32 timeStamps = 4;
 * @return {!Array<number>}
 */
proto.model.IntHItemValue.prototype.getTimestampsList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 4));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.model.IntHItemValue} returns this
 */
proto.model.IntHItemValue.prototype.setTimestampsList = function(value) {
  return jspb.Message.setField(this, 4, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.model.IntHItemValue} returns this
 */
proto.model.IntHItemValue.prototype.addTimestamps = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 4, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.IntHItemValue} returns this
 */
proto.model.IntHItemValue.prototype.clearTimestampsList = function() {
  return this.setTimestampsList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.StringHItemValues.repeatedFields_ = [3];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.StringHItemValues.prototype.toObject = function(opt_includeInstance) {
  return proto.model.StringHItemValues.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.StringHItemValues} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.StringHItemValues.toObject = function(includeInstance, msg) {
  var f, obj = {
    valuesList: jspb.Message.toObjectList(msg.getValuesList(),
    proto.model.StringHItemValue.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.StringHItemValues}
 */
proto.model.StringHItemValues.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.StringHItemValues;
  return proto.model.StringHItemValues.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.StringHItemValues} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.StringHItemValues}
 */
proto.model.StringHItemValues.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 3:
      var value = new proto.model.StringHItemValue;
      reader.readMessage(value,proto.model.StringHItemValue.deserializeBinaryFromReader);
      msg.addValues(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.StringHItemValues.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.StringHItemValues.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.StringHItemValues} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.StringHItemValues.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getValuesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      3,
      f,
      proto.model.StringHItemValue.serializeBinaryToWriter
    );
  }
};


/**
 * repeated StringHItemValue values = 3;
 * @return {!Array<!proto.model.StringHItemValue>}
 */
proto.model.StringHItemValues.prototype.getValuesList = function() {
  return /** @type{!Array<!proto.model.StringHItemValue>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.model.StringHItemValue, 3));
};


/**
 * @param {!Array<!proto.model.StringHItemValue>} value
 * @return {!proto.model.StringHItemValues} returns this
*/
proto.model.StringHItemValues.prototype.setValuesList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 3, value);
};


/**
 * @param {!proto.model.StringHItemValue=} opt_value
 * @param {number=} opt_index
 * @return {!proto.model.StringHItemValue}
 */
proto.model.StringHItemValues.prototype.addValues = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 3, opt_value, proto.model.StringHItemValue, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.StringHItemValues} returns this
 */
proto.model.StringHItemValues.prototype.clearValuesList = function() {
  return this.setValuesList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.StringHItemValue.repeatedFields_ = [2,3,4];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.StringHItemValue.prototype.toObject = function(opt_includeInstance) {
  return proto.model.StringHItemValue.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.StringHItemValue} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.StringHItemValue.toObject = function(includeInstance, msg) {
  var f, obj = {
    groupname: jspb.Message.getFieldWithDefault(msg, 1, ""),
    itemnameList: (f = jspb.Message.getRepeatedField(msg, 2)) == null ? undefined : f,
    itemvaluesList: (f = jspb.Message.getRepeatedField(msg, 3)) == null ? undefined : f,
    timestampsList: (f = jspb.Message.getRepeatedField(msg, 4)) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.StringHItemValue}
 */
proto.model.StringHItemValue.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.StringHItemValue;
  return proto.model.StringHItemValue.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.StringHItemValue} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.StringHItemValue}
 */
proto.model.StringHItemValue.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setGroupname(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.addItemname(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.addItemvalues(value);
      break;
    case 4:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedInt32() : [reader.readInt32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addTimestamps(values[i]);
      }
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.StringHItemValue.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.StringHItemValue.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.StringHItemValue} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.StringHItemValue.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getGroupname();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getItemnameList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      2,
      f
    );
  }
  f = message.getItemvaluesList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      3,
      f
    );
  }
  f = message.getTimestampsList();
  if (f.length > 0) {
    writer.writePackedInt32(
      4,
      f
    );
  }
};


/**
 * optional string groupName = 1;
 * @return {string}
 */
proto.model.StringHItemValue.prototype.getGroupname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.StringHItemValue} returns this
 */
proto.model.StringHItemValue.prototype.setGroupname = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * repeated string itemName = 2;
 * @return {!Array<string>}
 */
proto.model.StringHItemValue.prototype.getItemnameList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 2));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.model.StringHItemValue} returns this
 */
proto.model.StringHItemValue.prototype.setItemnameList = function(value) {
  return jspb.Message.setField(this, 2, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.model.StringHItemValue} returns this
 */
proto.model.StringHItemValue.prototype.addItemname = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 2, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.StringHItemValue} returns this
 */
proto.model.StringHItemValue.prototype.clearItemnameList = function() {
  return this.setItemnameList([]);
};


/**
 * repeated string itemValues = 3;
 * @return {!Array<string>}
 */
proto.model.StringHItemValue.prototype.getItemvaluesList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 3));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.model.StringHItemValue} returns this
 */
proto.model.StringHItemValue.prototype.setItemvaluesList = function(value) {
  return jspb.Message.setField(this, 3, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.model.StringHItemValue} returns this
 */
proto.model.StringHItemValue.prototype.addItemvalues = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 3, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.StringHItemValue} returns this
 */
proto.model.StringHItemValue.prototype.clearItemvaluesList = function() {
  return this.setItemvaluesList([]);
};


/**
 * repeated int32 timeStamps = 4;
 * @return {!Array<number>}
 */
proto.model.StringHItemValue.prototype.getTimestampsList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 4));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.model.StringHItemValue} returns this
 */
proto.model.StringHItemValue.prototype.setTimestampsList = function(value) {
  return jspb.Message.setField(this, 4, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.model.StringHItemValue} returns this
 */
proto.model.StringHItemValue.prototype.addTimestamps = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 4, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.StringHItemValue} returns this
 */
proto.model.StringHItemValue.prototype.clearTimestampsList = function() {
  return this.setTimestampsList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.BoolHItemValues.repeatedFields_ = [3];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.BoolHItemValues.prototype.toObject = function(opt_includeInstance) {
  return proto.model.BoolHItemValues.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.BoolHItemValues} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.BoolHItemValues.toObject = function(includeInstance, msg) {
  var f, obj = {
    valuesList: jspb.Message.toObjectList(msg.getValuesList(),
    proto.model.BoolHItemValue.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.BoolHItemValues}
 */
proto.model.BoolHItemValues.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.BoolHItemValues;
  return proto.model.BoolHItemValues.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.BoolHItemValues} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.BoolHItemValues}
 */
proto.model.BoolHItemValues.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 3:
      var value = new proto.model.BoolHItemValue;
      reader.readMessage(value,proto.model.BoolHItemValue.deserializeBinaryFromReader);
      msg.addValues(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.BoolHItemValues.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.BoolHItemValues.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.BoolHItemValues} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.BoolHItemValues.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getValuesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      3,
      f,
      proto.model.BoolHItemValue.serializeBinaryToWriter
    );
  }
};


/**
 * repeated BoolHItemValue values = 3;
 * @return {!Array<!proto.model.BoolHItemValue>}
 */
proto.model.BoolHItemValues.prototype.getValuesList = function() {
  return /** @type{!Array<!proto.model.BoolHItemValue>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.model.BoolHItemValue, 3));
};


/**
 * @param {!Array<!proto.model.BoolHItemValue>} value
 * @return {!proto.model.BoolHItemValues} returns this
*/
proto.model.BoolHItemValues.prototype.setValuesList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 3, value);
};


/**
 * @param {!proto.model.BoolHItemValue=} opt_value
 * @param {number=} opt_index
 * @return {!proto.model.BoolHItemValue}
 */
proto.model.BoolHItemValues.prototype.addValues = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 3, opt_value, proto.model.BoolHItemValue, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.BoolHItemValues} returns this
 */
proto.model.BoolHItemValues.prototype.clearValuesList = function() {
  return this.setValuesList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.BoolHItemValue.repeatedFields_ = [2,3,4];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.BoolHItemValue.prototype.toObject = function(opt_includeInstance) {
  return proto.model.BoolHItemValue.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.BoolHItemValue} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.BoolHItemValue.toObject = function(includeInstance, msg) {
  var f, obj = {
    groupname: jspb.Message.getFieldWithDefault(msg, 1, ""),
    itemnameList: (f = jspb.Message.getRepeatedField(msg, 2)) == null ? undefined : f,
    itemvaluesList: (f = jspb.Message.getRepeatedBooleanField(msg, 3)) == null ? undefined : f,
    timestampsList: (f = jspb.Message.getRepeatedField(msg, 4)) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.BoolHItemValue}
 */
proto.model.BoolHItemValue.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.BoolHItemValue;
  return proto.model.BoolHItemValue.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.BoolHItemValue} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.BoolHItemValue}
 */
proto.model.BoolHItemValue.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setGroupname(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.addItemname(value);
      break;
    case 3:
      var values = /** @type {!Array<boolean>} */ (reader.isDelimited() ? reader.readPackedBool() : [reader.readBool()]);
      for (var i = 0; i < values.length; i++) {
        msg.addItemvalues(values[i]);
      }
      break;
    case 4:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedInt32() : [reader.readInt32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addTimestamps(values[i]);
      }
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.BoolHItemValue.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.BoolHItemValue.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.BoolHItemValue} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.BoolHItemValue.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getGroupname();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getItemnameList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      2,
      f
    );
  }
  f = message.getItemvaluesList();
  if (f.length > 0) {
    writer.writePackedBool(
      3,
      f
    );
  }
  f = message.getTimestampsList();
  if (f.length > 0) {
    writer.writePackedInt32(
      4,
      f
    );
  }
};


/**
 * optional string groupName = 1;
 * @return {string}
 */
proto.model.BoolHItemValue.prototype.getGroupname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.BoolHItemValue} returns this
 */
proto.model.BoolHItemValue.prototype.setGroupname = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * repeated string itemName = 2;
 * @return {!Array<string>}
 */
proto.model.BoolHItemValue.prototype.getItemnameList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 2));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.model.BoolHItemValue} returns this
 */
proto.model.BoolHItemValue.prototype.setItemnameList = function(value) {
  return jspb.Message.setField(this, 2, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.model.BoolHItemValue} returns this
 */
proto.model.BoolHItemValue.prototype.addItemname = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 2, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.BoolHItemValue} returns this
 */
proto.model.BoolHItemValue.prototype.clearItemnameList = function() {
  return this.setItemnameList([]);
};


/**
 * repeated bool itemValues = 3;
 * @return {!Array<boolean>}
 */
proto.model.BoolHItemValue.prototype.getItemvaluesList = function() {
  return /** @type {!Array<boolean>} */ (jspb.Message.getRepeatedBooleanField(this, 3));
};


/**
 * @param {!Array<boolean>} value
 * @return {!proto.model.BoolHItemValue} returns this
 */
proto.model.BoolHItemValue.prototype.setItemvaluesList = function(value) {
  return jspb.Message.setField(this, 3, value || []);
};


/**
 * @param {boolean} value
 * @param {number=} opt_index
 * @return {!proto.model.BoolHItemValue} returns this
 */
proto.model.BoolHItemValue.prototype.addItemvalues = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 3, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.BoolHItemValue} returns this
 */
proto.model.BoolHItemValue.prototype.clearItemvaluesList = function() {
  return this.setItemvaluesList([]);
};


/**
 * repeated int32 timeStamps = 4;
 * @return {!Array<number>}
 */
proto.model.BoolHItemValue.prototype.getTimestampsList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 4));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.model.BoolHItemValue} returns this
 */
proto.model.BoolHItemValue.prototype.setTimestampsList = function(value) {
  return jspb.Message.setField(this, 4, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.model.BoolHItemValue} returns this
 */
proto.model.BoolHItemValue.prototype.addTimestamps = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 4, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.BoolHItemValue} returns this
 */
proto.model.BoolHItemValue.prototype.clearTimestampsList = function() {
  return this.setTimestampsList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.QueryRealTimeDataString.repeatedFields_ = [2,1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.QueryRealTimeDataString.prototype.toObject = function(opt_includeInstance) {
  return proto.model.QueryRealTimeDataString.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.QueryRealTimeDataString} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.QueryRealTimeDataString.toObject = function(includeInstance, msg) {
  var f, obj = {
    groupnamesList: (f = jspb.Message.getRepeatedField(msg, 2)) == null ? undefined : f,
    itemnamesList: (f = jspb.Message.getRepeatedField(msg, 1)) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.QueryRealTimeDataString}
 */
proto.model.QueryRealTimeDataString.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.QueryRealTimeDataString;
  return proto.model.QueryRealTimeDataString.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.QueryRealTimeDataString} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.QueryRealTimeDataString}
 */
proto.model.QueryRealTimeDataString.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.addGroupnames(value);
      break;
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.addItemnames(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.QueryRealTimeDataString.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.QueryRealTimeDataString.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.QueryRealTimeDataString} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.QueryRealTimeDataString.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getGroupnamesList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      2,
      f
    );
  }
  f = message.getItemnamesList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      1,
      f
    );
  }
};


/**
 * repeated string groupNames = 2;
 * @return {!Array<string>}
 */
proto.model.QueryRealTimeDataString.prototype.getGroupnamesList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 2));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.model.QueryRealTimeDataString} returns this
 */
proto.model.QueryRealTimeDataString.prototype.setGroupnamesList = function(value) {
  return jspb.Message.setField(this, 2, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.model.QueryRealTimeDataString} returns this
 */
proto.model.QueryRealTimeDataString.prototype.addGroupnames = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 2, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.QueryRealTimeDataString} returns this
 */
proto.model.QueryRealTimeDataString.prototype.clearGroupnamesList = function() {
  return this.setGroupnamesList([]);
};


/**
 * repeated string itemNames = 1;
 * @return {!Array<string>}
 */
proto.model.QueryRealTimeDataString.prototype.getItemnamesList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 1));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.model.QueryRealTimeDataString} returns this
 */
proto.model.QueryRealTimeDataString.prototype.setItemnamesList = function(value) {
  return jspb.Message.setField(this, 1, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.model.QueryRealTimeDataString} returns this
 */
proto.model.QueryRealTimeDataString.prototype.addItemnames = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 1, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.QueryRealTimeDataString} returns this
 */
proto.model.QueryRealTimeDataString.prototype.clearItemnamesList = function() {
  return this.setItemnamesList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.QueryHistoricalDataString.repeatedFields_ = [5,1,2,3,4];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.QueryHistoricalDataString.prototype.toObject = function(opt_includeInstance) {
  return proto.model.QueryHistoricalDataString.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.QueryHistoricalDataString} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.QueryHistoricalDataString.toObject = function(includeInstance, msg) {
  var f, obj = {
    groupnamesList: (f = jspb.Message.getRepeatedField(msg, 5)) == null ? undefined : f,
    itemnamesList: (f = jspb.Message.getRepeatedField(msg, 1)) == null ? undefined : f,
    starttimesList: (f = jspb.Message.getRepeatedField(msg, 2)) == null ? undefined : f,
    endtimesList: (f = jspb.Message.getRepeatedField(msg, 3)) == null ? undefined : f,
    intervalsList: (f = jspb.Message.getRepeatedField(msg, 4)) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.QueryHistoricalDataString}
 */
proto.model.QueryHistoricalDataString.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.QueryHistoricalDataString;
  return proto.model.QueryHistoricalDataString.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.QueryHistoricalDataString} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.QueryHistoricalDataString}
 */
proto.model.QueryHistoricalDataString.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.addGroupnames(value);
      break;
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.addItemnames(value);
      break;
    case 2:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedInt32() : [reader.readInt32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addStarttimes(values[i]);
      }
      break;
    case 3:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedInt32() : [reader.readInt32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addEndtimes(values[i]);
      }
      break;
    case 4:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedInt32() : [reader.readInt32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addIntervals(values[i]);
      }
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.QueryHistoricalDataString.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.QueryHistoricalDataString.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.QueryHistoricalDataString} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.QueryHistoricalDataString.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getGroupnamesList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      5,
      f
    );
  }
  f = message.getItemnamesList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      1,
      f
    );
  }
  f = message.getStarttimesList();
  if (f.length > 0) {
    writer.writePackedInt32(
      2,
      f
    );
  }
  f = message.getEndtimesList();
  if (f.length > 0) {
    writer.writePackedInt32(
      3,
      f
    );
  }
  f = message.getIntervalsList();
  if (f.length > 0) {
    writer.writePackedInt32(
      4,
      f
    );
  }
};


/**
 * repeated string groupNames = 5;
 * @return {!Array<string>}
 */
proto.model.QueryHistoricalDataString.prototype.getGroupnamesList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 5));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.model.QueryHistoricalDataString} returns this
 */
proto.model.QueryHistoricalDataString.prototype.setGroupnamesList = function(value) {
  return jspb.Message.setField(this, 5, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.model.QueryHistoricalDataString} returns this
 */
proto.model.QueryHistoricalDataString.prototype.addGroupnames = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 5, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.QueryHistoricalDataString} returns this
 */
proto.model.QueryHistoricalDataString.prototype.clearGroupnamesList = function() {
  return this.setGroupnamesList([]);
};


/**
 * repeated string itemNames = 1;
 * @return {!Array<string>}
 */
proto.model.QueryHistoricalDataString.prototype.getItemnamesList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 1));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.model.QueryHistoricalDataString} returns this
 */
proto.model.QueryHistoricalDataString.prototype.setItemnamesList = function(value) {
  return jspb.Message.setField(this, 1, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.model.QueryHistoricalDataString} returns this
 */
proto.model.QueryHistoricalDataString.prototype.addItemnames = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 1, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.QueryHistoricalDataString} returns this
 */
proto.model.QueryHistoricalDataString.prototype.clearItemnamesList = function() {
  return this.setItemnamesList([]);
};


/**
 * repeated int32 startTimes = 2;
 * @return {!Array<number>}
 */
proto.model.QueryHistoricalDataString.prototype.getStarttimesList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 2));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.model.QueryHistoricalDataString} returns this
 */
proto.model.QueryHistoricalDataString.prototype.setStarttimesList = function(value) {
  return jspb.Message.setField(this, 2, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.model.QueryHistoricalDataString} returns this
 */
proto.model.QueryHistoricalDataString.prototype.addStarttimes = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 2, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.QueryHistoricalDataString} returns this
 */
proto.model.QueryHistoricalDataString.prototype.clearStarttimesList = function() {
  return this.setStarttimesList([]);
};


/**
 * repeated int32 endTimes = 3;
 * @return {!Array<number>}
 */
proto.model.QueryHistoricalDataString.prototype.getEndtimesList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 3));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.model.QueryHistoricalDataString} returns this
 */
proto.model.QueryHistoricalDataString.prototype.setEndtimesList = function(value) {
  return jspb.Message.setField(this, 3, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.model.QueryHistoricalDataString} returns this
 */
proto.model.QueryHistoricalDataString.prototype.addEndtimes = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 3, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.QueryHistoricalDataString} returns this
 */
proto.model.QueryHistoricalDataString.prototype.clearEndtimesList = function() {
  return this.setEndtimesList([]);
};


/**
 * repeated int32 intervals = 4;
 * @return {!Array<number>}
 */
proto.model.QueryHistoricalDataString.prototype.getIntervalsList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 4));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.model.QueryHistoricalDataString} returns this
 */
proto.model.QueryHistoricalDataString.prototype.setIntervalsList = function(value) {
  return jspb.Message.setField(this, 4, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.model.QueryHistoricalDataString} returns this
 */
proto.model.QueryHistoricalDataString.prototype.addIntervals = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 4, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.QueryHistoricalDataString} returns this
 */
proto.model.QueryHistoricalDataString.prototype.clearIntervalsList = function() {
  return this.setIntervalsList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.QueryRawHistoricalDataString.repeatedFields_ = [1,2];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.QueryRawHistoricalDataString.prototype.toObject = function(opt_includeInstance) {
  return proto.model.QueryRawHistoricalDataString.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.QueryRawHistoricalDataString} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.QueryRawHistoricalDataString.toObject = function(includeInstance, msg) {
  var f, obj = {
    groupnamesList: (f = jspb.Message.getRepeatedField(msg, 1)) == null ? undefined : f,
    itemnamesList: (f = jspb.Message.getRepeatedField(msg, 2)) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.QueryRawHistoricalDataString}
 */
proto.model.QueryRawHistoricalDataString.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.QueryRawHistoricalDataString;
  return proto.model.QueryRawHistoricalDataString.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.QueryRawHistoricalDataString} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.QueryRawHistoricalDataString}
 */
proto.model.QueryRawHistoricalDataString.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.addGroupnames(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.addItemnames(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.QueryRawHistoricalDataString.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.QueryRawHistoricalDataString.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.QueryRawHistoricalDataString} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.QueryRawHistoricalDataString.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getGroupnamesList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      1,
      f
    );
  }
  f = message.getItemnamesList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      2,
      f
    );
  }
};


/**
 * repeated string groupNames = 1;
 * @return {!Array<string>}
 */
proto.model.QueryRawHistoricalDataString.prototype.getGroupnamesList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 1));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.model.QueryRawHistoricalDataString} returns this
 */
proto.model.QueryRawHistoricalDataString.prototype.setGroupnamesList = function(value) {
  return jspb.Message.setField(this, 1, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.model.QueryRawHistoricalDataString} returns this
 */
proto.model.QueryRawHistoricalDataString.prototype.addGroupnames = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 1, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.QueryRawHistoricalDataString} returns this
 */
proto.model.QueryRawHistoricalDataString.prototype.clearGroupnamesList = function() {
  return this.setGroupnamesList([]);
};


/**
 * repeated string itemNames = 2;
 * @return {!Array<string>}
 */
proto.model.QueryRawHistoricalDataString.prototype.getItemnamesList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 2));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.model.QueryRawHistoricalDataString} returns this
 */
proto.model.QueryRawHistoricalDataString.prototype.setItemnamesList = function(value) {
  return jspb.Message.setField(this, 2, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.model.QueryRawHistoricalDataString} returns this
 */
proto.model.QueryRawHistoricalDataString.prototype.addItemnames = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 2, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.QueryRawHistoricalDataString} returns this
 */
proto.model.QueryRawHistoricalDataString.prototype.clearItemnamesList = function() {
  return this.setItemnamesList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.QueryHistoricalDataWithStamp.repeatedFields_ = [3];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.QueryHistoricalDataWithStamp.prototype.toObject = function(opt_includeInstance) {
  return proto.model.QueryHistoricalDataWithStamp.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.QueryHistoricalDataWithStamp} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.QueryHistoricalDataWithStamp.toObject = function(includeInstance, msg) {
  var f, obj = {
    groupname: jspb.Message.getFieldWithDefault(msg, 1, ""),
    itemname: jspb.Message.getFieldWithDefault(msg, 2, ""),
    timestampsList: (f = jspb.Message.getRepeatedField(msg, 3)) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.QueryHistoricalDataWithStamp}
 */
proto.model.QueryHistoricalDataWithStamp.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.QueryHistoricalDataWithStamp;
  return proto.model.QueryHistoricalDataWithStamp.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.QueryHistoricalDataWithStamp} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.QueryHistoricalDataWithStamp}
 */
proto.model.QueryHistoricalDataWithStamp.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setGroupname(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setItemname(value);
      break;
    case 3:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedInt32() : [reader.readInt32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addTimestamps(values[i]);
      }
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.QueryHistoricalDataWithStamp.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.QueryHistoricalDataWithStamp.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.QueryHistoricalDataWithStamp} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.QueryHistoricalDataWithStamp.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getGroupname();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getItemname();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getTimestampsList();
  if (f.length > 0) {
    writer.writePackedInt32(
      3,
      f
    );
  }
};


/**
 * optional string groupName = 1;
 * @return {string}
 */
proto.model.QueryHistoricalDataWithStamp.prototype.getGroupname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.QueryHistoricalDataWithStamp} returns this
 */
proto.model.QueryHistoricalDataWithStamp.prototype.setGroupname = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string itemName = 2;
 * @return {string}
 */
proto.model.QueryHistoricalDataWithStamp.prototype.getItemname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.QueryHistoricalDataWithStamp} returns this
 */
proto.model.QueryHistoricalDataWithStamp.prototype.setItemname = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * repeated int32 timeStamps = 3;
 * @return {!Array<number>}
 */
proto.model.QueryHistoricalDataWithStamp.prototype.getTimestampsList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 3));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.model.QueryHistoricalDataWithStamp} returns this
 */
proto.model.QueryHistoricalDataWithStamp.prototype.setTimestampsList = function(value) {
  return jspb.Message.setField(this, 3, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.model.QueryHistoricalDataWithStamp} returns this
 */
proto.model.QueryHistoricalDataWithStamp.prototype.addTimestamps = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 3, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.QueryHistoricalDataWithStamp} returns this
 */
proto.model.QueryHistoricalDataWithStamp.prototype.clearTimestampsList = function() {
  return this.setTimestampsList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.QueryHistoricalDataWithStampString.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.QueryHistoricalDataWithStampString.prototype.toObject = function(opt_includeInstance) {
  return proto.model.QueryHistoricalDataWithStampString.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.QueryHistoricalDataWithStampString} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.QueryHistoricalDataWithStampString.toObject = function(includeInstance, msg) {
  var f, obj = {
    querystringList: jspb.Message.toObjectList(msg.getQuerystringList(),
    proto.model.QueryHistoricalDataWithStamp.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.QueryHistoricalDataWithStampString}
 */
proto.model.QueryHistoricalDataWithStampString.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.QueryHistoricalDataWithStampString;
  return proto.model.QueryHistoricalDataWithStampString.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.QueryHistoricalDataWithStampString} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.QueryHistoricalDataWithStampString}
 */
proto.model.QueryHistoricalDataWithStampString.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.model.QueryHistoricalDataWithStamp;
      reader.readMessage(value,proto.model.QueryHistoricalDataWithStamp.deserializeBinaryFromReader);
      msg.addQuerystring(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.QueryHistoricalDataWithStampString.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.QueryHistoricalDataWithStampString.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.QueryHistoricalDataWithStampString} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.QueryHistoricalDataWithStampString.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getQuerystringList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.model.QueryHistoricalDataWithStamp.serializeBinaryToWriter
    );
  }
};


/**
 * repeated QueryHistoricalDataWithStamp queryString = 1;
 * @return {!Array<!proto.model.QueryHistoricalDataWithStamp>}
 */
proto.model.QueryHistoricalDataWithStampString.prototype.getQuerystringList = function() {
  return /** @type{!Array<!proto.model.QueryHistoricalDataWithStamp>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.model.QueryHistoricalDataWithStamp, 1));
};


/**
 * @param {!Array<!proto.model.QueryHistoricalDataWithStamp>} value
 * @return {!proto.model.QueryHistoricalDataWithStampString} returns this
*/
proto.model.QueryHistoricalDataWithStampString.prototype.setQuerystringList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.model.QueryHistoricalDataWithStamp=} opt_value
 * @param {number=} opt_index
 * @return {!proto.model.QueryHistoricalDataWithStamp}
 */
proto.model.QueryHistoricalDataWithStampString.prototype.addQuerystring = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.model.QueryHistoricalDataWithStamp, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.QueryHistoricalDataWithStampString} returns this
 */
proto.model.QueryHistoricalDataWithStampString.prototype.clearQuerystringList = function() {
  return this.setQuerystringList([]);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.DeadZone.prototype.toObject = function(opt_includeInstance) {
  return proto.model.DeadZone.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.DeadZone} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.DeadZone.toObject = function(includeInstance, msg) {
  var f, obj = {
    itemname: jspb.Message.getFieldWithDefault(msg, 1, ""),
    deadzonecount: jspb.Message.getFieldWithDefault(msg, 2, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.DeadZone}
 */
proto.model.DeadZone.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.DeadZone;
  return proto.model.DeadZone.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.DeadZone} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.DeadZone}
 */
proto.model.DeadZone.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setItemname(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setDeadzonecount(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.DeadZone.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.DeadZone.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.DeadZone} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.DeadZone.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getItemname();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getDeadzonecount();
  if (f !== 0) {
    writer.writeInt32(
      2,
      f
    );
  }
};


/**
 * optional string itemName = 1;
 * @return {string}
 */
proto.model.DeadZone.prototype.getItemname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.DeadZone} returns this
 */
proto.model.DeadZone.prototype.setItemname = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional int32 deadZoneCount = 2;
 * @return {number}
 */
proto.model.DeadZone.prototype.getDeadzonecount = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.model.DeadZone} returns this
 */
proto.model.DeadZone.prototype.setDeadzonecount = function(value) {
  return jspb.Message.setProto3IntField(this, 2, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.QueryHistoricalDataWithConditionString.repeatedFields_ = [2,3,4,5,7];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.QueryHistoricalDataWithConditionString.prototype.toObject = function(opt_includeInstance) {
  return proto.model.QueryHistoricalDataWithConditionString.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.QueryHistoricalDataWithConditionString} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.QueryHistoricalDataWithConditionString.toObject = function(includeInstance, msg) {
  var f, obj = {
    groupname: jspb.Message.getFieldWithDefault(msg, 1, ""),
    itemnamesList: (f = jspb.Message.getRepeatedField(msg, 2)) == null ? undefined : f,
    starttimesList: (f = jspb.Message.getRepeatedField(msg, 3)) == null ? undefined : f,
    endtimesList: (f = jspb.Message.getRepeatedField(msg, 4)) == null ? undefined : f,
    intervalsList: (f = jspb.Message.getRepeatedField(msg, 5)) == null ? undefined : f,
    filtercondition: jspb.Message.getFieldWithDefault(msg, 6, ""),
    deadzonesList: jspb.Message.toObjectList(msg.getDeadzonesList(),
    proto.model.DeadZone.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.QueryHistoricalDataWithConditionString}
 */
proto.model.QueryHistoricalDataWithConditionString.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.QueryHistoricalDataWithConditionString;
  return proto.model.QueryHistoricalDataWithConditionString.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.QueryHistoricalDataWithConditionString} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.QueryHistoricalDataWithConditionString}
 */
proto.model.QueryHistoricalDataWithConditionString.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setGroupname(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.addItemnames(value);
      break;
    case 3:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedInt32() : [reader.readInt32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addStarttimes(values[i]);
      }
      break;
    case 4:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedInt32() : [reader.readInt32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addEndtimes(values[i]);
      }
      break;
    case 5:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedInt32() : [reader.readInt32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addIntervals(values[i]);
      }
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setFiltercondition(value);
      break;
    case 7:
      var value = new proto.model.DeadZone;
      reader.readMessage(value,proto.model.DeadZone.deserializeBinaryFromReader);
      msg.addDeadzones(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.QueryHistoricalDataWithConditionString.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.QueryHistoricalDataWithConditionString.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.QueryHistoricalDataWithConditionString} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.QueryHistoricalDataWithConditionString.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getGroupname();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getItemnamesList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      2,
      f
    );
  }
  f = message.getStarttimesList();
  if (f.length > 0) {
    writer.writePackedInt32(
      3,
      f
    );
  }
  f = message.getEndtimesList();
  if (f.length > 0) {
    writer.writePackedInt32(
      4,
      f
    );
  }
  f = message.getIntervalsList();
  if (f.length > 0) {
    writer.writePackedInt32(
      5,
      f
    );
  }
  f = message.getFiltercondition();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
  f = message.getDeadzonesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      7,
      f,
      proto.model.DeadZone.serializeBinaryToWriter
    );
  }
};


/**
 * optional string groupName = 1;
 * @return {string}
 */
proto.model.QueryHistoricalDataWithConditionString.prototype.getGroupname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.QueryHistoricalDataWithConditionString} returns this
 */
proto.model.QueryHistoricalDataWithConditionString.prototype.setGroupname = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * repeated string itemNames = 2;
 * @return {!Array<string>}
 */
proto.model.QueryHistoricalDataWithConditionString.prototype.getItemnamesList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 2));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.model.QueryHistoricalDataWithConditionString} returns this
 */
proto.model.QueryHistoricalDataWithConditionString.prototype.setItemnamesList = function(value) {
  return jspb.Message.setField(this, 2, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.model.QueryHistoricalDataWithConditionString} returns this
 */
proto.model.QueryHistoricalDataWithConditionString.prototype.addItemnames = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 2, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.QueryHistoricalDataWithConditionString} returns this
 */
proto.model.QueryHistoricalDataWithConditionString.prototype.clearItemnamesList = function() {
  return this.setItemnamesList([]);
};


/**
 * repeated int32 startTimes = 3;
 * @return {!Array<number>}
 */
proto.model.QueryHistoricalDataWithConditionString.prototype.getStarttimesList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 3));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.model.QueryHistoricalDataWithConditionString} returns this
 */
proto.model.QueryHistoricalDataWithConditionString.prototype.setStarttimesList = function(value) {
  return jspb.Message.setField(this, 3, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.model.QueryHistoricalDataWithConditionString} returns this
 */
proto.model.QueryHistoricalDataWithConditionString.prototype.addStarttimes = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 3, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.QueryHistoricalDataWithConditionString} returns this
 */
proto.model.QueryHistoricalDataWithConditionString.prototype.clearStarttimesList = function() {
  return this.setStarttimesList([]);
};


/**
 * repeated int32 endTimes = 4;
 * @return {!Array<number>}
 */
proto.model.QueryHistoricalDataWithConditionString.prototype.getEndtimesList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 4));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.model.QueryHistoricalDataWithConditionString} returns this
 */
proto.model.QueryHistoricalDataWithConditionString.prototype.setEndtimesList = function(value) {
  return jspb.Message.setField(this, 4, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.model.QueryHistoricalDataWithConditionString} returns this
 */
proto.model.QueryHistoricalDataWithConditionString.prototype.addEndtimes = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 4, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.QueryHistoricalDataWithConditionString} returns this
 */
proto.model.QueryHistoricalDataWithConditionString.prototype.clearEndtimesList = function() {
  return this.setEndtimesList([]);
};


/**
 * repeated int32 intervals = 5;
 * @return {!Array<number>}
 */
proto.model.QueryHistoricalDataWithConditionString.prototype.getIntervalsList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 5));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.model.QueryHistoricalDataWithConditionString} returns this
 */
proto.model.QueryHistoricalDataWithConditionString.prototype.setIntervalsList = function(value) {
  return jspb.Message.setField(this, 5, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.model.QueryHistoricalDataWithConditionString} returns this
 */
proto.model.QueryHistoricalDataWithConditionString.prototype.addIntervals = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 5, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.QueryHistoricalDataWithConditionString} returns this
 */
proto.model.QueryHistoricalDataWithConditionString.prototype.clearIntervalsList = function() {
  return this.setIntervalsList([]);
};


/**
 * optional string filterCondition = 6;
 * @return {string}
 */
proto.model.QueryHistoricalDataWithConditionString.prototype.getFiltercondition = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.QueryHistoricalDataWithConditionString} returns this
 */
proto.model.QueryHistoricalDataWithConditionString.prototype.setFiltercondition = function(value) {
  return jspb.Message.setProto3StringField(this, 6, value);
};


/**
 * repeated DeadZone deadZones = 7;
 * @return {!Array<!proto.model.DeadZone>}
 */
proto.model.QueryHistoricalDataWithConditionString.prototype.getDeadzonesList = function() {
  return /** @type{!Array<!proto.model.DeadZone>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.model.DeadZone, 7));
};


/**
 * @param {!Array<!proto.model.DeadZone>} value
 * @return {!proto.model.QueryHistoricalDataWithConditionString} returns this
*/
proto.model.QueryHistoricalDataWithConditionString.prototype.setDeadzonesList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 7, value);
};


/**
 * @param {!proto.model.DeadZone=} opt_value
 * @param {number=} opt_index
 * @return {!proto.model.DeadZone}
 */
proto.model.QueryHistoricalDataWithConditionString.prototype.addDeadzones = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 7, opt_value, proto.model.DeadZone, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.QueryHistoricalDataWithConditionString} returns this
 */
proto.model.QueryHistoricalDataWithConditionString.prototype.clearDeadzonesList = function() {
  return this.setDeadzonesList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.DeleteHistoricalDataString.repeatedFields_ = [1,2,3,4];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.DeleteHistoricalDataString.prototype.toObject = function(opt_includeInstance) {
  return proto.model.DeleteHistoricalDataString.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.DeleteHistoricalDataString} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.DeleteHistoricalDataString.toObject = function(includeInstance, msg) {
  var f, obj = {
    groupnamesList: (f = jspb.Message.getRepeatedField(msg, 1)) == null ? undefined : f,
    itemnamesList: (f = jspb.Message.getRepeatedField(msg, 2)) == null ? undefined : f,
    starttimesList: (f = jspb.Message.getRepeatedField(msg, 3)) == null ? undefined : f,
    endtimesList: (f = jspb.Message.getRepeatedField(msg, 4)) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.DeleteHistoricalDataString}
 */
proto.model.DeleteHistoricalDataString.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.DeleteHistoricalDataString;
  return proto.model.DeleteHistoricalDataString.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.DeleteHistoricalDataString} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.DeleteHistoricalDataString}
 */
proto.model.DeleteHistoricalDataString.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.addGroupnames(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.addItemnames(value);
      break;
    case 3:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedInt32() : [reader.readInt32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addStarttimes(values[i]);
      }
      break;
    case 4:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedInt32() : [reader.readInt32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addEndtimes(values[i]);
      }
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.DeleteHistoricalDataString.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.DeleteHistoricalDataString.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.DeleteHistoricalDataString} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.DeleteHistoricalDataString.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getGroupnamesList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      1,
      f
    );
  }
  f = message.getItemnamesList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      2,
      f
    );
  }
  f = message.getStarttimesList();
  if (f.length > 0) {
    writer.writePackedInt32(
      3,
      f
    );
  }
  f = message.getEndtimesList();
  if (f.length > 0) {
    writer.writePackedInt32(
      4,
      f
    );
  }
};


/**
 * repeated string groupNames = 1;
 * @return {!Array<string>}
 */
proto.model.DeleteHistoricalDataString.prototype.getGroupnamesList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 1));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.model.DeleteHistoricalDataString} returns this
 */
proto.model.DeleteHistoricalDataString.prototype.setGroupnamesList = function(value) {
  return jspb.Message.setField(this, 1, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.model.DeleteHistoricalDataString} returns this
 */
proto.model.DeleteHistoricalDataString.prototype.addGroupnames = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 1, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.DeleteHistoricalDataString} returns this
 */
proto.model.DeleteHistoricalDataString.prototype.clearGroupnamesList = function() {
  return this.setGroupnamesList([]);
};


/**
 * repeated string itemNames = 2;
 * @return {!Array<string>}
 */
proto.model.DeleteHistoricalDataString.prototype.getItemnamesList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 2));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.model.DeleteHistoricalDataString} returns this
 */
proto.model.DeleteHistoricalDataString.prototype.setItemnamesList = function(value) {
  return jspb.Message.setField(this, 2, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.model.DeleteHistoricalDataString} returns this
 */
proto.model.DeleteHistoricalDataString.prototype.addItemnames = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 2, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.DeleteHistoricalDataString} returns this
 */
proto.model.DeleteHistoricalDataString.prototype.clearItemnamesList = function() {
  return this.setItemnamesList([]);
};


/**
 * repeated int32 startTimes = 3;
 * @return {!Array<number>}
 */
proto.model.DeleteHistoricalDataString.prototype.getStarttimesList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 3));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.model.DeleteHistoricalDataString} returns this
 */
proto.model.DeleteHistoricalDataString.prototype.setStarttimesList = function(value) {
  return jspb.Message.setField(this, 3, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.model.DeleteHistoricalDataString} returns this
 */
proto.model.DeleteHistoricalDataString.prototype.addStarttimes = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 3, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.DeleteHistoricalDataString} returns this
 */
proto.model.DeleteHistoricalDataString.prototype.clearStarttimesList = function() {
  return this.setStarttimesList([]);
};


/**
 * repeated int32 endTimes = 4;
 * @return {!Array<number>}
 */
proto.model.DeleteHistoricalDataString.prototype.getEndtimesList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 4));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.model.DeleteHistoricalDataString} returns this
 */
proto.model.DeleteHistoricalDataString.prototype.setEndtimesList = function(value) {
  return jspb.Message.setField(this, 4, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.model.DeleteHistoricalDataString} returns this
 */
proto.model.DeleteHistoricalDataString.prototype.addEndtimes = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 4, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.DeleteHistoricalDataString} returns this
 */
proto.model.DeleteHistoricalDataString.prototype.clearEndtimesList = function() {
  return this.setEndtimesList([]);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.GdbRealTimeData.prototype.toObject = function(opt_includeInstance) {
  return proto.model.GdbRealTimeData.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.GdbRealTimeData} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.GdbRealTimeData.toObject = function(includeInstance, msg) {
  var f, obj = {
    realtimedata: jspb.Message.getFieldWithDefault(msg, 1, ""),
    times: jspb.Message.getFieldWithDefault(msg, 2, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.GdbRealTimeData}
 */
proto.model.GdbRealTimeData.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.GdbRealTimeData;
  return proto.model.GdbRealTimeData.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.GdbRealTimeData} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.GdbRealTimeData}
 */
proto.model.GdbRealTimeData.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setRealtimedata(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setTimes(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.GdbRealTimeData.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.GdbRealTimeData.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.GdbRealTimeData} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.GdbRealTimeData.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getRealtimedata();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getTimes();
  if (f !== 0) {
    writer.writeInt64(
      2,
      f
    );
  }
};


/**
 * optional string realTimeData = 1;
 * @return {string}
 */
proto.model.GdbRealTimeData.prototype.getRealtimedata = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.GdbRealTimeData} returns this
 */
proto.model.GdbRealTimeData.prototype.setRealtimedata = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional int64 times = 2;
 * @return {number}
 */
proto.model.GdbRealTimeData.prototype.getTimes = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.model.GdbRealTimeData} returns this
 */
proto.model.GdbRealTimeData.prototype.setTimes = function(value) {
  return jspb.Message.setProto3IntField(this, 2, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.GdbHistoricalData.prototype.toObject = function(opt_includeInstance) {
  return proto.model.GdbHistoricalData.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.GdbHistoricalData} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.GdbHistoricalData.toObject = function(includeInstance, msg) {
  var f, obj = {
    historicaldata: jspb.Message.getFieldWithDefault(msg, 1, ""),
    times: jspb.Message.getFieldWithDefault(msg, 2, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.GdbHistoricalData}
 */
proto.model.GdbHistoricalData.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.GdbHistoricalData;
  return proto.model.GdbHistoricalData.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.GdbHistoricalData} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.GdbHistoricalData}
 */
proto.model.GdbHistoricalData.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setHistoricaldata(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setTimes(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.GdbHistoricalData.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.GdbHistoricalData.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.GdbHistoricalData} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.GdbHistoricalData.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getHistoricaldata();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getTimes();
  if (f !== 0) {
    writer.writeInt64(
      2,
      f
    );
  }
};


/**
 * optional string historicalData = 1;
 * @return {string}
 */
proto.model.GdbHistoricalData.prototype.getHistoricaldata = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.GdbHistoricalData} returns this
 */
proto.model.GdbHistoricalData.prototype.setHistoricaldata = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional int64 times = 2;
 * @return {number}
 */
proto.model.GdbHistoricalData.prototype.getTimes = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.model.GdbHistoricalData} returns this
 */
proto.model.GdbHistoricalData.prototype.setTimes = function(value) {
  return jspb.Message.setProto3IntField(this, 2, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.AuthInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.model.AuthInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.AuthInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.AuthInfo.toObject = function(includeInstance, msg) {
  var f, obj = {
    username: jspb.Message.getFieldWithDefault(msg, 1, ""),
    password: jspb.Message.getFieldWithDefault(msg, 2, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.AuthInfo}
 */
proto.model.AuthInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.AuthInfo;
  return proto.model.AuthInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.AuthInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.AuthInfo}
 */
proto.model.AuthInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setUsername(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setPassword(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.AuthInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.AuthInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.AuthInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.AuthInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUsername();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getPassword();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * optional string userName = 1;
 * @return {string}
 */
proto.model.AuthInfo.prototype.getUsername = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.AuthInfo} returns this
 */
proto.model.AuthInfo.prototype.setUsername = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string passWord = 2;
 * @return {string}
 */
proto.model.AuthInfo.prototype.getPassword = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.AuthInfo} returns this
 */
proto.model.AuthInfo.prototype.setPassword = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.UserName.prototype.toObject = function(opt_includeInstance) {
  return proto.model.UserName.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.UserName} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.UserName.toObject = function(includeInstance, msg) {
  var f, obj = {
    name: jspb.Message.getFieldWithDefault(msg, 1, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.UserName}
 */
proto.model.UserName.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.UserName;
  return proto.model.UserName.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.UserName} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.UserName}
 */
proto.model.UserName.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.UserName.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.UserName.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.UserName} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.UserName.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string name = 1;
 * @return {string}
 */
proto.model.UserName.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.UserName} returns this
 */
proto.model.UserName.prototype.setName = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.AddedUserInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.model.AddedUserInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.AddedUserInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.AddedUserInfo.toObject = function(includeInstance, msg) {
  var f, obj = {
    name: jspb.Message.getFieldWithDefault(msg, 1, ""),
    role: jspb.Message.getFieldWithDefault(msg, 2, ""),
    password: jspb.Message.getFieldWithDefault(msg, 3, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.AddedUserInfo}
 */
proto.model.AddedUserInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.AddedUserInfo;
  return proto.model.AddedUserInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.AddedUserInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.AddedUserInfo}
 */
proto.model.AddedUserInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setRole(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setPassword(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.AddedUserInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.AddedUserInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.AddedUserInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.AddedUserInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getRole();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getPassword();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
};


/**
 * optional string name = 1;
 * @return {string}
 */
proto.model.AddedUserInfo.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.AddedUserInfo} returns this
 */
proto.model.AddedUserInfo.prototype.setName = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string role = 2;
 * @return {string}
 */
proto.model.AddedUserInfo.prototype.getRole = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.AddedUserInfo} returns this
 */
proto.model.AddedUserInfo.prototype.setRole = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string passWord = 3;
 * @return {string}
 */
proto.model.AddedUserInfo.prototype.getPassword = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.AddedUserInfo} returns this
 */
proto.model.AddedUserInfo.prototype.setPassword = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.UpdatedUserInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.model.UpdatedUserInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.UpdatedUserInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.UpdatedUserInfo.toObject = function(includeInstance, msg) {
  var f, obj = {
    username: jspb.Message.getFieldWithDefault(msg, 1, ""),
    newusername: jspb.Message.getFieldWithDefault(msg, 2, ""),
    newpassword: jspb.Message.getFieldWithDefault(msg, 3, ""),
    newrole: jspb.Message.getFieldWithDefault(msg, 4, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.UpdatedUserInfo}
 */
proto.model.UpdatedUserInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.UpdatedUserInfo;
  return proto.model.UpdatedUserInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.UpdatedUserInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.UpdatedUserInfo}
 */
proto.model.UpdatedUserInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setUsername(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setNewusername(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setNewpassword(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setNewrole(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.UpdatedUserInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.UpdatedUserInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.UpdatedUserInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.UpdatedUserInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUsername();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getNewusername();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getNewpassword();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getNewrole();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
};


/**
 * optional string userName = 1;
 * @return {string}
 */
proto.model.UpdatedUserInfo.prototype.getUsername = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.UpdatedUserInfo} returns this
 */
proto.model.UpdatedUserInfo.prototype.setUsername = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string newUserName = 2;
 * @return {string}
 */
proto.model.UpdatedUserInfo.prototype.getNewusername = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.UpdatedUserInfo} returns this
 */
proto.model.UpdatedUserInfo.prototype.setNewusername = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string newPassWord = 3;
 * @return {string}
 */
proto.model.UpdatedUserInfo.prototype.getNewpassword = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.UpdatedUserInfo} returns this
 */
proto.model.UpdatedUserInfo.prototype.setNewpassword = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string newRole = 4;
 * @return {string}
 */
proto.model.UpdatedUserInfo.prototype.getNewrole = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.UpdatedUserInfo} returns this
 */
proto.model.UpdatedUserInfo.prototype.setNewrole = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.FileInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.model.FileInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.FileInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.FileInfo.toObject = function(includeInstance, msg) {
  var f, obj = {
    filename: jspb.Message.getFieldWithDefault(msg, 1, ""),
    groupname: jspb.Message.getFieldWithDefault(msg, 2, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.FileInfo}
 */
proto.model.FileInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.FileInfo;
  return proto.model.FileInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.FileInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.FileInfo}
 */
proto.model.FileInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setFilename(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setGroupname(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.FileInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.FileInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.FileInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.FileInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getFilename();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getGroupname();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * optional string fileName = 1;
 * @return {string}
 */
proto.model.FileInfo.prototype.getFilename = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.FileInfo} returns this
 */
proto.model.FileInfo.prototype.setFilename = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string groupName = 2;
 * @return {string}
 */
proto.model.FileInfo.prototype.getGroupname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.FileInfo} returns this
 */
proto.model.FileInfo.prototype.setGroupname = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.HistoryFileInfo.repeatedFields_ = [2,3];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.HistoryFileInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.model.HistoryFileInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.HistoryFileInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.HistoryFileInfo.toObject = function(includeInstance, msg) {
  var f, obj = {
    groupname: jspb.Message.getFieldWithDefault(msg, 4, ""),
    filename: jspb.Message.getFieldWithDefault(msg, 1, ""),
    itemnamesList: (f = jspb.Message.getRepeatedField(msg, 2)) == null ? undefined : f,
    sheetnamesList: (f = jspb.Message.getRepeatedField(msg, 3)) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.HistoryFileInfo}
 */
proto.model.HistoryFileInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.HistoryFileInfo;
  return proto.model.HistoryFileInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.HistoryFileInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.HistoryFileInfo}
 */
proto.model.HistoryFileInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setGroupname(value);
      break;
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setFilename(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.addItemnames(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.addSheetnames(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.HistoryFileInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.HistoryFileInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.HistoryFileInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.HistoryFileInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getGroupname();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getFilename();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getItemnamesList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      2,
      f
    );
  }
  f = message.getSheetnamesList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      3,
      f
    );
  }
};


/**
 * optional string groupName = 4;
 * @return {string}
 */
proto.model.HistoryFileInfo.prototype.getGroupname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.HistoryFileInfo} returns this
 */
proto.model.HistoryFileInfo.prototype.setGroupname = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string fileName = 1;
 * @return {string}
 */
proto.model.HistoryFileInfo.prototype.getFilename = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.HistoryFileInfo} returns this
 */
proto.model.HistoryFileInfo.prototype.setFilename = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * repeated string itemNames = 2;
 * @return {!Array<string>}
 */
proto.model.HistoryFileInfo.prototype.getItemnamesList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 2));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.model.HistoryFileInfo} returns this
 */
proto.model.HistoryFileInfo.prototype.setItemnamesList = function(value) {
  return jspb.Message.setField(this, 2, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.model.HistoryFileInfo} returns this
 */
proto.model.HistoryFileInfo.prototype.addItemnames = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 2, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.HistoryFileInfo} returns this
 */
proto.model.HistoryFileInfo.prototype.clearItemnamesList = function() {
  return this.setItemnamesList([]);
};


/**
 * repeated string sheetNames = 3;
 * @return {!Array<string>}
 */
proto.model.HistoryFileInfo.prototype.getSheetnamesList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 3));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.model.HistoryFileInfo} returns this
 */
proto.model.HistoryFileInfo.prototype.setSheetnamesList = function(value) {
  return jspb.Message.setField(this, 3, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.model.HistoryFileInfo} returns this
 */
proto.model.HistoryFileInfo.prototype.addSheetnames = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 3, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.HistoryFileInfo} returns this
 */
proto.model.HistoryFileInfo.prototype.clearSheetnamesList = function() {
  return this.setSheetnamesList([]);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.QueryLogsInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.model.QueryLogsInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.QueryLogsInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.QueryLogsInfo.toObject = function(includeInstance, msg) {
  var f, obj = {
    level: jspb.Message.getFieldWithDefault(msg, 1, ""),
    starttime: jspb.Message.getFieldWithDefault(msg, 2, ""),
    endtime: jspb.Message.getFieldWithDefault(msg, 3, ""),
    startrow: jspb.Message.getFieldWithDefault(msg, 4, 0),
    rowcount: jspb.Message.getFieldWithDefault(msg, 5, 0),
    name: jspb.Message.getFieldWithDefault(msg, 6, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.QueryLogsInfo}
 */
proto.model.QueryLogsInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.QueryLogsInfo;
  return proto.model.QueryLogsInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.QueryLogsInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.QueryLogsInfo}
 */
proto.model.QueryLogsInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setLevel(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setStarttime(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setEndtime(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setStartrow(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setRowcount(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.QueryLogsInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.QueryLogsInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.QueryLogsInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.QueryLogsInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getLevel();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getStarttime();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getEndtime();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getStartrow();
  if (f !== 0) {
    writer.writeInt32(
      4,
      f
    );
  }
  f = message.getRowcount();
  if (f !== 0) {
    writer.writeInt32(
      5,
      f
    );
  }
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
};


/**
 * optional string level = 1;
 * @return {string}
 */
proto.model.QueryLogsInfo.prototype.getLevel = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.QueryLogsInfo} returns this
 */
proto.model.QueryLogsInfo.prototype.setLevel = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string startTime = 2;
 * @return {string}
 */
proto.model.QueryLogsInfo.prototype.getStarttime = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.QueryLogsInfo} returns this
 */
proto.model.QueryLogsInfo.prototype.setStarttime = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string endTime = 3;
 * @return {string}
 */
proto.model.QueryLogsInfo.prototype.getEndtime = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.QueryLogsInfo} returns this
 */
proto.model.QueryLogsInfo.prototype.setEndtime = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional int32 startRow = 4;
 * @return {number}
 */
proto.model.QueryLogsInfo.prototype.getStartrow = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/**
 * @param {number} value
 * @return {!proto.model.QueryLogsInfo} returns this
 */
proto.model.QueryLogsInfo.prototype.setStartrow = function(value) {
  return jspb.Message.setProto3IntField(this, 4, value);
};


/**
 * optional int32 rowCount = 5;
 * @return {number}
 */
proto.model.QueryLogsInfo.prototype.getRowcount = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/**
 * @param {number} value
 * @return {!proto.model.QueryLogsInfo} returns this
 */
proto.model.QueryLogsInfo.prototype.setRowcount = function(value) {
  return jspb.Message.setProto3IntField(this, 5, value);
};


/**
 * optional string name = 6;
 * @return {string}
 */
proto.model.QueryLogsInfo.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.QueryLogsInfo} returns this
 */
proto.model.QueryLogsInfo.prototype.setName = function(value) {
  return jspb.Message.setProto3StringField(this, 6, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.DeletedLogInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.model.DeletedLogInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.DeletedLogInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.DeletedLogInfo.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, ""),
    starttime: jspb.Message.getFieldWithDefault(msg, 2, ""),
    endtime: jspb.Message.getFieldWithDefault(msg, 3, ""),
    usernamecondition: jspb.Message.getFieldWithDefault(msg, 4, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.DeletedLogInfo}
 */
proto.model.DeletedLogInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.DeletedLogInfo;
  return proto.model.DeletedLogInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.DeletedLogInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.DeletedLogInfo}
 */
proto.model.DeletedLogInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setStarttime(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setEndtime(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setUsernamecondition(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.DeletedLogInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.DeletedLogInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.DeletedLogInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.DeletedLogInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getStarttime();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getEndtime();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getUsernamecondition();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
};


/**
 * optional string id = 1;
 * @return {string}
 */
proto.model.DeletedLogInfo.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.DeletedLogInfo} returns this
 */
proto.model.DeletedLogInfo.prototype.setId = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string startTime = 2;
 * @return {string}
 */
proto.model.DeletedLogInfo.prototype.getStarttime = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.DeletedLogInfo} returns this
 */
proto.model.DeletedLogInfo.prototype.setStarttime = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string endTime = 3;
 * @return {string}
 */
proto.model.DeletedLogInfo.prototype.getEndtime = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.DeletedLogInfo} returns this
 */
proto.model.DeletedLogInfo.prototype.setEndtime = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string userNameCondition = 4;
 * @return {string}
 */
proto.model.DeletedLogInfo.prototype.getUsernamecondition = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.DeletedLogInfo} returns this
 */
proto.model.DeletedLogInfo.prototype.setUsernamecondition = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.QuerySpeedHistoryDataString.repeatedFields_ = [2,3,4];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.QuerySpeedHistoryDataString.prototype.toObject = function(opt_includeInstance) {
  return proto.model.QuerySpeedHistoryDataString.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.QuerySpeedHistoryDataString} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.QuerySpeedHistoryDataString.toObject = function(includeInstance, msg) {
  var f, obj = {
    infotype: jspb.Message.getFieldWithDefault(msg, 5, ""),
    itemname: jspb.Message.getFieldWithDefault(msg, 1, ""),
    starttimesList: (f = jspb.Message.getRepeatedField(msg, 2)) == null ? undefined : f,
    endtimesList: (f = jspb.Message.getRepeatedField(msg, 3)) == null ? undefined : f,
    intervalsList: (f = jspb.Message.getRepeatedField(msg, 4)) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.QuerySpeedHistoryDataString}
 */
proto.model.QuerySpeedHistoryDataString.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.QuerySpeedHistoryDataString;
  return proto.model.QuerySpeedHistoryDataString.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.QuerySpeedHistoryDataString} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.QuerySpeedHistoryDataString}
 */
proto.model.QuerySpeedHistoryDataString.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setInfotype(value);
      break;
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setItemname(value);
      break;
    case 2:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedInt32() : [reader.readInt32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addStarttimes(values[i]);
      }
      break;
    case 3:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedInt32() : [reader.readInt32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addEndtimes(values[i]);
      }
      break;
    case 4:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedInt32() : [reader.readInt32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addIntervals(values[i]);
      }
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.QuerySpeedHistoryDataString.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.QuerySpeedHistoryDataString.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.QuerySpeedHistoryDataString} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.QuerySpeedHistoryDataString.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getInfotype();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getItemname();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getStarttimesList();
  if (f.length > 0) {
    writer.writePackedInt32(
      2,
      f
    );
  }
  f = message.getEndtimesList();
  if (f.length > 0) {
    writer.writePackedInt32(
      3,
      f
    );
  }
  f = message.getIntervalsList();
  if (f.length > 0) {
    writer.writePackedInt32(
      4,
      f
    );
  }
};


/**
 * optional string infoType = 5;
 * @return {string}
 */
proto.model.QuerySpeedHistoryDataString.prototype.getInfotype = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.QuerySpeedHistoryDataString} returns this
 */
proto.model.QuerySpeedHistoryDataString.prototype.setInfotype = function(value) {
  return jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional string itemName = 1;
 * @return {string}
 */
proto.model.QuerySpeedHistoryDataString.prototype.getItemname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.QuerySpeedHistoryDataString} returns this
 */
proto.model.QuerySpeedHistoryDataString.prototype.setItemname = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * repeated int32 startTimes = 2;
 * @return {!Array<number>}
 */
proto.model.QuerySpeedHistoryDataString.prototype.getStarttimesList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 2));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.model.QuerySpeedHistoryDataString} returns this
 */
proto.model.QuerySpeedHistoryDataString.prototype.setStarttimesList = function(value) {
  return jspb.Message.setField(this, 2, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.model.QuerySpeedHistoryDataString} returns this
 */
proto.model.QuerySpeedHistoryDataString.prototype.addStarttimes = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 2, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.QuerySpeedHistoryDataString} returns this
 */
proto.model.QuerySpeedHistoryDataString.prototype.clearStarttimesList = function() {
  return this.setStarttimesList([]);
};


/**
 * repeated int32 endTimes = 3;
 * @return {!Array<number>}
 */
proto.model.QuerySpeedHistoryDataString.prototype.getEndtimesList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 3));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.model.QuerySpeedHistoryDataString} returns this
 */
proto.model.QuerySpeedHistoryDataString.prototype.setEndtimesList = function(value) {
  return jspb.Message.setField(this, 3, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.model.QuerySpeedHistoryDataString} returns this
 */
proto.model.QuerySpeedHistoryDataString.prototype.addEndtimes = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 3, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.QuerySpeedHistoryDataString} returns this
 */
proto.model.QuerySpeedHistoryDataString.prototype.clearEndtimesList = function() {
  return this.setEndtimesList([]);
};


/**
 * repeated int32 intervals = 4;
 * @return {!Array<number>}
 */
proto.model.QuerySpeedHistoryDataString.prototype.getIntervalsList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 4));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.model.QuerySpeedHistoryDataString} returns this
 */
proto.model.QuerySpeedHistoryDataString.prototype.setIntervalsList = function(value) {
  return jspb.Message.setField(this, 4, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.model.QuerySpeedHistoryDataString} returns this
 */
proto.model.QuerySpeedHistoryDataString.prototype.addIntervals = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 4, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.QuerySpeedHistoryDataString} returns this
 */
proto.model.QuerySpeedHistoryDataString.prototype.clearIntervalsList = function() {
  return this.setIntervalsList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.RoutesInfo.repeatedFields_ = [2];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.RoutesInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.model.RoutesInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.RoutesInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.RoutesInfo.toObject = function(includeInstance, msg) {
  var f, obj = {
    name: jspb.Message.getFieldWithDefault(msg, 1, ""),
    routesList: (f = jspb.Message.getRepeatedField(msg, 2)) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.RoutesInfo}
 */
proto.model.RoutesInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.RoutesInfo;
  return proto.model.RoutesInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.RoutesInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.RoutesInfo}
 */
proto.model.RoutesInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.addRoutes(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.RoutesInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.RoutesInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.RoutesInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.RoutesInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getRoutesList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      2,
      f
    );
  }
};


/**
 * optional string name = 1;
 * @return {string}
 */
proto.model.RoutesInfo.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.RoutesInfo} returns this
 */
proto.model.RoutesInfo.prototype.setName = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * repeated string routes = 2;
 * @return {!Array<string>}
 */
proto.model.RoutesInfo.prototype.getRoutesList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 2));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.model.RoutesInfo} returns this
 */
proto.model.RoutesInfo.prototype.setRoutesList = function(value) {
  return jspb.Message.setField(this, 2, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.model.RoutesInfo} returns this
 */
proto.model.RoutesInfo.prototype.addRoutes = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 2, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.RoutesInfo} returns this
 */
proto.model.RoutesInfo.prototype.clearRoutesList = function() {
  return this.setRoutesList([]);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.GdbInfoData.prototype.toObject = function(opt_includeInstance) {
  return proto.model.GdbInfoData.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.GdbInfoData} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.GdbInfoData.toObject = function(includeInstance, msg) {
  var f, obj = {
    info: jspb.Message.getFieldWithDefault(msg, 1, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.GdbInfoData}
 */
proto.model.GdbInfoData.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.GdbInfoData;
  return proto.model.GdbInfoData.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.GdbInfoData} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.GdbInfoData}
 */
proto.model.GdbInfoData.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setInfo(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.GdbInfoData.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.GdbInfoData.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.GdbInfoData} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.GdbInfoData.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getInfo();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string info = 1;
 * @return {string}
 */
proto.model.GdbInfoData.prototype.getInfo = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.GdbInfoData} returns this
 */
proto.model.GdbInfoData.prototype.setInfo = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.UserToken.prototype.toObject = function(opt_includeInstance) {
  return proto.model.UserToken.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.UserToken} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.UserToken.toObject = function(includeInstance, msg) {
  var f, obj = {
    token: jspb.Message.getFieldWithDefault(msg, 1, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.UserToken}
 */
proto.model.UserToken.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.UserToken;
  return proto.model.UserToken.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.UserToken} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.UserToken}
 */
proto.model.UserToken.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setToken(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.UserToken.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.UserToken.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.UserToken} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.UserToken.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getToken();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string token = 1;
 * @return {string}
 */
proto.model.UserToken.prototype.getToken = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.UserToken} returns this
 */
proto.model.UserToken.prototype.setToken = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.UserInfo.repeatedFields_ = [2];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.UserInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.model.UserInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.UserInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.UserInfo.toObject = function(includeInstance, msg) {
  var f, obj = {
    username: jspb.Message.getFieldWithDefault(msg, 1, ""),
    roleList: (f = jspb.Message.getRepeatedField(msg, 2)) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.UserInfo}
 */
proto.model.UserInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.UserInfo;
  return proto.model.UserInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.UserInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.UserInfo}
 */
proto.model.UserInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setUsername(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.addRole(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.UserInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.UserInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.UserInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.UserInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUsername();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getRoleList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      2,
      f
    );
  }
};


/**
 * optional string userName = 1;
 * @return {string}
 */
proto.model.UserInfo.prototype.getUsername = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.UserInfo} returns this
 */
proto.model.UserInfo.prototype.setUsername = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * repeated string role = 2;
 * @return {!Array<string>}
 */
proto.model.UserInfo.prototype.getRoleList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 2));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.model.UserInfo} returns this
 */
proto.model.UserInfo.prototype.setRoleList = function(value) {
  return jspb.Message.setField(this, 2, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.model.UserInfo} returns this
 */
proto.model.UserInfo.prototype.addRole = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 2, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.UserInfo} returns this
 */
proto.model.UserInfo.prototype.clearRoleList = function() {
  return this.setRoleList([]);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.LogsInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.model.LogsInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.LogsInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.LogsInfo.toObject = function(includeInstance, msg) {
  var f, obj = {
    infos: jspb.Message.getFieldWithDefault(msg, 1, ""),
    count: jspb.Message.getFieldWithDefault(msg, 2, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.LogsInfo}
 */
proto.model.LogsInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.LogsInfo;
  return proto.model.LogsInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.LogsInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.LogsInfo}
 */
proto.model.LogsInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setInfos(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setCount(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.LogsInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.LogsInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.LogsInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.LogsInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getInfos();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getCount();
  if (f !== 0) {
    writer.writeInt32(
      2,
      f
    );
  }
};


/**
 * optional string infos = 1;
 * @return {string}
 */
proto.model.LogsInfo.prototype.getInfos = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.LogsInfo} returns this
 */
proto.model.LogsInfo.prototype.setInfos = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional int32 count = 2;
 * @return {number}
 */
proto.model.LogsInfo.prototype.getCount = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.model.LogsInfo} returns this
 */
proto.model.LogsInfo.prototype.setCount = function(value) {
  return jspb.Message.setProto3IntField(this, 2, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.UploadedFileInfo.repeatedFields_ = [2];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.UploadedFileInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.model.UploadedFileInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.UploadedFileInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.UploadedFileInfo.toObject = function(includeInstance, msg) {
  var f, obj = {
    filename: jspb.Message.getFieldWithDefault(msg, 1, ""),
    fileList: (f = jspb.Message.getRepeatedField(msg, 2)) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.UploadedFileInfo}
 */
proto.model.UploadedFileInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.UploadedFileInfo;
  return proto.model.UploadedFileInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.UploadedFileInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.UploadedFileInfo}
 */
proto.model.UploadedFileInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setFilename(value);
      break;
    case 2:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedInt32() : [reader.readInt32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addFile(values[i]);
      }
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.UploadedFileInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.UploadedFileInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.UploadedFileInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.UploadedFileInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getFilename();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getFileList();
  if (f.length > 0) {
    writer.writePackedInt32(
      2,
      f
    );
  }
};


/**
 * optional string fileName = 1;
 * @return {string}
 */
proto.model.UploadedFileInfo.prototype.getFilename = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.UploadedFileInfo} returns this
 */
proto.model.UploadedFileInfo.prototype.setFilename = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * repeated int32 file = 2;
 * @return {!Array<number>}
 */
proto.model.UploadedFileInfo.prototype.getFileList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 2));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.model.UploadedFileInfo} returns this
 */
proto.model.UploadedFileInfo.prototype.setFileList = function(value) {
  return jspb.Message.setField(this, 2, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.model.UploadedFileInfo} returns this
 */
proto.model.UploadedFileInfo.prototype.addFile = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 2, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.UploadedFileInfo} returns this
 */
proto.model.UploadedFileInfo.prototype.clearFileList = function() {
  return this.setFileList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.FileContents.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.FileContents.prototype.toObject = function(opt_includeInstance) {
  return proto.model.FileContents.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.FileContents} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.FileContents.toObject = function(includeInstance, msg) {
  var f, obj = {
    contentsList: (f = jspb.Message.getRepeatedField(msg, 1)) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.FileContents}
 */
proto.model.FileContents.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.FileContents;
  return proto.model.FileContents.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.FileContents} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.FileContents}
 */
proto.model.FileContents.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedInt32() : [reader.readInt32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addContents(values[i]);
      }
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.FileContents.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.FileContents.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.FileContents} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.FileContents.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getContentsList();
  if (f.length > 0) {
    writer.writePackedInt32(
      1,
      f
    );
  }
};


/**
 * repeated int32 contents = 1;
 * @return {!Array<number>}
 */
proto.model.FileContents.prototype.getContentsList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 1));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.model.FileContents} returns this
 */
proto.model.FileContents.prototype.setContentsList = function(value) {
  return jspb.Message.setField(this, 1, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.model.FileContents} returns this
 */
proto.model.FileContents.prototype.addContents = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 1, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.FileContents} returns this
 */
proto.model.FileContents.prototype.clearContentsList = function() {
  return this.setContentsList([]);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.FileSize.prototype.toObject = function(opt_includeInstance) {
  return proto.model.FileSize.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.FileSize} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.FileSize.toObject = function(includeInstance, msg) {
  var f, obj = {
    filesize: jspb.Message.getFieldWithDefault(msg, 1, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.FileSize}
 */
proto.model.FileSize.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.FileSize;
  return proto.model.FileSize.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.FileSize} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.FileSize}
 */
proto.model.FileSize.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setFilesize(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.FileSize.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.FileSize.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.FileSize} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.FileSize.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getFilesize();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string fileSize = 1;
 * @return {string}
 */
proto.model.FileSize.prototype.getFilesize = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.FileSize} returns this
 */
proto.model.FileSize.prototype.setFilesize = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.UserInfos.prototype.toObject = function(opt_includeInstance) {
  return proto.model.UserInfos.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.UserInfos} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.UserInfos.toObject = function(includeInstance, msg) {
  var f, obj = {
    userinfos: jspb.Message.getFieldWithDefault(msg, 1, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.UserInfos}
 */
proto.model.UserInfos.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.UserInfos;
  return proto.model.UserInfos.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.UserInfos} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.UserInfos}
 */
proto.model.UserInfos.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setUserinfos(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.UserInfos.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.UserInfos.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.UserInfos} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.UserInfos.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUserinfos();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string userInfos = 1;
 * @return {string}
 */
proto.model.UserInfos.prototype.getUserinfos = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.UserInfos} returns this
 */
proto.model.UserInfos.prototype.setUserinfos = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.Routes.prototype.toObject = function(opt_includeInstance) {
  return proto.model.Routes.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.Routes} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.Routes.toObject = function(includeInstance, msg) {
  var f, obj = {
    routes: jspb.Message.getFieldWithDefault(msg, 1, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.Routes}
 */
proto.model.Routes.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.Routes;
  return proto.model.Routes.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.Routes} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.Routes}
 */
proto.model.Routes.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setRoutes(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.Routes.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.Routes.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.Routes} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.Routes.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getRoutes();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string routes = 1;
 * @return {string}
 */
proto.model.Routes.prototype.getRoutes = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.Routes} returns this
 */
proto.model.Routes.prototype.setRoutes = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.CheckResult.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.CheckResult.prototype.toObject = function(opt_includeInstance) {
  return proto.model.CheckResult.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.CheckResult} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.CheckResult.toObject = function(includeInstance, msg) {
  var f, obj = {
    resultList: (f = jspb.Message.getRepeatedField(msg, 1)) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.CheckResult}
 */
proto.model.CheckResult.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.CheckResult;
  return proto.model.CheckResult.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.CheckResult} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.CheckResult}
 */
proto.model.CheckResult.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedInt32() : [reader.readInt32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addResult(values[i]);
      }
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.CheckResult.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.CheckResult.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.CheckResult} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.CheckResult.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getResultList();
  if (f.length > 0) {
    writer.writePackedInt32(
      1,
      f
    );
  }
};


/**
 * repeated int32 result = 1;
 * @return {!Array<number>}
 */
proto.model.CheckResult.prototype.getResultList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 1));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.model.CheckResult} returns this
 */
proto.model.CheckResult.prototype.setResultList = function(value) {
  return jspb.Message.setField(this, 1, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.model.CheckResult} returns this
 */
proto.model.CheckResult.prototype.addResult = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 1, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.CheckResult} returns this
 */
proto.model.CheckResult.prototype.clearResultList = function() {
  return this.setResultList([]);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.Code.prototype.toObject = function(opt_includeInstance) {
  return proto.model.Code.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.Code} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.Code.toObject = function(includeInstance, msg) {
  var f, obj = {
    code: jspb.Message.getFieldWithDefault(msg, 1, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.Code}
 */
proto.model.Code.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.Code;
  return proto.model.Code.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.Code} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.Code}
 */
proto.model.Code.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setCode(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.Code.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.Code.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.Code} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.Code.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getCode();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string code = 1;
 * @return {string}
 */
proto.model.Code.prototype.getCode = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.Code} returns this
 */
proto.model.Code.prototype.setCode = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.AddedCalcItemInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.model.AddedCalcItemInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.AddedCalcItemInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.AddedCalcItemInfo.toObject = function(includeInstance, msg) {
  var f, obj = {
    expression: jspb.Message.getFieldWithDefault(msg, 1, ""),
    flag: jspb.Message.getFieldWithDefault(msg, 2, ""),
    duration: jspb.Message.getFieldWithDefault(msg, 3, ""),
    description: jspb.Message.getFieldWithDefault(msg, 4, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.AddedCalcItemInfo}
 */
proto.model.AddedCalcItemInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.AddedCalcItemInfo;
  return proto.model.AddedCalcItemInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.AddedCalcItemInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.AddedCalcItemInfo}
 */
proto.model.AddedCalcItemInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setExpression(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setFlag(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setDuration(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setDescription(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.AddedCalcItemInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.AddedCalcItemInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.AddedCalcItemInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.AddedCalcItemInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getExpression();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getFlag();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getDuration();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getDescription();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
};


/**
 * optional string expression = 1;
 * @return {string}
 */
proto.model.AddedCalcItemInfo.prototype.getExpression = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.AddedCalcItemInfo} returns this
 */
proto.model.AddedCalcItemInfo.prototype.setExpression = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string flag = 2;
 * @return {string}
 */
proto.model.AddedCalcItemInfo.prototype.getFlag = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.AddedCalcItemInfo} returns this
 */
proto.model.AddedCalcItemInfo.prototype.setFlag = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string duration = 3;
 * @return {string}
 */
proto.model.AddedCalcItemInfo.prototype.getDuration = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.AddedCalcItemInfo} returns this
 */
proto.model.AddedCalcItemInfo.prototype.setDuration = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string description = 4;
 * @return {string}
 */
proto.model.AddedCalcItemInfo.prototype.getDescription = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.AddedCalcItemInfo} returns this
 */
proto.model.AddedCalcItemInfo.prototype.setDescription = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.QueryCalcItemsInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.model.QueryCalcItemsInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.QueryCalcItemsInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.QueryCalcItemsInfo.toObject = function(includeInstance, msg) {
  var f, obj = {
    condition: jspb.Message.getFieldWithDefault(msg, 1, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.QueryCalcItemsInfo}
 */
proto.model.QueryCalcItemsInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.QueryCalcItemsInfo;
  return proto.model.QueryCalcItemsInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.QueryCalcItemsInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.QueryCalcItemsInfo}
 */
proto.model.QueryCalcItemsInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setCondition(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.QueryCalcItemsInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.QueryCalcItemsInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.QueryCalcItemsInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.QueryCalcItemsInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getCondition();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string condition = 1;
 * @return {string}
 */
proto.model.QueryCalcItemsInfo.prototype.getCondition = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.QueryCalcItemsInfo} returns this
 */
proto.model.QueryCalcItemsInfo.prototype.setCondition = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.UpdatedCalcInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.model.UpdatedCalcInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.UpdatedCalcInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.UpdatedCalcInfo.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, ""),
    description: jspb.Message.getFieldWithDefault(msg, 2, ""),
    expression: jspb.Message.getFieldWithDefault(msg, 3, ""),
    duration: jspb.Message.getFieldWithDefault(msg, 4, ""),
    updatedtime: jspb.Message.getFieldWithDefault(msg, 5, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.UpdatedCalcInfo}
 */
proto.model.UpdatedCalcInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.UpdatedCalcInfo;
  return proto.model.UpdatedCalcInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.UpdatedCalcInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.UpdatedCalcInfo}
 */
proto.model.UpdatedCalcInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setDescription(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setExpression(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setDuration(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setUpdatedtime(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.UpdatedCalcInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.UpdatedCalcInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.UpdatedCalcInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.UpdatedCalcInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getDescription();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getExpression();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getDuration();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getUpdatedtime();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
};


/**
 * optional string id = 1;
 * @return {string}
 */
proto.model.UpdatedCalcInfo.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.UpdatedCalcInfo} returns this
 */
proto.model.UpdatedCalcInfo.prototype.setId = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string description = 2;
 * @return {string}
 */
proto.model.UpdatedCalcInfo.prototype.getDescription = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.UpdatedCalcInfo} returns this
 */
proto.model.UpdatedCalcInfo.prototype.setDescription = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string expression = 3;
 * @return {string}
 */
proto.model.UpdatedCalcInfo.prototype.getExpression = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.UpdatedCalcInfo} returns this
 */
proto.model.UpdatedCalcInfo.prototype.setExpression = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string duration = 4;
 * @return {string}
 */
proto.model.UpdatedCalcInfo.prototype.getDuration = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.UpdatedCalcInfo} returns this
 */
proto.model.UpdatedCalcInfo.prototype.setDuration = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string updatedTime = 5;
 * @return {string}
 */
proto.model.UpdatedCalcInfo.prototype.getUpdatedtime = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.UpdatedCalcInfo} returns this
 */
proto.model.UpdatedCalcInfo.prototype.setUpdatedtime = function(value) {
  return jspb.Message.setProto3StringField(this, 5, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.CalcId.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.CalcId.prototype.toObject = function(opt_includeInstance) {
  return proto.model.CalcId.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.CalcId} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.CalcId.toObject = function(includeInstance, msg) {
  var f, obj = {
    idList: (f = jspb.Message.getRepeatedField(msg, 1)) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.CalcId}
 */
proto.model.CalcId.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.CalcId;
  return proto.model.CalcId.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.CalcId} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.CalcId}
 */
proto.model.CalcId.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.addId(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.CalcId.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.CalcId.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.CalcId} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.CalcId.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getIdList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      1,
      f
    );
  }
};


/**
 * repeated string id = 1;
 * @return {!Array<string>}
 */
proto.model.CalcId.prototype.getIdList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 1));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.model.CalcId} returns this
 */
proto.model.CalcId.prototype.setIdList = function(value) {
  return jspb.Message.setField(this, 1, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.model.CalcId} returns this
 */
proto.model.CalcId.prototype.addId = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 1, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.CalcId} returns this
 */
proto.model.CalcId.prototype.clearIdList = function() {
  return this.setIdList([]);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.CalculationResult.prototype.toObject = function(opt_includeInstance) {
  return proto.model.CalculationResult.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.CalculationResult} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.CalculationResult.toObject = function(includeInstance, msg) {
  var f, obj = {
    result: jspb.Message.getFieldWithDefault(msg, 1, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.CalculationResult}
 */
proto.model.CalculationResult.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.CalculationResult;
  return proto.model.CalculationResult.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.CalculationResult} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.CalculationResult}
 */
proto.model.CalculationResult.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setResult(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.CalculationResult.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.CalculationResult.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.CalculationResult} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.CalculationResult.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getResult();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string result = 1;
 * @return {string}
 */
proto.model.CalculationResult.prototype.getResult = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.CalculationResult} returns this
 */
proto.model.CalculationResult.prototype.setResult = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.CalcItemsInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.model.CalcItemsInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.CalcItemsInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.CalcItemsInfo.toObject = function(includeInstance, msg) {
  var f, obj = {
    infos: jspb.Message.getFieldWithDefault(msg, 1, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.CalcItemsInfo}
 */
proto.model.CalcItemsInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.CalcItemsInfo;
  return proto.model.CalcItemsInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.CalcItemsInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.CalcItemsInfo}
 */
proto.model.CalcItemsInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setInfos(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.CalcItemsInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.CalcItemsInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.CalcItemsInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.CalcItemsInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getInfos();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string infos = 1;
 * @return {string}
 */
proto.model.CalcItemsInfo.prototype.getInfos = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.CalcItemsInfo} returns this
 */
proto.model.CalcItemsInfo.prototype.setInfos = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.CalculationResults.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.CalculationResults.prototype.toObject = function(opt_includeInstance) {
  return proto.model.CalculationResults.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.CalculationResults} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.CalculationResults.toObject = function(includeInstance, msg) {
  var f, obj = {
    resultsList: jspb.Message.toObjectList(msg.getResultsList(),
    proto.model.CalculationResult.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.CalculationResults}
 */
proto.model.CalculationResults.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.CalculationResults;
  return proto.model.CalculationResults.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.CalculationResults} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.CalculationResults}
 */
proto.model.CalculationResults.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.model.CalculationResult;
      reader.readMessage(value,proto.model.CalculationResult.deserializeBinaryFromReader);
      msg.addResults(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.CalculationResults.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.CalculationResults.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.CalculationResults} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.CalculationResults.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getResultsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.model.CalculationResult.serializeBinaryToWriter
    );
  }
};


/**
 * repeated CalculationResult results = 1;
 * @return {!Array<!proto.model.CalculationResult>}
 */
proto.model.CalculationResults.prototype.getResultsList = function() {
  return /** @type{!Array<!proto.model.CalculationResult>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.model.CalculationResult, 1));
};


/**
 * @param {!Array<!proto.model.CalculationResult>} value
 * @return {!proto.model.CalculationResults} returns this
*/
proto.model.CalculationResults.prototype.setResultsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.model.CalculationResult=} opt_value
 * @param {number=} opt_index
 * @return {!proto.model.CalculationResult}
 */
proto.model.CalculationResults.prototype.addResults = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.model.CalculationResult, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.CalculationResults} returns this
 */
proto.model.CalculationResults.prototype.clearResultsList = function() {
  return this.setResultsList([]);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.TestCalcItemInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.model.TestCalcItemInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.TestCalcItemInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.TestCalcItemInfo.toObject = function(includeInstance, msg) {
  var f, obj = {
    expression: jspb.Message.getFieldWithDefault(msg, 1, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.TestCalcItemInfo}
 */
proto.model.TestCalcItemInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.TestCalcItemInfo;
  return proto.model.TestCalcItemInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.TestCalcItemInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.TestCalcItemInfo}
 */
proto.model.TestCalcItemInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setExpression(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.TestCalcItemInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.TestCalcItemInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.TestCalcItemInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.TestCalcItemInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getExpression();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string expression = 1;
 * @return {string}
 */
proto.model.TestCalcItemInfo.prototype.getExpression = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.model.TestCalcItemInfo} returns this
 */
proto.model.TestCalcItemInfo.prototype.setExpression = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.FloatHistoricalData.repeatedFields_ = [1,2];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.FloatHistoricalData.prototype.toObject = function(opt_includeInstance) {
  return proto.model.FloatHistoricalData.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.FloatHistoricalData} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.FloatHistoricalData.toObject = function(includeInstance, msg) {
  var f, obj = {
    valuesList: (f = jspb.Message.getRepeatedFloatingPointField(msg, 1)) == null ? undefined : f,
    timestampsList: (f = jspb.Message.getRepeatedField(msg, 2)) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.FloatHistoricalData}
 */
proto.model.FloatHistoricalData.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.FloatHistoricalData;
  return proto.model.FloatHistoricalData.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.FloatHistoricalData} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.FloatHistoricalData}
 */
proto.model.FloatHistoricalData.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedFloat() : [reader.readFloat()]);
      for (var i = 0; i < values.length; i++) {
        msg.addValues(values[i]);
      }
      break;
    case 2:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedInt32() : [reader.readInt32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addTimestamps(values[i]);
      }
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.FloatHistoricalData.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.FloatHistoricalData.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.FloatHistoricalData} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.FloatHistoricalData.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getValuesList();
  if (f.length > 0) {
    writer.writePackedFloat(
      1,
      f
    );
  }
  f = message.getTimestampsList();
  if (f.length > 0) {
    writer.writePackedInt32(
      2,
      f
    );
  }
};


/**
 * repeated float values = 1;
 * @return {!Array<number>}
 */
proto.model.FloatHistoricalData.prototype.getValuesList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedFloatingPointField(this, 1));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.model.FloatHistoricalData} returns this
 */
proto.model.FloatHistoricalData.prototype.setValuesList = function(value) {
  return jspb.Message.setField(this, 1, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.model.FloatHistoricalData} returns this
 */
proto.model.FloatHistoricalData.prototype.addValues = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 1, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.FloatHistoricalData} returns this
 */
proto.model.FloatHistoricalData.prototype.clearValuesList = function() {
  return this.setValuesList([]);
};


/**
 * repeated int32 timeStamps = 2;
 * @return {!Array<number>}
 */
proto.model.FloatHistoricalData.prototype.getTimestampsList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 2));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.model.FloatHistoricalData} returns this
 */
proto.model.FloatHistoricalData.prototype.setTimestampsList = function(value) {
  return jspb.Message.setField(this, 2, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.model.FloatHistoricalData} returns this
 */
proto.model.FloatHistoricalData.prototype.addTimestamps = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 2, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.FloatHistoricalData} returns this
 */
proto.model.FloatHistoricalData.prototype.clearTimestampsList = function() {
  return this.setTimestampsList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.IntHistoricalData.repeatedFields_ = [1,2];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.IntHistoricalData.prototype.toObject = function(opt_includeInstance) {
  return proto.model.IntHistoricalData.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.IntHistoricalData} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.IntHistoricalData.toObject = function(includeInstance, msg) {
  var f, obj = {
    valuesList: (f = jspb.Message.getRepeatedField(msg, 1)) == null ? undefined : f,
    timestampsList: (f = jspb.Message.getRepeatedField(msg, 2)) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.IntHistoricalData}
 */
proto.model.IntHistoricalData.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.IntHistoricalData;
  return proto.model.IntHistoricalData.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.IntHistoricalData} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.IntHistoricalData}
 */
proto.model.IntHistoricalData.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedInt32() : [reader.readInt32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addValues(values[i]);
      }
      break;
    case 2:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedInt32() : [reader.readInt32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addTimestamps(values[i]);
      }
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.IntHistoricalData.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.IntHistoricalData.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.IntHistoricalData} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.IntHistoricalData.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getValuesList();
  if (f.length > 0) {
    writer.writePackedInt32(
      1,
      f
    );
  }
  f = message.getTimestampsList();
  if (f.length > 0) {
    writer.writePackedInt32(
      2,
      f
    );
  }
};


/**
 * repeated int32 values = 1;
 * @return {!Array<number>}
 */
proto.model.IntHistoricalData.prototype.getValuesList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 1));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.model.IntHistoricalData} returns this
 */
proto.model.IntHistoricalData.prototype.setValuesList = function(value) {
  return jspb.Message.setField(this, 1, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.model.IntHistoricalData} returns this
 */
proto.model.IntHistoricalData.prototype.addValues = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 1, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.IntHistoricalData} returns this
 */
proto.model.IntHistoricalData.prototype.clearValuesList = function() {
  return this.setValuesList([]);
};


/**
 * repeated int32 timeStamps = 2;
 * @return {!Array<number>}
 */
proto.model.IntHistoricalData.prototype.getTimestampsList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 2));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.model.IntHistoricalData} returns this
 */
proto.model.IntHistoricalData.prototype.setTimestampsList = function(value) {
  return jspb.Message.setField(this, 2, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.model.IntHistoricalData} returns this
 */
proto.model.IntHistoricalData.prototype.addTimestamps = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 2, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.IntHistoricalData} returns this
 */
proto.model.IntHistoricalData.prototype.clearTimestampsList = function() {
  return this.setTimestampsList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.StringHistoricalData.repeatedFields_ = [1,2];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.StringHistoricalData.prototype.toObject = function(opt_includeInstance) {
  return proto.model.StringHistoricalData.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.StringHistoricalData} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.StringHistoricalData.toObject = function(includeInstance, msg) {
  var f, obj = {
    valuesList: (f = jspb.Message.getRepeatedField(msg, 1)) == null ? undefined : f,
    timestampsList: (f = jspb.Message.getRepeatedField(msg, 2)) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.StringHistoricalData}
 */
proto.model.StringHistoricalData.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.StringHistoricalData;
  return proto.model.StringHistoricalData.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.StringHistoricalData} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.StringHistoricalData}
 */
proto.model.StringHistoricalData.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.addValues(value);
      break;
    case 2:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedInt32() : [reader.readInt32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addTimestamps(values[i]);
      }
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.StringHistoricalData.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.StringHistoricalData.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.StringHistoricalData} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.StringHistoricalData.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getValuesList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      1,
      f
    );
  }
  f = message.getTimestampsList();
  if (f.length > 0) {
    writer.writePackedInt32(
      2,
      f
    );
  }
};


/**
 * repeated string values = 1;
 * @return {!Array<string>}
 */
proto.model.StringHistoricalData.prototype.getValuesList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 1));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.model.StringHistoricalData} returns this
 */
proto.model.StringHistoricalData.prototype.setValuesList = function(value) {
  return jspb.Message.setField(this, 1, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.model.StringHistoricalData} returns this
 */
proto.model.StringHistoricalData.prototype.addValues = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 1, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.StringHistoricalData} returns this
 */
proto.model.StringHistoricalData.prototype.clearValuesList = function() {
  return this.setValuesList([]);
};


/**
 * repeated int32 timeStamps = 2;
 * @return {!Array<number>}
 */
proto.model.StringHistoricalData.prototype.getTimestampsList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 2));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.model.StringHistoricalData} returns this
 */
proto.model.StringHistoricalData.prototype.setTimestampsList = function(value) {
  return jspb.Message.setField(this, 2, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.model.StringHistoricalData} returns this
 */
proto.model.StringHistoricalData.prototype.addTimestamps = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 2, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.StringHistoricalData} returns this
 */
proto.model.StringHistoricalData.prototype.clearTimestampsList = function() {
  return this.setTimestampsList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.model.BoolHistoricalData.repeatedFields_ = [1,2];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.model.BoolHistoricalData.prototype.toObject = function(opt_includeInstance) {
  return proto.model.BoolHistoricalData.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.model.BoolHistoricalData} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.BoolHistoricalData.toObject = function(includeInstance, msg) {
  var f, obj = {
    valuesList: (f = jspb.Message.getRepeatedBooleanField(msg, 1)) == null ? undefined : f,
    timestampsList: (f = jspb.Message.getRepeatedField(msg, 2)) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.model.BoolHistoricalData}
 */
proto.model.BoolHistoricalData.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.model.BoolHistoricalData;
  return proto.model.BoolHistoricalData.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.model.BoolHistoricalData} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.model.BoolHistoricalData}
 */
proto.model.BoolHistoricalData.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var values = /** @type {!Array<boolean>} */ (reader.isDelimited() ? reader.readPackedBool() : [reader.readBool()]);
      for (var i = 0; i < values.length; i++) {
        msg.addValues(values[i]);
      }
      break;
    case 2:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedInt32() : [reader.readInt32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addTimestamps(values[i]);
      }
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.model.BoolHistoricalData.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.model.BoolHistoricalData.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.model.BoolHistoricalData} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.model.BoolHistoricalData.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getValuesList();
  if (f.length > 0) {
    writer.writePackedBool(
      1,
      f
    );
  }
  f = message.getTimestampsList();
  if (f.length > 0) {
    writer.writePackedInt32(
      2,
      f
    );
  }
};


/**
 * repeated bool values = 1;
 * @return {!Array<boolean>}
 */
proto.model.BoolHistoricalData.prototype.getValuesList = function() {
  return /** @type {!Array<boolean>} */ (jspb.Message.getRepeatedBooleanField(this, 1));
};


/**
 * @param {!Array<boolean>} value
 * @return {!proto.model.BoolHistoricalData} returns this
 */
proto.model.BoolHistoricalData.prototype.setValuesList = function(value) {
  return jspb.Message.setField(this, 1, value || []);
};


/**
 * @param {boolean} value
 * @param {number=} opt_index
 * @return {!proto.model.BoolHistoricalData} returns this
 */
proto.model.BoolHistoricalData.prototype.addValues = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 1, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.BoolHistoricalData} returns this
 */
proto.model.BoolHistoricalData.prototype.clearValuesList = function() {
  return this.setValuesList([]);
};


/**
 * repeated int32 timeStamps = 2;
 * @return {!Array<number>}
 */
proto.model.BoolHistoricalData.prototype.getTimestampsList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 2));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.model.BoolHistoricalData} returns this
 */
proto.model.BoolHistoricalData.prototype.setTimestampsList = function(value) {
  return jspb.Message.setField(this, 2, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.model.BoolHistoricalData} returns this
 */
proto.model.BoolHistoricalData.prototype.addTimestamps = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 2, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.model.BoolHistoricalData} returns this
 */
proto.model.BoolHistoricalData.prototype.clearTimestampsList = function() {
  return this.setTimestampsList([]);
};


goog.object.extend(exports, proto.model);
