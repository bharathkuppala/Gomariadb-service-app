package utility

import (
	"log"

	"github.com/Shopify/sarama"
	"gorm.io/gorm"
)

var (
	// Logger ...
	Logger *log.Logger
	// GormDatabase ...
	GormDatabase *gorm.DB
	// Producer ...
	Producer sarama.SyncProducer
)
