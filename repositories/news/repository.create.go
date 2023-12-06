package repositories

import (
	"net/http"
	"tempo/models"
	schema "tempo/serializers/api"
	serializers "tempo/serializers/news"

	"gorm.io/gorm"
)

type RepositoryCreateNews interface {
	CreateNewsRepository(input *serializers.NewsSerializer) (user *models.News, err schema.SchemaDatabaseError)
}

type repositoryCreateNews struct {
	db *gorm.DB
}

func NewRepositoryCreateNews(db *gorm.DB) *repositoryCreateNews {
	return &repositoryCreateNews{db: db}
}

func (r *repositoryCreateNews) CreateNewsRepository(input *serializers.NewsSerializer) (*models.News, schema.SchemaDatabaseError) {

	var news models.News
	db := r.db.Model(&news)

	news.Title = input.Title
	news.Body = input.Body
	news.CreatedBy = input.CreatedBy

	addNew := db.Debug().Create(&news).Commit()

	if addNew.RowsAffected < 1 {
		errorCode := schema.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_01",
		}
		return &news, errorCode
	}

	return &news, schema.SchemaDatabaseError{}
}
