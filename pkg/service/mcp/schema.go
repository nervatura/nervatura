package mcp

import (
	"errors"
	"fmt"
	"slices"

	"github.com/google/jsonschema-go/jsonschema"
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
	Name             string
	Prefix           string
	ResultType       func() any
	ResultListType   func() any
	InputSchema      func() (*jsonschema.Schema, error)
	ParameterSchema  func() (*jsonschema.Schema, error)
	ResultSchema     func() (*jsonschema.Schema, error)
	ResultListSchema func() (*jsonschema.Schema, error)
	SchemaModify     func(schemaType SchemaType, schema *jsonschema.Schema)
	LoadData         func(data any) (modelData, metaData any, err error)
	InsertValues     func(data any) (values cu.IM)
	Examples         map[string][]any
	PrimaryFields    []string
	Required         []string
}

type UpdateResponseData struct {
	Model string `json:"model" jsonschema:"The model name."`
	Code  string `json:"code" jsonschema:"The unique key of the model data."`
	ID    int64  `json:"id" jsonschema:"The database primary key of the model data."`
}

func getSchemaMap() (schemaMap map[string]*ModelSchema) {
	return map[string]*ModelSchema{
		"customer": CustomerSchema(),
		"product":  ProductSchema(),
	}
}

func makeModelSchema(model string, schemaType SchemaType) (schema *jsonschema.Schema) {
	schema = &jsonschema.Schema{}
	var sm *ModelSchema = getSchemaMap()[model]
	switch schemaType {

	case SchemaTypeInput:
		if inputType, err := sm.InputSchema(); err == nil {
			schema = inputType
		}
		sm.SchemaModify(schemaType, schema)

	case SchemaTypeParameter:
		if parameterType, err := sm.ParameterSchema(); err == nil {
			schema = parameterType
		}
		sm.SchemaModify(schemaType, schema)

	case SchemaTypeResult:
		if baseType, err := sm.ResultSchema(); err == nil {
			schema = baseType
		}
		sm.SchemaModify(schemaType, schema)

	case SchemaTypeResultList:
		if listType, err := sm.ResultListSchema(); err == nil {
			schema = listType
		}
		sm.SchemaModify(schemaType, schema.Items)

	}

	for property, examples := range sm.Examples {
		if _, found := schema.Properties[property]; found {
			schema.Properties[property].Examples = examples
		}
	}
	return schema
}

func makeModelSchemaList(schemaType SchemaType) (schemaList []*jsonschema.Schema) {
	schemaList = []*jsonschema.Schema{}
	for _, model := range getModelNames(schemaType) {
		schemaList = append(schemaList, makeModelSchema(model, schemaType))
	}
	return schemaList
}

func getModelNames(schemaType SchemaType) (models []string) {
	models = []string{}
	for model := range getSchemaMap() {
		if schemaType == SchemaTypeInput ||
			(schemaType != SchemaTypeInput && !slices.Contains([]string{"contact", "address", "event"}, model)) {
			models = append(models, model)
		}
	}
	return models
}

func getModelEnum(schemaType SchemaType) (models []any) {
	models = []any{}
	for _, model := range getModelNames(schemaType) {
		models = append(models, model)
	}
	return models
}

func getModelSchemaByPrefix(prefix string) (ms *ModelSchema, err error) {
	for _, model := range getModelNames(SchemaTypeResult) {
		if prefix == getSchemaMap()[model].Prefix {
			return getSchemaMap()[model], nil
		}
	}
	return nil, fmt.Errorf("invalid model prefix: %s", prefix)
}

func getSchemaData(data cu.IM, ms *ModelSchema) (modelData, metaData any, inputFields []string, metaFields []string, err error) {

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
