package model

import (
	cu "github.com/nervatura/component/pkg/util"
)

type Address struct {
	Country string  `json:"country" jsonschema:"Country. Example: Country."`
	State   string  `json:"state" jsonschema:"State. Example: State."`
	ZipCode ZipCode `json:"zip_code" jsonschema:"Zip code. Example: 12345"`
	City    string  `json:"city" jsonschema:"City. Example: City."`
	Street  string  `json:"street" jsonschema:"Street. Example: Street."`
	Notes   string  `json:"notes" jsonschema:"Notes. Example: Notes."`
	// Additional tags for the address
	Tags []string `json:"tags" jsonschema:"Additional tags for the contact. The value is an array of strings. Example: [TAG1, TAG2]"`
	// Flexible key-value map for additional metadata. The value is any json type.
	AddressMap cu.IM `json:"address_map" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
}

type Contact struct {
	FirstName string `json:"first_name" jsonschema:"First name."`
	Surname   string `json:"surname" jsonschema:"Surname."`
	Status    string `json:"status" jsonschema:"Status."`
	Phone     string `json:"phone" jsonschema:"Phone."`
	Mobile    string `json:"mobile" jsonschema:"Mobile."`
	Email     string `json:"email" jsonschema:"Email."`
	Notes     string `json:"notes" jsonschema:"Notes."`
	// Additional tags for the contact
	Tags []string `json:"tags" jsonschema:"Additional tags for the contact. The value is an array of strings. Example: [TAG1, TAG2]"`
	// Flexible key-value map for additional metadata. The value is any json type.
	ContactMap cu.IM `json:"contact_map" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
}

type Event struct {
	Uid         string `json:"uid" jsonschema:"Uid. Example: UID1731101982N123"`
	Subject     string `json:"subject" jsonschema:"Subject. Example: Subject."`
	StartTime   string `json:"start_time" jsonschema:"Event start time. ISO 8601 date-time. Example: 2025-01-01T00:00:00Z"`
	EndTime     string `json:"end_time" jsonschema:"Event end time. ISO 8601 date-time. Example: 2025-01-01T00:00:00Z"`
	Place       string `json:"place" jsonschema:"Place. Example: Place."`
	Description string `json:"description" jsonschema:"Description. Example: Description."`
	// Additional tags for the event
	Tags []string `json:"tags" jsonschema:"Additional tags for the event. The value is an array of strings. Example: [TAG1, TAG2]"`
	// Flexible key-value map for additional metadata. The value is any json type.
	EventMap cu.IM `json:"event_map" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
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
	TimeStamp string `json:"time_stamp" jsonschema:"Timestamp of data creation. ISO 8601 date-time. Example: 2025-01-01T00:00:00Z"`
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
	TimeStamp string `json:"time_stamp" jsonschema:"Timestamp of data creation. ISO 8601 date-time. Example: 2025-01-01T00:00:00Z"`
}

type CurrencyMeta struct {
	// The name of the currency.
	Description string `json:"description" jsonschema:"The name of the currency. Example: Euro"`
	// The number of decimal places used for recording and rounding by the program.
	Digit int64 `json:"digit" jsonschema:"The number of decimal places used for recording and rounding by the program. Example: 2"`
	// Rounding value for cash. Could be used in case the smallest banknote in circulation for that certain currency is not 1.
	CashRound int64 `json:"cash_round" jsonschema:"Rounding value for cash. Could be used in case the smallest banknote in circulation for that certain currency is not 1. Example: 1"`
	// Additional tags for the currency
	Tags []string `json:"tags" jsonschema:"Additional tags for the currency. The value is an array of strings. Example: [TAG1, TAG2]"`
}

type Currency struct {
	// Database primary key
	// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
	Id int64 `json:"id" jsonschema:"Database dependent serial number type. Its value is unique only at table level and may change during data migrations."`
	// Database independent unique external key. The ISO 4217 code of the currency. The value is always mandatory
	Code         string       `json:"code" jsonschema:"Database independent unique external key. The ISO 4217 code of the currency. The value is always mandatory. Example: EUR"`
	CurrencyMeta CurrencyMeta `json:"currency_meta"`
	// Flexible key-value map for additional metadata. The value is any json type.
	CurrencyMap cu.IM `json:"currency_map" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
	// Timestamp of data creation
	TimeStamp string `json:"time_stamp" jsonschema:"Timestamp of data creation. ISO 8601 date-time. Example: 2025-01-01T00:00:00Z"`
}

