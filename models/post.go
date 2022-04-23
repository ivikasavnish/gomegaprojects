package models

import (
	"html"
	"strings"
	"time"
)

type Post struct {
	ID        uint64
	Title     string
	Content   string
	Author    User
	AuthorID  uint
	CreatedAt time.Time
}

func (p *Post) Prepare() {
	p.Title = html.EscapeString(strings.TrimSpace(p.Title))
	p.Content = html.EscapeString(strings.TrimSpace(p.Title))
}

func (p *Post) Validate() {

}

func (p *Post) SavePost() {

}
func (p *Post) FindAllPost() {

}

func (p *Post) UpdateAPost() {

}
func (p *Post) DeletePost() {

}
