//
//  Generated code. Do not modify.
//  source: nervatura/store.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use userGroupDescriptor instead')
const UserGroup$json = {
  '1': 'UserGroup',
  '2': [
    {'1': 'GROUP_USER', '2': 0},
    {'1': 'GROUP_ADMIN', '2': 1},
    {'1': 'GROUP_GUEST', '2': 2},
  ],
};

/// Descriptor for `UserGroup`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List userGroupDescriptor = $convert.base64Decode(
    'CglVc2VyR3JvdXASDgoKR1JPVVBfVVNFUhAAEg8KC0dST1VQX0FETUlOEAESDwoLR1JPVVBfR1'
    'VFU1QQAg==');

@$core.Deprecated('Use customerTypeDescriptor instead')
const CustomerType$json = {
  '1': 'CustomerType',
  '2': [
    {'1': 'CUSTOMER_COMPANY', '2': 0},
    {'1': 'CUSTOMER_PRIVATE', '2': 1},
    {'1': 'CUSTOMER_OTHER', '2': 2},
    {'1': 'CUSTOMER_OWN', '2': 3},
  ],
};

/// Descriptor for `CustomerType`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List customerTypeDescriptor = $convert.base64Decode(
    'CgxDdXN0b21lclR5cGUSFAoQQ1VTVE9NRVJfQ09NUEFOWRAAEhQKEENVU1RPTUVSX1BSSVZBVE'
    'UQARISCg5DVVNUT01FUl9PVEhFUhACEhAKDENVU1RPTUVSX09XThAD');

@$core.Deprecated('Use linkTypeDescriptor instead')
const LinkType$json = {
  '1': 'LinkType',
  '2': [
    {'1': 'LINK_CUSTOMER', '2': 0},
    {'1': 'LINK_EMPLOYEE', '2': 1},
    {'1': 'LINK_ITEM', '2': 2},
    {'1': 'LINK_MOVEMENT', '2': 3},
    {'1': 'LINK_PAYMENT', '2': 4},
    {'1': 'LINK_PLACE', '2': 5},
    {'1': 'LINK_PRODUCT', '2': 6},
    {'1': 'LINK_PROJECT', '2': 7},
    {'1': 'LINK_TOOL', '2': 8},
    {'1': 'LINK_TRANS', '2': 9},
  ],
};

/// Descriptor for `LinkType`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List linkTypeDescriptor = $convert.base64Decode(
    'CghMaW5rVHlwZRIRCg1MSU5LX0NVU1RPTUVSEAASEQoNTElOS19FTVBMT1lFRRABEg0KCUxJTk'
    'tfSVRFTRACEhEKDUxJTktfTU9WRU1FTlQQAxIQCgxMSU5LX1BBWU1FTlQQBBIOCgpMSU5LX1BM'
    'QUNFEAUSEAoMTElOS19QUk9EVUNUEAYSEAoMTElOS19QUk9KRUNUEAcSDQoJTElOS19UT09MEA'
    'gSDgoKTElOS19UUkFOUxAJ');

@$core.Deprecated('Use logTypeDescriptor instead')
const LogType$json = {
  '1': 'LogType',
  '2': [
    {'1': 'LOG_INSERT', '2': 0},
    {'1': 'LOG_UPDATE', '2': 1},
    {'1': 'LOG_DELETE', '2': 2},
  ],
};

/// Descriptor for `LogType`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List logTypeDescriptor = $convert.base64Decode(
    'CgdMb2dUeXBlEg4KCkxPR19JTlNFUlQQABIOCgpMT0dfVVBEQVRFEAESDgoKTE9HX0RFTEVURR'
    'AC');

@$core.Deprecated('Use movementTypeDescriptor instead')
const MovementType$json = {
  '1': 'MovementType',
  '2': [
    {'1': 'MOVEMENT_INVENTORY', '2': 0},
    {'1': 'MOVEMENT_TOOL', '2': 1},
    {'1': 'MOVEMENT_PLAN', '2': 2},
    {'1': 'MOVEMENT_HEAD', '2': 3},
  ],
};

/// Descriptor for `MovementType`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List movementTypeDescriptor = $convert.base64Decode(
    'CgxNb3ZlbWVudFR5cGUSFgoSTU9WRU1FTlRfSU5WRU5UT1JZEAASEQoNTU9WRU1FTlRfVE9PTB'
    'ABEhEKDU1PVkVNRU5UX1BMQU4QAhIRCg1NT1ZFTUVOVF9IRUFEEAM=');

@$core.Deprecated('Use placeTypeDescriptor instead')
const PlaceType$json = {
  '1': 'PlaceType',
  '2': [
    {'1': 'PLACE_WAREHOUSE', '2': 0},
    {'1': 'PLACE_BANK', '2': 1},
    {'1': 'PLACE_CASH', '2': 2},
    {'1': 'PLACE_OTHER', '2': 3},
  ],
};

/// Descriptor for `PlaceType`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List placeTypeDescriptor = $convert.base64Decode(
    'CglQbGFjZVR5cGUSEwoPUExBQ0VfV0FSRUhPVVNFEAASDgoKUExBQ0VfQkFOSxABEg4KClBMQU'
    'NFX0NBU0gQAhIPCgtQTEFDRV9PVEhFUhAD');

@$core.Deprecated('Use priceTypeDescriptor instead')
const PriceType$json = {
  '1': 'PriceType',
  '2': [
    {'1': 'PRICE_CUSTOMER', '2': 0},
    {'1': 'PRICE_VENDOR', '2': 1},
  ],
};

/// Descriptor for `PriceType`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List priceTypeDescriptor = $convert.base64Decode(
    'CglQcmljZVR5cGUSEgoOUFJJQ0VfQ1VTVE9NRVIQABIQCgxQUklDRV9WRU5ET1IQAQ==');

@$core.Deprecated('Use barcodeTypeDescriptor instead')
const BarcodeType$json = {
  '1': 'BarcodeType',
  '2': [
    {'1': 'BARCODE_CODE_39', '2': 0},
    {'1': 'BARCODE_CODE_128', '2': 1},
    {'1': 'BARCODE_EAN_8', '2': 2},
    {'1': 'BARCODE_EAN_13', '2': 3},
    {'1': 'BARCODE_QR_CODE', '2': 4},
  ],
};

/// Descriptor for `BarcodeType`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List barcodeTypeDescriptor = $convert.base64Decode(
    'CgtCYXJjb2RlVHlwZRITCg9CQVJDT0RFX0NPREVfMzkQABIUChBCQVJDT0RFX0NPREVfMTI4EA'
    'ESEQoNQkFSQ09ERV9FQU5fOBACEhIKDkJBUkNPREVfRUFOXzEzEAMSEwoPQkFSQ09ERV9RUl9D'
    'T0RFEAQ=');

@$core.Deprecated('Use productTypeDescriptor instead')
const ProductType$json = {
  '1': 'ProductType',
  '2': [
    {'1': 'PRODUCT_ITEM', '2': 0},
    {'1': 'PRODUCT_SERVICE', '2': 1},
    {'1': 'PRODUCT_VIRTUAL', '2': 2},
  ],
};

/// Descriptor for `ProductType`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List productTypeDescriptor = $convert.base64Decode(
    'CgtQcm9kdWN0VHlwZRIQCgxQUk9EVUNUX0lURU0QABITCg9QUk9EVUNUX1NFUlZJQ0UQARITCg'
    '9QUk9EVUNUX1ZJUlRVQUwQAg==');

@$core.Deprecated('Use rateTypeDescriptor instead')
const RateType$json = {
  '1': 'RateType',
  '2': [
    {'1': 'RATE_RATE', '2': 0},
    {'1': 'RATE_BUY', '2': 1},
    {'1': 'RATE_SELL', '2': 2},
    {'1': 'RATE_AVERAGE', '2': 3},
  ],
};

/// Descriptor for `RateType`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List rateTypeDescriptor = $convert.base64Decode(
    'CghSYXRlVHlwZRINCglSQVRFX1JBVEUQABIMCghSQVRFX0JVWRABEg0KCVJBVEVfU0VMTBACEh'
    'AKDFJBVEVfQVZFUkFHRRAD');

@$core.Deprecated('Use paidTypeDescriptor instead')
const PaidType$json = {
  '1': 'PaidType',
  '2': [
    {'1': 'PAID_ONLINE', '2': 0},
    {'1': 'PAID_CARD', '2': 1},
    {'1': 'PAID_TRANSFER', '2': 2},
    {'1': 'PAID_CASH', '2': 3},
    {'1': 'PAID_OTHER', '2': 4},
  ],
};

/// Descriptor for `PaidType`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List paidTypeDescriptor = $convert.base64Decode(
    'CghQYWlkVHlwZRIPCgtQQUlEX09OTElORRAAEg0KCVBBSURfQ0FSRBABEhEKDVBBSURfVFJBTl'
    'NGRVIQAhINCglQQUlEX0NBU0gQAxIOCgpQQUlEX09USEVSEAQ=');

@$core.Deprecated('Use transStatusDescriptor instead')
const TransStatus$json = {
  '1': 'TransStatus',
  '2': [
    {'1': 'TRANS_NORMAL', '2': 0},
    {'1': 'TRANS_CANCELLATION', '2': 1},
    {'1': 'TRANS_AMENDMENT', '2': 2},
  ],
};

/// Descriptor for `TransStatus`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List transStatusDescriptor = $convert.base64Decode(
    'CgtUcmFuc1N0YXR1cxIQCgxUUkFOU19OT1JNQUwQABIWChJUUkFOU19DQU5DRUxMQVRJT04QAR'
    'ITCg9UUkFOU19BTUVORE1FTlQQAg==');

@$core.Deprecated('Use transStateDescriptor instead')
const TransState$json = {
  '1': 'TransState',
  '2': [
    {'1': 'TRANS_OK', '2': 0},
    {'1': 'TRANS_NEW', '2': 1},
    {'1': 'TRANS_BACK', '2': 2},
  ],
};

/// Descriptor for `TransState`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List transStateDescriptor = $convert.base64Decode(
    'CgpUcmFuc1N0YXRlEgwKCFRSQU5TX09LEAASDQoJVFJBTlNfTkVXEAESDgoKVFJBTlNfQkFDSx'
    'AC');

@$core.Deprecated('Use directionDescriptor instead')
const Direction$json = {
  '1': 'Direction',
  '2': [
    {'1': 'DIRECTION_OUT', '2': 0},
    {'1': 'DIRECTION_IN', '2': 1},
    {'1': 'DIRECTION_TRANSFER', '2': 2},
  ],
};

/// Descriptor for `Direction`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List directionDescriptor = $convert.base64Decode(
    'CglEaXJlY3Rpb24SEQoNRElSRUNUSU9OX09VVBAAEhAKDERJUkVDVElPTl9JThABEhYKEkRJUk'
    'VDVElPTl9UUkFOU0ZFUhAC');

@$core.Deprecated('Use transTypeDescriptor instead')
const TransType$json = {
  '1': 'TransType',
  '2': [
    {'1': 'TRANS_INVOICE', '2': 0},
    {'1': 'TRANS_RECEIPT', '2': 1},
    {'1': 'TRANS_ORDER', '2': 2},
    {'1': 'TRANS_OFFER', '2': 3},
    {'1': 'TRANS_WORKSHEET', '2': 4},
    {'1': 'TRANS_RENT', '2': 5},
    {'1': 'TRANS_DELIVERY', '2': 6},
    {'1': 'TRANS_INVENTORY', '2': 7},
    {'1': 'TRANS_WAYBILL', '2': 8},
    {'1': 'TRANS_PRODUCTION', '2': 9},
    {'1': 'TRANS_FORMULA', '2': 10},
    {'1': 'TRANS_BANK', '2': 11},
    {'1': 'TRANS_CASH', '2': 12},
  ],
};

/// Descriptor for `TransType`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List transTypeDescriptor = $convert.base64Decode(
    'CglUcmFuc1R5cGUSEQoNVFJBTlNfSU5WT0lDRRAAEhEKDVRSQU5TX1JFQ0VJUFQQARIPCgtUUk'
    'FOU19PUkRFUhACEg8KC1RSQU5TX09GRkVSEAMSEwoPVFJBTlNfV09SS1NIRUVUEAQSDgoKVFJB'
    'TlNfUkVOVBAFEhIKDlRSQU5TX0RFTElWRVJZEAYSEwoPVFJBTlNfSU5WRU5UT1JZEAcSEQoNVF'
    'JBTlNfV0FZQklMTBAIEhQKEFRSQU5TX1BST0RVQ1RJT04QCRIRCg1UUkFOU19GT1JNVUxBEAoS'
    'DgoKVFJBTlNfQkFOSxALEg4KClRSQU5TX0NBU0gQDA==');

@$core.Deprecated('Use configTypeDescriptor instead')
const ConfigType$json = {
  '1': 'ConfigType',
  '2': [
    {'1': 'CONFIG_MAP', '2': 0},
    {'1': 'CONFIG_SHORTCUT', '2': 1},
    {'1': 'CONFIG_MESSAGE', '2': 2},
    {'1': 'CONFIG_PATTERN', '2': 3},
    {'1': 'CONFIG_REPORT', '2': 4},
    {'1': 'CONFIG_PRINT_QUEUE', '2': 5},
    {'1': 'CONFIG_DATA', '2': 6},
  ],
};

/// Descriptor for `ConfigType`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List configTypeDescriptor = $convert.base64Decode(
    'CgpDb25maWdUeXBlEg4KCkNPTkZJR19NQVAQABITCg9DT05GSUdfU0hPUlRDVVQQARISCg5DT0'
    '5GSUdfTUVTU0FHRRACEhIKDkNPTkZJR19QQVRURVJOEAMSEQoNQ09ORklHX1JFUE9SVBAEEhYK'
    'EkNPTkZJR19QUklOVF9RVUVVRRAFEg8KC0NPTkZJR19EQVRBEAY=');

@$core.Deprecated('Use fieldTypeDescriptor instead')
const FieldType$json = {
  '1': 'FieldType',
  '2': [
    {'1': 'FIELD_STRING', '2': 0},
    {'1': 'FIELD_BOOL', '2': 1},
    {'1': 'FIELD_INTEGER', '2': 2},
    {'1': 'FIELD_NUMBER', '2': 3},
    {'1': 'FIELD_DATE', '2': 4},
    {'1': 'FIELD_DATETIME', '2': 5},
    {'1': 'FIELD_MEMO', '2': 6},
    {'1': 'FIELD_ENUM', '2': 7},
    {'1': 'FIELD_URL', '2': 8},
    {'1': 'FIELD_CUSTOMER', '2': 9},
    {'1': 'FIELD_EMPLOYEE', '2': 10},
    {'1': 'FIELD_PLACE', '2': 11},
    {'1': 'FIELD_PRODUCT', '2': 12},
    {'1': 'FIELD_PROJECT', '2': 13},
    {'1': 'FIELD_TOOL', '2': 14},
    {'1': 'FIELD_TRANS', '2': 15},
    {'1': 'FIELD_TRANS_ITEM', '2': 16},
    {'1': 'FIELD_TRANS_MOVEMENT', '2': 17},
    {'1': 'FIELD_TRANS_PAYMENT', '2': 18},
  ],
};

/// Descriptor for `FieldType`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List fieldTypeDescriptor = $convert.base64Decode(
    'CglGaWVsZFR5cGUSEAoMRklFTERfU1RSSU5HEAASDgoKRklFTERfQk9PTBABEhEKDUZJRUxEX0'
    'lOVEVHRVIQAhIQCgxGSUVMRF9OVU1CRVIQAxIOCgpGSUVMRF9EQVRFEAQSEgoORklFTERfREFU'
    'RVRJTUUQBRIOCgpGSUVMRF9NRU1PEAYSDgoKRklFTERfRU5VTRAHEg0KCUZJRUxEX1VSTBAIEh'
    'IKDkZJRUxEX0NVU1RPTUVSEAkSEgoORklFTERfRU1QTE9ZRUUQChIPCgtGSUVMRF9QTEFDRRAL'
    'EhEKDUZJRUxEX1BST0RVQ1QQDBIRCg1GSUVMRF9QUk9KRUNUEA0SDgoKRklFTERfVE9PTBAOEg'
    '8KC0ZJRUxEX1RSQU5TEA8SFAoQRklFTERfVFJBTlNfSVRFTRAQEhgKFEZJRUxEX1RSQU5TX01P'
    'VkVNRU5UEBESFwoTRklFTERfVFJBTlNfUEFZTUVOVBAS');

@$core.Deprecated('Use mapFilterDescriptor instead')
const MapFilter$json = {
  '1': 'MapFilter',
  '2': [
    {'1': 'FILTER_ADDRESS', '2': 0},
    {'1': 'FILTER_BARCODE', '2': 1},
    {'1': 'FILTER_CONTACT', '2': 2},
    {'1': 'FILTER_CURRENCY', '2': 3},
    {'1': 'FILTER_CUSTOMER', '2': 4},
    {'1': 'FILTER_EMPLOYEE', '2': 5},
    {'1': 'FILTER_EVENT', '2': 6},
    {'1': 'FILTER_ITEM', '2': 7},
    {'1': 'FILTER_MOVEMENT', '2': 8},
    {'1': 'FILTER_PAYMENT', '2': 9},
    {'1': 'FILTER_PLACE', '2': 10},
    {'1': 'FILTER_PRICE', '2': 11},
    {'1': 'FILTER_PRODUCT', '2': 12},
    {'1': 'FILTER_PROJECT', '2': 13},
    {'1': 'FILTER_RATE', '2': 14},
    {'1': 'FILTER_TAX', '2': 15},
    {'1': 'FILTER_TOOL', '2': 16},
    {'1': 'FILTER_USER', '2': 17},
    {'1': 'FILTER_TRANS', '2': 18},
    {'1': 'FILTER_INVOICE', '2': 19},
    {'1': 'FILTER_RECEIPT', '2': 20},
    {'1': 'FILTER_ORDER', '2': 21},
    {'1': 'FILTER_OFFER', '2': 22},
    {'1': 'FILTER_WORKSHEET', '2': 23},
    {'1': 'FILTER_RENT', '2': 24},
    {'1': 'FILTER_DELIVERY', '2': 25},
    {'1': 'FILTER_INVENTORY', '2': 26},
    {'1': 'FILTER_WAYBILL', '2': 27},
    {'1': 'FILTER_PRODUCTION', '2': 28},
    {'1': 'FILTER_FORMULA', '2': 29},
    {'1': 'FILTER_BANK', '2': 30},
    {'1': 'FILTER_CASH', '2': 31},
  ],
};

