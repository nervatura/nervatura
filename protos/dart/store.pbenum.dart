//
//  Generated code. Do not modify.
//  source: store.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class UserGroup extends $pb.ProtobufEnum {
  static const UserGroup GROUP_USER = UserGroup._(0, _omitEnumNames ? '' : 'GROUP_USER');
  static const UserGroup GROUP_ADMIN = UserGroup._(1, _omitEnumNames ? '' : 'GROUP_ADMIN');
  static const UserGroup GROUP_GUEST = UserGroup._(2, _omitEnumNames ? '' : 'GROUP_GUEST');

  static const $core.List<UserGroup> values = <UserGroup> [
    GROUP_USER,
    GROUP_ADMIN,
    GROUP_GUEST,
  ];

  static final $core.Map<$core.int, UserGroup> _byValue = $pb.ProtobufEnum.initByValue(values);
  static UserGroup? valueOf($core.int value) => _byValue[value];

  const UserGroup._($core.int v, $core.String n) : super(v, n);
}

class CustomerType extends $pb.ProtobufEnum {
  static const CustomerType CUSTOMER_COMPANY = CustomerType._(0, _omitEnumNames ? '' : 'CUSTOMER_COMPANY');
  static const CustomerType CUSTOMER_PRIVATE = CustomerType._(1, _omitEnumNames ? '' : 'CUSTOMER_PRIVATE');
  static const CustomerType CUSTOMER_OTHER = CustomerType._(2, _omitEnumNames ? '' : 'CUSTOMER_OTHER');
  static const CustomerType CUSTOMER_OWN = CustomerType._(3, _omitEnumNames ? '' : 'CUSTOMER_OWN');

  static const $core.List<CustomerType> values = <CustomerType> [
    CUSTOMER_COMPANY,
    CUSTOMER_PRIVATE,
    CUSTOMER_OTHER,
    CUSTOMER_OWN,
  ];

  static final $core.Map<$core.int, CustomerType> _byValue = $pb.ProtobufEnum.initByValue(values);
  static CustomerType? valueOf($core.int value) => _byValue[value];

  const CustomerType._($core.int v, $core.String n) : super(v, n);
}

class LinkType extends $pb.ProtobufEnum {
  static const LinkType LINK_CUSTOMER = LinkType._(0, _omitEnumNames ? '' : 'LINK_CUSTOMER');
  static const LinkType LINK_EMPLOYEE = LinkType._(1, _omitEnumNames ? '' : 'LINK_EMPLOYEE');
  static const LinkType LINK_ITEM = LinkType._(2, _omitEnumNames ? '' : 'LINK_ITEM');
  static const LinkType LINK_MOVEMENT = LinkType._(3, _omitEnumNames ? '' : 'LINK_MOVEMENT');
  static const LinkType LINK_PAYMENT = LinkType._(4, _omitEnumNames ? '' : 'LINK_PAYMENT');
  static const LinkType LINK_PLACE = LinkType._(5, _omitEnumNames ? '' : 'LINK_PLACE');
  static const LinkType LINK_PRODUCT = LinkType._(6, _omitEnumNames ? '' : 'LINK_PRODUCT');
  static const LinkType LINK_PROJECT = LinkType._(7, _omitEnumNames ? '' : 'LINK_PROJECT');
  static const LinkType LINK_TOOL = LinkType._(8, _omitEnumNames ? '' : 'LINK_TOOL');
  static const LinkType LINK_TRANS = LinkType._(9, _omitEnumNames ? '' : 'LINK_TRANS');

  static const $core.List<LinkType> values = <LinkType> [
    LINK_CUSTOMER,
    LINK_EMPLOYEE,
    LINK_ITEM,
    LINK_MOVEMENT,
    LINK_PAYMENT,
    LINK_PLACE,
    LINK_PRODUCT,
    LINK_PROJECT,
    LINK_TOOL,
    LINK_TRANS,
  ];

  static final $core.Map<$core.int, LinkType> _byValue = $pb.ProtobufEnum.initByValue(values);
  static LinkType? valueOf($core.int value) => _byValue[value];

  const LinkType._($core.int v, $core.String n) : super(v, n);
}

class LogType extends $pb.ProtobufEnum {
  static const LogType LOG_INSERT = LogType._(0, _omitEnumNames ? '' : 'LOG_INSERT');
  static const LogType LOG_UPDATE = LogType._(1, _omitEnumNames ? '' : 'LOG_UPDATE');
  static const LogType LOG_DELETE = LogType._(2, _omitEnumNames ? '' : 'LOG_DELETE');

  static const $core.List<LogType> values = <LogType> [
    LOG_INSERT,
    LOG_UPDATE,
    LOG_DELETE,
  ];

  static final $core.Map<$core.int, LogType> _byValue = $pb.ProtobufEnum.initByValue(values);
  static LogType? valueOf($core.int value) => _byValue[value];

  const LogType._($core.int v, $core.String n) : super(v, n);
}

class MovementType extends $pb.ProtobufEnum {
  static const MovementType MOVEMENT_INVENTORY = MovementType._(0, _omitEnumNames ? '' : 'MOVEMENT_INVENTORY');
  static const MovementType MOVEMENT_TOOL = MovementType._(1, _omitEnumNames ? '' : 'MOVEMENT_TOOL');
  static const MovementType MOVEMENT_PLAN = MovementType._(2, _omitEnumNames ? '' : 'MOVEMENT_PLAN');
  static const MovementType MOVEMENT_HEAD = MovementType._(3, _omitEnumNames ? '' : 'MOVEMENT_HEAD');

  static const $core.List<MovementType> values = <MovementType> [
    MOVEMENT_INVENTORY,
    MOVEMENT_TOOL,
    MOVEMENT_PLAN,
    MOVEMENT_HEAD,
  ];

  static final $core.Map<$core.int, MovementType> _byValue = $pb.ProtobufEnum.initByValue(values);
  static MovementType? valueOf($core.int value) => _byValue[value];

