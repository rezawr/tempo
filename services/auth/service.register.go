package services

import (
	"tempo/models"
	repositories "tempo/repositories/auth"
	schema "tempo/serializers/api"
	serializers "tempo/serializers/auth"
)

type ServiceRegister interface {
	RegisterService(input *serializers.RegisterSerializer) (*models.Users, schema.SchemaDatabaseError)
}

type serviceRegister struct {
	repository repositories.RepositoryRegister
}

func NewServiceRegister(repository repositories.RepositoryRegister) *serviceRegister {
	return &serviceRegister{repository: repository}
}

func (s *serviceRegister) RegisterService(input *serializers.RegisterSerializer) (*models.Users, schema.SchemaDatabaseError) {

	var schema serializers.RegisterSerializer
	schema.Fullname = input.Fullname
	schema.Email = input.Email
	schema.Password = input.Password

	res, err := s.repository.RegisterRepository(&schema)
	return res, err
}
