package table

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaOptionalColumn(it map[string]*schema.Schema) *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: it,
		},
	}
}

func schemaComputedColumn(it map[string]*schema.Schema) *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: it,
		},
	}
}

func Resource() *schema.Resource {
	return &schema.Resource{
		CreateContext: create,
		DeleteContext: delete,
		ReadContext:   read,
		UpdateContext: update,
		Schema: map[string]*schema.Schema{
			// TODO: follow the rest of the spec, add missing fields
			// https://dev.mysql.com/doc/refman/8.0/en/create-table.html
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true, // TODO We could technically get around this
				// https://stackoverflow.com/questions/67093/how-do-i-quickly-rename-a-mysql-database-change-schema-name
			},
			"binary":  schemaOptionalColumn(schemaBinary()),
			"bit":     schemaOptionalColumn(schemaBit()),
			"blob":    schemaOptionalColumn(schemaBlob()),
			"boolean": schemaOptionalColumn(schemaBoolean()),
			"char":    schemaOptionalColumn(schemaChar()),
			"enum":    schemaOptionalColumn(schemaEnum()),
			"float":   schemaOptionalColumn(schemaFloat()),
			"int":     schemaOptionalColumn(schemaInteger()),
			"set":     schemaOptionalColumn(schemaSet()),
			"text":    schemaOptionalColumn(schemaText()),
			// TODO "time":    {}
			// 		date
			// 		time
			// 		datetime
			// 		timestamp
			// 		year
		},
	}
}

func Data() *schema.Resource {
	return &schema.Resource{
		ReadContext: read,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"binary":  schemaComputedColumn(schemaBinary()),
			"bit":     schemaComputedColumn(schemaBit()),
			"blob":    schemaComputedColumn(schemaBlob()),
			"boolean": schemaComputedColumn(schemaBoolean()),
			"char":    schemaComputedColumn(schemaChar()),
			"enum":    schemaComputedColumn(schemaEnum()),
			"float":   schemaComputedColumn(schemaFloat()),
			"int":     schemaComputedColumn(schemaInteger()),
			"set":     schemaComputedColumn(schemaSet()),
			"text":    schemaComputedColumn(schemaText()),
		},
	}
}
