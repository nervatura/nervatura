package nervatura

import (
	"errors"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/alexedwards/argon2id"
	ver "github.com/mcuadros/go-version"
	ut "github.com/nervatura/nervatura/service/pkg/utils"
)

/*
API - Application Programming Interface

See more docs and examples: https://nervatura.github.io/nervatura/api
*/
type API struct {
	NStore *NervaStore
}

func (api *API) getHashvalue(refname string) (string, error) {
	hashtable := ut.ToString(api.NStore.config["NT_HASHTABLE"], "")
	err := api.NStore.ds.CheckHashtable(hashtable)
	if err != nil {
		return "", err
	}
	query := []Query{{
		Fields: []string{"*"}, From: hashtable, Filters: []Filter{
			{Field: "refname", Comp: "==", Value: refname},
		}}}
	rows, err := api.NStore.ds.Query(query, nil)
	if err != nil {
		return "", err
	}
	if len(rows) > 0 {
		return rows[0]["value"].(string), nil
	}
	return "", nil
}

func (api *API) authUser(options IM) error {

	if !api.NStore.ds.Connection().Connected {
		database := ut.ToString(options["database"], "")
		if database == "" {
			return errors.New(ut.GetMessage("missing_database"))
		}
		alias := ut.ToString(api.NStore.config["NT_ALIAS_"+strings.ToUpper(database)], os.Getenv("NT_ALIAS_"+strings.ToUpper(database)))
		if alias == "" {
			return errors.New(ut.GetMessage("missing_database"))
		}
		err := api.NStore.ds.CreateConnection(database, alias)
		if err != nil {
			return err
		}
	}

	if _, found := options["username"]; !found {
		return errors.New(ut.GetMessage("missing_user"))
	}

	rows, err := api.NStore.ds.QueryKey(IM{"qkey": "user", "username": options["username"]}, nil)
	if err != nil {
		return err
	}
	if len(rows) > 0 {
		api.NStore.User = &User{
			Id:         ut.ToInteger(rows[0]["id"], 0),
			Username:   ut.ToString(rows[0]["username"], ""),
			Empnumber:  ut.ToString(rows[0]["empnumber"], ""),
			Usergroup:  ut.ToInteger(rows[0]["usergroup"], 0),
			Scope:      ut.ToString(rows[0]["scope"], ""),
			Department: ut.ToString(rows[0]["department"], ""),
		}
	} else {
		query := []Query{{
			Fields: []string{"*"}, From: "customer", Filters: []Filter{
				{Field: "inactive", Comp: "==", Value: 0},
				{Field: "deleted", Comp: "==", Value: 0},
				{Field: "custnumber", Comp: "==", Value: options["username"]}}}}
		rows, err := api.NStore.ds.Query(query, nil)
		if err != nil {
			return err
		}
		if len(rows) > 0 {
			api.NStore.Customer = rows[0]
		} else {
			return errors.New(ut.GetMessage("unknown_user"))
		}
	}
	return nil
}

/*
checkVersion - database version check
*/
func (api *API) checkVersion() (err error) {
	version := ut.ToString(api.NStore.config["version"], "")

	dbsUpdate := func(dbVersion string) error {
		verUpdate := map[string]func() error{
			"": func() error {
				_, err = api.NStore.ds.Update(Update{Model: "fieldvalue", Values: IM{
					"fieldname": "version", "value": version,
				}})
				return err
			},
			"5.0.5": func() error {
				_, err = api.NStore.ds.QuerySQL(`delete from fieldvalue where id not in(
					select min(id) from fieldvalue group by fieldname, ref_id) and fieldname 
					in('trans_transcast','trans_custinvoice_compname','trans_custinvoice_comptax',
					'trans_custinvoice_compaddress','trans_custinvoice_custname','trans_custinvoice_custtax',
					'trans_custinvoice_custaddress')`, []interface{}{}, nil)
				return err
			},
		}
		return verUpdate[dbVersion]()
	}

	query := []Query{{
		Fields: []string{"value"},
		From:   "fieldvalue", Filters: []Filter{
			{Field: "ref_id", Comp: "is", Value: "null"},
			{Field: "fieldname", Comp: "==", Value: "version"}}}}
	result, err := api.NStore.ds.Query(query, nil)
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return dbsUpdate("")
	}
	if ut.ToString(result[0]["value"], "") == "dev" {
		return nil
	}
	if ver.Compare(ut.ToString(result[0]["value"], ""), "5.0.5", "<") {
		return dbsUpdate("5.0.5")
	}
	return nil
}

/*
TokenRefresh - create/refresh a JWT token
*/
func (api *API) TokenRefresh() (string, error) {
	conn := api.NStore.ds.Connection()
	if !conn.Connected {
		return "", errors.New(ut.GetMessage("not_connect"))
	}
	username := ut.ToString(api.NStore.Customer["custnumber"], api.NStore.User.Username)
	return ut.CreateToken(username, conn.Alias, api.NStore.config)
}

/*
TokenLogin - database JWT token auth.

Example:

  options := map[string]interface{}{"token": "JWT_token"}
  err := getAPI().TokenLogin(options)

*/
func (api *API) TokenLogin(options IM) error {
	tokenString := ut.ToString(options["token"], "")
	if tokenString == "" {
		return errors.New(ut.GetMessage("missing_required_field") + ": token")
	}
	keyMap := make(map[string]map[string]string)
	if keys, found := options["keys"].(map[string]map[string]string); found {
		keyMap = keys
	}
	data, err := ut.ParseToken(tokenString, keyMap, api.NStore.config)
	if err != nil {
		return err
	}
	return api.authUser(data)
}

