from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class UserGroup(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    GROUP_USER: _ClassVar[UserGroup]
    GROUP_ADMIN: _ClassVar[UserGroup]
    GROUP_GUEST: _ClassVar[UserGroup]

class CustomerType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    CUSTOMER_COMPANY: _ClassVar[CustomerType]
    CUSTOMER_PRIVATE: _ClassVar[CustomerType]
    CUSTOMER_OTHER: _ClassVar[CustomerType]
    CUSTOMER_OWN: _ClassVar[CustomerType]

class LinkType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    LINK_CUSTOMER: _ClassVar[LinkType]
    LINK_EMPLOYEE: _ClassVar[LinkType]
    LINK_ITEM: _ClassVar[LinkType]
    LINK_MOVEMENT: _ClassVar[LinkType]
    LINK_PAYMENT: _ClassVar[LinkType]
    LINK_PLACE: _ClassVar[LinkType]
    LINK_PRODUCT: _ClassVar[LinkType]
    LINK_PROJECT: _ClassVar[LinkType]
    LINK_TOOL: _ClassVar[LinkType]
    LINK_TRANS: _ClassVar[LinkType]

class LogType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    LOG_INSERT: _ClassVar[LogType]
    LOG_UPDATE: _ClassVar[LogType]
    LOG_DELETE: _ClassVar[LogType]

class MovementType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    MOVEMENT_INVENTORY: _ClassVar[MovementType]
    MOVEMENT_TOOL: _ClassVar[MovementType]
    MOVEMENT_PLAN: _ClassVar[MovementType]
    MOVEMENT_HEAD: _ClassVar[MovementType]

class PlaceType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    PLACE_WAREHOUSE: _ClassVar[PlaceType]
    PLACE_BANK: _ClassVar[PlaceType]
    PLACE_CASH: _ClassVar[PlaceType]
    PLACE_OTHER: _ClassVar[PlaceType]

class PriceType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    PRICE_CUSTOMER: _ClassVar[PriceType]
    PRICE_VENDOR: _ClassVar[PriceType]

class BarcodeType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    BARCODE_CODE_39: _ClassVar[BarcodeType]
    BARCODE_CODE_128: _ClassVar[BarcodeType]
    BARCODE_EAN_8: _ClassVar[BarcodeType]
    BARCODE_EAN_13: _ClassVar[BarcodeType]
    BARCODE_QR_CODE: _ClassVar[BarcodeType]

class ProductType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    PRODUCT_ITEM: _ClassVar[ProductType]
    PRODUCT_SERVICE: _ClassVar[ProductType]
    PRODUCT_VIRTUAL: _ClassVar[ProductType]

class RateType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    RATE_RATE: _ClassVar[RateType]
    RATE_BUY: _ClassVar[RateType]
    RATE_SELL: _ClassVar[RateType]
    RATE_AVERAGE: _ClassVar[RateType]

class PaidType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    PAID_ONLINE: _ClassVar[PaidType]
    PAID_CARD: _ClassVar[PaidType]
    PAID_TRANSFER: _ClassVar[PaidType]
    PAID_CASH: _ClassVar[PaidType]
    PAID_OTHER: _ClassVar[PaidType]

class TransStatus(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    TRANS_NORMAL: _ClassVar[TransStatus]
    TRANS_CANCELLATION: _ClassVar[TransStatus]
    TRANS_AMENDMENT: _ClassVar[TransStatus]

class TransState(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    TRANS_OK: _ClassVar[TransState]
    TRANS_NEW: _ClassVar[TransState]
    TRANS_BACK: _ClassVar[TransState]

class Direction(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    DIRECTION_OUT: _ClassVar[Direction]
    DIRECTION_IN: _ClassVar[Direction]
    DIRECTION_TRANSFER: _ClassVar[Direction]

class TransType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    TRANS_INVOICE: _ClassVar[TransType]
    TRANS_RECEIPT: _ClassVar[TransType]
    TRANS_ORDER: _ClassVar[TransType]
    TRANS_OFFER: _ClassVar[TransType]
    TRANS_WORKSHEET: _ClassVar[TransType]
    TRANS_RENT: _ClassVar[TransType]
    TRANS_DELIVERY: _ClassVar[TransType]
    TRANS_INVENTORY: _ClassVar[TransType]
    TRANS_WAYBILL: _ClassVar[TransType]
    TRANS_PRODUCTION: _ClassVar[TransType]
    TRANS_FORMULA: _ClassVar[TransType]
    TRANS_BANK: _ClassVar[TransType]
    TRANS_CASH: _ClassVar[TransType]

class ConfigType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    CONFIG_MAP: _ClassVar[ConfigType]
    CONFIG_SHORTCUT: _ClassVar[ConfigType]
    CONFIG_MESSAGE: _ClassVar[ConfigType]
    CONFIG_PATTERN: _ClassVar[ConfigType]
    CONFIG_REPORT: _ClassVar[ConfigType]
    CONFIG_PRINT_QUEUE: _ClassVar[ConfigType]
    CONFIG_DATA: _ClassVar[ConfigType]

class FieldType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    FIELD_STRING: _ClassVar[FieldType]
    FIELD_BOOL: _ClassVar[FieldType]
    FIELD_INTEGER: _ClassVar[FieldType]
    FIELD_NUMBER: _ClassVar[FieldType]
    FIELD_DATE: _ClassVar[FieldType]
    FIELD_DATETIME: _ClassVar[FieldType]
    FIELD_MEMO: _ClassVar[FieldType]
    FIELD_ENUM: _ClassVar[FieldType]
    FIELD_URL: _ClassVar[FieldType]
    FIELD_CUSTOMER: _ClassVar[FieldType]
    FIELD_EMPLOYEE: _ClassVar[FieldType]
    FIELD_PLACE: _ClassVar[FieldType]
    FIELD_PRODUCT: _ClassVar[FieldType]
    FIELD_PROJECT: _ClassVar[FieldType]
    FIELD_TOOL: _ClassVar[FieldType]
    FIELD_TRANS: _ClassVar[FieldType]
    FIELD_TRANS_ITEM: _ClassVar[FieldType]
    FIELD_TRANS_MOVEMENT: _ClassVar[FieldType]
    FIELD_TRANS_PAYMENT: _ClassVar[FieldType]

class MapFilter(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    FILTER_ADDRESS: _ClassVar[MapFilter]
    FILTER_BARCODE: _ClassVar[MapFilter]
    FILTER_CONTACT: _ClassVar[MapFilter]
    FILTER_CURRENCY: _ClassVar[MapFilter]
    FILTER_CUSTOMER: _ClassVar[MapFilter]
    FILTER_EMPLOYEE: _ClassVar[MapFilter]
    FILTER_EVENT: _ClassVar[MapFilter]
    FILTER_ITEM: _ClassVar[MapFilter]
    FILTER_MOVEMENT: _ClassVar[MapFilter]
    FILTER_PAYMENT: _ClassVar[MapFilter]
    FILTER_PLACE: _ClassVar[MapFilter]
    FILTER_PRICE: _ClassVar[MapFilter]
    FILTER_PRODUCT: _ClassVar[MapFilter]
    FILTER_PROJECT: _ClassVar[MapFilter]
    FILTER_RATE: _ClassVar[MapFilter]
    FILTER_TAX: _ClassVar[MapFilter]
    FILTER_TOOL: _ClassVar[MapFilter]
    FILTER_USER: _ClassVar[MapFilter]
    FILTER_TRANS: _ClassVar[MapFilter]
    FILTER_INVOICE: _ClassVar[MapFilter]
    FILTER_RECEIPT: _ClassVar[MapFilter]
    FILTER_ORDER: _ClassVar[MapFilter]
    FILTER_OFFER: _ClassVar[MapFilter]
    FILTER_WORKSHEET: _ClassVar[MapFilter]
    FILTER_RENT: _ClassVar[MapFilter]
    FILTER_DELIVERY: _ClassVar[MapFilter]
    FILTER_INVENTORY: _ClassVar[MapFilter]
    FILTER_WAYBILL: _ClassVar[MapFilter]
    FILTER_PRODUCTION: _ClassVar[MapFilter]
    FILTER_FORMULA: _ClassVar[MapFilter]
    FILTER_BANK: _ClassVar[MapFilter]
    FILTER_CASH: _ClassVar[MapFilter]

class ShortcutMethod(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    METHOD_GET: _ClassVar[ShortcutMethod]
    METHOD_POST: _ClassVar[ShortcutMethod]

class ShortcutField(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    SHORTCUT_STRING: _ClassVar[ShortcutField]
    SHORTCUT_BOOL: _ClassVar[ShortcutField]
    SHORTCUT_INTEGER: _ClassVar[ShortcutField]
    SHORTCUT_NUMBER: _ClassVar[ShortcutField]
    SHORTCUT_DATE: _ClassVar[ShortcutField]

class FileType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    FILE_PDF: _ClassVar[FileType]
    FILE_CSV: _ClassVar[FileType]

class Model(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    AUTH: _ClassVar[Model]
    CONFIG: _ClassVar[Model]
    CURRENCY: _ClassVar[Model]
    CUSTOMER: _ClassVar[Model]
    EMPLOYEE: _ClassVar[Model]
    ITEM: _ClassVar[Model]
    LINK: _ClassVar[Model]
    MOVEMENT: _ClassVar[Model]
    PAYMENT: _ClassVar[Model]
    PLACE: _ClassVar[Model]
    PRICE: _ClassVar[Model]
    PRODUCT: _ClassVar[Model]
    PROJECT: _ClassVar[Model]
    RATE: _ClassVar[Model]
    TAX: _ClassVar[Model]
    TOOL: _ClassVar[Model]
    TRANS: _ClassVar[Model]

class ViewName(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    VIEW_CONFIG_DATA: _ClassVar[ViewName]
    VIEW_CONFIG_MAP: _ClassVar[ViewName]
    VIEW_CONFIG_MESSAGE: _ClassVar[ViewName]
    VIEW_CONFIG_PATTERN: _ClassVar[ViewName]
    VIEW_CONFIG_PRINT_QUEUE: _ClassVar[ViewName]
    VIEW_CONFIG_REPORT: _ClassVar[ViewName]
    VIEW_CONFIG_SHORTCUT: _ClassVar[ViewName]
    VIEW_CURRENCY_MAP: _ClassVar[ViewName]
    VIEW_CURRENCY_TAGS: _ClassVar[ViewName]
    VIEW_CURRENCY_VIEW: _ClassVar[ViewName]
    VIEW_CUSTOMER_ADDRESSES: _ClassVar[ViewName]
    VIEW_CUSTOMER_CONTACTS: _ClassVar[ViewName]
    VIEW_CUSTOMER_EVENTS: _ClassVar[ViewName]
    VIEW_CUSTOMER_MAP: _ClassVar[ViewName]
    VIEW_CUSTOMER_TAGS: _ClassVar[ViewName]
    VIEW_CUSTOMER_VIEW: _ClassVar[ViewName]
    VIEW_EMPLOYEE_EVENTS: _ClassVar[ViewName]
    VIEW_EMPLOYEE_MAP: _ClassVar[ViewName]
    VIEW_EMPLOYEE_TAGS: _ClassVar[ViewName]
    VIEW_EMPLOYEE_VIEW: _ClassVar[ViewName]
    VIEW_ITEM_MAP: _ClassVar[ViewName]
    VIEW_ITEM_TAGS: _ClassVar[ViewName]
    VIEW_ITEM_VIEW: _ClassVar[ViewName]
    VIEW_LINK_MAP: _ClassVar[ViewName]
    VIEW_LINK_TAGS: _ClassVar[ViewName]
    VIEW_LINK_VIEW: _ClassVar[ViewName]
    VIEW_MOVEMENT_MAP: _ClassVar[ViewName]
    VIEW_MOVEMENT_TAGS: _ClassVar[ViewName]
    VIEW_MOVEMENT_STOCK: _ClassVar[ViewName]
    VIEW_MOVEMENT_INVENTORY: _ClassVar[ViewName]
    VIEW_MOVEMENT_WAYBILL: _ClassVar[ViewName]
    VIEW_MOVEMENT_FORMULA: _ClassVar[ViewName]
    VIEW_MOVEMENT_VIEW: _ClassVar[ViewName]
    VIEW_PAYMENT_MAP: _ClassVar[ViewName]
    VIEW_PAYMENT_INVOICE: _ClassVar[ViewName]
    VIEW_PAYMENT_TAGS: _ClassVar[ViewName]
    VIEW_PAYMENT_VIEW: _ClassVar[ViewName]
    VIEW_PLACE_MAP: _ClassVar[ViewName]
    VIEW_PLACE_TAGS: _ClassVar[ViewName]
    VIEW_PLACE_VIEW: _ClassVar[ViewName]
    VIEW_PRICE_MAP: _ClassVar[ViewName]
    VIEW_PRICE_TAGS: _ClassVar[ViewName]
    VIEW_PRICE_VIEW: _ClassVar[ViewName]
    VIEW_PRODUCT_MAP: _ClassVar[ViewName]
    VIEW_PRODUCT_TAGS: _ClassVar[ViewName]
    VIEW_PRODUCT_VIEW: _ClassVar[ViewName]
    VIEW_PRODUCT_COMPONENTS: _ClassVar[ViewName]
    VIEW_PROJECT_MAP: _ClassVar[ViewName]
    VIEW_PROJECT_TAGS: _ClassVar[ViewName]
    VIEW_PROJECT_VIEW: _ClassVar[ViewName]
    VIEW_RATE_MAP: _ClassVar[ViewName]
    VIEW_RATE_TAGS: _ClassVar[ViewName]
    VIEW_RATE_VIEW: _ClassVar[ViewName]
    VIEW_TAX_MAP: _ClassVar[ViewName]
    VIEW_TAX_TAGS: _ClassVar[ViewName]
    VIEW_TAX_VIEW: _ClassVar[ViewName]
    VIEW_TOOL_MAP: _ClassVar[ViewName]
    VIEW_TOOL_TAGS: _ClassVar[ViewName]
    VIEW_TOOL_VIEW: _ClassVar[ViewName]
    VIEW_TRANS_MAP: _ClassVar[ViewName]
    VIEW_TRANS_TAGS: _ClassVar[ViewName]
    VIEW_TRANS_VIEW: _ClassVar[ViewName]

class BookmarkType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    BOOKMARK_BROWSER: _ClassVar[BookmarkType]
    BOOKMARK_EDITOR: _ClassVar[BookmarkType]
GROUP_USER: UserGroup
GROUP_ADMIN: UserGroup
GROUP_GUEST: UserGroup
CUSTOMER_COMPANY: CustomerType
CUSTOMER_PRIVATE: CustomerType
CUSTOMER_OTHER: CustomerType
CUSTOMER_OWN: CustomerType
LINK_CUSTOMER: LinkType
LINK_EMPLOYEE: LinkType
LINK_ITEM: LinkType
LINK_MOVEMENT: LinkType
LINK_PAYMENT: LinkType
LINK_PLACE: LinkType
LINK_PRODUCT: LinkType
LINK_PROJECT: LinkType
LINK_TOOL: LinkType
LINK_TRANS: LinkType
LOG_INSERT: LogType
LOG_UPDATE: LogType
LOG_DELETE: LogType
MOVEMENT_INVENTORY: MovementType
MOVEMENT_TOOL: MovementType
MOVEMENT_PLAN: MovementType
MOVEMENT_HEAD: MovementType
PLACE_WAREHOUSE: PlaceType
PLACE_BANK: PlaceType
PLACE_CASH: PlaceType
PLACE_OTHER: PlaceType
PRICE_CUSTOMER: PriceType
PRICE_VENDOR: PriceType
BARCODE_CODE_39: BarcodeType
BARCODE_CODE_128: BarcodeType
BARCODE_EAN_8: BarcodeType
BARCODE_EAN_13: BarcodeType
BARCODE_QR_CODE: BarcodeType
PRODUCT_ITEM: ProductType
PRODUCT_SERVICE: ProductType
PRODUCT_VIRTUAL: ProductType
RATE_RATE: RateType
RATE_BUY: RateType
RATE_SELL: RateType
RATE_AVERAGE: RateType
PAID_ONLINE: PaidType
PAID_CARD: PaidType
PAID_TRANSFER: PaidType
PAID_CASH: PaidType
PAID_OTHER: PaidType
TRANS_NORMAL: TransStatus
TRANS_CANCELLATION: TransStatus
TRANS_AMENDMENT: TransStatus
TRANS_OK: TransState
TRANS_NEW: TransState
TRANS_BACK: TransState
DIRECTION_OUT: Direction
DIRECTION_IN: Direction
DIRECTION_TRANSFER: Direction
TRANS_INVOICE: TransType
TRANS_RECEIPT: TransType
TRANS_ORDER: TransType
TRANS_OFFER: TransType
TRANS_WORKSHEET: TransType
TRANS_RENT: TransType
TRANS_DELIVERY: TransType
TRANS_INVENTORY: TransType
TRANS_WAYBILL: TransType
TRANS_PRODUCTION: TransType
TRANS_FORMULA: TransType
TRANS_BANK: TransType
TRANS_CASH: TransType
CONFIG_MAP: ConfigType
CONFIG_SHORTCUT: ConfigType
CONFIG_MESSAGE: ConfigType
CONFIG_PATTERN: ConfigType
CONFIG_REPORT: ConfigType
CONFIG_PRINT_QUEUE: ConfigType
CONFIG_DATA: ConfigType
FIELD_STRING: FieldType
FIELD_BOOL: FieldType
FIELD_INTEGER: FieldType
FIELD_NUMBER: FieldType
FIELD_DATE: FieldType
FIELD_DATETIME: FieldType
FIELD_MEMO: FieldType
FIELD_ENUM: FieldType
FIELD_URL: FieldType
FIELD_CUSTOMER: FieldType
FIELD_EMPLOYEE: FieldType
FIELD_PLACE: FieldType
FIELD_PRODUCT: FieldType
FIELD_PROJECT: FieldType
FIELD_TOOL: FieldType
FIELD_TRANS: FieldType
FIELD_TRANS_ITEM: FieldType
FIELD_TRANS_MOVEMENT: FieldType
FIELD_TRANS_PAYMENT: FieldType
FILTER_ADDRESS: MapFilter
FILTER_BARCODE: MapFilter
FILTER_CONTACT: MapFilter
FILTER_CURRENCY: MapFilter
FILTER_CUSTOMER: MapFilter
FILTER_EMPLOYEE: MapFilter
FILTER_EVENT: MapFilter
FILTER_ITEM: MapFilter
FILTER_MOVEMENT: MapFilter
FILTER_PAYMENT: MapFilter
FILTER_PLACE: MapFilter
FILTER_PRICE: MapFilter
FILTER_PRODUCT: MapFilter
FILTER_PROJECT: MapFilter
FILTER_RATE: MapFilter
FILTER_TAX: MapFilter
FILTER_TOOL: MapFilter
FILTER_USER: MapFilter
FILTER_TRANS: MapFilter
FILTER_INVOICE: MapFilter
FILTER_RECEIPT: MapFilter
FILTER_ORDER: MapFilter
FILTER_OFFER: MapFilter
FILTER_WORKSHEET: MapFilter
FILTER_RENT: MapFilter
FILTER_DELIVERY: MapFilter
FILTER_INVENTORY: MapFilter
FILTER_WAYBILL: MapFilter
FILTER_PRODUCTION: MapFilter
FILTER_FORMULA: MapFilter
FILTER_BANK: MapFilter
FILTER_CASH: MapFilter
METHOD_GET: ShortcutMethod
METHOD_POST: ShortcutMethod
SHORTCUT_STRING: ShortcutField
SHORTCUT_BOOL: ShortcutField
SHORTCUT_INTEGER: ShortcutField
SHORTCUT_NUMBER: ShortcutField
SHORTCUT_DATE: ShortcutField
FILE_PDF: FileType
FILE_CSV: FileType
AUTH: Model
CONFIG: Model
CURRENCY: Model
CUSTOMER: Model
EMPLOYEE: Model
ITEM: Model
LINK: Model
MOVEMENT: Model
PAYMENT: Model
PLACE: Model
PRICE: Model
PRODUCT: Model
PROJECT: Model
RATE: Model
TAX: Model
TOOL: Model
TRANS: Model
VIEW_CONFIG_DATA: ViewName
VIEW_CONFIG_MAP: ViewName
VIEW_CONFIG_MESSAGE: ViewName
VIEW_CONFIG_PATTERN: ViewName
VIEW_CONFIG_PRINT_QUEUE: ViewName
VIEW_CONFIG_REPORT: ViewName
VIEW_CONFIG_SHORTCUT: ViewName
VIEW_CURRENCY_MAP: ViewName
VIEW_CURRENCY_TAGS: ViewName
VIEW_CURRENCY_VIEW: ViewName
VIEW_CUSTOMER_ADDRESSES: ViewName
VIEW_CUSTOMER_CONTACTS: ViewName
VIEW_CUSTOMER_EVENTS: ViewName
VIEW_CUSTOMER_MAP: ViewName
VIEW_CUSTOMER_TAGS: ViewName
VIEW_CUSTOMER_VIEW: ViewName
VIEW_EMPLOYEE_EVENTS: ViewName
VIEW_EMPLOYEE_MAP: ViewName
VIEW_EMPLOYEE_TAGS: ViewName
VIEW_EMPLOYEE_VIEW: ViewName
VIEW_ITEM_MAP: ViewName
VIEW_ITEM_TAGS: ViewName
VIEW_ITEM_VIEW: ViewName
VIEW_LINK_MAP: ViewName
VIEW_LINK_TAGS: ViewName
VIEW_LINK_VIEW: ViewName
VIEW_MOVEMENT_MAP: ViewName
VIEW_MOVEMENT_TAGS: ViewName
VIEW_MOVEMENT_STOCK: ViewName
VIEW_MOVEMENT_INVENTORY: ViewName
VIEW_MOVEMENT_WAYBILL: ViewName
VIEW_MOVEMENT_FORMULA: ViewName
VIEW_MOVEMENT_VIEW: ViewName
VIEW_PAYMENT_MAP: ViewName
VIEW_PAYMENT_INVOICE: ViewName
VIEW_PAYMENT_TAGS: ViewName
VIEW_PAYMENT_VIEW: ViewName
VIEW_PLACE_MAP: ViewName
VIEW_PLACE_TAGS: ViewName
VIEW_PLACE_VIEW: ViewName
VIEW_PRICE_MAP: ViewName
VIEW_PRICE_TAGS: ViewName
VIEW_PRICE_VIEW: ViewName
VIEW_PRODUCT_MAP: ViewName
VIEW_PRODUCT_TAGS: ViewName
VIEW_PRODUCT_VIEW: ViewName
VIEW_PRODUCT_COMPONENTS: ViewName
VIEW_PROJECT_MAP: ViewName
VIEW_PROJECT_TAGS: ViewName
VIEW_PROJECT_VIEW: ViewName
VIEW_RATE_MAP: ViewName
VIEW_RATE_TAGS: ViewName
VIEW_RATE_VIEW: ViewName
VIEW_TAX_MAP: ViewName
VIEW_TAX_TAGS: ViewName
VIEW_TAX_VIEW: ViewName
VIEW_TOOL_MAP: ViewName
VIEW_TOOL_TAGS: ViewName
VIEW_TOOL_VIEW: ViewName
VIEW_TRANS_MAP: ViewName
VIEW_TRANS_TAGS: ViewName
VIEW_TRANS_VIEW: ViewName
BOOKMARK_BROWSER: BookmarkType
BOOKMARK_EDITOR: BookmarkType

class RequestEmpty(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class ResponseStatus(_message.Message):
    __slots__ = ("success", "message")
    SUCCESS_FIELD_NUMBER: _ClassVar[int]
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    success: bool
    message: str
    def __init__(self, success: bool = ..., message: _Optional[str] = ...) -> None: ...

class RequestAuthLogin(_message.Message):
    __slots__ = ("user_name", "password", "database")
    USER_NAME_FIELD_NUMBER: _ClassVar[int]
    PASSWORD_FIELD_NUMBER: _ClassVar[int]
    DATABASE_FIELD_NUMBER: _ClassVar[int]
    user_name: str
    password: str
    database: str
    def __init__(self, user_name: _Optional[str] = ..., password: _Optional[str] = ..., database: _Optional[str] = ...) -> None: ...

class ResponseAuthLogin(_message.Message):
    __slots__ = ("token", "version")
    TOKEN_FIELD_NUMBER: _ClassVar[int]
    VERSION_FIELD_NUMBER: _ClassVar[int]
    token: str
    version: str
    def __init__(self, token: _Optional[str] = ..., version: _Optional[str] = ...) -> None: ...

class RequestPasswordChange(_message.Message):
    __slots__ = ("password", "confirm")
    PASSWORD_FIELD_NUMBER: _ClassVar[int]
    CONFIRM_FIELD_NUMBER: _ClassVar[int]
    password: str
    confirm: str
    def __init__(self, password: _Optional[str] = ..., confirm: _Optional[str] = ...) -> None: ...

class RequestTokenDecode(_message.Message):
    __slots__ = ("token",)
    TOKEN_FIELD_NUMBER: _ClassVar[int]
    token: str
    def __init__(self, token: _Optional[str] = ...) -> None: ...

class ResponseTokenDecode(_message.Message):
    __slots__ = ("code", "user_name", "database", "exp", "iss")
    CODE_FIELD_NUMBER: _ClassVar[int]
    USER_NAME_FIELD_NUMBER: _ClassVar[int]
    DATABASE_FIELD_NUMBER: _ClassVar[int]
    EXP_FIELD_NUMBER: _ClassVar[int]
    ISS_FIELD_NUMBER: _ClassVar[int]
    code: str
    user_name: str
    database: str
    exp: str
    iss: str
    def __init__(self, code: _Optional[str] = ..., user_name: _Optional[str] = ..., database: _Optional[str] = ..., exp: _Optional[str] = ..., iss: _Optional[str] = ...) -> None: ...

class RequestGet(_message.Message):
    __slots__ = ("id", "code")
    ID_FIELD_NUMBER: _ClassVar[int]
    CODE_FIELD_NUMBER: _ClassVar[int]
    id: int
    code: str
    def __init__(self, id: _Optional[int] = ..., code: _Optional[str] = ...) -> None: ...

class RequestDelete(_message.Message):
    __slots__ = ("id", "code", "model")
    ID_FIELD_NUMBER: _ClassVar[int]
    CODE_FIELD_NUMBER: _ClassVar[int]
    MODEL_FIELD_NUMBER: _ClassVar[int]
    id: int
    code: str
    model: Model
    def __init__(self, id: _Optional[int] = ..., code: _Optional[str] = ..., model: _Optional[_Union[Model, str]] = ...) -> None: ...

class RequestQueryFilter(_message.Message):
    __slots__ = ("field", "value")
    FIELD_FIELD_NUMBER: _ClassVar[int]
    VALUE_FIELD_NUMBER: _ClassVar[int]
    field: str
    value: str
    def __init__(self, field: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...

class RequestQuery(_message.Message):
    __slots__ = ("filters", "limit", "offset")
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    LIMIT_FIELD_NUMBER: _ClassVar[int]
    OFFSET_FIELD_NUMBER: _ClassVar[int]
    filters: _containers.RepeatedCompositeFieldContainer[RequestQueryFilter]
    limit: int
    offset: int
    def __init__(self, filters: _Optional[_Iterable[_Union[RequestQueryFilter, _Mapping]]] = ..., limit: _Optional[int] = ..., offset: _Optional[int] = ...) -> None: ...

class RequestFunction(_message.Message):
    __slots__ = ("function", "args")
    class ArgsEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    FUNCTION_FIELD_NUMBER: _ClassVar[int]
    ARGS_FIELD_NUMBER: _ClassVar[int]
    function: str
    args: _containers.ScalarMap[str, str]
    def __init__(self, function: _Optional[str] = ..., args: _Optional[_Mapping[str, str]] = ...) -> None: ...

class RequestView(_message.Message):
    __slots__ = ("name", "filters", "order_by", "limit", "offset")
    NAME_FIELD_NUMBER: _ClassVar[int]
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    ORDER_BY_FIELD_NUMBER: _ClassVar[int]
    LIMIT_FIELD_NUMBER: _ClassVar[int]
    OFFSET_FIELD_NUMBER: _ClassVar[int]
    name: ViewName
    filters: _containers.RepeatedCompositeFieldContainer[RequestQueryFilter]
    order_by: _containers.RepeatedScalarFieldContainer[str]
    limit: int
    offset: int
    def __init__(self, name: _Optional[_Union[ViewName, str]] = ..., filters: _Optional[_Iterable[_Union[RequestQueryFilter, _Mapping]]] = ..., order_by: _Optional[_Iterable[str]] = ..., limit: _Optional[int] = ..., offset: _Optional[int] = ...) -> None: ...

class RequestDatabase(_message.Message):
    __slots__ = ("alias", "demo")
    ALIAS_FIELD_NUMBER: _ClassVar[int]
    DEMO_FIELD_NUMBER: _ClassVar[int]
    alias: str
    demo: bool
    def __init__(self, alias: _Optional[str] = ..., demo: bool = ...) -> None: ...

class JsonString(_message.Message):
    __slots__ = ("data",)
    class DataEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    DATA_FIELD_NUMBER: _ClassVar[int]
    data: _containers.ScalarMap[str, str]
    def __init__(self, data: _Optional[_Mapping[str, str]] = ...) -> None: ...

class JsonBytes(_message.Message):
    __slots__ = ("data",)
    DATA_FIELD_NUMBER: _ClassVar[int]
    data: bytes
    def __init__(self, data: _Optional[bytes] = ...) -> None: ...

class Address(_message.Message):
    __slots__ = ("country", "state", "zip_code", "city", "street", "notes", "tags", "address_map")
    COUNTRY_FIELD_NUMBER: _ClassVar[int]
    STATE_FIELD_NUMBER: _ClassVar[int]
    ZIP_CODE_FIELD_NUMBER: _ClassVar[int]
    CITY_FIELD_NUMBER: _ClassVar[int]
    STREET_FIELD_NUMBER: _ClassVar[int]
    NOTES_FIELD_NUMBER: _ClassVar[int]
    TAGS_FIELD_NUMBER: _ClassVar[int]
    ADDRESS_MAP_FIELD_NUMBER: _ClassVar[int]
    country: str
    state: str
    zip_code: str
    city: str
    street: str
    notes: str
    tags: _containers.RepeatedScalarFieldContainer[str]
    address_map: JsonString
    def __init__(self, country: _Optional[str] = ..., state: _Optional[str] = ..., zip_code: _Optional[str] = ..., city: _Optional[str] = ..., street: _Optional[str] = ..., notes: _Optional[str] = ..., tags: _Optional[_Iterable[str]] = ..., address_map: _Optional[_Union[JsonString, _Mapping]] = ...) -> None: ...

class Contact(_message.Message):
    __slots__ = ("first_name", "surname", "status", "phone", "mobile", "email", "notes", "tags", "contact_map")
    FIRST_NAME_FIELD_NUMBER: _ClassVar[int]
    SURNAME_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    PHONE_FIELD_NUMBER: _ClassVar[int]
    MOBILE_FIELD_NUMBER: _ClassVar[int]
    EMAIL_FIELD_NUMBER: _ClassVar[int]
    NOTES_FIELD_NUMBER: _ClassVar[int]
    TAGS_FIELD_NUMBER: _ClassVar[int]
    CONTACT_MAP_FIELD_NUMBER: _ClassVar[int]
    first_name: str
    surname: str
    status: str
    phone: str
    mobile: str
    email: str
    notes: str
    tags: _containers.RepeatedScalarFieldContainer[str]
    contact_map: JsonString
    def __init__(self, first_name: _Optional[str] = ..., surname: _Optional[str] = ..., status: _Optional[str] = ..., phone: _Optional[str] = ..., mobile: _Optional[str] = ..., email: _Optional[str] = ..., notes: _Optional[str] = ..., tags: _Optional[_Iterable[str]] = ..., contact_map: _Optional[_Union[JsonString, _Mapping]] = ...) -> None: ...

class Event(_message.Message):
    __slots__ = ("uid", "subject", "start_time", "end_time", "place", "description", "tags", "event_map")
    UID_FIELD_NUMBER: _ClassVar[int]
    SUBJECT_FIELD_NUMBER: _ClassVar[int]
    START_TIME_FIELD_NUMBER: _ClassVar[int]
    END_TIME_FIELD_NUMBER: _ClassVar[int]
    PLACE_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    TAGS_FIELD_NUMBER: _ClassVar[int]
    EVENT_MAP_FIELD_NUMBER: _ClassVar[int]
    uid: str
    subject: str
    start_time: str
    end_time: str
    place: str
    description: str
    tags: _containers.RepeatedScalarFieldContainer[str]
    event_map: JsonString
    def __init__(self, uid: _Optional[str] = ..., subject: _Optional[str] = ..., start_time: _Optional[str] = ..., end_time: _Optional[str] = ..., place: _Optional[str] = ..., description: _Optional[str] = ..., tags: _Optional[_Iterable[str]] = ..., event_map: _Optional[_Union[JsonString, _Mapping]] = ...) -> None: ...

class Bookmark(_message.Message):
    __slots__ = ("bookmark_type", "label", "key", "code", "filters", "columns", "time_stamp")
    class ColumnsEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: bool
        def __init__(self, key: _Optional[str] = ..., value: bool = ...) -> None: ...
    BOOKMARK_TYPE_FIELD_NUMBER: _ClassVar[int]
    LABEL_FIELD_NUMBER: _ClassVar[int]
    KEY_FIELD_NUMBER: _ClassVar[int]
    CODE_FIELD_NUMBER: _ClassVar[int]
    FILTERS_FIELD_NUMBER: _ClassVar[int]
    COLUMNS_FIELD_NUMBER: _ClassVar[int]
    TIME_STAMP_FIELD_NUMBER: _ClassVar[int]
    bookmark_type: BookmarkType
    label: str
    key: str
    code: str
    filters: JsonString
    columns: _containers.ScalarMap[str, bool]
    time_stamp: str
    def __init__(self, bookmark_type: _Optional[_Union[BookmarkType, str]] = ..., label: _Optional[str] = ..., key: _Optional[str] = ..., code: _Optional[str] = ..., filters: _Optional[_Union[JsonString, _Mapping]] = ..., columns: _Optional[_Mapping[str, bool]] = ..., time_stamp: _Optional[str] = ...) -> None: ...

class AuthMeta(_message.Message):
    __slots__ = ("tags", "bookmarks")
    TAGS_FIELD_NUMBER: _ClassVar[int]
    BOOKMARKS_FIELD_NUMBER: _ClassVar[int]
    tags: _containers.RepeatedScalarFieldContainer[str]
    bookmarks: _containers.RepeatedCompositeFieldContainer[Bookmark]
    def __init__(self, tags: _Optional[_Iterable[str]] = ..., bookmarks: _Optional[_Iterable[_Union[Bookmark, _Mapping]]] = ...) -> None: ...

class Auth(_message.Message):
    __slots__ = ("id", "code", "user_name", "user_group", "disabled", "auth_meta", "auth_map", "time_stamp")
    ID_FIELD_NUMBER: _ClassVar[int]
    CODE_FIELD_NUMBER: _ClassVar[int]
    USER_NAME_FIELD_NUMBER: _ClassVar[int]
    USER_GROUP_FIELD_NUMBER: _ClassVar[int]
    DISABLED_FIELD_NUMBER: _ClassVar[int]
    AUTH_META_FIELD_NUMBER: _ClassVar[int]
    AUTH_MAP_FIELD_NUMBER: _ClassVar[int]
    TIME_STAMP_FIELD_NUMBER: _ClassVar[int]
    id: int
    code: str
    user_name: str
    user_group: UserGroup
    disabled: bool
    auth_meta: AuthMeta
    auth_map: JsonString
    time_stamp: str
    def __init__(self, id: _Optional[int] = ..., code: _Optional[str] = ..., user_name: _Optional[str] = ..., user_group: _Optional[_Union[UserGroup, str]] = ..., disabled: bool = ..., auth_meta: _Optional[_Union[AuthMeta, _Mapping]] = ..., auth_map: _Optional[_Union[JsonString, _Mapping]] = ..., time_stamp: _Optional[str] = ...) -> None: ...

class CurrencyMeta(_message.Message):
    __slots__ = ("description", "digit", "cash_round", "tags")
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    DIGIT_FIELD_NUMBER: _ClassVar[int]
    CASH_ROUND_FIELD_NUMBER: _ClassVar[int]
    TAGS_FIELD_NUMBER: _ClassVar[int]
    description: str
    digit: int
    cash_round: int
    tags: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, description: _Optional[str] = ..., digit: _Optional[int] = ..., cash_round: _Optional[int] = ..., tags: _Optional[_Iterable[str]] = ...) -> None: ...

class Currency(_message.Message):
    __slots__ = ("id", "code", "currency_meta", "currency_map", "time_stamp")
    ID_FIELD_NUMBER: _ClassVar[int]
    CODE_FIELD_NUMBER: _ClassVar[int]
    CURRENCY_META_FIELD_NUMBER: _ClassVar[int]
    CURRENCY_MAP_FIELD_NUMBER: _ClassVar[int]
    TIME_STAMP_FIELD_NUMBER: _ClassVar[int]
    id: int
    code: str
    currency_meta: CurrencyMeta
    currency_map: JsonString
    time_stamp: str
    def __init__(self, id: _Optional[int] = ..., code: _Optional[str] = ..., currency_meta: _Optional[_Union[CurrencyMeta, _Mapping]] = ..., currency_map: _Optional[_Union[JsonString, _Mapping]] = ..., time_stamp: _Optional[str] = ...) -> None: ...

class Currencies(_message.Message):
    __slots__ = ("data",)
    DATA_FIELD_NUMBER: _ClassVar[int]
    data: _containers.RepeatedCompositeFieldContainer[Currency]
    def __init__(self, data: _Optional[_Iterable[_Union[Currency, _Mapping]]] = ...) -> None: ...

class CustomerMeta(_message.Message):
    __slots__ = ("tax_number", "account", "tax_free", "terms", "credit_limit", "discount", "notes", "inactive", "tags")
    TAX_NUMBER_FIELD_NUMBER: _ClassVar[int]
    ACCOUNT_FIELD_NUMBER: _ClassVar[int]
    TAX_FREE_FIELD_NUMBER: _ClassVar[int]
    TERMS_FIELD_NUMBER: _ClassVar[int]
    CREDIT_LIMIT_FIELD_NUMBER: _ClassVar[int]
    DISCOUNT_FIELD_NUMBER: _ClassVar[int]
    NOTES_FIELD_NUMBER: _ClassVar[int]
    INACTIVE_FIELD_NUMBER: _ClassVar[int]
    TAGS_FIELD_NUMBER: _ClassVar[int]
    tax_number: str
    account: str
    tax_free: bool
    terms: int
    credit_limit: float
    discount: float
    notes: str
    inactive: bool
    tags: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, tax_number: _Optional[str] = ..., account: _Optional[str] = ..., tax_free: bool = ..., terms: _Optional[int] = ..., credit_limit: _Optional[float] = ..., discount: _Optional[float] = ..., notes: _Optional[str] = ..., inactive: bool = ..., tags: _Optional[_Iterable[str]] = ...) -> None: ...

class Customer(_message.Message):
    __slots__ = ("id", "code", "customer_type", "customer_name", "addresses", "contacts", "events", "customer_meta", "customer_map", "time_stamp")
    ID_FIELD_NUMBER: _ClassVar[int]
    CODE_FIELD_NUMBER: _ClassVar[int]
    CUSTOMER_TYPE_FIELD_NUMBER: _ClassVar[int]
    CUSTOMER_NAME_FIELD_NUMBER: _ClassVar[int]
    ADDRESSES_FIELD_NUMBER: _ClassVar[int]
    CONTACTS_FIELD_NUMBER: _ClassVar[int]
    EVENTS_FIELD_NUMBER: _ClassVar[int]
    CUSTOMER_META_FIELD_NUMBER: _ClassVar[int]
    CUSTOMER_MAP_FIELD_NUMBER: _ClassVar[int]
    TIME_STAMP_FIELD_NUMBER: _ClassVar[int]
    id: int
    code: str
    customer_type: CustomerType
    customer_name: str
    addresses: _containers.RepeatedCompositeFieldContainer[Address]
    contacts: _containers.RepeatedCompositeFieldContainer[Contact]
    events: _containers.RepeatedCompositeFieldContainer[Event]
    customer_meta: CustomerMeta
    customer_map: JsonString
    time_stamp: str
    def __init__(self, id: _Optional[int] = ..., code: _Optional[str] = ..., customer_type: _Optional[_Union[CustomerType, str]] = ..., customer_name: _Optional[str] = ..., addresses: _Optional[_Iterable[_Union[Address, _Mapping]]] = ..., contacts: _Optional[_Iterable[_Union[Contact, _Mapping]]] = ..., events: _Optional[_Iterable[_Union[Event, _Mapping]]] = ..., customer_meta: _Optional[_Union[CustomerMeta, _Mapping]] = ..., customer_map: _Optional[_Union[JsonString, _Mapping]] = ..., time_stamp: _Optional[str] = ...) -> None: ...

class Customers(_message.Message):
    __slots__ = ("data",)
    DATA_FIELD_NUMBER: _ClassVar[int]
    data: _containers.RepeatedCompositeFieldContainer[Customer]
    def __init__(self, data: _Optional[_Iterable[_Union[Customer, _Mapping]]] = ...) -> None: ...

class EmployeeMeta(_message.Message):
    __slots__ = ("start_date", "end_date", "inactive", "notes", "tags")
    START_DATE_FIELD_NUMBER: _ClassVar[int]
    END_DATE_FIELD_NUMBER: _ClassVar[int]
    INACTIVE_FIELD_NUMBER: _ClassVar[int]
    NOTES_FIELD_NUMBER: _ClassVar[int]
    TAGS_FIELD_NUMBER: _ClassVar[int]
    start_date: str
    end_date: str
    inactive: bool
    notes: str
    tags: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, start_date: _Optional[str] = ..., end_date: _Optional[str] = ..., inactive: bool = ..., notes: _Optional[str] = ..., tags: _Optional[_Iterable[str]] = ...) -> None: ...

class Employee(_message.Message):
    __slots__ = ("id", "code", "address", "contact", "events", "employee_meta", "employee_map", "time_stamp")
    ID_FIELD_NUMBER: _ClassVar[int]
    CODE_FIELD_NUMBER: _ClassVar[int]
    ADDRESS_FIELD_NUMBER: _ClassVar[int]
    CONTACT_FIELD_NUMBER: _ClassVar[int]
    EVENTS_FIELD_NUMBER: _ClassVar[int]
    EMPLOYEE_META_FIELD_NUMBER: _ClassVar[int]
    EMPLOYEE_MAP_FIELD_NUMBER: _ClassVar[int]
    TIME_STAMP_FIELD_NUMBER: _ClassVar[int]
    id: int
    code: str
    address: Address
    contact: Contact
    events: _containers.RepeatedCompositeFieldContainer[Event]
    employee_meta: EmployeeMeta
    employee_map: JsonString
    time_stamp: str
    def __init__(self, id: _Optional[int] = ..., code: _Optional[str] = ..., address: _Optional[_Union[Address, _Mapping]] = ..., contact: _Optional[_Union[Contact, _Mapping]] = ..., events: _Optional[_Iterable[_Union[Event, _Mapping]]] = ..., employee_meta: _Optional[_Union[EmployeeMeta, _Mapping]] = ..., employee_map: _Optional[_Union[JsonString, _Mapping]] = ..., time_stamp: _Optional[str] = ...) -> None: ...

class Employees(_message.Message):
    __slots__ = ("data",)
    DATA_FIELD_NUMBER: _ClassVar[int]
    data: _containers.RepeatedCompositeFieldContainer[Employee]
    def __init__(self, data: _Optional[_Iterable[_Union[Employee, _Mapping]]] = ...) -> None: ...

class ItemMeta(_message.Message):
    __slots__ = ("unit", "qty", "fx_price", "net_amount", "discount", "vat_amount", "amount", "description", "deposit", "own_stock", "action_price", "tags")
    UNIT_FIELD_NUMBER: _ClassVar[int]
    QTY_FIELD_NUMBER: _ClassVar[int]
    FX_PRICE_FIELD_NUMBER: _ClassVar[int]
    NET_AMOUNT_FIELD_NUMBER: _ClassVar[int]
    DISCOUNT_FIELD_NUMBER: _ClassVar[int]
    VAT_AMOUNT_FIELD_NUMBER: _ClassVar[int]
    AMOUNT_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    DEPOSIT_FIELD_NUMBER: _ClassVar[int]
    OWN_STOCK_FIELD_NUMBER: _ClassVar[int]
    ACTION_PRICE_FIELD_NUMBER: _ClassVar[int]
    TAGS_FIELD_NUMBER: _ClassVar[int]
    unit: str
    qty: float
    fx_price: float
    net_amount: float
    discount: float
    vat_amount: float
    amount: float
    description: str
    deposit: bool
    own_stock: float
    action_price: bool
    tags: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, unit: _Optional[str] = ..., qty: _Optional[float] = ..., fx_price: _Optional[float] = ..., net_amount: _Optional[float] = ..., discount: _Optional[float] = ..., vat_amount: _Optional[float] = ..., amount: _Optional[float] = ..., description: _Optional[str] = ..., deposit: bool = ..., own_stock: _Optional[float] = ..., action_price: bool = ..., tags: _Optional[_Iterable[str]] = ...) -> None: ...

class Item(_message.Message):
    __slots__ = ("id", "code", "trans_code", "product_code", "tax_code", "item_meta", "item_map", "time_stamp")
    ID_FIELD_NUMBER: _ClassVar[int]
    CODE_FIELD_NUMBER: _ClassVar[int]
    TRANS_CODE_FIELD_NUMBER: _ClassVar[int]
    PRODUCT_CODE_FIELD_NUMBER: _ClassVar[int]
    TAX_CODE_FIELD_NUMBER: _ClassVar[int]
    ITEM_META_FIELD_NUMBER: _ClassVar[int]
    ITEM_MAP_FIELD_NUMBER: _ClassVar[int]
    TIME_STAMP_FIELD_NUMBER: _ClassVar[int]
    id: int
    code: str
    trans_code: str
    product_code: str
    tax_code: str
    item_meta: ItemMeta
    item_map: JsonString
    time_stamp: str
    def __init__(self, id: _Optional[int] = ..., code: _Optional[str] = ..., trans_code: _Optional[str] = ..., product_code: _Optional[str] = ..., tax_code: _Optional[str] = ..., item_meta: _Optional[_Union[ItemMeta, _Mapping]] = ..., item_map: _Optional[_Union[JsonString, _Mapping]] = ..., time_stamp: _Optional[str] = ...) -> None: ...

class Items(_message.Message):
    __slots__ = ("data",)
    DATA_FIELD_NUMBER: _ClassVar[int]
    data: _containers.RepeatedCompositeFieldContainer[Item]
    def __init__(self, data: _Optional[_Iterable[_Union[Item, _Mapping]]] = ...) -> None: ...

class LinkMeta(_message.Message):
    __slots__ = ("qty", "rate", "tags")
    QTY_FIELD_NUMBER: _ClassVar[int]
    RATE_FIELD_NUMBER: _ClassVar[int]
    TAGS_FIELD_NUMBER: _ClassVar[int]
    qty: float
    rate: float
    tags: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, qty: _Optional[float] = ..., rate: _Optional[float] = ..., tags: _Optional[_Iterable[str]] = ...) -> None: ...

class Link(_message.Message):
    __slots__ = ("id", "code", "link_type_1", "link_code_1", "link_type_2", "link_code_2", "link_meta", "link_map", "time_stamp")
    ID_FIELD_NUMBER: _ClassVar[int]
    CODE_FIELD_NUMBER: _ClassVar[int]
    LINK_TYPE_1_FIELD_NUMBER: _ClassVar[int]
    LINK_CODE_1_FIELD_NUMBER: _ClassVar[int]
    LINK_TYPE_2_FIELD_NUMBER: _ClassVar[int]
    LINK_CODE_2_FIELD_NUMBER: _ClassVar[int]
    LINK_META_FIELD_NUMBER: _ClassVar[int]
    LINK_MAP_FIELD_NUMBER: _ClassVar[int]
    TIME_STAMP_FIELD_NUMBER: _ClassVar[int]
    id: int
    code: str
    link_type_1: LinkType
    link_code_1: str
    link_type_2: LinkType
    link_code_2: str
    link_meta: LinkMeta
    link_map: JsonString
    time_stamp: str
    def __init__(self, id: _Optional[int] = ..., code: _Optional[str] = ..., link_type_1: _Optional[_Union[LinkType, str]] = ..., link_code_1: _Optional[str] = ..., link_type_2: _Optional[_Union[LinkType, str]] = ..., link_code_2: _Optional[str] = ..., link_meta: _Optional[_Union[LinkMeta, _Mapping]] = ..., link_map: _Optional[_Union[JsonString, _Mapping]] = ..., time_stamp: _Optional[str] = ...) -> None: ...

class Links(_message.Message):
    __slots__ = ("data",)
    DATA_FIELD_NUMBER: _ClassVar[int]
    data: _containers.RepeatedCompositeFieldContainer[Link]
    def __init__(self, data: _Optional[_Iterable[_Union[Link, _Mapping]]] = ...) -> None: ...

class Log(_message.Message):
    __slots__ = ("id", "code", "log_type", "ref_type", "ref_code", "auth_code", "log_map", "time_stamp")
    ID_FIELD_NUMBER: _ClassVar[int]
    CODE_FIELD_NUMBER: _ClassVar[int]
    LOG_TYPE_FIELD_NUMBER: _ClassVar[int]
    REF_TYPE_FIELD_NUMBER: _ClassVar[int]
    REF_CODE_FIELD_NUMBER: _ClassVar[int]
    AUTH_CODE_FIELD_NUMBER: _ClassVar[int]
    LOG_MAP_FIELD_NUMBER: _ClassVar[int]
    TIME_STAMP_FIELD_NUMBER: _ClassVar[int]
    id: int
    code: str
    log_type: LogType
    ref_type: str
    ref_code: str
    auth_code: str
    log_map: JsonString
    time_stamp: str
    def __init__(self, id: _Optional[int] = ..., code: _Optional[str] = ..., log_type: _Optional[_Union[LogType, str]] = ..., ref_type: _Optional[str] = ..., ref_code: _Optional[str] = ..., auth_code: _Optional[str] = ..., log_map: _Optional[_Union[JsonString, _Mapping]] = ..., time_stamp: _Optional[str] = ...) -> None: ...

class Logs(_message.Message):
    __slots__ = ("data",)
    DATA_FIELD_NUMBER: _ClassVar[int]
    data: _containers.RepeatedCompositeFieldContainer[Log]
    def __init__(self, data: _Optional[_Iterable[_Union[Log, _Mapping]]] = ...) -> None: ...

class MovementMeta(_message.Message):
    __slots__ = ("qty", "notes", "shared", "tags")
    QTY_FIELD_NUMBER: _ClassVar[int]
    NOTES_FIELD_NUMBER: _ClassVar[int]
    SHARED_FIELD_NUMBER: _ClassVar[int]
    TAGS_FIELD_NUMBER: _ClassVar[int]
    qty: float
    notes: str
    shared: bool
    tags: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, qty: _Optional[float] = ..., notes: _Optional[str] = ..., shared: bool = ..., tags: _Optional[_Iterable[str]] = ...) -> None: ...

class Movement(_message.Message):
    __slots__ = ("id", "code", "movement_type", "shipping_time", "trans_code", "product_code", "tool_code", "place_code", "item_code", "movement_code", "movement_meta", "movement_map", "time_stamp")
    ID_FIELD_NUMBER: _ClassVar[int]
    CODE_FIELD_NUMBER: _ClassVar[int]
    MOVEMENT_TYPE_FIELD_NUMBER: _ClassVar[int]
    SHIPPING_TIME_FIELD_NUMBER: _ClassVar[int]
    TRANS_CODE_FIELD_NUMBER: _ClassVar[int]
    PRODUCT_CODE_FIELD_NUMBER: _ClassVar[int]
    TOOL_CODE_FIELD_NUMBER: _ClassVar[int]
    PLACE_CODE_FIELD_NUMBER: _ClassVar[int]
    ITEM_CODE_FIELD_NUMBER: _ClassVar[int]
    MOVEMENT_CODE_FIELD_NUMBER: _ClassVar[int]
    MOVEMENT_META_FIELD_NUMBER: _ClassVar[int]
    MOVEMENT_MAP_FIELD_NUMBER: _ClassVar[int]
    TIME_STAMP_FIELD_NUMBER: _ClassVar[int]
    id: int
    code: str
    movement_type: MovementType
    shipping_time: str
    trans_code: str
    product_code: str
    tool_code: str
    place_code: str
    item_code: str
    movement_code: str
    movement_meta: MovementMeta
    movement_map: JsonString
    time_stamp: str
    def __init__(self, id: _Optional[int] = ..., code: _Optional[str] = ..., movement_type: _Optional[_Union[MovementType, str]] = ..., shipping_time: _Optional[str] = ..., trans_code: _Optional[str] = ..., product_code: _Optional[str] = ..., tool_code: _Optional[str] = ..., place_code: _Optional[str] = ..., item_code: _Optional[str] = ..., movement_code: _Optional[str] = ..., movement_meta: _Optional[_Union[MovementMeta, _Mapping]] = ..., movement_map: _Optional[_Union[JsonString, _Mapping]] = ..., time_stamp: _Optional[str] = ...) -> None: ...

class Movements(_message.Message):
    __slots__ = ("data",)
    DATA_FIELD_NUMBER: _ClassVar[int]
    data: _containers.RepeatedCompositeFieldContainer[Movement]
    def __init__(self, data: _Optional[_Iterable[_Union[Movement, _Mapping]]] = ...) -> None: ...

class PaymentMeta(_message.Message):
    __slots__ = ("amount", "notes", "tags")
    AMOUNT_FIELD_NUMBER: _ClassVar[int]
    NOTES_FIELD_NUMBER: _ClassVar[int]
    TAGS_FIELD_NUMBER: _ClassVar[int]
    amount: float
    notes: str
    tags: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, amount: _Optional[float] = ..., notes: _Optional[str] = ..., tags: _Optional[_Iterable[str]] = ...) -> None: ...

class Payment(_message.Message):
    __slots__ = ("id", "code", "paid_date", "trans_code", "payment_meta", "payment_map", "time_stamp")
    ID_FIELD_NUMBER: _ClassVar[int]
    CODE_FIELD_NUMBER: _ClassVar[int]
    PAID_DATE_FIELD_NUMBER: _ClassVar[int]
    TRANS_CODE_FIELD_NUMBER: _ClassVar[int]
    PAYMENT_META_FIELD_NUMBER: _ClassVar[int]
    PAYMENT_MAP_FIELD_NUMBER: _ClassVar[int]
    TIME_STAMP_FIELD_NUMBER: _ClassVar[int]
    id: int
    code: str
    paid_date: str
    trans_code: str
    payment_meta: PaymentMeta
    payment_map: JsonString
    time_stamp: str
    def __init__(self, id: _Optional[int] = ..., code: _Optional[str] = ..., paid_date: _Optional[str] = ..., trans_code: _Optional[str] = ..., payment_meta: _Optional[_Union[PaymentMeta, _Mapping]] = ..., payment_map: _Optional[_Union[JsonString, _Mapping]] = ..., time_stamp: _Optional[str] = ...) -> None: ...

class Payments(_message.Message):
    __slots__ = ("data",)
    DATA_FIELD_NUMBER: _ClassVar[int]
    data: _containers.RepeatedCompositeFieldContainer[Payment]
    def __init__(self, data: _Optional[_Iterable[_Union[Payment, _Mapping]]] = ...) -> None: ...

class PlaceMeta(_message.Message):
    __slots__ = ("notes", "inactive", "tags")
    NOTES_FIELD_NUMBER: _ClassVar[int]
    INACTIVE_FIELD_NUMBER: _ClassVar[int]
    TAGS_FIELD_NUMBER: _ClassVar[int]
    notes: str
    inactive: bool
    tags: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, notes: _Optional[str] = ..., inactive: bool = ..., tags: _Optional[_Iterable[str]] = ...) -> None: ...

class Place(_message.Message):
    __slots__ = ("id", "code", "place_type", "place_name", "currency_code", "address", "contacts", "events", "place_meta", "place_map", "time_stamp")
    ID_FIELD_NUMBER: _ClassVar[int]
    CODE_FIELD_NUMBER: _ClassVar[int]
    PLACE_TYPE_FIELD_NUMBER: _ClassVar[int]
    PLACE_NAME_FIELD_NUMBER: _ClassVar[int]
    CURRENCY_CODE_FIELD_NUMBER: _ClassVar[int]
    ADDRESS_FIELD_NUMBER: _ClassVar[int]
    CONTACTS_FIELD_NUMBER: _ClassVar[int]
    EVENTS_FIELD_NUMBER: _ClassVar[int]
    PLACE_META_FIELD_NUMBER: _ClassVar[int]
    PLACE_MAP_FIELD_NUMBER: _ClassVar[int]
    TIME_STAMP_FIELD_NUMBER: _ClassVar[int]
    id: int
    code: str
    place_type: PlaceType
    place_name: str
    currency_code: str
    address: Address
    contacts: _containers.RepeatedCompositeFieldContainer[Contact]
    events: _containers.RepeatedCompositeFieldContainer[Event]
    place_meta: PlaceMeta
    place_map: JsonString
    time_stamp: str
    def __init__(self, id: _Optional[int] = ..., code: _Optional[str] = ..., place_type: _Optional[_Union[PlaceType, str]] = ..., place_name: _Optional[str] = ..., currency_code: _Optional[str] = ..., address: _Optional[_Union[Address, _Mapping]] = ..., contacts: _Optional[_Iterable[_Union[Contact, _Mapping]]] = ..., events: _Optional[_Iterable[_Union[Event, _Mapping]]] = ..., place_meta: _Optional[_Union[PlaceMeta, _Mapping]] = ..., place_map: _Optional[_Union[JsonString, _Mapping]] = ..., time_stamp: _Optional[str] = ...) -> None: ...

class Places(_message.Message):
    __slots__ = ("data",)
    DATA_FIELD_NUMBER: _ClassVar[int]
    data: _containers.RepeatedCompositeFieldContainer[Place]
    def __init__(self, data: _Optional[_Iterable[_Union[Place, _Mapping]]] = ...) -> None: ...

class PriceMeta(_message.Message):
    __slots__ = ("price_value", "tags")
    PRICE_VALUE_FIELD_NUMBER: _ClassVar[int]
    TAGS_FIELD_NUMBER: _ClassVar[int]
    price_value: float
    tags: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, price_value: _Optional[float] = ..., tags: _Optional[_Iterable[str]] = ...) -> None: ...

class Price(_message.Message):
    __slots__ = ("id", "code", "price_type", "valid_from", "valid_to", "product_code", "currency_code", "customer_code", "qty", "price_meta", "price_map", "time_stamp")
    ID_FIELD_NUMBER: _ClassVar[int]
    CODE_FIELD_NUMBER: _ClassVar[int]
    PRICE_TYPE_FIELD_NUMBER: _ClassVar[int]
    VALID_FROM_FIELD_NUMBER: _ClassVar[int]
    VALID_TO_FIELD_NUMBER: _ClassVar[int]
    PRODUCT_CODE_FIELD_NUMBER: _ClassVar[int]
    CURRENCY_CODE_FIELD_NUMBER: _ClassVar[int]
    CUSTOMER_CODE_FIELD_NUMBER: _ClassVar[int]
    QTY_FIELD_NUMBER: _ClassVar[int]
    PRICE_META_FIELD_NUMBER: _ClassVar[int]
    PRICE_MAP_FIELD_NUMBER: _ClassVar[int]
    TIME_STAMP_FIELD_NUMBER: _ClassVar[int]
    id: int
    code: str
    price_type: PriceType
    valid_from: str
    valid_to: str
    product_code: str
    currency_code: str
    customer_code: str
    qty: float
    price_meta: PriceMeta
    price_map: JsonString
    time_stamp: str
    def __init__(self, id: _Optional[int] = ..., code: _Optional[str] = ..., price_type: _Optional[_Union[PriceType, str]] = ..., valid_from: _Optional[str] = ..., valid_to: _Optional[str] = ..., product_code: _Optional[str] = ..., currency_code: _Optional[str] = ..., customer_code: _Optional[str] = ..., qty: _Optional[float] = ..., price_meta: _Optional[_Union[PriceMeta, _Mapping]] = ..., price_map: _Optional[_Union[JsonString, _Mapping]] = ..., time_stamp: _Optional[str] = ...) -> None: ...

class Prices(_message.Message):
    __slots__ = ("data",)
    DATA_FIELD_NUMBER: _ClassVar[int]
    data: _containers.RepeatedCompositeFieldContainer[Price]
    def __init__(self, data: _Optional[_Iterable[_Union[Price, _Mapping]]] = ...) -> None: ...

class ProductMeta(_message.Message):
    __slots__ = ("unit", "barcode_type", "barcode_data", "barcode_qty", "notes", "inactive", "tags")
    UNIT_FIELD_NUMBER: _ClassVar[int]
    BARCODE_TYPE_FIELD_NUMBER: _ClassVar[int]
    BARCODE_DATA_FIELD_NUMBER: _ClassVar[int]
    BARCODE_QTY_FIELD_NUMBER: _ClassVar[int]
    NOTES_FIELD_NUMBER: _ClassVar[int]
    INACTIVE_FIELD_NUMBER: _ClassVar[int]
    TAGS_FIELD_NUMBER: _ClassVar[int]
    unit: str
    barcode_type: BarcodeType
    barcode_data: str
    barcode_qty: float
    notes: str
    inactive: bool
    tags: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, unit: _Optional[str] = ..., barcode_type: _Optional[_Union[BarcodeType, str]] = ..., barcode_data: _Optional[str] = ..., barcode_qty: _Optional[float] = ..., notes: _Optional[str] = ..., inactive: bool = ..., tags: _Optional[_Iterable[str]] = ...) -> None: ...

class Product(_message.Message):
    __slots__ = ("id", "code", "product_type", "product_name", "tax_code", "events", "product_meta", "product_map", "time_stamp")
    ID_FIELD_NUMBER: _ClassVar[int]
    CODE_FIELD_NUMBER: _ClassVar[int]
    PRODUCT_TYPE_FIELD_NUMBER: _ClassVar[int]
    PRODUCT_NAME_FIELD_NUMBER: _ClassVar[int]
    TAX_CODE_FIELD_NUMBER: _ClassVar[int]
    EVENTS_FIELD_NUMBER: _ClassVar[int]
    PRODUCT_META_FIELD_NUMBER: _ClassVar[int]
    PRODUCT_MAP_FIELD_NUMBER: _ClassVar[int]
    TIME_STAMP_FIELD_NUMBER: _ClassVar[int]
    id: int
    code: str
    product_type: ProductType
    product_name: str
    tax_code: str
    events: _containers.RepeatedCompositeFieldContainer[Event]
    product_meta: ProductMeta
    product_map: JsonString
    time_stamp: str
    def __init__(self, id: _Optional[int] = ..., code: _Optional[str] = ..., product_type: _Optional[_Union[ProductType, str]] = ..., product_name: _Optional[str] = ..., tax_code: _Optional[str] = ..., events: _Optional[_Iterable[_Union[Event, _Mapping]]] = ..., product_meta: _Optional[_Union[ProductMeta, _Mapping]] = ..., product_map: _Optional[_Union[JsonString, _Mapping]] = ..., time_stamp: _Optional[str] = ...) -> None: ...

class Products(_message.Message):
    __slots__ = ("data",)
    DATA_FIELD_NUMBER: _ClassVar[int]
    data: _containers.RepeatedCompositeFieldContainer[Product]
    def __init__(self, data: _Optional[_Iterable[_Union[Product, _Mapping]]] = ...) -> None: ...

class ProjectMeta(_message.Message):
    __slots__ = ("start_date", "end_date", "notes", "inactive", "tags")
    START_DATE_FIELD_NUMBER: _ClassVar[int]
    END_DATE_FIELD_NUMBER: _ClassVar[int]
    NOTES_FIELD_NUMBER: _ClassVar[int]
    INACTIVE_FIELD_NUMBER: _ClassVar[int]
    TAGS_FIELD_NUMBER: _ClassVar[int]
    start_date: str
    end_date: str
    notes: str
    inactive: bool
    tags: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, start_date: _Optional[str] = ..., end_date: _Optional[str] = ..., notes: _Optional[str] = ..., inactive: bool = ..., tags: _Optional[_Iterable[str]] = ...) -> None: ...