/// Descriptor for `MapFilter`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List mapFilterDescriptor = $convert.base64Decode(
    'CglNYXBGaWx0ZXISEgoORklMVEVSX0FERFJFU1MQABISCg5GSUxURVJfQkFSQ09ERRABEhIKDk'
    'ZJTFRFUl9DT05UQUNUEAISEwoPRklMVEVSX0NVUlJFTkNZEAMSEwoPRklMVEVSX0NVU1RPTUVS'
    'EAQSEwoPRklMVEVSX0VNUExPWUVFEAUSEAoMRklMVEVSX0VWRU5UEAYSDwoLRklMVEVSX0lURU'
    '0QBxITCg9GSUxURVJfTU9WRU1FTlQQCBISCg5GSUxURVJfUEFZTUVOVBAJEhAKDEZJTFRFUl9Q'
    'TEFDRRAKEhAKDEZJTFRFUl9QUklDRRALEhIKDkZJTFRFUl9QUk9EVUNUEAwSEgoORklMVEVSX1'
    'BST0pFQ1QQDRIPCgtGSUxURVJfUkFURRAOEg4KCkZJTFRFUl9UQVgQDxIPCgtGSUxURVJfVE9P'
    'TBAQEg8KC0ZJTFRFUl9VU0VSEBESEAoMRklMVEVSX1RSQU5TEBISEgoORklMVEVSX0lOVk9JQ0'
    'UQExISCg5GSUxURVJfUkVDRUlQVBAUEhAKDEZJTFRFUl9PUkRFUhAVEhAKDEZJTFRFUl9PRkZF'
    'UhAWEhQKEEZJTFRFUl9XT1JLU0hFRVQQFxIPCgtGSUxURVJfUkVOVBAYEhMKD0ZJTFRFUl9ERU'
    'xJVkVSWRAZEhQKEEZJTFRFUl9JTlZFTlRPUlkQGhISCg5GSUxURVJfV0FZQklMTBAbEhUKEUZJ'
    'TFRFUl9QUk9EVUNUSU9OEBwSEgoORklMVEVSX0ZPUk1VTEEQHRIPCgtGSUxURVJfQkFOSxAeEg'
    '8KC0ZJTFRFUl9DQVNIEB8=');

@$core.Deprecated('Use shortcutMethodDescriptor instead')
const ShortcutMethod$json = {
  '1': 'ShortcutMethod',
  '2': [
    {'1': 'METHOD_GET', '2': 0},
    {'1': 'METHOD_POST', '2': 1},
  ],
};

/// Descriptor for `ShortcutMethod`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List shortcutMethodDescriptor = $convert.base64Decode(
    'Cg5TaG9ydGN1dE1ldGhvZBIOCgpNRVRIT0RfR0VUEAASDwoLTUVUSE9EX1BPU1QQAQ==');

@$core.Deprecated('Use shortcutFieldDescriptor instead')
const ShortcutField$json = {
  '1': 'ShortcutField',
  '2': [
    {'1': 'SHORTCUT_STRING', '2': 0},
    {'1': 'SHORTCUT_BOOL', '2': 1},
    {'1': 'SHORTCUT_INTEGER', '2': 2},
    {'1': 'SHORTCUT_NUMBER', '2': 3},
    {'1': 'SHORTCUT_DATE', '2': 4},
  ],
};

/// Descriptor for `ShortcutField`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List shortcutFieldDescriptor = $convert.base64Decode(
    'Cg1TaG9ydGN1dEZpZWxkEhMKD1NIT1JUQ1VUX1NUUklORxAAEhEKDVNIT1JUQ1VUX0JPT0wQAR'
    'IUChBTSE9SVENVVF9JTlRFR0VSEAISEwoPU0hPUlRDVVRfTlVNQkVSEAMSEQoNU0hPUlRDVVRf'
    'REFURRAE');

@$core.Deprecated('Use fileTypeDescriptor instead')
const FileType$json = {
  '1': 'FileType',
  '2': [
    {'1': 'FILE_PDF', '2': 0},
    {'1': 'FILE_CSV', '2': 1},
  ],
};

/// Descriptor for `FileType`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List fileTypeDescriptor = $convert.base64Decode(
    'CghGaWxlVHlwZRIMCghGSUxFX1BERhAAEgwKCEZJTEVfQ1NWEAE=');

@$core.Deprecated('Use modelDescriptor instead')
const Model$json = {
  '1': 'Model',
  '2': [
    {'1': 'AUTH', '2': 0},
    {'1': 'CONFIG', '2': 1},
    {'1': 'CURRENCY', '2': 2},
    {'1': 'CUSTOMER', '2': 3},
    {'1': 'EMPLOYEE', '2': 4},
    {'1': 'ITEM', '2': 5},
    {'1': 'LINK', '2': 6},
    {'1': 'MOVEMENT', '2': 7},
    {'1': 'PAYMENT', '2': 8},
    {'1': 'PLACE', '2': 9},
    {'1': 'PRICE', '2': 10},
    {'1': 'PRODUCT', '2': 11},
    {'1': 'PROJECT', '2': 12},
    {'1': 'RATE', '2': 13},
    {'1': 'TAX', '2': 14},
    {'1': 'TOOL', '2': 15},
    {'1': 'TRANS', '2': 16},
  ],
};

/// Descriptor for `Model`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List modelDescriptor = $convert.base64Decode(
    'CgVNb2RlbBIICgRBVVRIEAASCgoGQ09ORklHEAESDAoIQ1VSUkVOQ1kQAhIMCghDVVNUT01FUh'
    'ADEgwKCEVNUExPWUVFEAQSCAoESVRFTRAFEggKBExJTksQBhIMCghNT1ZFTUVOVBAHEgsKB1BB'
    'WU1FTlQQCBIJCgVQTEFDRRAJEgkKBVBSSUNFEAoSCwoHUFJPRFVDVBALEgsKB1BST0pFQ1QQDB'
    'IICgRSQVRFEA0SBwoDVEFYEA4SCAoEVE9PTBAPEgkKBVRSQU5TEBA=');

@$core.Deprecated('Use viewNameDescriptor instead')
const ViewName$json = {
  '1': 'ViewName',
  '2': [
    {'1': 'VIEW_CONFIG_DATA', '2': 0},
    {'1': 'VIEW_CONFIG_MAP', '2': 1},
    {'1': 'VIEW_CONFIG_MESSAGE', '2': 2},
    {'1': 'VIEW_CONFIG_PATTERN', '2': 3},
    {'1': 'VIEW_CONFIG_PRINT_QUEUE', '2': 4},
    {'1': 'VIEW_CONFIG_REPORT', '2': 5},
    {'1': 'VIEW_CONFIG_SHORTCUT', '2': 6},
    {'1': 'VIEW_CURRENCY_MAP', '2': 7},
    {'1': 'VIEW_CURRENCY_TAGS', '2': 8},
    {'1': 'VIEW_CURRENCY_VIEW', '2': 9},
    {'1': 'VIEW_CUSTOMER_ADDRESSES', '2': 10},
    {'1': 'VIEW_CUSTOMER_CONTACTS', '2': 11},
    {'1': 'VIEW_CUSTOMER_EVENTS', '2': 12},
    {'1': 'VIEW_CUSTOMER_MAP', '2': 13},
    {'1': 'VIEW_CUSTOMER_TAGS', '2': 14},
    {'1': 'VIEW_CUSTOMER_VIEW', '2': 15},
    {'1': 'VIEW_EMPLOYEE_EVENTS', '2': 16},
    {'1': 'VIEW_EMPLOYEE_MAP', '2': 17},
    {'1': 'VIEW_EMPLOYEE_TAGS', '2': 18},
    {'1': 'VIEW_EMPLOYEE_VIEW', '2': 19},
    {'1': 'VIEW_ITEM_MAP', '2': 20},
    {'1': 'VIEW_ITEM_TAGS', '2': 21},
    {'1': 'VIEW_ITEM_VIEW', '2': 22},
    {'1': 'VIEW_LINK_MAP', '2': 23},
    {'1': 'VIEW_LINK_TAGS', '2': 24},
    {'1': 'VIEW_LINK_VIEW', '2': 25},
    {'1': 'VIEW_MOVEMENT_MAP', '2': 26},
    {'1': 'VIEW_MOVEMENT_TAGS', '2': 27},
    {'1': 'VIEW_MOVEMENT_STOCK', '2': 28},
    {'1': 'VIEW_MOVEMENT_INVENTORY', '2': 29},
    {'1': 'VIEW_MOVEMENT_WAYBILL', '2': 30},
    {'1': 'VIEW_MOVEMENT_FORMULA', '2': 31},
    {'1': 'VIEW_MOVEMENT_VIEW', '2': 32},
    {'1': 'VIEW_PAYMENT_MAP', '2': 33},
    {'1': 'VIEW_PAYMENT_INVOICE', '2': 34},
    {'1': 'VIEW_PAYMENT_TAGS', '2': 35},
    {'1': 'VIEW_PAYMENT_VIEW', '2': 36},
    {'1': 'VIEW_PLACE_MAP', '2': 37},
    {'1': 'VIEW_PLACE_TAGS', '2': 38},
    {'1': 'VIEW_PLACE_VIEW', '2': 39},
    {'1': 'VIEW_PRICE_MAP', '2': 40},
    {'1': 'VIEW_PRICE_TAGS', '2': 41},
    {'1': 'VIEW_PRICE_VIEW', '2': 42},
    {'1': 'VIEW_PRODUCT_MAP', '2': 43},
    {'1': 'VIEW_PRODUCT_TAGS', '2': 44},
    {'1': 'VIEW_PRODUCT_VIEW', '2': 45},
    {'1': 'VIEW_PRODUCT_COMPONENTS', '2': 46},
    {'1': 'VIEW_PROJECT_MAP', '2': 47},
    {'1': 'VIEW_PROJECT_TAGS', '2': 48},
    {'1': 'VIEW_PROJECT_VIEW', '2': 49},
    {'1': 'VIEW_RATE_MAP', '2': 50},
    {'1': 'VIEW_RATE_TAGS', '2': 51},
    {'1': 'VIEW_RATE_VIEW', '2': 52},
    {'1': 'VIEW_TAX_MAP', '2': 53},
    {'1': 'VIEW_TAX_TAGS', '2': 54},
    {'1': 'VIEW_TAX_VIEW', '2': 55},
    {'1': 'VIEW_TOOL_MAP', '2': 56},
    {'1': 'VIEW_TOOL_TAGS', '2': 57},
    {'1': 'VIEW_TOOL_VIEW', '2': 58},
    {'1': 'VIEW_TRANS_MAP', '2': 59},
    {'1': 'VIEW_TRANS_TAGS', '2': 60},
    {'1': 'VIEW_TRANS_VIEW', '2': 61},
  ],
};

/// Descriptor for `ViewName`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List viewNameDescriptor = $convert.base64Decode(
    'CghWaWV3TmFtZRIUChBWSUVXX0NPTkZJR19EQVRBEAASEwoPVklFV19DT05GSUdfTUFQEAESFw'
    'oTVklFV19DT05GSUdfTUVTU0FHRRACEhcKE1ZJRVdfQ09ORklHX1BBVFRFUk4QAxIbChdWSUVX'
    'X0NPTkZJR19QUklOVF9RVUVVRRAEEhYKElZJRVdfQ09ORklHX1JFUE9SVBAFEhgKFFZJRVdfQ0'
    '9ORklHX1NIT1JUQ1VUEAYSFQoRVklFV19DVVJSRU5DWV9NQVAQBxIWChJWSUVXX0NVUlJFTkNZ'
    'X1RBR1MQCBIWChJWSUVXX0NVUlJFTkNZX1ZJRVcQCRIbChdWSUVXX0NVU1RPTUVSX0FERFJFU1'
    'NFUxAKEhoKFlZJRVdfQ1VTVE9NRVJfQ09OVEFDVFMQCxIYChRWSUVXX0NVU1RPTUVSX0VWRU5U'
    'UxAMEhUKEVZJRVdfQ1VTVE9NRVJfTUFQEA0SFgoSVklFV19DVVNUT01FUl9UQUdTEA4SFgoSVk'
    'lFV19DVVNUT01FUl9WSUVXEA8SGAoUVklFV19FTVBMT1lFRV9FVkVOVFMQEBIVChFWSUVXX0VN'
    'UExPWUVFX01BUBAREhYKElZJRVdfRU1QTE9ZRUVfVEFHUxASEhYKElZJRVdfRU1QTE9ZRUVfVk'
    'lFVxATEhEKDVZJRVdfSVRFTV9NQVAQFBISCg5WSUVXX0lURU1fVEFHUxAVEhIKDlZJRVdfSVRF'
    'TV9WSUVXEBYSEQoNVklFV19MSU5LX01BUBAXEhIKDlZJRVdfTElOS19UQUdTEBgSEgoOVklFV1'
    '9MSU5LX1ZJRVcQGRIVChFWSUVXX01PVkVNRU5UX01BUBAaEhYKElZJRVdfTU9WRU1FTlRfVEFH'
    'UxAbEhcKE1ZJRVdfTU9WRU1FTlRfU1RPQ0sQHBIbChdWSUVXX01PVkVNRU5UX0lOVkVOVE9SWR'
    'AdEhkKFVZJRVdfTU9WRU1FTlRfV0FZQklMTBAeEhkKFVZJRVdfTU9WRU1FTlRfRk9STVVMQRAf'
    'EhYKElZJRVdfTU9WRU1FTlRfVklFVxAgEhQKEFZJRVdfUEFZTUVOVF9NQVAQIRIYChRWSUVXX1'
    'BBWU1FTlRfSU5WT0lDRRAiEhUKEVZJRVdfUEFZTUVOVF9UQUdTECMSFQoRVklFV19QQVlNRU5U'
    'X1ZJRVcQJBISCg5WSUVXX1BMQUNFX01BUBAlEhMKD1ZJRVdfUExBQ0VfVEFHUxAmEhMKD1ZJRV'
    'dfUExBQ0VfVklFVxAnEhIKDlZJRVdfUFJJQ0VfTUFQECgSEwoPVklFV19QUklDRV9UQUdTECkS'
    'EwoPVklFV19QUklDRV9WSUVXECoSFAoQVklFV19QUk9EVUNUX01BUBArEhUKEVZJRVdfUFJPRF'
    'VDVF9UQUdTECwSFQoRVklFV19QUk9EVUNUX1ZJRVcQLRIbChdWSUVXX1BST0RVQ1RfQ09NUE9O'
    'RU5UUxAuEhQKEFZJRVdfUFJPSkVDVF9NQVAQLxIVChFWSUVXX1BST0pFQ1RfVEFHUxAwEhUKEV'
    'ZJRVdfUFJPSkVDVF9WSUVXEDESEQoNVklFV19SQVRFX01BUBAyEhIKDlZJRVdfUkFURV9UQUdT'
    'EDMSEgoOVklFV19SQVRFX1ZJRVcQNBIQCgxWSUVXX1RBWF9NQVAQNRIRCg1WSUVXX1RBWF9UQU'
    'dTEDYSEQoNVklFV19UQVhfVklFVxA3EhEKDVZJRVdfVE9PTF9NQVAQOBISCg5WSUVXX1RPT0xf'
    'VEFHUxA5EhIKDlZJRVdfVE9PTF9WSUVXEDoSEgoOVklFV19UUkFOU19NQVAQOxITCg9WSUVXX1'
    'RSQU5TX1RBR1MQPBITCg9WSUVXX1RSQU5TX1ZJRVcQPQ==');

@$core.Deprecated('Use bookmarkTypeDescriptor instead')
const BookmarkType$json = {
  '1': 'BookmarkType',
  '2': [
    {'1': 'BOOKMARK_BROWSER', '2': 0},
    {'1': 'BOOKMARK_EDITOR', '2': 1},
  ],
};

/// Descriptor for `BookmarkType`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List bookmarkTypeDescriptor = $convert.base64Decode(
    'CgxCb29rbWFya1R5cGUSFAoQQk9PS01BUktfQlJPV1NFUhAAEhMKD0JPT0tNQVJLX0VESVRPUh'
    'AB');

@$core.Deprecated('Use requestEmptyDescriptor instead')
const RequestEmpty$json = {
  '1': 'RequestEmpty',
};

/// Descriptor for `RequestEmpty`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List requestEmptyDescriptor = $convert.base64Decode(
    'CgxSZXF1ZXN0RW1wdHk=');

@$core.Deprecated('Use responseStatusDescriptor instead')
const ResponseStatus$json = {
  '1': 'ResponseStatus',
  '2': [
    {'1': 'success', '3': 1, '4': 1, '5': 8, '10': 'success'},
    {'1': 'message', '3': 2, '4': 1, '5': 9, '10': 'message'},
  ],
};

/// Descriptor for `ResponseStatus`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List responseStatusDescriptor = $convert.base64Decode(
    'Cg5SZXNwb25zZVN0YXR1cxIYCgdzdWNjZXNzGAEgASgIUgdzdWNjZXNzEhgKB21lc3NhZ2UYAi'
    'ABKAlSB21lc3NhZ2U=');

