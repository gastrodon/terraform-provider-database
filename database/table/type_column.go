package table

import (
	"strings"
)

type column struct {
	Field   string
	Type    string
	Null    string
	Key     string
	Default string
	Extra   string
}

func (it column) Map() map[string]interface{} {
	values := map[string]interface{}{
		"field":          it.Field,
		"type":           it.Type,
		"default":        nil,
		"primary":        false,
		"auto_increment": false,
		"nullable":       true,
	}

	if it.Null == "NO" {
		values["null"] = false
	}

	if it.Default != "NULL" {
		// TODO There's probably a safer way to do this
		values["default"] = it.Default
	}

	// TODO foreign key and other keys ( idk them right now )
	if it.Key != "" {
		switch it.Key {
		case "PRI":
			values["primary"] = true
		}
	}

	for _, extra := range strings.Split(it.Extra, " ") {
		switch extra {
		case "auto_increment":
			values["auto_increment"] = true
			break
		}
	}

	return values
}

func (it column) String() string {
	return columnMapString(it.Map())
}

func columnMapString(it map[string]interface{}) string {
	parts := []string{it["field"].(string), it["type"].(string)}

	if value, ok := it["default"].(string); ok {
		parts = append(parts, "DEFAULT '"+value+"'")
	}

	if value, ok := it["primary"].(bool); value && ok {
		parts = append(parts, "PRIMARY KEY")
	}

	if value, ok := it["auto_increment"].(bool); value && ok {
		parts = append(parts, "AUTO_INCREMENT")
	}

	if value, ok := it["nullable"].(bool); !value && ok {
		parts = append(parts, "NOT NULL")
	}

	return strings.Join(parts, " ")
}

func columnMaps(columns []column) []map[string]interface{} {
	values := make([]map[string]interface{}, len(columns))
	for index, it := range columns {
		values[index] = it.Map()
	}

	return values
}

func columnStrings(columns []column) []string {
	values := make([]string, len(columns))
	for index, it := range columns {
		values[index] = it.String()
	}

	return values
}
