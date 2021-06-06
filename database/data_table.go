package database

import (
	"context"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gorm.io/gorm"
)

const (
	TABLE_NOT_EXIST = "Error 1146:"
)

type tableColumn struct {
	Field   string
	Type    string
	Null    string
	Key     string
	Default string
	Extra   string
}

func (it tableColumn) Map() map[string]interface{} {
	var null bool
	if it.Null == "YES" {
		null = true
	} else {
		null = false
	}

	var defaultValue *string
	if it.Default == "NULL" {
		defaultValue = nil
	} else {
		defaultValue = &it.Default
	}

	var extra *string
	if it.Extra == "" {
		extra = nil
	} else {
		extra = &it.Extra
	}

	return map[string]interface{}{
		"field":   it.Field,
		"type":    it.Type,
		"null":    null,
		"key":     it.Key,
		"default": defaultValue,
		"extra":   extra,
	}
}

func tableColumnMaps(columns []tableColumn) []map[string]interface{} {
	maps := make([]map[string]interface{}, len(columns))
	for index, it := range columns {
		maps[index] = it.Map()
	}

	return maps
}

func dataTable() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataTableRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"columns": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"null": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"default": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"extra": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataTableRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := data.Get("name").(string)
	log.Println("execute DESCRIBE " + name)

	var columns []tableColumn
	if err := meta.(Config).Connection.Raw("DESCRIBE " + name).Scan(&columns).Error; err != nil {
		if err == gorm.ErrRecordNotFound || strings.HasPrefix(err.Error(), TABLE_NOT_EXIST) {
			log.Printf("Table %s doesn't exist\n", name)
			data.SetId("")
			return nil
		}

		log.Println("got another error: ", err)
		return diag.FromErr(err)
	}

	log.Printf("DESCRIBE %s: %#v", name, columns)

	data.Set("columns", tableColumnMaps(columns))
	data.SetId("table_" + name)
	return nil
}
