package repositories

import (
	"net/http"
	"tempo/models"
	schema "tempo/serializers/api"
	serializers "tempo/serializers/news"

	"gorm.io/gorm"
)

type RepositoryUpdateNews interface {
	UpdateNewsRepository(input *serializers.NewsSerializer, ID string) (user *models.News, err schema.SchemaDatabaseError)
}

type repositoryUpdateNews struct {
	db *gorm.DB
}

func NewRepositoryUpdateNews(db *gorm.DB) *repositoryUpdateNews {
	return &repositoryUpdateNews{db: db}
}

func (r *repositoryUpdateNews) UpdateNewsRepository(input *serializers.NewsSerializer, ID string) (*models.News, schema.SchemaDatabaseError) {

	var news models.News
	db := r.db.Model(&news)

	news.ID = ID

	checkNewsId := db.Debug().First(&news)
	if checkNewsId.RowsAffected < 1 {
		errorCode := schema.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_01",
		}
		return &news, errorCode
	}

	news.Title = input.Title
	news.Body = input.Body
	news.UpdatedBy = input.UpdatedBy

	updateNews := db.Debug().Where("id = ?", ID).Updates(news)

	if updateNews.RowsAffected < 1 {
		errorCode := schema.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_02",
		}
		return &news, errorCode
	}

	return &news, schema.SchemaDatabaseError{}
}
