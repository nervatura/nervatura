//go:build grpc || all
// +build grpc all

package service

import (
	"context"
	"errors"
	"strconv"
	"strings"

	nt "github.com/nervatura/nervatura/service/pkg/nervatura"
	pb "github.com/nervatura/nervatura/service/pkg/proto"
	ut "github.com/nervatura/nervatura/service/pkg/utils"
)

// RPCService implements the Nervatura API service
type RPCService struct {
	Config        map[string]interface{}
	GetNervaStore func(database string) *nt.NervaStore
	GetTokenKeys  func() map[string]map[string]string
	pb.UnimplementedAPIServer
}

func (srv *RPCService) itemMap(key string, data nt.IM) *pb.ResponseGet_Value {
	metaMap := func(data interface{}) []*pb.MetaData {
		metadata := []*pb.MetaData{}
		if mdata, valid := data.([]nt.IM); valid {
			for i := 0; i < len(mdata); i++ {
				metadata = append(metadata, &pb.MetaData{
					Id:        ut.ToInteger(mdata[i]["id"], 0),
					Fieldname: ut.ToString(mdata[i]["fieldname"], ""),
					Fieldtype: ut.ToString(mdata[i]["fieldtype"], ""),
					Value:     ut.ToString(mdata[i]["value"], ""),
					Notes:     ut.ToString(mdata[i]["notes"], ""),
				})
			}
		}
		return metadata
	}

	itemMap := map[string]func(data nt.IM) *pb.ResponseGet_Value{
		"address": func(data nt.IM) *pb.ResponseGet_Value {
			return &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Address{Address: &pb.Address{
					Id:        ut.ToInteger(data["id"], 0),
					Nervatype: ut.ToInteger(data["nervatype"], 0),
					RefId:     ut.ToInteger(data["ref_id"], 0),
					Country:   ut.ToString(data["country"], ""),
					State:     ut.ToString(data["state"], ""),
					Zipcode:   ut.ToString(data["zipcode"], ""),
					City:      ut.ToString(data["city"], ""),
					Street:    ut.ToString(data["street"], ""),
					Notes:     ut.ToString(data["notes"], ""),
					Metadata:  metaMap(data["metadata"]),
				}}}
		},
		"barcode": func(data nt.IM) *pb.ResponseGet_Value {
			return &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Barcode{Barcode: &pb.Barcode{
					Id:          ut.ToInteger(data["id"], 0),
					Code:        ut.ToString(data["code"], ""),
					ProductId:   ut.ToInteger(data["product_id"], 0),
					Description: ut.ToString(data["description"], ""),
					Barcodetype: ut.ToInteger(data["barcodetype"], 0),
					Qty:         ut.ToFloat(data["qty"], 0),
					Defcode:     ut.ToBoolean(data["defcode"], false),
				}}}
		},
		"contact": func(data nt.IM) *pb.ResponseGet_Value {
			return &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Contact{Contact: &pb.Contact{
					Id:        ut.ToInteger(data["id"], 0),
					Nervatype: ut.ToInteger(data["nervatype"], 0),
					RefId:     ut.ToInteger(data["ref_id"], 0),
					Firstname: ut.ToString(data["firstname"], ""),
					Surname:   ut.ToString(data["surname"], ""),
					Status:    ut.ToString(data["status"], ""),
					Phone:     ut.ToString(data["phone"], ""),
					Fax:       ut.ToString(data["fax"], ""),
					Mobil:     ut.ToString(data["mobil"], ""),
					Email:     ut.ToString(data["email"], ""),
					Notes:     ut.ToString(data["notes"], ""),
					Metadata:  metaMap(data["metadata"]),
				}}}
		},
		"currency": func(data nt.IM) *pb.ResponseGet_Value {
			return &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Currency{Currency: &pb.Currency{
					Id:          ut.ToInteger(data["id"], 0),
					Curr:        ut.ToString(data["curr"], ""),
					Description: ut.ToString(data["description"], ""),
					Digit:       ut.ToInteger(data["digit"], 0),
					Defrate:     ut.ToFloat(data["defrate"], 0),
					Cround:      ut.ToInteger(data["cround"], 0),
					Metadata:    metaMap(data["metadata"]),
				}}}
		},
		"customer": func(data nt.IM) *pb.ResponseGet_Value {
			return &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Customer{Customer: &pb.Customer{
					Id:          ut.ToInteger(data["id"], 0),
					Custtype:    ut.ToInteger(data["custtype"], 0),
					Custnumber:  ut.ToString(data["custnumber"], ""),
					Custname:    ut.ToString(data["custname"], ""),
					Taxnumber:   ut.ToString(data["taxnumber"], ""),
					Account:     ut.ToString(data["account"], ""),
					Notax:       ut.ToBoolean(data["notax"], false),
					Terms:       ut.ToInteger(data["terms"], 0),
					Creditlimit: ut.ToFloat(data["creditlimit"], 0),
					Discount:    ut.ToFloat(data["discount"], 0),
					Notes:       ut.ToString(data["notes"], ""),
					Inactive:    ut.ToBoolean(data["inactive"], false),
					Metadata:    metaMap(data["metadata"]),
				}}}
		},
		"deffield": func(data nt.IM) *pb.ResponseGet_Value {
			return &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Deffield{Deffield: &pb.Deffield{
					Id:          ut.ToInteger(data["id"], 0),
					Fieldname:   ut.ToString(data["fieldname"], ""),
					Nervatype:   ut.ToInteger(data["nervatype"], 0),
					Subtype:     ut.ToIntPointer(data["subtype"], 0),
					Fieldtype:   ut.ToInteger(data["fieldtype"], 0),
					Description: ut.ToString(data["description"], ""),
					Valuelist:   ut.ToString(data["valuelist"], ""),
					Addnew:      ut.ToBoolean(data["addnew"], false),
					Visible:     ut.ToBoolean(data["visible"], false),
					Readonly:    ut.ToBoolean(data["readonly"], false),
				}}}
		},
		"employee": func(data nt.IM) *pb.ResponseGet_Value {
			return &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Employee{Employee: &pb.Employee{
					Id:              ut.ToInteger(data["id"], 0),
					Empnumber:       ut.ToString(data["empnumber"], ""),
					Username:        ut.ToStringPointer(data["username"], ""),
					Usergroup:       ut.ToInteger(data["Usergroup"], 0),
					Startdate:       ut.ToStringPointer(data["startdate"], ""),
					Enddate:         ut.ToStringPointer(data["enddate"], ""),
					Department:      ut.ToIntPointer(data["department"], 0),
					RegistrationKey: ut.ToString(data["registration_key"], ""),
					Inactive:        ut.ToBoolean(data["inactive"], false),
					Metadata:        metaMap(data["metadata"]),
				}}}
		},
		"event": func(data nt.IM) *pb.ResponseGet_Value {
			return &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Event{Event: &pb.Event{
					Id:          ut.ToInteger(data["id"], 0),
					Calnumber:   ut.ToString(data["calnumber"], ""),
					Nervatype:   ut.ToInteger(data["nervatype"], 0),
					RefId:       ut.ToInteger(data["ref_id"], 0),
					Uid:         ut.ToString(data["uid"], ""),
					Eventgroup:  ut.ToIntPointer(data["eventgroup"], 0),
					Fromdate:    ut.ToString(data["fromdate"], ""),
					Todate:      ut.ToStringPointer(data["todate"], ""),
					Subject:     ut.ToString(data["subject"], ""),
					Place:       ut.ToString(data["place"], ""),
					Description: ut.ToString(data["description"], ""),
					Metadata:    metaMap(data["metadata"]),
				}}}
		},
		"fieldvalue": func(data nt.IM) *pb.ResponseGet_Value {
			return &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Fieldvalue{Fieldvalue: &pb.Fieldvalue{
					Id:        ut.ToInteger(data["id"], 0),
					RefId:     ut.ToIntPointer(data["ref_id"], 0),
					Fieldname: ut.ToString(data["fieldname"], ""),
					Value:     ut.ToString(data["value"], ""),
					Notes:     ut.ToString(data["notes"], ""),
				}}}
		},
		"groups": func(data nt.IM) *pb.ResponseGet_Value {
			return &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Groups{Groups: &pb.Groups{
					Id:          ut.ToInteger(data["id"], 0),
					Groupname:   ut.ToString(data["groupname"], ""),
					Groupvalue:  ut.ToString(data["groupvalue"], ""),
					Description: ut.ToString(data["description"], ""),
					Inactive:    ut.ToBoolean(data["inactive"], false),
				}}}
		},
		"item": func(data nt.IM) *pb.ResponseGet_Value {
			return &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Item{Item: &pb.Item{
					Id:          ut.ToInteger(data["id"], 0),
					TransId:     ut.ToInteger(data["trans_id"], 0),
					ProductId:   ut.ToInteger(data["product_id"], 0),
					Unit:        ut.ToString(data["unit"], ""),
					Qty:         ut.ToFloat(data["qty"], 0),
					Fxprice:     ut.ToFloat(data["fxprice"], 0),
					Netamount:   ut.ToFloat(data["netamount"], 0),
					Discount:    ut.ToFloat(data["discount"], 0),
					TaxId:       ut.ToInteger(data["tax_id"], 0),
					Vatamount:   ut.ToFloat(data["vatamount"], 0),
					Amount:      ut.ToFloat(data["amount"], 0),
					Description: ut.ToString(data["description"], ""),
					Deposit:     ut.ToBoolean(data["deposit"], false),
					Ownstock:    ut.ToFloat(data["ownstock"], 0),
					Actionprice: ut.ToBoolean(data["actionprice"], false),
					Metadata:    metaMap(data["metadata"]),
				}}}
		},
		"link": func(data nt.IM) *pb.ResponseGet_Value {
			return &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Link{Link: &pb.Link{
					Id:          ut.ToInteger(data["id"], 0),
					Nervatype_1: ut.ToInteger(data["nervatype_1"], 0),
					RefId_1:     ut.ToInteger(data["ref_id_1"], 0),
					Nervatype_2: ut.ToInteger(data["nervatype_2"], 0),
					RefId_2:     ut.ToInteger(data["ref_id_2"], 0),
					Metadata:    metaMap(data["metadata"]),
				}}}
		},
		"log": func(data nt.IM) *pb.ResponseGet_Value {
			return &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Log{Log: &pb.Log{
					Id:         ut.ToInteger(data["id"], 0),
					Nervatype:  ut.ToIntPointer(data["nervatype"], 0),
					RefId:      ut.ToIntPointer(data["ref_id"], 0),
					EmployeeId: ut.ToInteger(data["employee_id"], 0),
					Crdate:     ut.ToString(data["crdate"], ""),
					Logstate:   ut.ToInteger(data["logstate"], 0),
					Metadata:   metaMap(data["metadata"]),
				}}}
		},
		"movement": func(data nt.IM) *pb.ResponseGet_Value {
			return &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Movement{Movement: &pb.Movement{
					Id:           ut.ToInteger(data["id"], 0),
					TransId:      ut.ToInteger(data["trans_id"], 0),
					Shippingdate: ut.ToString(data["shippingdate"], ""),
					Movetype:     ut.ToInteger(data["movetype"], 0),
					ProductId:    ut.ToIntPointer(data["product_id"], 0),
					ToolId:       ut.ToIntPointer(data["tool_id"], 0),
					PlaceId:      ut.ToIntPointer(data["place_id"], 0),
					Qty:          ut.ToFloat(data["qty"], 0),
					Description:  ut.ToString(data["description"], ""),
					Shared:       ut.ToBoolean(data["shared"], false),
					Metadata:     metaMap(data["metadata"]),
				}}}
		},
		"numberdef": func(data nt.IM) *pb.ResponseGet_Value {
			return &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Numberdef{Numberdef: &pb.Numberdef{
					Id:          ut.ToInteger(data["id"], 0),
					Numberkey:   ut.ToString(data["numberkey"], ""),
					Prefix:      ut.ToString(data["prefix"], ""),
					Curvalue:    ut.ToInteger(data["curvalue"], 0),
					Isyear:      ut.ToBoolean(data["isyear"], false),
					Sep:         ut.ToString(data["sep"], ""),
					Len:         ut.ToInteger(data["len"], 0),
					Description: ut.ToString(data["description"], ""),
					Visible:     ut.ToBoolean(data["visible"], false),
					Readonly:    ut.ToBoolean(data["readonly"], false),
					Orderby:     ut.ToInteger(data["orderby"], 0),
				}}}
		},
		"pattern": func(data nt.IM) *pb.ResponseGet_Value {
			return &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Pattern{Pattern: &pb.Pattern{
					Id:          ut.ToInteger(data["id"], 0),
					Description: ut.ToString(data["description"], ""),
					Transtype:   ut.ToInteger(data["transtype"], 0),
					Notes:       ut.ToString(data["notes"], ""),
					Defpattern:  ut.ToBoolean(data["defpattern"], false),
				}}}
		},
		"payment": func(data nt.IM) *pb.ResponseGet_Value {
			return &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Payment{Payment: &pb.Payment{
					Id:       ut.ToInteger(data["id"], 0),
					TransId:  ut.ToInteger(data["trans_id"], 0),
					Paiddate: ut.ToString(data["paiddate"], ""),
					Amount:   ut.ToFloat(data["amount"], 0),
					Notes:    ut.ToString(data["notes"], ""),
					Metadata: metaMap(data["metadata"]),
				}}}
		},
		"place": func(data nt.IM) *pb.ResponseGet_Value {
			return &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Place{Place: &pb.Place{
					Id:          ut.ToInteger(data["id"], 0),
					Planumber:   ut.ToString(data["planumber"], ""),
					Placetype:   ut.ToInteger(data["placetype"], 0),
					Description: ut.ToString(data["description"], ""),
					Curr:        ut.ToStringPointer(data["curr"], ""),
					Defplace:    ut.ToBoolean(data["defplace"], false),
					Notes:       ut.ToString(data["notes"], ""),
					Inactive:    ut.ToBoolean(data["inactive"], false),
					Metadata:    metaMap(data["metadata"]),
				}}}
		},
		"price": func(data nt.IM) *pb.ResponseGet_Value {
			return &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Price{Price: &pb.Price{
					Id:          ut.ToInteger(data["id"], 0),
					ProductId:   ut.ToInteger(data["product_id"], 0),
					Validfrom:   ut.ToString(data["validfrom"], ""),
					Validto:     ut.ToStringPointer(data["validto"], ""),
					Curr:        ut.ToString(data["curr"], ""),
					Qty:         ut.ToFloat(data["qty"], 0),
					Pricevalue:  ut.ToFloat(data["pricevalue"], 0),
					Vendorprice: ut.ToBoolean(data["vendorprice"], false),
					Metadata:    metaMap(data["metadata"]),
				}}}
		},
		"product": func(data nt.IM) *pb.ResponseGet_Value {
			return &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Product{Product: &pb.Product{
					Id:          ut.ToInteger(data["id"], 0),
					Partnumber:  ut.ToString(data["partnumber"], ""),
					Protype:     ut.ToInteger(data["protype"], 0),
					Description: ut.ToString(data["description"], ""),
					Unit:        ut.ToString(data["unit"], ""),
					TaxId:       ut.ToInteger(data["tax_id"], 0),
					Notes:       ut.ToString(data["notes"], ""),
					Webitem:     ut.ToBoolean(data["webitem"], false),
					Inactive:    ut.ToBoolean(data["inactive"], false),
					Metadata:    metaMap(data["metadata"]),
				}}}
		},
		"project": func(data nt.IM) *pb.ResponseGet_Value {
			return &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Project{Project: &pb.Project{
					Id:          ut.ToInteger(data["id"], 0),
					Pronumber:   ut.ToString(data["pronumber"], ""),
					Description: ut.ToString(data["description"], ""),
					CustomerId:  ut.ToIntPointer(data["customer_id"], 0),
					Startdate:   ut.ToStringPointer(data["startdate"], ""),
					Enddate:     ut.ToStringPointer(data["enddate"], ""),
					Notes:       ut.ToString(data["notes"], ""),
					Inactive:    ut.ToBoolean(data["inactive"], false),
					Metadata:    metaMap(data["metadata"]),
				}}}
		},
		"rate": func(data nt.IM) *pb.ResponseGet_Value {
			return &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Rate{Rate: &pb.Rate{
					Id:        ut.ToInteger(data["id"], 0),
					Ratetype:  ut.ToInteger(data["ratetype"], 0),
					Ratedate:  ut.ToString(data["ratedate"], ""),
					Curr:      ut.ToString(data["curr"], ""),
					PlaceId:   ut.ToIntPointer(data["place_id"], 0),
					Rategroup: ut.ToIntPointer(data["rategroup"], 0),
					Ratevalue: ut.ToFloat(data["ratevalue"], 0),
					Metadata:  metaMap(data["metadata"]),
				}}}
		},
		"tax": func(data nt.IM) *pb.ResponseGet_Value {
			return &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Tax{Tax: &pb.Tax{
					Id:          ut.ToInteger(data["id"], 0),
					Taxcode:     ut.ToString(data["taxcode"], ""),
					Description: ut.ToString(data["description"], ""),
					Rate:        ut.ToFloat(data["rate"], 0),
					Inactive:    ut.ToBoolean(data["inactive"], false),
					Metadata:    metaMap(data["metadata"]),
				}}}
		},
		"tool": func(data nt.IM) *pb.ResponseGet_Value {
			return &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Tool{Tool: &pb.Tool{
					Id:          ut.ToInteger(data["id"], 0),
					Serial:      ut.ToString(data["serial"], ""),
					Description: ut.ToString(data["description"], ""),
					ProductId:   ut.ToInteger(data["product_id"], 0),
					Toolgroup:   ut.ToIntPointer(data["toolgroup"], 0),
					Notes:       ut.ToString(data["notes"], ""),
					Inactive:    ut.ToBoolean(data["inactive"], false),
					Metadata:    metaMap(data["metadata"]),
				}}}
		},
		"trans": func(data nt.IM) *pb.ResponseGet_Value {
			return &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_Trans{Trans: &pb.Trans{
					Id:             ut.ToInteger(data["id"], 0),
					Transnumber:    ut.ToString(data["transnumber"], ""),
					Transtype:      ut.ToInteger(data["transtype"], 0),
					Direction:      ut.ToInteger(data["direction"], 0),
					RefTransnumber: ut.ToStringPointer(data["ref_transnumber"], ""),
					Crdate:         ut.ToString(data["crdate"], ""),
					Transdate:      ut.ToString(data["transdate"], ""),
					Duedate:        ut.ToStringPointer(data["duedate"], ""),
					CustomerId:     ut.ToIntPointer(data["customer_id"], 0),
					EmployeeId:     ut.ToIntPointer(data["employee_id"], 0),
					Department:     ut.ToIntPointer(data["department"], 0),
					ProjectId:      ut.ToIntPointer(data["project_id"], 0),
					PlaceId:        ut.ToIntPointer(data["place_id"], 0),
					Paidtype:       ut.ToIntPointer(data["paidtype"], 0),
					Curr:           ut.ToStringPointer(data["curr"], ""),
					Notax:          ut.ToBoolean(data["notax"], false),
					Paid:           ut.ToBoolean(data["paid"], false),
					Acrate:         ut.ToFloat(data["acrate"], 0),
					Notes:          ut.ToString(data["notes"], ""),
					Intnotes:       ut.ToString(data["intnotes"], ""),
					Fnote:          ut.ToString(data["fnote"], ""),
					Transtate:      ut.ToInteger(data["transtate"], 0),
					Closed:         ut.ToBoolean(data["closed"], false),
					Metadata:       metaMap(data["metadata"]),
				}}}
		},
		"ui_audit": func(data nt.IM) *pb.ResponseGet_Value {
			return &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_UiAudit{UiAudit: &pb.UiAudit{
					Id:          ut.ToInteger(data["id"], 0),
					Usergroup:   ut.ToInteger(data["usergroup"], 0),
					Nervatype:   ut.ToInteger(data["nervatype"], 0),
					Subtype:     ut.ToIntPointer(data["subtype"], 0),
					Inputfilter: ut.ToInteger(data["inputfilter"], 0),
					Supervisor:  ut.ToBoolean(data["supervisor"], false),
				}}}
		},
		"ui_menu": func(data nt.IM) *pb.ResponseGet_Value {
			return &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_UiMenu{UiMenu: &pb.UiMenu{
					Id:          ut.ToInteger(data["id"], 0),
					Menukey:     ut.ToString(data["menukey"], ""),
					Description: ut.ToString(data["description"], ""),
					Modul:       ut.ToString(data["modul"], ""),
					Icon:        ut.ToString(data["icon"], ""),
					Method:      ut.ToInteger(data["method"], 0),
					Funcname:    ut.ToString(data["funcname"], ""),
					Address:     ut.ToString(data["address"], ""),
				}}}
		},
		"ui_menufields": func(data nt.IM) *pb.ResponseGet_Value {
			return &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_UiMenufields{UiMenufields: &pb.UiMenufields{
					Id:          ut.ToInteger(data["id"], 0),
					MenuId:      ut.ToInteger(data["menu_id"], 0),
					Fieldname:   ut.ToString(data["fieldname"], ""),
					Description: ut.ToString(data["description"], ""),
					Fieldtype:   ut.ToInteger(data["fieldtype"], 0),
					Orderby:     ut.ToInteger(data["orderby"], 0),
				}}}
		},
		"ui_message": func(data nt.IM) *pb.ResponseGet_Value {
			return &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_UiMessage{UiMessage: &pb.UiMessage{
					Id:        ut.ToInteger(data["id"], 0),
					Secname:   ut.ToString(data["secname"], ""),
					Fieldname: ut.ToString(data["fieldname"], ""),
					Lang:      ut.ToString(data["lang"], ""),
					Msg:       ut.ToString(data["msg"], ""),
				}}}
		},
		"ui_printqueue": func(data nt.IM) *pb.ResponseGet_Value {
			return &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_UiPrintqueue{UiPrintqueue: &pb.UiPrintqueue{
					Id:         ut.ToInteger(data["id"], 0),
					Nervatype:  ut.ToIntPointer(data["nervatype"], 0),
					RefId:      ut.ToInteger(data["ref_id"], 0),
					Qty:        ut.ToFloat(data["qty"], 0),
					EmployeeId: ut.ToIntPointer(data["employee_id"], 0),
					ReportId:   ut.ToInteger(data["report_id"], 0),
					Crdate:     ut.ToString(data["crdate"], ""),
				}}}
		},
		"ui_report": func(data nt.IM) *pb.ResponseGet_Value {
			return &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_UiReport{UiReport: &pb.UiReport{
					Id:          ut.ToInteger(data["id"], 0),
					Reportkey:   ut.ToString(data["reportkey"], ""),
					Nervatype:   ut.ToInteger(data["nervatype"], 0),
					Transtype:   ut.ToIntPointer(data["transtype"], 0),
					Direction:   ut.ToIntPointer(data["direction"], 0),
					Repname:     ut.ToString(data["repname"], ""),
					Description: ut.ToString(data["description"], ""),
					Label:       ut.ToString(data["label"], ""),
					Filetype:    ut.ToInteger(data["filetype"], 0),
					Report:      ut.ToString(data["report"], ""),
				}}}
		},
		"ui_userconfig": func(data nt.IM) *pb.ResponseGet_Value {
			return &pb.ResponseGet_Value{
				Value: &pb.ResponseGet_Value_UiUserconfig{UiUserconfig: &pb.UiUserconfig{
					Id:         ut.ToInteger(data["id"], 0),
					EmployeeId: ut.ToIntPointer(data["employee_id"], 0),
					Section:    ut.ToStringPointer(data["section"], ""),
					Cfgroup:    ut.ToString(data["cfgroup"], ""),
					Cfname:     ut.ToString(data["cfname"], ""),
					Cfvalue:    ut.ToStringPointer(data["cfvalue"], ""),
					Orderby:    ut.ToInteger(data["orderby"], 0),
				}}}
		},
	}
	return itemMap[key](data)
}