type CustomerMeta struct {
	TaxNumber string `json:"tax_number" jsonschema:"Tax number."`
	Account   string `json:"account" jsonschema:"Account number."`
	// Tax-free
	TaxFree bool `json:"tax_free" jsonschema:"Tax-free customer"`
	// Payment per.
	Terms int64 `json:"terms" jsonschema:"Payment period in days. Example: 10"`
	// Customer's credit limit. Data is used by financial reports.
	CreditLimit float64 `json:"credit_limit" jsonschema:"Customer's credit limit. Data is used by financial reports. Example: 1000"`
	// If new product line is added (offer, order, invoice etc.) all products will receive the discount percentage specified in this field. If the product has a separate customer price, the value specified here will not be considered by the program.
	Discount float64 `json:"discount" jsonschema:"If new product line is added (offer, order, invoice etc.) all products will receive the discount percentage specified in this field."`
	Notes    string  `json:"notes" jsonschema:"Additional notes for the customer."`
	Inactive bool    `json:"inactive" jsonschema:"Inactive customer"`
	// Additional tags for the customer
	Tags []string `json:"tags" jsonschema:"Additional tags for the customer. The value is an array of strings. Example: [TAG1, TAG2]"`
}

type Customer struct {
	// Database primary key
	// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
	Id int64 `json:"id" jsonschema:"Database dependent serial number type. Its value is unique only at table level and may change during data migrations."`
	// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
	// Example: CUS1731101982N123 ("CUS" + UNIX Time stamp + "N" + current ID)
	Code string `json:"code" jsonschema:"Database independent unique external key. It cannot be modified after creation."`
	// ENUM field. Valid values: OWN, COMPANY, PRIVATE, OTHER
	CustomerType CustomerType `json:"customer_type" jsonschema:"Customer type. Enum values. Example: CUSTOMER_COMPANY"`
	// Full name of the customer
	CustomerName string       `json:"customer_name" jsonschema:"Full name of the customer. Required when creating a new customer."`
	Addresses    []Address    `json:"addresses" jsonschema:"Customer address data"`
	Contacts     []Contact    `json:"contacts" jsonschema:"Customer contact data"`
	Events       []Event      `json:"events" jsonschema:"Customer event data"`
	CustomerMeta CustomerMeta `json:"customer_meta" jsonschema:"Customer additional data"`
	// Flexible key-value map for additional metadata. The value is any json type.
	CustomerMap cu.IM `json:"customer_map" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
	// Timestamp of data creation
	TimeStamp string `json:"time_stamp" jsonschema:"Timestamp of data creation. ISO 8601 date-time. Example: 2025-01-01T00:00:00Z"`
}

type EmployeeMeta struct {
	StartDate string `json:"start_date" jsonschema:"Employee start date. Example: 2025-01-01"`
	EndDate   string `json:"end_date" jsonschema:"Employee end date. Example: 2025-01-01"`
	Inactive  bool   `json:"inactive" jsonschema:"Inactive employee"`
	Notes     string `json:"notes" jsonschema:"Additional notes for the employee."`
	// Additional tags for the employee
	Tags []string `json:"tags" jsonschema:"Additional tags for the employee. The value is an array of strings. Example: [TAG1, TAG2]"`
}

type Employee struct {
	// Database primary key
	// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
	Id int64 `json:"id" jsonschema:"Database dependent serial number type. Its value is unique only at table level and may change during data migrations."`
	// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
	// Example: EMP1731101982N123 ("EMP" + UNIX Time stamp + "N" + current ID)
	Code         string       `json:"code" jsonschema:"Database independent unique external key. It cannot be modified after creation. Example: EMP1731101982N123"`
	Address      Address      `json:"address"`
	Contact      Contact      `json:"contact"`
	Events       []Event      `json:"events"`
	EmployeeMeta EmployeeMeta `json:"employee_meta"`
	// Flexible key-value map for additional metadata. The value is any json type.
	EmployeeMap cu.IM `json:"employee_map" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
	// Timestamp of data creation
	TimeStamp string `json:"time_stamp" jsonschema:"Timestamp of data creation. ISO 8601 date-time. Example: 2025-01-01T00:00:00Z"`
}

type ItemMeta struct {
	Unit        string  `json:"unit" jsonschema:"Unit. Example: piece"`
	Qty         float64 `json:"qty" jsonschema:"Quantity."`
	FxPrice     float64 `json:"fx_price" jsonschema:"Foreign price."`
	NetAmount   float64 `json:"net_amount" jsonschema:"Net amount."`
	Discount    float64 `json:"discount" jsonschema:"Discount."`
	VatAmount   float64 `json:"vat_amount" jsonschema:"Vat amount."`
	Amount      float64 `json:"amount" jsonschema:"Amount."`
	Description string  `json:"description" jsonschema:"Description."`
	Deposit     bool    `json:"deposit" jsonschema:"Deposit."`
	OwnStock    float64 `json:"own_stock" jsonschema:"Own stock."`
	ActionPrice bool    `json:"action_price" jsonschema:"Action price."`
	// Additional tags for the item
	Tags []string `json:"tags" jsonschema:"Additional tags for the item. The value is an array of strings. Example: [TAG1, TAG2]"`
}

