package service

import (
	"fmt"
	"strings"
	"time"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	api "github.com/nervatura/nervatura/v6/pkg/api"
	cp "github.com/nervatura/nervatura/v6/pkg/component/client/component"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
	st "github.com/nervatura/nervatura/v6/pkg/static"
)

var searchQuery map[string]md.Query = map[string]md.Query{
	"transitem_simple": {
		Fields: []string{
			"t.*", "c.customer_name as customer_name", "COALESCE(i.amount, 0) as amount",
			"'trans' as editor", "t.id as editor_id", "'trans' as editor_view"},
		From: `trans_view t inner join customer c on t.customer_code = c.code  
		left join(select trans_code, sum(amount) as amount from item_view group by trans_code) i on t.code = i.trans_code`,
		Filter: fmt.Sprintf("trans_type in('%s','%s','%s','%s','%s','%s')",
			md.TransTypeInvoice.String(), md.TransTypeReceipt.String(), md.TransTypeOrder.String(),
			md.TransTypeOffer.String(), md.TransTypeWorksheet.String(), md.TransTypeRent.String()),
		OrderBy: []string{"t.id"},
		Limit:   st.BrowserRowLimit,
	},
	"transitem": {
		Fields: []string{
			"t.*", "c.customer_name", "COALESCE(i.amount, 0) as amount",
			"'trans' as editor", "t.id as editor_id", "'trans' as editor_view"},
		From: `trans_view t inner join customer c on t.customer_code = c.code  
		left join(select trans_code, sum(amount) as amount from item_view group by trans_code) i on t.code = i.trans_code`,
		Filter: fmt.Sprintf("trans_type in('%s','%s','%s','%s','%s','%s')",
			md.TransTypeInvoice.String(), md.TransTypeReceipt.String(), md.TransTypeOrder.String(),
			md.TransTypeOffer.String(), md.TransTypeWorksheet.String(), md.TransTypeRent.String()),
		OrderBy: []string{"t.id"},
		Limit:   st.BrowserRowLimit,
	},
	"transitem_map": {
		Fields: []string{"tm.id", "tm.code", "tm.trans_type", "tm.direction", "tm.trans_date",
			"tm.map_key as field_name",
			"COALESCE(cf.description, tm.map_key) as description",
			"tm.map_value as value", "COALESCE(cf.field_type, 'FIELD_STRING') as field_type",
			`case when cf.field_type = 'FIELD_BOOL' then 'bool'
				when cf.field_type = 'FIELD_INTEGER' then 'integer'
			when cf.field_type = 'FIELD_NUMBER' then 'float'
			when cf.field_type = 'FIELD_DATE' then 'date'
			when cf.field_type = 'FIELD_DATETIME' then 'datetime'
			when cf.field_type in (
				'FIELD_URL', 'FIELD_CUSTOMER','FIELD_EMPLOYEE','FIELD_PLACE','FIELD_PRODUCT','FIELD_PROJECT',
				'FIELD_TOOL', 'FIELD_TRANS_ITEM', 'FIELD_TRANS_MOVEMENT', 'FIELD_TRANS_PAYMENT') then 'link'
			else 'string' end as value_meta`,
			"tm.id as trans_id", "'trans' as editor", "'maps' as editor_view"},
		From:    "trans_map tm left join config_map cf on tm.map_key = cf.field_name",
		Filters: []md.Filter{},
		Filter: fmt.Sprintf("trans_type in('%s','%s','%s','%s','%s','%s')",
			md.TransTypeInvoice.String(), md.TransTypeReceipt.String(), md.TransTypeOrder.String(),
			md.TransTypeOffer.String(), md.TransTypeWorksheet.String(), md.TransTypeRent.String()),
		OrderBy: []string{"tm.id"},
		Limit:   st.BrowserRowLimit,
	},
	"transitem_item": {
		Fields: []string{"i.*", "i.code as item_code", "t.trans_date", "t.currency_code",
			"t.id as trans_id", "'trans' as editor", "'items' as editor_view"},
		From:    "item_view i inner join trans t on i.trans_code = t.code",
		OrderBy: []string{"i.id"},
		Limit:   st.BrowserRowLimit,
	},
	"customer_simple": {
		Fields: []string{
			"c.*", "'customer' as editor", "c.id as editor_id", "'customer' as editor_view"},
		From:    "customer_view c",
		OrderBy: []string{"c.id"},
		Limit:   st.BrowserRowLimit,
	},
	"customer": {
		Fields: []string{
			"c.*", "'customer' as editor", "c.id as editor_id", "'customer' as editor_view"},
		From:    "customer_view c",
		OrderBy: []string{"c.id"},
		Limit:   st.BrowserRowLimit,
	},
	"customer_map": {
		Fields: []string{"cm.id", "cm.code", "cm.customer_name",
			"cm.map_key as field_name",
			"COALESCE(cf.description, cm.map_key) as description",
			"cm.map_value as value", "COALESCE(cf.field_type, 'FIELD_STRING') as field_type",
			`case when cf.field_type = 'FIELD_BOOL' then 'bool'
				when cf.field_type = 'FIELD_INTEGER' then 'integer'
			when cf.field_type = 'FIELD_NUMBER' then 'float'
			when cf.field_type = 'FIELD_DATE' then 'date'
			when cf.field_type = 'FIELD_DATETIME' then 'datetime'
			when cf.field_type in (
				'FIELD_URL', 'FIELD_CUSTOMER','FIELD_EMPLOYEE','FIELD_PLACE','FIELD_PRODUCT','FIELD_PROJECT',
				'FIELD_TOOL', 'FIELD_TRANS_ITEM', 'FIELD_TRANS_MOVEMENT', 'FIELD_TRANS_PAYMENT') then 'link'
			else 'string' end as value_meta`,
			"cm.id as customer_id", "'customer' as editor", "'maps' as editor_view"},
		From:    "customer_map cm left join config_map cf on cm.map_key = cf.field_name",
		OrderBy: []string{"cm.id"},
		Limit:   st.BrowserRowLimit,
	},
	"customer_addresses": {
		Fields:  []string{"c.*", "c.id as customer_id", "'customer' as editor", "'addresses' as editor_view"},
		From:    "customer_addresses c",
		OrderBy: []string{"c.id"},
		Limit:   st.BrowserRowLimit,
	},
	"customer_contacts": {
		Fields:  []string{"c.*", "c.id as customer_id", "'customer' as editor", "'contacts' as editor_view"},
		From:    "customer_contacts c",
		OrderBy: []string{"c.id"},
		Limit:   st.BrowserRowLimit,
	},
	"customer_events": {
		Fields: []string{"c.*", "c.start_time as start_date", "c.end_time as end_date",
			"c.id as customer_id", "'customer' as editor", "'events' as editor_view"},
		From:    "customer_events c",
		OrderBy: []string{"c.id"},
		Limit:   st.BrowserRowLimit,
	},
	"product_simple": {
		Fields: []string{
			"p.*", "'product' as editor", "p.id as editor_id", "'product' as editor_view"},
		From:    "product_view p",
		OrderBy: []string{"p.id"},
		Limit:   st.BrowserRowLimit,
	},
	"product": {
		Fields: []string{
			"p.*", "'product' as editor", "p.id as editor_id", "'product' as editor_view"},
		From:    "product_view p",
		OrderBy: []string{"p.id"},
		Limit:   st.BrowserRowLimit,
	},
	"product_map": {
		Fields: []string{"pm.id", "pm.code", "pm.product_name",
			"pm.map_key as field_name",
			"COALESCE(cf.description, pm.map_key) as description",
			"pm.map_value as value", "COALESCE(cf.field_type, 'FIELD_STRING') as field_type",
			`case when cf.field_type = 'FIELD_BOOL' then 'bool'
				when cf.field_type = 'FIELD_INTEGER' then 'integer'
			when cf.field_type = 'FIELD_NUMBER' then 'float'
			when cf.field_type = 'FIELD_DATE' then 'date'
			when cf.field_type = 'FIELD_DATETIME' then 'datetime'
			when cf.field_type in (
				'FIELD_URL', 'FIELD_CUSTOMER','FIELD_EMPLOYEE','FIELD_PLACE','FIELD_PRODUCT','FIELD_PROJECT',
				'FIELD_TOOL', 'FIELD_TRANS_ITEM', 'FIELD_TRANS_MOVEMENT', 'FIELD_TRANS_PAYMENT') then 'link'
			else 'string' end as value_meta`,
			"pm.id as product_id", "'product' as editor", "'maps' as editor_view"},
		From:    "product_map pm left join config_map cf on pm.map_key = cf.field_name",
		OrderBy: []string{"pm.id"},
		Limit:   st.BrowserRowLimit,
	},
	"product_events": {
		Fields: []string{"p.*", "p.start_time as start_date", "p.end_time as end_date",
			"p.id as product_id", "'product' as editor", "'events' as editor_view"},
		From:    "product_events p",
		OrderBy: []string{"p.id"},
		Limit:   st.BrowserRowLimit,
	},
	"product_prices": {
		Fields: []string{"pr.*", "p.code as code", "pr.code as price_code", "p.product_name",
			"p.id as product_id", "'product' as editor", "'prices' as editor_view"},
		From:    "price_view pr inner join product p on pr.product_code = p.code",
		OrderBy: []string{"p.id"},
		Limit:   st.BrowserRowLimit,
	},
	"tool_simple": {
		Fields: []string{
			"t.*", "'tool' as editor", "t.id as editor_id", "'tool' as editor_view"},
		From:    "tool_view t",
		OrderBy: []string{"t.id"},
		Limit:   st.BrowserRowLimit,
	},
	"tool": {
		Fields: []string{
			"t.*", "'tool' as editor", "t.id as editor_id", "'tool' as editor_view"},
		From:    "tool_view t",
		OrderBy: []string{"t.id"},
		Limit:   st.BrowserRowLimit,
	},
	"tool_map": {
		Fields: []string{"tm.id", "tm.code", "tm.tool_description",
			"tm.map_key as field_name",
			"COALESCE(cf.description, tm.map_key) as description",
			"tm.map_value as value", "COALESCE(cf.field_type, 'FIELD_STRING') as field_type",
			`case when cf.field_type = 'FIELD_BOOL' then 'bool'
				when cf.field_type = 'FIELD_INTEGER' then 'integer'
			when cf.field_type = 'FIELD_NUMBER' then 'float'
			when cf.field_type = 'FIELD_DATE' then 'date'
			when cf.field_type = 'FIELD_DATETIME' then 'datetime'
			when cf.field_type in (
				'FIELD_URL', 'FIELD_CUSTOMER','FIELD_EMPLOYEE','FIELD_PLACE','FIELD_PRODUCT','FIELD_PROJECT',
				'FIELD_TOOL', 'FIELD_TRANS_ITEM', 'FIELD_TRANS_MOVEMENT', 'FIELD_TRANS_PAYMENT') then 'link'
			else 'string' end as value_meta`,
			"tm.id as tool_id", "'tool' as editor", "'maps' as editor_view"},
		From:    "tool_map tm left join config_map cf on tm.map_key = cf.field_name",
		OrderBy: []string{"tm.id"},
		Limit:   st.BrowserRowLimit,
	},
	"tool_events": {
		Fields: []string{"t.*", "t.start_time as start_date", "t.end_time as end_date",
			"t.id as tool_id", "'tool' as editor", "'events' as editor_view"},
		From:    "tool_events t",
		OrderBy: []string{"t.id"},
		Limit:   st.BrowserRowLimit,
	},
	"project_simple": {
		Fields: []string{
			"p.*", "'project' as editor", "p.id as editor_id", "'project' as editor_view"},
		From:    "project_view p",
		OrderBy: []string{"p.id"},
		Limit:   st.BrowserRowLimit,
	},
	"project": {
		Fields: []string{
			"p.*", "'project' as editor", "p.id as editor_id", "'project' as editor_view"},
		From:    "project_view p",
		OrderBy: []string{"p.id"},
		Limit:   st.BrowserRowLimit,
	},
	"project_map": {
		Fields: []string{"pm.id", "pm.code", "pm.project_name",
			"pm.map_key as field_name",
			"COALESCE(cf.description, pm.map_key) as description",
			"pm.map_value as value", "COALESCE(cf.field_type, 'FIELD_STRING') as field_type",
			`case when cf.field_type = 'FIELD_BOOL' then 'bool'
				when cf.field_type = 'FIELD_INTEGER' then 'integer'
			when cf.field_type = 'FIELD_NUMBER' then 'float'
			when cf.field_type = 'FIELD_DATE' then 'date'
			when cf.field_type = 'FIELD_DATETIME' then 'datetime'
			when cf.field_type in (
				'FIELD_URL', 'FIELD_CUSTOMER','FIELD_EMPLOYEE','FIELD_PLACE','FIELD_PRODUCT','FIELD_PROJECT',
				'FIELD_TOOL', 'FIELD_TRANS_ITEM', 'FIELD_TRANS_MOVEMENT', 'FIELD_TRANS_PAYMENT') then 'link'
			else 'string' end as value_meta`,
			"pm.id as project_id", "'project' as editor", "'maps' as editor_view"},
		From:    "project_map pm left join config_map cf on pm.map_key = cf.field_name",
		OrderBy: []string{"pm.id"},
		Limit:   st.BrowserRowLimit,
	},
	"project_addresses": {
		Fields:  []string{"p.*", "p.id as project_id", "'project' as editor", "'addresses' as editor_view"},
		From:    "project_addresses p",
		OrderBy: []string{"p.id"},
		Limit:   st.BrowserRowLimit,
	},
	"project_contacts": {
		Fields:  []string{"p.*", "p.id as project_id", "'project' as editor", "'contacts' as editor_view"},
		From:    "project_contacts p",
		OrderBy: []string{"p.id"},
		Limit:   st.BrowserRowLimit,
	},
	"project_events": {
		Fields: []string{"p.*", "p.start_time as start_date", "p.end_time as end_date",
			"p.id as project_id", "'project' as editor", "'events' as editor_view"},
		From:    "project_events p",
		OrderBy: []string{"p.id"},
		Limit:   st.BrowserRowLimit,
	},
	"employee_simple": {
		Fields: []string{
			"e.*", "'employee' as editor", "e.id as editor_id", "'employee' as editor_view"},
		From:    "employee_view e",
		OrderBy: []string{"e.id"},
		Limit:   st.BrowserRowLimit,
	},
	"employee": {
		Fields: []string{
			"e.*", "'employee' as editor", "e.id as editor_id", "'employee' as editor_view"},
		From:    "employee_view e",
		OrderBy: []string{"e.id"},
		Limit:   st.BrowserRowLimit,
	},
	"employee_map": {
		Fields: []string{"em.id", "em.code", "em.first_name", "em.surname",
			"em.map_key as field_name",
			"COALESCE(cf.description, em.map_key) as description",
			"em.map_value as value", "COALESCE(cf.field_type, 'FIELD_STRING') as field_type",
			`case when cf.field_type = 'FIELD_BOOL' then 'bool'
				when cf.field_type = 'FIELD_INTEGER' then 'integer'
			when cf.field_type = 'FIELD_NUMBER' then 'float'
			when cf.field_type = 'FIELD_DATE' then 'date'
			when cf.field_type = 'FIELD_DATETIME' then 'datetime'
			when cf.field_type in (
				'FIELD_URL', 'FIELD_CUSTOMER','FIELD_EMPLOYEE','FIELD_PLACE','FIELD_PRODUCT','FIELD_PROJECT',
				'FIELD_TOOL', 'FIELD_TRANS_ITEM', 'FIELD_TRANS_MOVEMENT', 'FIELD_TRANS_PAYMENT') then 'link'
			else 'string' end as value_meta`,
			"em.id as employee_id", "'employee' as editor", "'maps' as editor_view"},
		From:    "employee_map em left join config_map cf on em.map_key = cf.field_name",
		OrderBy: []string{"em.id"},
		Limit:   st.BrowserRowLimit,
	},
	"employee_events": {
		Fields: []string{"e.*", "e.start_time as start_date", "e.end_time as end_date",
			"e.id as employee_id", "'employee' as editor", "'events' as editor_view"},
		From:    "employee_events e",
		OrderBy: []string{"e.id"},
		Limit:   st.BrowserRowLimit,
	},
	"place_simple": {
		Fields: []string{
			"p.*", "'place' as editor", "p.id as editor_id", "'place' as editor_view"},
		From:    "place_view p",
		OrderBy: []string{"p.id"},
		Limit:   st.BrowserRowLimit,
	},
	"place": {
		Fields: []string{
			"p.*", "'place' as editor", "p.id as editor_id", "'place' as editor_view"},
		From:    "place_view p",
		OrderBy: []string{"p.id"},
		Limit:   st.BrowserRowLimit,
	},
	"place_map": {
		Fields: []string{"pm.id", "pm.code", "pm.place_name",
			"pm.map_key as field_name",
			"COALESCE(cf.description, pm.map_key) as description",
			"pm.map_value as value", "COALESCE(cf.field_type, 'FIELD_STRING') as field_type",
			`case when cf.field_type = 'FIELD_BOOL' then 'bool'
				when cf.field_type = 'FIELD_INTEGER' then 'integer'
			when cf.field_type = 'FIELD_NUMBER' then 'float'
			when cf.field_type = 'FIELD_DATE' then 'date'
			when cf.field_type = 'FIELD_DATETIME' then 'datetime'
			when cf.field_type in (
				'FIELD_URL', 'FIELD_CUSTOMER','FIELD_EMPLOYEE','FIELD_PLACE','FIELD_PRODUCT','FIELD_PROJECT',
				'FIELD_TOOL', 'FIELD_TRANS_ITEM', 'FIELD_TRANS_MOVEMENT', 'FIELD_TRANS_PAYMENT') then 'link'
			else 'string' end as value_meta`,
			"pm.id as place_id", "'place' as editor", "'maps' as editor_view"},
		From:    "place_map pm left join config_map cf on pm.map_key = cf.field_name",
		OrderBy: []string{"pm.id"},
		Limit:   st.BrowserRowLimit,
	},
	"place_contacts": {
		Fields:  []string{"p.*", "p.id as place_id", "'place' as editor", "'contacts' as editor_view"},
		From:    "place_contacts p",
		OrderBy: []string{"p.id"},
		Limit:   st.BrowserRowLimit,
	},
}

