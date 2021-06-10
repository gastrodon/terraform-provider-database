package database

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gastrodon/terraform-provider-database/database/table"
	"github.com/gastrodon/terraform-provider-database/database/types"
)

func configure(data *schema.ResourceData) (interface{}, error) {
	var connectionString string
	log.Printf(
		"connecting to database/table %s://%s/%s\n",
		data.Get("protocol"),
		data.Get("address"),
		data.Get("database"),
	)

	if _, ok := data.GetOk("password"); ok {
		connectionString = fmt.Sprintf(
			"%s:%s@%s(%s)/%s",
			data.Get("username"),
			data.Get("password"),
			data.Get("protocol"),
			data.Get("address"),
			data.Get("database"),
		)
	} else {
		connectionString = fmt.Sprintf(
			"%s@%s(%s)/%s",
			data.Get("username"),
			data.Get("protocol"),
			data.Get("address"),
			data.Get("database"),
		)
	}

	connection, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return types.Config{connection}, nil
}

func Provider() *schema.Provider {
	return &schema.Provider{
		ConfigureFunc: configure,
		ResourcesMap: map[string]*schema.Resource{
			"database_table": table.Resource(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"database_table": table.Data(),
		},
		Schema: map[string]*schema.Schema{
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
			},
			"address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"database": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
