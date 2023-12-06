package services

import (
	"tempo/models"
	repositories "tempo/repositories/news"
	schema "tempo/serializers/api"
	serializers "tempo/serializers/news"
)

type ServiceCreateNews interface {
	CreateNewsService(input *serializers.NewsSerializer, userID string) (news *models.News, err schema.SchemaDatabaseError)
}

type serviceCreateNews struct {
	repository repositories.RepositoryCreateNews
}

func NewServiceCreateNews(repository repositories.RepositoryCreateNews) *serviceCreateNews {
	return &serviceCreateNews{repository: repository}
}

func (s *serviceCreateNews) CreateNewsService(input *serializers.NewsSerializer, userID string) (*models.News, schema.SchemaDatabaseError) {

	var schema serializers.NewsSerializer
	schema.Title = input.Title
	schema.Body = input.Body
	schema.CreatedBy = userID

	res, err := s.repository.CreateNewsRepository(&schema)
	return res, err
}