  const MovementType._($core.int v, $core.String n) : super(v, n);
}

class PlaceType extends $pb.ProtobufEnum {
  static const PlaceType PLACE_WAREHOUSE = PlaceType._(0, _omitEnumNames ? '' : 'PLACE_WAREHOUSE');
  static const PlaceType PLACE_BANK = PlaceType._(1, _omitEnumNames ? '' : 'PLACE_BANK');
  static const PlaceType PLACE_CASH = PlaceType._(2, _omitEnumNames ? '' : 'PLACE_CASH');
  static const PlaceType PLACE_OTHER = PlaceType._(3, _omitEnumNames ? '' : 'PLACE_OTHER');

  static const $core.List<PlaceType> values = <PlaceType> [
    PLACE_WAREHOUSE,
    PLACE_BANK,
    PLACE_CASH,
    PLACE_OTHER,
  ];

  static final $core.Map<$core.int, PlaceType> _byValue = $pb.ProtobufEnum.initByValue(values);
  static PlaceType? valueOf($core.int value) => _byValue[value];

  const PlaceType._($core.int v, $core.String n) : super(v, n);
}

class PriceType extends $pb.ProtobufEnum {
  static const PriceType PRICE_CUSTOMER = PriceType._(0, _omitEnumNames ? '' : 'PRICE_CUSTOMER');
  static const PriceType PRICE_VENDOR = PriceType._(1, _omitEnumNames ? '' : 'PRICE_VENDOR');

  static const $core.List<PriceType> values = <PriceType> [
    PRICE_CUSTOMER,
    PRICE_VENDOR,
  ];

  static final $core.Map<$core.int, PriceType> _byValue = $pb.ProtobufEnum.initByValue(values);
  static PriceType? valueOf($core.int value) => _byValue[value];

  const PriceType._($core.int v, $core.String n) : super(v, n);
}

class BarcodeType extends $pb.ProtobufEnum {
  static const BarcodeType BARCODE_CODE_39 = BarcodeType._(0, _omitEnumNames ? '' : 'BARCODE_CODE_39');
  static const BarcodeType BARCODE_CODE_128 = BarcodeType._(1, _omitEnumNames ? '' : 'BARCODE_CODE_128');
  static const BarcodeType BARCODE_EAN_8 = BarcodeType._(2, _omitEnumNames ? '' : 'BARCODE_EAN_8');
  static const BarcodeType BARCODE_EAN_13 = BarcodeType._(3, _omitEnumNames ? '' : 'BARCODE_EAN_13');
  static const BarcodeType BARCODE_QR_CODE = BarcodeType._(4, _omitEnumNames ? '' : 'BARCODE_QR_CODE');

  static const $core.List<BarcodeType> values = <BarcodeType> [
    BARCODE_CODE_39,
    BARCODE_CODE_128,
    BARCODE_EAN_8,
    BARCODE_EAN_13,
    BARCODE_QR_CODE,
  ];

  static final $core.Map<$core.int, BarcodeType> _byValue = $pb.ProtobufEnum.initByValue(values);
  static BarcodeType? valueOf($core.int value) => _byValue[value];

  const BarcodeType._($core.int v, $core.String n) : super(v, n);
}

class ProductType extends $pb.ProtobufEnum {
  static const ProductType PRODUCT_ITEM = ProductType._(0, _omitEnumNames ? '' : 'PRODUCT_ITEM');
  static const ProductType PRODUCT_SERVICE = ProductType._(1, _omitEnumNames ? '' : 'PRODUCT_SERVICE');
  static const ProductType PRODUCT_VIRTUAL = ProductType._(2, _omitEnumNames ? '' : 'PRODUCT_VIRTUAL');

  static const $core.List<ProductType> values = <ProductType> [
    PRODUCT_ITEM,
    PRODUCT_SERVICE,
    PRODUCT_VIRTUAL,
  ];

  static final $core.Map<$core.int, ProductType> _byValue = $pb.ProtobufEnum.initByValue(values);
  static ProductType? valueOf($core.int value) => _byValue[value];

  const ProductType._($core.int v, $core.String n) : super(v, n);
}

class RateType extends $pb.ProtobufEnum {
  static const RateType RATE_RATE = RateType._(0, _omitEnumNames ? '' : 'RATE_RATE');
  static const RateType RATE_BUY = RateType._(1, _omitEnumNames ? '' : 'RATE_BUY');
  static const RateType RATE_SELL = RateType._(2, _omitEnumNames ? '' : 'RATE_SELL');
  static const RateType RATE_AVERAGE = RateType._(3, _omitEnumNames ? '' : 'RATE_AVERAGE');

  static const $core.List<RateType> values = <RateType> [
    RATE_RATE,
    RATE_BUY,
    RATE_SELL,
    RATE_AVERAGE,
  ];

  static final $core.Map<$core.int, RateType> _byValue = $pb.ProtobufEnum.initByValue(values);
  static RateType? valueOf($core.int value) => _byValue[value];

  const RateType._($core.int v, $core.String n) : super(v, n);
}

class PaidType extends $pb.ProtobufEnum {
  static const PaidType PAID_ONLINE = PaidType._(0, _omitEnumNames ? '' : 'PAID_ONLINE');
  static const PaidType PAID_CARD = PaidType._(1, _omitEnumNames ? '' : 'PAID_CARD');
  static const PaidType PAID_TRANSFER = PaidType._(2, _omitEnumNames ? '' : 'PAID_TRANSFER');
  static const PaidType PAID_CASH = PaidType._(3, _omitEnumNames ? '' : 'PAID_CASH');
  static const PaidType PAID_OTHER = PaidType._(4, _omitEnumNames ? '' : 'PAID_OTHER');

  static const $core.List<PaidType> values = <PaidType> [
    PAID_ONLINE,
    PAID_CARD,
    PAID_TRANSFER,
    PAID_CASH,
    PAID_OTHER,
  ];

