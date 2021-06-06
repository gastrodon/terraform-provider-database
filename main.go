package main

import (
	"log"

	"github.com/gastrodon/terraform-provider-database/database"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	log.SetFlags(0)
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: database.Provider})
}
