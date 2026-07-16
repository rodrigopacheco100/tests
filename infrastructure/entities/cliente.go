package infra_entities

import "time"

type Cliente struct {
	Id        string     `db:"id"`
	Nome      string     `db:"nome"`
	Email     string     `db:"email"`
	Tipo      string     `db:"tipo"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