func getValue(value interface{}) *pb.Value {
	if value == nil {
		return &pb.Value{Value: &pb.Value_Text{Text: "null"}}
	}
	if boolValue, valid := value.(bool); valid {
		return &pb.Value{Value: &pb.Value_Boolean{Boolean: boolValue}}
	}
	if intValue, valid := value.(int64); valid {
		return &pb.Value{Value: &pb.Value_Number{Number: ut.ToFloat(intValue, 0)}}
	}
	if floatValue, valid := value.(float64); valid {
		return &pb.Value{Value: &pb.Value_Number{Number: floatValue}}
	}
	return &pb.Value{Value: &pb.Value_Text{Text: ut.ToString(value, "")}}
}

func getIL(slist []string) []interface{} {
	ilist := make([]interface{}, len(slist))
	for i, v := range slist {
		ilist[i] = v
	}
	return ilist
}

func getIValue(value *pb.Value) interface{} {
	switch v := value.Value.(type) {
	case *pb.Value_Boolean:
		return v.Boolean
	case *pb.Value_Number:
		return v.Number
	case *pb.Value_Text:
		if v.Text == "null" || v.Text == "" {
			return nil
		}
		if strings.HasPrefix(v.Text, "numberdef,") {
			return getIL(strings.Split(v.Text, ","))
		}
		return v.Text
	}
	return nil
}