@$core.Deprecated('Use requestAuthLoginDescriptor instead')
const RequestAuthLogin$json = {
  '1': 'RequestAuthLogin',
  '2': [
    {'1': 'user_name', '3': 1, '4': 1, '5': 9, '10': 'userName'},
    {'1': 'password', '3': 2, '4': 1, '5': 9, '10': 'password'},
    {'1': 'database', '3': 3, '4': 1, '5': 9, '10': 'database'},
  ],
};

/// Descriptor for `RequestAuthLogin`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List requestAuthLoginDescriptor = $convert.base64Decode(
    'ChBSZXF1ZXN0QXV0aExvZ2luEhsKCXVzZXJfbmFtZRgBIAEoCVIIdXNlck5hbWUSGgoIcGFzc3'
    'dvcmQYAiABKAlSCHBhc3N3b3JkEhoKCGRhdGFiYXNlGAMgASgJUghkYXRhYmFzZQ==');

@$core.Deprecated('Use responseAuthLoginDescriptor instead')
const ResponseAuthLogin$json = {
  '1': 'ResponseAuthLogin',
  '2': [
    {'1': 'token', '3': 1, '4': 1, '5': 9, '10': 'token'},
    {'1': 'version', '3': 2, '4': 1, '5': 9, '10': 'version'},
  ],
};

/// Descriptor for `ResponseAuthLogin`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List responseAuthLoginDescriptor = $convert.base64Decode(
    'ChFSZXNwb25zZUF1dGhMb2dpbhIUCgV0b2tlbhgBIAEoCVIFdG9rZW4SGAoHdmVyc2lvbhgCIA'
    'EoCVIHdmVyc2lvbg==');

@$core.Deprecated('Use requestPasswordChangeDescriptor instead')
const RequestPasswordChange$json = {
  '1': 'RequestPasswordChange',
  '2': [
    {'1': 'password', '3': 1, '4': 1, '5': 9, '10': 'password'},
    {'1': 'confirm', '3': 2, '4': 1, '5': 9, '10': 'confirm'},
  ],
};

/// Descriptor for `RequestPasswordChange`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List requestPasswordChangeDescriptor = $convert.base64Decode(
    'ChVSZXF1ZXN0UGFzc3dvcmRDaGFuZ2USGgoIcGFzc3dvcmQYASABKAlSCHBhc3N3b3JkEhgKB2'
    'NvbmZpcm0YAiABKAlSB2NvbmZpcm0=');

@$core.Deprecated('Use requestTokenDecodeDescriptor instead')
const RequestTokenDecode$json = {
  '1': 'RequestTokenDecode',
  '2': [
    {'1': 'token', '3': 1, '4': 1, '5': 9, '10': 'token'},
  ],
};

/// Descriptor for `RequestTokenDecode`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List requestTokenDecodeDescriptor = $convert.base64Decode(
    'ChJSZXF1ZXN0VG9rZW5EZWNvZGUSFAoFdG9rZW4YASABKAlSBXRva2Vu');

@$core.Deprecated('Use responseTokenDecodeDescriptor instead')
const ResponseTokenDecode$json = {
  '1': 'ResponseTokenDecode',
  '2': [
    {'1': 'code', '3': 1, '4': 1, '5': 9, '10': 'code'},
    {'1': 'user_name', '3': 2, '4': 1, '5': 9, '10': 'userName'},
    {'1': 'database', '3': 3, '4': 1, '5': 9, '10': 'database'},
    {'1': 'exp', '3': 4, '4': 1, '5': 9, '10': 'exp'},
    {'1': 'iss', '3': 5, '4': 1, '5': 9, '10': 'iss'},
  ],
};

/// Descriptor for `ResponseTokenDecode`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List responseTokenDecodeDescriptor = $convert.base64Decode(
    'ChNSZXNwb25zZVRva2VuRGVjb2RlEhIKBGNvZGUYASABKAlSBGNvZGUSGwoJdXNlcl9uYW1lGA'
    'IgASgJUgh1c2VyTmFtZRIaCghkYXRhYmFzZRgDIAEoCVIIZGF0YWJhc2USEAoDZXhwGAQgASgJ'
    'UgNleHASEAoDaXNzGAUgASgJUgNpc3M=');

@$core.Deprecated('Use requestGetDescriptor instead')
const RequestGet$json = {
  '1': 'RequestGet',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    {'1': 'code', '3': 2, '4': 1, '5': 9, '10': 'code'},
  ],
};

/// Descriptor for `RequestGet`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List requestGetDescriptor = $convert.base64Decode(
    'CgpSZXF1ZXN0R2V0Eg4KAmlkGAEgASgDUgJpZBISCgRjb2RlGAIgASgJUgRjb2Rl');

@$core.Deprecated('Use requestDeleteDescriptor instead')
const RequestDelete$json = {
  '1': 'RequestDelete',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    {'1': 'code', '3': 2, '4': 1, '5': 9, '10': 'code'},
    {'1': 'model', '3': 3, '4': 1, '5': 14, '6': '.nervatura.Model', '10': 'model'},
  ],
};

/// Descriptor for `RequestDelete`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List requestDeleteDescriptor = $convert.base64Decode(
    'Cg1SZXF1ZXN0RGVsZXRlEg4KAmlkGAEgASgDUgJpZBISCgRjb2RlGAIgASgJUgRjb2RlEiYKBW'
    '1vZGVsGAMgASgOMhAubmVydmF0dXJhLk1vZGVsUgVtb2RlbA==');

@$core.Deprecated('Use requestQueryFilterDescriptor instead')
const RequestQueryFilter$json = {
  '1': 'RequestQueryFilter',
  '2': [
    {'1': 'field', '3': 1, '4': 1, '5': 9, '10': 'field'},
    {'1': 'value', '3': 2, '4': 1, '5': 9, '10': 'value'},
  ],
};

/// Descriptor for `RequestQueryFilter`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List requestQueryFilterDescriptor = $convert.base64Decode(
    'ChJSZXF1ZXN0UXVlcnlGaWx0ZXISFAoFZmllbGQYASABKAlSBWZpZWxkEhQKBXZhbHVlGAIgAS'
    'gJUgV2YWx1ZQ==');

@$core.Deprecated('Use requestQueryDescriptor instead')
const RequestQuery$json = {
  '1': 'RequestQuery',
  '2': [
    {'1': 'filters', '3': 1, '4': 3, '5': 11, '6': '.nervatura.RequestQueryFilter', '10': 'filters'},
    {'1': 'limit', '3': 2, '4': 1, '5': 3, '10': 'limit'},
    {'1': 'offset', '3': 3, '4': 1, '5': 3, '10': 'offset'},
  ],
};

/// Descriptor for `RequestQuery`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List requestQueryDescriptor = $convert.base64Decode(
    'CgxSZXF1ZXN0UXVlcnkSNwoHZmlsdGVycxgBIAMoCzIdLm5lcnZhdHVyYS5SZXF1ZXN0UXVlcn'
    'lGaWx0ZXJSB2ZpbHRlcnMSFAoFbGltaXQYAiABKANSBWxpbWl0EhYKBm9mZnNldBgDIAEoA1IG'
    'b2Zmc2V0');

@$core.Deprecated('Use requestFunctionDescriptor instead')
const RequestFunction$json = {
  '1': 'RequestFunction',
  '2': [
    {'1': 'function', '3': 1, '4': 1, '5': 9, '10': 'function'},
    {'1': 'args', '3': 2, '4': 3, '5': 11, '6': '.nervatura.RequestFunction.ArgsEntry', '10': 'args'},
  ],
  '3': [RequestFunction_ArgsEntry$json],
};

@$core.Deprecated('Use requestFunctionDescriptor instead')
const RequestFunction_ArgsEntry$json = {
  '1': 'ArgsEntry',
  '2': [
    {'1': 'key', '3': 1, '4': 1, '5': 9, '10': 'key'},
    {'1': 'value', '3': 2, '4': 1, '5': 9, '10': 'value'},
  ],
  '7': {'7': true},
};

/// Descriptor for `RequestFunction`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List requestFunctionDescriptor = $convert.base64Decode(
    'Cg9SZXF1ZXN0RnVuY3Rpb24SGgoIZnVuY3Rpb24YASABKAlSCGZ1bmN0aW9uEjgKBGFyZ3MYAi'
    'ADKAsyJC5uZXJ2YXR1cmEuUmVxdWVzdEZ1bmN0aW9uLkFyZ3NFbnRyeVIEYXJncxo3CglBcmdz'
    'RW50cnkSEAoDa2V5GAEgASgJUgNrZXkSFAoFdmFsdWUYAiABKAlSBXZhbHVlOgI4AQ==');

@$core.Deprecated('Use requestViewDescriptor instead')
const RequestView$json = {
  '1': 'RequestView',
  '2': [
    {'1': 'name', '3': 1, '4': 1, '5': 14, '6': '.nervatura.ViewName', '10': 'name'},
    {'1': 'filter', '3': 2, '4': 1, '5': 9, '10': 'filter'},
    {'1': 'order_by', '3': 3, '4': 3, '5': 9, '10': 'orderBy'},
    {'1': 'limit', '3': 4, '4': 1, '5': 3, '10': 'limit'},
    {'1': 'offset', '3': 5, '4': 1, '5': 3, '10': 'offset'},
  ],
};

/// Descriptor for `RequestView`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List requestViewDescriptor = $convert.base64Decode(
    'CgtSZXF1ZXN0VmlldxInCgRuYW1lGAEgASgOMhMubmVydmF0dXJhLlZpZXdOYW1lUgRuYW1lEh'
    'YKBmZpbHRlchgCIAEoCVIGZmlsdGVyEhkKCG9yZGVyX2J5GAMgAygJUgdvcmRlckJ5EhQKBWxp'
    'bWl0GAQgASgDUgVsaW1pdBIWCgZvZmZzZXQYBSABKANSBm9mZnNldA==');

@$core.Deprecated('Use requestDatabaseDescriptor instead')
const RequestDatabase$json = {
  '1': 'RequestDatabase',
  '2': [
    {'1': 'alias', '3': 1, '4': 1, '5': 9, '10': 'alias'},
    {'1': 'demo', '3': 2, '4': 1, '5': 8, '10': 'demo'},
  ],
};

/// Descriptor for `RequestDatabase`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List requestDatabaseDescriptor = $convert.base64Decode(
    'Cg9SZXF1ZXN0RGF0YWJhc2USFAoFYWxpYXMYASABKAlSBWFsaWFzEhIKBGRlbW8YAiABKAhSBG'
    'RlbW8=');

@$core.Deprecated('Use jsonStringDescriptor instead')
const JsonString$json = {
  '1': 'JsonString',
  '2': [
    {'1': 'data', '3': 1, '4': 3, '5': 11, '6': '.nervatura.JsonString.DataEntry', '10': 'data'},
  ],
  '3': [JsonString_DataEntry$json],
};

@$core.Deprecated('Use jsonStringDescriptor instead')
const JsonString_DataEntry$json = {
  '1': 'DataEntry',
  '2': [
    {'1': 'key', '3': 1, '4': 1, '5': 9, '10': 'key'},
    {'1': 'value', '3': 2, '4': 1, '5': 9, '10': 'value'},
  ],
  '7': {'7': true},
};

/// Descriptor for `JsonString`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List jsonStringDescriptor = $convert.base64Decode(
    'CgpKc29uU3RyaW5nEjMKBGRhdGEYASADKAsyHy5uZXJ2YXR1cmEuSnNvblN0cmluZy5EYXRhRW'
    '50cnlSBGRhdGEaNwoJRGF0YUVudHJ5EhAKA2tleRgBIAEoCVIDa2V5EhQKBXZhbHVlGAIgASgJ'
    'UgV2YWx1ZToCOAE=');

@$core.Deprecated('Use jsonBytesDescriptor instead')
const JsonBytes$json = {
  '1': 'JsonBytes',
  '2': [
    {'1': 'data', '3': 1, '4': 1, '5': 12, '10': 'data'},
  ],
};

/// Descriptor for `JsonBytes`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List jsonBytesDescriptor = $convert.base64Decode(
    'CglKc29uQnl0ZXMSEgoEZGF0YRgBIAEoDFIEZGF0YQ==');

@$core.Deprecated('Use addressDescriptor instead')
const Address$json = {
  '1': 'Address',
  '2': [
    {'1': 'country', '3': 1, '4': 1, '5': 9, '10': 'country'},
    {'1': 'state', '3': 2, '4': 1, '5': 9, '10': 'state'},
    {'1': 'zip_code', '3': 3, '4': 1, '5': 9, '10': 'zipCode'},
    {'1': 'city', '3': 4, '4': 1, '5': 9, '10': 'city'},
    {'1': 'street', '3': 5, '4': 1, '5': 9, '10': 'street'},
    {'1': 'notes', '3': 6, '4': 1, '5': 9, '10': 'notes'},
    {'1': 'tags', '3': 7, '4': 3, '5': 9, '10': 'tags'},
    {'1': 'address_map', '3': 8, '4': 1, '5': 11, '6': '.nervatura.JsonString', '10': 'addressMap'},
  ],
};

/// Descriptor for `Address`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List addressDescriptor = $convert.base64Decode(
    'CgdBZGRyZXNzEhgKB2NvdW50cnkYASABKAlSB2NvdW50cnkSFAoFc3RhdGUYAiABKAlSBXN0YX'
    'RlEhkKCHppcF9jb2RlGAMgASgJUgd6aXBDb2RlEhIKBGNpdHkYBCABKAlSBGNpdHkSFgoGc3Ry'
    'ZWV0GAUgASgJUgZzdHJlZXQSFAoFbm90ZXMYBiABKAlSBW5vdGVzEhIKBHRhZ3MYByADKAlSBH'
    'RhZ3MSNgoLYWRkcmVzc19tYXAYCCABKAsyFS5uZXJ2YXR1cmEuSnNvblN0cmluZ1IKYWRkcmVz'
    'c01hcA==');

@$core.Deprecated('Use contactDescriptor instead')
const Contact$json = {
  '1': 'Contact',
  '2': [
    {'1': 'first_name', '3': 1, '4': 1, '5': 9, '10': 'firstName'},
    {'1': 'surname', '3': 2, '4': 1, '5': 9, '10': 'surname'},
    {'1': 'status', '3': 3, '4': 1, '5': 9, '10': 'status'},
    {'1': 'phone', '3': 4, '4': 1, '5': 9, '10': 'phone'},
    {'1': 'mobile', '3': 5, '4': 1, '5': 9, '10': 'mobile'},
    {'1': 'email', '3': 6, '4': 1, '5': 9, '10': 'email'},
    {'1': 'notes', '3': 7, '4': 1, '5': 9, '10': 'notes'},
    {'1': 'tags', '3': 8, '4': 3, '5': 9, '10': 'tags'},
    {'1': 'contact_map', '3': 9, '4': 1, '5': 11, '6': '.nervatura.JsonString', '10': 'contactMap'},
  ],
};

/// Descriptor for `Contact`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List contactDescriptor = $convert.base64Decode(
    'CgdDb250YWN0Eh0KCmZpcnN0X25hbWUYASABKAlSCWZpcnN0TmFtZRIYCgdzdXJuYW1lGAIgAS'
    'gJUgdzdXJuYW1lEhYKBnN0YXR1cxgDIAEoCVIGc3RhdHVzEhQKBXBob25lGAQgASgJUgVwaG9u'
    'ZRIWCgZtb2JpbGUYBSABKAlSBm1vYmlsZRIUCgVlbWFpbBgGIAEoCVIFZW1haWwSFAoFbm90ZX'
    'MYByABKAlSBW5vdGVzEhIKBHRhZ3MYCCADKAlSBHRhZ3MSNgoLY29udGFjdF9tYXAYCSABKAsy'
    'FS5uZXJ2YXR1cmEuSnNvblN0cmluZ1IKY29udGFjdE1hcA==');

@$core.Deprecated('Use eventDescriptor instead')
const Event$json = {
  '1': 'Event',
  '2': [
    {'1': 'uid', '3': 1, '4': 1, '5': 9, '10': 'uid'},
    {'1': 'subject', '3': 2, '4': 1, '5': 9, '10': 'subject'},
    {'1': 'start_time', '3': 3, '4': 1, '5': 9, '10': 'startTime'},
    {'1': 'end_time', '3': 4, '4': 1, '5': 9, '10': 'endTime'},
    {'1': 'place', '3': 5, '4': 1, '5': 9, '10': 'place'},
    {'1': 'description', '3': 6, '4': 1, '5': 9, '10': 'description'},
    {'1': 'tags', '3': 7, '4': 3, '5': 9, '10': 'tags'},
    {'1': 'event_map', '3': 8, '4': 1, '5': 11, '6': '.nervatura.JsonString', '10': 'eventMap'},
  ],
};

/// Descriptor for `Event`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List eventDescriptor = $convert.base64Decode(
    'CgVFdmVudBIQCgN1aWQYASABKAlSA3VpZBIYCgdzdWJqZWN0GAIgASgJUgdzdWJqZWN0Eh0KCn'
    'N0YXJ0X3RpbWUYAyABKAlSCXN0YXJ0VGltZRIZCghlbmRfdGltZRgEIAEoCVIHZW5kVGltZRIU'
    'CgVwbGFjZRgFIAEoCVIFcGxhY2USIAoLZGVzY3JpcHRpb24YBiABKAlSC2Rlc2NyaXB0aW9uEh'
    'IKBHRhZ3MYByADKAlSBHRhZ3MSMgoJZXZlbnRfbWFwGAggASgLMhUubmVydmF0dXJhLkpzb25T'
    'dHJpbmdSCGV2ZW50TWFw');

