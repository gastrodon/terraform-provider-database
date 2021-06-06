package database

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type statementKind string

const (
	ALTER  statementKind = "ALTER"
	CREATE statementKind = "CREATE"
	DELETE statementKind = "DELETE"
)

func resourceTable() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTableCreate,
		DeleteContext: resourceTableDelete,
		ReadContext:   resourceTableRead,
		UpdateContext: resourceTableUpdate,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true, // TODO We could technically get around this
				// https://stackoverflow.com/questions/67093/how-do-i-quickly-rename-a-mysql-database-change-schema-name
			},
			"column": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"null": {
							Type:     schema.TypeString,
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

func resourceTableCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	statement := statementTable(
		CREATE,
		data,
	)

	log.Println("executing " + statement)

	if err := meta.(Config).Connection.Exec(statement).Error; err != nil {
		return diag.FromErr(err)
	}

	data.SetId("table_" + data.Get("name").(string))
	return resourceTableRead(ctx, data, meta)
}

func resourceTableUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceTableRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return dataTableRead(ctx, data, meta)
}

func resourceTableDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func formatColumn(it map[string]interface{}) string {
	parts := []string{it["field"].(string), it["type"].(string)}

	return strings.Join(parts, " ")
}

func statementTable(kind statementKind, data *schema.ResourceData) string {
	columnsRaw := data.Get("column").(*schema.Set)
	serialized := make([]string, columnsRaw.Len())
	for index, it := range columnsRaw.List() {
		serialized[index] = formatColumn(it.(map[string]interface{}))
	}

	return fmt.Sprintf(
		"%s TABLE %s (%s)",
		kind,
		data.Get("name"),
		strings.Join(serialized, ", "),
	)
}