var compMap = cu.SM{
	"==": "=", "!=": "<>", "<": "<", "<=": "<=", ">": ">", ">=": ">=",
}

var compMapString = cu.SM{
	"==": "like", "!=": "not like",
}

var pre = func(or bool) string {
	if or {
		return "or"
	}
	return "and"
}

func searchFilterCustomer(view string, filter ct.BrowserFilter, queryFilters []string) []string {
	result := map[string]func() []string{
		"customer_simple": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
		"customer": func() []string {
			switch filter.Field {
			case "tax_free", "inactive":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %t)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToBoolean(filter.Value, false)))
			case "code", "customer_name", "customer_type", "tax_number", "account", "notes", "tag_lst":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
			case "id", "terms", "credit_limit", "discount":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %s)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToString(filter.Value, "")))
			default:
				return queryFilters
			}

		},
		"customer_map": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
		"customer_addresses": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
		"customer_contacts": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
		"customer_events": func() []string {
			switch filter.Field {
			case "start_time", "end_time":
				return append(queryFilters,
					fmt.Sprintf("%s (start_time %s '%s')", pre(filter.Or), compMap[filter.Comp], cu.ToString(filter.Value, "")))
			default:
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
			}
		},
	}
	return result[view]()
}

func searchFilterProduct(view string, filter ct.BrowserFilter, queryFilters []string) []string {
	result := map[string]func() []string{
		"product_simple": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
		"product": func() []string {
			switch filter.Field {
			case "inactive":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %t)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToBoolean(filter.Value, false)))
			case "code", "product_name", "product_type", "tax_code", "unit", "notes", "tag_lst", "barcode_type", "barcode":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
			case "id", "barcode_qty":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %s)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToString(filter.Value, "")))
			default:
				return queryFilters
			}
		},
		"product_map": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
		"product_events": func() []string {
			switch filter.Field {
			case "start_time", "end_time":
				return append(queryFilters,
					fmt.Sprintf("%s (start_time %s '%s')", pre(filter.Or), compMap[filter.Comp], cu.ToString(filter.Value, "")))
			default:
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
			}
		},
		"product_prices": func() []string {
			switch filter.Field {
			case "valid_from", "valid_to":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToString(filter.Value, "")))
			case "id", "qty", "price_value":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %s)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToString(filter.Value, "")))
			default:
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
			}
		},
	}
	return result[view]()
}

