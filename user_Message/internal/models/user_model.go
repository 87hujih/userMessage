package models

import "database/sql"

type User struct {
	UserId       sql.NullInt64
	UserPhone    sql.NullString
	UserName     sql.NullString
	IsAdmin      sql.NullInt64
	Age          sql.NullInt64
	Gender       sql.NullString
	RegisterData sql.NullString
	Email        sql.NullString
	AvatarURL    sql.NullString
}
