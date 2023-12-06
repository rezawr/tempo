package serializers

type NewsSerializer struct {
	Title     string `json:"title" validate:"required"`
	Body      string `json:"body" validate:"required"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
}
