package modcon

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"time"
)

type Post struct {
	Id          int64
	UserId      int64
	Title       string
	Body        string
	Draft       bool
	PublishedAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
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
