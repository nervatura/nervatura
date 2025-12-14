package mcp

import (
	"context"
	"errors"
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

type ModelSchemaInterface interface {
	CreateInputSchema(scope string) *jsonschema.Schema
	UpdateInputSchema(scope string) *jsonschema.Schema
	QueryInputSchema(scope string) *jsonschema.Schema
	QueryOutputSchema(scope string) *jsonschema.Schema
}

type ModelSchema struct {
	Name              string
	Prefix            string
	CustomFrom        string
	CustomParameters  cu.IM
	CustomFilter      string
	CreateInputSchema func(scope string) *jsonschema.Schema
	UpdateInputSchema func(scope string) *jsonschema.Schema
	QueryInputSchema  func(scope string) *jsonschema.Schema
	QueryOutputSchema func(scope string) *jsonschema.Schema
	LoadData          func(data any) (modelData, metaData any, err error)
	LoadList          func(rows []cu.IM) (items any, err error)
	Validate          func(ctx context.Context, input cu.IM) (data cu.IM, err error)
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

func getModelSchemaByPrefix(prefix string) (ms *ModelSchema, err error) {
	for _, td := range toolDataMap {
		if !td.Extend && td.ModelSchema != nil && prefix == td.ModelSchema.Prefix {
			return td.ModelSchema, nil
		}
	}
	return nil, errors.New("invalid model prefix: " + prefix)
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