func (srv *RPCService) rowMap(values interface{}) *pb.ResponseRows {
	encodeMap := func(values interface{}) *pb.ResponseRows_Item {
		row := &pb.ResponseRows_Item{
			Values: make(map[string]*pb.Value),
		}

		switch v := values.(type) {
		case nt.SM:
			for fieldName, sValue := range v {
				row.Values[fieldName] = getValue(sValue)
			}
		case nt.IM:
			for fieldName, iValue := range v {
				row.Values[fieldName] = getValue(iValue)
			}
		}
		return row
	}

	rows := &pb.ResponseRows{}
	if imap, valid := values.([]nt.IM); valid {
		for index := 0; index < len(imap); index++ {
			rows.Items = append(rows.Items, encodeMap(imap[index]))
		}
	}
	if smap, valid := values.([]nt.SM); valid {
		for index := 0; index < len(smap); index++ {
			rows.Items = append(rows.Items, encodeMap(smap[index]))
		}
	}
	return rows
}

func (srv *RPCService) fieldsToIMap(values map[string]*pb.Value) nt.IM {
	iMap := make(nt.IM)
	for fieldname, value := range values {
		iMap[fieldname] = getIValue(value)
	}
	return iMap
}

func (srv *RPCService) TokenAuth(authorization []string, parent context.Context) (ctx context.Context, err error) {
	if len(authorization) < 1 {
		return ctx, errors.New(ut.GetMessage("error_unauthorized"))
	}
	tokenStr := strings.TrimPrefix(authorization[0], "Bearer ")
	if tokenStr == "" {
		return ctx, errors.New(ut.GetMessage("error_unauthorized"))
	}
	claim, err := ut.TokenDecode(tokenStr)
	if err != nil {
		return ctx, err
	}
	tokenCtx := context.WithValue(parent, TokenCtxKey, tokenStr)

	database := ""
	if _, found := claim["database"]; found {
		database = claim["database"].(string)
	}
	nstore := srv.GetNervaStore(database)
	err = (&nt.API{NStore: nstore}).TokenLogin(nt.IM{"token": tokenStr, "keys": srv.GetTokenKeys()})
	if err != nil {
		return ctx, err
	}
	ctx = context.WithValue(tokenCtx, NstoreCtxKey, &nt.API{NStore: nstore})
	return ctx, nil
}