/*
UserPassword - set/change a user password

Example:

  options = map[string]interface{}{
    "username": "demo",
    "password": "321",
    "confirm": "321"}
  err = api.UserPassword(options)

*/
func (api *API) UserPassword(options IM) error {
	if _, found := options["username"]; !found {
		if _, found := options["custnumber"]; !found {
			return errors.New(ut.GetMessage("missing_required_field") + ": username or custnumber")
		}
	}
	if _, found := options["password"]; !found {
		return errors.New(ut.GetMessage("missing_required_field") + ": password")
	}
	if _, found := options["confirm"]; !found {
		return errors.New(ut.GetMessage("missing_required_field") + ": confirm")
	}
	if options["password"] == "" {
		return errors.New(ut.GetMessage("empty_password"))
	}
	if options["password"] != options["confirm"] {
		return errors.New(ut.GetMessage("verify_password"))
	}
	if !api.NStore.ds.Connection().Connected {
		return errors.New(ut.GetMessage("not_connect"))
	}
	refname := ""
	if _, found := options["custnumber"]; found && api.NStore.Customer != nil {
		if options["custnumber"] == api.NStore.Customer["custnumber"] {
			refname = "customer" + strconv.FormatInt(api.NStore.Customer["id"].(int64), 10)
		}
	} else if _, found := options["username"]; found && api.NStore.User != nil {
		if options["username"] == api.NStore.User.Username {
			refname = "employee" + strconv.FormatInt(api.NStore.User.Id, 10)
		}
	}
	if refname == "" {
		var query []Query
		if _, found := options["custnumber"]; found && options["custnumber"] != "" {
			query = []Query{{
				Fields: []string{"*"}, From: "customer", Filters: []Filter{
					{Field: "inactive", Comp: "==", Value: 0},
					{Field: "deleted", Comp: "==", Value: 0},
					{Field: "custnumber", Comp: "==", Value: options["custnumber"]},
				}}}
			refname = "customer"
		} else {
			query = []Query{{
				Fields: []string{"*"}, From: "employee", Filters: []Filter{
					{Field: "deleted", Comp: "==", Value: 0},
					{Field: "username", Comp: "==", Value: options["username"]},
				}}}
			refname = "employee"
		}
		rows, err := api.NStore.ds.Query(query, nil)
		if err != nil {
			return err
		}
		if len(rows) > 0 {
			refname += ut.ToString(rows[0]["id"], "")
		} else {
			return errors.New(ut.GetMessage("unknown_user"))
		}
	}
	refname = ut.GetHash(refname)
	hash, err := argon2id.CreateHash(ut.ToString(options["password"], ""), argon2id.DefaultParams)
	if err != nil {
		return err
	}
	return api.NStore.ds.UpdateHashtable(ut.ToString(api.NStore.config["NT_HASHTABLE"], ""), refname, hash)
}

/*
UserLogin - database user login

Returns a access token and the type of database.

  options := map[string]interface{}{
    "database": "alias_name",
    "username": "username",
    "password": "password"}
  token, engine, err := getAPI().UserLogin(options)

*/
func (api *API) UserLogin(options IM) (string, string, error) {
	if !ut.ToBoolean(api.NStore.config["NT_PASSWORD_LOGIN"], true) {
		return "", "", errors.New(ut.GetMessage("password_login_disabled"))
	}
	if _, found := options["database"]; !found {
		if ut.ToString(api.NStore.config["NT_ALIAS_DEFAULT"], "") == "" {
			return "", "", errors.New(ut.GetMessage("missing_database"))
		}
		options["database"] = strings.ToLower(ut.ToString(api.NStore.config["NT_ALIAS_DEFAULT"], ""))
	}
	password := ut.ToString(options["password"], "")
	if err := api.authUser(options); err != nil {
		return "", "", err
	}
	if err := api.checkVersion(); err != nil {
		return "", "", err
	}

	refname := "employee" + strconv.FormatInt(api.NStore.User.Id, 10)
	if api.NStore.Customer != nil {
		refname = "customer" + strconv.FormatInt(api.NStore.Customer["id"].(int64), 10)
	}
	refname = ut.GetHash(refname)
	hash, err := api.getHashvalue(refname)
	if err != nil {
		return "", "", err
	}

	if password != "" && hash != "" {
		match, err := argon2id.ComparePasswordAndHash(password, hash)
		if err != nil {
			return "", "", err
		}
		if !match {
			return "", "", errors.New(ut.GetMessage("wrong_password"))
		}
	} else if password != hash {
		return "", "", errors.New(ut.GetMessage("wrong_password"))
	}

	token, err := api.TokenRefresh()
	return token, api.NStore.ds.Connection().Engine, err
}

/*
DatabaseCreate - create a Nervatura Database

All data will be destroyed!

Example:

  options := map[string]interface{}{
    "database": alias,
    "demo": true,
    "report_dir": ""}
  _, err := getAPI().DatabaseCreate(options)

*/
func (api *API) DatabaseCreate(options IM) ([]SM, error) {
	logData := []SM{}
	database := ut.ToString(options["database"], "")
	if database == "" {
		return logData, errors.New(ut.GetMessage("missing_required_field") + ": database")
	}

	//check connect
	alias := ut.ToString(api.NStore.config["NT_ALIAS_"+strings.ToUpper(database)], os.Getenv("NT_ALIAS_"+strings.ToUpper(database)))
	if err := api.NStore.ds.CreateConnection(database, alias); err != nil {
		logData = append(logData, SM{
			"stamp":   time.Now().Format(TimeLayout),
			"state":   "err",
			"message": ut.GetMessage("not_connect")})
		return logData, errors.New(ut.GetMessage("not_connect"))
	}

	logData, err := api.NStore.ds.CreateDatabase(logData)
	if err != nil {
		logData = append(logData, SM{
			"stamp":   time.Now().Format(TimeLayout),
			"state":   "err",
			"message": err.Error()})
		return logData, err
	}

	if _, found := options["demo"]; found {
		if ut.ToBoolean(options["demo"], false) {
			options["logData"] = logData
			logData, err = api.demoDatabase(options)
			if err != nil {
				logData = append(logData, SM{
					"stamp":   time.Now().Format(TimeLayout),
					"state":   "err",
					"message": err.Error()})
				return logData, err
			}
		}
	}

	logData = append(logData, SM{
		"stamp":   time.Now().Format(TimeLayout),
		"state":   "log",
		"message": ut.GetMessage("info_create_ok")})

	return logData, nil
}

