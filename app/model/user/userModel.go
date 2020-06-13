package user

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type UserData map[string]interface{}

func (p UserData) Value() (driver.Value, error) {
	j, err := json.Marshal(p)
	return j, err
}

func (p *UserData) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}

	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	*p, ok = i.(map[string]interface{})
	if !ok {
		return errors.New("Type assertion .(map[string]interface{}) failed.")
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
