package table

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func collectAllowed(data schema.Set) string {
	collected := make([]string, data.Len())
	for index, it := range data.List() {
		collected[index] = fmt.Sprintf("'%s'", it.(string))
	}

	return strings.Join(collected, ", ")
}

func determineType(kind string, data map[string]interface{}) string {
	if explicit, ok := data["kind"].(string); ok {
		if signed, ok := data["signed"].(bool); ok && !signed {
			return fmt.Sprintf("%s UNSIGNED", explicit)
		} else {
			return explicit
		}
	}

	switch strings.ToLower(kind) {
	case "binary":
		if data["variable"].(bool) {
			return fmt.Sprintf("VARBINARY(%d)", data["size"].(int))
		} else {
			return fmt.Sprintf("BINARY(%d)", data["size"].(int))
		}
	case "char":
		if data["variable"].(bool) {
			return fmt.Sprintf("VARCHAR(%d)", data["size"].(int))
		} else {
			return fmt.Sprintf("CHAR(%d)", data["size"].(int))
		}
	case "bit":
		return fmt.Sprintf("BIT(%d)", data["size"].(int))
	case "boolean":
		return "BOOLEAN"
	case "enum":
		allowed := collectAllowed(data["allowed"].(schema.Set))
		return fmt.Sprintf("ENUM(%s)", allowed)
	case "set":
		allowed := collectAllowed(data["allowed"].(schema.Set))
		return fmt.Sprintf("SET(%s)", allowed)
	}

	panic("Can't find type for " + kind)
}

func determineDefault(data map[string]interface{}) string {
	// TODO
	return ""
}

func columnStatement(kind string, data map[string]interface{}) string {
	log.Printf("Generating column statement for %s: %#v\n", kind, data)

	parts := []string{
		data["name"].(string),
		determineType(kind, data),
	}

	if nullable, ok := data["nullable"].(bool); ok && nullable {
		parts = append(parts, "NULL")
	} else {
		parts = append(parts, "NOT NULL")
	}

	if value, ok := data["default"]; ok && value != nil {
		parts = append(parts, determineDefault(data))
	}

	if auto, ok := data["auto_increment"].(bool); ok && auto {
		parts = append(parts, "AUTO_INCREMENT")
	}

	if primary, ok := data["primary"].(bool); ok && primary {
		parts = append(parts, "PRIMARY KEY")
	}

	statement := strings.Join(parts, " ")
	log.Printf("Generated column statement %s\n", statement)
	return statement
}