func (srv *RPCService) ApiKeyAuth(authorization []string, parent context.Context) (ctx context.Context, err error) {
	if len(authorization) < 1 {
		return ctx, errors.New(ut.GetMessage("error_unauthorized"))
	}
	apiKey := strings.Trim(authorization[0], " ")
	if apiKey == "" {
		return ctx, errors.New(ut.GetMessage("error_unauthorized"))
	}
	if srv.Config["NT_API_KEY"] != apiKey {
		return ctx, errors.New(ut.GetMessage("error_unauthorized"))
	}
	nstore := srv.GetNervaStore("")
	ctx = context.WithValue(parent, NstoreCtxKey, &nt.API{NStore: nstore})
	return ctx, nil
}

// UserLogin - Logs in user by username and password
func (srv *RPCService) UserLogin(ctx context.Context, req *pb.RequestUserLogin) (res *pb.ResponseUserLogin, err error) {
	if req.Database == "" {
		return res, errors.New(ut.GetMessage("missing_database"))
	}
	nstore := srv.GetNervaStore(req.Database)
	login := nt.IM{"username": req.Username, "password": req.Password, "database": req.Database}
	token, engine, err := (&nt.API{NStore: nstore}).UserLogin(login)
	return &pb.ResponseUserLogin{Token: token, Engine: engine, Version: srv.Config["version"].(string)}, err
}

