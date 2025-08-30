package model

import (
	"io"
	"log/slog"
	"net/http"
	"net/smtp"

	cu "github.com/nervatura/component/pkg/util"
)

const TimeLayout = "2006-01-02 15:04"

type ContextKey struct {
	Name string
}

func (k *ContextKey) String() string {
	return k.Name
}

var AuthOptionsCtxKey = &ContextKey{Name: "auth_options"}
var AuthUserCtxKey = &ContextKey{Name: "auth_user"}
var DataStoreCtxKey = &ContextKey{Name: "data_store"}
var ClientServiceCtxKey = &ContextKey{Name: "client_service"}

// Filter query filter type
type Filter struct {
	Or    bool   // and (False) or (True)
	Field string //Fieldname and alias
	Comp  string //==,!=,<,<=,>,>=,in,is
	Value interface{}
}

// Query data desc. type
type Query struct {
	Fields  []string //Returns fields
	From    string   //Table or doc. name and alias
	Filters []Filter
	Filter  string //filter string (eg. "id=1 and field='value'")
	OrderBy []string
	Limit   int64
	Offset  int64
}

// Update data desc. type
type Update struct {
	Values cu.IM
	Model  string
	IDKey  int64       //Update, delete or insert exec
	Trans  interface{} //Transaction
}

type UpdateDataOptions struct {
	Model      string
	IDKey      int64
	Code       string
	Data       any
	Meta       any
	Fields     []string
	MetaFields []string
	SetValue   func(modelName string, values cu.IM, inputName string, fieldValue interface{}) cu.IM
}

type AuthOptions struct {
	Request           *http.Request
	Config            cu.IM
	AppLog            *slog.Logger
	ParseToken        func(tokenString string, keyMap []cu.SM, config cu.IM) (cu.IM, error)
	ConvertFromReader func(data io.Reader, result interface{}) error
}

type View struct {
	Name    ViewName `json:"name"`
	Filter  string   `json:"filter"`
	OrderBy []string `json:"order_by"`
	Limit   int64    `json:"limit"`
	Offset  int64    `json:"offset"`
}

type SmtpClient interface {
	Auth(smtp.Auth) error
	Mail(string) error
	Rcpt(string) error
	Data() (io.WriteCloser, error)
	Close() error
	Quit() error
}

func JSONBMap(fieldName, value string) any {
	typeMap := map[string]func() any{
		/*
			"tags": func() any {
				var result []string
				_ = cu.ConvertFromByte([]byte(value), &result)
				return result
			},
				"address": func() any {
					var result Address
					_ = ut.ConvertFromByte([]byte(value), &result)
					return result
				},
				"addresses": func() any {
					var result []Address = []Address{}
					_ = ut.ConvertFromByte([]byte(value), &result)
					return result
				},
		*/
	}
	if resultData, found := typeMap[fieldName]; found {
		return resultData()
	}
	var result interface{}
	_ = cu.ConvertFromByte([]byte(value), &result)
	return result
}

func GetEnumString(enumType string, enumValue interface{}) string {
	enumMap := map[string]func() string{
		"user_group": func() string {
			if value, found := enumValue.(UserGroup); found {
				return value.String()
			}
			return ""
		},
		"barcode_type": func() string {
			if value, found := enumValue.(BarcodeType); found {
				return value.String()
			}
			return ""
		},
		"customer_type": func() string {
			if value, found := enumValue.(CustomerType); found {
				return value.String()
			}
			return ""
		},
		"field_type": func() string {
			if value, found := enumValue.(FieldType); found {
				return value.String()
			}
			return ""
		},
		"log_type": func() string {
			if value, found := enumValue.(LogType); found {
				return value.String()
			}
			return ""
		},
		"movement_type": func() string {
			if value, found := enumValue.(MovementType); found {
				return value.String()
			}
			return ""
		},
		"trans_type": func() string {
			if value, found := enumValue.(TransType); found {
				return value.String()
			}
			return ""
		},
		"place_type": func() string {
			if value, found := enumValue.(PlaceType); found {
				return value.String()
			}
			return ""
		},
		"product_type": func() string {
			if value, found := enumValue.(ProductType); found {
				return value.String()
			}
			return ""
		},
		"price_type": func() string {
			if value, found := enumValue.(PriceType); found {
				return value.String()
			}
			return ""
		},
		"rate_type": func() string {
			if value, found := enumValue.(RateType); found {
				return value.String()
			}
			return ""
		},
		"direction": func() string {
			if value, found := enumValue.(Direction); found {
				return value.String()
			}
			return ""
		},
		"paid_type": func() string {
			if value, found := enumValue.(PaidType); found {
				return value.String()
			}
			return ""
		},
		"trans_state": func() string {
			if value, found := enumValue.(TransState); found {
				return value.String()
			}
			return ""
		},
		"trans_status": func() string {
			if value, found := enumValue.(TransStatus); found {
				return value.String()
			}
			return ""
		},
		"config_type": func() string {
			if value, found := enumValue.(ConfigType); found {
				return value.String()
			}
			return ""
		},
		"shortcut_method": func() string {
			if value, found := enumValue.(ShortcutMethod); found {
				return value.String()
			}
			return ""
		},
		"shortcut_field": func() string {
			if value, found := enumValue.(ShortcutField); found {
				return value.String()
			}
			return ""
		},
		"map_filter": func() string {
			if value, found := enumValue.(MapFilter); found {
				return value.String()
			}
			return ""
		},
		"file_type": func() string {
			if value, found := enumValue.(FileType); found {
				return value.String()
			}
			return ""
		},
		"link_type": func() string {
			if value, found := enumValue.(LinkType); found {
				return value.String()
			}
			return ""
		},
	}
	return cu.ToString(enumMap[enumType](), "")
}
