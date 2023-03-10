package generator

import (
	"GenFromStruct/GenFromPackage/model"
	"GenFromStruct/utils"
	"bytes"
	"fmt"
	. "github.com/dave/jennifer/jen"
	"go/format"
	"go/types"
	"log"
	"os"
	"path/filepath"
)

func GenerateModel(si model.StructInfo) error {

	modelPath := filepath.Join(si.FilePath, "model")
	modelFileName := filepath.Join(modelPath, utils.Camel2camel(si.StructName)+".go")

	// make path if not exist
	_, err := os.Stat(modelPath)
	if os.IsNotExist(err) {
		os.Mkdir(modelPath, os.ModePerm)
	}

	// create model file
	f := NewFile("model")
	f.PackageComment("Package model Code generated by generator.")

	// Generate struct
	err = genStruct(si, f)
	if err != nil {
		return fmt.Errorf("error generating struct: %v", err)
	}

	// Write struct to file
	err = f.Save(modelFileName)
	if err != nil {
		return fmt.Errorf("error saving model file: %v", err)
	}

	/////////////////////////////////////////////////////////////////////////////////////
	// Generate schema
	buf := genSchema(si)

	// append schema to file
	file, err := os.OpenFile(modelFileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening model file: %v", err)
	}
	defer file.Close()

	// format the output
	src, err := format.Source(buf.Bytes())
	if err != nil {
		log.Printf("formatting source: %s", err)
		src = buf.Bytes()
	}
	if _, err = file.Write(src); err != nil {
		return fmt.Errorf("error writing schema to model file: %v", err)
	}

	/////////////////////////////////////////////////////////////////////////////////////
	// Generate FieldMap
	buf = genFieldMap(si)
	src, err = format.Source(buf.Bytes())
	if err != nil {
		log.Printf("formatting source: %s", err)
		src = buf.Bytes()
	}
	if _, err = file.Write(src); err != nil {
		return fmt.Errorf("error writing schema to model file: %v", err)
	}

	return nil
}

func genSchema(si model.StructInfo) bytes.Buffer {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, " \n")
	fmt.Fprintf(&buf, "const Schema%s = `\n", si.StructName)
	// create table
	fmt.Fprintf(&buf, "\t\tCREATE TABLE %s ( \n", si.TableName)

	// write fields
	for _, f := range si.Fields {
		fmt.Fprintf(&buf, "\t\t\t%s\t%s", f.TagMap["db"], f.TagMap["pgType"])

		if pgA, ok := f.TagMap["pgA"]; ok {
			fmt.Fprintf(&buf, " %s", pgA)
		}
		// end of line
		fmt.Fprintf(&buf, ", \n")
	}
	// end of create table
	fmt.Fprintf(&buf, "); \n")
	// end of schema
	fmt.Fprintf(&buf, "` \n\n")
	return buf
}

func genStruct(si model.StructInfo, f *File) error {
	structFields := make([]Code, len(si.Fields))

	// Iterate over struct fields
	for i := 0; i < len(si.Fields); i++ {
		field := si.Fields[i]

		// Generate code for each changeset field
		code := Id(field.Name)
		switch v := field.Type.(type) {
		case *types.Basic:
			code.Id(v.String())
		case *types.Named:
			typeName := v.Obj()
			// Qual automatically imports packages
			code.Qual(
				typeName.Pkg().Path(),
				typeName.Name(),
			)
		default:
			return fmt.Errorf("struct field type not hanled: %T", v)
		}

		tags := make(map[string]string, 2)
		tags["json"] = field.TagMap["json"]
		tags["db"] = field.TagMap["db"]
		code.Tag(tags)

		structFields[i] = code
	}

	// Generate struct type
	f.Type().Id(si.StructName).Struct(structFields...)

	return nil
}

func genFieldMap(si model.StructInfo) bytes.Buffer {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, " \n")
	fmt.Fprintf(&buf, "var FieldMap%s = map[string]string{ \n", si.StructName)
	for _, f := range si.Fields {
		fieldType := ""
		switch v := f.Type.(type) {
		case *types.Basic:
			fieldType = v.String()
		case *types.Named:
			typeName := v.Obj()
			fieldType = typeName.Pkg().Path() + "." + typeName.Name()
		default:
			log.Fatalf("struct field type not hanled: %T", v)
		}

		fmt.Fprintf(&buf, "\t\t\t\"%s\": \"%s\", \n", f.TagMap["json"], fieldType)
	}
	fmt.Fprintf(&buf, "} \n\n")
	return buf
}
