package plugin

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type SchemaField struct {
	DataType   string
	Nullable   string
	ColumnName string
}

func GenerateContentSchema(Name string, Column []TableColumn) string {
	modelSchema := map[string]map[string]SchemaField{}
	modelField := map[string]SchemaField{}

	modelSchemaGeneration := make(map[string]string)

	for _, column := range Column {
		modelField[column.ColumnName] = SchemaField{
			DataType:   MappingType[column.DataType],
			ColumnName: column.ColumnName,
			Nullable: func() string {
				if column.IsNullable == "NO" {
					return "!"
				}
				return ""
			}(),
		}
	}

	modelSchema[Name] = modelField

	for key, val := range modelSchema {
		content := `type %s {
{{s}}
}`

		content = fmt.Sprintf(content, key)

		content = strings.Replace(content, "{{s}}", "%s", -1)
		contentField := make([]string, 0)
		for _, val2 := range val {
			contentField = append(contentField, fmt.Sprintf(" %s: %s%s ", val2.ColumnName, val2.DataType, val2.Nullable))
		}
		content = fmt.Sprintf(content, strings.Join(contentField, "\n"))

		modelSchemaGeneration[key] = content
	}

	content := ""
	for _, val := range modelSchemaGeneration {
		content = content + val
		content = content + "\n"
	}
	return content
}

func GenerateSchemaFile(content string, fileName string) {
	err := ioutil.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println(fmt.Sprintf("File '%s' created successfully.", fileName))
}
