package table

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var (
	KINDS_INTEGER = []string{
		"BIGINT",
		"INT",
		"MEDIUMINT",
		"SMALLINT",
		"TINYINT",
	}
	KINDS_FLOAT = []string{
		"FLOAT",
		"DOUBLE",
		"DECIMAL",
	}
	KINDS_BLOB = []string{
		"LONGBLOB",
		"BLOB",
		"MEDIUMBLOB",
		"TINYBLOB",
	}
	KINDS_TEXT = []string{
		"LONGTEXT",
		"TEXT",
		"MEDIUMTEXT",
		"TINYTEXT",
	}
)

func schemaBinary() map[string]*schema.Schema {
	return merge(
		schemaColumn(),
		schemaDefaultString(),
		schemaSized(255),
		schemaVariable(),
	)
}

func schemaBit() map[string]*schema.Schema {
	it := map[string]*schema.Schema{
		"size": {
			Type:         schema.TypeInt,
			Required:     true,
			ValidateFunc: validation.IntBetween(1, 64),
		},
	}

	return merge(
		schemaColumn(),
		schemaDefaultInt(),
		it,
	)
}

func schemaBlob() map[string]*schema.Schema {
	return merge(
		schemaColumn(),
		schemaDefaultString(),
		schemaKinds("BLOB", KINDS_BLOB),
	)
}

func schemaBoolean() map[string]*schema.Schema {
	return merge(
		schemaColumn(),
		schemaDefaultBool(),
	)
}

func schemaChar() map[string]*schema.Schema {
	return merge(
		schemaColumn(),
		schemaDefaultString(),
		schemaSized(255),
		schemaVariable(),
	)
}

func schemaEnum() map[string]*schema.Schema {
	return merge(
		schemaColumn(),
		schemaDefaultString(),
		schemaAllowedSet(schema.TypeString),
	)
}

func schemaFloat() map[string]*schema.Schema {
	return merge(
		schemaColumn(),
		schemaDefaultFloat(),
		schemaKinds("FLOAT", KINDS_FLOAT),
	)
}

func schemaInteger() map[string]*schema.Schema {
	it := map[string]*schema.Schema{
		"auto_increment": {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		"signed": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
	}

	return merge(
		schemaColumn(),
		schemaDefaultInt(),
		schemaKinds("INT", KINDS_INTEGER),
		it,
	)
}

func schemaSet() map[string]*schema.Schema {
	return merge(
		schemaColumn(),
		schemaDefaultSet(schema.TypeString),
		schemaAllowedSet(schema.TypeString),
	)
}

func schemaText() map[string]*schema.Schema {
	return merge(
		schemaColumn(),
		schemaDefaultString(),
		schemaKinds("TEXT", KINDS_TEXT),
	)
}