func searchFilterTool(view string, filter ct.BrowserFilter, queryFilters []string) []string {
	result := map[string]func() []string{
		"tool_simple": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
		"tool": func() []string {
			switch filter.Field {
			case "inactive":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %t)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToBoolean(filter.Value, false)))
			case "code", "description", "product_code", "notes", "tag_lst", "serial_number":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
			case "id":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %s)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToString(filter.Value, "")))
			default:
				return queryFilters
			}
		},
		"tool_map": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
		"tool_events": func() []string {
			switch filter.Field {
			case "start_time", "end_time":
				return append(queryFilters,
					fmt.Sprintf("%s (start_time %s '%s')", pre(filter.Or), compMap[filter.Comp], cu.ToString(filter.Value, "")))
			default:
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
			}
		},
	}
	return result[view]()
}

func searchFilterProject(view string, filter ct.BrowserFilter, queryFilters []string) []string {
	result := map[string]func() []string{
		"project_simple": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
		"project": func() []string {
			switch filter.Field {
			case "inactive":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %t)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToBoolean(filter.Value, false)))
			case "start_date", "end_date":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToString(filter.Value, "")))
			case "code", "project_name", "customer_code", "notes", "tag_lst":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
			case "id":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %s)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToString(filter.Value, "")))
			default:
				return queryFilters
			}
		},
		"project_map": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
		"project_addresses": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
		"project_contacts": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
		"project_events": func() []string {
			switch filter.Field {
			case "start_time", "end_time":
				return append(queryFilters,
					fmt.Sprintf("%s (start_time %s '%s')", pre(filter.Or), compMap[filter.Comp], cu.ToString(filter.Value, "")))
			default:
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
			}
		},
	}
	return result[view]()
}

