package processor

import (
	"GenFromStruct/GenFromPackage/model"
	"GenFromStruct/utils"
)

func ProcessTag(si model.StructInfo) (psi model.StructInfo) {
	psi = model.StructInfo{
		PackagePath: si.PackagePath,
		FilePath:    si.FilePath,
		StructName:  si.StructName,
		TableName:   utils.CamelToSnake(si.StructName),
		Fields:      make([]model.StructField, len(si.Fields)),
	}

	for f := range si.Fields {
		psi.Fields[f] = model.StructField{
			Name:   si.Fields[f].Name,
			Type:   si.Fields[f].Type,
			TagMap: make(map[string]string),
		}
		fieldName := utils.CamelToSnake(si.Fields[f].Name)
		fieldPgType := go2Pg[si.Fields[f].Type.String()]

		for k, v := range si.Fields[f].TagMap {
			psi.Fields[f].TagMap[k] = v
		}

		if len(psi.Fields[f].TagMap["db"]) <= 0 {
			psi.Fields[f].TagMap["db"] = fieldName
		} else {
			psi.Fields[f].TagMap["db"] = psi.Fields[f].TagMap["db"]
		}

		if len(psi.Fields[f].TagMap["json"]) <= 0 {
			psi.Fields[f].TagMap["json"] = fieldName
		} else {
			psi.Fields[f].TagMap["json"] = psi.Fields[f].TagMap["json"]
		}

		if len(psi.Fields[f].TagMap["pgType"]) <= 0 {
			psi.Fields[f].TagMap["pgType"] = fieldPgType
		} else {
			psi.Fields[f].TagMap["pgType"] = psi.Fields[f].TagMap["pgType"]
		}

	}
	return psi
}
