// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Customer struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	LicenseID   string `json:"license_id"`
	PhoneNumber string `json:"phone_number"`
}

type CustomerInput struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	LicenseID   string `json:"license_id"`
	PhoneNumber string `json:"phone_number"`
}

type LoginInfo struct {
	TokenString string `json:"tokenString"`
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID       string    `json:"id"`
	Username string    `json:"username"`
	Role     string    `json:"role"`
	Customer *Customer `json:"customer,omitempty"`
}

type UserInput struct {
	Username string         `json:"username"`
	Password string         `json:"password"`
	Role     string         `json:"role"`
	Customer *CustomerInput `json:"customer,omitempty"`
}

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
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
