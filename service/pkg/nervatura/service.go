package nervatura

import (
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"net/smtp"
	"strconv"
	"strings"
	"time"

	ut "github.com/nervatura/nervatura/service/pkg/utils"
)

//GetService - call Nervatura server side functions and services
func (nstore *NervaStore) GetService(key string, options IM) (interface{}, error) {

	switch key {

	case "nextNumber":
		return nstore.nextNumber(options)

	case "getPriceValue":
		return nstore.getPriceValue(options)

	case "sendEmail":
		return nstore.sendEmail(options)

	}
	return nil, errors.New(ut.GetMessage("unknown_method") + ": " + key)
}

//nextNumber - get the next value from the numberdef table (transnumber, custnumber, partnumber etc.)
func (nstore *NervaStore) nextNumber(options IM) (retnumber string, err error) {

	numberkey := ut.ToString(options["numberkey"], "")
	if numberkey == "" {
		return retnumber, errors.New(ut.GetMessage("missing_required_field") + ": numberkey")
	}
	step := ut.ToBoolean(options["step"], true)
	insertKey := ut.ToBoolean(options["insert_key"], true)

	if ok, err := nstore.connected(); !ok || err != nil {
		return retnumber, errors.New(ut.GetMessage("not_connect"))
	}

	var trans interface{}
	if _, found := options["trans"]; found {
		trans = options["trans"]
	} else if nstore.ds.Properties().Transaction {
		trans, err = nstore.ds.BeginTransaction()
		if err != nil {
			return retnumber, err
		}
	}
	defer func() {
		pe := recover()
		if trans != nil {
			if _, found := options["trans"]; !found {
				if err != nil || pe != nil {
					if rb_err := nstore.ds.RollbackTransaction(trans); rb_err != nil {
						return
					}
				} else {
					err = nstore.ds.CommitTransaction(trans)
				}
			}
		}
		if pe != nil {
			panic(pe)
		}
	}()

	query := []Query{{
		Fields: []string{"*"}, From: "numberdef", Filters: []Filter{
			{Field: "numberkey", Comp: "==", Value: numberkey}}}}
	result, err := nstore.ds.Query(query, trans)
	if err != nil {
		return retnumber, err
	}

	var values IM
	id, curvalue, length := int64(0), int64(0), 5
	if len(result) == 0 {
		if insertKey {
			values = IM{"numberkey": numberkey,
				"prefix": strings.ToUpper(numberkey[:3]), "curvalue": curvalue,
				"isyear": 1, "sep": "/", "len": length, "visible": 1, "readonly": 0}
			data := Update{Values: values, Model: "numberdef", Trans: trans}
			id, err = nstore.ds.Update(data)
			if err != nil {
				return retnumber, err
			}
			values["id"] = id
		} else {
			return retnumber, errors.New(ut.GetMessage("invalid_value") + ": refnumber")
		}
	} else {
		id = result[0]["id"].(int64)
		curvalue = result[0]["curvalue"].(int64)
		length = int(result[0]["len"].(int64))
		values = result[0]
	}

	if values["prefix"] != "" && values["prefix"] != nil {
		retnumber = values["prefix"].(string) + values["sep"].(string)
	}
	if values["isyear"] == 1 || values["isyear"] == "1" {
		transyear := time.Now().Format("2006")
		query := []Query{{
			Fields: []string{"value"}, From: "fieldvalue", Filters: []Filter{
				{Field: "fieldname", Comp: "==", Value: "transyear"},
				{Field: "ref_id", Comp: "is", Value: "null"}}}}
		result, err = nstore.ds.Query(query, trans)
		if err != nil {
			return retnumber, err
		}
		if len(result) > 0 {
			transyear = result[0]["value"].(string)
		}
		retnumber += transyear + values["sep"].(string)
	}
	value := strings.Repeat("0", length)
	value += strconv.FormatInt((curvalue + 1), 10)
	vlen := len(value)
	retnumber += value[vlen-length : vlen]
	if step {
		data := Update{Values: IM{"curvalue": curvalue + 1}, Model: "numberdef",
			IDKey: id, Trans: trans}
		_, err := nstore.ds.Update(data)
		if err != nil {
			return retnumber, err
		}
	}
	return retnumber, nil
}

