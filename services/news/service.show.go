package services

import (
	"tempo/models"
	repositories "tempo/repositories/news"
	schema "tempo/serializers/api"
)

type ServiceShowNews interface {
	ShowNewsService(ID string) (news *models.News, err schema.SchemaDatabaseError)
}

type serviceShowNews struct {
	repository repositories.RepositoryShowNews
}

func NewServiceShowNews(repository repositories.RepositoryShowNews) *serviceShowNews {
	return &serviceShowNews{repository: repository}
}

func (s *serviceShowNews) ShowNewsService(ID string) (*models.News, schema.SchemaDatabaseError) {
	res, err := s.repository.ShowNewsRepository(ID)
	return res, err
}