type Item struct {
	// Database primary key
	// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
	Id int64 `json:"id" jsonschema:"Database dependent serial number type. Its value is unique only at table level and may change during data migrations."`
	// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
	// Example: ITM1731101982N123 ("ITM" + UNIX Time stamp + "N" + current ID)
	Code string `json:"code" jsonschema:"Database independent unique external key. It cannot be modified after creation. Example: ITM1731101982N123"`
	// Reference to [trans](#trans).code
	TransCode string `json:"trans_code" jsonschema:"Reference to trans.code. Example: INV1731101982N123"`
	// Reference to [product](#product).code
	ProductCode string `json:"product_code" jsonschema:"Reference to product.code. Example: PRD1731101982N123"`
	// Reference to [Tax](#tax).code
	TaxCode  string   `json:"tax_code" jsonschema:"Reference to tax.code. Example: VAT20"`
	ItemMeta ItemMeta `json:"item_meta"`
	// Flexible key-value map for additional metadata. The value is any json type.
	ItemMap cu.IM `json:"item_map" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
	// Timestamp of data creation
	TimeStamp string `json:"time_stamp" jsonschema:"Timestamp of data creation. ISO 8601 date-time. Example: 2025-01-01T00:00:00Z"`
}

type LinkMeta struct {
	Qty    float64 `json:"qty"`
	Amount float64 `json:"amount"`
	Rate   float64 `json:"rate"`
	Notes  string  `json:"notes"`
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
	TimeStamp string `json:"time_stamp" jsonschema:"Timestamp of data creation. ISO 8601 date-time. Example: 2025-01-01T00:00:00Z"`
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
	TimeStamp string `json:"time_stamp" jsonschema:"Timestamp of data creation. ISO 8601 date-time. Example: 2025-01-01T00:00:00Z"`
}

type MovementMeta struct {
	Qty    float64 `json:"qty" jsonschema:"Quantity."`
	Notes  string  `json:"notes" jsonschema:"Additional notes for the movement."`
	Shared bool    `json:"shared" jsonschema:"Shared movement"`
	// Additional tags for the movement
	Tags []string `json:"tags" jsonschema:"Additional tags for the movement. The value is an array of strings. Example: [TAG1, TAG2]"`
}

type Movement struct {
	// Database primary key
	// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
	Id int64 `json:"id" jsonschema:"Database dependent serial number type. Its value is unique only at table level and may change during data migrations."`
	// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
	// Example: MOV1731101982N123 ("MOV" + UNIX Time stamp + "N" + current ID)
	Code string `json:"code" jsonschema:"Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation. Example: MOV1731101982N123"`
	// ENUM field. Valid values: INVENTORY, TOOL, PLAN, HEAD
	MovementType MovementType `json:"movement_type" jsonschema:"Movement type. Enum values. Example: MOVEMENT_INVENTORY"`
	// Date-time
	ShippingTime string `json:"shipping_time" jsonschema:"Shipping time. ISO 8601 date-time. Example: 2025-01-01T00:00:00Z"`
	// Reference to [trans](#trans).code
	TransCode string `json:"trans_code" jsonschema:"Reference to trans.code. Example: INV1731101982N123"`
	// Reference to [Product](#product).code
	ProductCode string `json:"product_code" jsonschema:"Reference to product.code. Example: PRD1731101982N123"`
	// Reference to [Tool](#tool).code
	ToolCode string `json:"tool_code" jsonschema:"Optional reference to tool.code (MOVEMENT_TOOL). Example: SER1731101982N123"`
	// Reference to [Place](#place).code
	PlaceCode string `json:"place_code" jsonschema:"Optional reference to place.code (MOVEMENT_INVENTORY). Example: PLA1731101982N123"`
	// Reference to [Item](#item).code
	ItemCode string `json:"item_code" jsonschema:"Optional reference to item.code (MOVEMENT_INVENTORY). Example: ITM1731101982N123"`
	// Reference to [Movement](#movement).code
	MovementCode string       `json:"movement_code" jsonschema:"Optional reference to movement.code (MOVEMENT_INVENTORY). Example: MOV1731101982N123"`
	MovementMeta MovementMeta `json:"movement_meta"`
	// Flexible key-value map for additional metadata. The value is any json type.
	MovementMap cu.IM `json:"movement_map" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
	// Timestamp of data creation
	TimeStamp string `json:"time_stamp" jsonschema:"Timestamp of data creation. ISO 8601 date-time. Example: 2025-01-01T00:00:00Z"`
}

type PaymentMeta struct {
	Amount float64 `json:"amount" jsonschema:"Bank or cash amount."`
	Notes  string  `json:"notes" jsonschema:"Additional notes for the payment."`
	// Additional tags for the payment
	Tags []string `json:"tags" jsonschema:"Additional tags for the payment. The value is an array of strings. Example: [TAG1, TAG2]"`
}

