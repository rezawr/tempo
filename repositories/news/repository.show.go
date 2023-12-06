package repositories

import (
	"net/http"
	"tempo/models"
	schema "tempo/serializers/api"

	"gorm.io/gorm"
)

type RepositoryShowNews interface {
	ShowNewsRepository(ID string) (user *models.News, err schema.SchemaDatabaseError)
}

type repositoryShowNews struct {
	db *gorm.DB
}

func NewRepositoryShowNews(db *gorm.DB) *repositoryShowNews {
	return &repositoryShowNews{db: db}
}

func (r *repositoryShowNews) ShowNewsRepository(ID string) (*models.News, schema.SchemaDatabaseError) {

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

	return &news, schema.SchemaDatabaseError{}
}