class Project(_message.Message):
    __slots__ = ("id", "code", "project_name", "customer_code", "addresses", "contacts", "events", "project_meta", "project_map", "time_stamp")
    ID_FIELD_NUMBER: _ClassVar[int]
    CODE_FIELD_NUMBER: _ClassVar[int]
    PROJECT_NAME_FIELD_NUMBER: _ClassVar[int]
    CUSTOMER_CODE_FIELD_NUMBER: _ClassVar[int]
    ADDRESSES_FIELD_NUMBER: _ClassVar[int]
    CONTACTS_FIELD_NUMBER: _ClassVar[int]
    EVENTS_FIELD_NUMBER: _ClassVar[int]
    PROJECT_META_FIELD_NUMBER: _ClassVar[int]
    PROJECT_MAP_FIELD_NUMBER: _ClassVar[int]
    TIME_STAMP_FIELD_NUMBER: _ClassVar[int]
    id: int
    code: str
    project_name: str
    customer_code: str
    addresses: _containers.RepeatedCompositeFieldContainer[Address]
    contacts: _containers.RepeatedCompositeFieldContainer[Contact]
    events: _containers.RepeatedCompositeFieldContainer[Event]
    project_meta: ProjectMeta
    project_map: JsonString
    time_stamp: str
    def __init__(self, id: _Optional[int] = ..., code: _Optional[str] = ..., project_name: _Optional[str] = ..., customer_code: _Optional[str] = ..., addresses: _Optional[_Iterable[_Union[Address, _Mapping]]] = ..., contacts: _Optional[_Iterable[_Union[Contact, _Mapping]]] = ..., events: _Optional[_Iterable[_Union[Event, _Mapping]]] = ..., project_meta: _Optional[_Union[ProjectMeta, _Mapping]] = ..., project_map: _Optional[_Union[JsonString, _Mapping]] = ..., time_stamp: _Optional[str] = ...) -> None: ...