@$core.Deprecated('Use bookmarkDescriptor instead')
const Bookmark$json = {
  '1': 'Bookmark',
  '2': [
    {'1': 'bookmark_type', '3': 1, '4': 1, '5': 14, '6': '.nervatura.BookmarkType', '10': 'bookmarkType'},
    {'1': 'label', '3': 2, '4': 1, '5': 9, '10': 'label'},
    {'1': 'key', '3': 3, '4': 1, '5': 9, '10': 'key'},
    {'1': 'code', '3': 4, '4': 1, '5': 9, '10': 'code'},
    {'1': 'filters', '3': 5, '4': 1, '5': 11, '6': '.nervatura.JsonString', '10': 'filters'},
    {'1': 'columns', '3': 6, '4': 3, '5': 11, '6': '.nervatura.Bookmark.ColumnsEntry', '10': 'columns'},
    {'1': 'time_stamp', '3': 7, '4': 1, '5': 9, '10': 'timeStamp'},
  ],
  '3': [Bookmark_ColumnsEntry$json],
};

@$core.Deprecated('Use bookmarkDescriptor instead')
const Bookmark_ColumnsEntry$json = {
  '1': 'ColumnsEntry',
  '2': [
    {'1': 'key', '3': 1, '4': 1, '5': 9, '10': 'key'},
    {'1': 'value', '3': 2, '4': 1, '5': 8, '10': 'value'},
  ],
  '7': {'7': true},
};

/// Descriptor for `Bookmark`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List bookmarkDescriptor = $convert.base64Decode(
    'CghCb29rbWFyaxI8Cg1ib29rbWFya190eXBlGAEgASgOMhcubmVydmF0dXJhLkJvb2ttYXJrVH'
    'lwZVIMYm9va21hcmtUeXBlEhQKBWxhYmVsGAIgASgJUgVsYWJlbBIQCgNrZXkYAyABKAlSA2tl'
    'eRISCgRjb2RlGAQgASgJUgRjb2RlEi8KB2ZpbHRlcnMYBSABKAsyFS5uZXJ2YXR1cmEuSnNvbl'
    'N0cmluZ1IHZmlsdGVycxI6Cgdjb2x1bW5zGAYgAygLMiAubmVydmF0dXJhLkJvb2ttYXJrLkNv'
    'bHVtbnNFbnRyeVIHY29sdW1ucxIdCgp0aW1lX3N0YW1wGAcgASgJUgl0aW1lU3RhbXAaOgoMQ2'
    '9sdW1uc0VudHJ5EhAKA2tleRgBIAEoCVIDa2V5EhQKBXZhbHVlGAIgASgIUgV2YWx1ZToCOAE=');

@$core.Deprecated('Use authMetaDescriptor instead')
const AuthMeta$json = {
  '1': 'AuthMeta',
  '2': [
    {'1': 'tags', '3': 1, '4': 3, '5': 9, '10': 'tags'},
    {'1': 'bookmarks', '3': 2, '4': 3, '5': 11, '6': '.nervatura.Bookmark', '10': 'bookmarks'},
  ],
};

/// Descriptor for `AuthMeta`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List authMetaDescriptor = $convert.base64Decode(
    'CghBdXRoTWV0YRISCgR0YWdzGAEgAygJUgR0YWdzEjEKCWJvb2ttYXJrcxgCIAMoCzITLm5lcn'
    'ZhdHVyYS5Cb29rbWFya1IJYm9va21hcmtz');

@$core.Deprecated('Use authDescriptor instead')
const Auth$json = {
  '1': 'Auth',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    {'1': 'code', '3': 2, '4': 1, '5': 9, '10': 'code'},
    {'1': 'user_name', '3': 3, '4': 1, '5': 9, '10': 'userName'},
    {'1': 'user_group', '3': 4, '4': 1, '5': 14, '6': '.nervatura.UserGroup', '10': 'userGroup'},
    {'1': 'disabled', '3': 5, '4': 1, '5': 8, '10': 'disabled'},
    {'1': 'auth_meta', '3': 6, '4': 1, '5': 11, '6': '.nervatura.AuthMeta', '10': 'authMeta'},
    {'1': 'auth_map', '3': 7, '4': 1, '5': 11, '6': '.nervatura.JsonString', '10': 'authMap'},
    {'1': 'time_stamp', '3': 8, '4': 1, '5': 9, '10': 'timeStamp'},
  ],
};

/// Descriptor for `Auth`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List authDescriptor = $convert.base64Decode(
    'CgRBdXRoEg4KAmlkGAEgASgDUgJpZBISCgRjb2RlGAIgASgJUgRjb2RlEhsKCXVzZXJfbmFtZR'
    'gDIAEoCVIIdXNlck5hbWUSMwoKdXNlcl9ncm91cBgEIAEoDjIULm5lcnZhdHVyYS5Vc2VyR3Jv'
    'dXBSCXVzZXJHcm91cBIaCghkaXNhYmxlZBgFIAEoCFIIZGlzYWJsZWQSMAoJYXV0aF9tZXRhGA'
    'YgASgLMhMubmVydmF0dXJhLkF1dGhNZXRhUghhdXRoTWV0YRIwCghhdXRoX21hcBgHIAEoCzIV'
    'Lm5lcnZhdHVyYS5Kc29uU3RyaW5nUgdhdXRoTWFwEh0KCnRpbWVfc3RhbXAYCCABKAlSCXRpbW'
    'VTdGFtcA==');

@$core.Deprecated('Use currencyMetaDescriptor instead')
const CurrencyMeta$json = {
  '1': 'CurrencyMeta',
  '2': [
    {'1': 'description', '3': 1, '4': 1, '5': 9, '10': 'description'},
    {'1': 'digit', '3': 2, '4': 1, '5': 3, '10': 'digit'},
    {'1': 'cash_round', '3': 3, '4': 1, '5': 3, '10': 'cashRound'},
    {'1': 'tags', '3': 4, '4': 3, '5': 9, '10': 'tags'},
  ],
};

/// Descriptor for `CurrencyMeta`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List currencyMetaDescriptor = $convert.base64Decode(
    'CgxDdXJyZW5jeU1ldGESIAoLZGVzY3JpcHRpb24YASABKAlSC2Rlc2NyaXB0aW9uEhQKBWRpZ2'
    'l0GAIgASgDUgVkaWdpdBIdCgpjYXNoX3JvdW5kGAMgASgDUgljYXNoUm91bmQSEgoEdGFncxgE'
    'IAMoCVIEdGFncw==');

@$core.Deprecated('Use currencyDescriptor instead')
const Currency$json = {
  '1': 'Currency',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    {'1': 'code', '3': 2, '4': 1, '5': 9, '10': 'code'},
    {'1': 'currency_meta', '3': 3, '4': 1, '5': 11, '6': '.nervatura.CurrencyMeta', '10': 'currencyMeta'},
    {'1': 'currency_map', '3': 4, '4': 1, '5': 11, '6': '.nervatura.JsonString', '10': 'currencyMap'},
    {'1': 'time_stamp', '3': 5, '4': 1, '5': 9, '10': 'timeStamp'},
  ],
};

/// Descriptor for `Currency`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List currencyDescriptor = $convert.base64Decode(
    'CghDdXJyZW5jeRIOCgJpZBgBIAEoA1ICaWQSEgoEY29kZRgCIAEoCVIEY29kZRI8Cg1jdXJyZW'
    '5jeV9tZXRhGAMgASgLMhcubmVydmF0dXJhLkN1cnJlbmN5TWV0YVIMY3VycmVuY3lNZXRhEjgK'
    'DGN1cnJlbmN5X21hcBgEIAEoCzIVLm5lcnZhdHVyYS5Kc29uU3RyaW5nUgtjdXJyZW5jeU1hcB'
    'IdCgp0aW1lX3N0YW1wGAUgASgJUgl0aW1lU3RhbXA=');

@$core.Deprecated('Use currenciesDescriptor instead')
const Currencies$json = {
  '1': 'Currencies',
  '2': [
    {'1': 'data', '3': 1, '4': 3, '5': 11, '6': '.nervatura.Currency', '10': 'data'},
  ],
};

/// Descriptor for `Currencies`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List currenciesDescriptor = $convert.base64Decode(
    'CgpDdXJyZW5jaWVzEicKBGRhdGEYASADKAsyEy5uZXJ2YXR1cmEuQ3VycmVuY3lSBGRhdGE=');

@$core.Deprecated('Use customerMetaDescriptor instead')
const CustomerMeta$json = {
  '1': 'CustomerMeta',
  '2': [
    {'1': 'tax_number', '3': 1, '4': 1, '5': 9, '10': 'taxNumber'},
    {'1': 'account', '3': 2, '4': 1, '5': 9, '10': 'account'},
    {'1': 'tax_free', '3': 3, '4': 1, '5': 8, '10': 'taxFree'},
    {'1': 'terms', '3': 4, '4': 1, '5': 3, '10': 'terms'},
    {'1': 'credit_limit', '3': 5, '4': 1, '5': 1, '10': 'creditLimit'},
    {'1': 'discount', '3': 6, '4': 1, '5': 1, '10': 'discount'},
    {'1': 'notes', '3': 7, '4': 1, '5': 9, '10': 'notes'},
    {'1': 'inactive', '3': 8, '4': 1, '5': 8, '10': 'inactive'},
    {'1': 'tags', '3': 9, '4': 3, '5': 9, '10': 'tags'},
  ],
};

/// Descriptor for `CustomerMeta`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List customerMetaDescriptor = $convert.base64Decode(
    'CgxDdXN0b21lck1ldGESHQoKdGF4X251bWJlchgBIAEoCVIJdGF4TnVtYmVyEhgKB2FjY291bn'
    'QYAiABKAlSB2FjY291bnQSGQoIdGF4X2ZyZWUYAyABKAhSB3RheEZyZWUSFAoFdGVybXMYBCAB'
    'KANSBXRlcm1zEiEKDGNyZWRpdF9saW1pdBgFIAEoAVILY3JlZGl0TGltaXQSGgoIZGlzY291bn'
    'QYBiABKAFSCGRpc2NvdW50EhQKBW5vdGVzGAcgASgJUgVub3RlcxIaCghpbmFjdGl2ZRgIIAEo'
    'CFIIaW5hY3RpdmUSEgoEdGFncxgJIAMoCVIEdGFncw==');

@$core.Deprecated('Use customerDescriptor instead')
const Customer$json = {
  '1': 'Customer',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    {'1': 'code', '3': 2, '4': 1, '5': 9, '10': 'code'},
    {'1': 'customer_type', '3': 3, '4': 1, '5': 14, '6': '.nervatura.CustomerType', '10': 'customerType'},
    {'1': 'customer_name', '3': 4, '4': 1, '5': 9, '10': 'customerName'},
    {'1': 'addresses', '3': 5, '4': 3, '5': 11, '6': '.nervatura.Address', '10': 'addresses'},
    {'1': 'contacts', '3': 6, '4': 3, '5': 11, '6': '.nervatura.Contact', '10': 'contacts'},
    {'1': 'events', '3': 7, '4': 3, '5': 11, '6': '.nervatura.Event', '10': 'events'},
    {'1': 'customer_meta', '3': 8, '4': 1, '5': 11, '6': '.nervatura.CustomerMeta', '10': 'customerMeta'},
    {'1': 'customer_map', '3': 9, '4': 1, '5': 11, '6': '.nervatura.JsonString', '10': 'customerMap'},
    {'1': 'time_stamp', '3': 10, '4': 1, '5': 9, '10': 'timeStamp'},
  ],
};

/// Descriptor for `Customer`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List customerDescriptor = $convert.base64Decode(
    'CghDdXN0b21lchIOCgJpZBgBIAEoA1ICaWQSEgoEY29kZRgCIAEoCVIEY29kZRI8Cg1jdXN0b2'
    '1lcl90eXBlGAMgASgOMhcubmVydmF0dXJhLkN1c3RvbWVyVHlwZVIMY3VzdG9tZXJUeXBlEiMK'
    'DWN1c3RvbWVyX25hbWUYBCABKAlSDGN1c3RvbWVyTmFtZRIwCglhZGRyZXNzZXMYBSADKAsyEi'
    '5uZXJ2YXR1cmEuQWRkcmVzc1IJYWRkcmVzc2VzEi4KCGNvbnRhY3RzGAYgAygLMhIubmVydmF0'
    'dXJhLkNvbnRhY3RSCGNvbnRhY3RzEigKBmV2ZW50cxgHIAMoCzIQLm5lcnZhdHVyYS5FdmVudF'
    'IGZXZlbnRzEjwKDWN1c3RvbWVyX21ldGEYCCABKAsyFy5uZXJ2YXR1cmEuQ3VzdG9tZXJNZXRh'
    'UgxjdXN0b21lck1ldGESOAoMY3VzdG9tZXJfbWFwGAkgASgLMhUubmVydmF0dXJhLkpzb25TdH'
    'JpbmdSC2N1c3RvbWVyTWFwEh0KCnRpbWVfc3RhbXAYCiABKAlSCXRpbWVTdGFtcA==');

@$core.Deprecated('Use customersDescriptor instead')
const Customers$json = {
  '1': 'Customers',
  '2': [
    {'1': 'data', '3': 1, '4': 3, '5': 11, '6': '.nervatura.Customer', '10': 'data'},
  ],
};

/// Descriptor for `Customers`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List customersDescriptor = $convert.base64Decode(
    'CglDdXN0b21lcnMSJwoEZGF0YRgBIAMoCzITLm5lcnZhdHVyYS5DdXN0b21lclIEZGF0YQ==');

@$core.Deprecated('Use employeeMetaDescriptor instead')
const EmployeeMeta$json = {
  '1': 'EmployeeMeta',
  '2': [
    {'1': 'start_date', '3': 1, '4': 1, '5': 9, '10': 'startDate'},
    {'1': 'end_date', '3': 2, '4': 1, '5': 9, '10': 'endDate'},
    {'1': 'inactive', '3': 3, '4': 1, '5': 8, '10': 'inactive'},
    {'1': 'notes', '3': 4, '4': 1, '5': 9, '10': 'notes'},
    {'1': 'tags', '3': 5, '4': 3, '5': 9, '10': 'tags'},
  ],
};

/// Descriptor for `EmployeeMeta`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List employeeMetaDescriptor = $convert.base64Decode(
    'CgxFbXBsb3llZU1ldGESHQoKc3RhcnRfZGF0ZRgBIAEoCVIJc3RhcnREYXRlEhkKCGVuZF9kYX'
    'RlGAIgASgJUgdlbmREYXRlEhoKCGluYWN0aXZlGAMgASgIUghpbmFjdGl2ZRIUCgVub3RlcxgE'
    'IAEoCVIFbm90ZXMSEgoEdGFncxgFIAMoCVIEdGFncw==');

@$core.Deprecated('Use employeeDescriptor instead')
const Employee$json = {
  '1': 'Employee',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    {'1': 'code', '3': 2, '4': 1, '5': 9, '10': 'code'},
    {'1': 'address', '3': 3, '4': 1, '5': 11, '6': '.nervatura.Address', '10': 'address'},
    {'1': 'contact', '3': 4, '4': 1, '5': 11, '6': '.nervatura.Contact', '10': 'contact'},
    {'1': 'events', '3': 5, '4': 3, '5': 11, '6': '.nervatura.Event', '10': 'events'},
    {'1': 'employee_meta', '3': 6, '4': 1, '5': 11, '6': '.nervatura.EmployeeMeta', '10': 'employeeMeta'},
    {'1': 'employee_map', '3': 7, '4': 1, '5': 11, '6': '.nervatura.JsonString', '10': 'employeeMap'},
    {'1': 'time_stamp', '3': 8, '4': 1, '5': 9, '10': 'timeStamp'},
  ],
};

/// Descriptor for `Employee`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List employeeDescriptor = $convert.base64Decode(
    'CghFbXBsb3llZRIOCgJpZBgBIAEoA1ICaWQSEgoEY29kZRgCIAEoCVIEY29kZRIsCgdhZGRyZX'
    'NzGAMgASgLMhIubmVydmF0dXJhLkFkZHJlc3NSB2FkZHJlc3MSLAoHY29udGFjdBgEIAEoCzIS'
    'Lm5lcnZhdHVyYS5Db250YWN0Ugdjb250YWN0EigKBmV2ZW50cxgFIAMoCzIQLm5lcnZhdHVyYS'
    '5FdmVudFIGZXZlbnRzEjwKDWVtcGxveWVlX21ldGEYBiABKAsyFy5uZXJ2YXR1cmEuRW1wbG95'
    'ZWVNZXRhUgxlbXBsb3llZU1ldGESOAoMZW1wbG95ZWVfbWFwGAcgASgLMhUubmVydmF0dXJhLk'
    'pzb25TdHJpbmdSC2VtcGxveWVlTWFwEh0KCnRpbWVfc3RhbXAYCCABKAlSCXRpbWVTdGFtcA==');

@$core.Deprecated('Use employeesDescriptor instead')
const Employees$json = {
  '1': 'Employees',
  '2': [
    {'1': 'data', '3': 1, '4': 3, '5': 11, '6': '.nervatura.Employee', '10': 'data'},
  ],
};

/// Descriptor for `Employees`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List employeesDescriptor = $convert.base64Decode(
    'CglFbXBsb3llZXMSJwoEZGF0YRgBIAMoCzITLm5lcnZhdHVyYS5FbXBsb3llZVIEZGF0YQ==');

