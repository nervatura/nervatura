package mcp

import (
	"testing"

	cu "github.com/nervatura/component/pkg/util"
)

func Test_getSchemaData(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		data    cu.IM
		ms      *ModelSchema
		wantErr bool
	}{
		{
			name: "success",
			data: cu.IM{
				"code": "CUS1731101982N123", "customer_name": "First Customer LTD", "customer_type": "CUSTOMER_COMPANY",
				"customer_map": cu.IM{
					"key1": "value1",
					"key2": "value2",
				},
				"customer_meta": cu.IM{
					"key1": "value1",
					"key2": "value2",
				},
			},
			ms:      CustomerSchema(),
			wantErr: false,
		},
		{
			name: "product",
			data: cu.IM{
				"code": "PRD1731101982N123", "product_name": "Big product", "product_type": "PRODUCT_ITEM", "tax_code": "VAT20",
				"product_map": cu.IM{
					"key1": "value1",
					"key2": "value2",
				},
			},
			ms:      ProductSchema(),
			wantErr: false,
		},
		{
			name: "missing required field",
			data: cu.IM{
				"code": "", "customer_type": "CUSTOMER_COMPANY",
				"customer_map": cu.IM{
					"key1": "value1",
					"key2": "value2",
				},
				"customer_meta": cu.IM{
					"key1": "value1",
					"key2": "value2",
				},
			},
			ms: &ModelSchema{
				Name:          "customer",
				PrimaryFields: []string{"customer_name"},
				Required:      []string{"customer_name"},
				LoadData: func(data any) (modelData, metaData any, err error) {
					return data, nil, nil
				},
				LoadList: func(rows []cu.IM) (items any, err error) {
					return rows, nil
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, _, _, gotErr := getSchemaData(tt.data, tt.ms)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("getSchemaData() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("getSchemaData() succeeded unexpectedly")
			}
		})
	}
}