class Projects(_message.Message):
    __slots__ = ("data",)
    DATA_FIELD_NUMBER: _ClassVar[int]
    data: _containers.RepeatedCompositeFieldContainer[Project]
    def __init__(self, data: _Optional[_Iterable[_Union[Project, _Mapping]]] = ...) -> None: ...

class RateMeta(_message.Message):
    __slots__ = ("rate_value", "tags")
    RATE_VALUE_FIELD_NUMBER: _ClassVar[int]
    TAGS_FIELD_NUMBER: _ClassVar[int]
    rate_value: float
    tags: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, rate_value: _Optional[float] = ..., tags: _Optional[_Iterable[str]] = ...) -> None: ...

class Rate(_message.Message):
    __slots__ = ("id", "code", "rate_type", "rate_date", "place_code", "currency_code", "rate_meta", "rate_map", "time_stamp")
    ID_FIELD_NUMBER: _ClassVar[int]
    CODE_FIELD_NUMBER: _ClassVar[int]
    RATE_TYPE_FIELD_NUMBER: _ClassVar[int]
    RATE_DATE_FIELD_NUMBER: _ClassVar[int]
    PLACE_CODE_FIELD_NUMBER: _ClassVar[int]
    CURRENCY_CODE_FIELD_NUMBER: _ClassVar[int]
    RATE_META_FIELD_NUMBER: _ClassVar[int]
    RATE_MAP_FIELD_NUMBER: _ClassVar[int]
    TIME_STAMP_FIELD_NUMBER: _ClassVar[int]
    id: int
    code: str
    rate_type: RateType
    rate_date: str
    place_code: str
    currency_code: str
    rate_meta: RateMeta
    rate_map: JsonString
    time_stamp: str
    def __init__(self, id: _Optional[int] = ..., code: _Optional[str] = ..., rate_type: _Optional[_Union[RateType, str]] = ..., rate_date: _Optional[str] = ..., place_code: _Optional[str] = ..., currency_code: _Optional[str] = ..., rate_meta: _Optional[_Union[RateMeta, _Mapping]] = ..., rate_map: _Optional[_Union[JsonString, _Mapping]] = ..., time_stamp: _Optional[str] = ...) -> None: ...