func (api *API) demoDatabase(options IM) ([]SM, error) {
	var err error
	logData := options["logData"].([]SM)
	data := demoData()
	type item struct {
		section, datatype string
		data              []IM
	}

	postData := func(pdata item) error {
		result, err := api.Update(pdata.datatype, pdata.data)
		if err != nil {
			return err
		}
		resultStr := ""
		for index := 0; index < len(result); index++ {
			resultStr += "," + strconv.FormatInt(result[index], 10)
		}
		log := SM{
			"stamp":    time.Now().Format(TimeLayout),
			"state":    "demo",
			"datatype": pdata.datatype,
			"result":   resultStr[1:],
		}
		if pdata.section != "" {
			log["section"] = pdata.section
		}
		logData = append(logData, log)
		return nil
	}

	//load general reports and other templates
	reports, err := api.ReportList(options)
	if err != nil {
		return logData, err
	}
	resultStr := ""
	for index := 0; index < len(reports); index++ {
		params := IM{"reportkey": reports[index]["reportkey"]}
		if _, found := options["report_dir"]; found {
			params["report_dir"] = options["report_dir"]
		}
		_, err = api.ReportInstall(params)
		if err != nil {
			return logData, err
		}
		resultStr += "," + ut.ToString(reports[index]["reportkey"], "")
	}
	if len(resultStr) > 0 {
		resultStr = resultStr[1:]
	}
	logData = append(logData, SM{
		"stamp":   time.Now().Format(TimeLayout),
		"state":   "demo",
		"section": "report templates",
		"result":  resultStr,
	})

	items := []item{
		//create 3 departments and 3 eventgroups
		{section: "", datatype: "groups", data: data["groups"].([]IM)},
		//customer
		//-> def. 4 customer additional data (float,date,valuelist,customer types),
		//-> create 3 customers,
		//-> and more create and link to contacts, addresses and events
		{section: "customer", datatype: "deffield", data: data["customer"].(IM)["deffield"].([]IM)},
		{section: "customer", datatype: "customer", data: data["customer"].(IM)["customer"].([]IM)},
		{section: "customer", datatype: "address", data: data["customer"].(IM)["address"].([]IM)},
		{section: "customer", datatype: "contact", data: data["customer"].(IM)["contact"].([]IM)},
		{section: "customer", datatype: "event", data: data["customer"].(IM)["event"].([]IM)},
		//employee
		//-> def. 1 employee additional data (integer type),
		//->create 1 employee,
		//->and more create and link to contact, address and event
		{section: "employee", datatype: "deffield", data: data["employee"].(IM)["deffield"].([]IM)},
		{section: "employee", datatype: "employee", data: data["employee"].(IM)["employee"].([]IM)},
		{section: "employee", datatype: "address", data: data["employee"].(IM)["address"].([]IM)},
		{section: "employee", datatype: "contact", data: data["employee"].(IM)["contact"].([]IM)},
		{section: "employee", datatype: "event", data: data["employee"].(IM)["event"].([]IM)},
		//product
		//-> def. 3 product additional data (product,integer and valulist types),
		//->create 13 products,
		//->and more create and link to barcodes, events, prices, additional data
		{section: "product", datatype: "deffield", data: data["product"].(IM)["deffield"].([]IM)},
		{section: "product", datatype: "product", data: data["product"].(IM)["product"].([]IM)},
		{section: "product", datatype: "barcode", data: data["product"].(IM)["barcode"].([]IM)},
		{section: "product", datatype: "price", data: data["product"].(IM)["price"].([]IM)},
		{section: "product", datatype: "event", data: data["product"].(IM)["event"].([]IM)},
		//project
		//-> def. 2 project additional data,
		//->create 1 project,
		//->and more create and link to contact, address and event
		{section: "project", datatype: "deffield", data: data["project"].(IM)["deffield"].([]IM)},
		{section: "project", datatype: "project", data: data["project"].(IM)["project"].([]IM)},
		{section: "project", datatype: "address", data: data["project"].(IM)["address"].([]IM)},
		{section: "project", datatype: "contact", data: data["project"].(IM)["contact"].([]IM)},
		{section: "project", datatype: "event", data: data["project"].(IM)["event"].([]IM)},
		//tool
		//-> def. 2 tool additional data,
		//->create 3 tools,
		//->and more create and link to event and additional data
		{section: "tool", datatype: "deffield", data: data["tool"].(IM)["deffield"].([]IM)},
		{section: "tool", datatype: "tool", data: data["tool"].(IM)["tool"].([]IM)},
		{section: "tool", datatype: "event", data: data["tool"].(IM)["event"].([]IM)},
		//create +1 warehouse
		{section: "", datatype: "place", data: data["place"].([]IM)},
		//documents
		//offer, order, invoice, worksheet, rent
		{section: "document(offer,order,invoice,rent,worksheet)",
			datatype: "trans", data: data["trans_item"].(IM)["trans"].([]IM)},
		{section: "", datatype: "item", data: data["trans_item"].(IM)["item"].([]IM)},
		{section: "", datatype: "link", data: data["trans_item"].(IM)["link"].([]IM)},
		//payments
		//bank and petty cash
		{section: "payment(bank,petty cash)",
			datatype: "trans", data: data["trans_payment"].(IM)["trans"].([]IM)},
		{section: "", datatype: "payment", data: data["trans_payment"].(IM)["payment"].([]IM)},
		{section: "", datatype: "link", data: data["trans_payment"].(IM)["link"].([]IM)},
		//stock control
		//tool movement (for employee)
		//create delivery,stock transfer,correction
		//formula and production
		{section: "pstock control(tool movement,delivery,stock transfer,correction,formula,production)",
			datatype: "trans", data: data["trans_movement"].(IM)["trans"].([]IM)},
		{section: "", datatype: "movement", data: data["trans_movement"].(IM)["movement"].([]IM)},
		{section: "", datatype: "link", data: data["trans_movement"].(IM)["link"].([]IM)},
		//sample menus and menufields
		{section: "sample menus", datatype: "ui_menu", data: data["menu"].(IM)["ui_menu"].([]IM)},
		{section: "", datatype: "ui_menufields", data: data["menu"].(IM)["ui_menufields"].([]IM)},
	}
	for _, item := range items {
		err = postData(item)
		if err != nil {
			return logData, err
		}
	}

	return logData, err
}

