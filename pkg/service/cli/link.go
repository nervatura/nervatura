package cli

import (
	"bytes"
	"errors"
	"io"
	"net/http"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/pkg/api"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

// LinkInsert - create new link
func (cli *CLIService) LinkInsert(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)

	// convert request body to struct and schema validation
	var data md.Link = md.Link{
		LinkType1: md.LinkType(md.LinkTypeTrans),
		LinkType2: md.LinkType(md.LinkTypeTrans),
		LinkMeta: md.LinkMeta{
			Tags: []string{},
		},
		LinkMap: cu.IM{},
	}
	err := ds.ConvertFromByte([]byte(requestData), &data)
	if err != nil {
		return cli.respondString(http.StatusUnprocessableEntity, nil, http.StatusUnprocessableEntity, err)
	}

	if data.LinkCode1 == "" || data.LinkCode2 == "" {
		err = errors.New("link link_code_1 and link_code_2 are required")
		return cli.respondString(http.StatusUnprocessableEntity, nil, http.StatusUnprocessableEntity, err)
	}

	// prepare values for database update
	values := cu.IM{
		"link_type_1": data.LinkType1.String(),
		"link_code_1": data.LinkCode1,
		"link_type_2": data.LinkType2.String(),
		"link_code_2": data.LinkCode2,
	}
	if data.Code != "" {
		values["code"] = data.Code
	}

	ut.ConvertByteToIMData(data.LinkMeta, values, "link_meta")
	ut.ConvertByteToIMData(data.LinkMap, values, "link_map")

	// database insert
	var rows []cu.IM
	var result cu.IM
	var linkID int64
	if linkID, err = ds.StoreDataUpdate(md.Update{Values: values, Model: "link"}); err == nil {
		if rows, err = ds.StoreDataGet(cu.IM{"id": linkID, "model": "link"}, true); err == nil {
			result = rows[0]
		}
	}
	return cli.respondString(http.StatusCreated, result, http.StatusUnprocessableEntity, err)
}

// LinkUpdate - update link
func (cli *CLIService) LinkUpdate(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	linkID := cu.ToInteger(options["id"], 0)
	linkCode := cu.ToString(options["code"], "")

	var link md.Link
	var inputFields, metaFields []string
	var err error
	reader := io.NopCloser(bytes.NewReader([]byte(requestData)))
	if _, inputFields, metaFields, err = ds.GetBodyData("link", reader, &link); err == nil {
		_, err = ds.UpdateData(md.UpdateDataOptions{
			Model: "link", IDKey: linkID, Code: linkCode,
			Data: link, Meta: link.LinkMeta, Fields: inputFields, MetaFields: metaFields,
		})
	}
	return cli.respondString(http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

func (cli *CLIService) LinkDelete(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	linkID := cu.ToInteger(options["id"], 0)
	linkCode := cu.ToString(options["code"], "")
	err := ds.DataDelete("link", linkID, linkCode)
	return cli.respondString(http.StatusNoContent, nil, http.StatusUnprocessableEntity, err)
}

// LinkQuery - get links
func (cli *CLIService) LinkQuery(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)

	var params cu.IM = cu.IM{"model": "link"}
	for _, v := range []string{"link_type_1", "link_code_1", "link_type_2", "link_code_2", "tag"} {
		if options[v] != nil {
			params[v] = options[v]
		}
	}
	response, err := ds.StoreDataGet(params, false)
	return cli.respondString(http.StatusOK, response, http.StatusUnprocessableEntity, err)
}

// LinkGet - get link by id or code
func (cli *CLIService) LinkGet(options cu.IM, requestData string) string {
	ds := options["ds"].(*api.DataStore)
	linkID := cu.ToInteger(options["id"], 0)
	linkCode := cu.ToString(options["code"], "")
	var err error
	var links []cu.IM
	var response interface{}
	errStatus := http.StatusUnprocessableEntity
	if links, err = ds.GetDataByID("link", linkID, linkCode, true); err == nil {
		response = links[0]
	} else if err.Error() == http.StatusText(http.StatusNotFound) {
		errStatus = http.StatusNotFound
	}
	return cli.respondString(http.StatusOK, response, errStatus, err)
}
