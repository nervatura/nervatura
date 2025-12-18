package api

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"
	"os"
	"slices"
	"strings"
	"time"

	cu "github.com/nervatura/component/pkg/util"
	drv "github.com/nervatura/nervatura/v6/pkg/driver"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

// DataDriver a general database interface
type DataDriver interface {
	Properties() struct {
		SQL, Transaction bool
	} //DataDriver features
	Connection() struct {
		Alias     string
		Connected bool
		Engine    string
	} //Returns the database connection
	CreateConnection(string, string) error                                                     //Create a new database connection
	CloseConnection() error                                                                    //Create a IESF Database
	Query(queries []md.Query, transaction interface{}) ([]cu.IM, error)                        //Query is a basic nosql friendly queries the database
	QuerySQL(sqlString string, params []interface{}, transaction interface{}) ([]cu.IM, error) //Executes a SQL query
	Update(options md.Update) (int64, error)                                                   //Update is a basic nosql friendly update/insert/delete and returns the update/insert id
	UpdateSQL(sqlString string, transaction interface{}) error                                 //Executes a SQL query string
	BeginTransaction() (interface{}, error)                                                    //Begins a transaction and returns an it
	CommitTransaction(trans interface{}) error                                                 //Commit a transaction
	RollbackTransaction(trans interface{}) error                                               //Rollback a transaction
}

// DataStore is the core structure of the data connection
type DataStore struct {
	Db                     DataDriver
	Alias                  string
	Config                 cu.IM
	AppLog                 *slog.Logger
	ReadAll                func(r io.Reader) ([]byte, error)
	ConvertToByte          func(data interface{}) ([]byte, error)
	ConvertFromByte        func(data []byte, result interface{}) error
	ConvertFromReader      func(data io.Reader, result interface{}) error
	ConvertToType          func(data interface{}, result any) (err error)
	GetDataField           func(data any, JSONName string) (fieldName string, fieldValue interface{})
	NewRequest             func(method string, url string, body io.Reader) (*http.Request, error)
	RequestDo              func(req *http.Request) (*http.Response, error)
	CreateLoginToken       func(params cu.SM, config cu.IM) (result string, err error)
	ParseToken             func(token string, keyMap []cu.SM, config cu.IM) (cu.IM, error)
	CreatePasswordHash     func(password string) (hash string, err error)
	ComparePasswordAndHash func(password string, hash string) (err error)
	ReadFile               func(name string) ([]byte, error)
	NewSmtpClient          func(conn net.Conn, host string) (md.SmtpClient, error)
}

func NewDataStore(config cu.IM, alias string, appLog *slog.Logger) *DataStore {
	ds := &DataStore{
		Config:                 config,
		Db:                     &drv.SQLDriver{Config: config},
		Alias:                  cu.ToString(alias, cu.ToString(config["NT_DEFAULT_ALIAS"], "")),
		AppLog:                 appLog,
		ReadAll:                io.ReadAll,
		ConvertToByte:          cu.ConvertToByte,
		ConvertFromByte:        cu.ConvertFromByte,
		ConvertFromReader:      cu.ConvertFromReader,
		GetDataField:           ut.GetDataField,
		NewRequest:             http.NewRequest,
		RequestDo:              (&http.Client{Timeout: time.Second * 20}).Do,
		CreatePasswordHash:     ut.CreatePasswordHash,
		ComparePasswordAndHash: ut.ComparePasswordAndHash,
		ConvertToType:          ut.ConvertToType,
		CreateLoginToken:       ut.CreateLoginToken,
		ParseToken:             ut.ParseToken,
		ReadFile:               os.ReadFile,
		NewSmtpClient:          ut.SmtpClient,
	}
	if db, found := ds.Config["db"].(DataDriver); found {
		ds.Db = db
	}
	return ds
}

func (ds *DataStore) SetError(sysErr error, publicErr error) error {
	hideError := cu.ToBoolean(ds.Config["NT_DEV_HIDE_ERROR"], false)
	ds.AppLog.Error(sysErr.Error())
	if hideError {
		return publicErr
	}
	return sysErr
}