type Payment struct {
	// Database primary key
	// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
	Id int64 `json:"id" jsonschema:"Database dependent serial number type. Its value is unique only at table level and may change during data migrations."`
	// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
	// Example: PMT1731101982N123 ("PMT" + UNIX Time stamp + "N" + current ID)
	Code     string `json:"code" jsonschema:"Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation. Example: PMT1731101982N123"`
	PaidDate string `json:"paid_date" jsonschema:"Date when payment was made. ISO 8601 date. Example: 2025-01-01"`
	// Reference to [trans](#trans).code
	TransCode   string      `json:"trans_code" jsonschema:"Reference to trans.code. Example: INV1731101982N123"`
	PaymentMeta PaymentMeta `json:"payment_meta"`
	// Flexible key-value map for additional metadata. The value is any json type.
	PaymentMap cu.IM `json:"payment_map" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
	// Timestamp of data creation
	TimeStamp string `json:"time_stamp" jsonschema:"Timestamp of data creation. ISO 8601 date-time. Example: 2025-01-01T00:00:00Z"`
}

type PlaceMeta struct {
	Notes    string `json:"notes" jsonschema:"Additional notes for the place."`
	Inactive bool   `json:"inactive" jsonschema:"Inactive place"`
	// Additional tags for the place
	Tags []string `json:"tags" jsonschema:"Additional tags for the place. The value is an array of strings. Example: [TAG1, TAG2]"`
}

type Place struct {
	// Database primary key
	// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
	Id int64 `json:"id" jsonschema:"Database dependent serial number type. Its value is unique only at table level and may change during data migrations."`
	// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
	// Example: PLA1731101982N123 ("PLA" + UNIX Time stamp + "N" + current ID)
	Code string `json:"code" jsonschema:"Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation. Example: PLA1731101982N123"`
	// ENUM field. Valid values: BANK, CASH, WAREHOUSE, OTHER
	PlaceType PlaceType `json:"place_type" jsonschema:"Place type. Enum values. Example: PLACE_WAREHOUSE"`
	// The full name of the place.
	PlaceName string `json:"place_name" jsonschema:"Place name. Example: Main warehouse"`
	// Reference to [currency](#currency).code
	CurrencyCode string    `json:"currency_code" jsonschema:"Optional reference to currency.code (PLACE_BANK, PLACE_CASH). Example: EUR"`
	Address      Address   `json:"address"`
	Contacts     []Contact `json:"contacts"`
	Events       []Event   `json:"events"`
	PlaceMeta    PlaceMeta `json:"place_meta"`
	// Flexible key-value map for additional metadata. The value is any json type.
	PlaceMap cu.IM `json:"place_map" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
	// Timestamp of data creation
	TimeStamp string `json:"time_stamp" jsonschema:"Timestamp of data creation. ISO 8601 date-time. Example: 2025-01-01T00:00:00Z"`
}

type PriceMeta struct {
	// Price value
	PriceValue float64 `json:"price_value" jsonschema:"Price value."`
	// Additional tags for the price
	Tags []string `json:"tags" jsonschema:"Additional tags for the price. The value is an array of strings. Example: [TAG1, TAG2]"`
}

