package table

import (
	"testing"
)

var (
	cases = []struct {
		Column string
		Kind   string
		Data   map[string]interface{}
	}{
		{
			"intlike BIGINT UNSIGNED NOT NULL",
			"int",
			map[string]interface{}{
				"auto_increment": nil,
				"default":        nil,
				"kind":           "BIGINT",
				"name":           "intlike",
				"nullable":       false,
				"primary":        nil,
				"signed":         false,
			},
		},
		{
			"intlike TINYINT UNSIGNED NOT NULL",
			"int",
			map[string]interface{}{
				"auto_increment": nil,
				"default":        nil,
				"kind":           "TINYINT",
				"name":           "intlike",
				"nullable":       false,
				"primary":        nil,
				"signed":         false,
			},
		},
		{
			"intlike TINYINT UNSIGNED NOT NULL",
			"int",
			map[string]interface{}{
				"auto_increment": false,
				"default":        nil,
				"kind":           "TINYINT",
				"name":           "intlike",
				"nullable":       false,
				"primary":        false,
				"signed":         false,
			},
		},
		{
			"intlike TINYINT UNSIGNED NOT NULL AUTO_INCREMENT",
			"int",
			map[string]interface{}{
				"auto_increment": true,
				"default":        nil,
				"kind":           "TINYINT",
				"name":           "intlike",
				"nullable":       false,
				"primary":        nil,
				"signed":         false,
			},
		},
		{
			"intlike INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY",
			"int",
			map[string]interface{}{
				"auto_increment": true,
				"default":        nil,
				"kind":           "INT",
				"name":           "intlike",
				"nullable":       false,
				"primary":        true,
				"signed":         false,
			},
		},
	}
)

func Test_columnStatement_generation(test *testing.T) {
	for index, single := range cases {
		generated := columnStatement(single.Kind, single.Data)

		if generated != single.Column {
			test.Fatalf("Incorrect column at %d!\n\t%s", index, generated)
		}
	}
}
