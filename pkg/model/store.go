package model

import (
	cu "github.com/nervatura/component/pkg/util"
)

type Address struct {
	Country string `json:"country"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
	City    string `json:"city"`
	Street  string `json:"street"`
	Notes   string `json:"notes"`
	// Additional tags for the address
	Tags []string `json:"tags"`
	// Flexible key-value map for additional metadata. The value is any json type.
	AddressMap cu.IM `json:"address_map"`
}

type Contact struct {
	FirstName string `json:"first_name"`
	Surname   string `json:"surname"`
	Status    string `json:"status"`
	Phone     string `json:"phone"`
	Mobile    string `json:"mobile"`
	Email     string `json:"email"`
	Notes     string `json:"notes"`
	// Additional tags for the contact
	Tags []string `json:"tags"`
	// Flexible key-value map for additional metadata. The value is any json type.
	ContactMap cu.IM `json:"contact_map"`
}

type Event struct {
	Uid         string       `json:"uid"`
	Subject     string       `json:"subject"`
	StartTime   TimeDateTime `json:"start_time"`
	EndTime     TimeDateTime `json:"end_time"`
	Place       string       `json:"place"`
	Description string       `json:"description"`
	// Additional tags for the event
	Tags []string `json:"tags"`
	// Flexible key-value map for additional metadata. The value is any json type.
	EventMap cu.IM `json:"event_map"`
}

type Bookmark struct {
	// ENUM field. Valid values: BROWSER, EDITOR
	BookmarkType BookmarkType `json:"bookmark_type"`
	Label        string       `json:"label"`
	// Editor model or browser view name
	Key string `json:"key"`
	// Model code
	Code string `json:"code"`
	// Browser filters
	Filters any `json:"filters"`
	// Browser visible columns
	Columns map[string]bool `json:"columns"`
	// Timestamp of data creation
	TimeStamp TimeDateTime `json:"time_stamp"`
}

type AuthMeta struct {
	// The description of the user
	Description string `json:"description"`
	// Additional tags for the user
	Tags []string `json:"tags"`
	// User's UI Filters
	Filter []AuthFilter `json:"filter"`
	// User's Bookmarks
	Bookmarks []Bookmark `json:"bookmarks"`
}

type Auth struct {
	// Database primary key
	// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
	Id int64 `json:"id"`
	// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
	// Example: USR1731101982N123 ("USR" + UNIX Time stamp + "N" + current ID)
	Code string `json:"code"`
	// Friendly username, nickname or email for login.
	// It can be changed, but it is a unique identifier at the database level.
	UserName string `json:"user_name"`
	// ENUM field. Valid values: ADMIN, USER, GUEST
	UserGroup UserGroup `json:"user_group"`
	Disabled  bool      `json:"disabled"`
	AuthMeta  AuthMeta  `json:"auth_meta"`
	// Flexible key-value map for additional metadata. The value is any json type.
	AuthMap cu.IM `json:"auth_map"`
	// Timestamp of data creation
	TimeStamp TimeDateTime `json:"time_stamp"`
}

type CurrencyMeta struct {
	// The name of the currency.
	Description string `json:"description"`
	// The number of decimal places used for recording and rounding by the program.
	Digit int64 `json:"digit"`
	// Rounding value for cash. Could be used in case the smallest banknote in circulation for that certain currency is not 1.
	CashRound int64 `json:"cash_round"`
	// Additional tags for the currency
	Tags []string `json:"tags"`
}

type Currency struct {
	// Database primary key
	// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
	Id int64 `json:"id"`
	// Database independent unique external key. The ISO 4217 code of the currency. The value is always mandatory
	Code         string       `json:"code"`
	CurrencyMeta CurrencyMeta `json:"currency_meta"`
	// Flexible key-value map for additional metadata. The value is any json type.
	CurrencyMap cu.IM `json:"currency_map"`
	// Timestamp of data creation
	TimeStamp TimeDateTime `json:"time_stamp"`
}

type CustomerMeta struct {
	TaxNumber string `json:"tax_number"`
	Account   string `json:"account"`
	// Tax-free
	TaxFree bool `json:"tax_free"`
	// Payment per.
	Terms int64 `json:"terms"`
	// Customer's credit limit. Data is used by financial reports.
	CreditLimit float64 `json:"credit_limit"`
	// If new product line is added (offer, order, invoice etc.) all products will receive the discount percentage specified in this field. If the product has a separate customer price, the value specified here will not be considered by the program.
	Discount float64 `json:"discount"`
	Notes    string  `json:"notes"`
	Inactive bool    `json:"inactive"`
	// Additional tags for the customer
	Tags []string `json:"tags"`
}

type Customer struct {
	// Database primary key
	// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
	Id int64 `json:"id"`
	// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
	// Example: CUS1731101982N123 ("CUS" + UNIX Time stamp + "N" + current ID)
	Code string `json:"code"`
	// ENUM field. Valid values: OWN, COMPANY, PRIVATE, OTHER
	CustomerType CustomerType `json:"customer_type"`
	// Full name of the customer
	CustomerName string       `json:"customer_name"`
	Addresses    []Address    `json:"addresses"`
	Contacts     []Contact    `json:"contacts"`
	Events       []Event      `json:"events"`
	CustomerMeta CustomerMeta `json:"customer_meta"`
	// Flexible key-value map for additional metadata. The value is any json type.
	CustomerMap cu.IM `json:"customer_map"`
	// Timestamp of data creation
	TimeStamp TimeDateTime `json:"time_stamp"`
}

type EmployeeMeta struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Inactive  bool   `json:"inactive"`
	Notes     string `json:"notes"`
	// Additional tags for the employee
	Tags []string `json:"tags"`
}

type Employee struct {
	// Database primary key
	// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
	Id int64 `json:"id"`
	// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
	// Example: EMP1731101982N123 ("EMP" + UNIX Time stamp + "N" + current ID)
	Code         string       `json:"code"`
	Address      Address      `json:"address"`
	Contact      Contact      `json:"contact"`
	Events       []Event      `json:"events"`
	EmployeeMeta EmployeeMeta `json:"employee_meta"`
	// Flexible key-value map for additional metadata. The value is any json type.
	EmployeeMap cu.IM `json:"employee_map"`
	// Timestamp of data creation
	TimeStamp TimeDateTime `json:"time_stamp"`
}

type ItemMeta struct {
	Unit        string  `json:"unit"`
	Qty         float64 `json:"qty"`
	FxPrice     float64 `json:"fx_price"`
	NetAmount   float64 `json:"net_amount"`
	Discount    float64 `json:"discount"`
	VatAmount   float64 `json:"vat_amount"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	Deposit     bool    `json:"deposit"`
	OwnStock    float64 `json:"own_stock"`
	ActionPrice bool    `json:"action_price"`
	// Additional tags for the item
	Tags []string `json:"tags"`
}