/*
Delete - delete a record

Examples:

  Delete data by ID:

  options = map[string]interface{}{"nervatype": "address", "id": 2}
  err = api.Delete(options)

  Delete data by Key:

  options = map[string]interface{}{"nervatype": "address", "key": "customer/DMCUST/00001~1"}
  err = api.Delete(options)
*/
func (api *API) Delete(options IM) error {
	if _, found := options["id"]; found {
		options["ref_id"] = ut.ToInteger(options["id"], 0)
	}
	return api.NStore.DeleteData(IM{
		"nervatype": options["nervatype"],
		"ref_id":    options["ref_id"],
		"refnumber": options["key"],
	})
}

/*
Get - returns one or more records

Examples:

  Find data by Filter:

  options = map[string]interface{}{"nervatype": "customer", "metadata": true,
    "filter": "custname;==;First Customer Co.|custnumber;in;DMCUST/00001,DMCUST/00002"}
  _, err = api.Get(options)

  Find data by IDs:

  options = map[string]interface{}{"nervatype": "customer", "metadata": true, "ids": "2,4"}
  _, err = api.Get(options)

*/
func (api *API) Get(options IM) (results []IM, err error) {
	nervatype := ut.ToString(options["nervatype"], "")
	if nervatype == "" {
		return results, errors.New(ut.GetMessage("missing_required_field") + ": nervatype")
	}
	if _, found := api.NStore.models[nervatype]; !found {
		return results, errors.New(ut.GetMessage("invalid_nervatype") + " " + nervatype)
	}
	metadata := ut.ToBoolean(options["metadata"], false)

	query := []Query{{
		Fields: []string{"*"}, From: nervatype, Filters: []Filter{}}}
	if _, found := api.NStore.models[nervatype].(IM)["deleted"]; found {
		query[0].Filters = append(query[0].Filters, Filter{Field: "deleted", Comp: "==", Value: 0})
	}
	if ut.ToString(options["ids"], "") != "" {
		query[0].Filters = append(query[0].Filters, Filter{Field: "id", Comp: "in", Value: ut.ToString(options["ids"], "")})
	} else if _, found := options["filter"]; found {
		filters := strings.Split(ut.ToString(options["filter"], ""), "|")
		for index := 0; index < len(filters); index++ {
			fields := strings.Split(filters[index], ";")
			if len(fields) != 3 {
				return results, errors.New(ut.GetMessage("invalid_value") + "- filter: " + filters[index])
			}
			if _, found := api.NStore.models[nervatype].(IM)[fields[0]]; !found {
				return results, errors.New(ut.GetMessage("invalid_value") + "- fieldname: " + fields[0])
			}
			switch fields[1] {
			case "==", "!=", "<", "<=", ">", ">=", "in":
			default:
				return results, errors.New(ut.GetMessage("invalid_value") + "- comparison: " + fields[1])
			}
			value := fields[2]
			query[0].Filters = append(query[0].Filters, Filter{Field: fields[0], Comp: fields[1], Value: value})
		}
	} else {
		return results, errors.New(ut.GetMessage("missing_required_field") + ": filter or ids")
	}

	results, err = api.NStore.ds.Query(query, nil)
	if err != nil {
		return results, err
	}
	if len(results) > 0 && metadata {
		switch nervatype {
		case "address", "barcode", "contact", "currency", "customer", "employee", "event", "groups",
			"item", "link", "log", "movement", "price", "place", "product", "project", "rate",
			"tax", "tool", "trans":
			ids := ""
			for index := 0; index < len(results); index++ {
				ids += "," + ut.ToString(results[index]["id"], "")
			}
			ids = ids[1:]
			metadata, err := api.NStore.ds.QueryKey(IM{"qkey": "metadata", "nervatype": nervatype, "ids": ids}, nil)
			if err != nil {
				return results, err
			}
			if len(metadata) > 0 {
				for index := 0; index < len(results); index++ {
					results[index]["metadata"] = []IM{}
					for mi := 0; mi < len(metadata); mi++ {
						if metadata[mi]["ref_id"] == results[index]["id"] {
							results[index]["metadata"] = append(results[index]["metadata"].([]IM), IM{
								"id":        metadata[mi]["id"],
								"fieldname": metadata[mi]["fieldname"],
								"fieldtype": metadata[mi]["fieldtype"],
								"value":     metadata[mi]["value"],
								"notes":     metadata[mi]["notes"],
							})
						}
					}
				}
			}
		}
	}
	return results, err
}

