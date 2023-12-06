package repositories

import (
	"net/http"
	"tempo/models"
	schema "tempo/serializers/api"

	"gorm.io/gorm"
)

type RepositoryDeleteNews interface {
	DeleteNewsRepository(ID string) (user *models.News, err schema.SchemaDatabaseError)
}

type repositoryDeleteNews struct {
	db *gorm.DB
}

func NewRepositoryDeleteNews(db *gorm.DB) *repositoryDeleteNews {
	return &repositoryDeleteNews{db: db}
}

func (r *repositoryDeleteNews) DeleteNewsRepository(ID string) (*models.News, schema.SchemaDatabaseError) {

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

	deleteNewsId := db.Debug().Delete(&news)
	if deleteNewsId.RowsAffected < 1 {
		errorCode := schema.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_02",
		}
		return &news, errorCode
	}

	return &news, schema.SchemaDatabaseError{}
}
