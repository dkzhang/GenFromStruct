package model

type StructField struct {
	Name    string
	Type    string
	TagMap  map[string]string
	Comment string
}

type StructInfo struct {
	ClassName string
	TableName string
	Fields    []StructField
}

type StructFileInfo struct {
	OriginFileName string
	FileName       string
	PackageName    string
	Imports        []string
	Structs        []StructInfo
}