class Rates(_message.Message):
    __slots__ = ("data",)
    DATA_FIELD_NUMBER: _ClassVar[int]
    data: _containers.RepeatedCompositeFieldContainer[Rate]
    def __init__(self, data: _Optional[_Iterable[_Union[Rate, _Mapping]]] = ...) -> None: ...

class TaxMeta(_message.Message):
    __slots__ = ("description", "rate_value", "tags")
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    RATE_VALUE_FIELD_NUMBER: _ClassVar[int]
    TAGS_FIELD_NUMBER: _ClassVar[int]
    description: str
    rate_value: float
    tags: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, description: _Optional[str] = ..., rate_value: _Optional[float] = ..., tags: _Optional[_Iterable[str]] = ...) -> None: ...

class Tax(_message.Message):
    __slots__ = ("id", "code", "tax_meta", "tax_map", "time_stamp")
    ID_FIELD_NUMBER: _ClassVar[int]
    CODE_FIELD_NUMBER: _ClassVar[int]
    TAX_META_FIELD_NUMBER: _ClassVar[int]
    TAX_MAP_FIELD_NUMBER: _ClassVar[int]
    TIME_STAMP_FIELD_NUMBER: _ClassVar[int]
    id: int
    code: str
    tax_meta: TaxMeta
    tax_map: JsonString
    time_stamp: str
    def __init__(self, id: _Optional[int] = ..., code: _Optional[str] = ..., tax_meta: _Optional[_Union[TaxMeta, _Mapping]] = ..., tax_map: _Optional[_Union[JsonString, _Mapping]] = ..., time_stamp: _Optional[str] = ...) -> None: ...

