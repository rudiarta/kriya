package role

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

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

type RoleData struct {
	RoleName    string `json:"role_name,omitempty"`
	Description string `json:"description,omitempty"`
}

func (p RoleData) Value() (driver.Value, error) {
	j, err := json.Marshal(p)
	return j, err
}

func (p *RoleData) Scan(src interface{}) error {
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

type Role struct {
	Data RoleData `grom:"column:data;type:jsonb;"`
}

func (Role) TableName() string {
	return "roles"
}