type Item struct {
	// Database primary key
	// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
	Id int64 `json:"id"`
	// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
	// Example: ITM1731101982N123 ("ITM" + UNIX Time stamp + "N" + current ID)
	Code string `json:"code"`
	// Reference to [trans](#trans).code
	TransCode string `json:"trans_code"`
	// Reference to [product](#product).code
	ProductCode string `json:"product_code"`
	// Reference to [Tax](#tax).code
	TaxCode  string   `json:"tax_code"`
	ItemMeta ItemMeta `json:"item_meta"`
	// Flexible key-value map for additional metadata. The value is any json type.
	ItemMap cu.IM `json:"item_map"`
	// Timestamp of data creation
	TimeStamp TimeDateTime `json:"time_stamp"`
}

type LinkMeta struct {
	Qty    float64 `json:"qty"`
	Amount float64 `json:"amount"`
	Rate   float64 `json:"rate"`
	// Additional tags for the link
	Tags []string `json:"tags"`
}

type Link struct {
	// Database primary key
	// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
	Id int64 `json:"id"`
	// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
	// Example: LNK1731101982N123 ("LNK" + UNIX Time stamp + "N" + current ID)
	Code string `json:"code"`
	// ENUM field. Valid values: CUSTOMER, EMPLOYEE, ITEM, MOVEMENT, PAYMENT, PLACE, PRODUCT, PROJECT, TOOL, TRANS
	LinkType1 LinkType `json:"link_type_1"`
	// Reference to LinkType1.code
	LinkCode1 string `json:"link_code_1"`
	// ENUM field. Valid values: CUSTOMER, EMPLOYEE, ITEM, MOVEMENT, PAYMENT, PLACE, PRODUCT, PROJECT, TOOL, TRANS
	LinkType2 LinkType `json:"link_type_2"`
	// Reference to LinkType2.code
	LinkCode2 string   `json:"link_code_2"`
	LinkMeta  LinkMeta `json:"link_meta"`
	// Flexible key-value map for additional metadata. The value is any json type.
	LinkMap cu.IM `json:"link_map"`
	// Timestamp of data creation
	TimeStamp TimeDateTime `json:"time_stamp"`
}

