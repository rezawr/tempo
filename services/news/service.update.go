package services

import (
	"tempo/models"
	repositories "tempo/repositories/news"
	schema "tempo/serializers/api"
	serializers "tempo/serializers/news"
)

type ServiceUpdateNews interface {
	UpdateNewsService(input *serializers.NewsSerializer, ID string, userID string) (news *models.News, err schema.SchemaDatabaseError)
}

type serviceUpdateNews struct {
	repository repositories.RepositoryUpdateNews
}

func NewServiceUpdateNews(repository repositories.RepositoryUpdateNews) *serviceUpdateNews {
	return &serviceUpdateNews{repository: repository}
}

func (s *serviceUpdateNews) UpdateNewsService(input *serializers.NewsSerializer, ID string, userID string) (*models.News, schema.SchemaDatabaseError) {

	var schema serializers.NewsSerializer
	schema.Title = input.Title
	schema.Body = input.Body
	schema.UpdatedBy = userID

	res, err := s.repository.UpdateNewsRepository(&schema, ID)
	return res, err
}
