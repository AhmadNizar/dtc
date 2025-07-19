package models

import (
	"time"

	uuid "github.com/google/uuid"
)

type AuditLog struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	TableName string    `gorm:"not null"`
	RecordID  uuid.UUID `gorm:"not null"`
	Action    string    `gorm:"not null"` // CREATE, UPDATE, DELETE
	UserID    *uuid.UUID
	IPAddress string
	RequestID *uuid.UUID
	OldData   map[string]interface{} `gorm:"type:jsonb"`
	NewData   map[string]interface{} `gorm:"type:jsonb"`
	CreatedAt time.Time

	User *User `gorm:"foreignKey:UserID"`
}