@$core.Deprecated('Use itemMetaDescriptor instead')
const ItemMeta$json = {
  '1': 'ItemMeta',
  '2': [
    {'1': 'unit', '3': 1, '4': 1, '5': 9, '10': 'unit'},
    {'1': 'qty', '3': 2, '4': 1, '5': 1, '10': 'qty'},
    {'1': 'fx_price', '3': 3, '4': 1, '5': 1, '10': 'fxPrice'},
    {'1': 'net_amount', '3': 4, '4': 1, '5': 1, '10': 'netAmount'},
    {'1': 'discount', '3': 5, '4': 1, '5': 1, '10': 'discount'},
    {'1': 'vat_amount', '3': 6, '4': 1, '5': 1, '10': 'vatAmount'},
    {'1': 'amount', '3': 7, '4': 1, '5': 1, '10': 'amount'},
    {'1': 'description', '3': 8, '4': 1, '5': 9, '10': 'description'},
    {'1': 'deposit', '3': 9, '4': 1, '5': 8, '10': 'deposit'},
    {'1': 'own_stock', '3': 10, '4': 1, '5': 1, '10': 'ownStock'},
    {'1': 'action_price', '3': 11, '4': 1, '5': 8, '10': 'actionPrice'},
    {'1': 'tags', '3': 12, '4': 3, '5': 9, '10': 'tags'},
  ],
};

/// Descriptor for `ItemMeta`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List itemMetaDescriptor = $convert.base64Decode(
    'CghJdGVtTWV0YRISCgR1bml0GAEgASgJUgR1bml0EhAKA3F0eRgCIAEoAVIDcXR5EhkKCGZ4X3'
    'ByaWNlGAMgASgBUgdmeFByaWNlEh0KCm5ldF9hbW91bnQYBCABKAFSCW5ldEFtb3VudBIaCghk'
    'aXNjb3VudBgFIAEoAVIIZGlzY291bnQSHQoKdmF0X2Ftb3VudBgGIAEoAVIJdmF0QW1vdW50Eh'
    'YKBmFtb3VudBgHIAEoAVIGYW1vdW50EiAKC2Rlc2NyaXB0aW9uGAggASgJUgtkZXNjcmlwdGlv'
    'bhIYCgdkZXBvc2l0GAkgASgIUgdkZXBvc2l0EhsKCW93bl9zdG9jaxgKIAEoAVIIb3duU3RvY2'
    'sSIQoMYWN0aW9uX3ByaWNlGAsgASgIUgthY3Rpb25QcmljZRISCgR0YWdzGAwgAygJUgR0YWdz');

@$core.Deprecated('Use itemDescriptor instead')
const Item$json = {
  '1': 'Item',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    {'1': 'code', '3': 2, '4': 1, '5': 9, '10': 'code'},
    {'1': 'trans_code', '3': 3, '4': 1, '5': 9, '10': 'transCode'},
    {'1': 'product_code', '3': 4, '4': 1, '5': 9, '10': 'productCode'},
    {'1': 'tax_code', '3': 5, '4': 1, '5': 9, '10': 'taxCode'},
    {'1': 'item_meta', '3': 6, '4': 1, '5': 11, '6': '.nervatura.ItemMeta', '10': 'itemMeta'},
    {'1': 'item_map', '3': 7, '4': 1, '5': 11, '6': '.nervatura.JsonString', '10': 'itemMap'},
    {'1': 'time_stamp', '3': 8, '4': 1, '5': 9, '10': 'timeStamp'},
  ],
};

/// Descriptor for `Item`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List itemDescriptor = $convert.base64Decode(
    'CgRJdGVtEg4KAmlkGAEgASgDUgJpZBISCgRjb2RlGAIgASgJUgRjb2RlEh0KCnRyYW5zX2NvZG'
    'UYAyABKAlSCXRyYW5zQ29kZRIhCgxwcm9kdWN0X2NvZGUYBCABKAlSC3Byb2R1Y3RDb2RlEhkK'
    'CHRheF9jb2RlGAUgASgJUgd0YXhDb2RlEjAKCWl0ZW1fbWV0YRgGIAEoCzITLm5lcnZhdHVyYS'
    '5JdGVtTWV0YVIIaXRlbU1ldGESMAoIaXRlbV9tYXAYByABKAsyFS5uZXJ2YXR1cmEuSnNvblN0'
    'cmluZ1IHaXRlbU1hcBIdCgp0aW1lX3N0YW1wGAggASgJUgl0aW1lU3RhbXA=');

@$core.Deprecated('Use itemsDescriptor instead')
const Items$json = {
  '1': 'Items',
  '2': [
    {'1': 'data', '3': 1, '4': 3, '5': 11, '6': '.nervatura.Item', '10': 'data'},
  ],
};

/// Descriptor for `Items`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List itemsDescriptor = $convert.base64Decode(
    'CgVJdGVtcxIjCgRkYXRhGAEgAygLMg8ubmVydmF0dXJhLkl0ZW1SBGRhdGE=');

@$core.Deprecated('Use linkMetaDescriptor instead')
const LinkMeta$json = {
  '1': 'LinkMeta',
  '2': [
    {'1': 'qty', '3': 1, '4': 1, '5': 1, '10': 'qty'},
    {'1': 'rate', '3': 2, '4': 1, '5': 1, '10': 'rate'},
    {'1': 'tags', '3': 3, '4': 3, '5': 9, '10': 'tags'},
  ],
};

/// Descriptor for `LinkMeta`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List linkMetaDescriptor = $convert.base64Decode(
    'CghMaW5rTWV0YRIQCgNxdHkYASABKAFSA3F0eRISCgRyYXRlGAIgASgBUgRyYXRlEhIKBHRhZ3'
    'MYAyADKAlSBHRhZ3M=');

@$core.Deprecated('Use linkDescriptor instead')
const Link$json = {
  '1': 'Link',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    {'1': 'code', '3': 2, '4': 1, '5': 9, '10': 'code'},
    {'1': 'link_type_1', '3': 3, '4': 1, '5': 14, '6': '.nervatura.LinkType', '10': 'linkType1'},
    {'1': 'link_code_1', '3': 4, '4': 1, '5': 9, '10': 'linkCode1'},
    {'1': 'link_type_2', '3': 5, '4': 1, '5': 14, '6': '.nervatura.LinkType', '10': 'linkType2'},
    {'1': 'link_code_2', '3': 6, '4': 1, '5': 9, '10': 'linkCode2'},
    {'1': 'link_meta', '3': 7, '4': 1, '5': 11, '6': '.nervatura.LinkMeta', '10': 'linkMeta'},
    {'1': 'link_map', '3': 8, '4': 1, '5': 11, '6': '.nervatura.JsonString', '10': 'linkMap'},
    {'1': 'time_stamp', '3': 9, '4': 1, '5': 9, '10': 'timeStamp'},
  ],
};

/// Descriptor for `Link`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List linkDescriptor = $convert.base64Decode(
    'CgRMaW5rEg4KAmlkGAEgASgDUgJpZBISCgRjb2RlGAIgASgJUgRjb2RlEjMKC2xpbmtfdHlwZV'
    '8xGAMgASgOMhMubmVydmF0dXJhLkxpbmtUeXBlUglsaW5rVHlwZTESHgoLbGlua19jb2RlXzEY'
    'BCABKAlSCWxpbmtDb2RlMRIzCgtsaW5rX3R5cGVfMhgFIAEoDjITLm5lcnZhdHVyYS5MaW5rVH'
    'lwZVIJbGlua1R5cGUyEh4KC2xpbmtfY29kZV8yGAYgASgJUglsaW5rQ29kZTISMAoJbGlua19t'
    'ZXRhGAcgASgLMhMubmVydmF0dXJhLkxpbmtNZXRhUghsaW5rTWV0YRIwCghsaW5rX21hcBgIIA'
    'EoCzIVLm5lcnZhdHVyYS5Kc29uU3RyaW5nUgdsaW5rTWFwEh0KCnRpbWVfc3RhbXAYCSABKAlS'
    'CXRpbWVTdGFtcA==');

@$core.Deprecated('Use linksDescriptor instead')
const Links$json = {
  '1': 'Links',
  '2': [
    {'1': 'data', '3': 1, '4': 3, '5': 11, '6': '.nervatura.Link', '10': 'data'},
  ],
};

/// Descriptor for `Links`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List linksDescriptor = $convert.base64Decode(
    'CgVMaW5rcxIjCgRkYXRhGAEgAygLMg8ubmVydmF0dXJhLkxpbmtSBGRhdGE=');

@$core.Deprecated('Use logDescriptor instead')
const Log$json = {
  '1': 'Log',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    {'1': 'code', '3': 2, '4': 1, '5': 9, '10': 'code'},
    {'1': 'log_type', '3': 3, '4': 1, '5': 14, '6': '.nervatura.LogType', '10': 'logType'},
    {'1': 'ref_type', '3': 4, '4': 1, '5': 9, '10': 'refType'},
    {'1': 'ref_code', '3': 5, '4': 1, '5': 9, '10': 'refCode'},
    {'1': 'auth_code', '3': 6, '4': 1, '5': 9, '10': 'authCode'},
    {'1': 'log_map', '3': 7, '4': 1, '5': 11, '6': '.nervatura.JsonString', '10': 'logMap'},
    {'1': 'time_stamp', '3': 8, '4': 1, '5': 9, '10': 'timeStamp'},
  ],
};

/// Descriptor for `Log`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List logDescriptor = $convert.base64Decode(
    'CgNMb2cSDgoCaWQYASABKANSAmlkEhIKBGNvZGUYAiABKAlSBGNvZGUSLQoIbG9nX3R5cGUYAy'
    'ABKA4yEi5uZXJ2YXR1cmEuTG9nVHlwZVIHbG9nVHlwZRIZCghyZWZfdHlwZRgEIAEoCVIHcmVm'
    'VHlwZRIZCghyZWZfY29kZRgFIAEoCVIHcmVmQ29kZRIbCglhdXRoX2NvZGUYBiABKAlSCGF1dG'
    'hDb2RlEi4KB2xvZ19tYXAYByABKAsyFS5uZXJ2YXR1cmEuSnNvblN0cmluZ1IGbG9nTWFwEh0K'
    'CnRpbWVfc3RhbXAYCCABKAlSCXRpbWVTdGFtcA==');

@$core.Deprecated('Use logsDescriptor instead')
const Logs$json = {
  '1': 'Logs',
  '2': [
    {'1': 'data', '3': 1, '4': 3, '5': 11, '6': '.nervatura.Log', '10': 'data'},
  ],
};

/// Descriptor for `Logs`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List logsDescriptor = $convert.base64Decode(
    'CgRMb2dzEiIKBGRhdGEYASADKAsyDi5uZXJ2YXR1cmEuTG9nUgRkYXRh');

@$core.Deprecated('Use movementMetaDescriptor instead')
const MovementMeta$json = {
  '1': 'MovementMeta',
  '2': [
    {'1': 'qty', '3': 1, '4': 1, '5': 1, '10': 'qty'},
    {'1': 'notes', '3': 2, '4': 1, '5': 9, '10': 'notes'},
    {'1': 'shared', '3': 3, '4': 1, '5': 8, '10': 'shared'},
    {'1': 'tags', '3': 4, '4': 3, '5': 9, '10': 'tags'},
  ],
};

/// Descriptor for `MovementMeta`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List movementMetaDescriptor = $convert.base64Decode(
    'CgxNb3ZlbWVudE1ldGESEAoDcXR5GAEgASgBUgNxdHkSFAoFbm90ZXMYAiABKAlSBW5vdGVzEh'
    'YKBnNoYXJlZBgDIAEoCFIGc2hhcmVkEhIKBHRhZ3MYBCADKAlSBHRhZ3M=');

@$core.Deprecated('Use movementDescriptor instead')
const Movement$json = {
  '1': 'Movement',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    {'1': 'code', '3': 2, '4': 1, '5': 9, '10': 'code'},
    {'1': 'movement_type', '3': 3, '4': 1, '5': 14, '6': '.nervatura.MovementType', '10': 'movementType'},
    {'1': 'shipping_time', '3': 4, '4': 1, '5': 9, '10': 'shippingTime'},
    {'1': 'trans_code', '3': 5, '4': 1, '5': 9, '10': 'transCode'},
    {'1': 'product_code', '3': 6, '4': 1, '5': 9, '10': 'productCode'},
    {'1': 'tool_code', '3': 7, '4': 1, '5': 9, '10': 'toolCode'},
    {'1': 'place_code', '3': 8, '4': 1, '5': 9, '10': 'placeCode'},
    {'1': 'item_code', '3': 9, '4': 1, '5': 9, '10': 'itemCode'},
    {'1': 'movement_code', '3': 10, '4': 1, '5': 9, '10': 'movementCode'},
    {'1': 'movement_meta', '3': 11, '4': 1, '5': 11, '6': '.nervatura.MovementMeta', '10': 'movementMeta'},
    {'1': 'movement_map', '3': 12, '4': 1, '5': 11, '6': '.nervatura.JsonString', '10': 'movementMap'},
    {'1': 'time_stamp', '3': 13, '4': 1, '5': 9, '10': 'timeStamp'},
  ],
};

/// Descriptor for `Movement`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List movementDescriptor = $convert.base64Decode(
    'CghNb3ZlbWVudBIOCgJpZBgBIAEoA1ICaWQSEgoEY29kZRgCIAEoCVIEY29kZRI8Cg1tb3ZlbW'
    'VudF90eXBlGAMgASgOMhcubmVydmF0dXJhLk1vdmVtZW50VHlwZVIMbW92ZW1lbnRUeXBlEiMK'
    'DXNoaXBwaW5nX3RpbWUYBCABKAlSDHNoaXBwaW5nVGltZRIdCgp0cmFuc19jb2RlGAUgASgJUg'
    'l0cmFuc0NvZGUSIQoMcHJvZHVjdF9jb2RlGAYgASgJUgtwcm9kdWN0Q29kZRIbCgl0b29sX2Nv'
    'ZGUYByABKAlSCHRvb2xDb2RlEh0KCnBsYWNlX2NvZGUYCCABKAlSCXBsYWNlQ29kZRIbCglpdG'
    'VtX2NvZGUYCSABKAlSCGl0ZW1Db2RlEiMKDW1vdmVtZW50X2NvZGUYCiABKAlSDG1vdmVtZW50'
    'Q29kZRI8Cg1tb3ZlbWVudF9tZXRhGAsgASgLMhcubmVydmF0dXJhLk1vdmVtZW50TWV0YVIMbW'
    '92ZW1lbnRNZXRhEjgKDG1vdmVtZW50X21hcBgMIAEoCzIVLm5lcnZhdHVyYS5Kc29uU3RyaW5n'
    'Ugttb3ZlbWVudE1hcBIdCgp0aW1lX3N0YW1wGA0gASgJUgl0aW1lU3RhbXA=');

@$core.Deprecated('Use movementsDescriptor instead')
const Movements$json = {
  '1': 'Movements',
  '2': [
    {'1': 'data', '3': 1, '4': 3, '5': 11, '6': '.nervatura.Movement', '10': 'data'},
  ],
};

/// Descriptor for `Movements`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List movementsDescriptor = $convert.base64Decode(
    'CglNb3ZlbWVudHMSJwoEZGF0YRgBIAMoCzITLm5lcnZhdHVyYS5Nb3ZlbWVudFIEZGF0YQ==');

@$core.Deprecated('Use paymentMetaDescriptor instead')
const PaymentMeta$json = {
  '1': 'PaymentMeta',
  '2': [
    {'1': 'amount', '3': 1, '4': 1, '5': 1, '10': 'amount'},
    {'1': 'notes', '3': 2, '4': 1, '5': 9, '10': 'notes'},
    {'1': 'tags', '3': 3, '4': 3, '5': 9, '10': 'tags'},
  ],
};

/// Descriptor for `PaymentMeta`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List paymentMetaDescriptor = $convert.base64Decode(
    'CgtQYXltZW50TWV0YRIWCgZhbW91bnQYASABKAFSBmFtb3VudBIUCgVub3RlcxgCIAEoCVIFbm'
    '90ZXMSEgoEdGFncxgDIAMoCVIEdGFncw==');

@$core.Deprecated('Use paymentDescriptor instead')
const Payment$json = {
  '1': 'Payment',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    {'1': 'code', '3': 2, '4': 1, '5': 9, '10': 'code'},
    {'1': 'paid_date', '3': 3, '4': 1, '5': 9, '10': 'paidDate'},
    {'1': 'trans_code', '3': 4, '4': 1, '5': 9, '10': 'transCode'},
    {'1': 'payment_meta', '3': 5, '4': 1, '5': 11, '6': '.nervatura.PaymentMeta', '10': 'paymentMeta'},
    {'1': 'payment_map', '3': 6, '4': 1, '5': 11, '6': '.nervatura.JsonString', '10': 'paymentMap'},
    {'1': 'time_stamp', '3': 7, '4': 1, '5': 9, '10': 'timeStamp'},
  ],
};

/// Descriptor for `Payment`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List paymentDescriptor = $convert.base64Decode(
    'CgdQYXltZW50Eg4KAmlkGAEgASgDUgJpZBISCgRjb2RlGAIgASgJUgRjb2RlEhsKCXBhaWRfZG'
    'F0ZRgDIAEoCVIIcGFpZERhdGUSHQoKdHJhbnNfY29kZRgEIAEoCVIJdHJhbnNDb2RlEjkKDHBh'
    'eW1lbnRfbWV0YRgFIAEoCzIWLm5lcnZhdHVyYS5QYXltZW50TWV0YVILcGF5bWVudE1ldGESNg'
    'oLcGF5bWVudF9tYXAYBiABKAsyFS5uZXJ2YXR1cmEuSnNvblN0cmluZ1IKcGF5bWVudE1hcBId'
    'Cgp0aW1lX3N0YW1wGAcgASgJUgl0aW1lU3RhbXA=');

@$core.Deprecated('Use paymentsDescriptor instead')
const Payments$json = {
  '1': 'Payments',
  '2': [
    {'1': 'data', '3': 1, '4': 3, '5': 11, '6': '.nervatura.Payment', '10': 'data'},
  ],
};