class Taxes(_message.Message):
    __slots__ = ("data",)
    DATA_FIELD_NUMBER: _ClassVar[int]
    data: _containers.RepeatedCompositeFieldContainer[Tax]
    def __init__(self, data: _Optional[_Iterable[_Union[Tax, _Mapping]]] = ...) -> None: ...

class ToolMeta(_message.Message):
    __slots__ = ("serial_number", "notes", "inactive", "tags")
    SERIAL_NUMBER_FIELD_NUMBER: _ClassVar[int]
    NOTES_FIELD_NUMBER: _ClassVar[int]
    INACTIVE_FIELD_NUMBER: _ClassVar[int]
    TAGS_FIELD_NUMBER: _ClassVar[int]
    serial_number: str
    notes: str
    inactive: bool
    tags: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, serial_number: _Optional[str] = ..., notes: _Optional[str] = ..., inactive: bool = ..., tags: _Optional[_Iterable[str]] = ...) -> None: ...

class Tool(_message.Message):
    __slots__ = ("id", "code", "description", "product_code", "events", "tool_meta", "tool_map", "time_stamp")
    ID_FIELD_NUMBER: _ClassVar[int]
    CODE_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    PRODUCT_CODE_FIELD_NUMBER: _ClassVar[int]
    EVENTS_FIELD_NUMBER: _ClassVar[int]
    TOOL_META_FIELD_NUMBER: _ClassVar[int]
    TOOL_MAP_FIELD_NUMBER: _ClassVar[int]
    TIME_STAMP_FIELD_NUMBER: _ClassVar[int]
    id: int
    code: str
    description: str
    product_code: str
    events: _containers.RepeatedCompositeFieldContainer[Event]
    tool_meta: ToolMeta
    tool_map: JsonString
    time_stamp: str
    def __init__(self, id: _Optional[int] = ..., code: _Optional[str] = ..., description: _Optional[str] = ..., product_code: _Optional[str] = ..., events: _Optional[_Iterable[_Union[Event, _Mapping]]] = ..., tool_meta: _Optional[_Union[ToolMeta, _Mapping]] = ..., tool_map: _Optional[_Union[JsonString, _Mapping]] = ..., time_stamp: _Optional[str] = ...) -> None: ...

