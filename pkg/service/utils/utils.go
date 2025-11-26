package utils

import (
	"encoding/base64"
	"net"
	"net/smtp"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	cu "github.com/nervatura/component/pkg/util"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	st "github.com/nervatura/nervatura/v6/pkg/static"
)

// SMToJS - convert map[string]string to JSON compatible map[string]interface{}
func SMToJS(sm cu.SM) cu.IM {
	result := cu.IM{}
	mapValue := func(value string) (ivalue interface{}) {
		if value == "true" || value == "false" {
			return cu.ToBoolean(value, false)
		}
		if fvalue, err := strconv.ParseFloat(value, 64); err == nil {
			return fvalue
		}
		if err := cu.ConvertFromByte([]byte(value), &ivalue); err == nil {
			return ivalue
		}
		return value
	}
	for key, value := range sm {
		result[key] = mapValue(value)
	}
	return result
}

// ConvertToType - convert interface to any type
func ConvertToType(data interface{}, result any) (err error) {
	var dt []byte
	if dt, err = cu.ConvertToByte(data); err == nil {
		err = cu.ConvertFromByte(dt, result)
	}
	return err
}

// IntArrayToString - convert []int64 to string
func IntArrayToString(arr []int64, prefix bool) string {
	pg := []string{}
	for _, value := range arr {
		pg = append(pg, cu.ToString(value, ""))
	}
	if prefix {
		return "{" + strings.Join(pg, ",") + "}"
	}
	return strings.Join(pg, ",")
}

// ToIntArray - convert interface to []int64
func ToIntArray(arr interface{}) (result []int64) {
	result = []int64{}
	if result, found := arr.([]int64); found {
		return result
	}
	if ifa, found := arr.(string); found {
		for _, value := range strings.Split(ifa, ",") {
			result = append(result, cu.ToInteger(value, 0))
		}
	}
	if ifa, found := arr.([]string); found {
		for _, value := range ifa {
			result = append(result, cu.ToInteger(value, 0))
		}
	}
	if ifa, found := arr.([]interface{}); found {
		for _, value := range ifa {
			result = append(result, cu.ToInteger(value, 0))
		}
	}
	return result
}

func ToStringArray(arr interface{}) (result []string) {
	result = []string{}
	if result, found := arr.([]string); found {
		return result
	}
	if ifa, found := arr.([]interface{}); found {
		for _, value := range ifa {
			result = append(result, cu.ToString(value, ""))
		}
	}
	return result
}

// GetMessage - application error and info messages
func GetMessage(lang, key string) string {
	var messages map[string]map[string]string
	var jsonMessages, _ = st.Static.ReadFile("message.json")
	_ = cu.ConvertFromByte(jsonMessages, &messages)
	lang = cu.ToString(lang, "en")
	langMsg := messages[lang]
	if value, found := langMsg[key]; found {
		return value
	}
	return ""
}

// GetMessages - get messages from message.json
func GetMessages() map[string]map[string]string {
	var messages map[string]map[string]string
	var jsonMessages, _ = st.Static.ReadFile("message.json")
	_ = cu.ConvertFromByte(jsonMessages, &messages)
	return messages
}

func GetLangMessages(lang string) map[string]string {
	msg := GetMessages()
	if value, found := msg[lang]; found {
		return value
	}
	return map[string]string{}
}

// GetDataField - get field name and value from struct
func GetDataField(data any, JSONName string) (fieldName string, fieldValue interface{}) {
	for _, sfield := range reflect.VisibleFields(reflect.TypeOf(data)) {
		jtag := strings.Trim(strings.Split(sfield.Tag.Get("json"), ",")[0], " ")
		if jtag == JSONName {
			fieldName = sfield.Name
			fieldValue = reflect.ValueOf(data).FieldByName(fieldName).Interface()
			return fieldName, fieldValue
		}
	}
	return fieldName, fieldValue
}

func GetSessionID() string {
	return cu.RandString(5) + base64.URLEncoding.WithPadding(-1).EncodeToString([]byte(
		cu.ToString(time.Now().Unix(), "")))
}

func SmtpClient(conn net.Conn, host string) (md.SmtpClient, error) {
	return smtp.NewClient(conn, host)
}

func ConvertByteToIMValue(data, initValue any, imap cu.IM, key string) {
	v := reflect.ValueOf(data)
	if data == nil || v.IsNil() || v.Elem().IsNil() {
		data = initValue
	}
	if valueData, err := cu.ConvertToByte(data); err == nil {
		imap[key] = string(valueData[:])
	}
}

func ConvertByteToIMData(data any, imap cu.IM, key string) {
	if valueData, err := cu.ConvertToByte(data); err == nil {
		imap[key] = string(valueData[:])
	}
}

// ToBoolMap - safe map[string]bool conversion
func ToBoolMap(im interface{}, defValue map[string]bool) (result map[string]bool) {
	result = map[string]bool{}
	if im == nil {
		return defValue
	}
	if values, valid := im.(map[string]bool); valid && len(values) > 0 {
		return values
	}
	if values, valid := im.(map[string]interface{}); valid {
		for key, value := range values {
			result[key] = cu.ToBoolean(value, false)
		}
		return result
	}
	return defValue
}

func SortIMData(data []map[string]interface{}, sortField string) {
	sort.Slice(data, func(i, j int) bool {
		a := cu.ToString(data[i][sortField], "")
		b := cu.ToString(data[j][sortField], "")
		return a < b
	})
}

func ToTagList(tags []string) []cu.IM {
	rows := []cu.IM{}
	for _, tag := range tags {
		rows = append(rows, cu.IM{"tag": tag})
	}
	return rows
}

func MetaName(mp cu.IM, key string) string {
	for field := range mp {
		if strings.HasSuffix(field, key) {
			return field
		}
	}
	return ""
}

func AnyPointer[T any](v T) *T {
	return &v
}