func (ds *DataStore) checkConnection() error {
	if !ds.Db.Connection().Connected {
		connStr := cu.ToString(ds.Config["NT_ALIAS_"+strings.ToUpper(ds.Alias)], os.Getenv("NT_ALIAS_"+strings.ToUpper(ds.Alias)))
		if connStr == "" {
			return errors.New(ut.GetMessage("en", "missing_database"))
		}
		err := ds.Db.CreateConnection(ds.Alias, connStr)
		if err != nil {
			return err
		}
	}
	return nil
}

/*
func (ds *DataStore) StoreUpdateLog(resultID int64, update md.Update) {
	if cu.ToBoolean(ds.Config["NT_UPDATE_LOG"], false) {
		values := cu.IM{
			"log_type":  "LOG_UPDATE",
			"ref_type":  update.Model,
			"ref_code":  "",
			"auth_code": "",
		}
		if update.IDKey == 0 {
			values["log_type"] = "LOG_INSERT"
		}
		if (update.Values == nil) && (update.IDKey > 0) {
			values["log_type"] = "LOG_DELETE"
		}
		logValues := cu.ToString(ds.Config["NT_UPDATE_LOG_"+strings.ToUpper(update.Model)], "")
		if strings.Contains(logValues, strings.Split(cu.ToString(update.Values["log_type"], ""), "_")[1]) {
			if update.Values != nil {
				if values_byte, err := ds.ConvertToByte(update.Values); err == nil {
					values["data"] = values_byte
				}
			}
			ds.Db.Update(md.Update{Model: "update_log", Values: values})
		}
	}
}
*/

func (ds *DataStore) StoreDataUpdate(update md.Update) (storeID int64, err error) {
	err = ds.checkConnection()
	if err != nil {
		err = ds.SetError(err, errors.New(http.StatusText(http.StatusInternalServerError)))
		return storeID, err
	}

	resultID, err := ds.Db.Update(update)
	if err != nil {
		err = ds.SetError(err, errors.New(http.StatusText(http.StatusNotModified)))
		return storeID, err
	}
	if resultID == 0 {
		return storeID, errors.New(http.StatusText(http.StatusNotFound))
	}
	return resultID, err
}

func (ds *DataStore) StoreDataQuery(query md.Query, foundErr bool) (rows []cu.IM, err error) {
	rows = []cu.IM{}
	err = ds.checkConnection()
	if err != nil {
		err = ds.SetError(err, errors.New(http.StatusText(http.StatusInternalServerError)))
		return rows, err
	}
	rows, err = ds.Db.Query([]md.Query{query}, nil)
	if err != nil {
		err = ds.SetError(err, errors.New(http.StatusText(http.StatusInternalServerError)))
		return rows, err
	}
	if foundErr && len(rows) == 0 {
		err = errors.New(http.StatusText(http.StatusNotFound))
	}
	return rows, err
}

func (ds *DataStore) StoreDataQueries(queries []md.Query) (rows []cu.IM, err error) {
	rows = []cu.IM{}
	err = ds.checkConnection()
	if err != nil {
		err = ds.SetError(err, errors.New(http.StatusText(http.StatusInternalServerError)))
		return rows, err
	}
	rows, err = ds.Db.Query(queries, nil)
	if err != nil {
		err = ds.SetError(err, errors.New(http.StatusText(http.StatusInternalServerError)))
		return rows, err
	}
	return rows, err
}