func (srv *RPCService) getApi(ctx context.Context, admin bool) (*nt.API, error) {
	api := ctx.Value(NstoreCtxKey).(*nt.API)
	if api.NStore.User == nil {
		return api, errors.New(ut.GetMessage("error_unauthorized"))
	}
	if admin && api.NStore.User.Scope != "admin" {
		return api, errors.New(ut.GetMessage("error_unauthorized"))
	}
	return api, nil
}

// User (employee or customer) password change.
func (srv *RPCService) UserPassword(ctx context.Context, req *pb.RequestUserPassword) (res *pb.ResponseEmpty, err error) {
	api, err := srv.getApi(ctx, false)
	if err != nil || ((req.Username != api.NStore.User.Username && api.NStore.User.Scope != "admin") || (req.Custnumber != "" && api.NStore.User.Scope != "admin")) {
		return res, errors.New(ut.GetMessage("error_unauthorized"))
	}
	options := nt.IM{"username": req.Username, "custnumber": req.Custnumber,
		"password": req.Password, "confirm": req.Confirm}
	if req.Username == "" && req.Custnumber == "" {
		options["username"] = api.NStore.User.Username
	}
	err = api.UserPassword(options)
	return &pb.ResponseEmpty{}, err
}

// TokenDecode - decoded JWT token but doesn't validate the signature.
func (srv *RPCService) TokenDecode(ctx context.Context, req *pb.RequestTokenDecode) (*pb.ResponseTokenDecode, error) {
	mClaims, err := ut.TokenDecode(req.Value)
	if err != nil {
		return nil, err
	}
	claims := &pb.ResponseTokenDecode{
		Username: mClaims["username"].(string), Database: mClaims["database"].(string),
		Exp: mClaims["exp"].(float64), Iss: mClaims["iss"].(string),
	}
	return claims, err
}

