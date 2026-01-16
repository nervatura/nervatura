package model_test

import (
	"testing"

	"github.com/nervatura/nervatura/v6/pkg/model"
)

func TestZipCode_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		b       []byte
		wantErr bool
	}{
		{
			name:    "test_ZipCode_UnmarshalJSON",
			b:       []byte("\"12345\""),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var z model.ZipCode
			gotErr := z.UnmarshalJSON(tt.b)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("UnmarshalJSON() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("UnmarshalJSON() succeeded unexpectedly")
			}
		})
	}
}

func TestUserGroup_String(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		ug   model.UserGroup
		want string
	}{
		{
			name: "test_UserGroup_String",
			ug:   model.UserGroupUser,
			want: "GROUP_USER",
		},
		{
			name: "test_UserGroup_String_missing",
			ug:   model.UserGroup(100),
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.ug.String() != tt.want {
				t.Errorf("String() = %v, want %v", tt.ug.String(), tt.want)
			}
			tt.ug.Keys()
			tt.ug.UnmarshalJSON([]byte("\"GROUP_USER\""))
			tt.ug.UnmarshalJSON([]byte("\"INVALID\""))
			tt.ug.MarshalJSON()
		})
	}
}

func TestBarcodeType_String(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		bt   model.BarcodeType
		want string
	}{
		{
			name: "test_BarcodeType_String",
			bt:   model.BarcodeTypeCode39,
			want: "BARCODE_CODE_39",
		},
		{
			name: "test_BarcodeType_String_missing",
			bt:   model.BarcodeType(100),
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.bt.String() != tt.want {
				t.Errorf("String() = %v, want %v", tt.bt.String(), tt.want)
			}
			tt.bt.UnmarshalJSON([]byte("\"BARCODE_CODE_39\""))
			tt.bt.UnmarshalJSON([]byte("INVALID\""))
			tt.bt.UnmarshalJSON([]byte("\"INVALID\""))
			tt.bt.MarshalJSON()
		})
	}
}

func TestBookmarkType_String(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		bt   model.BookmarkType
		want string
	}{
		{
			name: "test_BookmarkType_String",
			bt:   model.BookmarkTypeBrowser,
			want: "BOOKMARK_BROWSER",
		},
		{
			name: "test_BookmarkType_String_missing",
			bt:   model.BookmarkType(100),
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.bt.String() != tt.want {
				t.Errorf("String() = %v, want %v", tt.bt.String(), tt.want)
			}
			tt.bt.UnmarshalJSON([]byte("\"BOOKMARK_BROWSER\""))
			tt.bt.UnmarshalJSON([]byte("INVALID\""))
			tt.bt.UnmarshalJSON([]byte("\"INVALID\""))
			tt.bt.MarshalJSON()
		})
	}
}

func TestCustomerType_String(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		ct   model.CustomerType
		want string
	}{
		{
			name: "test_CustomerType_String",
			ct:   model.CustomerTypeCompany,
			want: "CUSTOMER_COMPANY",
		},
		{
			name: "test_CustomerType_String_missing",
			ct:   model.CustomerType(100),
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.ct.String() != tt.want {
				t.Errorf("String() = %v, want %v", tt.ct.String(), tt.want)
			}
			tt.ct.Keys()
			tt.ct.Get("CUSTOMER_COMPANY")
			tt.ct.Get("INVALID")
			tt.ct.UnmarshalJSON([]byte("\"CUSTOMER_COMPANY\""))
			tt.ct.UnmarshalJSON([]byte("INVALID\""))
			tt.ct.UnmarshalJSON([]byte("\"INVALID\""))
			tt.ct.MarshalJSON()
		})
	}
}