type Price struct {
	// Database primary key
	// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
	Id int64 `json:"id" jsonschema:"Database dependent serial number type. Its value is unique only at table level and may change during data migrations."`
	// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
	// Example: PRC1731101982N123 ("PRC" + UNIX Time stamp + "N" + current ID)
	Code string `json:"code" jsonschema:"Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation. Example: PRC1731101982N123"`
	// ENUM field. Valid values: CUSTOMER, VENDOR
	PriceType PriceType `json:"price_type" jsonschema:"Price type. Enum values. Example: PRICE_CUSTOMER"`
	// Start of validity, mandatory data.
	ValidFrom string `json:"valid_from" jsonschema:"Valid from date. ISO 8601 date. Example: 2025-01-01"`
	// End of validity, can be left empty.
	ValidTo string `json:"valid_to" jsonschema:"Valid validity date. ISO 8601 date. Example: 2025-01-01"`
	// Reference to [Product](#product).code
	ProductCode string `json:"product_code" jsonschema:"Reference to product.code. Example: PRD1731101982N123"`
	// Reference to [Currency](#currency).code
	CurrencyCode string `json:"currency_code" jsonschema:"Reference to currency.code (PLACE_BANK, PLACE_CASH). Example: EUR"`
	// Reference to [Customer](#customer).code
	CustomerCode string `json:"customer_code" jsonschema:"Optional reference to customer.code. Example: CUS1731101982N123"`
	// Price ranges can also be specified, thus different price can be set for a smaller and bigger quantity of the same product.
	// The quantity should be used as the lower threshold, ie. this should be the minimum quantity for the price set.
	Qty       float64   `json:"qty" jsonschema:"Quantity. Price ranges can also be specified, thus different price can be set for a smaller and bigger quantity of the same product. The quantity should be used as the lower threshold, ie. this should be the minimum quantity for the price set."`
	PriceMeta PriceMeta `json:"price_meta"`
	// Flexible key-value map for additional metadata. The value is any json type.
	PriceMap cu.IM `json:"price_map" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
	// Timestamp of data creation
	TimeStamp string `json:"time_stamp" jsonschema:"Timestamp of data creation. ISO 8601 date-time. Example: 2025-01-01T00:00:00Z"`
}

type ProductMeta struct {
	Unit string `json:"unit" jsonschema:"Unit of measurement. Example: piece"`
	// ENUM field. Valid values: CODE_128, CODE_39, EAN_13, EAN_8, QR_CODE
	BarcodeType BarcodeType `json:"barcode_type" jsonschema:"Barcode type. Enum values."`
	// Any barcode or QR code data
	Barcode string `json:"barcode" jsonschema:"Barcode or QR code data."`
	// The actual amount of the products identified by the barcode. For example can be used for packaged goods, tray packaging.
	BarcodeQty float64 `json:"barcode_qty" jsonschema:"Barcode quantity."`
	Notes      string  `json:"notes" jsonschema:"Notes."`
	Inactive   bool    `json:"inactive" jsonschema:"Inactive."`
	// Additional tags for the product
	Tags []string `json:"tags" jsonschema:"Additional tags for the product. The value is an array of strings. Example: [TAG1, TAG2]"`
}

type Product struct {
	// Database primary key
	// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
	Id int64 `json:"id" jsonschema:"Database dependent serial number type. Its value is unique only at table level and may change during data migrations."`
	// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
	// Example: PRD1731101982N123 ("PRD" + UNIX Time stamp + "N" + current ID)
	Code string `json:"code" jsonschema:"Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation."`
	// ENUM field. Valid values: ITEM, SERVICE
	ProductType ProductType `json:"product_type" jsonschema:"Product type. Enum values. Example: PRODUCT_ITEM"`
	// The full name of the product or short description.
	ProductName string `json:"product_name" jsonschema:"Product name. Required when creating a new product."`
	// Reference to [Tax](#tax).code
	TaxCode     string      `json:"tax_code" jsonschema:"Reference to tax.code. Example: VAT25"`
	Events      []Event     `json:"events" jsonschema:"Product events"`
	ProductMeta ProductMeta `json:"product_meta" jsonschema:"Product metadata"`
	// Flexible key-value map for additional metadata. The value is any json type.
	ProductMap cu.IM `json:"product_map" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
	// Timestamp of data creation
	TimeStamp string `json:"time_stamp" jsonschema:"Timestamp of data creation. ISO 8601 date-time. Example: 2025-01-01T00:00:00Z"`
}

type ProjectMeta struct {
	StartDate string `json:"start_date" jsonschema:"Start date. ISO 8601 date. Example: 2025-01-01"`
	EndDate   string `json:"end_date" jsonschema:"End date. ISO 8601 date. Example: 2025-01-01"`
	Notes     string `json:"notes" jsonschema:"Additional notes for the project."`
	Inactive  bool   `json:"inactive" jsonschema:"Inactive project"`
	// Additional tags for the project
	Tags []string `json:"tags" jsonschema:"Additional tags for the project. The value is an array of strings. Example: [TAG1, TAG2]"`
}

