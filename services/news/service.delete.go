package services

import (
	"tempo/models"
	repositories "tempo/repositories/news"
	schema "tempo/serializers/api"
)

type ServiceDeleteNews interface {
	DeleteNewsService(ID string) (news *models.News, err schema.SchemaDatabaseError)
}

type serviceDeleteNews struct {
	repository repositories.RepositoryDeleteNews
}

func NewServiceDeleteNews(repository repositories.RepositoryDeleteNews) *serviceDeleteNews {
	return &serviceDeleteNews{repository: repository}
}

func (s *serviceDeleteNews) DeleteNewsService(ID string) (*models.News, schema.SchemaDatabaseError) {
	res, err := s.repository.DeleteNewsRepository(ID)
	return res, err
}