func searchFilterEmployee(view string, filter ct.BrowserFilter, queryFilters []string) []string {
	result := map[string]func() []string{
		"employee_simple": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
		"employee": func() []string {
			switch filter.Field {
			case "inactive":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %t)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToBoolean(filter.Value, false)))
			case "start_date", "end_date":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToString(filter.Value, "")))
			case "code", "notes", "tag_lst",
				"first_name", "surname", "status", "email", "phone", "mobile",
				"street", "city", "state", "zip_code", "country":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
			case "id":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %s)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToString(filter.Value, "")))
			default:
				return queryFilters
			}
		},
		"employee_map": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
		"employee_events": func() []string {
			switch filter.Field {
			case "start_time", "end_time":
				return append(queryFilters,
					fmt.Sprintf("%s (start_time %s '%s')", pre(filter.Or), compMap[filter.Comp], cu.ToString(filter.Value, "")))
			default:
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
			}
		},
	}
	return result[view]()
}

func searchFilterPlace(view string, filter ct.BrowserFilter, queryFilters []string) []string {
	result := map[string]func() []string{
		"place_simple": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
		"place": func() []string {
			switch filter.Field {
			case "inactive":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %t)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToBoolean(filter.Value, false)))
			case "code", "place_name", "place_type", "currency_code", "notes", "tag_lst",
				"country", "state", "zip_code", "city", "street":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
			case "id":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %s)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToString(filter.Value, "")))
			default:
				return queryFilters
			}
		},
		"place_map": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
		"place_contacts": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
	}
	return result[view]()
}

