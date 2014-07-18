package modcon

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"time"
)

type Post struct {
	Id        int64
	UserId    int64
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (p *Post) Published() string {
	return p.CreatedAt.Format("Mon Jan _2, 2006")
}

func (p *Post) LinkTitle() string {
	return "Post Link Title"
}

func (p *Post) Url() string {
	return "http://www.post-link.com"
}

func (p *Post) AuthorName(db *gorm.DB) string {
	user, err := p.getUser(db)
	if err != nil {
		return "Author Missing"
	}
	return fmt.Sprintf("%s %s", user.FirstName, user.LastName)
}

func (p *Post) getUser(db *gorm.DB) (*User, error) {
	user := &User{}
	query := map[string]interface{}{"id": p.UserId}
	if err := FindByMap(db, query, user, true); err != nil {
		return nil, err
	}
	return user, nil
}
