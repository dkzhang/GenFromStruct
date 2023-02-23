package model

import "go/types"

type StructField struct {
	Name    string
	Type    types.Type
	TagMap  map[string]string
	Comment string
}

type StructInfo struct {
	ClassName string
	TableName string
	Fields    []StructField
}