func searchFilterTransItem(view string, filter ct.BrowserFilter, queryFilters []string) []string {
	result := map[string]func() []string{
		"transitem_simple": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
		"transitem": func() []string {
			switch filter.Field {
			case "tax_free", "paid", "closed":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %t)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToBoolean(filter.Value, false)))
			case "trans_date", "due_time":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToString(filter.Value, "")))
			case "id", "rate", "worksheet_distance", "worksheet_repair", "worksheet_total", "rent_holiday", "rent_bad_tool", "rent_other", "amount":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %s)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToString(filter.Value, "")))
			default:
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
			}
		},
		"transitem_map": func() []string {
			return append(queryFilters,
				fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
		},
		"transitem_item": func() []string {
			switch filter.Field {
			case "deposit", "action_price":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %t)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToBoolean(filter.Value, false)))
			case "id", "qty", "fx_price", "net_amount", "discount", "vat_amount", "amount", "own_stock":
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s %s)", pre(filter.Or), filter.Field, compMap[filter.Comp], cu.ToString(filter.Value, "")))
			default:
				return append(queryFilters,
					fmt.Sprintf("%s (%s %s '%s')", pre(filter.Or), filter.Field, compMapString[filter.Comp], "%"+cu.ToString(filter.Value, "")+"%"))
			}
		},
	}
	return result[view]()
}