func TestFieldType_String(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		ft   model.FieldType
		want string
	}{
		{
			name: "test_FieldType_String",
			ft:   model.FieldTypeString,
			want: "FIELD_STRING",
		},
		{
			name: "test_FieldType_String_missing",
			ft:   model.FieldType(100),
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.ft.String() != tt.want {
				t.Errorf("String() = %v, want %v", tt.ft.String(), tt.want)
			}
			tt.ft.Keys()
			tt.ft.UnmarshalJSON([]byte("\"FIELD_STRING\""))
			tt.ft.UnmarshalJSON([]byte("INVALID\""))
			tt.ft.UnmarshalJSON([]byte("\"INVALID\""))
			tt.ft.MarshalJSON()
		})
	}
}

func TestLogType_String(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		lt   model.LogType
		want string
	}{
		{
			name: "test_LogType_String",
			lt:   model.LogTypeInsert,
			want: "LOG_INSERT",
		},
		{
			name: "test_LogType_String_missing",
			lt:   model.LogType(100),
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.lt.String() != tt.want {
				t.Errorf("String() = %v, want %v", tt.lt.String(), tt.want)
			}
			tt.lt.UnmarshalJSON([]byte("\"LOG_INSERT\""))
			tt.lt.UnmarshalJSON([]byte("INVALID\""))
			tt.lt.UnmarshalJSON([]byte("\"INVALID\""))
			tt.lt.MarshalJSON()
		})
	}
}

func TestMovementType_String(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		mt   model.MovementType
		want string
	}{
		{
			name: "test_MovementType_String",
			mt:   model.MovementTypeInventory,
			want: "MOVEMENT_INVENTORY",
		},
		{
			name: "test_MovementType_String_missing",
			mt:   model.MovementType(100),
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mt.String() != tt.want {
				t.Errorf("String() = %v, want %v", tt.mt.String(), tt.want)
			}
			tt.mt.UnmarshalJSON([]byte("\"MOVEMENT_INVENTORY\""))
			tt.mt.UnmarshalJSON([]byte("INVALID\""))
			tt.mt.UnmarshalJSON([]byte("\"INVALID\""))
			tt.mt.MarshalJSON()
		})
	}
}

func TestTransType_String(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		tt   model.TransType
		want string
	}{
		{
			name: "test_TransType_String",
			tt:   model.TransTypeInvoice,
			want: "TRANS_INVOICE",
		},
		{
			name: "test_TransType_String_missing",
			tt:   model.TransType(100),
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.tt.String() != tt.want {
				t.Errorf("String() = %v, want %v", tt.tt.String(), tt.want)
			}
			tt.tt.Keys()
			tt.tt.UnmarshalJSON([]byte("\"TRANS_INVOICE\""))
			tt.tt.UnmarshalJSON([]byte("INVALID\""))
			tt.tt.UnmarshalJSON([]byte("\"INVALID\""))
			tt.tt.MarshalJSON()
		})
	}
}

func TestPlaceType_String(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		pt   model.PlaceType
		want string
	}{
		{
			name: "test_PlaceType_String",
			pt:   model.PlaceTypeWarehouse,
			want: "PLACE_WAREHOUSE",
		},
		{
			name: "test_PlaceType_String_missing",
			pt:   model.PlaceType(100),
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.pt.String() != tt.want {
				t.Errorf("String() = %v, want %v", tt.pt.String(), tt.want)
			}
			tt.pt.Keys()
			tt.pt.UnmarshalJSON([]byte("\"PLACE_WAREHOUSE\""))
			tt.pt.UnmarshalJSON([]byte("INVALID\""))
			tt.pt.UnmarshalJSON([]byte("\"INVALID\""))
			tt.pt.MarshalJSON()
		})
	}
}

func TestProductType_String(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		pt   model.ProductType
		want string
	}{
		{
			name: "test_ProductType_String",
			pt:   model.ProductTypeItem,
			want: "PRODUCT_ITEM",
		},
		{
			name: "test_ProductType_String_missing",
			pt:   model.ProductType(100),
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.pt.String() != tt.want {
				t.Errorf("String() = %v, want %v", tt.pt.String(), tt.want)
			}
			tt.pt.Keys()
			tt.pt.UnmarshalJSON([]byte("\"PRODUCT_ITEM\""))
			tt.pt.UnmarshalJSON([]byte("INVALID\""))
			tt.pt.UnmarshalJSON([]byte("\"INVALID\""))
			tt.pt.MarshalJSON()
		})
	}
}