/*
View - run raw SQL queries in safe mode

Only "select" queries and functions can be executed. Changes to the data are not saved in the database.

Examples:

  options := []map[string]interface{}{
    map[string]interface{}{
      "key":    "customers",
      "text":   "select c.id, ct.groupvalue as custtype, c.custnumber, c.custname from customer c inner join groups ct on c.custtype = ct.id where c.deleted = 0 and c.custnumber <> 'HOME'",
    },
    map[string]interface{}{
      "key":    "invoices",
      "text":   "select t.id, t.transnumber, tt.groupvalue as transtype, td.groupvalue as direction, t.transdate, c.custname, t.curr, items.amount from trans t inner join groups tt on t.transtype = tt.id inner join groups td on t.direction = td.id inner join customer c on t.customer_id = c.id inner join ( select trans_id, sum(amount) amount from item where deleted = 0 group by trans_id) items on t.id = items.trans_id where t.deleted = 0 and tt.groupvalue = 'invoice'",
    },
  }
  _, err = api.View(options)
*/
func (api *API) View(options []IM) (results IM, err error) {
	results = IM{}
	var trans interface{}
	if api.NStore.ds.Properties().Transaction {
		trans, err = api.NStore.ds.BeginTransaction()
		if err != nil {
			return results, err
		}
	}

	defer func() {
		pe := recover()
		if trans != nil {
			if rb_err := api.NStore.ds.RollbackTransaction(trans); rb_err != nil {
				return
			}
		}
		if pe != nil {
			panic(pe)
		}
	}()

	for index := 0; index < len(options); index++ {
		key := ut.ToString(options[index]["key"], "")
		if key == "" {
			return results, errors.New(ut.GetMessage("missing_required_field") + ": key")
		}
		text := ut.ToString(options[index]["text"], "")
		if text == "" {
			return results, errors.New(ut.GetMessage("missing_required_field") + ": text")
		}
		if _, valid := options[index]["values"].([]interface{}); !valid {
			return results, errors.New(ut.GetMessage("missing_required_field") + ": values")
		}
		result, err := api.NStore.ds.QuerySQL(
			text, options[index]["values"].([]interface{}), trans)
		if err != nil {
			return results, err
		}
		results[key] = result
	}
	return results, err
}

/*
Function - call a server-side function

Examples:

  The next value from the numberdef table (customer, product, invoice etc.):

  options := map[string]interface{}{
    "key": "nextNumber",
    "values": map[string]interface{}{
      "numberkey": "custnumber",
      "step":      false,
    },
  }
  _, err = api.Function(options)

  Product price (current date, all customer, all qty):

  options = map[string]interface{}{
    "key": "getPriceValue",
    "values": map[string]interface{}{
      "curr":        "EUR",
      "product_id":  2,
      "customer_id": 2,
    },
  }
  _, err = api.Function(options)

  Email sending with attached report:

	options := map[string]interface{}{
		"key": "sendEmail",
		"values": map[string]interface{}{
			"provider": "smtp",
			"email": map[string]interface{}{
				"from": "info@nervatura.com", "name": "Nervatura",
				"recipients": []interface{}{
					map[string]interface{}{"email": "sample@company.com"}},
				"subject": "Demo Invoice",
				"text":    "Email sending with attached invoice",
				"attachments": []interface{}{
					map[string]interface{}{
						"reportkey": "ntr_invoice_en",
						"nervatype": "trans",
						"refnumber": "DMINV/00001"}},
			},
		},
	}

*/
func (api *API) Function(options IM) (results interface{}, err error) {
	key := ut.ToString(options["key"], "")
	if key == "" {
		return results, errors.New(ut.GetMessage("missing_required_field") + ": key")
	}
	if _, valid := options["values"].(IM); !valid {
		return results, errors.New(ut.GetMessage("missing_required_field") + ": values")
	}
	return api.NStore.GetService(key, options["values"].(IM))
}

func (api *API) updateTransInfo(data []IM) ([]IM, error) {
	for index := 0; index < len(data); index++ {
		_, fkeys := data[index]["keys"]
		ftranstype := false
		fcustomer := false
		if fkeys {
			_, ftranstype = data[index]["keys"].(IM)["transtype"]
			_, fcustomer = data[index]["keys"].(IM)["customer_id"]
		}
		if !(fkeys && ftranstype && fcustomer) {
			options := IM{"qkey": "post_transtype"}
			options["transtype_id"] = nil
			transtypeId := ut.ToInteger(data[index]["transtype"], 0)
			if transtypeId > 0 {
				options["transtype_id"] = transtypeId
			}
			options["transtype_key"] = nil
			if fkeys && ftranstype {
				transtype := ut.ToString(data[index]["keys"].(IM)["transtype"], "")
				if transtype != "" {
					options["transtype_key"] = transtype
				}
			}
			options["customer_id"] = nil
			if _, found := data[index]["customer_id"]; found {
				customerId := ut.ToInteger(data[index]["customer_id"], 0)
				if customerId > 0 {
					options["customer_id"] = customerId
				}
			}
			options["custnumber"] = nil
			if fkeys && fcustomer {
				custnumber := ut.ToString(data[index]["keys"].(IM)["customer_id"], "")
				if custnumber != "" {
					options["custnumber"] = custnumber
				}
			}
			options["trans_id"] = nil
			if _, found := data[index]["id"]; found {
				transId := ut.ToInteger(data[index]["id"], 0)
				if transId > 0 {
					options["trans_id"] = transId
				}
			}
			info, err := api.NStore.ds.QueryKey(options, nil)
			if err != nil {
				return data, err
			}
			if len(info) > 0 {
				if !fkeys {
					data[index]["keys"] = IM{}
				}
				keys := map[string][]interface{}{}
				for index := 0; index < len(info); index++ {
					keys[info[index]["rtype"].(string)] = IL{info[index]["transtype"], info[index]["custnumber"]}
				}
				if _, found := keys["groups"]; found {
					if !ftranstype {
						data[index]["keys"].(IM)["transtype"] = keys["groups"][0]
					}
				} else if _, found := keys["trans"]; found {
					if !ftranstype {
						data[index]["keys"].(IM)["transtype"] = keys["trans"][0]
					}
				}
				if _, found := keys["customer"]; found {
					if !fcustomer {
						data[index]["keys"].(IM)["customer_id"] = keys["customer"][1]
					}
				} else if _, found := keys["trans"]; found {
					if !fcustomer && keys["trans"][1] != nil {
						data[index]["keys"].(IM)["customer_id"] = keys["trans"][1]
					}
				}
			}
		}
	}
	return data, nil
}