type Log struct {
	// Database primary key
	// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
	Id int64 `json:"id"`
	// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
	// Example: LOG1731101982N123 ("LOG" + UNIX Time stamp + "N" + current ID)
	Code string `json:"code"`
	// ENUM field. Valid values: INSERT, UPDATE, DELETE
	LogType LogType `json:"log_type"`
	RefType string  `json:"ref_type"`
	RefCode string  `json:"ref_code"`
	// Reference to [Auth](#auth).code
	AuthCode string      `json:"auth_code"`
	Data     interface{} `json:"data"`
	// Timestamp of data creation
	TimeStamp TimeDateTime `json:"time_stamp"`
}

type MovementMeta struct {
	Qty    float64 `json:"qty"`
	Notes  string  `json:"notes"`
	Shared bool    `json:"shared"`
	// Additional tags for the movement
	Tags []string `json:"tags"`
}

type Movement struct {
	// Database primary key
	// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
	Id int64 `json:"id"`
	// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
	// Example: MOV1731101982N123 ("MOV" + UNIX Time stamp + "N" + current ID)
	Code string `json:"code"`
	// ENUM field. Valid values: INVENTORY, TOOL, PLAN, HEAD
	MovementType MovementType `json:"movement_type"`
	// Date-time
	ShippingTime TimeDateTime `json:"shipping_time"`
	// Reference to [trans](#trans).code
	TransCode string `json:"trans_code"`
	// Reference to [Product](#product).code
	ProductCode string `json:"product_code"`
	// Reference to [Tool](#tool).code
	ToolCode string `json:"tool_code"`
	// Reference to [Place](#place).code
	PlaceCode string `json:"place_code"`
	// Reference to [Item](#item).code
	ItemCode string `json:"item_code"`
	// Reference to [Movement](#movement).code
	MovementCode string       `json:"movement_code"`
	MovementMeta MovementMeta `json:"movement_meta"`
	// Flexible key-value map for additional metadata. The value is any json type.
	MovementMap cu.IM `json:"movement_map"`
	// Timestamp of data creation
	TimeStamp TimeDateTime `json:"time_stamp"`
}

type PaymentMeta struct {
	Amount float64 `json:"amount"`
	Notes  string  `json:"notes"`
	// Additional tags for the payment
	Tags []string `json:"tags"`
}

type Payment struct {
	// Database primary key
	// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
	Id int64 `json:"id"`
	// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
	// Example: PMT1731101982N123 ("PMT" + UNIX Time stamp + "N" + current ID)
	Code     string   `json:"code"`
	PaidDate TimeDate `json:"paid_date"`
	// Reference to [trans](#trans).code
	TransCode   string      `json:"trans_code"`
	PaymentMeta PaymentMeta `json:"payment_meta"`
	// Flexible key-value map for additional metadata. The value is any json type.
	PaymentMap cu.IM `json:"payment_map"`
	// Timestamp of data creation
	TimeStamp TimeDateTime `json:"time_stamp"`
}

