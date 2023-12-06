package services

import (
	"tempo/models"
	repositories "tempo/repositories/news"
	schema "tempo/serializers/api"
)

type ServiceGetNews interface {
	GetNewsService() (news *[]models.News, err schema.SchemaDatabaseError)
}

type serviceGetNews struct {
	repository repositories.RepositoryGetNews
}

func NewServiceGetNews(repository repositories.RepositoryGetNews) *serviceGetNews {
	return &serviceGetNews{repository: repository}
}

func (s *serviceGetNews) GetNewsService() (*[]models.News, schema.SchemaDatabaseError) {
	res, err := s.repository.GetNewsRepository()
	return res, err
}