func TestPriceType_String(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		pt   model.PriceType
		want string
	}{
		{
			name: "test_PriceType_String",
			pt:   model.PriceTypeCustomer,
			want: "PRICE_CUSTOMER",
		},
		{
			name: "test_PriceType_String_missing",
			pt:   model.PriceType(100),
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.pt.String() != tt.want {
				t.Errorf("String() = %v, want %v", tt.pt.String(), tt.want)
			}
			tt.pt.Keys()
			tt.pt.UnmarshalJSON([]byte("\"PRICE_CUSTOMER\""))
			tt.pt.UnmarshalJSON([]byte("INVALID\""))
			tt.pt.UnmarshalJSON([]byte("\"INVALID\""))
			tt.pt.MarshalJSON()
		})
	}
}

func TestRateType_String(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		rt   model.RateType
		want string
	}{
		{
			name: "test_RateType_String",
			rt:   model.RateTypeRate,
			want: "RATE_RATE",
		},
		{
			name: "test_RateType_String_missing",
			rt:   model.RateType(100),
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.rt.String() != tt.want {
				t.Errorf("String() = %v, want %v", tt.rt.String(), tt.want)
			}
			tt.rt.Keys()
			tt.rt.UnmarshalJSON([]byte("\"RATE_RATE\""))
			tt.rt.UnmarshalJSON([]byte("INVALID\""))
			tt.rt.UnmarshalJSON([]byte("\"INVALID\""))
			tt.rt.MarshalJSON()
		})
	}
}

func TestDirection_String(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		dr   model.Direction
		want string
	}{
		{
			name: "test_Direction_String",
			dr:   model.DirectionIn,
			want: "DIRECTION_IN",
		},
		{
			name: "test_Direction_String_missing",
			dr:   model.Direction(100),
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.dr.String() != tt.want {
				t.Errorf("String() = %v, want %v", tt.dr.String(), tt.want)
			}
			tt.dr.UnmarshalJSON([]byte("\"DIRECTION_IN\""))
			tt.dr.UnmarshalJSON([]byte("INVALID\""))
			tt.dr.UnmarshalJSON([]byte("\"INVALID\""))
			tt.dr.MarshalJSON()
		})
	}
}

func TestPaidType_String(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		pt   model.PaidType
		want string
	}{
		{
			name: "test_PaidType_String",
			pt:   model.PaidTypeOnline,
			want: "PAID_ONLINE",
		},
		{
			name: "test_PaidType_String_missing",
			pt:   model.PaidType(100),
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.pt.String() != tt.want {
				t.Errorf("String() = %v, want %v", tt.pt.String(), tt.want)
			}
			tt.pt.Keys()
			tt.pt.UnmarshalJSON([]byte("\"PAID_ONLINE\""))
			tt.pt.UnmarshalJSON([]byte("INVALID\""))
			tt.pt.UnmarshalJSON([]byte("\"INVALID\""))
			tt.pt.MarshalJSON()
		})
	}
}

func TestTransState_String(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		ts   model.TransState
		want string
	}{
		{
			name: "test_TransState_String",
			ts:   model.TransStateOK,
			want: "STATE_OK",
		},
		{
			name: "test_TransState_String_missing",
			ts:   model.TransState(100),
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.ts.String() != tt.want {
				t.Errorf("String() = %v, want %v", tt.ts.String(), tt.want)
			}
			tt.ts.Keys()
			tt.ts.UnmarshalJSON([]byte("\"STATE_OK\""))
			tt.ts.UnmarshalJSON([]byte("INVALID\""))
			tt.ts.UnmarshalJSON([]byte("\"INVALID\""))
			tt.ts.MarshalJSON()
		})
	}
}