/// Descriptor for `Payments`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List paymentsDescriptor = $convert.base64Decode(
    'CghQYXltZW50cxImCgRkYXRhGAEgAygLMhIubmVydmF0dXJhLlBheW1lbnRSBGRhdGE=');

@$core.Deprecated('Use placeMetaDescriptor instead')
const PlaceMeta$json = {
  '1': 'PlaceMeta',
  '2': [
    {'1': 'notes', '3': 1, '4': 1, '5': 9, '10': 'notes'},
    {'1': 'inactive', '3': 2, '4': 1, '5': 8, '10': 'inactive'},
    {'1': 'tags', '3': 3, '4': 3, '5': 9, '10': 'tags'},
  ],
};

/// Descriptor for `PlaceMeta`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List placeMetaDescriptor = $convert.base64Decode(
    'CglQbGFjZU1ldGESFAoFbm90ZXMYASABKAlSBW5vdGVzEhoKCGluYWN0aXZlGAIgASgIUghpbm'
    'FjdGl2ZRISCgR0YWdzGAMgAygJUgR0YWdz');

@$core.Deprecated('Use placeDescriptor instead')
const Place$json = {
  '1': 'Place',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    {'1': 'code', '3': 2, '4': 1, '5': 9, '10': 'code'},
    {'1': 'place_type', '3': 3, '4': 1, '5': 14, '6': '.nervatura.PlaceType', '10': 'placeType'},
    {'1': 'place_name', '3': 4, '4': 1, '5': 9, '10': 'placeName'},
    {'1': 'currency_code', '3': 5, '4': 1, '5': 9, '10': 'currencyCode'},
    {'1': 'address', '3': 6, '4': 1, '5': 11, '6': '.nervatura.Address', '10': 'address'},
    {'1': 'contacts', '3': 7, '4': 3, '5': 11, '6': '.nervatura.Contact', '10': 'contacts'},
    {'1': 'events', '3': 8, '4': 3, '5': 11, '6': '.nervatura.Event', '10': 'events'},
    {'1': 'place_meta', '3': 9, '4': 1, '5': 11, '6': '.nervatura.PlaceMeta', '10': 'placeMeta'},
    {'1': 'place_map', '3': 10, '4': 1, '5': 11, '6': '.nervatura.JsonString', '10': 'placeMap'},
    {'1': 'time_stamp', '3': 11, '4': 1, '5': 9, '10': 'timeStamp'},
  ],
};

/// Descriptor for `Place`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List placeDescriptor = $convert.base64Decode(
    'CgVQbGFjZRIOCgJpZBgBIAEoA1ICaWQSEgoEY29kZRgCIAEoCVIEY29kZRIzCgpwbGFjZV90eX'
    'BlGAMgASgOMhQubmVydmF0dXJhLlBsYWNlVHlwZVIJcGxhY2VUeXBlEh0KCnBsYWNlX25hbWUY'
    'BCABKAlSCXBsYWNlTmFtZRIjCg1jdXJyZW5jeV9jb2RlGAUgASgJUgxjdXJyZW5jeUNvZGUSLA'
    'oHYWRkcmVzcxgGIAEoCzISLm5lcnZhdHVyYS5BZGRyZXNzUgdhZGRyZXNzEi4KCGNvbnRhY3Rz'
    'GAcgAygLMhIubmVydmF0dXJhLkNvbnRhY3RSCGNvbnRhY3RzEigKBmV2ZW50cxgIIAMoCzIQLm'
    '5lcnZhdHVyYS5FdmVudFIGZXZlbnRzEjMKCnBsYWNlX21ldGEYCSABKAsyFC5uZXJ2YXR1cmEu'
    'UGxhY2VNZXRhUglwbGFjZU1ldGESMgoJcGxhY2VfbWFwGAogASgLMhUubmVydmF0dXJhLkpzb2'
    '5TdHJpbmdSCHBsYWNlTWFwEh0KCnRpbWVfc3RhbXAYCyABKAlSCXRpbWVTdGFtcA==');

@$core.Deprecated('Use placesDescriptor instead')
const Places$json = {
  '1': 'Places',
  '2': [
    {'1': 'data', '3': 1, '4': 3, '5': 11, '6': '.nervatura.Place', '10': 'data'},
  ],
};

/// Descriptor for `Places`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List placesDescriptor = $convert.base64Decode(
    'CgZQbGFjZXMSJAoEZGF0YRgBIAMoCzIQLm5lcnZhdHVyYS5QbGFjZVIEZGF0YQ==');

@$core.Deprecated('Use priceMetaDescriptor instead')
const PriceMeta$json = {
  '1': 'PriceMeta',
  '2': [
    {'1': 'price_value', '3': 1, '4': 1, '5': 1, '10': 'priceValue'},
    {'1': 'tags', '3': 2, '4': 3, '5': 9, '10': 'tags'},
  ],
};

/// Descriptor for `PriceMeta`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List priceMetaDescriptor = $convert.base64Decode(
    'CglQcmljZU1ldGESHwoLcHJpY2VfdmFsdWUYASABKAFSCnByaWNlVmFsdWUSEgoEdGFncxgCIA'
    'MoCVIEdGFncw==');

@$core.Deprecated('Use priceDescriptor instead')
const Price$json = {
  '1': 'Price',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    {'1': 'code', '3': 2, '4': 1, '5': 9, '10': 'code'},
    {'1': 'price_type', '3': 3, '4': 1, '5': 14, '6': '.nervatura.PriceType', '10': 'priceType'},
    {'1': 'valid_from', '3': 4, '4': 1, '5': 9, '10': 'validFrom'},
    {'1': 'valid_to', '3': 5, '4': 1, '5': 9, '10': 'validTo'},
    {'1': 'product_code', '3': 6, '4': 1, '5': 9, '10': 'productCode'},
    {'1': 'currency_code', '3': 7, '4': 1, '5': 9, '10': 'currencyCode'},
    {'1': 'customer_code', '3': 8, '4': 1, '5': 9, '10': 'customerCode'},
    {'1': 'qty', '3': 9, '4': 1, '5': 1, '10': 'qty'},
    {'1': 'price_meta', '3': 10, '4': 1, '5': 11, '6': '.nervatura.PriceMeta', '10': 'priceMeta'},
    {'1': 'price_map', '3': 11, '4': 1, '5': 11, '6': '.nervatura.JsonString', '10': 'priceMap'},
    {'1': 'time_stamp', '3': 12, '4': 1, '5': 9, '10': 'timeStamp'},
  ],
};

/// Descriptor for `Price`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List priceDescriptor = $convert.base64Decode(
    'CgVQcmljZRIOCgJpZBgBIAEoA1ICaWQSEgoEY29kZRgCIAEoCVIEY29kZRIzCgpwcmljZV90eX'
    'BlGAMgASgOMhQubmVydmF0dXJhLlByaWNlVHlwZVIJcHJpY2VUeXBlEh0KCnZhbGlkX2Zyb20Y'
    'BCABKAlSCXZhbGlkRnJvbRIZCgh2YWxpZF90bxgFIAEoCVIHdmFsaWRUbxIhCgxwcm9kdWN0X2'
    'NvZGUYBiABKAlSC3Byb2R1Y3RDb2RlEiMKDWN1cnJlbmN5X2NvZGUYByABKAlSDGN1cnJlbmN5'
    'Q29kZRIjCg1jdXN0b21lcl9jb2RlGAggASgJUgxjdXN0b21lckNvZGUSEAoDcXR5GAkgASgBUg'
    'NxdHkSMwoKcHJpY2VfbWV0YRgKIAEoCzIULm5lcnZhdHVyYS5QcmljZU1ldGFSCXByaWNlTWV0'
    'YRIyCglwcmljZV9tYXAYCyABKAsyFS5uZXJ2YXR1cmEuSnNvblN0cmluZ1IIcHJpY2VNYXASHQ'
    'oKdGltZV9zdGFtcBgMIAEoCVIJdGltZVN0YW1w');

@$core.Deprecated('Use pricesDescriptor instead')
const Prices$json = {
  '1': 'Prices',
  '2': [
    {'1': 'data', '3': 1, '4': 3, '5': 11, '6': '.nervatura.Price', '10': 'data'},
  ],
};

/// Descriptor for `Prices`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pricesDescriptor = $convert.base64Decode(
    'CgZQcmljZXMSJAoEZGF0YRgBIAMoCzIQLm5lcnZhdHVyYS5QcmljZVIEZGF0YQ==');

@$core.Deprecated('Use productMetaDescriptor instead')
const ProductMeta$json = {
  '1': 'ProductMeta',
  '2': [
    {'1': 'unit', '3': 1, '4': 1, '5': 9, '10': 'unit'},
    {'1': 'barcode_type', '3': 2, '4': 1, '5': 14, '6': '.nervatura.BarcodeType', '10': 'barcodeType'},
    {'1': 'barcode_data', '3': 3, '4': 1, '5': 9, '10': 'barcodeData'},
    {'1': 'barcode_qty', '3': 4, '4': 1, '5': 1, '10': 'barcodeQty'},
    {'1': 'notes', '3': 5, '4': 1, '5': 9, '10': 'notes'},
    {'1': 'inactive', '3': 6, '4': 1, '5': 8, '10': 'inactive'},
    {'1': 'tags', '3': 7, '4': 3, '5': 9, '10': 'tags'},
  ],
};

/// Descriptor for `ProductMeta`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List productMetaDescriptor = $convert.base64Decode(
    'CgtQcm9kdWN0TWV0YRISCgR1bml0GAEgASgJUgR1bml0EjkKDGJhcmNvZGVfdHlwZRgCIAEoDj'
    'IWLm5lcnZhdHVyYS5CYXJjb2RlVHlwZVILYmFyY29kZVR5cGUSIQoMYmFyY29kZV9kYXRhGAMg'
    'ASgJUgtiYXJjb2RlRGF0YRIfCgtiYXJjb2RlX3F0eRgEIAEoAVIKYmFyY29kZVF0eRIUCgVub3'
    'RlcxgFIAEoCVIFbm90ZXMSGgoIaW5hY3RpdmUYBiABKAhSCGluYWN0aXZlEhIKBHRhZ3MYByAD'
    'KAlSBHRhZ3M=');

@$core.Deprecated('Use productDescriptor instead')
const Product$json = {
  '1': 'Product',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    {'1': 'code', '3': 2, '4': 1, '5': 9, '10': 'code'},
    {'1': 'product_type', '3': 3, '4': 1, '5': 14, '6': '.nervatura.ProductType', '10': 'productType'},
    {'1': 'product_name', '3': 4, '4': 1, '5': 9, '10': 'productName'},
    {'1': 'tax_code', '3': 5, '4': 1, '5': 9, '10': 'taxCode'},
    {'1': 'events', '3': 6, '4': 3, '5': 11, '6': '.nervatura.Event', '10': 'events'},
    {'1': 'product_meta', '3': 7, '4': 1, '5': 11, '6': '.nervatura.ProductMeta', '10': 'productMeta'},
    {'1': 'product_map', '3': 8, '4': 1, '5': 11, '6': '.nervatura.JsonString', '10': 'productMap'},
    {'1': 'time_stamp', '3': 9, '4': 1, '5': 9, '10': 'timeStamp'},
  ],
};

/// Descriptor for `Product`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List productDescriptor = $convert.base64Decode(
    'CgdQcm9kdWN0Eg4KAmlkGAEgASgDUgJpZBISCgRjb2RlGAIgASgJUgRjb2RlEjkKDHByb2R1Y3'
    'RfdHlwZRgDIAEoDjIWLm5lcnZhdHVyYS5Qcm9kdWN0VHlwZVILcHJvZHVjdFR5cGUSIQoMcHJv'
    'ZHVjdF9uYW1lGAQgASgJUgtwcm9kdWN0TmFtZRIZCgh0YXhfY29kZRgFIAEoCVIHdGF4Q29kZR'
    'IoCgZldmVudHMYBiADKAsyEC5uZXJ2YXR1cmEuRXZlbnRSBmV2ZW50cxI5Cgxwcm9kdWN0X21l'
    'dGEYByABKAsyFi5uZXJ2YXR1cmEuUHJvZHVjdE1ldGFSC3Byb2R1Y3RNZXRhEjYKC3Byb2R1Y3'
    'RfbWFwGAggASgLMhUubmVydmF0dXJhLkpzb25TdHJpbmdSCnByb2R1Y3RNYXASHQoKdGltZV9z'
    'dGFtcBgJIAEoCVIJdGltZVN0YW1w');

@$core.Deprecated('Use productsDescriptor instead')
const Products$json = {
  '1': 'Products',
  '2': [
    {'1': 'data', '3': 1, '4': 3, '5': 11, '6': '.nervatura.Product', '10': 'data'},
  ],
};

/// Descriptor for `Products`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List productsDescriptor = $convert.base64Decode(
    'CghQcm9kdWN0cxImCgRkYXRhGAEgAygLMhIubmVydmF0dXJhLlByb2R1Y3RSBGRhdGE=');

@$core.Deprecated('Use projectMetaDescriptor instead')
const ProjectMeta$json = {
  '1': 'ProjectMeta',
  '2': [
    {'1': 'start_date', '3': 1, '4': 1, '5': 9, '10': 'startDate'},
    {'1': 'end_date', '3': 2, '4': 1, '5': 9, '10': 'endDate'},
    {'1': 'notes', '3': 3, '4': 1, '5': 9, '10': 'notes'},
    {'1': 'inactive', '3': 4, '4': 1, '5': 8, '10': 'inactive'},
    {'1': 'tags', '3': 5, '4': 3, '5': 9, '10': 'tags'},
  ],
};

/// Descriptor for `ProjectMeta`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List projectMetaDescriptor = $convert.base64Decode(
    'CgtQcm9qZWN0TWV0YRIdCgpzdGFydF9kYXRlGAEgASgJUglzdGFydERhdGUSGQoIZW5kX2RhdG'
    'UYAiABKAlSB2VuZERhdGUSFAoFbm90ZXMYAyABKAlSBW5vdGVzEhoKCGluYWN0aXZlGAQgASgI'
    'UghpbmFjdGl2ZRISCgR0YWdzGAUgAygJUgR0YWdz');

@$core.Deprecated('Use projectDescriptor instead')
const Project$json = {
  '1': 'Project',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    {'1': 'code', '3': 2, '4': 1, '5': 9, '10': 'code'},
    {'1': 'project_name', '3': 3, '4': 1, '5': 9, '10': 'projectName'},
    {'1': 'customer_code', '3': 4, '4': 1, '5': 9, '10': 'customerCode'},
    {'1': 'addresses', '3': 5, '4': 3, '5': 11, '6': '.nervatura.Address', '10': 'addresses'},
    {'1': 'contacts', '3': 6, '4': 3, '5': 11, '6': '.nervatura.Contact', '10': 'contacts'},
    {'1': 'events', '3': 7, '4': 3, '5': 11, '6': '.nervatura.Event', '10': 'events'},
    {'1': 'project_meta', '3': 8, '4': 1, '5': 11, '6': '.nervatura.ProjectMeta', '10': 'projectMeta'},
    {'1': 'project_map', '3': 9, '4': 1, '5': 11, '6': '.nervatura.JsonString', '10': 'projectMap'},
    {'1': 'time_stamp', '3': 10, '4': 1, '5': 9, '10': 'timeStamp'},
  ],
};

/// Descriptor for `Project`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List projectDescriptor = $convert.base64Decode(
    'CgdQcm9qZWN0Eg4KAmlkGAEgASgDUgJpZBISCgRjb2RlGAIgASgJUgRjb2RlEiEKDHByb2plY3'
    'RfbmFtZRgDIAEoCVILcHJvamVjdE5hbWUSIwoNY3VzdG9tZXJfY29kZRgEIAEoCVIMY3VzdG9t'
    'ZXJDb2RlEjAKCWFkZHJlc3NlcxgFIAMoCzISLm5lcnZhdHVyYS5BZGRyZXNzUglhZGRyZXNzZX'
    'MSLgoIY29udGFjdHMYBiADKAsyEi5uZXJ2YXR1cmEuQ29udGFjdFIIY29udGFjdHMSKAoGZXZl'
    'bnRzGAcgAygLMhAubmVydmF0dXJhLkV2ZW50UgZldmVudHMSOQoMcHJvamVjdF9tZXRhGAggAS'
    'gLMhYubmVydmF0dXJhLlByb2plY3RNZXRhUgtwcm9qZWN0TWV0YRI2Cgtwcm9qZWN0X21hcBgJ'
    'IAEoCzIVLm5lcnZhdHVyYS5Kc29uU3RyaW5nUgpwcm9qZWN0TWFwEh0KCnRpbWVfc3RhbXAYCi'
    'ABKAlSCXRpbWVTdGFtcA==');

@$core.Deprecated('Use projectsDescriptor instead')
const Projects$json = {
  '1': 'Projects',
  '2': [
    {'1': 'data', '3': 1, '4': 3, '5': 11, '6': '.nervatura.Project', '10': 'data'},
  ],
};

/// Descriptor for `Projects`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List projectsDescriptor = $convert.base64Decode(
    'CghQcm9qZWN0cxImCgRkYXRhGAEgAygLMhIubmVydmF0dXJhLlByb2plY3RSBGRhdGE=');

@$core.Deprecated('Use rateMetaDescriptor instead')
const RateMeta$json = {
  '1': 'RateMeta',
  '2': [
    {'1': 'rate_value', '3': 1, '4': 1, '5': 1, '10': 'rateValue'},
    {'1': 'tags', '3': 2, '4': 3, '5': 9, '10': 'tags'},
  ],
};

/// Descriptor for `RateMeta`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List rateMetaDescriptor = $convert.base64Decode(
    'CghSYXRlTWV0YRIdCgpyYXRlX3ZhbHVlGAEgASgBUglyYXRlVmFsdWUSEgoEdGFncxgCIAMoCV'
    'IEdGFncw==');

