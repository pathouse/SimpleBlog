package modcon

import (
	"code.google.com/p/go.crypto/bcrypt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"time"
)

type User struct {
	Id           int64
	Byline       string `sql:"size:255"`
	Email        string `sql:"type:varchar(100)";`
	PasswordHash string `sql:"size:255"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time

	Posts []Post
}

func (u *User) GenerateHashedPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(hash[:])
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func (u *User) GetPosts(db *gorm.DB) error {
	query := map[string]interface{}{"user_id": u.Id}
	if err := FindByMap(db, query, u.Posts, false); err != nil {
		return err
	}
	return nil
}