func TestTransStatus_String(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		ts   model.TransStatus
		want string
	}{
		{
			name: "test_TransStatus_String",
			ts:   model.TransStatusNormal,
			want: "STATUS_NORMAL",
		},
		{
			name: "test_TransStatus_String_missing",
			ts:   model.TransStatus(100),
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.ts.String() != tt.want {
				t.Errorf("String() = %v, want %v", tt.ts.String(), tt.want)
			}
			tt.ts.Get("STATUS_NORMAL")
			tt.ts.Get("INVALID")
			tt.ts.UnmarshalJSON([]byte("\"STATUS_NORMAL\""))
			tt.ts.UnmarshalJSON([]byte("INVALID\""))
			tt.ts.UnmarshalJSON([]byte("\"INVALID\""))
			tt.ts.MarshalJSON()
		})
	}
}

func TestConfigType_String(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		ct   model.ConfigType
		want string
	}{
		{
			name: "test_ConfigType_String",
			ct:   model.ConfigTypeMap,
			want: "CONFIG_MAP",
		},
		{
			name: "test_ConfigType_String_missing",
			ct:   model.ConfigType(100),
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.ct.String() != tt.want {
				t.Errorf("String() = %v, want %v", tt.ct.String(), tt.want)
			}
			tt.ct.UnmarshalJSON([]byte("\"CONFIG_MAP\""))
			tt.ct.UnmarshalJSON([]byte("INVALID\""))
			tt.ct.UnmarshalJSON([]byte("\"INVALID\""))
			tt.ct.MarshalJSON()
		})
	}
}

func TestShortcutMethod_String(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		sm   model.ShortcutMethod
		want string
	}{
		{
			name: "test_ShortcutMethod_String",
			sm:   model.ShortcutMethodGET,
			want: "METHOD_GET",
		},
		{
			name: "test_ShortcutMethod_String_missing",
			sm:   model.ShortcutMethod(100),
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.sm.String() != tt.want {
				t.Errorf("String() = %v, want %v", tt.sm.String(), tt.want)
			}
			tt.sm.Keys()
			tt.sm.UnmarshalJSON([]byte("\"METHOD_GET\""))
			tt.sm.UnmarshalJSON([]byte("INVALID\""))
			tt.sm.UnmarshalJSON([]byte("\"INVALID\""))
			tt.sm.MarshalJSON()
		})
	}
}

func TestShortcutField_String(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		sf   model.ShortcutField
		want string
	}{
		{
			name: "test_ShortcutField_String",
			sf:   model.ShortcutFieldInteger,
			want: "SHORTCUT_INTEGER",
		},
		{
			name: "test_ShortcutField_String_missing",
			sf:   model.ShortcutField(100),
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.sf.String() != tt.want {
				t.Errorf("String() = %v, want %v", tt.sf.String(), tt.want)
			}
			tt.sf.Keys()
			tt.sf.UnmarshalJSON([]byte("\"SHORTCUT_INTEGER\""))
			tt.sf.UnmarshalJSON([]byte("INVALID\""))
			tt.sf.UnmarshalJSON([]byte("\"INVALID\""))
			tt.sf.MarshalJSON()
		})
	}
}

func TestMapFilter_String(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		mf   model.MapFilter
		want string
	}{
		{
			name: "test_MapFilter_String",
			mf:   model.MapFilterAddress,
			want: "FILTER_ADDRESS",
		},
		{
			name: "test_MapFilter_String_missing",
			mf:   model.MapFilter(100),
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mf.String() != tt.want {
				t.Errorf("String() = %v, want %v", tt.mf.String(), tt.want)
			}
			tt.mf.Keys()
			tt.mf.UnmarshalJSON([]byte("\"FILTER_ADDRESS\""))
			tt.mf.UnmarshalJSON([]byte("INVALID\""))
			tt.mf.UnmarshalJSON([]byte("\"INVALID\""))
			tt.mf.MarshalJSON()
		})
	}
}