func (cls *ClientService) searchFilter(view string, filter ct.BrowserFilter, queryFilters []string) []string {
	result := map[string]func() []string{
		"customer_simple": func() []string {
			return searchFilterCustomer(view, filter, queryFilters)
		},
		"customer": func() []string {
			return searchFilterCustomer(view, filter, queryFilters)
		},
		"customer_map": func() []string {
			return searchFilterCustomer(view, filter, queryFilters)
		},
		"customer_addresses": func() []string {
			return searchFilterCustomer(view, filter, queryFilters)
		},
		"customer_contacts": func() []string {
			return searchFilterCustomer(view, filter, queryFilters)
		},
		"customer_events": func() []string {
			return searchFilterCustomer(view, filter, queryFilters)
		},
		"product_simple": func() []string {
			return searchFilterProduct(view, filter, queryFilters)
		},
		"product": func() []string {
			return searchFilterProduct(view, filter, queryFilters)
		},
		"product_map": func() []string {
			return searchFilterProduct(view, filter, queryFilters)
		},
		"product_events": func() []string {
			return searchFilterProduct(view, filter, queryFilters)
		},
		"product_prices": func() []string {
			return searchFilterProduct(view, filter, queryFilters)
		},
		"tool_simple": func() []string {
			return searchFilterTool(view, filter, queryFilters)
		},
		"tool": func() []string {
			return searchFilterTool(view, filter, queryFilters)
		},
		"tool_map": func() []string {
			return searchFilterTool(view, filter, queryFilters)
		},
		"tool_events": func() []string {
			return searchFilterTool(view, filter, queryFilters)
		},
		"project_simple": func() []string {
			return searchFilterProject(view, filter, queryFilters)
		},
		"project": func() []string {
			return searchFilterProject(view, filter, queryFilters)
		},
		"project_map": func() []string {
			return searchFilterProject(view, filter, queryFilters)
		},
		"project_addresses": func() []string {
			return searchFilterProject(view, filter, queryFilters)
		},
		"project_contacts": func() []string {
			return searchFilterProject(view, filter, queryFilters)
		},
		"project_events": func() []string {
			return searchFilterProject(view, filter, queryFilters)
		},
		"employee_simple": func() []string {
			return searchFilterEmployee(view, filter, queryFilters)
		},
		"employee": func() []string {
			return searchFilterEmployee(view, filter, queryFilters)
		},
		"employee_map": func() []string {
			return searchFilterEmployee(view, filter, queryFilters)
		},
		"employee_events": func() []string {
			return searchFilterEmployee(view, filter, queryFilters)
		},
		"transitem_simple": func() []string {
			return searchFilterTransItem(view, filter, queryFilters)
		},
		"transitem": func() []string {
			return searchFilterTransItem(view, filter, queryFilters)
		},
		"transitem_map": func() []string {
			return searchFilterTransItem(view, filter, queryFilters)
		},
		"transitem_item": func() []string {
			return searchFilterTransItem(view, filter, queryFilters)
		},
		"place_simple": func() []string {
			return searchFilterPlace(view, filter, queryFilters)
		},
		"place": func() []string {
			return searchFilterPlace(view, filter, queryFilters)
		},
		"place_map": func() []string {
			return searchFilterPlace(view, filter, queryFilters)
		},
		"place_contacts": func() []string {
			return searchFilterPlace(view, filter, queryFilters)
		},
	}

	return result[view]()
}

