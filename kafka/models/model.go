package models

import "github.com/francescoforesti/go-demo/gin/models"

type KafkaMessage struct {
	Content   models.GinMessage
	Timestamp string
}
