package main

import (
	"GenFromStruct/GenFromPackage/generator"
	"GenFromStruct/GenFromPackage/parser"
	"GenFromStruct/GenFromPackage/processor"
)

func main() {
	// ...
	sourceType := "GenFromStruct/sampleData/mypack.Project_ToG"
	//cwd, err := os.Getwd()
	//if err != nil {
	//	log.Fatal(err)
	//}
	cwd := "D:/GolandProjects/GenFromStruct/sampleData/mypack"

	si := parser.Parse(sourceType, cwd)

	psi := processor.ProcessTag(si)
	generator.GenerateModel(psi)
	generator.GenerateModelPtr(psi)
	generator.GenerateDbOps(psi)
}
