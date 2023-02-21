package parseStruct

type StructField struct {
	Name    string
	Type    string
	Tag     string
	TagMap  map[string]string
	Comment string
}

type StructInfo struct {
	Name   string
	Fields []StructField
}

type StructFileInfo struct {
	FileName string
	Structs  []StructInfo
}
