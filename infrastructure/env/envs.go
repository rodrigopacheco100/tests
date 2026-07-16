package env

import (
	"os"

	validate "github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type Envs struct {
	DatabaseUrl string `validate:"required,url"`
}

func LoadEnvs() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func GetEnvs() (*Envs, error) {
	validator := validate.New()
	envs := &Envs{
		DatabaseUrl: os.Getenv("DATABASE_URL"),
	}
	if err := validator.Struct(envs); err != nil {
		return nil, err
	}

	return envs, nil
}