type PlaceMeta struct {
	Notes    string `json:"notes"`
	Inactive bool   `json:"inactive"`
	// Additional tags for the place
	Tags []string `json:"tags"`
}

type Place struct {
	// Database primary key
	// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
	Id int64 `json:"id"`
	// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
	// Example: PLA1731101982N123 ("PLA" + UNIX Time stamp + "N" + current ID)
	Code string `json:"code"`
	// ENUM field. Valid values: BANK, CASH, WAREHOUSE, OTHER
	PlaceType PlaceType `json:"place_type"`
	// The full name of the place.
	PlaceName string `json:"place_name"`
	// Reference to [currency](#currency).code
	CurrencyCode string    `json:"currency_code"`
	Address      Address   `json:"address"`
	Contacts     []Contact `json:"contacts"`
	PlaceMeta    PlaceMeta `json:"place_meta"`
	// Flexible key-value map for additional metadata. The value is any json type.
	PlaceMap cu.IM `json:"place_map"`
	// Timestamp of data creation
	TimeStamp TimeDateTime `json:"time_stamp"`
}

type PriceMeta struct {
	// Price value
	PriceValue float64 `json:"price_value"`
	// Additional tags for the price
	Tags []string `json:"tags"`
}

type Price struct {
	// Database primary key
	// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
	Id int64 `json:"id"`
	// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
	// Example: PRC1731101982N123 ("PRC" + UNIX Time stamp + "N" + current ID)
	Code string `json:"code"`
	// ENUM field. Valid values: CUSTOMER, VENDOR
	PriceType PriceType `json:"price_type"`
	// Start of validity, mandatory data.
	ValidFrom TimeDate `json:"valid_from"`
	// End of validity, can be left empty.
	ValidTo TimeDate `json:"valid_to"`
	// Reference to [Product](#product).code
	ProductCode string `json:"product_code"`
	// Reference to [Currency](#currency).code
	CurrencyCode string `json:"currency_code"`
	// Reference to [Customer](#customer).code
	CustomerCode string `json:"customer_code"`
	// Price ranges can also be specified, thus different price can be set for a smaller and bigger quantity of the same product.
	// The quantity should be used as the lower threshold, ie. this should be the minimum quantity for the price set.
	Qty       float64   `json:"qty"`
	PriceMeta PriceMeta `json:"price_meta"`
	// Flexible key-value map for additional metadata. The value is any json type.
	PriceMap cu.IM `json:"price_map"`
	// Timestamp of data creation
	TimeStamp TimeDateTime `json:"time_stamp"`
}

type ProductMeta struct {
	Unit string `json:"unit"`
	// ENUM field. Valid values: CODE_128, CODE_39, EAN_13, EAN_8, QR_CODE
	BarcodeType BarcodeType `json:"barcode_type"`
	// Any barcode or QR code data
	Barcode string `json:"barcode"`
	// The actual amount of the products identified by the barcode. For example can be used for packaged goods, tray packaging.
	BarcodeQty float64 `json:"barcode_qty"`
	Notes      string  `json:"notes"`
	Inactive   bool    `json:"inactive"`
	// Additional tags for the product
	Tags []string `json:"tags"`
}

type Product struct {
	// Database primary key
	// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
	Id int64 `json:"id"`
	// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
	// Example: PRD1731101982N123 ("PRD" + UNIX Time stamp + "N" + current ID)
	Code string `json:"code"`
	// ENUM field. Valid values: ITEM, SERVICE
	ProductType ProductType `json:"product_type"`
	// The full name of the product or short description.
	ProductName string `json:"product_name"`
	// Reference to [Tax](#tax).code
	TaxCode     string      `json:"tax_code"`
	Events      []Event     `json:"events"`
	ProductMeta ProductMeta `json:"product_meta"`
	// Flexible key-value map for additional metadata. The value is any json type.
	ProductMap cu.IM `json:"product_map"`
	// Timestamp of data creation
	TimeStamp TimeDateTime `json:"time_stamp"`
	Deleted   bool         `json:"deleted"`
}