class Tools(_message.Message):
    __slots__ = ("data",)
    DATA_FIELD_NUMBER: _ClassVar[int]
    data: _containers.RepeatedCompositeFieldContainer[Tool]
    def __init__(self, data: _Optional[_Iterable[_Union[Tool, _Mapping]]] = ...) -> None: ...

class Worksheet(_message.Message):
    __slots__ = ("distance", "repair", "total", "justification")
    DISTANCE_FIELD_NUMBER: _ClassVar[int]
    REPAIR_FIELD_NUMBER: _ClassVar[int]
    TOTAL_FIELD_NUMBER: _ClassVar[int]
    JUSTIFICATION_FIELD_NUMBER: _ClassVar[int]
    distance: float
    repair: float
    total: float
    justification: str
    def __init__(self, distance: _Optional[float] = ..., repair: _Optional[float] = ..., total: _Optional[float] = ..., justification: _Optional[str] = ...) -> None: ...

class Rent(_message.Message):
    __slots__ = ("holiday", "bad_tool", "other", "justification")
    HOLIDAY_FIELD_NUMBER: _ClassVar[int]
    BAD_TOOL_FIELD_NUMBER: _ClassVar[int]
    OTHER_FIELD_NUMBER: _ClassVar[int]
    JUSTIFICATION_FIELD_NUMBER: _ClassVar[int]
    holiday: float
    bad_tool: float
    other: float
    justification: str
    def __init__(self, holiday: _Optional[float] = ..., bad_tool: _Optional[float] = ..., other: _Optional[float] = ..., justification: _Optional[str] = ...) -> None: ...