  static final $core.Map<$core.int, PaidType> _byValue = $pb.ProtobufEnum.initByValue(values);
  static PaidType? valueOf($core.int value) => _byValue[value];

  const PaidType._($core.int v, $core.String n) : super(v, n);
}

class TransStatus extends $pb.ProtobufEnum {
  static const TransStatus TRANS_NORMAL = TransStatus._(0, _omitEnumNames ? '' : 'TRANS_NORMAL');
  static const TransStatus TRANS_CANCELLATION = TransStatus._(1, _omitEnumNames ? '' : 'TRANS_CANCELLATION');
  static const TransStatus TRANS_AMENDMENT = TransStatus._(2, _omitEnumNames ? '' : 'TRANS_AMENDMENT');

  static const $core.List<TransStatus> values = <TransStatus> [
    TRANS_NORMAL,
    TRANS_CANCELLATION,
    TRANS_AMENDMENT,
  ];

  static final $core.Map<$core.int, TransStatus> _byValue = $pb.ProtobufEnum.initByValue(values);
  static TransStatus? valueOf($core.int value) => _byValue[value];

  const TransStatus._($core.int v, $core.String n) : super(v, n);
}

class TransState extends $pb.ProtobufEnum {
  static const TransState TRANS_OK = TransState._(0, _omitEnumNames ? '' : 'TRANS_OK');
  static const TransState TRANS_NEW = TransState._(1, _omitEnumNames ? '' : 'TRANS_NEW');
  static const TransState TRANS_BACK = TransState._(2, _omitEnumNames ? '' : 'TRANS_BACK');

  static const $core.List<TransState> values = <TransState> [
    TRANS_OK,
    TRANS_NEW,
    TRANS_BACK,
  ];

  static final $core.Map<$core.int, TransState> _byValue = $pb.ProtobufEnum.initByValue(values);
  static TransState? valueOf($core.int value) => _byValue[value];

  const TransState._($core.int v, $core.String n) : super(v, n);
}

class Direction extends $pb.ProtobufEnum {
  static const Direction DIRECTION_OUT = Direction._(0, _omitEnumNames ? '' : 'DIRECTION_OUT');
  static const Direction DIRECTION_IN = Direction._(1, _omitEnumNames ? '' : 'DIRECTION_IN');
  static const Direction DIRECTION_TRANSFER = Direction._(2, _omitEnumNames ? '' : 'DIRECTION_TRANSFER');

  static const $core.List<Direction> values = <Direction> [
    DIRECTION_OUT,
    DIRECTION_IN,
    DIRECTION_TRANSFER,
  ];

  static final $core.Map<$core.int, Direction> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Direction? valueOf($core.int value) => _byValue[value];

  const Direction._($core.int v, $core.String n) : super(v, n);
}

class TransType extends $pb.ProtobufEnum {
  static const TransType TRANS_INVOICE = TransType._(0, _omitEnumNames ? '' : 'TRANS_INVOICE');
  static const TransType TRANS_RECEIPT = TransType._(1, _omitEnumNames ? '' : 'TRANS_RECEIPT');
  static const TransType TRANS_ORDER = TransType._(2, _omitEnumNames ? '' : 'TRANS_ORDER');
  static const TransType TRANS_OFFER = TransType._(3, _omitEnumNames ? '' : 'TRANS_OFFER');
  static const TransType TRANS_WORKSHEET = TransType._(4, _omitEnumNames ? '' : 'TRANS_WORKSHEET');
  static const TransType TRANS_RENT = TransType._(5, _omitEnumNames ? '' : 'TRANS_RENT');
  static const TransType TRANS_DELIVERY = TransType._(6, _omitEnumNames ? '' : 'TRANS_DELIVERY');
  static const TransType TRANS_INVENTORY = TransType._(7, _omitEnumNames ? '' : 'TRANS_INVENTORY');
  static const TransType TRANS_WAYBILL = TransType._(8, _omitEnumNames ? '' : 'TRANS_WAYBILL');
  static const TransType TRANS_PRODUCTION = TransType._(9, _omitEnumNames ? '' : 'TRANS_PRODUCTION');
  static const TransType TRANS_FORMULA = TransType._(10, _omitEnumNames ? '' : 'TRANS_FORMULA');
  static const TransType TRANS_BANK = TransType._(11, _omitEnumNames ? '' : 'TRANS_BANK');
  static const TransType TRANS_CASH = TransType._(12, _omitEnumNames ? '' : 'TRANS_CASH');

  static const $core.List<TransType> values = <TransType> [
    TRANS_INVOICE,
    TRANS_RECEIPT,
    TRANS_ORDER,
    TRANS_OFFER,
    TRANS_WORKSHEET,
    TRANS_RENT,
    TRANS_DELIVERY,
    TRANS_INVENTORY,
    TRANS_WAYBILL,
    TRANS_PRODUCTION,
    TRANS_FORMULA,
    TRANS_BANK,
    TRANS_CASH,
  ];

  static final $core.Map<$core.int, TransType> _byValue = $pb.ProtobufEnum.initByValue(values);
  static TransType? valueOf($core.int value) => _byValue[value];

  const TransType._($core.int v, $core.String n) : super(v, n);
}

class ConfigType extends $pb.ProtobufEnum {
  static const ConfigType CONFIG_MAP = ConfigType._(0, _omitEnumNames ? '' : 'CONFIG_MAP');
  static const ConfigType CONFIG_SHORTCUT = ConfigType._(1, _omitEnumNames ? '' : 'CONFIG_SHORTCUT');
  static const ConfigType CONFIG_MESSAGE = ConfigType._(2, _omitEnumNames ? '' : 'CONFIG_MESSAGE');
  static const ConfigType CONFIG_PATTERN = ConfigType._(3, _omitEnumNames ? '' : 'CONFIG_PATTERN');
  static const ConfigType CONFIG_REPORT = ConfigType._(4, _omitEnumNames ? '' : 'CONFIG_REPORT');
  static const ConfigType CONFIG_PRINT_QUEUE = ConfigType._(5, _omitEnumNames ? '' : 'CONFIG_PRINT_QUEUE');
  static const ConfigType CONFIG_DATA = ConfigType._(6, _omitEnumNames ? '' : 'CONFIG_DATA');

