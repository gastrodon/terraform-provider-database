package types

import (
	"gorm.io/gorm"
)

type Config struct {
	Connection *gorm.DB
}
