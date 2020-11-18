package models

import (
	"github.com/francescoforesti/go-demo/gin/models"
	"time"
)

type KafkaMessage struct {
	Content   models.GinMessage
	Reversed  models.GinMessage
	Timestamp *time.Time
}