  static const $core.List<ConfigType> values = <ConfigType> [
    CONFIG_MAP,
    CONFIG_SHORTCUT,
    CONFIG_MESSAGE,
    CONFIG_PATTERN,
    CONFIG_REPORT,
    CONFIG_PRINT_QUEUE,
    CONFIG_DATA,
  ];

  static final $core.Map<$core.int, ConfigType> _byValue = $pb.ProtobufEnum.initByValue(values);
  static ConfigType? valueOf($core.int value) => _byValue[value];

  const ConfigType._($core.int v, $core.String n) : super(v, n);
}

class FieldType extends $pb.ProtobufEnum {
  static const FieldType FIELD_STRING = FieldType._(0, _omitEnumNames ? '' : 'FIELD_STRING');
  static const FieldType FIELD_BOOL = FieldType._(1, _omitEnumNames ? '' : 'FIELD_BOOL');
  static const FieldType FIELD_INTEGER = FieldType._(2, _omitEnumNames ? '' : 'FIELD_INTEGER');
  static const FieldType FIELD_NUMBER = FieldType._(3, _omitEnumNames ? '' : 'FIELD_NUMBER');
  static const FieldType FIELD_DATE = FieldType._(4, _omitEnumNames ? '' : 'FIELD_DATE');
  static const FieldType FIELD_DATETIME = FieldType._(5, _omitEnumNames ? '' : 'FIELD_DATETIME');
  static const FieldType FIELD_MEMO = FieldType._(6, _omitEnumNames ? '' : 'FIELD_MEMO');
  static const FieldType FIELD_ENUM = FieldType._(7, _omitEnumNames ? '' : 'FIELD_ENUM');
  static const FieldType FIELD_URL = FieldType._(8, _omitEnumNames ? '' : 'FIELD_URL');
  static const FieldType FIELD_CUSTOMER = FieldType._(9, _omitEnumNames ? '' : 'FIELD_CUSTOMER');
  static const FieldType FIELD_EMPLOYEE = FieldType._(10, _omitEnumNames ? '' : 'FIELD_EMPLOYEE');
  static const FieldType FIELD_PLACE = FieldType._(11, _omitEnumNames ? '' : 'FIELD_PLACE');
  static const FieldType FIELD_PRODUCT = FieldType._(12, _omitEnumNames ? '' : 'FIELD_PRODUCT');
  static const FieldType FIELD_PROJECT = FieldType._(13, _omitEnumNames ? '' : 'FIELD_PROJECT');
  static const FieldType FIELD_TOOL = FieldType._(14, _omitEnumNames ? '' : 'FIELD_TOOL');
  static const FieldType FIELD_TRANS = FieldType._(15, _omitEnumNames ? '' : 'FIELD_TRANS');
  static const FieldType FIELD_TRANS_ITEM = FieldType._(16, _omitEnumNames ? '' : 'FIELD_TRANS_ITEM');
  static const FieldType FIELD_TRANS_MOVEMENT = FieldType._(17, _omitEnumNames ? '' : 'FIELD_TRANS_MOVEMENT');
  static const FieldType FIELD_TRANS_PAYMENT = FieldType._(18, _omitEnumNames ? '' : 'FIELD_TRANS_PAYMENT');

  static const $core.List<FieldType> values = <FieldType> [
    FIELD_STRING,
    FIELD_BOOL,
    FIELD_INTEGER,
    FIELD_NUMBER,
    FIELD_DATE,
    FIELD_DATETIME,
    FIELD_MEMO,
    FIELD_ENUM,
    FIELD_URL,
    FIELD_CUSTOMER,
    FIELD_EMPLOYEE,
    FIELD_PLACE,
    FIELD_PRODUCT,
    FIELD_PROJECT,
    FIELD_TOOL,
    FIELD_TRANS,
    FIELD_TRANS_ITEM,
    FIELD_TRANS_MOVEMENT,
    FIELD_TRANS_PAYMENT,
  ];

  static final $core.Map<$core.int, FieldType> _byValue = $pb.ProtobufEnum.initByValue(values);
  static FieldType? valueOf($core.int value) => _byValue[value];

  const FieldType._($core.int v, $core.String n) : super(v, n);
}

