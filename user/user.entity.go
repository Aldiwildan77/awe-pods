package users

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// Users entity
type Users struct {
	ID          uuid.UUID `gorm:"type:uuid"`
	Name        string
	Username    string
	Email       string
	Password    string
	Age         int
	Gender      string `sql:"type:gender"`
	PhotoURL    string
	Description string
	Verified    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
}

// BeforeCreate for do some action before the data stored into database
func (u *Users) BeforeCreate(scope *gorm.Scope) (err error) {
	if password, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost); err == nil {
		scope.SetColumn("Password", password)
	}

	return
}