func (api *API) updateSetKeys(nervatype string, data []IM) ([]IM, error) {
	for index := 0; index < len(data); index++ {
		if _, found := data[index]["keys"]; found {
			for key, value := range data[index]["keys"].(IM) {
				info := IM{"fieldname": key, "reftype": "id"}
				switch key {
				case "id":
					info["nervatype"] = nervatype
					info["refnumber"] = value

				case "ref_id", "ref_id_1", "ref_id_2":
					info["nervatype"] = strings.Split(value.(string), "/")[0]
					info["refnumber"] = strings.ReplaceAll(value.(string), strings.Split(value.(string), "/")[0]+"/", "")

				default:
					if _, found := api.NStore.models[nervatype].(IM)[key]; found {
						if api.NStore.models[nervatype].(IM)[key].(MF).References != nil {
							info["nervatype"] = api.NStore.models[nervatype].(IM)[key].(MF).References[0]
							if info["nervatype"] == "groups" {
								switch key {
								case "nervatype_1", "nervatype_2":
									info["refnumber"] = "nervatype~" + value.(string)
								default:
									info["refnumber"] = key + "~" + value.(string)
								}
							} else {
								info["refnumber"] = value
								if key == "customer_id" && data[index]["keys"].(IM)["transtype"] == "invoice" {
									info["extra_info"] = true
								}
							}
						} else if api.NStore.models[nervatype].(IM)["_key"].(SL)[0] == key {
							if svalue, valid := value.(string); valid && svalue == "numberdef" {
								info["reftype"] = "numberdef"
								info["numberkey"] = key
								info["step"] = true
								info["insert_key"] = false
							} else if ivalue, valid := value.(IL); valid {
								info["reftype"] = "numberdef"
								if len(value.(IL)) > 1 {
									info["numberkey"] = ivalue[1]
								}
								info["step"] = true
								info["insert_key"] = false
							} else {
								info["nervatype"] = nervatype
								info["refnumber"] = value
								info["fieldname"] = "id"
							}
						}
					}
					if _, found := info["nervatype"]; !found {
						if info["reftype"] == "id" {
							info["nervatype"] = "invalid"
							info["refnumber"] = value
						}
					}
				}
				if info["reftype"] == "numberdef" {
					retnumber, err := api.NStore.nextNumber(info)
					if err != nil {
						return data, err
					}
					data[index][ut.ToString(info["fieldname"], "")] = retnumber
				} else {
					refValues, err := api.NStore.GetInfofromRefnumber(info)
					if err != nil {
						return data, err
					}
					data[index][ut.ToString(info["fieldname"], "")] = refValues["id"]
					extraInfo := ut.ToBoolean(info["extra_info"], false)
					if extraInfo {
						data[index]["trans_custinvoice_compname"] = refValues["compname"]
						data[index]["trans_custinvoice_comptax"] = refValues["comptax"]
						data[index]["trans_custinvoice_compaddress"] = refValues["compaddress"]
						data[index]["trans_custinvoice_custname"] = refValues["custname"]
						data[index]["trans_custinvoice_custtax"] = refValues["custtax"]
						data[index]["trans_custinvoice_custaddress"] = refValues["custaddress"]
					}
				}
			}
		}
	}
	return data, nil
}

func (api *API) updateCheckInfo(nervatype string, data []IM) ([]IM, error) {
	model := api.NStore.models[nervatype].(IM)
	for index := 0; index < len(data); index++ {
		delete(data[index], "keys")
		if _, found := data[index]["id"]; !found || (data[index]["id"] == nil) {
			for ikey, ifield := range model {
				switch ikey {
				case "_access", "_key", "_fields", "id":

				case "crdate":
					if ifield.(MF).Type == "datetime" {
						data[index]["crdate"] = time.Now().Format("2006-01-02T15:04:05-0700")
					} else if ifield.(MF).Type == "date" {
						if _, found := data[index]["crdate"]; !found {
							data[index]["crdate"] = time.Now().Format("2006-01-02")
						}
					}

				case "cruser_id":
					if api.NStore.User != nil {
						data[index]["cruser_id"] = api.NStore.User.Id
					} else {
						data[index]["cruser_id"] = 1
					}

				default:
					if ifield.(MF).NotNull && ifield.(MF).Default == nil {
						if _, found := data[index][ikey]; !found {
							return data, errors.New(ut.GetMessage("missing_required_field") + " " + ikey)
						}
					}
				}
			}
			if nervatype == "trans" {
				if _, found := data[index]["trans_transcast"]; !found {
					data[index]["trans_transcast"] = "normal"
				}
			}
		} else {
			if _, found := data[index]["crdate"]; found {
				if _, found := model["crdate"]; found {
					delete(data[index], "crdate")
				}
			}
		}
	}
	return data, nil
}