@$core.Deprecated('Use rateDescriptor instead')
const Rate$json = {
  '1': 'Rate',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    {'1': 'code', '3': 2, '4': 1, '5': 9, '10': 'code'},
    {'1': 'rate_type', '3': 3, '4': 1, '5': 14, '6': '.nervatura.RateType', '10': 'rateType'},
    {'1': 'rate_date', '3': 4, '4': 1, '5': 9, '10': 'rateDate'},
    {'1': 'place_code', '3': 5, '4': 1, '5': 9, '10': 'placeCode'},
    {'1': 'currency_code', '3': 6, '4': 1, '5': 9, '10': 'currencyCode'},
    {'1': 'rate_meta', '3': 7, '4': 1, '5': 11, '6': '.nervatura.RateMeta', '10': 'rateMeta'},
    {'1': 'rate_map', '3': 8, '4': 1, '5': 11, '6': '.nervatura.JsonString', '10': 'rateMap'},
    {'1': 'time_stamp', '3': 9, '4': 1, '5': 9, '10': 'timeStamp'},
  ],
};

/// Descriptor for `Rate`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List rateDescriptor = $convert.base64Decode(
    'CgRSYXRlEg4KAmlkGAEgASgDUgJpZBISCgRjb2RlGAIgASgJUgRjb2RlEjAKCXJhdGVfdHlwZR'
    'gDIAEoDjITLm5lcnZhdHVyYS5SYXRlVHlwZVIIcmF0ZVR5cGUSGwoJcmF0ZV9kYXRlGAQgASgJ'
    'UghyYXRlRGF0ZRIdCgpwbGFjZV9jb2RlGAUgASgJUglwbGFjZUNvZGUSIwoNY3VycmVuY3lfY2'
    '9kZRgGIAEoCVIMY3VycmVuY3lDb2RlEjAKCXJhdGVfbWV0YRgHIAEoCzITLm5lcnZhdHVyYS5S'
    'YXRlTWV0YVIIcmF0ZU1ldGESMAoIcmF0ZV9tYXAYCCABKAsyFS5uZXJ2YXR1cmEuSnNvblN0cm'
    'luZ1IHcmF0ZU1hcBIdCgp0aW1lX3N0YW1wGAkgASgJUgl0aW1lU3RhbXA=');

@$core.Deprecated('Use ratesDescriptor instead')
const Rates$json = {
  '1': 'Rates',
  '2': [
    {'1': 'data', '3': 1, '4': 3, '5': 11, '6': '.nervatura.Rate', '10': 'data'},
  ],
};

/// Descriptor for `Rates`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List ratesDescriptor = $convert.base64Decode(
    'CgVSYXRlcxIjCgRkYXRhGAEgAygLMg8ubmVydmF0dXJhLlJhdGVSBGRhdGE=');

@$core.Deprecated('Use taxMetaDescriptor instead')
const TaxMeta$json = {
  '1': 'TaxMeta',
  '2': [
    {'1': 'description', '3': 1, '4': 1, '5': 9, '10': 'description'},
    {'1': 'rate_value', '3': 2, '4': 1, '5': 1, '10': 'rateValue'},
    {'1': 'tags', '3': 3, '4': 3, '5': 9, '10': 'tags'},
  ],
};

/// Descriptor for `TaxMeta`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List taxMetaDescriptor = $convert.base64Decode(
    'CgdUYXhNZXRhEiAKC2Rlc2NyaXB0aW9uGAEgASgJUgtkZXNjcmlwdGlvbhIdCgpyYXRlX3ZhbH'
    'VlGAIgASgBUglyYXRlVmFsdWUSEgoEdGFncxgDIAMoCVIEdGFncw==');

@$core.Deprecated('Use taxDescriptor instead')
const Tax$json = {
  '1': 'Tax',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    {'1': 'code', '3': 2, '4': 1, '5': 9, '10': 'code'},
    {'1': 'tax_meta', '3': 3, '4': 1, '5': 11, '6': '.nervatura.TaxMeta', '10': 'taxMeta'},
    {'1': 'tax_map', '3': 4, '4': 1, '5': 11, '6': '.nervatura.JsonString', '10': 'taxMap'},
    {'1': 'time_stamp', '3': 5, '4': 1, '5': 9, '10': 'timeStamp'},
  ],
};

/// Descriptor for `Tax`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List taxDescriptor = $convert.base64Decode(
    'CgNUYXgSDgoCaWQYASABKANSAmlkEhIKBGNvZGUYAiABKAlSBGNvZGUSLQoIdGF4X21ldGEYAy'
    'ABKAsyEi5uZXJ2YXR1cmEuVGF4TWV0YVIHdGF4TWV0YRIuCgd0YXhfbWFwGAQgASgLMhUubmVy'
    'dmF0dXJhLkpzb25TdHJpbmdSBnRheE1hcBIdCgp0aW1lX3N0YW1wGAUgASgJUgl0aW1lU3RhbX'
    'A=');

@$core.Deprecated('Use taxesDescriptor instead')
const Taxes$json = {
  '1': 'Taxes',
  '2': [
    {'1': 'data', '3': 1, '4': 3, '5': 11, '6': '.nervatura.Tax', '10': 'data'},
  ],
};

/// Descriptor for `Taxes`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List taxesDescriptor = $convert.base64Decode(
    'CgVUYXhlcxIiCgRkYXRhGAEgAygLMg4ubmVydmF0dXJhLlRheFIEZGF0YQ==');

@$core.Deprecated('Use toolMetaDescriptor instead')
const ToolMeta$json = {
  '1': 'ToolMeta',
  '2': [
    {'1': 'serial_number', '3': 1, '4': 1, '5': 9, '10': 'serialNumber'},
    {'1': 'notes', '3': 2, '4': 1, '5': 9, '10': 'notes'},
    {'1': 'inactive', '3': 3, '4': 1, '5': 8, '10': 'inactive'},
    {'1': 'tags', '3': 4, '4': 3, '5': 9, '10': 'tags'},
  ],
};

/// Descriptor for `ToolMeta`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List toolMetaDescriptor = $convert.base64Decode(
    'CghUb29sTWV0YRIjCg1zZXJpYWxfbnVtYmVyGAEgASgJUgxzZXJpYWxOdW1iZXISFAoFbm90ZX'
    'MYAiABKAlSBW5vdGVzEhoKCGluYWN0aXZlGAMgASgIUghpbmFjdGl2ZRISCgR0YWdzGAQgAygJ'
    'UgR0YWdz');

@$core.Deprecated('Use toolDescriptor instead')
const Tool$json = {
  '1': 'Tool',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    {'1': 'code', '3': 2, '4': 1, '5': 9, '10': 'code'},
    {'1': 'description', '3': 3, '4': 1, '5': 9, '10': 'description'},
    {'1': 'product_code', '3': 4, '4': 1, '5': 9, '10': 'productCode'},
    {'1': 'events', '3': 5, '4': 3, '5': 11, '6': '.nervatura.Event', '10': 'events'},
    {'1': 'tool_meta', '3': 6, '4': 1, '5': 11, '6': '.nervatura.ToolMeta', '10': 'toolMeta'},
    {'1': 'tool_map', '3': 7, '4': 1, '5': 11, '6': '.nervatura.JsonString', '10': 'toolMap'},
    {'1': 'time_stamp', '3': 8, '4': 1, '5': 9, '10': 'timeStamp'},
  ],
};

/// Descriptor for `Tool`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List toolDescriptor = $convert.base64Decode(
    'CgRUb29sEg4KAmlkGAEgASgDUgJpZBISCgRjb2RlGAIgASgJUgRjb2RlEiAKC2Rlc2NyaXB0aW'
    '9uGAMgASgJUgtkZXNjcmlwdGlvbhIhCgxwcm9kdWN0X2NvZGUYBCABKAlSC3Byb2R1Y3RDb2Rl'
    'EigKBmV2ZW50cxgFIAMoCzIQLm5lcnZhdHVyYS5FdmVudFIGZXZlbnRzEjAKCXRvb2xfbWV0YR'
    'gGIAEoCzITLm5lcnZhdHVyYS5Ub29sTWV0YVIIdG9vbE1ldGESMAoIdG9vbF9tYXAYByABKAsy'
    'FS5uZXJ2YXR1cmEuSnNvblN0cmluZ1IHdG9vbE1hcBIdCgp0aW1lX3N0YW1wGAggASgJUgl0aW'
    '1lU3RhbXA=');

@$core.Deprecated('Use toolsDescriptor instead')
const Tools$json = {
  '1': 'Tools',
  '2': [
    {'1': 'data', '3': 1, '4': 3, '5': 11, '6': '.nervatura.Tool', '10': 'data'},
  ],
};

/// Descriptor for `Tools`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List toolsDescriptor = $convert.base64Decode(
    'CgVUb29scxIjCgRkYXRhGAEgAygLMg8ubmVydmF0dXJhLlRvb2xSBGRhdGE=');

@$core.Deprecated('Use worksheetDescriptor instead')
const Worksheet$json = {
  '1': 'Worksheet',
  '2': [
    {'1': 'distance', '3': 1, '4': 1, '5': 1, '10': 'distance'},
    {'1': 'repair', '3': 2, '4': 1, '5': 1, '10': 'repair'},
    {'1': 'total', '3': 3, '4': 1, '5': 1, '10': 'total'},
    {'1': 'justification', '3': 4, '4': 1, '5': 9, '10': 'justification'},
  ],
};

/// Descriptor for `Worksheet`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List worksheetDescriptor = $convert.base64Decode(
    'CglXb3Jrc2hlZXQSGgoIZGlzdGFuY2UYASABKAFSCGRpc3RhbmNlEhYKBnJlcGFpchgCIAEoAV'
    'IGcmVwYWlyEhQKBXRvdGFsGAMgASgBUgV0b3RhbBIkCg1qdXN0aWZpY2F0aW9uGAQgASgJUg1q'
    'dXN0aWZpY2F0aW9u');

@$core.Deprecated('Use rentDescriptor instead')
const Rent$json = {
  '1': 'Rent',
  '2': [
    {'1': 'holiday', '3': 1, '4': 1, '5': 1, '10': 'holiday'},
    {'1': 'bad_tool', '3': 2, '4': 1, '5': 1, '10': 'badTool'},
    {'1': 'other', '3': 3, '4': 1, '5': 1, '10': 'other'},
    {'1': 'justification', '3': 4, '4': 1, '5': 9, '10': 'justification'},
  ],
};

/// Descriptor for `Rent`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List rentDescriptor = $convert.base64Decode(
    'CgRSZW50EhgKB2hvbGlkYXkYASABKAFSB2hvbGlkYXkSGQoIYmFkX3Rvb2wYAiABKAFSB2JhZF'
    'Rvb2wSFAoFb3RoZXIYAyABKAFSBW90aGVyEiQKDWp1c3RpZmljYXRpb24YBCABKAlSDWp1c3Rp'
    'ZmljYXRpb24=');

@$core.Deprecated('Use invoiceDescriptor instead')
const Invoice$json = {
  '1': 'Invoice',
  '2': [
    {'1': 'company_name', '3': 1, '4': 1, '5': 9, '10': 'companyName'},
    {'1': 'company_address', '3': 2, '4': 1, '5': 9, '10': 'companyAddress'},
    {'1': 'company_tax_number', '3': 3, '4': 1, '5': 9, '10': 'companyTaxNumber'},
    {'1': 'customer_name', '3': 4, '4': 1, '5': 9, '10': 'customerName'},
    {'1': 'customer_address', '3': 5, '4': 1, '5': 9, '10': 'customerAddress'},
    {'1': 'customer_tax_number', '3': 6, '4': 1, '5': 9, '10': 'customerTaxNumber'},
  ],
};

/// Descriptor for `Invoice`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List invoiceDescriptor = $convert.base64Decode(
    'CgdJbnZvaWNlEiEKDGNvbXBhbnlfbmFtZRgBIAEoCVILY29tcGFueU5hbWUSJwoPY29tcGFueV'
    '9hZGRyZXNzGAIgASgJUg5jb21wYW55QWRkcmVzcxIsChJjb21wYW55X3RheF9udW1iZXIYAyAB'
    'KAlSEGNvbXBhbnlUYXhOdW1iZXISIwoNY3VzdG9tZXJfbmFtZRgEIAEoCVIMY3VzdG9tZXJOYW'
    '1lEikKEGN1c3RvbWVyX2FkZHJlc3MYBSABKAlSD2N1c3RvbWVyQWRkcmVzcxIuChNjdXN0b21l'
    'cl90YXhfbnVtYmVyGAYgASgJUhFjdXN0b21lclRheE51bWJlcg==');

@$core.Deprecated('Use transMetaDescriptor instead')
const TransMeta$json = {
  '1': 'TransMeta',
  '2': [
    {'1': 'due_time', '3': 1, '4': 1, '5': 9, '10': 'dueTime'},
    {'1': 'ref_number', '3': 2, '4': 1, '5': 9, '10': 'refNumber'},
    {'1': 'paid_type', '3': 3, '4': 1, '5': 14, '6': '.nervatura.PaidType', '10': 'paidType'},
    {'1': 'tax_free', '3': 4, '4': 1, '5': 8, '10': 'taxFree'},
    {'1': 'paid', '3': 5, '4': 1, '5': 8, '10': 'paid'},
    {'1': 'rate', '3': 6, '4': 1, '5': 1, '10': 'rate'},
    {'1': 'closed', '3': 7, '4': 1, '5': 8, '10': 'closed'},
    {'1': 'status', '3': 8, '4': 1, '5': 14, '6': '.nervatura.TransStatus', '10': 'status'},
    {'1': 'trans_state', '3': 9, '4': 1, '5': 14, '6': '.nervatura.TransState', '10': 'transState'},
    {'1': 'notes', '3': 10, '4': 1, '5': 9, '10': 'notes'},
    {'1': 'internal_notes', '3': 11, '4': 1, '5': 9, '10': 'internalNotes'},
    {'1': 'report_notes', '3': 12, '4': 1, '5': 9, '10': 'reportNotes'},
    {'1': 'worksheet', '3': 13, '4': 1, '5': 11, '6': '.nervatura.Worksheet', '10': 'worksheet'},
    {'1': 'rent', '3': 14, '4': 1, '5': 11, '6': '.nervatura.Rent', '10': 'rent'},
    {'1': 'invoice', '3': 15, '4': 1, '5': 11, '6': '.nervatura.Invoice', '10': 'invoice'},
    {'1': 'tags', '3': 16, '4': 3, '5': 9, '10': 'tags'},
  ],
};

/// Descriptor for `TransMeta`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List transMetaDescriptor = $convert.base64Decode(
    'CglUcmFuc01ldGESGQoIZHVlX3RpbWUYASABKAlSB2R1ZVRpbWUSHQoKcmVmX251bWJlchgCIA'
    'EoCVIJcmVmTnVtYmVyEjAKCXBhaWRfdHlwZRgDIAEoDjITLm5lcnZhdHVyYS5QYWlkVHlwZVII'
    'cGFpZFR5cGUSGQoIdGF4X2ZyZWUYBCABKAhSB3RheEZyZWUSEgoEcGFpZBgFIAEoCFIEcGFpZB'
    'ISCgRyYXRlGAYgASgBUgRyYXRlEhYKBmNsb3NlZBgHIAEoCFIGY2xvc2VkEi4KBnN0YXR1cxgI'
    'IAEoDjIWLm5lcnZhdHVyYS5UcmFuc1N0YXR1c1IGc3RhdHVzEjYKC3RyYW5zX3N0YXRlGAkgAS'
    'gOMhUubmVydmF0dXJhLlRyYW5zU3RhdGVSCnRyYW5zU3RhdGUSFAoFbm90ZXMYCiABKAlSBW5v'
    'dGVzEiUKDmludGVybmFsX25vdGVzGAsgASgJUg1pbnRlcm5hbE5vdGVzEiEKDHJlcG9ydF9ub3'
    'RlcxgMIAEoCVILcmVwb3J0Tm90ZXMSMgoJd29ya3NoZWV0GA0gASgLMhQubmVydmF0dXJhLldv'
    'cmtzaGVldFIJd29ya3NoZWV0EiMKBHJlbnQYDiABKAsyDy5uZXJ2YXR1cmEuUmVudFIEcmVudB'
    'IsCgdpbnZvaWNlGA8gASgLMhIubmVydmF0dXJhLkludm9pY2VSB2ludm9pY2USEgoEdGFncxgQ'
    'IAMoCVIEdGFncw==');

@$core.Deprecated('Use transDescriptor instead')
const Trans$json = {
  '1': 'Trans',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    {'1': 'code', '3': 2, '4': 1, '5': 9, '10': 'code'},
    {'1': 'trans_type', '3': 3, '4': 1, '5': 14, '6': '.nervatura.TransType', '10': 'transType'},
    {'1': 'trans_date', '3': 4, '4': 1, '5': 9, '10': 'transDate'},
    {'1': 'direction', '3': 5, '4': 1, '5': 14, '6': '.nervatura.Direction', '10': 'direction'},
    {'1': 'trans_code', '3': 6, '4': 1, '5': 9, '10': 'transCode'},
    {'1': 'customer_code', '3': 7, '4': 1, '5': 9, '10': 'customerCode'},
    {'1': 'employee_code', '3': 8, '4': 1, '5': 9, '10': 'employeeCode'},
    {'1': 'project_code', '3': 9, '4': 1, '5': 9, '10': 'projectCode'},
    {'1': 'place_code', '3': 10, '4': 1, '5': 9, '10': 'placeCode'},
    {'1': 'currency_code', '3': 11, '4': 1, '5': 9, '10': 'currencyCode'},
    {'1': 'auth_code', '3': 12, '4': 1, '5': 9, '10': 'authCode'},
    {'1': 'trans_meta', '3': 13, '4': 1, '5': 11, '6': '.nervatura.TransMeta', '10': 'transMeta'},
    {'1': 'trans_map', '3': 14, '4': 1, '5': 11, '6': '.nervatura.JsonString', '10': 'transMap'},
    {'1': 'time_stamp', '3': 15, '4': 1, '5': 9, '10': 'timeStamp'},
  ],
};