// TokenLogin - JWT token auth.
func (srv *RPCService) TokenLogin(ctx context.Context, req *pb.RequestEmpty) (*pb.ResponseTokenLogin, error) {
	api := ctx.Value(NstoreCtxKey).(*nt.API)
	if api.NStore.User != nil {
		return &pb.ResponseTokenLogin{
			Id: api.NStore.User.Id, Username: api.NStore.User.Username, Empnumber: api.NStore.User.Empnumber,
			Usergroup: api.NStore.User.Usergroup, Scope: api.NStore.User.Scope, Department: api.NStore.User.Department,
		}, nil
	}
	return &pb.ResponseTokenLogin{
		Id: ut.ToInteger(api.NStore.Customer["id"], 0), Username: ut.ToString(api.NStore.Customer["custnumber"], ""),
	}, nil
}

// TokenRefresh - Refreshes JWT token by checking at database whether refresh token exists.
func (srv *RPCService) TokenRefresh(ctx context.Context, req *pb.RequestEmpty) (res *pb.ResponseTokenRefresh, err error) {
	if ctx.Value(NstoreCtxKey) == nil {
		return res, errors.New(ut.GetMessage("error_unauthorized"))
	}
	api := ctx.Value(NstoreCtxKey).(*nt.API)
	token, err := api.TokenRefresh()
	return &pb.ResponseTokenRefresh{Value: token}, err
}

