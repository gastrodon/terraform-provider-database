package table

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func merge(source ...map[string]*schema.Schema) map[string]*schema.Schema {
	size := 0
	for _, it := range source {
		size += len(it)
	}

	log.Printf("Merging %d items: %#v", size, source)
	collected := make(map[string]*schema.Schema, size)
	for _, it := range source {
		for key, value := range it {
			collected[key] = value
		}
	}

	log.Printf("Merged into %#v\n", collected)
	return collected
}

func schemaColumn() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"primary": {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		"nullable": {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
	}
}

func schemaDefaultBool() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"default": {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
			Default:  nil,
		},
	}
}

func schemaDefaultFloat() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"default": {
			Type:     schema.TypeFloat,
			Optional: true,
			Computed: true,
			Default:  nil,
		},
	}
}

func schemaDefaultInt() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"default": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
			Default:  nil,
		},
	}
}

func schemaDefaultSet(kind schema.ValueType) map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"default": {
			Type:     schema.TypeSet,
			Elem:     kind,
			Optional: true,
			Computed: true,
			Default:  nil,
		},
	}
}

func schemaDefaultString() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"default": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Default:  nil,
		},
	}
}

func schemaAllowedSet(kind schema.ValueType) map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"allowed": {
			Type:     schema.TypeSet,
			Elem:     kind,
			Required: true,
		},
	}
}

func schemaKinds(fallback string, kinds []string) map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"kind": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      fallback,
			ValidateFunc: validation.StringInSlice(kinds, true),
		},
	}
}

func schemaVariable() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"variable": {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
	}
}

func schemaSized(size int) map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"size": {
			Type:         schema.TypeInt,
			Required:     true,
			ValidateFunc: validation.IntBetween(1, size),
		},
	}
}
