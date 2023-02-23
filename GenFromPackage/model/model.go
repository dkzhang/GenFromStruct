package model

import "go/types"

type StructField struct {
	Name   string
	Type   types.Type
	TagMap map[string]string
}

type StructInfo struct {
	StructName string
	TableName  string
	Fields     []StructField
}