class MapFilter extends $pb.ProtobufEnum {
  static const MapFilter FILTER_ADDRESS = MapFilter._(0, _omitEnumNames ? '' : 'FILTER_ADDRESS');
  static const MapFilter FILTER_BARCODE = MapFilter._(1, _omitEnumNames ? '' : 'FILTER_BARCODE');
  static const MapFilter FILTER_CONTACT = MapFilter._(2, _omitEnumNames ? '' : 'FILTER_CONTACT');
  static const MapFilter FILTER_CURRENCY = MapFilter._(3, _omitEnumNames ? '' : 'FILTER_CURRENCY');
  static const MapFilter FILTER_CUSTOMER = MapFilter._(4, _omitEnumNames ? '' : 'FILTER_CUSTOMER');
  static const MapFilter FILTER_EMPLOYEE = MapFilter._(5, _omitEnumNames ? '' : 'FILTER_EMPLOYEE');
  static const MapFilter FILTER_EVENT = MapFilter._(6, _omitEnumNames ? '' : 'FILTER_EVENT');
  static const MapFilter FILTER_ITEM = MapFilter._(7, _omitEnumNames ? '' : 'FILTER_ITEM');
  static const MapFilter FILTER_MOVEMENT = MapFilter._(8, _omitEnumNames ? '' : 'FILTER_MOVEMENT');
  static const MapFilter FILTER_PAYMENT = MapFilter._(9, _omitEnumNames ? '' : 'FILTER_PAYMENT');
  static const MapFilter FILTER_PLACE = MapFilter._(10, _omitEnumNames ? '' : 'FILTER_PLACE');
  static const MapFilter FILTER_PRICE = MapFilter._(11, _omitEnumNames ? '' : 'FILTER_PRICE');
  static const MapFilter FILTER_PRODUCT = MapFilter._(12, _omitEnumNames ? '' : 'FILTER_PRODUCT');
  static const MapFilter FILTER_PROJECT = MapFilter._(13, _omitEnumNames ? '' : 'FILTER_PROJECT');
  static const MapFilter FILTER_RATE = MapFilter._(14, _omitEnumNames ? '' : 'FILTER_RATE');
  static const MapFilter FILTER_TAX = MapFilter._(15, _omitEnumNames ? '' : 'FILTER_TAX');
  static const MapFilter FILTER_TOOL = MapFilter._(16, _omitEnumNames ? '' : 'FILTER_TOOL');
  static const MapFilter FILTER_USER = MapFilter._(17, _omitEnumNames ? '' : 'FILTER_USER');
  static const MapFilter FILTER_TRANS = MapFilter._(18, _omitEnumNames ? '' : 'FILTER_TRANS');
  static const MapFilter FILTER_INVOICE = MapFilter._(19, _omitEnumNames ? '' : 'FILTER_INVOICE');
  static const MapFilter FILTER_RECEIPT = MapFilter._(20, _omitEnumNames ? '' : 'FILTER_RECEIPT');
  static const MapFilter FILTER_ORDER = MapFilter._(21, _omitEnumNames ? '' : 'FILTER_ORDER');
  static const MapFilter FILTER_OFFER = MapFilter._(22, _omitEnumNames ? '' : 'FILTER_OFFER');
  static const MapFilter FILTER_WORKSHEET = MapFilter._(23, _omitEnumNames ? '' : 'FILTER_WORKSHEET');
  static const MapFilter FILTER_RENT = MapFilter._(24, _omitEnumNames ? '' : 'FILTER_RENT');
  static const MapFilter FILTER_DELIVERY = MapFilter._(25, _omitEnumNames ? '' : 'FILTER_DELIVERY');
  static const MapFilter FILTER_INVENTORY = MapFilter._(26, _omitEnumNames ? '' : 'FILTER_INVENTORY');
  static const MapFilter FILTER_WAYBILL = MapFilter._(27, _omitEnumNames ? '' : 'FILTER_WAYBILL');
  static const MapFilter FILTER_PRODUCTION = MapFilter._(28, _omitEnumNames ? '' : 'FILTER_PRODUCTION');
  static const MapFilter FILTER_FORMULA = MapFilter._(29, _omitEnumNames ? '' : 'FILTER_FORMULA');
  static const MapFilter FILTER_BANK = MapFilter._(30, _omitEnumNames ? '' : 'FILTER_BANK');
  static const MapFilter FILTER_CASH = MapFilter._(31, _omitEnumNames ? '' : 'FILTER_CASH');

  static const $core.List<MapFilter> values = <MapFilter> [
    FILTER_ADDRESS,
    FILTER_BARCODE,
    FILTER_CONTACT,
    FILTER_CURRENCY,
    FILTER_CUSTOMER,
    FILTER_EMPLOYEE,
    FILTER_EVENT,
    FILTER_ITEM,
    FILTER_MOVEMENT,
    FILTER_PAYMENT,
    FILTER_PLACE,
    FILTER_PRICE,
    FILTER_PRODUCT,
    FILTER_PROJECT,
    FILTER_RATE,
    FILTER_TAX,
    FILTER_TOOL,
    FILTER_USER,
    FILTER_TRANS,
    FILTER_INVOICE,
    FILTER_RECEIPT,
    FILTER_ORDER,
    FILTER_OFFER,
    FILTER_WORKSHEET,
    FILTER_RENT,
    FILTER_DELIVERY,
    FILTER_INVENTORY,
    FILTER_WAYBILL,
    FILTER_PRODUCTION,
    FILTER_FORMULA,
    FILTER_BANK,
    FILTER_CASH,
  ];

  static final $core.Map<$core.int, MapFilter> _byValue = $pb.ProtobufEnum.initByValue(values);
  static MapFilter? valueOf($core.int value) => _byValue[value];

  const MapFilter._($core.int v, $core.String n) : super(v, n);
}

class ShortcutMethod extends $pb.ProtobufEnum {
  static const ShortcutMethod METHOD_GET = ShortcutMethod._(0, _omitEnumNames ? '' : 'METHOD_GET');
  static const ShortcutMethod METHOD_POST = ShortcutMethod._(1, _omitEnumNames ? '' : 'METHOD_POST');

  static const $core.List<ShortcutMethod> values = <ShortcutMethod> [
    METHOD_GET,
    METHOD_POST,
  ];

  static final $core.Map<$core.int, ShortcutMethod> _byValue = $pb.ProtobufEnum.initByValue(values);
  static ShortcutMethod? valueOf($core.int value) => _byValue[value];

  const ShortcutMethod._($core.int v, $core.String n) : super(v, n);
}

class ShortcutField extends $pb.ProtobufEnum {
  static const ShortcutField SHORTCUT_STRING = ShortcutField._(0, _omitEnumNames ? '' : 'SHORTCUT_STRING');
  static const ShortcutField SHORTCUT_BOOL = ShortcutField._(1, _omitEnumNames ? '' : 'SHORTCUT_BOOL');
  static const ShortcutField SHORTCUT_INTEGER = ShortcutField._(2, _omitEnumNames ? '' : 'SHORTCUT_INTEGER');
  static const ShortcutField SHORTCUT_NUMBER = ShortcutField._(3, _omitEnumNames ? '' : 'SHORTCUT_NUMBER');
  static const ShortcutField SHORTCUT_DATE = ShortcutField._(4, _omitEnumNames ? '' : 'SHORTCUT_DATE');

