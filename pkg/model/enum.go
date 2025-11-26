package model

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"

	cu "github.com/nervatura/component/pkg/util"
)

func JSONString(value []byte) string {
	return strings.Trim(string(value), "\"")
	/*
		var s string
		if err := json.Unmarshal(b, &s); err != nil {
			return err
		}
		return s
	*/
}

type TimeDate struct {
	time.Time
}

func (td *TimeDate) UnmarshalJSON(b []byte) (err error) {
	s := JSONString(b)
	if s == "null" || s == "" {
		td.Time = time.Time{}
		return
	}
	if len(s) > 10 {
		s = s[:10]
	}
	td.Time, err = time.Parse(time.DateOnly, s)
	return
}

func (td *TimeDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(td.String())
}

func (td *TimeDate) String() string {
	if td.IsZero() {
		return ""
	}
	return td.Format(time.DateOnly)
}

type TimeDateTime struct {
	time.Time
}

func (td *TimeDateTime) UnmarshalJSON(b []byte) (err error) {
	s := JSONString(b)
	if s == "null" || s == "" {
		td.Time = time.Time{}
		return
	}
	td.Time, err = cu.StringToDateTime(s)
	return
}

func (td *TimeDateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(td.Format(time.RFC3339))
}

func (td *TimeDateTime) String() string {
	if td.IsZero() {
		return ""
	}
	return td.Format(TimeLayout)
}

type UserGroup int

const (
	UserGroupUser UserGroup = iota
	UserGroupAdmin
	UserGroupGuest
)

var userGroupMap = map[string]UserGroup{
	"GROUP_USER":  UserGroupUser,
	"GROUP_ADMIN": UserGroupAdmin,
	"GROUP_GUEST": UserGroupGuest,
}

func (ug UserGroup) String() string {
	for k, v := range userGroupMap {
		if v == ug {
			return k
		}
	}
	return ""
}

