package mcp

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/google/jsonschema-go/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	cu "github.com/nervatura/component/pkg/util"
)

type SchemaType int

const (
	SchemaTypeInput SchemaType = iota
	SchemaTypeParameter
	SchemaTypeResult
	SchemaTypeResultList
)

type ModelSchema struct {
	Name              string
	Prefix            string
	CreateInputSchema func(scope string) *jsonschema.Schema
	UpdateInputSchema func(scope string) *jsonschema.Schema
	QueryInputSchema  func(scope string) *jsonschema.Schema
	QueryOutputSchema func(scope string) *jsonschema.Schema
	LoadData          func(data any) (modelData, metaData any, err error)
	LoadList          func(rows []cu.IM) (items any, err error)
	Examples          map[string][]any
	PrimaryFields     []string
	Required          []string
}

type ModelExtendSchema struct {
	Model             string
	ViewName          cu.SM
	ModelFromCode     func(code string) (model, field string, err error)
	CreateInputSchema func(scope string) *jsonschema.Schema
	UpdateInputSchema func(scope string) *jsonschema.Schema
	QueryInputSchema  func(scope string) *jsonschema.Schema
	QueryOutputSchema func(scope string) *jsonschema.Schema
	LoadData          func(data any) (modelData any, err error)
	LoadList          func(model string, rows []cu.IM) (items any, err error)
}

type UpdateResponseData struct {
	Model string `json:"model" jsonschema:"The model name."`
	Code  string `json:"code" jsonschema:"The unique key of the model data."`
	ID    int64  `json:"id" jsonschema:"The database primary key of the model data."`
}

func getSchemaMap() (schemaMap map[string]*ModelSchema) {
	return map[string]*ModelSchema{
		"customer": CustomerSchema(),
	}
}

func getExtendSchemaMap() (schemaMap map[string]*ModelExtendSchema) {
	return map[string]*ModelExtendSchema{
		"contact": ContactSchema(),
	}
}

func getModelSchemaByPrefix(prefix string) (ms *ModelSchema, err error) {
	for model := range getSchemaMap() {
		if prefix == getSchemaMap()[model].Prefix {
			return getSchemaMap()[model], nil
		}
	}
	return nil, fmt.Errorf("invalid model prefix: %s", prefix)
}

func getParamsMeta(req *mcp.CallToolRequest) (meta cu.IM) {
	meta = cu.IM{}
	for key, value := range req.Params.Meta {
		if !strings.Contains(strings.ToLower(key), "token") {
			meta[key] = value
		}
	}
	return meta
}

func getSchemaData(data cu.IM, ms *ModelSchema, paramsMeta cu.IM) (modelData, metaData any, inputFields []string, metaFields []string, err error) {

	inputFields = []string{}
	metaFields = []string{}

	// get input fields
	modelMeta := cu.IM{}
	for key, value := range data {
		if slices.Contains(ms.PrimaryFields, key) {
			inputFields = append(inputFields, key)
		} else {
			metaFields = append(metaFields, key)
			modelMeta[key] = value
		}
	}
	data[ms.Name+"_meta"] = modelMeta
	if len(metaFields) > 0 {
		inputFields = append(inputFields, ms.Name+"_meta")
	}
	if len(paramsMeta) > 0 {
		inputFields = append(inputFields, ms.Name+"_map")
		data[ms.Name+"_map"] = paramsMeta
	}

	if modelData, metaData, err = ms.LoadData(data); err == nil {
		if cu.ToString(data["code"], "") == "" {
			for _, field := range ms.Required {
				if !slices.Contains(inputFields, field) {
					return modelData, metaData, inputFields, metaFields, errors.New(field + " is required")
				}
			}
		}
	}

	return modelData, metaData, inputFields, metaFields, err
}
