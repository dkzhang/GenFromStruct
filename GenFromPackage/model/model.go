package model

import "go/types"

type StructField struct {
	Name   string
	Type   types.Type
	TagMap map[string]string
}

type StructInfo struct {
	PackagePath string
	FilePath    string

	StructName string
	TableName  string
	Fields     []StructField
}
