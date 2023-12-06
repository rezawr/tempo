package services

import (
	"tempo/models"
	repositories "tempo/repositories/auth"
	schema "tempo/serializers/api"
	serializers "tempo/serializers/auth"
)

type ServiceLogin interface {
	LoginService(input *serializers.LoginSerializer) (*models.Users, schema.SchemaDatabaseError)
}

type serviceLogin struct {
	repository repositories.RepositoryLogin
}

func NewServiceLogin(repository repositories.RepositoryLogin) *serviceLogin {
	return &serviceLogin{repository: repository}
}

func (s *serviceLogin) LoginService(input *serializers.LoginSerializer) (*models.Users, schema.SchemaDatabaseError) {

	var schema serializers.LoginSerializer
	schema.Email = input.Email
	schema.Password = input.Password

	res, err := s.repository.LoginRepository(&schema)
	return res, err
}