func (ds *DataStore) StoreDataGet(params cu.IM, foundErr bool) (result []cu.IM, err error) {
	query := md.Query{
		Fields:  []string{"*"},
		From:    cu.ToString(params["model"], ""),
		Filters: []md.Filter{},
	}
	if limit := cu.ToInteger(params["limit"], 0); limit > 0 {
		query.Limit = limit
	}
	if offset := cu.ToInteger(params["offset"], 0); offset > 0 {
		query.Offset = offset
	}
	if !strings.Contains(query.From, "_") {
		query.Filters = append(query.Filters, md.Filter{Field: "deleted", Comp: "==", Value: false})
	}
	if fields, found := params["fields"].([]string); found {
		query.Fields = fields
	}
	queryFilters := []string{}
	if filter, found := params["filter"].(string); found {
		queryFilters = append(queryFilters, filter)
	}
	for key, value := range params {
		if !slices.Contains([]string{"model", "fields", "tag", "limit", "offset", "filter"}, key) {
			query.Filters = append(query.Filters, md.Filter{Field: key, Comp: "==", Value: value})
		}
		if key == "tag" {
			queryFilters = append(queryFilters,
				fmt.Sprintf("code in (select code from %s_tags where tag='%s')",
					strings.Split(strings.Split(query.From, "_")[0], " ")[0], cu.ToString(value, "")))
		}
	}
	if len(queryFilters) > 0 {
		query.Filter = "(" + strings.Join(queryFilters, " and ") + ")"
		if len(query.Filters) > 0 {
			query.Filter = " and " + query.Filter
		}
	}
	return ds.StoreDataQuery(query, foundErr)
}

// GetBodyData - convert body to bytes and struct
func (ds *DataStore) GetBodyData(modelName string, body io.ReadCloser, modelData any) (data cu.IM, inputFields []string, metaFields []string, err error) {
	data = cu.IM{}
	inputFields = []string{}
	metaFields = []string{}
	// convert body to bytes
	var bodyBytes []byte
	if bodyBytes, err = ds.ReadAll(body); err != nil {
		return data, inputFields, metaFields, err
	}
	// convert body bytes to map struct
	if err = ds.ConvertFromByte(bodyBytes, &data); err != nil {
		return data, inputFields, metaFields, err
	}
	// convert body bytes to model
	if err = ds.ConvertFromByte(bodyBytes, modelData); err != nil {
		return data, inputFields, metaFields, err
	}

	// get input fields
	for key := range data {
		inputFields = append(inputFields, key)
	}
	// get meta fields
	if metaField, found := data[modelName+"_meta"].(cu.IM); found {
		for key := range metaField {
			metaFields = append(metaFields, key)
		}
	}

	return data, inputFields, metaFields, err
}

func (ds *DataStore) defaultSetValue(modelName string, itemRow cu.IM, values cu.IM, inputName string, fieldValue interface{}) cu.IM {
	switch inputName {
	case "id", "code", modelName + "_meta", "time_stamp":
	// protected fields
	case modelName + "_map":
		if mapValue, found := fieldValue.(cu.IM); found {
			value := cu.MergeIM(cu.ToIM(itemRow[inputName], cu.IM{}), mapValue)
			jvalue, err := ds.ConvertToByte(value)
			if err == nil {
				values[inputName] = string(jvalue[:])
			}
		}
	case "addresses", "contacts", "events", "address", "contact", "data":
		// json fields
		jvalue, err := ds.ConvertToByte(fieldValue)
		if err == nil {
			values[inputName] = string(jvalue[:])
		}
	case "user_group", "barcode_type", "customer_type", "log_type", "movement_type", "place_type", "price_type",
		"product_type", "rate_type", "trans_type", "direction", "config_type":
		if _, found := fieldValue.(string); !found {
			values[inputName] = md.GetEnumString(inputName, fieldValue)
		} else {
			values[inputName] = fieldValue
		}
	case "link_type_1", "link_type_2":
		if _, found := fieldValue.(string); !found {
			values[inputName] = md.GetEnumString("link_type", fieldValue)
		} else {
			values[inputName] = fieldValue
		}
	case "paid_date", "valid_from", "valid_to", "rate_date", "trans_date", "shipping_time":
		values[inputName] = cu.ToString(fieldValue, "")
	default:
		values[inputName] = fieldValue
	}
	return values
}

// SetUpdateValue - default set update value function
func (ds *DataStore) SetUpdateValue(
	modelName string, item cu.IM, inputData any, inputFields []string,
	setValue func(modelName string, itemRow cu.IM, values cu.IM, inputName string, fieldValue interface{}) cu.IM,
) (values cu.IM, err error) {
	values = cu.IM{}
	if setValue == nil {
		// default
		setValue = ds.defaultSetValue
	}
	for _, inputName := range inputFields {
		if fieldName, fieldValue := ds.GetDataField(inputData, inputName); fieldName != "" {
			values = setValue(modelName, item, values, inputName, fieldValue)
		} else {
			return values, errors.New(http.StatusText(http.StatusUnprocessableEntity) + ": " + inputName)
		}
	}
	return values, err
}