func TestFileType_String(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		ft   model.FileType
		want string
	}{
		{
			name: "test_FileType_String",
			ft:   model.FileTypePDF,
			want: "FILE_PDF",
		},
		{
			name: "test_FileType_String_missing",
			ft:   model.FileType(100),
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.ft.String() != tt.want {
				t.Errorf("String() = %v, want %v", tt.ft.String(), tt.want)
			}
			tt.ft.Value("FILE_PDF")
			tt.ft.Value("INVALID")
			tt.ft.UnmarshalJSON([]byte("\"FILE_PDF\""))
			tt.ft.UnmarshalJSON([]byte("INVALID\""))
			tt.ft.UnmarshalJSON([]byte("\"INVALID\""))
			tt.ft.MarshalJSON()
		})
	}
}

func TestLinkType_String(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		lt   model.LinkType
		want string
	}{
		{
			name: "test_LinkType_String",
			lt:   model.LinkTypeCustomer,
			want: "LINK_CUSTOMER",
		},
		{
			name: "test_LinkType_String_missing",
			lt:   model.LinkType(100),
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.lt.String() != tt.want {
				t.Errorf("String() = %v, want %v", tt.lt.String(), tt.want)
			}
			tt.lt.UnmarshalJSON([]byte("\"LINK_CUSTOMER\""))
			tt.lt.UnmarshalJSON([]byte("INVALID\""))
			tt.lt.UnmarshalJSON([]byte("\"INVALID\""))
			tt.lt.MarshalJSON()
		})
	}
}

func TestViewName_String(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		vn   model.ViewName
		want string
	}{
		{
			name: "test_ViewName_String",
			vn:   model.ViewConfigMap,
			want: "VIEW_CONFIG_MAP",
		},
		{
			name: "test_ViewName_String_missing",
			vn:   model.ViewName(100),
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.vn.String() != tt.want {
				t.Errorf("String() = %v, want %v", tt.vn.String(), tt.want)
			}
			tt.vn.UnmarshalJSON([]byte("\"VIEW_CONFIG_MAP\""))
			tt.vn.UnmarshalJSON([]byte("INVALID\""))
			tt.vn.UnmarshalJSON([]byte("\"INVALID\""))
			tt.vn.MarshalJSON()
		})
	}
}

func TestAuthFilter_String(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		af   model.AuthFilter
		want string
	}{
		{
			name: "test_AuthFilter_String",
			af:   model.AuthFilterCustomer,
			want: "AUTH_CUSTOMER",
		},
		{
			name: "test_AuthFilter_String_missing",
			af:   model.AuthFilter(100),
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.af.String() != tt.want {
				t.Errorf("String() = %v, want %v", tt.af.String(), tt.want)
			}
			tt.af.Keys()
			tt.af.UnmarshalJSON([]byte("\"AUTH_CUSTOMER\""))
			tt.af.UnmarshalJSON([]byte("INVALID\""))
			tt.af.UnmarshalJSON([]byte("\"INVALID\""))
			tt.af.MarshalJSON()
		})
	}
}

func TestSessionMethod_String(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		sm   model.SessionMethod
		want string
	}{
		{
			name: "test_SessionMethod_String",
			sm:   model.SessionMethodAuto,
			want: "SESSION_AUTO",
		},
		{
			name: "test_SessionMethod_String_missing",
			sm:   model.SessionMethod(100),
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.sm.String() != tt.want {
				t.Errorf("String() = %v, want %v", tt.sm.String(), tt.want)
			}
			tt.sm.UnmarshalJSON([]byte("\"SESSION_AUTO\""))
			tt.sm.UnmarshalJSON([]byte("INVALID\""))
			tt.sm.UnmarshalJSON([]byte("\"INVALID\""))
			tt.sm.MarshalJSON()
		})
	}
}
