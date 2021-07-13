package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"k8s.io/gengo/parser"
	"k8s.io/gengo/types"
	"strings"
	"text/template"
)

const (
	// The package to generate merge methods for
	PKGName = "github.com/germanoeich/nirn/discord"
	// If this is present on a comment above a type, it will have merge methods generated for it
	CommentMarker = "++gen_merge"
	// Merge methods will only include fields that contain this field tag
	FieldMarker = "merge:\"true\""
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func extractMergeableFields(typearr *types.Type) []types.Member {
	var ret []types.Member
	for _, t := range typearr.Members {
		if strings.Contains(t.Tags, FieldMarker) {
			ret = append(ret, t)
		}
	}
	return ret
}

func zeroValue(f *types.Type) (v string) {
	switch f.Kind {
	case types.Slice:
		v = "nil"
	case types.Pointer, types.Interface:
		// TODO: check for non-pointers
		v = "nil"
	case types.Struct:
		v = f.Name.Name + "{}"
	case types.Alias:
		v = zeroValue(f.Underlying)
	case types.Builtin:
		switch f.Name.Name {
		case "bool":
			v = "false"
		case "int", "uint", "int8", "int16", "int32", "int64", "uint8", "uint16", "uint32", "uint64":
			v = "0"
		case "float64", "float32":
			v = "0.0"
		case "string":
			v = "\"\""
		default:
			v = "0 // ++"
		}
	default:
		v = "0  // -"
	}
	return v
}

func main() {
	builder := parser.New()
	if err := builder.AddDir(PKGName); err != nil {
		fmt.Println("unable to add nirn package to gengo-parser builder.", err)
		return
	}

	universe, err := builder.FindTypes()
	if err != nil {
		fmt.Println("unable to find types for nirn package.", err)
		return
	}

	nirn := universe.Package(PKGName)
	mergeableFields := make(map[*types.Type][]types.Member)
	for _, typeData := range nirn.Types {
		if contains(typeData.CommentLines, CommentMarker) {
			mergeableFields[typeData] = extractMergeableFields(typeData)
		}
	}

	tplVars := TemplateVariables{
		Package: nirn.Name,
	}
	for t, fields := range mergeableFields {
		var templateType = TemplateType{
			Name: t.Name.Name,
		}
		var fieldList []TemplateField
		for _, field := range fields {
			fieldList = append(fieldList, TemplateField{
				Name:      field.Name,
				ZeroValue: zeroValue(field.Type),
			})
		}
		templateType.Fields = fieldList
		tplVars.TypeList = append(tplVars.TypeList, templateType)
	}

	tpl, err := template.New("merge.gotpl").ParseFiles("generate/merge.gotpl")
	if err != nil {
		fmt.Println("Failed to load template", tpl)
		return
	}
	var b bytes.Buffer
	err = tpl.Execute(&b, tplVars)
	if err != nil {
		fmt.Println("Error when executing template", err)
		return
	}
	err = ioutil.WriteFile("discord/merge_gen.go", b.Bytes(), 0644)
	if err != nil {
		fmt.Println("Error creating file", err)
		return
	}
	fmt.Println("All done :)")
}

type TemplateVariables struct {
	Package string
	TypeList []TemplateType
}

type TemplateType struct {
	Name string
	Fields []TemplateField
}

type TemplateField struct {
	Name string
	ZeroValue string
}
