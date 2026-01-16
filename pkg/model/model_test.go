package model_test

import (
	"testing"

	"github.com/nervatura/nervatura/v6/pkg/model"
)

func TestJSONBMap(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		fieldName string
		value     string
		want      any
	}{
		{
			name:      "test_JSONBMap",
			fieldName: "test",
			value:     "test",
			want:      "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			model.JSONBMap(tt.fieldName, tt.value)
		})
	}
}

func TestGetEnumString(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		enumType  string
		enumValue interface{}
		want      string
	}{
		{
			name:      "user_group",
			enumType:  "user_group",
			enumValue: model.UserGroupUser,
			want:      "GROUP_USER",
		},
		{
			name:      "user_group_missing",
			enumType:  "user_group",
			enumValue: "test",
			want:      "",
		},
		{
			name:      "barcode_type",
			enumType:  "barcode_type",
			enumValue: model.BarcodeTypeCode39,
			want:      "BARCODE_CODE_39",
		},
		{
			name:      "barcode_type_missing",
			enumType:  "barcode_type",
			enumValue: "test",
			want:      "",
		},
		{
			name:      "customer_type",
			enumType:  "customer_type",
			enumValue: model.CustomerTypeCompany,
			want:      "CUSTOMER_COMPANY",
		},
		{
			name:      "customer_type_missing",
			enumType:  "customer_type",
			enumValue: "test",
			want:      "",
		},
		{
			name:      "field_type",
			enumType:  "field_type",
			enumValue: model.FieldTypeString,
			want:      "FIELD_STRING",
		},
		{
			name:      "field_type_missing",
			enumType:  "field_type",
			enumValue: "test",
			want:      "",
		},
		{
			name:      "log_type",
			enumType:  "log_type",
			enumValue: model.LogTypeInsert,
			want:      "LOG_INSERT",
		},
		{
			name:      "log_type_missing",
			enumType:  "log_type",
			enumValue: "test",
			want:      "",
		},
		{
			name:      "movement_type",
			enumType:  "movement_type",
			enumValue: model.MovementTypeInventory,
			want:      "MOVEMENT_INVENTORY",
		},
		{
			name:      "movement_type_missing",
			enumType:  "movement_type",
			enumValue: "test",
			want:      "",
		},
		{
			name:      "trans_type",
			enumType:  "trans_type",
			enumValue: model.TransTypeInvoice,
			want:      "TRANS_INVOICE",
		},
		{
			name:      "trans_type_missing",
			enumType:  "trans_type",
			enumValue: "test",
			want:      "",
		},
		{
			name:      "place_type",
			enumType:  "place_type",
			enumValue: model.PlaceTypeWarehouse,
			want:      "PLACE_WAREHOUSE",
		},
		{
			name:      "place_type_missing",
			enumType:  "place_type",
			enumValue: "test",
			want:      "",
		},
		{
			name:      "product_type",
			enumType:  "product_type",
			enumValue: model.ProductTypeItem,
			want:      "PRODUCT_ITEM",
		},
		{
			name:      "product_type_missing",
			enumType:  "product_type",
			enumValue: "test",
			want:      "",
		},
		{
			name:      "price_type",
			enumType:  "price_type",
			enumValue: model.PriceTypeCustomer,
			want:      "PRICE_CUSTOMER",
		},
		{
			name:      "price_type_missing",
			enumType:  "price_type",
			enumValue: "test",
			want:      "",
		},
		{
			name:      "rate_type",
			enumType:  "rate_type",
			enumValue: model.RateTypeRate,
			want:      "RATE_RATE",
		},
		{
			name:      "rate_type_missing",
			enumType:  "rate_type",
			enumValue: "test",
			want:      "",
		},
		{
			name:      "direction",
			enumType:  "direction",
			enumValue: model.DirectionIn,
			want:      "DIRECTION_IN",
		},
		{
			name:      "direction_missing",
			enumType:  "direction",
			enumValue: "test",
			want:      "",
		},
		{
			name:      "paid_type",
			enumType:  "paid_type",
			enumValue: model.PaidTypeOnline,
			want:      "PAID_ONLINE",
		},
		{
			name:      "paid_type_missing",
			enumType:  "paid_type",
			enumValue: "test",
			want:      "",
		},
		{
			name:      "trans_state",
			enumType:  "trans_state",
			enumValue: model.TransStateOK,
			want:      "STATE_OK",
		},
		{
			name:      "trans_state_missing",
			enumType:  "trans_state",
			enumValue: "test",
			want:      "",
		},
		{
			name:      "trans_status",
			enumType:  "trans_status",
			enumValue: model.TransStatusNormal,
			want:      "STATUS_NORMAL",
		},
		{
			name:      "trans_status_missing",
			enumType:  "trans_status",
			enumValue: "test",
			want:      "",
		},
		{
			name:      "config_type",
			enumType:  "config_type",
			enumValue: model.ConfigTypeMap,
			want:      "CONFIG_MAP",
		},
		{
			name:      "config_type_missing",
			enumType:  "config_type",
			enumValue: "test",
			want:      "",
		},
		{
			name:      "shortcut_method",
			enumType:  "shortcut_method",
			enumValue: model.ShortcutMethodGET,
			want:      "METHOD_GET",
		},
		{
			name:      "shortcut_method_missing",
			enumType:  "shortcut_method",
			enumValue: "test",
			want:      "",
		},
		{
			name:      "shortcut_field",
			enumType:  "shortcut_field",
			enumValue: model.ShortcutFieldInteger,
			want:      "SHORTCUT_INTEGER",
		},
		{
			name:      "shortcut_field_missing",
			enumType:  "shortcut_field",
			enumValue: "test",
			want:      "",
		},
		{
			name:      "map_filter",
			enumType:  "map_filter",
			enumValue: model.MapFilterAddress,
			want:      "FILTER_ADDRESS",
		},
		{
			name:      "map_filter_missing",
			enumType:  "map_filter",
			enumValue: "test",
			want:      "",
		},
		{
			name:      "file_type",
			enumType:  "file_type",
			enumValue: model.FileTypePDF,
			want:      "FILE_PDF",
		},
		{
			name:      "file_type_missing",
			enumType:  "file_type",
			enumValue: "test",
			want:      "",
		},
		{
			name:      "link_type",
			enumType:  "link_type",
			enumValue: model.LinkTypeCustomer,
			want:      "LINK_CUSTOMER",
		},
		{
			name:      "link_type_missing",
			enumType:  "link_type",
			enumValue: "test",
			want:      "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := model.GetEnumString(tt.enumType, tt.enumValue)
			if got != tt.want {
				t.Errorf("GetEnumString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContextKey_String(t *testing.T) {
	tests := []struct {
		name string // description of this test case
	}{
		{
			name: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var k model.ContextKey
			if k.String() != tt.name {
				t.Errorf("ContextKey_String() = %v, want %v", k.String(), tt.name)
			}
		})
	}
}