/*
Update - Add or update one or more items

If the ID (or Key) value is missing, it creates a new item. Returns the all new/updated IDs values.

Examples:

  addressData := []map[string]interface{}{
    map[string]interface{}{
      "nervatype":         10,
      "ref_id":            2,
      "zipcode":           "12345",
      "city":              "BigCity",
      "notes":             "Create a new item by IDs",
      "address_metadata1": "value1",
      "address_metadata2": "value2~note2"},
    map[string]interface{}{
      "id":                6,
      "zipcode":           "54321",
      "city":              "BigCity",
      "notes":             "Update an item by IDs",
      "address_metadata1": "value1",
      "address_metadata2": "value2~note2"},
    map[string]interface{}{
      "keys": map[string]interface{}{
        "nervatype": "customer",
        "ref_id":    "customer/DMCUST/00001"},
      "zipcode":           "12345",
      "city":              "BigCity",
      "notes":             "Create a new item by Keys",
      "address_metadata1": "value1",
      "address_metadata2": "value2~note2"},
    map[string]interface{}{
      "keys": map[string]interface{}{
        "id": "customer/DMCUST/00001~1"},
      "zipcode":           "54321",
      "city":              "BigCity",
      "notes":             "Update an item by Keys",
      "address_metadata1": "value1",
      "address_metadata2": "value2~note2"}}

  _, err = api.Update("address", addressData)

*/
func (api *API) Update(nervatype string, data []IM) (results []int64, err error) {
	if _, found := api.NStore.models[nervatype]; !found {
		return results, errors.New(ut.GetMessage("invalid_nervatype") + " " + nervatype)
	}

	if nervatype == "trans" {
		data, err = api.updateTransInfo(data)
		if err != nil {
			return results, err
		}
	}

	data, err = api.updateSetKeys(nervatype, data)
	if err != nil {
		return results, err
	}

	data, err = api.updateCheckInfo(nervatype, data)
	if err != nil {
		return results, err
	}

	var trans interface{}
	if api.NStore.ds.Properties().Transaction {
		trans, err = api.NStore.ds.BeginTransaction()
		if err != nil {
			return results, err
		}
	}

	defer func() {
		pe := recover()
		if trans != nil {
			if err != nil || pe != nil {
				if rb_err := api.NStore.ds.RollbackTransaction(trans); rb_err != nil {
					return
				}
			} else {
				err = api.NStore.ds.CommitTransaction(trans)
			}
		}
		if pe != nil {
			panic(pe)
		}
	}()

	for index := 0; index < len(data); index++ {
		var id int64
		id, err = api.NStore.UpdateData(IM{
			"nervatype":    nervatype,
			"values":       data[index],
			"validate":     true,
			"insert_row":   true,
			"insert_field": true,
			"trans":        trans,
		})
		if err != nil {
			return results, err
		}
		results = append(results, id)
	}

	return results, err
}

/*
Report - server-side PDF and CSV report generation

Examples:

  Customer PDF invoice:

  options := map[string]interface{}{
    "reportkey":   "ntr_invoice_en",
    "orientation": "portrait",
    "size":        "a4",
    "nervatype":   "trans",
    "refnumber":   "DMINV/00001",
  }
  _, err = api.Report(options)

  Customer invoice XML data:

  options = map[string]interface{}{
    "reportkey": "ntr_invoice_en",
    "output":    "xml",
    "nervatype": "trans",
    "refnumber": "DMINV/00001",
  }
  _, err = api.Report(options)

  CSV report:

  options = map[string]interface{}{
    "reportkey": "csv_vat_en",
    "filters": map[string]interface{}{
      "date_from": "2014-01-01",
      "date_to":   "2019-01-01",
      "curr":      "EUR",
    },
  }
  _, err = api.Report(options)

*/
func (api *API) Report(options IM) (results IM, err error) {
	return api.NStore.getReport(options)
}

/*
ReportList - returns all installable files from the NT_REPORT_DIR (environment variable) or report_dir (options) directory

Example:

  options := map[string]interface{}{
    "report_dir": "",
  }
  _, err = api.ReportList(options)

*/
func (api *API) ReportList(options IM) (results []IM, err error) {
	query := []Query{{
		Fields: []string{"id", "reportkey"}, From: "ui_report"}}
	reports, err := api.NStore.ds.Query(query, nil)
	if err != nil {
		return results, err
	}
	dbReports := IM{}
	for index := 0; index < len(reports); index++ {
		dbReports[ut.ToString(reports[index]["reportkey"], "")] = reports[index]["id"]
	}
	reportDir := ut.ToString(options["report_dir"], ut.ToString(api.NStore.config["NT_REPORT_DIR"], ""))
	filter := ut.ToString(options["label"], "")
	results = []IM{}

	fileInfo := func(file []byte, fileName string) {
		temp := IM{}
		if err = ut.ConvertFromByte(file, &temp); err == nil {
			if meta, found := temp["meta"].(IM); found {
				report := IM{"installed": false, "label": ""}
				report["reportkey"] = ut.ToString(meta["reportkey"], "")
				if _, found := dbReports[ut.ToString(meta["reportkey"], "")]; found {
					report["installed"] = true
				}
				report["repname"] = ut.ToString(meta["repname"], "")
				report["description"] = ut.ToString(meta["description"], "")
				report["reptype"] = ut.ToString(meta["filetype"], "")
				report["label"] = ut.ToString(meta["nervatype"], "")
				report["label"] = ut.ToString(meta["transtype"], "")
				report["filename"] = fileName
				if (filter == "") || (filter == report["label"]) {
					if (report["reptype"] == "csv") || (report["reptype"] == "pdf") {
						results = append(results, report)
					}
				}
			}
		}
	}

	if reportDir == "" {
		err = fs.WalkDir(ut.Report, path.Join("static", "templates"), func(path string, d fs.DirEntry, err error) error {
			if filepath.Ext(path) == ".json" {
				if file, err := ut.Report.ReadFile(path); err == nil {
					fileInfo(file, d.Name())
				} else {
					return err
				}
			}
			return err
		})
	} else {
		err = filepath.Walk(reportDir, func(path string, info os.FileInfo, err error) error {
			if filepath.Ext(path) == ".json" {
				if file, err := ioutil.ReadFile(filepath.Clean(path)); err == nil {
					fileInfo(file, info.Name())
				} else {
					return err
				}
			}
			return err
		})
	}

	if err != nil {
		return results, err
	}
	return results, err
}

