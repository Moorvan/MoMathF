package request

type UpdateLevel struct {
	UUID      string `json:"uuid" validate:"required,uuid"`
	UpdatedId string `json:"updatedId" validate:"required,uuid"`
}
