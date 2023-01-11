package repository

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"git.ecobin.ir/ecomicro/template/domain"
	"gorm.io/gorm"
)

type User struct {
	Id        uint64 `gorm:"primaryKey;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Roles     StringArray
	Allow     StringArray
	Deny      StringArray
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func FromDomainUser(user domain.User) User {
	return User{
		Id:        user.Id,
		CreatedAt: user.CreatedDate,
		UpdatedAt: user.UpdatedDate,
		Roles:     user.Roles,
		Allow:     user.Allow,
		Deny:      user.Deny,
	}
}

func (u *User) ToDomainUser() domain.User {
	return domain.User{
		Id:          u.Id,
		CreatedDate: u.CreatedAt,
		UpdatedDate: u.UpdatedAt,
		Roles:       u.Roles,
		Allow:       u.Allow,
		Deny:        u.Deny,
	}
}

// for jsonb type
type StringArray []string

// Value Marshal
func (sa StringArray) Value() (driver.Value, error) {
	return json.Marshal(sa)
}

// Scan Unmarshal
func (sa *StringArray) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := []string{}
	err := json.Unmarshal(bytes, &result)
	*sa = StringArray(result)
	return err
}