type Project struct {
	// Database primary key
	// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
	Id int64 `json:"id" jsonschema:"Database dependent serial number type. Its value is unique only at table level and may change during data migrations."`
	// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
	// Example: PRJ1731101982N123 ("PRJ" + UNIX Time stamp + "N" + current ID)
	Code string `json:"code" jsonschema:"Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation. Example: PRJ1731101982N123"`
	// The name of the project.
	ProjectName string `json:"project_name" jsonschema:"Project name. Example: Project 1"`
	// Reference to [Customer](#customer).code
	CustomerCode string      `json:"customer_code" jsonschema:"Optional reference to customer.code. Example: CUS1731101982N123"`
	Addresses    []Address   `json:"addresses"`
	Contacts     []Contact   `json:"contacts"`
	Events       []Event     `json:"events"`
	ProjectMeta  ProjectMeta `json:"project_meta"`
	// Flexible key-value map for additional metadata. The value is any json type.
	ProjectMap cu.IM `json:"project_map" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
	// Timestamp of data creation
	TimeStamp string `json:"time_stamp" jsonschema:"Timestamp of data creation. ISO 8601 date-time. Example: 2025-01-01T00:00:00Z"`
}

type RateMeta struct {
	RateValue float64 `json:"rate_value" jsonschema:"Rate value."`
	Notes     string  `json:"notes" jsonschema:"Additional notes for the rate."`
	// Additional tags for the rate
	Tags []string `json:"tags" jsonschema:"Additional tags for the rate. The value is an array of strings. Example: [TAG1, TAG2]"`
}

type Rate struct {
	// Database primary key
	// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
	Id int64 `json:"id" jsonschema:"Database dependent serial number type. Its value is unique only at table level and may change during data migrations."`
	// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
	// Example: RAT1731101982N123 ("RAT" + UNIX Time stamp + "N" + current ID)
	Code string `json:"code" jsonschema:"Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation. Example: RAT1731101982N123"`
	// ENUM field. Valid values: RATE, BUY, SELL, AVERAGE
	RateType RateType `json:"rate_type" jsonschema:"Rate type. Enum values. Example: RATE_RATE"`
	RateDate string   `json:"rate_date" jsonschema:"Rate date. ISO 8601 date. Example: 2025-01-01"`
	// Reference to [Place](#place).code
	PlaceCode string `json:"place_code" jsonschema:"Optional reference to place.code (bank rate). Example: PLA1731101982N123"`
	// Reference to [Currency](#currency).code
	CurrencyCode string   `json:"currency_code" jsonschema:"Reference to currency.code. Example: CUS1731101982N123"`
	RateMeta     RateMeta `json:"rate_meta"`
	// Flexible key-value map for additional metadata. The value is any json type.
	RateMap cu.IM `json:"rate_map" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
	// Timestamp of data creation
	TimeStamp string `json:"time_stamp" jsonschema:"Timestamp of data creation. ISO 8601 date-time. Example: 2025-01-01T00:00:00Z"`
}

type TaxMeta struct {
	Description string  `json:"description" jsonschema:"Description."`
	RateValue   float64 `json:"rate_value" jsonschema:"Rate value. Percentage value expressed between 0 and 1."`
	// Additional tags for the tax
	Tags []string `json:"tags" jsonschema:"Additional tags for the tax. The value is an array of strings. Example: [TAG1, TAG2]"`
}

type Tax struct {
	// Database primary key
	// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
	Id int64 `json:"id" jsonschema:"Database dependent serial number type. Its value is unique only at table level and may change during data migrations."`
	// Unique tax ID.
	Code    string  `json:"code" jsonschema:"Database independent unique external key. Required field. It cannot be modified after creation. Example: VAT25"`
	TaxMeta TaxMeta `json:"tax_meta"`
	// Flexible key-value map for additional metadata. The value is any json type.
	TaxMap cu.IM `json:"tax_map" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
	// Timestamp of data creation
	TimeStamp string `json:"time_stamp" jsonschema:"Timestamp of data creation. ISO 8601 date-time. Example: 2025-01-01T00:00:00Z"`
}

type ToolMeta struct {
	// Other specific serial number for the tool. Example: IMEI-023456789, ABC-123, etc.
	SerialNumber string `json:"serial_number" jsonschema:"Serial number or other unique identifier"`
	Notes        string `json:"notes" jsonschema:"Tool notes"`
	Inactive     bool   `json:"inactive" jsonschema:"Inactive tool"`
	// Additional tags for the tool
	Tags []string `json:"tags" jsonschema:"Additional tags for the tax. The value is an array of strings. Example: [TAG1, TAG2]"`
}

type Tool struct {
	// Database primary key
	// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
	Id int64 `json:"id" jsonschema:"Database dependent serial number type. Its value is unique only at table level and may change during data migrations."`
	// Unique serial ID. If not specified, it is generated on creation. It cannot be modified after creation.
	// Example: SER1731101982N123 ("SER" + UNIX Time stamp + "N" + current ID)
	Code string `json:"code" jsonschema:"Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation. Example: SER1731101982N123"`
	// The description of the tool.
	Description string `json:"description" jsonschema:"Tool name or description"`
	// Reference to [Product](#product).code
	ProductCode string   `json:"product_code" jsonschema:"Reference to product.code. Example: PRD1731101982N12"`
	Events      []Event  `json:"events"`
	ToolMeta    ToolMeta `json:"tool_meta"`
	// Flexible key-value map for additional metadata. The value is any json type.
	ToolMap cu.IM `json:"tool_map" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
	// Timestamp of data creation
	TimeStamp string `json:"time_stamp" jsonschema:"Timestamp of data creation. ISO 8601 date-time. Example: 2025-01-01T00:00:00Z"`
}

type TransMetaWorksheet struct {
	// Distance (km)
	Distance float64 `json:"distance" jsonschema:"Distance"`
	// Repair time (hour)
	Repair float64 `json:"repair" jsonschema:"Repair"`
	// Total time (hour)
	Total float64 `json:"total" jsonschema:"Total"`
	// Worksheet justification
	Justification string `json:"justification" jsonschema:"Justification"`
}