class Invoice(_message.Message):
    __slots__ = ("company_name", "company_address", "company_tax_number", "customer_name", "customer_address", "customer_tax_number")
    COMPANY_NAME_FIELD_NUMBER: _ClassVar[int]
    COMPANY_ADDRESS_FIELD_NUMBER: _ClassVar[int]
    COMPANY_TAX_NUMBER_FIELD_NUMBER: _ClassVar[int]
    CUSTOMER_NAME_FIELD_NUMBER: _ClassVar[int]
    CUSTOMER_ADDRESS_FIELD_NUMBER: _ClassVar[int]
    CUSTOMER_TAX_NUMBER_FIELD_NUMBER: _ClassVar[int]
    company_name: str
    company_address: str
    company_tax_number: str
    customer_name: str
    customer_address: str
    customer_tax_number: str
    def __init__(self, company_name: _Optional[str] = ..., company_address: _Optional[str] = ..., company_tax_number: _Optional[str] = ..., customer_name: _Optional[str] = ..., customer_address: _Optional[str] = ..., customer_tax_number: _Optional[str] = ...) -> None: ...

class TransMeta(_message.Message):
    __slots__ = ("due_time", "ref_number", "paid_type", "tax_free", "paid", "rate", "closed", "status", "trans_state", "notes", "internal_notes", "report_notes", "worksheet", "rent", "invoice", "tags")
    DUE_TIME_FIELD_NUMBER: _ClassVar[int]
    REF_NUMBER_FIELD_NUMBER: _ClassVar[int]
    PAID_TYPE_FIELD_NUMBER: _ClassVar[int]
    TAX_FREE_FIELD_NUMBER: _ClassVar[int]
    PAID_FIELD_NUMBER: _ClassVar[int]
    RATE_FIELD_NUMBER: _ClassVar[int]
    CLOSED_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    TRANS_STATE_FIELD_NUMBER: _ClassVar[int]
    NOTES_FIELD_NUMBER: _ClassVar[int]
    INTERNAL_NOTES_FIELD_NUMBER: _ClassVar[int]
    REPORT_NOTES_FIELD_NUMBER: _ClassVar[int]
    WORKSHEET_FIELD_NUMBER: _ClassVar[int]
    RENT_FIELD_NUMBER: _ClassVar[int]
    INVOICE_FIELD_NUMBER: _ClassVar[int]
    TAGS_FIELD_NUMBER: _ClassVar[int]
    due_time: str
    ref_number: str
    paid_type: PaidType
    tax_free: bool
    paid: bool
    rate: float
    closed: bool
    status: TransStatus
    trans_state: TransState
    notes: str
    internal_notes: str
    report_notes: str
    worksheet: Worksheet
    rent: Rent
    invoice: Invoice
    tags: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, due_time: _Optional[str] = ..., ref_number: _Optional[str] = ..., paid_type: _Optional[_Union[PaidType, str]] = ..., tax_free: bool = ..., paid: bool = ..., rate: _Optional[float] = ..., closed: bool = ..., status: _Optional[_Union[TransStatus, str]] = ..., trans_state: _Optional[_Union[TransState, str]] = ..., notes: _Optional[str] = ..., internal_notes: _Optional[str] = ..., report_notes: _Optional[str] = ..., worksheet: _Optional[_Union[Worksheet, _Mapping]] = ..., rent: _Optional[_Union[Rent, _Mapping]] = ..., invoice: _Optional[_Union[Invoice, _Mapping]] = ..., tags: _Optional[_Iterable[str]] = ...) -> None: ...

