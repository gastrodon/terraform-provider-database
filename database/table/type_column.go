package table

import (
	"strings"
)

type column struct {
	Name    string
	Type    string
	Null    string
	Key     string
	Default string
	Extra   string
}

func (it column) SchemaMap() map[string]interface{} {
	switch strings.ToUpper(it.Type) {
	case "BINARY":
		panic("BINARY")
	case "BIT":
		panic("BIT")
	case "BLOB", "TINYBLOB", "MEDIUMBLOB", "LONGBLOB":
		panic("BLOB")
	case "CHAR", "VARCHAR":
		panic("CHAR")
	case "ENUM":
		panic("ENUM")
	case "FLOAT", "DOUBLE", "DECIMAL":
		panic("FLOAT")
	case "INT", "TINYINT", "SMALLINT", "MEDIUMINT", "BIGINT":
		panic("INT")
	case "SET":
		panic("SET")
	case "TEXT", "TINYTEXT", "MEDIUMTEXT", "LONGTEXT":
		panic("TEXT")
	}

	return nil
}

func schemaMaps(them []column) []map[string]interface{} {
	maps := make([]map[string]interface{}, len(them))
	for index, it := range them {
		maps[index] = it.SchemaMap()
	}

	return maps
}