type TransMetaRent struct {
	// Holidays (day)
	Holiday float64 `json:"holiday" jsonschema:"Holiday"`
	// Bad tool / machine (hour)
	BadTool float64 `json:"bad_tool" jsonschema:"Bad tool"`
	// Other non-eligible
	Other float64 `json:"other" jsonschema:"Other"`
	// Rent justification
	Justification string `json:"justification" jsonschema:"Justification"`
}

type TransMetaInvoice struct {
	CompanyName       string `json:"company_name" jsonschema:"Company name. Its value is set automatically and is read-only."`
	CompanyAddress    string `json:"company_address" jsonschema:"Company address. Its value is set automatically and is read-only."`
	CompanyTaxNumber  string `json:"company_tax_number" jsonschema:"Company tax number. Its value is set automatically and is read-only."`
	CompanyAccount    string `json:"company_account" jsonschema:"Company account. Its value is set automatically and is read-only."`
	CustomerName      string `json:"customer_name" jsonschema:"Customer name. Its value is set automatically and is read-only."`
	CustomerAddress   string `json:"customer_address" jsonschema:"Customer address. Its value is set automatically and is read-only."`
	CustomerTaxNumber string `json:"customer_tax_number" jsonschema:"Customer tax number. Its value is set automatically and is read-only."`
	CustomerAccount   string `json:"customer_account" jsonschema:"Customer account. Its value is set automatically and is read-only."`
}

type TransMeta struct {
	DueTime   string `json:"due_time" jsonschema:"Due time. ISO 8601 date-time. Example: 2025-01-01T00:00:00Z"`
	RefNumber string `json:"ref_number" jsonschema:"Ref number. Example: REF1731101982N123"`
	// ENUM field. Valid values: CASH, TRANSFER, CARD, ONLINE, OTHER
	PaidType PaidType `json:"paid_type" jsonschema:"Paid type. Enum values. Example: PAID_TYPE_CASH"`
	TaxFree  bool     `json:"tax_free" jsonschema:"Tax free invoice"`
	Paid     bool     `json:"paid" jsonschema:"Paid invoice"`
	Rate     float64  `json:"rate" jsonschema:"Rate"`
	Closed   bool     `json:"closed" jsonschema:"Closed invoice"`
	// ENUM field. Valid values: NORMAL, CANCELLATION, AMENDMENT
	Status TransStatus `json:"status" jsonschema:"Status. Enum values. Example: TRANS_STATUS_NORMAL"`
	// ENUM field. Valid values: OK, NEW, BACK
	TransState    TransState         `json:"trans_state" jsonschema:"Trans state. Enum values. Example: TRANS_STATE_OK"`
	Notes         string             `json:"notes" jsonschema:"Notes. Example: Notes."`
	InternalNotes string             `json:"internal_notes" jsonschema:"Internal notes. Example: Internal notes."`
	ReportNotes   string             `json:"report_notes" jsonschema:"Report notes. Example: Report notes."`
	Worksheet     TransMetaWorksheet `json:"worksheet,omitempty"`
	Rent          TransMetaRent      `json:"rent,omitempty"`
	Invoice       TransMetaInvoice   `json:"invoice,omitempty"`
	// Additional tags for the trans
	Tags []string `json:"tags" jsonschema:"Additional tags for the customer. The value is an array of strings. Example: [TAG1, TAG2]"`
}

type Trans struct {
	// Database primary key
	// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
	Id int64 `json:"id"`
	// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
	// Example: INV1731101982N123 ("INV"/"REC"/"ORD"/"WOR"... + UNIX Time stamp + "N" + current ID)
	Code string `json:"code" jsonschema:"Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation."`
	// ENUM field. Valid values:
	// INVOICE, RECEIPT, ORDER, OFFER, WORKSHEET, RENT, DELIVERY,
	// INVENTORY, WAYBILL, PRODUCTION, FORMULA, BANK, CASH
	TransType TransType `json:"trans_type" jsonschema:"Trans type. Enum values. Example: TRANS_TYPE_INVOICE"`
	TransDate string    `json:"trans_date" jsonschema:"Trans date. ISO 8601 date. Example: 2025-01-01"`
	// ENUM field. Valid values: OUT, IN, TRANSFER
	Direction Direction `json:"direction" jsonschema:"Direction. Enum values. Example: DIRECTION_OUT"`
	// Reference to [Trans](#trans).code
	TransCode string `json:"trans_code" jsonschema:"Optional reference to other transaction (invoice, receipt, offer, order, worksheet, rent etc.) reference. Example: ORD1731101982N123"`
	// Reference to [Customer](#customer).code
	CustomerCode string `json:"customer_code" jsonschema:"Optional reference to customer.code. Example: CUS1731101982N123"`
	// Reference to [Employee](#employee).code
	EmployeeCode string `json:"employee_code" jsonschema:"Optional reference to employee.code. Example: EMP1731101982N123"`
	// Reference to [Project](#project).code
	ProjectCode string `json:"project_code" jsonschema:"Optional reference to project.code. Example: PRJ1731101982N123"`
	// Reference to [Place](#place).code
	PlaceCode string `json:"place_code" jsonschema:"Optional reference to place.code. Example: PLA1731101982N123"`
	// Reference to [currency](#currency).code
	CurrencyCode string `json:"currency_code" jsonschema:"Optional reference to currency.code. Example: EUR"`
	// Reference to [Auth](#auth).code
	AuthCode  string    `json:"auth_code" jsonschema:"Auth reference. Example: AUT1731101982N123"`
	TransMeta TransMeta `json:"trans_meta"`
	// Flexible key-value map for additional metadata. The value is any json type.
	TransMap cu.IM `json:"trans_map" jsonschema:"Flexible key-value map for additional metadata. The value is any json type."`
	// Timestamp of data creation
	TimeStamp string `json:"time_stamp" jsonschema:"Timestamp of data creation. ISO 8601 date-time. Example: 2025-01-01T00:00:00Z"`
}