type ProjectMeta struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Notes     string `json:"notes"`
	Inactive  bool   `json:"inactive"`
	// Additional tags for the project
	Tags []string `json:"tags"`
}

type Project struct {
	// Database primary key
	// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
	Id int64 `json:"id"`
	// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
	// Example: PRJ1731101982N123 ("PRJ" + UNIX Time stamp + "N" + current ID)
	Code string `json:"code"`
	// The name of the project.
	ProjectName string `json:"project_name"`
	// Reference to [Customer](#customer).code
	CustomerCode string      `json:"customer_code"`
	Addresses    []Address   `json:"addresses"`
	Contacts     []Contact   `json:"contacts"`
	Events       []Event     `json:"events"`
	ProjectMeta  ProjectMeta `json:"project_meta"`
	// Flexible key-value map for additional metadata. The value is any json type.
	ProjectMap cu.IM `json:"project_map"`
	// Timestamp of data creation
	TimeStamp TimeDateTime `json:"time_stamp"`
}

type RateMeta struct {
	RateValue float64 `json:"rate_value"`
	// Additional tags for the rate
	Tags []string `json:"tags"`
}

type Rate struct {
	// Database primary key
	// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
	Id int64 `json:"id"`
	// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
	// Example: RAT1731101982N123 ("RAT" + UNIX Time stamp + "N" + current ID)
	Code string `json:"code"`
	// ENUM field. Valid values: RATE, BUY, SELL, AVERAGE
	RateType RateType `json:"rate_type"`
	RateDate TimeDate `json:"rate_date"`
	// Reference to [Place](#place).code
	PlaceCode string `json:"place_code"`
	// Reference to [Currency](#currency).code
	CurrencyCode string   `json:"currency_code"`
	RateMeta     RateMeta `json:"rate_meta"`
	// Flexible key-value map for additional metadata. The value is any json type.
	RateMap cu.IM `json:"rate_map"`
	// Timestamp of data creation
	TimeStamp TimeDateTime `json:"time_stamp"`
}

type TaxMeta struct {
	Description string  `json:"description"`
	RateValue   float64 `json:"rate_value"`
	// Additional tags for the tax
	Tags []string `json:"tags"`
}

type Tax struct {
	// Database primary key
	// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
	Id int64 `json:"id"`
	// Unique tax ID.
	Code    string  `json:"code"`
	TaxMeta TaxMeta `json:"tax_meta"`
	// Flexible key-value map for additional metadata. The value is any json type.
	TaxMap cu.IM `json:"tax_map"`
	// Timestamp of data creation
	TimeStamp TimeDateTime `json:"time_stamp"`
}

type ToolMeta struct {
	// Other specific serial number for the tool. Example: IMEI-023456789, ABC-123, etc.
	SerialNumber string `json:"serial_number"`
	Notes        string `json:"notes"`
	Inactive     bool   `json:"inactive"`
	// Additional tags for the tool
	Tags []string `json:"tags"`
}

type Tool struct {
	// Database primary key
	// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
	Id int64 `json:"id"`
	// Unique serial ID. If not specified, it is generated on creation. It cannot be modified after creation.
	// Example: SER1731101982N123 ("SER" + UNIX Time stamp + "N" + current ID)
	Code string `json:"code"`
	// The description of the tool.
	Description string `json:"description"`
	// Reference to [Product](#product).code
	ProductCode string   `json:"product_code"`
	Events      []Event  `json:"events"`
	ToolMeta    ToolMeta `json:"tool_meta"`
	// Flexible key-value map for additional metadata. The value is any json type.
	ToolMap cu.IM `json:"tool_map"`
	// Timestamp of data creation
	TimeStamp TimeDateTime `json:"time_stamp"`
}

