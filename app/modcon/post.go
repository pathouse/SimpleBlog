package modcon

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"time"
)

type Post struct {
	Id          int64     `json:"id"`
	UserId      int64     `json:"user_id"`
	Title       string    `json:"title"`
	Body        string    `json:"body"`
	Draft       bool      `json:"draft"`
	PublishedAt time.Time `json:"-"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"-"`
}

func (p *Post) Published() string {
	return p.PublishedAt.Format("Mon Jan _2, 2006")
}

func (p *Post) GetUser(db *gorm.DB) (*User, error) {
	user := &User{}
	query := map[string]interface{}{"id": p.UserId}
	if err := FindByMap(db, query, user, true); err != nil {
		return nil, err
	}
	return user, nil
}
