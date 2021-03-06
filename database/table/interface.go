package table

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gorm.io/gorm"

	"github.com/gastrodon/terraform-provider-database/database/types"
)

const (
	TABLE_NOT_EXIST = "Error 1146:"
)

var (
	columnKinds = []string{
		"binary",
		"bit",
		"blob",
		"boolean",
		"char",
		"enum",
		"float",
		"int",
		"set",
		"text",
	}
)

func handleMaybeExisted(data *schema.ResourceData, err error) diag.Diagnostics {
	if err == gorm.ErrRecordNotFound || strings.HasPrefix(err.Error(), TABLE_NOT_EXIST) {
		log.Printf("Table %s doesn't exist\n", data.Get("name"))
		data.SetId("")
		return nil
	}

	return diag.FromErr(err)
}

func createTable(data *schema.ResourceData) string {
	size := 0
	for _, kind := range columnKinds {
		size += data.Get(kind).(*schema.Set).Len()
	}

	index := 0
	columns := make([]string, size)
	for _, kind := range columnKinds {
		for _, raw := range data.Get(kind).(*schema.Set).List() {
			columns[index] = columnStatement(kind, raw.(map[string]interface{}))
			index++
		}
	}

	statement := fmt.Sprintf(
		"CREATE TABLE %s (%s)",
		data.Get("name"),
		strings.Join(columns, ", "),
	)

	log.Println("generated sql " + statement)
	return statement
}

func dropTable(name string) string {
	statement := "DROP TABLE " + name
	log.Println("generated sql " + statement)
	return statement
}

func describeTable(name string) string {
	statement := "DESCRIBE TABLE " + name
	log.Println("generated sql " + statement)
	return statement
}

func create(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if err := meta.(types.Config).Connection.Exec(createTable(data)).Error; err != nil {
		return diag.FromErr(err)
	}

	return read(ctx, data, meta)
}

func delete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if err := meta.(types.Config).Connection.Exec(dropTable(data.Get("name").(string))).Error; err != nil {
		handleMaybeExisted(data, err)
	}

	return read(ctx, data, meta)
}

func read(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := data.Get("name").(string)
	columns := *new([]column)
	if err := meta.(types.Config).Connection.Raw(describeTable(name)).Scan(&columns).Error; err != nil {
		handleMaybeExisted(data, err)
	}

	data.Set("columns", schemaMaps(columns))
	data.SetId("table_" + name)
	return nil
}

func update(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}