// MergeMetaData - merge meta data
func (ds *DataStore) MergeMetaData(modelName string, item cu.IM, inputMeta any, metaFields []string) (result string, err error) {
	metaData := cu.IM{}
	if meta, found := item[modelName+"_meta"].(cu.IM); found {
		metaData = meta
	}

	if len(metaFields) > 0 {
		for _, inputName := range metaFields {
			if fieldName, fieldValue := ds.GetDataField(inputMeta, inputName); fieldName != "" {
				metaData[inputName] = fieldValue
			} else {
				return result, errors.New(http.StatusText(http.StatusUnprocessableEntity) + ": " + inputName)
			}
		}
	}
	var metaByte []byte
	if metaByte, err = ds.ConvertToByte(metaData); err == nil {
		result = string(metaByte[:])
	}

	return result, err
}

func (ds *DataStore) GetDataByID(model string, id int64, code string, foundErr bool) (result []cu.IM, err error) {
	if id == 0 && code == "" {
		return result, errors.New(http.StatusText(http.StatusUnprocessableEntity) + ": id or code is required")
	}
	query := md.Query{
		Fields:  []string{"*"},
		From:    model,
		Filters: []md.Filter{},
	}
	if id > 0 {
		query.Filters = append(query.Filters, md.Filter{Field: "id", Comp: "==", Value: id})
	}
	if id == 0 && code != "" {
		query.Filters = append(query.Filters, md.Filter{Field: "code", Comp: "==", Value: code})
	}
	if !strings.Contains(model, "_") {
		query.Filters = append(query.Filters, md.Filter{Field: "deleted", Comp: "==", Value: false})
	}
	return ds.StoreDataQuery(query, foundErr)
}

func (ds *DataStore) UpdateData(options md.UpdateDataOptions) (storeID int64, err error) {
	var rows []cu.IM
	if rows, err = ds.GetDataByID(options.Model, options.IDKey, options.Code, true); err != nil {
		return storeID, err
	}

	var values cu.IM
	if values, err = ds.SetUpdateValue(options.Model, rows[0], options.Data, options.Fields, options.SetValue); err == nil {
		if len(options.MetaFields) > 0 {
			if values[options.Model+"_meta"], err = ds.MergeMetaData(options.Model, rows[0], options.Meta, options.MetaFields); err != nil {
				return storeID, err
			}
		}
		modelID := cu.ToInteger(rows[0]["id"], 0)
		storeID, err = ds.StoreDataUpdate(md.Update{Values: values, Model: options.Model, IDKey: modelID})
	}

	return storeID, err
}

func (ds *DataStore) DataDelete(model string, id int64, code string) (err error) {
	var rows []cu.IM
	if rows, err = ds.GetDataByID(model, id, code, true); err != nil {
		return err
	}
	modelID := cu.ToInteger(rows[0]["id"], 0)
	values := cu.IM{"deleted": true}
	if model == "trans" {
		transType := cu.ToString(rows[0]["trans_type"], "")
		direction := cu.ToString(rows[0]["direction"], "")
		if (slices.Contains([]string{md.TransTypeReceipt.String(), md.TransTypeInvoice.String()}, transType) &&
			direction == md.DirectionOut.String()) || transType == md.TransTypeCash.String() {
			values = cu.IM{}
			transMeta := cu.ToIM(rows[0]["trans_meta"], cu.IM{})
			transMeta["status"] = md.TransStatusDeleted.String()
			ut.ConvertByteToIMData(transMeta, values, "trans_meta")
		}
	}
	_, err = ds.StoreDataUpdate(md.Update{Model: model, Values: values, IDKey: modelID})
	return err
}

func (ds *DataStore) GetData(query md.Query, result any) (err error) {
	var rows []cu.IM
	if rows, err = ds.StoreDataQuery(query, false); err == nil {
		if len(rows) > 0 {
			return ds.ConvertToType(rows, result)
		}
	}
	return err
}