func (cls *ClientService) searchData(ds *api.DataStore, view string, query md.Query, filters []ct.BrowserFilter) (result []cu.IM, err error) {
	queryFilters := []string{}

	for _, filter := range filters {
		queryFilters = cls.searchFilter(view, filter, queryFilters)
	}

	if len(queryFilters) > 0 {
		query.Filter = strings.Join(queryFilters, " ")
		query.Filter, _ = strings.CutPrefix(query.Filter, "or ")
		query.Filter, _ = strings.CutPrefix(query.Filter, "and ")
		query.Filter = "(" + query.Filter + ")"
		if len(query.Filters) > 0 {
			query.Filter = " and " + query.Filter
		}
	}

	return ds.StoreDataQuery(query, false)
}

func (cls *ClientService) searchResponseSideMenu(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, stateKey, stateData := client.GetStateData()
	switch strings.Split(cu.ToString(evt.Value, ""), "_")[0] {

	case "group":
		if cu.ToString(stateData["side_group"], "") == cu.ToString(evt.Value, "") {
			stateData["side_group"] = ""
		} else {
			stateData["side_group"] = evt.Value
		}
		client.SetSearch(stateKey, stateData, cu.ToBoolean(stateData["simple"], false))

	default:
		stateData["rows"] = []cu.IM{}
		stateData["filter_value"] = ""
		client.SetSearch(cu.ToString(evt.Value, ""), stateData, cp.SearchViewConfig(cu.ToString(evt.Value, ""), cu.SM{}).Simple)

	}
	return evt, err
}

