package repositories

import (
	"net/http"
	"tempo/models"
	schema "tempo/serializers/api"
	serializers "tempo/serializers/auth"

	"gorm.io/gorm"
)

type RepositoryRegister interface {
	RegisterRepository(input *serializers.RegisterSerializer) (user *models.Users, err schema.SchemaDatabaseError)
}

type repositoryRegister struct {
	db *gorm.DB
}

func NewRepositoryRegister(db *gorm.DB) *repositoryRegister {
	return &repositoryRegister{db: db}
}

func (r *repositoryRegister) RegisterRepository(input *serializers.RegisterSerializer) (*models.Users, schema.SchemaDatabaseError) {

	var user models.Users
	db := r.db.Model(&user)

	checkUserAccount := db.Debug().First(&user, "email = ?", input.Email)

	if checkUserAccount.RowsAffected > 0 {
		errorCode := schema.SchemaDatabaseError{
			Code: http.StatusConflict,
			Type: "error_01",
		}
		return &user, errorCode
	}

	user.Fullname = input.Fullname
	user.Email = input.Email
	user.Password = input.Password

	addNewUser := db.Debug().Create(&user).Commit()

	if addNewUser.RowsAffected < 1 {
		errorCode := schema.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_02",
		}
		return &user, errorCode
	}

	return &user, schema.SchemaDatabaseError{}
}
