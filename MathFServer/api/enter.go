package api

import "github.com/go-playground/validator/v10"

type APIGroup struct {
	MathAPI
	UserAPI
}

var APIGroupApp = new(APIGroup)

var validate = validator.New()
