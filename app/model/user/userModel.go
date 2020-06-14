package user

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type StatusData struct {
	IsActive bool `json:"is_active,omitempty"`
}

func (p StatusData) Value() (driver.Value, error) {
	j, err := json.Marshal(p)
	return j, err
}

func (p *StatusData) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}

	err := json.Unmarshal(source, &p)
	if err != nil {
		return err
	}

	return nil
}

type UserData struct {
	Username string     `json:"username,omitempty"`
	Email    string     `json:"email,omitempty"`
	Status   StatusData `json:"status,omitempty"`
}

func (p UserData) Value() (driver.Value, error) {
	j, err := json.Marshal(p)
	return j, err
}

func (p *UserData) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}

	err := json.Unmarshal(source, &p)
	if err != nil {
		return err
	}

	return nil
}

type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;column:id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("ID", uuid)
}

type User struct {
	Base
	Data   UserData `gorm:"column:data;type:jsonb;"`
	RoleID string   `gorm:"column:role_id;"`
}

func (User) TableName() string {
	return "users"
}