  static const $core.List<ShortcutField> values = <ShortcutField> [
    SHORTCUT_STRING,
    SHORTCUT_BOOL,
    SHORTCUT_INTEGER,
    SHORTCUT_NUMBER,
    SHORTCUT_DATE,
  ];

  static final $core.Map<$core.int, ShortcutField> _byValue = $pb.ProtobufEnum.initByValue(values);
  static ShortcutField? valueOf($core.int value) => _byValue[value];

  const ShortcutField._($core.int v, $core.String n) : super(v, n);
}

class FileType extends $pb.ProtobufEnum {
  static const FileType FILE_PDF = FileType._(0, _omitEnumNames ? '' : 'FILE_PDF');
  static const FileType FILE_CSV = FileType._(1, _omitEnumNames ? '' : 'FILE_CSV');

  static const $core.List<FileType> values = <FileType> [
    FILE_PDF,
    FILE_CSV,
  ];

  static final $core.Map<$core.int, FileType> _byValue = $pb.ProtobufEnum.initByValue(values);
  static FileType? valueOf($core.int value) => _byValue[value];

  const FileType._($core.int v, $core.String n) : super(v, n);
}

class Model extends $pb.ProtobufEnum {
  static const Model AUTH = Model._(0, _omitEnumNames ? '' : 'AUTH');
  static const Model CONFIG = Model._(1, _omitEnumNames ? '' : 'CONFIG');
  static const Model CURRENCY = Model._(2, _omitEnumNames ? '' : 'CURRENCY');
  static const Model CUSTOMER = Model._(3, _omitEnumNames ? '' : 'CUSTOMER');
  static const Model EMPLOYEE = Model._(4, _omitEnumNames ? '' : 'EMPLOYEE');
  static const Model ITEM = Model._(5, _omitEnumNames ? '' : 'ITEM');
  static const Model LINK = Model._(6, _omitEnumNames ? '' : 'LINK');
  static const Model MOVEMENT = Model._(7, _omitEnumNames ? '' : 'MOVEMENT');
  static const Model PAYMENT = Model._(8, _omitEnumNames ? '' : 'PAYMENT');
  static const Model PLACE = Model._(9, _omitEnumNames ? '' : 'PLACE');
  static const Model PRICE = Model._(10, _omitEnumNames ? '' : 'PRICE');
  static const Model PRODUCT = Model._(11, _omitEnumNames ? '' : 'PRODUCT');
  static const Model PROJECT = Model._(12, _omitEnumNames ? '' : 'PROJECT');
  static const Model RATE = Model._(13, _omitEnumNames ? '' : 'RATE');
  static const Model TAX = Model._(14, _omitEnumNames ? '' : 'TAX');
  static const Model TOOL = Model._(15, _omitEnumNames ? '' : 'TOOL');
  static const Model TRANS = Model._(16, _omitEnumNames ? '' : 'TRANS');

  static const $core.List<Model> values = <Model> [
    AUTH,
    CONFIG,
    CURRENCY,
    CUSTOMER,
    EMPLOYEE,
    ITEM,
    LINK,
    MOVEMENT,
    PAYMENT,
    PLACE,
    PRICE,
    PRODUCT,
    PROJECT,
    RATE,
    TAX,
    TOOL,
    TRANS,
  ];

  static final $core.Map<$core.int, Model> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Model? valueOf($core.int value) => _byValue[value];

  const Model._($core.int v, $core.String n) : super(v, n);
}