/*
ReportDelete - delete a report from the database

Example:

  options := nt.IM{
    "reportkey": "ntr_cash_in_en",
  }
  err = api.ReportDelete(options)

*/
func (api *API) ReportDelete(options IM) (err error) {
	reportkey := ut.ToString(options["reportkey"], "")
	if reportkey == "" {
		return errors.New(ut.GetMessage("missing_required_field") + ": reportkey")
	}

	query := []Query{{
		Fields: []string{"*"}, From: "ui_report", Filters: []Filter{
			{Field: "reportkey", Comp: "==", Value: reportkey},
		}}}
	rows, err := api.NStore.ds.Query(query, nil)
	if err != nil {
		return err
	}
	if len(rows) == 0 {
		return errors.New(ut.GetMessage("missing_reportkey") + ": " + reportkey)
	}
	refID := ut.ToInteger(rows[0]["id"], 0)

	_, err = api.NStore.ds.Update(Update{IDKey: refID, Model: "ui_report"})
	return err
}

/*
ReportInstall - install a report to the database

Example:

  options := nt.IM{
    "report_dir": "",
    "reportkey":  "ntr_cash_in_en",
  }
  _, err = api.ReportInstall(options)

*/
func (api *API) ReportInstall(options IM) (result int64, err error) {
	reportkey := ut.ToString(options["reportkey"], "")
	if reportkey == "" {
		return result, errors.New(ut.GetMessage("missing_required_field") + ": reportkey")
	}
	reportDir := ut.ToString(options["report_dir"], ut.ToString(api.NStore.config["NT_REPORT_DIR"], ""))
	var file []byte
	if reportDir == "" {
		file, err = ut.Report.ReadFile(path.Join("static", "templates", reportkey+".json"))
	} else {
		file, err = ioutil.ReadFile(filepath.Clean(filepath.Join(reportDir, reportkey+".json")))
	}
	if err != nil {
		return result, err
	}

	temp := IM{}
	if err = ut.ConvertFromByte(file, &temp); err != nil {
		return result, err
	}

	report := IM{}
	meta, found := temp["meta"].(IM)
	if found {
		reportkey := ut.ToString(meta["reportkey"], "")
		if reportkey != "" {
			report["reportkey"] = reportkey
			query := []Query{{
				Fields: []string{"*"}, From: "ui_report", Filters: []Filter{
					{Field: "reportkey", Comp: "==", Value: reportkey},
				}}}
			rows, err := api.NStore.ds.Query(query, nil)
			if err != nil {
				return result, err
			}
			if len(rows) > 0 {
				return result, errors.New(ut.GetMessage("exists_template"))
			}
		} else {
			return result, errors.New(ut.GetMessage("invalid_template"))
		}
	} else {
		return result, errors.New(ut.GetMessage("invalid_template"))
	}

	groups := IM{}
	query := []Query{{
		Fields: []string{"*"}, From: "groups", Filters: []Filter{
			{Field: "groupname", Comp: "in",
				Value: "nervatype,transtype,direction,filetype,fieldtype,wheretype"},
		}}}
	rows, err := api.NStore.ds.Query(query, nil)
	if err != nil {
		return result, err
	}
	for index := 0; index < len(rows); index++ {
		if _, found := groups[rows[index]["groupname"].(string)]; !found {
			groups[rows[index]["groupname"].(string)] = IM{}
		}
		groups[rows[index]["groupname"].(string)].(IM)[rows[index]["groupvalue"].(string)] = rows[index]["id"]
	}

	report["repname"] = ut.ToString(meta["repname"], "")
	report["description"] = ut.ToString(meta["description"], "")
	if _, found := groups["nervatype"].(IM)[ut.ToString(meta["nervatype"], "")]; found {
		report["nervatype"] = ut.ToFloat(groups["nervatype"].(IM)[ut.ToString(meta["nervatype"], "")], 0)
	}
	if _, found := groups["filetype"].(IM)[ut.ToString(meta["filetype"], "")]; found {
		report["filetype"] = ut.ToFloat(groups["filetype"].(IM)[ut.ToString(meta["filetype"], "")], 0)
	}
	if _, found := groups["transtype"].(IM)[ut.ToString(meta["transtype"], "")]; found {
		report["transtype"] = ut.ToFloat(groups["transtype"].(IM)[ut.ToString(meta["transtype"], "")], 0)
	}
	if _, found := groups["direction"].(IM)[ut.ToString(meta["direction"], "")]; found {
		report["direction"] = ut.ToFloat(groups["direction"].(IM)[ut.ToString(meta["direction"], "")], 0)
	}
	report["label"] = ut.ToString(meta["label"], "")
	report["report"] = string(file)

	return api.NStore.ds.Update(Update{Model: "ui_report", Values: report})
}