func (cls *ClientService) searchResponse(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, stateKey, stateData := client.GetStateData()

	switch evt.Name {
	case ct.ClientEventSideMenu:
		return cls.searchResponseSideMenu(evt)

	case ct.BrowserEventBookmark:
		modal := cu.IM{
			"title":         client.Msg("bookmark_new"),
			"icon":          ct.IconStar,
			"label":         client.Msg("bookmark_enter"),
			"placeholder":   "",
			"field_name":    "value",
			"default_value": "",
			"required":      false,
			"next":          "bookmark_add",
		}
		client.SetForm("input_string", modal, 0, true)

	case ct.FormEventOK:
		sConf := cp.SearchViewConfig(stateKey, client.ClientLabels(client.Lang))
		visibleColumns := client.GetSearchVisibleColumns(ut.ToBoolMap(sConf.VisibleColumns, map[string]bool{}))
		viewData := cu.ToIM(stateData[stateKey], cu.IM{})
		frmValues := cu.ToIM(evt.Value, cu.IM{})
		frmValue := cu.ToIM(frmValues["value"], cu.IM{})
		label := cu.ToString(frmValue["value"], "")
		bookmark := md.Bookmark{
			BookmarkType: md.BookmarkTypeBrowser,
			Label:        label,
			Key:          stateKey,
			Code:         sConf.Title,
			Filters:      viewData["filters"],
			Columns:      visibleColumns,
			TimeStamp:    md.TimeDateTime{Time: time.Now()},
		}
		return cls.addBookmark(evt, bookmark), nil

	}
	return evt, err
}
