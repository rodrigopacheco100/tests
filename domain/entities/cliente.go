package entities

import (
	"time"

	"github.com/samborkent/uuidv7"
)

type ClienteTipo = string

const (
	PessoaFisica   ClienteTipo = "fisico"
	PessoaJuridica ClienteTipo = "juridico"
)

type Cliente struct {
	Id        string
	Nome      string
	Email     string
	Tipo      ClienteTipo
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

func NewCliente(name, email string, tipo ClienteTipo) *Cliente {
	now := time.Now().In(time.FixedZone("America/Sao_Paulo", -3*60*60))

	return &Cliente{
		Id:        uuidv7.New().String(),
		Nome:      name,
		Email:     email,
		Tipo:      tipo,
		CreatedAt: &now,
		UpdatedAt: &now,
		DeletedAt: nil,
	}
}

func (c *Cliente) Destroy() {
	now := time.Now()
	c.DeletedAt = &now
}

func (c *Cliente) Restore() {
	c.DeletedAt = nil
}