class ViewName extends $pb.ProtobufEnum {
  static const ViewName VIEW_CONFIG_DATA = ViewName._(0, _omitEnumNames ? '' : 'VIEW_CONFIG_DATA');
  static const ViewName VIEW_CONFIG_MAP = ViewName._(1, _omitEnumNames ? '' : 'VIEW_CONFIG_MAP');
  static const ViewName VIEW_CONFIG_MESSAGE = ViewName._(2, _omitEnumNames ? '' : 'VIEW_CONFIG_MESSAGE');
  static const ViewName VIEW_CONFIG_PATTERN = ViewName._(3, _omitEnumNames ? '' : 'VIEW_CONFIG_PATTERN');
  static const ViewName VIEW_CONFIG_PRINT_QUEUE = ViewName._(4, _omitEnumNames ? '' : 'VIEW_CONFIG_PRINT_QUEUE');
  static const ViewName VIEW_CONFIG_REPORT = ViewName._(5, _omitEnumNames ? '' : 'VIEW_CONFIG_REPORT');
  static const ViewName VIEW_CONFIG_SHORTCUT = ViewName._(6, _omitEnumNames ? '' : 'VIEW_CONFIG_SHORTCUT');
  static const ViewName VIEW_CURRENCY_MAP = ViewName._(7, _omitEnumNames ? '' : 'VIEW_CURRENCY_MAP');
  static const ViewName VIEW_CURRENCY_TAGS = ViewName._(8, _omitEnumNames ? '' : 'VIEW_CURRENCY_TAGS');
  static const ViewName VIEW_CURRENCY_VIEW = ViewName._(9, _omitEnumNames ? '' : 'VIEW_CURRENCY_VIEW');
  static const ViewName VIEW_CUSTOMER_ADDRESSES = ViewName._(10, _omitEnumNames ? '' : 'VIEW_CUSTOMER_ADDRESSES');
  static const ViewName VIEW_CUSTOMER_CONTACTS = ViewName._(11, _omitEnumNames ? '' : 'VIEW_CUSTOMER_CONTACTS');
  static const ViewName VIEW_CUSTOMER_EVENTS = ViewName._(12, _omitEnumNames ? '' : 'VIEW_CUSTOMER_EVENTS');
  static const ViewName VIEW_CUSTOMER_MAP = ViewName._(13, _omitEnumNames ? '' : 'VIEW_CUSTOMER_MAP');
  static const ViewName VIEW_CUSTOMER_TAGS = ViewName._(14, _omitEnumNames ? '' : 'VIEW_CUSTOMER_TAGS');
  static const ViewName VIEW_CUSTOMER_VIEW = ViewName._(15, _omitEnumNames ? '' : 'VIEW_CUSTOMER_VIEW');
  static const ViewName VIEW_EMPLOYEE_EVENTS = ViewName._(16, _omitEnumNames ? '' : 'VIEW_EMPLOYEE_EVENTS');
  static const ViewName VIEW_EMPLOYEE_MAP = ViewName._(17, _omitEnumNames ? '' : 'VIEW_EMPLOYEE_MAP');
  static const ViewName VIEW_EMPLOYEE_TAGS = ViewName._(18, _omitEnumNames ? '' : 'VIEW_EMPLOYEE_TAGS');
  static const ViewName VIEW_EMPLOYEE_VIEW = ViewName._(19, _omitEnumNames ? '' : 'VIEW_EMPLOYEE_VIEW');
  static const ViewName VIEW_ITEM_MAP = ViewName._(20, _omitEnumNames ? '' : 'VIEW_ITEM_MAP');
  static const ViewName VIEW_ITEM_TAGS = ViewName._(21, _omitEnumNames ? '' : 'VIEW_ITEM_TAGS');
  static const ViewName VIEW_ITEM_VIEW = ViewName._(22, _omitEnumNames ? '' : 'VIEW_ITEM_VIEW');
  static const ViewName VIEW_LINK_MAP = ViewName._(23, _omitEnumNames ? '' : 'VIEW_LINK_MAP');
  static const ViewName VIEW_LINK_TAGS = ViewName._(24, _omitEnumNames ? '' : 'VIEW_LINK_TAGS');
  static const ViewName VIEW_LINK_VIEW = ViewName._(25, _omitEnumNames ? '' : 'VIEW_LINK_VIEW');
  static const ViewName VIEW_MOVEMENT_MAP = ViewName._(26, _omitEnumNames ? '' : 'VIEW_MOVEMENT_MAP');
  static const ViewName VIEW_MOVEMENT_TAGS = ViewName._(27, _omitEnumNames ? '' : 'VIEW_MOVEMENT_TAGS');
  static const ViewName VIEW_MOVEMENT_STOCK = ViewName._(28, _omitEnumNames ? '' : 'VIEW_MOVEMENT_STOCK');
  static const ViewName VIEW_MOVEMENT_INVENTORY = ViewName._(29, _omitEnumNames ? '' : 'VIEW_MOVEMENT_INVENTORY');
  static const ViewName VIEW_MOVEMENT_WAYBILL = ViewName._(30, _omitEnumNames ? '' : 'VIEW_MOVEMENT_WAYBILL');
  static const ViewName VIEW_MOVEMENT_FORMULA = ViewName._(31, _omitEnumNames ? '' : 'VIEW_MOVEMENT_FORMULA');
  static const ViewName VIEW_MOVEMENT_VIEW = ViewName._(32, _omitEnumNames ? '' : 'VIEW_MOVEMENT_VIEW');
  static const ViewName VIEW_PAYMENT_MAP = ViewName._(33, _omitEnumNames ? '' : 'VIEW_PAYMENT_MAP');
  static const ViewName VIEW_PAYMENT_INVOICE = ViewName._(34, _omitEnumNames ? '' : 'VIEW_PAYMENT_INVOICE');
  static const ViewName VIEW_PAYMENT_TAGS = ViewName._(35, _omitEnumNames ? '' : 'VIEW_PAYMENT_TAGS');
  static const ViewName VIEW_PAYMENT_VIEW = ViewName._(36, _omitEnumNames ? '' : 'VIEW_PAYMENT_VIEW');
  static const ViewName VIEW_PLACE_MAP = ViewName._(37, _omitEnumNames ? '' : 'VIEW_PLACE_MAP');
  static const ViewName VIEW_PLACE_TAGS = ViewName._(38, _omitEnumNames ? '' : 'VIEW_PLACE_TAGS');
  static const ViewName VIEW_PLACE_VIEW = ViewName._(39, _omitEnumNames ? '' : 'VIEW_PLACE_VIEW');
  static const ViewName VIEW_PRICE_MAP = ViewName._(40, _omitEnumNames ? '' : 'VIEW_PRICE_MAP');
  static const ViewName VIEW_PRICE_TAGS = ViewName._(41, _omitEnumNames ? '' : 'VIEW_PRICE_TAGS');
  static const ViewName VIEW_PRICE_VIEW = ViewName._(42, _omitEnumNames ? '' : 'VIEW_PRICE_VIEW');
  static const ViewName VIEW_PRODUCT_MAP = ViewName._(43, _omitEnumNames ? '' : 'VIEW_PRODUCT_MAP');
  static const ViewName VIEW_PRODUCT_TAGS = ViewName._(44, _omitEnumNames ? '' : 'VIEW_PRODUCT_TAGS');
  static const ViewName VIEW_PRODUCT_VIEW = ViewName._(45, _omitEnumNames ? '' : 'VIEW_PRODUCT_VIEW');
  static const ViewName VIEW_PRODUCT_COMPONENTS = ViewName._(46, _omitEnumNames ? '' : 'VIEW_PRODUCT_COMPONENTS');
  static const ViewName VIEW_PROJECT_MAP = ViewName._(47, _omitEnumNames ? '' : 'VIEW_PROJECT_MAP');
  static const ViewName VIEW_PROJECT_TAGS = ViewName._(48, _omitEnumNames ? '' : 'VIEW_PROJECT_TAGS');
  static const ViewName VIEW_PROJECT_VIEW = ViewName._(49, _omitEnumNames ? '' : 'VIEW_PROJECT_VIEW');
  static const ViewName VIEW_RATE_MAP = ViewName._(50, _omitEnumNames ? '' : 'VIEW_RATE_MAP');
  static const ViewName VIEW_RATE_TAGS = ViewName._(51, _omitEnumNames ? '' : 'VIEW_RATE_TAGS');
  static const ViewName VIEW_RATE_VIEW = ViewName._(52, _omitEnumNames ? '' : 'VIEW_RATE_VIEW');
  static const ViewName VIEW_TAX_MAP = ViewName._(53, _omitEnumNames ? '' : 'VIEW_TAX_MAP');
  static const ViewName VIEW_TAX_TAGS = ViewName._(54, _omitEnumNames ? '' : 'VIEW_TAX_TAGS');
  static const ViewName VIEW_TAX_VIEW = ViewName._(55, _omitEnumNames ? '' : 'VIEW_TAX_VIEW');
  static const ViewName VIEW_TOOL_MAP = ViewName._(56, _omitEnumNames ? '' : 'VIEW_TOOL_MAP');
  static const ViewName VIEW_TOOL_TAGS = ViewName._(57, _omitEnumNames ? '' : 'VIEW_TOOL_TAGS');
  static const ViewName VIEW_TOOL_VIEW = ViewName._(58, _omitEnumNames ? '' : 'VIEW_TOOL_VIEW');
  static const ViewName VIEW_TRANS_MAP = ViewName._(59, _omitEnumNames ? '' : 'VIEW_TRANS_MAP');
  static const ViewName VIEW_TRANS_TAGS = ViewName._(60, _omitEnumNames ? '' : 'VIEW_TRANS_TAGS');
  static const ViewName VIEW_TRANS_VIEW = ViewName._(61, _omitEnumNames ? '' : 'VIEW_TRANS_VIEW');

