package repositories

import (
	"net/http"
	"tempo/models"
	schema "tempo/serializers/api"

	"gorm.io/gorm"
)

type RepositoryGetNews interface {
	GetNewsRepository() (user *[]models.News, err schema.SchemaDatabaseError)
}

type repositoryGetNews struct {
	db *gorm.DB
}

func NewRepositoryGetNews(db *gorm.DB) *repositoryGetNews {
	return &repositoryGetNews{db: db}
}

func (r *repositoryGetNews) GetNewsRepository() (*[]models.News, schema.SchemaDatabaseError) {

	var news []models.News
	db := r.db.Model(&news)

	resultNews := db.Debug().Find(&news)
	if resultNews.RowsAffected < 1 {
		errorCode := schema.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_01",
		}
		return &news, errorCode
	}

	return &news, schema.SchemaDatabaseError{}
}