class Trans(_message.Message):
    __slots__ = ("id", "code", "trans_type", "trans_date", "direction", "trans_code", "customer_code", "employee_code", "project_code", "place_code", "currency_code", "auth_code", "trans_meta", "trans_map", "time_stamp")
    ID_FIELD_NUMBER: _ClassVar[int]
    CODE_FIELD_NUMBER: _ClassVar[int]
    TRANS_TYPE_FIELD_NUMBER: _ClassVar[int]
    TRANS_DATE_FIELD_NUMBER: _ClassVar[int]
    DIRECTION_FIELD_NUMBER: _ClassVar[int]
    TRANS_CODE_FIELD_NUMBER: _ClassVar[int]
    CUSTOMER_CODE_FIELD_NUMBER: _ClassVar[int]
    EMPLOYEE_CODE_FIELD_NUMBER: _ClassVar[int]
    PROJECT_CODE_FIELD_NUMBER: _ClassVar[int]
    PLACE_CODE_FIELD_NUMBER: _ClassVar[int]
    CURRENCY_CODE_FIELD_NUMBER: _ClassVar[int]
    AUTH_CODE_FIELD_NUMBER: _ClassVar[int]
    TRANS_META_FIELD_NUMBER: _ClassVar[int]
    TRANS_MAP_FIELD_NUMBER: _ClassVar[int]
    TIME_STAMP_FIELD_NUMBER: _ClassVar[int]
    id: int
    code: str
    trans_type: TransType
    trans_date: str
    direction: Direction
    trans_code: str
    customer_code: str
    employee_code: str
    project_code: str
    place_code: str
    currency_code: str
    auth_code: str
    trans_meta: TransMeta
    trans_map: JsonString
    time_stamp: str
    def __init__(self, id: _Optional[int] = ..., code: _Optional[str] = ..., trans_type: _Optional[_Union[TransType, str]] = ..., trans_date: _Optional[str] = ..., direction: _Optional[_Union[Direction, str]] = ..., trans_code: _Optional[str] = ..., customer_code: _Optional[str] = ..., employee_code: _Optional[str] = ..., project_code: _Optional[str] = ..., place_code: _Optional[str] = ..., currency_code: _Optional[str] = ..., auth_code: _Optional[str] = ..., trans_meta: _Optional[_Union[TransMeta, _Mapping]] = ..., trans_map: _Optional[_Union[JsonString, _Mapping]] = ..., time_stamp: _Optional[str] = ...) -> None: ...

class Transactions(_message.Message):
    __slots__ = ("data",)
    DATA_FIELD_NUMBER: _ClassVar[int]
    data: _containers.RepeatedCompositeFieldContainer[Trans]
    def __init__(self, data: _Optional[_Iterable[_Union[Trans, _Mapping]]] = ...) -> None: ...

class Config(_message.Message):
    __slots__ = ("id", "code", "config_type", "map", "shortcut", "message", "pattern", "report", "print_queue", "config_data", "time_stamp")
    ID_FIELD_NUMBER: _ClassVar[int]
    CODE_FIELD_NUMBER: _ClassVar[int]
    CONFIG_TYPE_FIELD_NUMBER: _ClassVar[int]
    MAP_FIELD_NUMBER: _ClassVar[int]
    SHORTCUT_FIELD_NUMBER: _ClassVar[int]
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    PATTERN_FIELD_NUMBER: _ClassVar[int]
    REPORT_FIELD_NUMBER: _ClassVar[int]
    PRINT_QUEUE_FIELD_NUMBER: _ClassVar[int]
    CONFIG_DATA_FIELD_NUMBER: _ClassVar[int]
    TIME_STAMP_FIELD_NUMBER: _ClassVar[int]
    id: int
    code: str
    config_type: ConfigType
    map: ConfigMap
    shortcut: ConfigShortcut
    message: ConfigMessage
    pattern: ConfigPattern
    report: ConfigReport
    print_queue: ConfigPrintQueue
    config_data: JsonString
    time_stamp: str
    def __init__(self, id: _Optional[int] = ..., code: _Optional[str] = ..., config_type: _Optional[_Union[ConfigType, str]] = ..., map: _Optional[_Union[ConfigMap, _Mapping]] = ..., shortcut: _Optional[_Union[ConfigShortcut, _Mapping]] = ..., message: _Optional[_Union[ConfigMessage, _Mapping]] = ..., pattern: _Optional[_Union[ConfigPattern, _Mapping]] = ..., report: _Optional[_Union[ConfigReport, _Mapping]] = ..., print_queue: _Optional[_Union[ConfigPrintQueue, _Mapping]] = ..., config_data: _Optional[_Union[JsonString, _Mapping]] = ..., time_stamp: _Optional[str] = ...) -> None: ...

class Configs(_message.Message):
    __slots__ = ("data",)
    DATA_FIELD_NUMBER: _ClassVar[int]
    data: _containers.RepeatedCompositeFieldContainer[Config]
    def __init__(self, data: _Optional[_Iterable[_Union[Config, _Mapping]]] = ...) -> None: ...

class ConfigMap(_message.Message):
    __slots__ = ("field_name", "field_type", "description", "tags", "filter")
    FIELD_NAME_FIELD_NUMBER: _ClassVar[int]
    FIELD_TYPE_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    TAGS_FIELD_NUMBER: _ClassVar[int]
    FILTER_FIELD_NUMBER: _ClassVar[int]
    field_name: str
    field_type: FieldType
    description: str
    tags: _containers.RepeatedScalarFieldContainer[str]
    filter: _containers.RepeatedScalarFieldContainer[MapFilter]
    def __init__(self, field_name: _Optional[str] = ..., field_type: _Optional[_Union[FieldType, str]] = ..., description: _Optional[str] = ..., tags: _Optional[_Iterable[str]] = ..., filter: _Optional[_Iterable[_Union[MapFilter, str]]] = ...) -> None: ...

class ConfigShortcutField(_message.Message):
    __slots__ = ("field_name", "description", "field_type", "order")
    FIELD_NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    FIELD_TYPE_FIELD_NUMBER: _ClassVar[int]
    ORDER_FIELD_NUMBER: _ClassVar[int]
    field_name: str
    description: str
    field_type: ShortcutField
    order: int
    def __init__(self, field_name: _Optional[str] = ..., description: _Optional[str] = ..., field_type: _Optional[_Union[ShortcutField, str]] = ..., order: _Optional[int] = ...) -> None: ...

class ConfigShortcut(_message.Message):
    __slots__ = ("shortcut_key", "description", "modul", "icon", "method", "funcname", "address", "fields")
    SHORTCUT_KEY_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    MODUL_FIELD_NUMBER: _ClassVar[int]
    ICON_FIELD_NUMBER: _ClassVar[int]
    METHOD_FIELD_NUMBER: _ClassVar[int]
    FUNCNAME_FIELD_NUMBER: _ClassVar[int]
    ADDRESS_FIELD_NUMBER: _ClassVar[int]
    FIELDS_FIELD_NUMBER: _ClassVar[int]
    shortcut_key: str
    description: str
    modul: str
    icon: str
    method: ShortcutMethod
    funcname: str
    address: str
    fields: _containers.RepeatedCompositeFieldContainer[ConfigShortcutField]
    def __init__(self, shortcut_key: _Optional[str] = ..., description: _Optional[str] = ..., modul: _Optional[str] = ..., icon: _Optional[str] = ..., method: _Optional[_Union[ShortcutMethod, str]] = ..., funcname: _Optional[str] = ..., address: _Optional[str] = ..., fields: _Optional[_Iterable[_Union[ConfigShortcutField, _Mapping]]] = ...) -> None: ...

class ConfigMessage(_message.Message):
    __slots__ = ("section", "key", "lang", "value")
    SECTION_FIELD_NUMBER: _ClassVar[int]
    KEY_FIELD_NUMBER: _ClassVar[int]
    LANG_FIELD_NUMBER: _ClassVar[int]
    VALUE_FIELD_NUMBER: _ClassVar[int]
    section: str
    key: str
    lang: str
    value: str
    def __init__(self, section: _Optional[str] = ..., key: _Optional[str] = ..., lang: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...

class ConfigPattern(_message.Message):
    __slots__ = ("trans_type", "description", "notes", "default_pattern")
    TRANS_TYPE_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    NOTES_FIELD_NUMBER: _ClassVar[int]
    DEFAULT_PATTERN_FIELD_NUMBER: _ClassVar[int]
    trans_type: TransType
    description: str
    notes: str
    default_pattern: bool
    def __init__(self, trans_type: _Optional[_Union[TransType, str]] = ..., description: _Optional[str] = ..., notes: _Optional[str] = ..., default_pattern: bool = ...) -> None: ...

class ConfigPrintQueue(_message.Message):
    __slots__ = ("ref_type", "ref_code", "qty", "report_code", "auth_code", "time_stamp")
    REF_TYPE_FIELD_NUMBER: _ClassVar[int]
    REF_CODE_FIELD_NUMBER: _ClassVar[int]
    QTY_FIELD_NUMBER: _ClassVar[int]
    REPORT_CODE_FIELD_NUMBER: _ClassVar[int]
    AUTH_CODE_FIELD_NUMBER: _ClassVar[int]
    TIME_STAMP_FIELD_NUMBER: _ClassVar[int]
    ref_type: str
    ref_code: str
    qty: int
    report_code: str
    auth_code: str
    time_stamp: str
    def __init__(self, ref_type: _Optional[str] = ..., ref_code: _Optional[str] = ..., qty: _Optional[int] = ..., report_code: _Optional[str] = ..., auth_code: _Optional[str] = ..., time_stamp: _Optional[str] = ...) -> None: ...

class ConfigReport(_message.Message):
    __slots__ = ("report_key", "report_type", "trans_type", "direction", "file_type", "report_name", "description", "label", "template")
    REPORT_KEY_FIELD_NUMBER: _ClassVar[int]
    REPORT_TYPE_FIELD_NUMBER: _ClassVar[int]
    TRANS_TYPE_FIELD_NUMBER: _ClassVar[int]
    DIRECTION_FIELD_NUMBER: _ClassVar[int]
    FILE_TYPE_FIELD_NUMBER: _ClassVar[int]
    REPORT_NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    LABEL_FIELD_NUMBER: _ClassVar[int]
    TEMPLATE_FIELD_NUMBER: _ClassVar[int]
    report_key: str
    report_type: str
    trans_type: TransType
    direction: Direction
    file_type: FileType
    report_name: str
    description: str
    label: str
    template: str
    def __init__(self, report_key: _Optional[str] = ..., report_type: _Optional[str] = ..., trans_type: _Optional[_Union[TransType, str]] = ..., direction: _Optional[_Union[Direction, str]] = ..., file_type: _Optional[_Union[FileType, str]] = ..., report_name: _Optional[str] = ..., description: _Optional[str] = ..., label: _Optional[str] = ..., template: _Optional[str] = ...) -> None: ...
