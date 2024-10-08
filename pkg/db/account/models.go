// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Role string

const (
	RoleUser       Role = "user"
	RoleSubscriber Role = "subscriber"
	RoleModerator  Role = "moderator"
	RoleCreator    Role = "creator"
	RoleAdmin      Role = "admin"
	RoleOrganizer  Role = "organizer"
	RoleRecruiter  Role = "recruiter"
)

func (e *Role) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Role(s)
	case string:
		*e = Role(s)
	default:
		return fmt.Errorf("unsupported scan type for Role: %T", src)
	}
	return nil
}

type NullRole struct {
	Role  Role `json:"Role"`
	Valid bool `json:"valid"` // Valid is true if Role is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullRole) Scan(value interface{}) error {
	if value == nil {
		ns.Role, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Role.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullRole) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Role), nil
}

type TokenType string

const (
	TokenTypeRefreshToken TokenType = "refresh_token"
	TokenTypeAccessToken  TokenType = "access_token"
	TokenTypeMagicLink    TokenType = "magic_link"
)

func (e *TokenType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = TokenType(s)
	case string:
		*e = TokenType(s)
	default:
		return fmt.Errorf("unsupported scan type for TokenType: %T", src)
	}
	return nil
}

type NullTokenType struct {
	TokenType TokenType `json:"token_type"`
	Valid     bool      `json:"valid"` // Valid is true if TokenType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullTokenType) Scan(value interface{}) error {
	if value == nil {
		ns.TokenType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.TokenType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullTokenType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.TokenType), nil
}

type Account struct {
	ID           uuid.UUID   `json:"id"`
	Username     string      `json:"username"`
	Email        string      `json:"email"`
	PasswordHash pgtype.Text `json:"password_hash"`
	Bio          pgtype.Text `json:"bio"`
	ProfileImage pgtype.Text `json:"profile_image"`
	Position     pgtype.Text `json:"position"`
	Organisation string      `json:"organisation"`
	Birthday     pgtype.Date `json:"birthday"`
}

type Token struct {
	ID     uuid.UUID `json:"id"`
	Token  string    `json:"token"`
	UserID int32     `json:"user_id"`
	Expiry time.Time `json:"expiry"`
}
