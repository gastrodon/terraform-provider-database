package table

import (
	"hash/fnv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func stringItemHash(it interface{}) int {
	computed := fnv.New32a()
	computed.Write([]byte(it.(string)))
	return int(computed.Sum32())
}

var (
	cases = []struct {
		Column string
		Kind   string
		Data   map[string]interface{}
	}{
		{
			"binlike BINARY(12) NULL",
			"binary",
			map[string]interface{}{
				"name":     "binlike",
				"size":     12,
				"variable": false,
				"nullable": true,
			},
		},
		{
			"binlike BINARY(21) NOT NULL",
			"binary",
			map[string]interface{}{
				"name":     "binlike",
				"variable": false,
				"size":     21,
				"nullable": false,
			},
		},
		{
			"binlike VARBINARY(13) NULL",
			"binary",
			map[string]interface{}{
				"name":     "binlike",
				"variable": true,
				"size":     13,
				"nullable": true,
			},
		},
		{
			"binlike VARBINARY(31) NOT NULL",
			"binary",
			map[string]interface{}{
				"name":     "binlike",
				"variable": true,
				"nullable": false,
				"size":     31,
			},
		},

		// Default havers
		{
			"binlike BINARY(51) NULL DEFAULT ('42069')",
			"binary",
			map[string]interface{}{
				"name":     "binlike",
				"size":     51,
				"variable": false,
				"nullable": true,
				"default":  42069,
			},
		},
		{
			"bitlike BIT(52) NULL DEFAULT ('69420')",
			"bit",
			map[string]interface{}{
				"name":     "bitlike",
				"size":     52,
				"nullable": true,
				"default":  69420,
			},
		},
		{
			"bloblike BLOB NULL DEFAULT ('say hello to bengis!')",
			"blob",
			map[string]interface{}{
				"name":     "bloblike",
				"kind":     "BLOB",
				"nullable": true,
				"default":  "say hello to bengis!",
			},
		},
		{
			"boollike BOOLEAN NOT NULL DEFAULT ('true')",
			"boolean",
			map[string]interface{}{
				"name":     "boollike",
				"nullable": false,
				"default":  true,
			},
		},
		{
			"enumlike ENUM('bar', 'foo') NULL DEFAULT ('foo')",
			"enum",
			map[string]interface{}{
				"name":     "enumlike",
				"nullable": true,
				"allowed":  *schema.NewSet(stringItemHash, []interface{}{"foo", "bar"}),
				"default":  "foo",
			},
		},
		{
			"floatlike FLOAT NULL DEFAULT ('3.141500')",
			"float",
			map[string]interface{}{
				"name":     "floatlike",
				"nullable": true,
				"kind":     "FLOAT",
				"default":  3.141500,
			},
		},
		{
			"setlike SET('baz', 'bar', 'foo') NULL DEFAULT ('bar,foo')",
			"set",
			map[string]interface{}{
				"name":     "setlike",
				"nullable": true,
				"allowed":  *schema.NewSet(stringItemHash, []interface{}{"bar", "baz", "foo"}),
				"default":  *schema.NewSet(stringItemHash, []interface{}{"foo", "bar"}),
			},
		},

		{
			"bitlike BIT(24) NULL",
			"bit",
			map[string]interface{}{
				"name":     "bitlike",
				"size":     24,
				"nullable": true,
			},
		},
		{
			"bitlike BIT(1) NOT NULL",
			"bit",
			map[string]interface{}{
				"name":     "bitlike",
				"size":     1,
				"nullable": false,
			},
		},

		{
			"bloblike BLOB NOT NULL",
			"blob",
			map[string]interface{}{
				"name":     "bloblike",
				"kind":     "BLOB",
				"nullable": false,
			},
		},
		{
			"bloblike BLOB NULL",
			"blob",
			map[string]interface{}{
				"name":     "bloblike",
				"kind":     "BLOB",
				"nullable": true,
			},
		},
		{
			"bloblike SMALLBLOB NOT NULL",
			"blob",
			map[string]interface{}{
				"name": "bloblike",
				"kind": "SMALLBLOB",
			},
		},
		{
			"bloblike MEDIUMBLOB NOT NULL",
			"blob",
			map[string]interface{}{
				"name": "bloblike",
				"kind": "MEDIUMBLOB",
			},
		},
		{
			"bloblike LONGBLOB NOT NULL",
			"blob",
			map[string]interface{}{
				"name": "bloblike",
				"kind": "LONGBLOB",
			},
		},

		{
			"boollike BOOLEAN NULL",
			"boolean",
			map[string]interface{}{
				"name":     "boollike",
				"nullable": true,
			},
		},
		{
			"boollike BOOLEAN NOT NULL",
			"boolean",
			map[string]interface{}{
				"name":     "boollike",
				"nullable": false,
			},
		},

		{
			"charlike CHAR(51) NOT NULL",
			"char",
			map[string]interface{}{
				"name":     "charlike",
				"nullable": false,
				"variable": false,
				"size":     51,
			},
		},
		{
			"charlike VARCHAR(15) NOT NULL",
			"char",
			map[string]interface{}{
				"name":     "charlike",
				"nullable": false,
				"variable": true,
				"size":     15,
			},
		},

		// NOTE: because of how the collection is implemented.
		// the ENUM items are backwards
		{
			"enumlike ENUM('bar', 'foo') NOT NULL",
			"enum",
			map[string]interface{}{
				"name":     "enumlike",
				"nullable": false,
				"allowed":  *schema.NewSet(stringItemHash, []interface{}{"foo", "bar"}),
			},
		},
		{
			"enumlike ENUM() NOT NULL",
			"enum",
			map[string]interface{}{
				"name":     "enumlike",
				"nullable": false,
				"allowed":  *schema.NewSet(stringItemHash, nil),
			},
		},
		{
			"enumlike ENUM('baz', 'bar') NOT NULL",
			"enum",
			map[string]interface{}{
				"name":     "enumlike",
				"nullable": false,
				"allowed":  *schema.NewSet(stringItemHash, []interface{}{"bar", "baz", "bar"}),
			},
		},

		{
			"floatlike FLOAT NOT NULL",
			"float",
			map[string]interface{}{
				"name":     "floatlike",
				"nullable": false,
				"kind":     "FLOAT",
			},
		},
		{
			"floatlike DOUBLE NOT NULL",
			"float",
			map[string]interface{}{
				"name":     "floatlike",
				"nullable": false,
				"kind":     "DOUBLE",
			},
		},
		{
			"floatlike DECIMAL NOT NULL",
			"float",
			map[string]interface{}{
				"name":     "floatlike",
				"nullable": false,
				"kind":     "DECIMAL",
			},
		},

		{
			"intlike INT NOT NULL",
			"int",
			map[string]interface{}{
				"name":     "intlike",
				"nullable": false,
				"kind":     "INT",
				"signed":   true,
			},
		},
		{
			"intlike INT UNSIGNED NOT NULL",
			"int",
			map[string]interface{}{
				"name":     "intlike",
				"nullable": false,
				"kind":     "INT",
				"signed":   false,
			},
		},
		{
			"intlike INT NOT NULL AUTO_INCREMENT",
			"int",
			map[string]interface{}{
				"name":           "intlike",
				"nullable":       false,
				"kind":           "INT",
				"auto_increment": true,
				"signed":         true,
			},
		},
		{
			"intlike INT UNSIGNED NOT NULL AUTO_INCREMENT",
			"int",
			map[string]interface{}{
				"name":           "intlike",
				"nullable":       false,
				"kind":           "INT",
				"auto_increment": true,
				"signed":         false,
			},
		},
		{
			"intlike TINYINT NOT NULL",
			"int",
			map[string]interface{}{
				"name":           "intlike",
				"nullable":       false,
				"kind":           "TINYINT",
				"auto_increment": false,
				"signed":         true,
			},
		},
		{
			"intlike TINYINT UNSIGNED NOT NULL",
			"int",
			map[string]interface{}{
				"name":           "intlike",
				"nullable":       false,
				"kind":           "TINYINT",
				"auto_increment": false,
				"signed":         false,
			},
		},
		{
			"intlike MEDIUMINT NOT NULL",
			"int",
			map[string]interface{}{
				"name":           "intlike",
				"nullable":       false,
				"kind":           "MEDIUMINT",
				"auto_increment": false,
				"signed":         true,
			},
		},
		{
			"intlike MEDIUMINT UNSIGNED NOT NULL",
			"int",
			map[string]interface{}{
				"name":           "intlike",
				"nullable":       false,
				"kind":           "MEDIUMINT",
				"auto_increment": false,
				"signed":         false,
			},
		},
		{
			"intlike BIGINT NOT NULL",
			"int",
			map[string]interface{}{
				"name":           "intlike",
				"nullable":       false,
				"kind":           "BIGINT",
				"auto_increment": false,
				"signed":         true,
			},
		},
		{
			"intlike BIGINT UNSIGNED NOT NULL",
			"int",
			map[string]interface{}{
				"name":           "intlike",
				"nullable":       false,
				"kind":           "BIGINT",
				"auto_increment": false,
				"signed":         false,
			},
		},

		{
			"setlike SET('bar', 'foo') NOT NULL",
			"set",
			map[string]interface{}{
				"name":     "setlike",
				"nullable": false,
				"allowed":  *schema.NewSet(stringItemHash, []interface{}{"bar", "foo"}),
			},
		},
		{
			"setlike SET() NOT NULL",
			"set",
			map[string]interface{}{
				"name":     "setlike",
				"nullable": false,
				"allowed":  *schema.NewSet(stringItemHash, nil),
			},
		},
		{
			"setlike SET('bar', 'foo') NOT NULL",
			"set",
			map[string]interface{}{
				"name":     "setlike",
				"nullable": false,
				"allowed":  *schema.NewSet(stringItemHash, []interface{}{"bar", "foo", "foo", "bar"}),
			},
		},

		{
			"textlike TEXT NOT NULL",
			"text",
			map[string]interface{}{
				"name":     "textlike",
				"nullable": false,
				"kind":     "TEXT",
			},
		},
		{
			"textlike TINYTEXT NOT NULL",
			"text",
			map[string]interface{}{
				"name":     "textlike",
				"nullable": false,
				"kind":     "TINYTEXT",
			},
		},
		{
			"textlike MEDIUMTEXT NOT NULL",
			"text",
			map[string]interface{}{
				"name":     "textlike",
				"nullable": false,
				"kind":     "MEDIUMTEXT",
			},
		},
		{
			"textlike BIGTEXT NOT NULL",
			"text",
			map[string]interface{}{
				"name":     "textlike",
				"nullable": false,
				"kind":     "BIGTEXT",
			},
		},
	}
)

func Test_columnStatement_generation(test *testing.T) {
	for index, single := range cases {
		generated := columnStatement(single.Kind, single.Data)

		if generated != single.Column {
			test.Fatalf(
				"Incorrect column at %d!\n have: %s \n want: %s ",
				index,
				generated,
				single.Column,
			)
		}
	}
}