type TransMetaWorksheet struct {
	// Distance (km)
	Distance float64 `json:"distance"`
	// Repair time (hour)
	Repair float64 `json:"repair"`
	// Total time (hour)
	Total float64 `json:"total"`
	// Worksheet justification
	Notes string `json:"notes"`
}

type TransMetaRent struct {
	// Holidays (day)
	Holiday float64 `json:"holiday"`
	// Bad tool / machine (hour)
	BadTool float64 `json:"bad_tool"`
	// Other non-eligible
	Other float64 `json:"other"`
	// Rent justification
	Notes string `json:"notes"`
}

type TransMetaInvoice struct {
	CompanyName       string `json:"company_name"`
	CompanyAddress    string `json:"company_address"`
	CompanyTaxNumber  string `json:"company_tax_number"`
	CompanyAccount    string `json:"company_account"`
	CustomerName      string `json:"customer_name"`
	CustomerAddress   string `json:"customer_address"`
	CustomerTaxNumber string `json:"customer_tax_number"`
	CustomerAccount   string `json:"customer_account"`
}

type TransMeta struct {
	DueTime   TimeDateTime `json:"due_time"`
	RefNumber string       `json:"ref_number"`
	// ENUM field. Valid values: CASH, TRANSFER, CARD, ONLINE, OTHER
	PaidType PaidType `json:"paid_type"`
	TaxFree  bool     `json:"tax_free"`
	Paid     bool     `json:"paid"`
	Rate     float64  `json:"rate"`
	Closed   bool     `json:"closed"`
	// ENUM field. Valid values: NORMAL, CANCELLATION, AMENDMENT
	Status TransStatus `json:"status"`
	// ENUM field. Valid values: OK, NEW, BACK
	TransState    TransState         `json:"trans_state"`
	Notes         string             `json:"notes"`
	InternalNotes string             `json:"internal_notes"`
	ReportNotes   string             `json:"report_notes"`
	Worksheet     TransMetaWorksheet `json:"worksheet,omitempty"`
	Rent          TransMetaRent      `json:"rent,omitempty"`
	Invoice       TransMetaInvoice   `json:"invoice,omitempty"`
	// Additional tags for the trans
	Tags []string `json:"tags"`
}

type Trans struct {
	// Database primary key
	// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
	Id int64 `json:"id"`
	// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
	// Example: INV1731101982N123 ("INV"/"REC"/"ORD"/"WOR"... + UNIX Time stamp + "N" + current ID)
	Code string `json:"code"`
	// ENUM field. Valid values:
	// INVOICE, RECEIPT, ORDER, OFFER, WORKSHEET, RENT, DELIVERY,
	// INVENTORY, WAYBILL, PRODUCTION, FORMULA, BANK, CASH
	TransType TransType `json:"trans_type"`
	TransDate TimeDate  `json:"trans_date"`
	// ENUM field. Valid values: OUT, IN, TRANSFER
	Direction Direction `json:"direction"`
	// Reference to [Trans](#trans).code
	TransCode string `json:"trans_code"`
	// Reference to [Customer](#customer).code
	CustomerCode string `json:"customer_code"`
	// Reference to [Employee](#employee).code
	EmployeeCode string `json:"employee_code"`
	// Reference to [Project](#project).code
	ProjectCode string `json:"project_code"`
	// Reference to [Place](#place).code
	PlaceCode string `json:"place_code"`
	// Reference to [currency](#currency).code
	CurrencyCode string `json:"currency_code"`
	// Reference to [Auth](#auth).code
	AuthCode  string    `json:"auth_code"`
	TransMeta TransMeta `json:"trans_meta"`
	// Flexible key-value map for additional metadata. The value is any json type.
	TransMap cu.IM `json:"trans_map"`
	// Timestamp of data creation
	TimeStamp TimeDateTime `json:"time_stamp"`
	Deleted   bool         `json:"deleted"`
}

