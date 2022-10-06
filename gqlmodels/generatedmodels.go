// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gqlmodels

import (
	"fmt"
	"io"
	"strconv"
)

type Article struct {
	ID        string  `json:"id"`
	Title     *string `json:"title"`
	CreatedAt *int    `json:"createdAt"`
	UpdatedAt *int    `json:"updatedAt"`
	DeletedAt *int    `json:"deletedAt"`
}

type ArticleCreateInput struct {
	Title    string `json:"title"`
	AuthorID string `json:"authorId"`
}

type ArticleDeleteInput struct {
	ID string `json:"id"`
}

type ArticleDeletePayload struct {
	ID string `json:"id"`
}

type ArticlePayload struct {
	Article *Article `json:"article"`
}

type ArticleUpdateInput struct {
	ID    string  `json:"id"`
	Title *string `json:"title"`
}

type ArticlesPayload struct {
	Articles []*Article `json:"articles"`
	Total    int        `json:"total"`
}

type Author struct {
	ID        string     `json:"id"`
	FirstName *string    `json:"firstName"`
	LastName  *string    `json:"lastName"`
	Username  *string    `json:"username"`
	Password  *string    `json:"password"`
	Active    *bool      `json:"active"`
	Token     *string    `json:"token"`
	Role      *string    `json:"role"`
	Articles  []*Article `json:"articles"`
	CreateAt  *int       `json:"createAt"`
	UpdatedAt *int       `json:"updatedAt"`
	DeletedAt *int       `json:"deletedAt"`
}

type AuthorCreateInput struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Active    *bool  `json:"active"`
	Role      string `json:"role"`
}

type AuthorDeleteInput struct {
	ID string `json:"id"`
}

type AuthorDeletePayload struct {
	ID string `json:"id"`
}

type AuthorPayload struct {
	Author *Author `json:"author"`
}

type AuthorUpdateInput struct {
	ID        string  `json:"id"`
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
	Username  *string `json:"username"`
	Password  *string `json:"password"`
	Active    *bool   `json:"active"`
	CreatedAt *int    `json:"createdAt"`
	UpdatedAt *int    `json:"updatedAt"`
	DeletedAt *int    `json:"deletedAt"`
}

type AuthorsPayload struct {
	Authors []*Author `json:"authors"`
	Total   int       `json:"total"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type Role string

const (
	RoleAdmin Role = "ADMIN"
	RoleUser  Role = "USER"
)

var AllRole = []Role{
	RoleAdmin,
	RoleUser,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleAdmin, RoleUser:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
