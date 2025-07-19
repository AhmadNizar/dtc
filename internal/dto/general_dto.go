package dto

type GeneralResponseDTO struct {
	OK      bool        `json:"ok"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