type Config struct {
	// Database primary key
	// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
	Id int64 `json:"id"`
	// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
	// Example: CNF1731101982N123 ("CNF" + UNIX Time stamp + "N" + current ID)
	Code string `json:"code"`
	// ENUM field. Valid values: MAP, SHORTCUT, MESSAGE, PATTERN, REPORT, PRINT_QUEUE, DATA
	ConfigType ConfigType   `json:"config_type"`
	Data       interface{}  `json:"data"`
	TimeStamp  TimeDateTime `json:"time_stamp"`
}

type ConfigMap struct {
	FieldName string `json:"field_name"`
	// Enum field. Valid values: BOOL, INTEGER, NUMBER, DATE, DATETIME, STRING, MEMO, ENUM, URL,
	//	CUSTOMER, EMPLOYEE, PLACE, PRODUCT, PROJECT, TOOL, TRANS_ITEM, TRANS_MOVEMENT, TRANS_PAYMENT
	FieldType   FieldType `json:"field_type"`
	Description string    `json:"description"`
	// ENUM list if FieldType is ENUM
	Tags []string `json:"tags"`
	// ENUM field. Valid values: ADDRESS, BARCODE, CONTACT, CURRENCY, CUSTOMER, EMPLOYEE, EVENT, ITEM,
	// MOVEMENT, PAYMENT, PLACE, PRICE, PRODUCT, PROJECT, RATE, TAX, TOOL, USER, TRANS,
	// INVOICE, RECEIPT, ORDER, OFFER, WORKSHEET, RENT, DELIVERY,
	// INVENTORY, WAYBILL, PRODUCTION, FORMULA, BANK, CASH
	Filter []MapFilter `json:"filter"`
}

type ConfigShortcut struct {
	ShortcutKey string         `json:"shortcut_key"`
	Description string         `json:"description"`
	Modul       string         `json:"modul"`
	Icon        string         `json:"icon"`
	Method      ShortcutMethod `json:"method"`
	Funcname    string         `json:"func_name"`
	Address     string         `json:"address"`
	Fields      []struct {
		FieldName   string `json:"field_name"`
		Description string `json:"description"`
		// Enum field. Valid values: BOOL, INTEGER, NUMBER, DATE, DATETIME, STRING, MEMO, ENUM, URL
		FieldType ShortcutField `json:"field_type"`
		Order     int64         `json:"order"`
	} `json:"fields"`
}

type ConfigMessage struct {
	Section string `json:"section"`
	Key     string `json:"key"`
	Lang    string `json:"lang"`
	Value   string `json:"value"`
}

type ConfigPattern struct {
	// ENUM field. Valid values:
	// INVOICE, RECEIPT, ORDER, OFFER, WORKSHEET, RENT, DELIVERY,
	// INVENTORY, WAYBILL, PRODUCTION, FORMULA, BANK, CASH
	TransType      TransType `json:"trans_type"`
	Description    string    `json:"description"`
	Notes          string    `json:"notes"`
	DefaultPattern bool      `json:"default_pattern"`
}

type ConfigPrintQueue struct {
	RefType     string `json:"ref_type"`
	RefCode     string `json:"ref_code"`
	Qty         int64  `json:"qty"`
	ReportCode  string `json:"report_code"`
	Orientation string `json:"orientation"`
	PaperSize   string `json:"paper_size"`
	AuthCode    string `json:"auth_code"`
}

type ConfigReport struct {
	ReportKey string `json:"report_key"`
	// Required
	ReportType string `json:"report_type"`
	// Optional. Valid values:
	// INVOICE, RECEIPT, ORDER, OFFER, WORKSHEET, RENT, DELIVERY,
	// INVENTORY, WAYBILL, PRODUCTION, FORMULA, BANK, CASH
	TransType string `json:"trans_type"`
	// Optional. Valid values: OUT, IN, TRANSFER
	Direction   string `json:"direction"`
	ReportName  string `json:"report_name"`
	Description string `json:"description"`
	Label       string `json:"label"`
	// ENUM field. Valid values: PDF, CSV
	FileType FileType `json:"file_type"`
	Template string   `json:"template"`
}