// DatabaseCreate - Create a new Nervatura database
func (srv *RPCService) DatabaseCreate(ctx context.Context, req *pb.RequestDatabaseCreate) (log *pb.ResponseDatabaseCreate, err error) {
	options := nt.IM{
		"database": req.Alias, "demo": strconv.FormatBool(req.Demo),
	}
	api := ctx.Value(NstoreCtxKey).(*nt.API)
	results, err := api.DatabaseCreate(options)
	if err != nil {
		return log, err
	}
	log = &pb.ResponseDatabaseCreate{Details: srv.rowMap(results)}
	return log, err
}

// Get - returns one or more records
func (srv *RPCService) Get(ctx context.Context, req *pb.RequestGet) (res *pb.ResponseGet, err error) {
	api, err := srv.getApi(ctx, false)
	if err != nil {
		return res, err
	}
	res = &pb.ResponseGet{
		Values: []*pb.ResponseGet_Value{},
	}
	options := nt.IM{
		"nervatype": pb.DataType_name[int32(req.Nervatype)], "metadata": req.Metadata,
	}
	if len(req.Ids) > 0 {
		ids := []string{}
		for i := 0; i < len(req.Ids); i++ {
			ids = append(ids, strconv.FormatInt(req.Ids[i], 10))
		}
		options["ids"] = strings.Join(ids, ",")
	} else if len(req.Filter) > 0 {
		options["filter"] = strings.Join(req.Filter, "|")
	}
	results, err := api.Get(options)

	for i := 0; i < len(results); i++ {
		res.Values = append(res.Values, srv.itemMap(pb.DataType_name[int32(req.Nervatype)], results[i]))
	}

	return res, err
}

// Add/update one or more items
func (srv *RPCService) Update(ctx context.Context, req *pb.RequestUpdate) (res *pb.ResponseUpdate, err error) {
	api, err := srv.getApi(ctx, false)
	if err != nil {
		return res, err
	}
	res = &pb.ResponseUpdate{}
	options := []nt.IM{}
	for i := 0; i < len(req.Items); i++ {
		item := srv.fieldsToIMap(req.Items[i].Values)
		item["keys"] = srv.fieldsToIMap(req.Items[i].Keys)
		options = append(options, item)
	}
	res.Values, err = api.Update(pb.DataType_name[int32(req.Nervatype)], options)
	return res, err
}