  static const $core.List<ViewName> values = <ViewName> [
    VIEW_CONFIG_DATA,
    VIEW_CONFIG_MAP,
    VIEW_CONFIG_MESSAGE,
    VIEW_CONFIG_PATTERN,
    VIEW_CONFIG_PRINT_QUEUE,
    VIEW_CONFIG_REPORT,
    VIEW_CONFIG_SHORTCUT,
    VIEW_CURRENCY_MAP,
    VIEW_CURRENCY_TAGS,
    VIEW_CURRENCY_VIEW,
    VIEW_CUSTOMER_ADDRESSES,
    VIEW_CUSTOMER_CONTACTS,
    VIEW_CUSTOMER_EVENTS,
    VIEW_CUSTOMER_MAP,
    VIEW_CUSTOMER_TAGS,
    VIEW_CUSTOMER_VIEW,
    VIEW_EMPLOYEE_EVENTS,
    VIEW_EMPLOYEE_MAP,
    VIEW_EMPLOYEE_TAGS,
    VIEW_EMPLOYEE_VIEW,
    VIEW_ITEM_MAP,
    VIEW_ITEM_TAGS,
    VIEW_ITEM_VIEW,
    VIEW_LINK_MAP,
    VIEW_LINK_TAGS,
    VIEW_LINK_VIEW,
    VIEW_MOVEMENT_MAP,
    VIEW_MOVEMENT_TAGS,
    VIEW_MOVEMENT_STOCK,
    VIEW_MOVEMENT_INVENTORY,
    VIEW_MOVEMENT_WAYBILL,
    VIEW_MOVEMENT_FORMULA,
    VIEW_MOVEMENT_VIEW,
    VIEW_PAYMENT_MAP,
    VIEW_PAYMENT_INVOICE,
    VIEW_PAYMENT_TAGS,
    VIEW_PAYMENT_VIEW,
    VIEW_PLACE_MAP,
    VIEW_PLACE_TAGS,
    VIEW_PLACE_VIEW,
    VIEW_PRICE_MAP,
    VIEW_PRICE_TAGS,
    VIEW_PRICE_VIEW,
    VIEW_PRODUCT_MAP,
    VIEW_PRODUCT_TAGS,
    VIEW_PRODUCT_VIEW,
    VIEW_PRODUCT_COMPONENTS,
    VIEW_PROJECT_MAP,
    VIEW_PROJECT_TAGS,
    VIEW_PROJECT_VIEW,
    VIEW_RATE_MAP,
    VIEW_RATE_TAGS,
    VIEW_RATE_VIEW,
    VIEW_TAX_MAP,
    VIEW_TAX_TAGS,
    VIEW_TAX_VIEW,
    VIEW_TOOL_MAP,
    VIEW_TOOL_TAGS,
    VIEW_TOOL_VIEW,
    VIEW_TRANS_MAP,
    VIEW_TRANS_TAGS,
    VIEW_TRANS_VIEW,
  ];

  static final $core.Map<$core.int, ViewName> _byValue = $pb.ProtobufEnum.initByValue(values);
  static ViewName? valueOf($core.int value) => _byValue[value];

  const ViewName._($core.int v, $core.String n) : super(v, n);
}

class BookmarkType extends $pb.ProtobufEnum {
  static const BookmarkType BOOKMARK_BROWSER = BookmarkType._(0, _omitEnumNames ? '' : 'BOOKMARK_BROWSER');
  static const BookmarkType BOOKMARK_EDITOR = BookmarkType._(1, _omitEnumNames ? '' : 'BOOKMARK_EDITOR');

  static const $core.List<BookmarkType> values = <BookmarkType> [
    BOOKMARK_BROWSER,
    BOOKMARK_EDITOR,
  ];

  static final $core.Map<$core.int, BookmarkType> _byValue = $pb.ProtobufEnum.initByValue(values);
  static BookmarkType? valueOf($core.int value) => _byValue[value];

  const BookmarkType._($core.int v, $core.String n) : super(v, n);
}


const _omitEnumNames = $core.bool.fromEnvironment('protobuf.omit_enum_names');