func (ug UserGroup) Keys() []string {
	keys := []string{}
	for k := range userGroupMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (ug *UserGroup) UnmarshalJSON(b []byte) error {
	s := JSONString(b)

	if result, found := userGroupMap[s]; found {
		*ug = result
	} else {
		return fmt.Errorf("invalid user group")
	}
	return nil
}

func (ug UserGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(ug.String())
}

type BarcodeType int

const (
	BarcodeTypeCode39 BarcodeType = iota
	BarcodeTypeCode128
	BarcodeTypeEan8
	BarcodeTypeEan13
	BarcodeTypeQRCode
)

var barcodeTypeMap = map[string]BarcodeType{
	"BARCODE_CODE_39":  BarcodeTypeCode39,
	"BARCODE_CODE_128": BarcodeTypeCode128,
	"BARCODE_EAN_8":    BarcodeTypeEan8,
	"BARCODE_EAN_13":   BarcodeTypeEan13,
	"BARCODE_QR_CODE":  BarcodeTypeQRCode,
}

func (bt BarcodeType) String() string {
	for k, v := range barcodeTypeMap {
		if v == bt {
			return k
		}
	}
	return ""
}

func (bt *BarcodeType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	if result, found := barcodeTypeMap[s]; found {
		*bt = result
	} else {
		return fmt.Errorf("invalid barcode type")
	}
	return nil
}

func (bt BarcodeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(bt.String())
}

type BookmarkType int

const (
	BookmarkTypeBrowser BookmarkType = iota
	BookmarkTypeEditor
)

var bookmarkTypeMap = map[string]BookmarkType{
	"BOOKMARK_BROWSER": BookmarkTypeBrowser,
	"BOOKMARK_EDITOR":  BookmarkTypeEditor,
}

func (ct BookmarkType) String() string {
	for k, v := range bookmarkTypeMap {
		if v == ct {
			return k
		}
	}
	return ""
}

func (ct *BookmarkType) UnmarshalJSON(b []byte) error {
	s := JSONString(b)

	if result, found := bookmarkTypeMap[s]; found {
		*ct = result
	} else {
		return fmt.Errorf("invalid bookmark type")
	}
	return nil
}

func (ct BookmarkType) MarshalJSON() ([]byte, error) {
	return json.Marshal(ct.String())
}

type CustomerType int

const (
	CustomerTypeCompany CustomerType = iota
	CustomerTypePrivate
	CustomerTypeOther
	CustomerTypeOwn
)

var customerTypeMap = map[string]CustomerType{
	"CUSTOMER_COMPANY": CustomerTypeCompany,
	"CUSTOMER_PRIVATE": CustomerTypePrivate,
	"CUSTOMER_OTHER":   CustomerTypeOther,
	"CUSTOMER_OWN":     CustomerTypeOwn,
}

func (ct CustomerType) String() string {
	for k, v := range customerTypeMap {
		if v == ct {
			return k
		}
	}
	return ""
}

func (ct CustomerType) Keys() []string {
	keys := []string{}
	for k := range customerTypeMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (ct CustomerType) Get(value string) CustomerType {
	if result, found := customerTypeMap[value]; found {
		return result
	}
	return CustomerTypeCompany
}

func (ct *CustomerType) UnmarshalJSON(b []byte) error {
	s := JSONString(b)

	if result, found := customerTypeMap[s]; found {
		*ct = result
	} else {
		return fmt.Errorf("invalid customer type")
	}
	return nil
}

func (ct CustomerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(ct.String())
}

type FieldType int

const (
	FieldTypeString FieldType = iota
	FieldTypeBool
	FieldTypeInteger
	FieldTypeNumber
	FieldTypeDate
	FieldTypeDateTime
	FieldTypeMemo
	FieldTypeEnum
	FieldTypeURL
	FieldTypeCustomer
	FieldTypeEmployee
	FieldTypePlace
	FieldTypeProduct
	FieldTypeProject
	FieldTypeTool
	FieldTypeTransItem
	FieldTypeTransMovement
	FieldTypeTransPayment
)

var fieldTypeMap = map[string]FieldType{
	"FIELD_STRING":         FieldTypeString,
	"FIELD_BOOL":           FieldTypeBool,
	"FIELD_INTEGER":        FieldTypeInteger,
	"FIELD_NUMBER":         FieldTypeNumber,
	"FIELD_DATE":           FieldTypeDate,
	"FIELD_DATETIME":       FieldTypeDateTime,
	"FIELD_MEMO":           FieldTypeMemo,
	"FIELD_ENUM":           FieldTypeEnum,
	"FIELD_URL":            FieldTypeURL,
	"FIELD_CUSTOMER":       FieldTypeCustomer,
	"FIELD_EMPLOYEE":       FieldTypeEmployee,
	"FIELD_PLACE":          FieldTypePlace,
	"FIELD_PRODUCT":        FieldTypeProduct,
	"FIELD_PROJECT":        FieldTypeProject,
	"FIELD_TOOL":           FieldTypeTool,
	"FIELD_TRANS_ITEM":     FieldTypeTransItem,
	"FIELD_TRANS_MOVEMENT": FieldTypeTransMovement,
	"FIELD_TRANS_PAYMENT":  FieldTypeTransPayment,
}

func (ft FieldType) String() string {
	for k, v := range fieldTypeMap {
		if v == ft {
			return k
		}
	}
	return ""
}

func (ft FieldType) Keys() []string {
	keys := []string{}
	for k := range fieldTypeMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (ft *FieldType) UnmarshalJSON(b []byte) error {
	s := JSONString(b)

	if result, found := fieldTypeMap[s]; found {
		*ft = result
	} else {
		return fmt.Errorf("invalid field type")
	}
	return nil
}

func (ft FieldType) MarshalJSON() ([]byte, error) {
	return json.Marshal(ft.String())
}

type LogType int

const (
	LogTypeInsert LogType = iota
	LogTypeUpdate
	LogTypeDelete
)

var logTypeMap = map[string]LogType{
	"LOG_INSERT": LogTypeInsert,
	"LOG_UPDATE": LogTypeUpdate,
	"LOG_DELETE": LogTypeDelete,
}

func (lt LogType) String() string {
	for k, v := range logTypeMap {
		if v == lt {
			return k
		}
	}
	return ""
}

func (lt *LogType) UnmarshalJSON(b []byte) error {
	s := JSONString(b)

	if result, found := logTypeMap[s]; found {
		*lt = result
	} else {
		return fmt.Errorf("invalid log type")
	}
	return nil
}

func (lt LogType) MarshalJSON() ([]byte, error) {
	return json.Marshal(lt.String())
}

type MovementType int

const (
	MovementTypeInventory MovementType = iota
	MovementTypeTool
	MovementTypePlan
	MovementTypeHead
)

var movementTypeMap = map[string]MovementType{
	"MOVEMENT_INVENTORY": MovementTypeInventory,
	"MOVEMENT_TOOL":      MovementTypeTool,
	"MOVEMENT_PLAN":      MovementTypePlan,
	"MOVEMENT_HEAD":      MovementTypeHead,
}

func (mt MovementType) String() string {
	for k, v := range movementTypeMap {
		if v == mt {
			return k
		}
	}
	return ""
}

func (mt *MovementType) UnmarshalJSON(b []byte) error {
	s := JSONString(b)

	if result, found := movementTypeMap[s]; found {
		*mt = result
	} else {
		return fmt.Errorf("invalid movement type")
	}
	return nil
}

func (mt MovementType) MarshalJSON() ([]byte, error) {
	return json.Marshal(mt.String())
}

type TransType int

const (
	TransTypeInvoice TransType = iota
	TransTypeReceipt
	TransTypeOrder
	TransTypeOffer
	TransTypeWorksheet
	TransTypeRent
	TransTypeDelivery
	TransTypeInventory
	TransTypeWaybill
	TransTypeProduction
	TransTypeFormula
	TransTypeBank
	TransTypeCash
)

var transTypeMap = map[string]TransType{
	"TRANS_INVOICE":    TransTypeInvoice,
	"TRANS_RECEIPT":    TransTypeReceipt,
	"TRANS_ORDER":      TransTypeOrder,
	"TRANS_OFFER":      TransTypeOffer,
	"TRANS_WORKSHEET":  TransTypeWorksheet,
	"TRANS_RENT":       TransTypeRent,
	"TRANS_DELIVERY":   TransTypeDelivery,
	"TRANS_INVENTORY":  TransTypeInventory,
	"TRANS_WAYBILL":    TransTypeWaybill,
	"TRANS_PRODUCTION": TransTypeProduction,
	"TRANS_FORMULA":    TransTypeFormula,
	"TRANS_BANK":       TransTypeBank,
	"TRANS_CASH":       TransTypeCash,
}

func (tt TransType) String() string {
	for k, v := range transTypeMap {
		if v == tt {
			return k
		}
	}
	return ""
}

func (tt TransType) Keys() []string {
	keys := []string{}
	for k := range transTypeMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (tt *TransType) UnmarshalJSON(b []byte) error {
	s := JSONString(b)

	if result, found := transTypeMap[s]; found {
		*tt = result
	} else {
		return fmt.Errorf("invalid trans type")
	}
	return nil
}

func (tt TransType) MarshalJSON() ([]byte, error) {
	return json.Marshal(tt.String())
}

type PlaceType int

const (
	PlaceTypeWarehouse PlaceType = iota
	PlaceTypeBank
	PlaceTypeCash
	PlaceTypeOther
)

var placeTypeMap = map[string]PlaceType{
	"PLACE_WAREHOUSE": PlaceTypeWarehouse,
	"PLACE_BANK":      PlaceTypeBank,
	"PLACE_CASH":      PlaceTypeCash,
	"PLACE_OTHER":     PlaceTypeOther,
}

func (pt PlaceType) String() string {
	for k, v := range placeTypeMap {
		if v == pt {
			return k
		}
	}
	return ""
}

func (pt *PlaceType) UnmarshalJSON(b []byte) error {
	s := JSONString(b)

	if result, found := placeTypeMap[s]; found {
		*pt = result
	} else {
		return fmt.Errorf("invalid place type")
	}
	return nil
}

func (pt PlaceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(pt.String())
}

type ProductType int

const (
	ProductTypeItem ProductType = iota
	ProductTypeService
	ProductTypeVirtual
)

var productTypeMap = map[string]ProductType{
	"PRODUCT_ITEM":    ProductTypeItem,
	"PRODUCT_SERVICE": ProductTypeService,
	"PRODUCT_VIRTUAL": ProductTypeVirtual,
}

func (pt ProductType) String() string {
	for k, v := range productTypeMap {
		if v == pt {
			return k
		}
	}
	return ""
}

func (pt *ProductType) UnmarshalJSON(b []byte) error {
	s := JSONString(b)

	if result, found := productTypeMap[s]; found {
		*pt = result
	} else {
		return fmt.Errorf("invalid product type")
	}
	return nil
}

func (pt ProductType) MarshalJSON() ([]byte, error) {
	return json.Marshal(pt.String())
}

type PriceType int

const (
	PriceTypeCustomer PriceType = iota
	PriceTypeVendor
)

var priceTypeMap = map[string]PriceType{
	"PRICE_CUSTOMER": PriceTypeCustomer,
	"PRICE_VENDOR":   PriceTypeVendor,
}

func (pt PriceType) String() string {
	for k, v := range priceTypeMap {
		if v == pt {
			return k
		}
	}
	return ""
}

func (pt *PriceType) UnmarshalJSON(b []byte) error {
	s := JSONString(b)

	if result, found := priceTypeMap[s]; found {
		*pt = result
	} else {
		return fmt.Errorf("invalid price type")
	}
	return nil
}

func (pt PriceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(pt.String())
}

type RateType int

const (
	RateTypeRate RateType = iota
	RateTypeBuy
	RateTypeSell
	RateTypeAverage
)

var rateTypeMap = map[string]RateType{
	"RATE_RATE":    RateTypeRate,
	"RATE_BUY":     RateTypeBuy,
	"RATE_SELL":    RateTypeSell,
	"RATE_AVERAGE": RateTypeAverage,
}

func (rt RateType) String() string {
	for k, v := range rateTypeMap {
		if v == rt {
			return k
		}
	}
	return ""
}

func (rt *RateType) UnmarshalJSON(b []byte) error {
	s := JSONString(b)

	if result, found := rateTypeMap[s]; found {
		*rt = result
	} else {
		return fmt.Errorf("invalid rate type")
	}
	return nil
}

func (rt RateType) MarshalJSON() ([]byte, error) {
	return json.Marshal(rt.String())
}

type Direction int

const (
	DirectionOut Direction = iota
	DirectionIn
	DirectionTransfer
)

var directionMap = map[string]Direction{
	"DIRECTION_OUT":      DirectionOut,
	"DIRECTION_IN":       DirectionIn,
	"DIRECTION_TRANSFER": DirectionTransfer,
}

func (dr Direction) String() string {
	for k, v := range directionMap {
		if v == dr {
			return k
		}
	}
	return ""
}

func (dr *Direction) UnmarshalJSON(b []byte) error {
	s := JSONString(b)

	if result, found := directionMap[s]; found {
		*dr = result
	} else {
		return fmt.Errorf("invalid direction")
	}
	return nil
}

func (dr Direction) MarshalJSON() ([]byte, error) {
	return json.Marshal(dr.String())
}

type PaidType int

const (
	PaidTypeOnline PaidType = iota
	PaidTypeCard
	PaidTypeTransfer
	PaidTypeCash
	PaidTypeOther
)

var paidTypeMap = map[string]PaidType{
	"PAID_ONLINE":   PaidTypeOnline,
	"PAID_CARD":     PaidTypeCard,
	"PAID_TRANSFER": PaidTypeTransfer,
	"PAID_CASH":     PaidTypeCash,
	"PAID_OTHER":    PaidTypeOther,
}

func (pt PaidType) String() string {
	for k, v := range paidTypeMap {
		if v == pt {
			return k
		}
	}
	return ""
}

func (pt *PaidType) UnmarshalJSON(b []byte) error {
	s := JSONString(b)

	if result, found := paidTypeMap[s]; found {
		*pt = result
	} else {
		return fmt.Errorf("invalid paid type")
	}
	return nil
}

func (pt PaidType) MarshalJSON() ([]byte, error) {
	return json.Marshal(pt.String())
}

type TransState int

const (
	TransStateOK TransState = iota
	TransStateNew
	TransStateBack
)

var transStateMap = map[string]TransState{
	"STATE_OK":   TransStateOK,
	"STATE_NEW":  TransStateNew,
	"STATE_BACK": TransStateBack,
}

func (tt TransState) String() string {
	for k, v := range transStateMap {
		if v == tt {
			return k
		}
	}
	return ""
}

func (tt *TransState) UnmarshalJSON(b []byte) error {
	s := JSONString(b)

	if result, found := transStateMap[s]; found {
		*tt = result
	} else {
		return fmt.Errorf("invalid trans state")
	}
	return nil
}

func (tt TransState) MarshalJSON() ([]byte, error) {
	return json.Marshal(tt.String())
}

type TransStatus int

const (
	TransStatusNormal TransStatus = iota
	TransStatusCancellation
	TransStatusAmendment
)

var transStatusMap = map[string]TransStatus{
	"STATUS_NORMAL":       TransStatusNormal,
	"STATUS_CANCELLATION": TransStatusCancellation,
	"STATUS_AMENDMENT":    TransStatusAmendment,
}

func (tc TransStatus) String() string {
	for k, v := range transStatusMap {
		if v == tc {
			return k
		}
	}
	return ""
}

func (tc TransStatus) Get(value string) TransStatus {
	if result, found := transStatusMap[value]; found {
		return result
	}
	return TransStatusNormal
}

func (tc *TransStatus) UnmarshalJSON(b []byte) error {
	s := JSONString(b)

	if result, found := transStatusMap[s]; found {
		*tc = result
	} else {
		return fmt.Errorf("invalid trans status")
	}
	return nil
}

func (tc TransStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(tc.String())
}

type ConfigType int

const (
	ConfigTypeMap ConfigType = iota
	ConfigTypeShortcut
	ConfigTypeMessage
	ConfigTypePattern
	ConfigTypeReport
	ConfigTypePrintQueue
	ConfigTypeData
)

var configTypeMap = map[string]ConfigType{
	"CONFIG_MAP":         ConfigTypeMap,
	"CONFIG_SHORTCUT":    ConfigTypeShortcut,
	"CONFIG_MESSAGE":     ConfigTypeMessage,
	"CONFIG_PATTERN":     ConfigTypePattern,
	"CONFIG_REPORT":      ConfigTypeReport,
	"CONFIG_PRINT_QUEUE": ConfigTypePrintQueue,
	"CONFIG_DATA":        ConfigTypeData,
}

func (ct ConfigType) String() string {
	for k, v := range configTypeMap {
		if v == ct {
			return k
		}
	}
	return ""
}

func (ct *ConfigType) UnmarshalJSON(b []byte) error {
	s := JSONString(b)

	if result, found := configTypeMap[s]; found {
		*ct = result
	} else {
		return fmt.Errorf("invalid config type")
	}
	return nil
}

func (ct ConfigType) MarshalJSON() ([]byte, error) {
	return json.Marshal(ct.String())
}

type ShortcutMethod int

const (
	ShortcutMethodGET ShortcutMethod = iota
	ShortcutMethodPOST
)

var shortcutMethodMap = map[string]ShortcutMethod{
	"METHOD_GET":  ShortcutMethodGET,
	"METHOD_POST": ShortcutMethodPOST,
}

func (mm ShortcutMethod) String() string {
	for k, v := range shortcutMethodMap {
		if v == mm {
			return k
		}
	}
	return ""
}

func (mm ShortcutMethod) Keys() []string {
	keys := []string{}
	for k := range shortcutMethodMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (mm *ShortcutMethod) UnmarshalJSON(b []byte) error {
	s := JSONString(b)

	if result, found := shortcutMethodMap[s]; found {
		*mm = result
	} else {
		return fmt.Errorf("invalid menu method")
	}
	return nil
}

func (mm ShortcutMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(mm.String())
}

type ShortcutField int

const (
	ShortcutFieldString ShortcutField = iota
	ShortcutFieldBool
	ShortcutFieldDate
	ShortcutFieldInteger
	ShortcutFieldNumber
)

var shortcutFieldMap = map[string]ShortcutField{
	"SHORTCUT_STRING":  ShortcutFieldString,
	"SHORTCUT_BOOL":    ShortcutFieldBool,
	"SHORTCUT_DATE":    ShortcutFieldDate,
	"SHORTCUT_INTEGER": ShortcutFieldInteger,
	"SHORTCUT_NUMBER":  ShortcutFieldNumber,
}

func (ft ShortcutField) String() string {
	for k, v := range shortcutFieldMap {
		if v == ft {
			return k
		}
	}
	return ""
}

func (ft ShortcutField) Keys() []string {
	keys := []string{}
	for k := range shortcutFieldMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (ft *ShortcutField) UnmarshalJSON(b []byte) error {
	s := JSONString(b)

	if result, found := shortcutFieldMap[s]; found {
		*ft = result
	} else {
		return fmt.Errorf("invalid field type")
	}
	return nil
}

func (ft ShortcutField) MarshalJSON() ([]byte, error) {
	return json.Marshal(ft.String())
}

type MapFilter int

const (
	MapFilterAddress MapFilter = iota
	MapFilterBarcode
	MapFilterCurrency
	MapFilterContact
	MapFilterCustomer
	MapFilterEmployee
	MapFilterEvent
	MapFilterItem
	MapFilterMovement
	MapFilterPayment
	MapFilterPlace
	MapFilterPrice
	MapFilterProduct
	MapFilterProject
	MapFilterRate
	MapFilterTax
	MapFilterTool
	MapFilterUser
	MapFilterTrans
	MapFilterInvoice
	MapFilterReceipt
	MapFilterOrder
	MapFilterOffer
	MapFilterWorksheet
	MapFilterRent
	MapFilterDelivery
	MapFilterInventory
	MapFilterWaybill
	MapFilterProduction
	MapFilterFormula
	MapFilterBank
	MapFilterCash
)

var mapFilterMap = map[string]MapFilter{
	"FILTER_ADDRESS":    MapFilterAddress,
	"FILTER_BARCODE":    MapFilterBarcode,
	"FILTER_CONTACT":    MapFilterContact,
	"FILTER_CURRENCY":   MapFilterCurrency,
	"FILTER_CUSTOMER":   MapFilterCustomer,
	"FILTER_EMPLOYEE":   MapFilterEmployee,
	"FILTER_EVENT":      MapFilterEvent,
	"FILTER_ITEM":       MapFilterItem,
	"FILTER_MOVEMENT":   MapFilterMovement,
	"FILTER_PAYMENT":    MapFilterPayment,
	"FILTER_PLACE":      MapFilterPlace,
	"FILTER_PRICE":      MapFilterPrice,
	"FILTER_PRODUCT":    MapFilterProduct,
	"FILTER_PROJECT":    MapFilterProject,
	"FILTER_RATE":       MapFilterRate,
	"FILTER_TAX":        MapFilterTax,
	"FILTER_TOOL":       MapFilterTool,
	"FILTER_USER":       MapFilterUser,
	"FILTER_TRANS":      MapFilterTrans,
	"FILTER_INVOICE":    MapFilterInvoice,
	"FILTER_RECEIPT":    MapFilterReceipt,
	"FILTER_ORDER":      MapFilterOrder,
	"FILTER_OFFER":      MapFilterOffer,
	"FILTER_WORKSHEET":  MapFilterWorksheet,
	"FILTER_RENT":       MapFilterRent,
	"FILTER_DELIVERY":   MapFilterDelivery,
	"FILTER_INVENTORY":  MapFilterInventory,
	"FILTER_WAYBILL":    MapFilterWaybill,
	"FILTER_PRODUCTION": MapFilterProduction,
	"FILTER_FORMULA":    MapFilterFormula,
	"FILTER_BANK":       MapFilterBank,
	"FILTER_CASH":       MapFilterCash,
}

func (mf MapFilter) String() string {
	for k, v := range mapFilterMap {
		if v == mf {
			return k
		}
	}
	return ""
}

func (mf MapFilter) Keys() []string {
	keys := []string{}
	for k := range mapFilterMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (mf *MapFilter) UnmarshalJSON(b []byte) error {
	s := JSONString(b)
	if result, found := mapFilterMap[s]; found {
		*mf = result
	} else {
		return fmt.Errorf("invalid map filter")
	}
	return nil
}

func (mf MapFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(mf.String())
}

type FileType int

const (
	FileTypePDF FileType = iota
	FileTypeCSV
)

var fileTypeMap = map[string]FileType{
	"FILE_PDF": FileTypePDF,
	"FILE_CSV": FileTypeCSV,
}

func (ft FileType) Value(stringValue string) FileType {
	if result, found := fileTypeMap[stringValue]; found {
		return result
	}
	return FileTypePDF
}

func (ft FileType) String() string {
	for k, v := range fileTypeMap {
		if v == ft {
			return k
		}
	}
	return ""
}

func (ft *FileType) UnmarshalJSON(b []byte) error {
	s := JSONString(b)

	if result, found := fileTypeMap[s]; found {
		*ft = result
	} else {
		return fmt.Errorf("invalid file type")
	}
	return nil
}

func (ft FileType) MarshalJSON() ([]byte, error) {
	return json.Marshal(ft.String())
}

type LinkType int

const (
	LinkTypeCustomer LinkType = iota
	LinkTypeEmployee
	LinkTypeItem
	LinkTypeMovement
	LinkTypePayment
	LinkTypePlace
	LinkTypeProduct
	LinkTypeProject
	LinkTypeTool
	LinkTypeTrans
)

var linkTypeMap = map[string]LinkType{
	"LINK_CUSTOMER": LinkTypeCustomer,
	"LINK_EMPLOYEE": LinkTypeEmployee,
	"LINK_ITEM":     LinkTypeItem,
	"LINK_MOVEMENT": LinkTypeMovement,
	"LINK_PAYMENT":  LinkTypePayment,
	"LINK_PLACE":    LinkTypePlace,
	"LINK_PRODUCT":  LinkTypeProduct,
	"LINK_PROJECT":  LinkTypeProject,
	"LINK_TOOL":     LinkTypeTool,
	"LINK_TRANS":    LinkTypeTrans,
}

func (lt LinkType) String() string {
	for k, v := range linkTypeMap {
		if v == lt {
			return k
		}
	}
	return ""
}

func (lt *LinkType) UnmarshalJSON(b []byte) error {
	s := JSONString(b)

	if result, found := linkTypeMap[s]; found {
		*lt = result
	} else {
		return fmt.Errorf("invalid link type")
	}
	return nil
}

func (lt LinkType) MarshalJSON() ([]byte, error) {
	return json.Marshal(lt.String())
}

type ViewName int

const (
	ViewConfigData ViewName = iota
	ViewConfigMap
	ViewConfigMessage
	ViewConfigPattern
	ViewConfigPrintQueue
	ViewConfigReport
	ViewConfigShortcut
	ViewCurrencyMap
	ViewCurrencyTags
	ViewCurrencyView
	ViewCustomerAddresses
	ViewCustomerContacts
	ViewCustomerEvents
	ViewCustomerMap
	ViewCustomerTags
	ViewCustomerView
	ViewEmployeeEvents
	ViewEmployeeMap
	ViewEmployeeTags
	ViewEmployeeView
	ViewItemMap
	ViewItemTags
	ViewItemView
	ViewLinkMap
	ViewLinkTags
	ViewLinkView
	ViewMovementMap
	ViewMovementTags
	ViewMovementStock
	ViewMovementInventory
	ViewMovementWaybill
	ViewMovementFormula
	ViewMovementView
	ViewPaymentMap
	ViewPaymentInvoice
	ViewPaymentTags
	ViewPaymentView
	ViewPlaceContacts
	ViewPlaceMap
	ViewPlaceTags
	ViewPlaceView
	ViewPriceMap
	ViewPriceTags
	ViewPriceView
	ViewProductEvents
	ViewProductMap
	ViewProductTags
	ViewProductView
	ViewProductComponents
	ViewProjectAddresse
	ViewProjectContacts
	ViewProjectEvents
	ViewProjectMap
	ViewProjectTags
	ViewProjectView
	ViewRateMap
	ViewRateTags
	ViewRateView
	ViewTaxMap
	ViewTaxTags
	ViewTaxView
	ViewToolEvents
	ViewToolMap
	ViewToolTags
	ViewToolView
	ViewTransMap
	ViewTransTags
	ViewTransView
)

var viewNameMap = map[string]ViewName{
	"VIEW_CONFIG_DATA":        ViewConfigData,
	"VIEW_CONFIG_MAP":         ViewConfigMap,
	"VIEW_CONFIG_MESSAGE":     ViewConfigMessage,
	"VIEW_CONFIG_PATTERN":     ViewConfigPattern,
	"VIEW_CONFIG_PRINT_QUEUE": ViewConfigPrintQueue,
	"VIEW_CONFIG_REPORT":      ViewConfigReport,
	"VIEW_CONFIG_SHORTCUT":    ViewConfigShortcut,
	"VIEW_CURRENCY_MAP":       ViewCurrencyMap,
	"VIEW_CURRENCY_TAGS":      ViewCurrencyTags,
	"VIEW_CURRENCY_VIEW":      ViewCurrencyView,
	"VIEW_CUSTOMER_ADDRESSES": ViewCustomerAddresses,
	"VIEW_CUSTOMER_CONTACTS":  ViewCustomerContacts,
	"VIEW_CUSTOMER_EVENTS":    ViewCustomerEvents,
	"VIEW_CUSTOMER_MAP":       ViewCustomerMap,
	"VIEW_CUSTOMER_TAGS":      ViewCustomerTags,
	"VIEW_CUSTOMER_VIEW":      ViewCustomerView,
	"VIEW_EMPLOYEE_EVENTS":    ViewEmployeeEvents,
	"VIEW_EMPLOYEE_MAP":       ViewEmployeeMap,
	"VIEW_EMPLOYEE_TAGS":      ViewEmployeeTags,
	"VIEW_EMPLOYEE_VIEW":      ViewEmployeeView,
	"VIEW_ITEM_MAP":           ViewItemMap,
	"VIEW_ITEM_TAGS":          ViewItemTags,
	"VIEW_ITEM_VIEW":          ViewItemView,
	"VIEW_LINK_MAP":           ViewLinkMap,
	"VIEW_LINK_TAGS":          ViewLinkTags,
	"VIEW_LINK_VIEW":          ViewLinkView,
	"VIEW_MOVEMENT_MAP":       ViewMovementMap,
	"VIEW_MOVEMENT_TAGS":      ViewMovementTags,
	"VIEW_MOVEMENT_STOCK":     ViewMovementStock,
	"VIEW_MOVEMENT_INVENTORY": ViewMovementInventory,
	"VIEW_MOVEMENT_WAYBILL":   ViewMovementWaybill,
	"VIEW_MOVEMENT_FORMULA":   ViewMovementFormula,
	"VIEW_MOVEMENT_VIEW":      ViewMovementView,
	"VIEW_PAYMENT_MAP":        ViewPaymentMap,
	"VIEW_PAYMENT_INVOICE":    ViewPaymentInvoice,
	"VIEW_PAYMENT_TAGS":       ViewPaymentTags,
	"VIEW_PAYMENT_VIEW":       ViewPaymentView,
	"VIEW_PLACE_CONTACTS":     ViewPlaceContacts,
	"VIEW_PLACE_MAP":          ViewPlaceMap,
	"VIEW_PLACE_TAGS":         ViewPlaceTags,
	"VIEW_PLACE_VIEW":         ViewPlaceView,
	"VIEW_PRICE_MAP":          ViewPriceMap,
	"VIEW_PRICE_TAGS":         ViewPriceTags,
	"VIEW_PRICE_VIEW":         ViewPriceView,
	"VIEW_PRODUCT_EVENTS":     ViewProductEvents,
	"VIEW_PRODUCT_MAP":        ViewProductMap,
	"VIEW_PRODUCT_TAGS":       ViewProductTags,
	"VIEW_PRODUCT_VIEW":       ViewProductView,
	"VIEW_PRODUCT_COMPONENTS": ViewProductComponents,
	"VIEW_PROJECT_ADDRESSES":  ViewProjectAddresse,
	"VIEW_PROJECT_CONTACTS":   ViewProjectContacts,
	"VIEW_PROJECT_EVENTS":     ViewProjectEvents,
	"VIEW_PROJECT_MAP":        ViewProjectMap,
	"VIEW_PROJECT_TAGS":       ViewProjectTags,
	"VIEW_PROJECT_VIEW":       ViewProjectView,
	"VIEW_RATE_MAP":           ViewRateMap,
	"VIEW_RATE_TAGS":          ViewRateTags,
	"VIEW_RATE_VIEW":          ViewRateView,
	"VIEW_TAX_MAP":            ViewTaxMap,
	"VIEW_TAX_TAGS":           ViewTaxTags,
	"VIEW_TAX_VIEW":           ViewTaxView,
	"VIEW_TOOL_EVENTS":        ViewToolEvents,
	"VIEW_TOOL_MAP":           ViewToolMap,
	"VIEW_TOOL_TAGS":          ViewToolTags,
	"VIEW_TOOL_VIEW":          ViewToolView,
	"VIEW_TRANS_MAP":          ViewTransMap,
	"VIEW_TRANS_TAGS":         ViewTransTags,
	"VIEW_TRANS_VIEW":         ViewTransView,
}

func (lt ViewName) String() string {
	for k, v := range viewNameMap {
		if v == lt {
			return k
		}
	}
	return ""
}

func (lt *ViewName) UnmarshalJSON(b []byte) error {
	s := JSONString(b)

	if result, found := viewNameMap[s]; found {
		*lt = result
	} else {
		return fmt.Errorf("invalid view name")
	}
	return nil
}

func (lt ViewName) MarshalJSON() ([]byte, error) {
	return json.Marshal(lt.String())
}

type AuthFilter int

const (
	AuthFilterCustomer AuthFilter = iota
	AuthFilterEmployee
	AuthFilterPlace
	AuthFilterProduct
	AuthFilterProject
	AuthFilterTool
	AuthFilterTransItem
	AuthFilterTransMovement
	AuthFilterTransPayment
	AuthFilterOffice
)

var authFilterMap = map[string]AuthFilter{
	"AUTH_CUSTOMER":       AuthFilterCustomer,
	"AUTH_EMPLOYEE":       AuthFilterEmployee,
	"AUTH_PLACE":          AuthFilterPlace,
	"AUTH_PRODUCT":        AuthFilterProduct,
	"AUTH_PROJECT":        AuthFilterProject,
	"AUTH_TOOL":           AuthFilterTool,
	"AUTH_TRANS_ITEM":     AuthFilterTransItem,
	"AUTH_TRANS_MOVEMENT": AuthFilterTransMovement,
	"AUTH_TRANS_PAYMENT":  AuthFilterTransPayment,
	"AUTH_OFFICE":         AuthFilterOffice,
}

func (ft AuthFilter) String() string {
	for k, v := range authFilterMap {
		if v == ft {
			return k
		}
	}
	return ""
}

func (ft AuthFilter) Keys() []string {
	keys := []string{}
	for k := range authFilterMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (ft *AuthFilter) UnmarshalJSON(b []byte) error {
	s := JSONString(b)

	if result, found := authFilterMap[s]; found {
		*ft = result
	} else {
		return fmt.Errorf("invalid field type")
	}
	return nil
}

func (ft AuthFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(ft.String())
}

type SessionMethod int

const (
	SessionMethodAuto SessionMethod = iota
	SessionMethodMemory
	SessionMethodFile
	SessionMethodDatabase
)

var sessionMethodMap = map[string]SessionMethod{
	"SESSION_AUTO":     SessionMethodAuto,
	"SESSION_MEMORY":   SessionMethodMemory,
	"SESSION_FILE":     SessionMethodFile,
	"SESSION_DATABASE": SessionMethodDatabase,
}

func (sm SessionMethod) String() string {
	for k, v := range sessionMethodMap {
		if v == sm {
			return k
		}
	}
	return ""
}

func (sm *SessionMethod) UnmarshalJSON(b []byte) error {
	s := JSONString(b)

	if result, found := sessionMethodMap[s]; found {
		*sm = result
	} else {
		return fmt.Errorf("invalid session method")
	}
	return nil
}

func (sm SessionMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(sm.String())
}
