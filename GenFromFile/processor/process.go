package processor

import (
	"GenFromStruct/GenFromFile/model"
	"GenFromStruct/utils"
)

func ProcessTag(sfi model.StructFileInfo) (psfi model.StructFileInfo) {
	psfi = model.StructFileInfo{
		OriginFileName: sfi.OriginFileName,
		FileName:       sfi.FileName,
		PackageName:    sfi.PackageName,
		Imports:        sfi.Imports,
		Structs:        make([]model.StructInfo, len(sfi.Structs)),
	}
	for s := range sfi.Structs {
		psfi.Structs[s] = model.StructInfo{
			ClassName: sfi.Structs[s].ClassName,
			TableName: utils.CamelToSnake(sfi.Structs[s].ClassName),
			Fields:    make([]model.StructField, len(sfi.Structs[s].Fields)),
		}
		for f := range sfi.Structs[s].Fields {
			psfi.Structs[s].Fields[f] = model.StructField{
				Name:    sfi.Structs[s].Fields[f].Name,
				Type:    sfi.Structs[s].Fields[f].Type,
				TagMap:  make(map[string]string),
				Comment: sfi.Structs[s].Fields[f].Comment,
			}

			fieldName := utils.CamelToSnake(sfi.Structs[s].Fields[f].Name)
			fieldPgType := utils.go2Pg[sfi.Structs[s].Fields[f].Type]

			for k, v := range sfi.Structs[s].Fields[f].TagMap {
				psfi.Structs[s].Fields[f].TagMap[k] = v
			}

			if len(psfi.Structs[s].Fields[f].TagMap["db"]) <= 0 {
				psfi.Structs[s].Fields[f].TagMap["db"] = fieldName
			} else {
				psfi.Structs[s].Fields[f].TagMap["db"] = psfi.Structs[s].Fields[f].TagMap["db"]
			}

			if len(psfi.Structs[s].Fields[f].TagMap["json"]) <= 0 {
				psfi.Structs[s].Fields[f].TagMap["json"] = fieldName
			} else {
				psfi.Structs[s].Fields[f].TagMap["json"] = psfi.Structs[s].Fields[f].TagMap["json"]
			}

			if len(psfi.Structs[s].Fields[f].TagMap["pgType"]) <= 0 {
				psfi.Structs[s].Fields[f].TagMap["pgType"] = fieldPgType
			} else {
				psfi.Structs[s].Fields[f].TagMap["pgType"] = psfi.Structs[s].Fields[f].TagMap["pgType"]
			}

		}
	}
	return psfi
}