//getPriceValue - get product price
func (nstore *NervaStore) getPriceValue(options IM) (results IM, err error) {
	results = IM{"price": float64(0), "discount": float64(0)}
	params := IM{
		"qkey":        "listprice",
		"curr":        options["curr"],
		"product_id":  options["product_id"],
		"vendorprice": options["vendorprice"],
		"posdate":     options["posdate"],
		"qty":         options["qty"],
		"customer_id": options["customer_id"],
	}
	if _, found := options["curr"]; !found {
		return results, errors.New(ut.GetMessage("missing_required_field") + ": curr")
	}

	if _, found := options["product_id"]; !found {
		return results, errors.New(ut.GetMessage("missing_required_field") + ": product_id")
	}

	params["vendorprice"] = ut.ToFloat(options["vendorprice"], 0)
	params["posdate"] = ut.ToString(options["posdate"], time.Now().Format("2006-01-02"))
	params["qty"] = ut.ToFloat(options["qty"], 0)

	//best_listprice
	pdata, err := nstore.ds.QueryKey(params, nil)
	if err != nil {
		return results, err
	}
	if len(pdata) > 0 {
		if pdata[0]["mp"] != nil {
			results["price"] = ut.ToFloat(pdata[0]["mp"], 0)
		}
	}

	if _, found := options["customer_id"]; found {
		//customer discount
		query := []Query{{
			Fields: []string{"*"}, From: "customer", Filters: []Filter{
				{Field: "id", Comp: "==", Value: params["customer_id"]},
			}}}
		discount, err := nstore.ds.Query(query, nil)
		if err != nil {
			return results, err
		}
		if len(discount) > 0 {
			if discount[0]["discount"] != nil {
				results["discount"] = discount[0]["discount"]
			}
		}
	}

	if _, found := options["customer_id"]; found {
		//best_custprice
		params["qkey"] = "custprice"
		pdata, err := nstore.ds.QueryKey(params, nil)
		if err != nil {
			return results, err
		}
		if len(pdata) > 0 {
			if pdata[0]["mp"] != nil {
				price := ut.ToFloat(pdata[0]["mp"], 0)
				if results["price"].(float64) > price || results["price"] == 0 {
					results["price"] = price
					results["discount"] = 0
				}
			}
		}
	}

	if _, found := options["customer_id"]; found {
		//best_grouprice
		params["qkey"] = "grouprice"
		pdata, err := nstore.ds.QueryKey(params, nil)
		if err != nil {
			return results, err
		}
		if len(pdata) > 0 {
			if pdata[0]["mp"] != nil {
				price := ut.ToFloat(pdata[0]["mp"], 0)
				if results["price"].(float64) > price || results["price"] == 0 {
					results["price"] = price
					results["discount"] = 0
				}
			}
		}
	}

	return results, nil
}