// Delete - delete a record
func (srv *RPCService) Delete(ctx context.Context, req *pb.RequestDelete) (res *pb.ResponseEmpty, err error) {
	api, err := srv.getApi(ctx, false)
	if err != nil {
		return res, err
	}
	options := nt.IM{"nervatype": pb.DataType_name[int32(req.Nervatype)], "id": int(req.Id), "key": req.Key}
	err = api.Delete(options)
	return &pb.ResponseEmpty{}, err
}

// Run raw SQL queries in safe mode
func (srv *RPCService) View(ctx context.Context, req *pb.RequestView) (res *pb.ResponseView, err error) {
	api, err := srv.getApi(ctx, false)
	if err != nil {
		return res, err
	}
	res = &pb.ResponseView{
		Values: make(map[string]*pb.ResponseRows),
	}
	options := []nt.IM{}
	for i := 0; i < len(req.Options); i++ {
		values := []interface{}{}
		for vi := 0; vi < len(req.Options[i].Values); vi++ {
			values = append(values, getIValue(req.Options[i].Values[vi]))
		}
		prm := nt.IM{
			"key":    req.Options[i].Key,
			"text":   req.Options[i].Text,
			"values": values,
		}
		options = append(options, prm)
	}
	results, err := api.View(options)
	if err == nil {
		for fieldname, values := range results {
			res.Values[fieldname] = srv.rowMap(values)
		}
	}
	return res, err
}

// Call a server-side function
func (srv *RPCService) Function(ctx context.Context, req *pb.RequestFunction) (res *pb.ResponseFunction, err error) {
	api, err := srv.getApi(ctx, false)
	if err != nil {
		return res, err
	}
	res = &pb.ResponseFunction{}
	options := nt.IM{
		"key":    req.Key,
		"values": srv.fieldsToIMap(req.Values),
	}
	if req.Value != nil {
		var values interface{}
		err = ut.ConvertFromByte(req.Value, &values)
		if err == nil {
			options["values"] = values
		}
	}
	result, err := api.Function(options)
	if err != nil {
		return res, err
	}
	res.Value, err = ut.ConvertToByte(result)
	return res, err
}

// List all available Nervatura Report.
func (srv *RPCService) ReportList(ctx context.Context, req *pb.RequestReportList) (res *pb.ResponseReportList, err error) {
	api, err := srv.getApi(ctx, true)
	if err != nil {
		return res, err
	}
	res = &pb.ResponseReportList{
		Items: []*pb.ResponseReportList_Info{},
	}
	options := nt.IM{
		"label": req.Label,
	}
	results, err := api.ReportList(options)
	if err != nil {
		return res, err
	}
	for i := 0; i < len(results); i++ {
		res.Items = append(res.Items, &pb.ResponseReportList_Info{
			Reportkey:   results[i]["reportkey"].(string),
			Repname:     results[i]["repname"].(string),
			Description: results[i]["description"].(string),
			Label:       results[i]["label"].(string),
			Reptype:     results[i]["reptype"].(string),
			Filename:    results[i]["filename"].(string),
			Installed:   results[i]["installed"].(bool),
		})
	}
	return res, err
}

// Install a report to the database.
func (srv *RPCService) ReportInstall(ctx context.Context, req *pb.RequestReportInstall) (res *pb.ResponseReportInstall, err error) {
	api, err := srv.getApi(ctx, true)
	if err != nil {
		return res, err
	}
	res = &pb.ResponseReportInstall{}
	options := nt.IM{
		"reportkey": req.Reportkey,
	}
	res.Id, err = api.ReportInstall(options)
	if err != nil {
		return res, err
	}
	return res, err
}

// Delete a report from the database.
func (srv *RPCService) ReportDelete(ctx context.Context, req *pb.RequestReportDelete) (res *pb.ResponseEmpty, err error) {
	api, err := srv.getApi(ctx, true)
	if err != nil {
		return res, err
	}
	res = &pb.ResponseEmpty{}
	options := nt.IM{
		"reportkey": req.Reportkey,
	}
	err = api.ReportDelete(options)
	return res, err
}

func (srv *RPCService) Report(ctx context.Context, req *pb.RequestReport) (res *pb.ResponseReport, err error) {
	api := ctx.Value(NstoreCtxKey).(*nt.API)
	res = &pb.ResponseReport{}
	options := nt.IM{
		"reportkey":   req.Reportkey,
		"orientation": pb.ReportOrientation_name[int32(req.Orientation)],
		"size":        pb.ReportSize_name[int32(req.Size)],
		"output":      pb.ReportOutput_name[int32(req.Output)],
		"refnumber":   req.Refnumber,
		"template":    req.Template,
		"filters":     srv.fieldsToIMap(req.Filters),
	}
	if req.Type != pb.ReportType_report_none {
		options["nervatype"] = strings.TrimPrefix(pb.ReportType_name[int32(req.Type)], "report_")
	}
	results, err := api.Report(options)
	if err == nil {
		res.Value, err = ut.ConvertToByte(results)
	}
	return res, err
}
