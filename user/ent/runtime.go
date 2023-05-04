// Code generated by ent, DO NOT EDIT.

package ent

import (
	"flookybooky/user/ent/schema"
	"flookybooky/user/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[0].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[1].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescRole is the schema descriptor for role field.
	userDescRole := userFields[2].Descriptor()
	// user.RoleValidator is a validator for the "role" field. It is called by the builders before save.
	user.RoleValidator = userDescRole.Validators[0].(func(string) error)
}