/// Descriptor for `Trans`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List transDescriptor = $convert.base64Decode(
    'CgVUcmFucxIOCgJpZBgBIAEoA1ICaWQSEgoEY29kZRgCIAEoCVIEY29kZRIzCgp0cmFuc190eX'
    'BlGAMgASgOMhQubmVydmF0dXJhLlRyYW5zVHlwZVIJdHJhbnNUeXBlEh0KCnRyYW5zX2RhdGUY'
    'BCABKAlSCXRyYW5zRGF0ZRIyCglkaXJlY3Rpb24YBSABKA4yFC5uZXJ2YXR1cmEuRGlyZWN0aW'
    '9uUglkaXJlY3Rpb24SHQoKdHJhbnNfY29kZRgGIAEoCVIJdHJhbnNDb2RlEiMKDWN1c3RvbWVy'
    'X2NvZGUYByABKAlSDGN1c3RvbWVyQ29kZRIjCg1lbXBsb3llZV9jb2RlGAggASgJUgxlbXBsb3'
    'llZUNvZGUSIQoMcHJvamVjdF9jb2RlGAkgASgJUgtwcm9qZWN0Q29kZRIdCgpwbGFjZV9jb2Rl'
    'GAogASgJUglwbGFjZUNvZGUSIwoNY3VycmVuY3lfY29kZRgLIAEoCVIMY3VycmVuY3lDb2RlEh'
    'sKCWF1dGhfY29kZRgMIAEoCVIIYXV0aENvZGUSMwoKdHJhbnNfbWV0YRgNIAEoCzIULm5lcnZh'
    'dHVyYS5UcmFuc01ldGFSCXRyYW5zTWV0YRIyCgl0cmFuc19tYXAYDiABKAsyFS5uZXJ2YXR1cm'
    'EuSnNvblN0cmluZ1IIdHJhbnNNYXASHQoKdGltZV9zdGFtcBgPIAEoCVIJdGltZVN0YW1w');

@$core.Deprecated('Use transactionsDescriptor instead')
const Transactions$json = {
  '1': 'Transactions',
  '2': [
    {'1': 'data', '3': 1, '4': 3, '5': 11, '6': '.nervatura.Trans', '10': 'data'},
  ],
};

/// Descriptor for `Transactions`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List transactionsDescriptor = $convert.base64Decode(
    'CgxUcmFuc2FjdGlvbnMSJAoEZGF0YRgBIAMoCzIQLm5lcnZhdHVyYS5UcmFuc1IEZGF0YQ==');

@$core.Deprecated('Use configDescriptor instead')
const Config$json = {
  '1': 'Config',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    {'1': 'code', '3': 2, '4': 1, '5': 9, '10': 'code'},
    {'1': 'config_type', '3': 3, '4': 1, '5': 14, '6': '.nervatura.ConfigType', '10': 'configType'},
    {'1': 'map', '3': 4, '4': 1, '5': 11, '6': '.nervatura.ConfigMap', '9': 0, '10': 'map'},
    {'1': 'shortcut', '3': 5, '4': 1, '5': 11, '6': '.nervatura.ConfigShortcut', '9': 0, '10': 'shortcut'},
    {'1': 'message', '3': 6, '4': 1, '5': 11, '6': '.nervatura.ConfigMessage', '9': 0, '10': 'message'},
    {'1': 'pattern', '3': 7, '4': 1, '5': 11, '6': '.nervatura.ConfigPattern', '9': 0, '10': 'pattern'},
    {'1': 'report', '3': 8, '4': 1, '5': 11, '6': '.nervatura.ConfigReport', '9': 0, '10': 'report'},
    {'1': 'print_queue', '3': 9, '4': 1, '5': 11, '6': '.nervatura.ConfigPrintQueue', '9': 0, '10': 'printQueue'},
    {'1': 'config_data', '3': 10, '4': 1, '5': 11, '6': '.nervatura.JsonString', '9': 0, '10': 'configData'},
    {'1': 'time_stamp', '3': 11, '4': 1, '5': 9, '10': 'timeStamp'},
  ],
  '8': [
    {'1': 'data'},
  ],
};

/// Descriptor for `Config`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List configDescriptor = $convert.base64Decode(
    'CgZDb25maWcSDgoCaWQYASABKANSAmlkEhIKBGNvZGUYAiABKAlSBGNvZGUSNgoLY29uZmlnX3'
    'R5cGUYAyABKA4yFS5uZXJ2YXR1cmEuQ29uZmlnVHlwZVIKY29uZmlnVHlwZRIoCgNtYXAYBCAB'
    'KAsyFC5uZXJ2YXR1cmEuQ29uZmlnTWFwSABSA21hcBI3CghzaG9ydGN1dBgFIAEoCzIZLm5lcn'
    'ZhdHVyYS5Db25maWdTaG9ydGN1dEgAUghzaG9ydGN1dBI0CgdtZXNzYWdlGAYgASgLMhgubmVy'
    'dmF0dXJhLkNvbmZpZ01lc3NhZ2VIAFIHbWVzc2FnZRI0CgdwYXR0ZXJuGAcgASgLMhgubmVydm'
    'F0dXJhLkNvbmZpZ1BhdHRlcm5IAFIHcGF0dGVybhIxCgZyZXBvcnQYCCABKAsyFy5uZXJ2YXR1'
    'cmEuQ29uZmlnUmVwb3J0SABSBnJlcG9ydBI+CgtwcmludF9xdWV1ZRgJIAEoCzIbLm5lcnZhdH'
    'VyYS5Db25maWdQcmludFF1ZXVlSABSCnByaW50UXVldWUSOAoLY29uZmlnX2RhdGEYCiABKAsy'
    'FS5uZXJ2YXR1cmEuSnNvblN0cmluZ0gAUgpjb25maWdEYXRhEh0KCnRpbWVfc3RhbXAYCyABKA'
    'lSCXRpbWVTdGFtcEIGCgRkYXRh');

@$core.Deprecated('Use configsDescriptor instead')
const Configs$json = {
  '1': 'Configs',
  '2': [
    {'1': 'data', '3': 1, '4': 3, '5': 11, '6': '.nervatura.Config', '10': 'data'},
  ],
};

/// Descriptor for `Configs`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List configsDescriptor = $convert.base64Decode(
    'CgdDb25maWdzEiUKBGRhdGEYASADKAsyES5uZXJ2YXR1cmEuQ29uZmlnUgRkYXRh');

@$core.Deprecated('Use configMapDescriptor instead')
const ConfigMap$json = {
  '1': 'ConfigMap',
  '2': [
    {'1': 'field_name', '3': 1, '4': 1, '5': 9, '10': 'fieldName'},
    {'1': 'field_type', '3': 2, '4': 1, '5': 14, '6': '.nervatura.FieldType', '10': 'fieldType'},
    {'1': 'description', '3': 3, '4': 1, '5': 9, '10': 'description'},
    {'1': 'tags', '3': 4, '4': 3, '5': 9, '10': 'tags'},
    {'1': 'filter', '3': 5, '4': 3, '5': 14, '6': '.nervatura.MapFilter', '10': 'filter'},
  ],
};

/// Descriptor for `ConfigMap`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List configMapDescriptor = $convert.base64Decode(
    'CglDb25maWdNYXASHQoKZmllbGRfbmFtZRgBIAEoCVIJZmllbGROYW1lEjMKCmZpZWxkX3R5cG'
    'UYAiABKA4yFC5uZXJ2YXR1cmEuRmllbGRUeXBlUglmaWVsZFR5cGUSIAoLZGVzY3JpcHRpb24Y'
    'AyABKAlSC2Rlc2NyaXB0aW9uEhIKBHRhZ3MYBCADKAlSBHRhZ3MSLAoGZmlsdGVyGAUgAygOMh'
    'QubmVydmF0dXJhLk1hcEZpbHRlclIGZmlsdGVy');

@$core.Deprecated('Use configShortcutFieldDescriptor instead')
const ConfigShortcutField$json = {
  '1': 'ConfigShortcutField',
  '2': [
    {'1': 'field_name', '3': 1, '4': 1, '5': 9, '10': 'fieldName'},
    {'1': 'description', '3': 2, '4': 1, '5': 9, '10': 'description'},
    {'1': 'field_type', '3': 3, '4': 1, '5': 14, '6': '.nervatura.ShortcutField', '10': 'fieldType'},
    {'1': 'order', '3': 4, '4': 1, '5': 3, '10': 'order'},
  ],
};

/// Descriptor for `ConfigShortcutField`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List configShortcutFieldDescriptor = $convert.base64Decode(
    'ChNDb25maWdTaG9ydGN1dEZpZWxkEh0KCmZpZWxkX25hbWUYASABKAlSCWZpZWxkTmFtZRIgCg'
    'tkZXNjcmlwdGlvbhgCIAEoCVILZGVzY3JpcHRpb24SNwoKZmllbGRfdHlwZRgDIAEoDjIYLm5l'
    'cnZhdHVyYS5TaG9ydGN1dEZpZWxkUglmaWVsZFR5cGUSFAoFb3JkZXIYBCABKANSBW9yZGVy');

@$core.Deprecated('Use configShortcutDescriptor instead')
const ConfigShortcut$json = {
  '1': 'ConfigShortcut',
  '2': [
    {'1': 'shortcut_key', '3': 1, '4': 1, '5': 9, '10': 'shortcutKey'},
    {'1': 'description', '3': 2, '4': 1, '5': 9, '10': 'description'},
    {'1': 'modul', '3': 3, '4': 1, '5': 9, '10': 'modul'},
    {'1': 'icon', '3': 4, '4': 1, '5': 9, '10': 'icon'},
    {'1': 'method', '3': 5, '4': 1, '5': 14, '6': '.nervatura.ShortcutMethod', '10': 'method'},
    {'1': 'funcname', '3': 6, '4': 1, '5': 9, '10': 'funcname'},
    {'1': 'address', '3': 7, '4': 1, '5': 9, '10': 'address'},
    {'1': 'fields', '3': 8, '4': 3, '5': 11, '6': '.nervatura.ConfigShortcutField', '10': 'fields'},
  ],
};

/// Descriptor for `ConfigShortcut`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List configShortcutDescriptor = $convert.base64Decode(
    'Cg5Db25maWdTaG9ydGN1dBIhCgxzaG9ydGN1dF9rZXkYASABKAlSC3Nob3J0Y3V0S2V5EiAKC2'
    'Rlc2NyaXB0aW9uGAIgASgJUgtkZXNjcmlwdGlvbhIUCgVtb2R1bBgDIAEoCVIFbW9kdWwSEgoE'
    'aWNvbhgEIAEoCVIEaWNvbhIxCgZtZXRob2QYBSABKA4yGS5uZXJ2YXR1cmEuU2hvcnRjdXRNZX'
    'Rob2RSBm1ldGhvZBIaCghmdW5jbmFtZRgGIAEoCVIIZnVuY25hbWUSGAoHYWRkcmVzcxgHIAEo'
    'CVIHYWRkcmVzcxI2CgZmaWVsZHMYCCADKAsyHi5uZXJ2YXR1cmEuQ29uZmlnU2hvcnRjdXRGaW'
    'VsZFIGZmllbGRz');

@$core.Deprecated('Use configMessageDescriptor instead')
const ConfigMessage$json = {
  '1': 'ConfigMessage',
  '2': [
    {'1': 'section', '3': 1, '4': 1, '5': 9, '10': 'section'},
    {'1': 'key', '3': 2, '4': 1, '5': 9, '10': 'key'},
    {'1': 'lang', '3': 3, '4': 1, '5': 9, '10': 'lang'},
    {'1': 'value', '3': 4, '4': 1, '5': 9, '10': 'value'},
  ],
};

/// Descriptor for `ConfigMessage`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List configMessageDescriptor = $convert.base64Decode(
    'Cg1Db25maWdNZXNzYWdlEhgKB3NlY3Rpb24YASABKAlSB3NlY3Rpb24SEAoDa2V5GAIgASgJUg'
    'NrZXkSEgoEbGFuZxgDIAEoCVIEbGFuZxIUCgV2YWx1ZRgEIAEoCVIFdmFsdWU=');

@$core.Deprecated('Use configPatternDescriptor instead')
const ConfigPattern$json = {
  '1': 'ConfigPattern',
  '2': [
    {'1': 'trans_type', '3': 1, '4': 1, '5': 14, '6': '.nervatura.TransType', '10': 'transType'},
    {'1': 'description', '3': 2, '4': 1, '5': 9, '10': 'description'},
    {'1': 'notes', '3': 3, '4': 1, '5': 9, '10': 'notes'},
    {'1': 'default_pattern', '3': 4, '4': 1, '5': 8, '10': 'defaultPattern'},
  ],
};

/// Descriptor for `ConfigPattern`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List configPatternDescriptor = $convert.base64Decode(
    'Cg1Db25maWdQYXR0ZXJuEjMKCnRyYW5zX3R5cGUYASABKA4yFC5uZXJ2YXR1cmEuVHJhbnNUeX'
    'BlUgl0cmFuc1R5cGUSIAoLZGVzY3JpcHRpb24YAiABKAlSC2Rlc2NyaXB0aW9uEhQKBW5vdGVz'
    'GAMgASgJUgVub3RlcxInCg9kZWZhdWx0X3BhdHRlcm4YBCABKAhSDmRlZmF1bHRQYXR0ZXJu');

@$core.Deprecated('Use configPrintQueueDescriptor instead')
const ConfigPrintQueue$json = {
  '1': 'ConfigPrintQueue',
  '2': [
    {'1': 'ref_type', '3': 1, '4': 1, '5': 9, '10': 'refType'},
    {'1': 'ref_code', '3': 2, '4': 1, '5': 9, '10': 'refCode'},
    {'1': 'qty', '3': 3, '4': 1, '5': 3, '10': 'qty'},
    {'1': 'report_code', '3': 4, '4': 1, '5': 9, '10': 'reportCode'},
    {'1': 'auth_code', '3': 5, '4': 1, '5': 9, '10': 'authCode'},
    {'1': 'time_stamp', '3': 6, '4': 1, '5': 9, '10': 'timeStamp'},
  ],
};

/// Descriptor for `ConfigPrintQueue`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List configPrintQueueDescriptor = $convert.base64Decode(
    'ChBDb25maWdQcmludFF1ZXVlEhkKCHJlZl90eXBlGAEgASgJUgdyZWZUeXBlEhkKCHJlZl9jb2'
    'RlGAIgASgJUgdyZWZDb2RlEhAKA3F0eRgDIAEoA1IDcXR5Eh8KC3JlcG9ydF9jb2RlGAQgASgJ'
    'UgpyZXBvcnRDb2RlEhsKCWF1dGhfY29kZRgFIAEoCVIIYXV0aENvZGUSHQoKdGltZV9zdGFtcB'
    'gGIAEoCVIJdGltZVN0YW1w');

@$core.Deprecated('Use configReportDescriptor instead')
const ConfigReport$json = {
  '1': 'ConfigReport',
  '2': [
    {'1': 'report_key', '3': 1, '4': 1, '5': 9, '10': 'reportKey'},
    {'1': 'report_type', '3': 2, '4': 1, '5': 9, '10': 'reportType'},
    {'1': 'trans_type', '3': 3, '4': 1, '5': 14, '6': '.nervatura.TransType', '9': 0, '10': 'transType', '17': true},
    {'1': 'direction', '3': 4, '4': 1, '5': 14, '6': '.nervatura.Direction', '9': 1, '10': 'direction', '17': true},
    {'1': 'file_type', '3': 5, '4': 1, '5': 14, '6': '.nervatura.FileType', '10': 'fileType'},
    {'1': 'report_name', '3': 6, '4': 1, '5': 9, '10': 'reportName'},
    {'1': 'description', '3': 7, '4': 1, '5': 9, '10': 'description'},
    {'1': 'label', '3': 8, '4': 1, '5': 9, '10': 'label'},
    {'1': 'template', '3': 9, '4': 1, '5': 9, '10': 'template'},
  ],
  '8': [
    {'1': '_trans_type'},
    {'1': '_direction'},
  ],
};

/// Descriptor for `ConfigReport`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List configReportDescriptor = $convert.base64Decode(
    'CgxDb25maWdSZXBvcnQSHQoKcmVwb3J0X2tleRgBIAEoCVIJcmVwb3J0S2V5Eh8KC3JlcG9ydF'
    '90eXBlGAIgASgJUgpyZXBvcnRUeXBlEjgKCnRyYW5zX3R5cGUYAyABKA4yFC5uZXJ2YXR1cmEu'
    'VHJhbnNUeXBlSABSCXRyYW5zVHlwZYgBARI3CglkaXJlY3Rpb24YBCABKA4yFC5uZXJ2YXR1cm'
    'EuRGlyZWN0aW9uSAFSCWRpcmVjdGlvbogBARIwCglmaWxlX3R5cGUYBSABKA4yEy5uZXJ2YXR1'
    'cmEuRmlsZVR5cGVSCGZpbGVUeXBlEh8KC3JlcG9ydF9uYW1lGAYgASgJUgpyZXBvcnROYW1lEi'
    'AKC2Rlc2NyaXB0aW9uGAcgASgJUgtkZXNjcmlwdGlvbhIUCgVsYWJlbBgIIAEoCVIFbGFiZWwS'
    'GgoIdGVtcGxhdGUYCSABKAlSCHRlbXBsYXRlQg0KC190cmFuc190eXBlQgwKCl9kaXJlY3Rpb2'
    '4=');