type Config struct {
	// Database primary key
	// Database dependent serial number type. Its value is unique only at table level and may change during data migrations.
	Id int64 `json:"id" jsonschema:"Database dependent serial number type. Its value is unique only at table level and may change during data migrations."`
	// Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation.
	// Example: CNF1731101982N123 ("CNF" + UNIX Time stamp + "N" + current ID)
	Code string `json:"code" jsonschema:"Database independent unique external key. If not specified, it is generated on creation. It cannot be modified after creation."`
	// ENUM field. Valid values: MAP, SHORTCUT, MESSAGE, PATTERN, REPORT, PRINT_QUEUE, DATA
	ConfigType ConfigType  `json:"config_type" jsonschema:"Config type. Enum values. Example: CONFIG_TYPE_MAP"`
	Data       interface{} `json:"data"`
	TimeStamp  string      `json:"time_stamp" jsonschema:"Timestamp of data creation. ISO 8601 date-time. Example: 2025-01-01T00:00:00Z"`
}

type ConfigMap struct {
	FieldName string `json:"field_name" jsonschema:"Field name. Example: car_color"`
	// Enum field. Valid values: BOOL, INTEGER, NUMBER, DATE, DATETIME, STRING, MEMO, ENUM, URL,
	//	CUSTOMER, EMPLOYEE, PLACE, PRODUCT, PROJECT, TOOL, TRANS_ITEM, TRANS_MOVEMENT, TRANS_PAYMENT
	FieldType   FieldType `json:"field_type" jsonschema:"Field type. Enum values. Example: FIELD_TYPE_ENUM"`
	Description string    `json:"description" jsonschema:"Description. Example: Car color."`
	// ENUM list if FieldType is ENUM
	Tags []string `json:"tags" jsonschema:"Enum list if FieldType is ENUM. The value is an array of strings. Example: [RED, BLUE]"`
	// ENUM field. Valid values: ADDRESS, BARCODE, CONTACT, CURRENCY, CUSTOMER, EMPLOYEE, EVENT, ITEM,
	// MOVEMENT, PAYMENT, PLACE, PRICE, PRODUCT, PROJECT, RATE, TAX, TOOL, USER, TRANS,
	// INVOICE, RECEIPT, ORDER, OFFER, WORKSHEET, RENT, DELIVERY,
	// INVENTORY, WAYBILL, PRODUCTION, FORMULA, BANK, CASH
	Filter []MapFilter `json:"filter" jsonschema:"Limiting the visibility of data in the user interface. Array of enum values. Example: [FILTER_PRODUCT, FILTER_EMPLOYEE]"`
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
	ReportKey string `json:"report_key" jsonschema:"Report key."`
	// Required
	ReportType string `json:"report_type" jsonschema:"Report type."`
	// Optional. Valid values:
	// INVOICE, RECEIPT, ORDER, OFFER, WORKSHEET, RENT, DELIVERY,
	// INVENTORY, WAYBILL, PRODUCTION, FORMULA, BANK, CASH
	TransType string `json:"trans_type" jsonschema:"Trans type."`
	// Optional. Valid values: OUT, IN, TRANSFER
	Direction   string `json:"direction" jsonschema:"Direction."`
	ReportName  string `json:"report_name" jsonschema:"Report name."`
	Description string `json:"description" jsonschema:"Description."`
	Label       string `json:"label" jsonschema:"Label."`
	// ENUM field. Valid values: PDF, CSV
	FileType FileType `json:"file_type" jsonschema:"File type."`
	Template string   `json:"template" jsonschema:"Template."`
}
