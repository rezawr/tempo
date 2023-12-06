package repositories

import (
	"net/http"
	"tempo/models"
	"tempo/pkg"
	schema "tempo/serializers/api"
	serializers "tempo/serializers/auth"

	"gorm.io/gorm"
)

type RepositoryLogin interface {
	LoginRepository(input *serializers.LoginSerializer) (user *models.Users, err schema.SchemaDatabaseError)
}

type repositoryLogin struct {
	db *gorm.DB
}

func NewRepositoryLogin(db *gorm.DB) *repositoryLogin {
	return &repositoryLogin{db: db}
}

func (r *repositoryLogin) LoginRepository(input *serializers.LoginSerializer) (*models.Users, schema.SchemaDatabaseError) {

	var user models.Users
	db := r.db.Model(&user)

	user.Email = input.Email
	user.Password = input.Password

	checkUserAccount := db.Debug().First(&user, "email = ?", input.Email)

	if checkUserAccount.RowsAffected < 1 {
		errorCode := schema.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_01",
		}
		return &user, errorCode
	}

	comparePassword := pkg.ComparePassword(user.Password, input.Password)

	if comparePassword != nil {
		errorCode := schema.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_02",
		}
		return &user, errorCode
	}

	return &user, schema.SchemaDatabaseError{}
}