func (nstore *NervaStore) createEmail(from string, emailTo []string, emailOpt IM) (string, error) {
	delimeter := "**=myohmy689407924327"

	emailMsg := fmt.Sprintf("From: %s\r\n", from)
	emailMsg += fmt.Sprintf("To: %s\r\n", strings.Join(emailTo, ";"))
	emailMsg += fmt.Sprintf("Subject: %s\r\n", ut.ToString(emailOpt["subject"], ""))

	emailMsg += "MIME-Version: 1.0\r\n"
	emailMsg += fmt.Sprintf("Content-Type: multipart/mixed; boundary=\"%s\"\r\n", delimeter)

	emailMsg += fmt.Sprintf("\r\n--%s\r\n", delimeter)
	emailMsg += "Content-Type: text/html; charset=\"utf-8\"\r\n"
	emailMsg += "Content-Transfer-Encoding: 7bit\r\n"
	if _, found := emailOpt["html"]; found {
		emailMsg += fmt.Sprintf("\r\n%s\r\n", ut.ToString(emailOpt["html"], ""))
	} else {
		emailMsg += fmt.Sprintf("\r\n%s\r\n", ut.ToString(emailOpt["text"], ""))
	}

	if attachments, withAttachments := emailOpt["attachments"].([]interface{}); withAttachments {
		for index := 0; index < len(attachments); index++ {
			attachment := attachments[index].(IM)
			filename := "docs_" + strconv.Itoa(index+1) + ".pdf"
			if _, found := attachment["filename"]; found {
				filename = ut.ToString(attachment["filename"], "")
			}
			emailMsg += fmt.Sprintf("\r\n--%s\r\n", delimeter)
			emailMsg += "Content-Type: application/pdf; charset=\"utf-8\"\r\n"
			emailMsg += "Content-Transfer-Encoding: base64\r\n"
			emailMsg += "Content-Disposition: attachment;filename=\"" + filename + "\"\r\n"

			params := IM{"output": "pdf"}
			if _, found := attachment["reportkey"]; found {
				params["reportkey"] = attachment["reportkey"]
			}
			if _, found := attachment["report_id"]; found {
				params["report_id"] = attachment["report_id"]
			}
			if _, found := attachment["nervatype"]; found {
				params["nervatype"] = attachment["nervatype"]
			}
			if _, found := attachment["refnumber"]; found {
				params["refnumber"] = attachment["refnumber"]
			}
			if _, found := attachment["ref_id"]; found {
				params["filters"] = IM{"@id": attachment["ref_id"]}
			}
			report, err := nstore.getReport(params)
			if err != nil {
				return "", err
			}
			emailMsg += "\r\n" + base64.StdEncoding.EncodeToString(report["template"].([]uint8))
		}
	}
	return emailMsg, nil
}

func (nstore *NervaStore) sendEmail(options IM) (results IM, err error) {
	results = IM{"result": "OK"}

	emailOpt, valid := options["email"].(IM)
	if !valid {
		return results, errors.New(ut.GetMessage("missing_required_field") + ": email")
	}
	if ut.ToString(options["provider"], "smtp") != "smtp" {
		return results, errors.New(ut.GetMessage("invalid_provider"))
	}

	username := ut.ToString(nstore.config["NT_SMTP_USER"], "")
	password := ut.ToString(nstore.config["NT_SMTP_PASSWORD"], "")
	host := ut.ToString(nstore.config["NT_SMTP_HOST"], "")
	port := ut.ToInteger(nstore.config["NT_SMTP_PORT"], 465)
	tlsMin := uint16(ut.ToInteger(nstore.config["NT_SMTP_TLS_MIN_VERSION"], 0))

	tlsConfig := tls.Config{ServerName: host, InsecureSkipVerify: false, MinVersion: tls.VersionTLS13}
	if tlsMin > 0 {
		tlsConfig.MinVersion = tlsMin
	}
	conn, connErr := tls.Dial("tcp", fmt.Sprintf("%s:%d", host, port), &tlsConfig)
	if connErr != nil {
		return results, connErr
	}
	defer conn.Close()

	client, clientErr := smtp.NewClient(conn, host)
	if clientErr != nil {
		return results, clientErr
	}
	defer client.Close()

	auth := smtp.PlainAuth("", username, password, host)
	if err := client.Auth(auth); err != nil {
		return results, err
	}

	from := ut.ToString(emailOpt["from"], username)
	if err := client.Mail(from); err != nil {
		return results, err
	}
	emailTo := []string{}
	if recipients, valid := emailOpt["recipients"].([]interface{}); valid {
		for index := 0; index < len(recipients); index++ {
			if email, valid := recipients[index].(IM)["email"].(string); valid {
				emailTo = append(emailTo, email)
				if err := client.Rcpt(email); err != nil {
					return results, err
				}
			}
		}
	} else {
		return results, errors.New(ut.GetMessage("missing_required_field") + ": recipients")
	}

	writer, writerErr := client.Data()
	if writerErr != nil {
		return results, writerErr
	}
	emailMsg, err := nstore.createEmail(from, emailTo, emailOpt)
	if err != nil {
		return results, err
	}

	if _, err := writer.Write([]byte(emailMsg)); err != nil {
		return results, err
	}

	if err = writer.Close(); err != nil {
		return results, err
	}

	return results, client.Quit()
}
