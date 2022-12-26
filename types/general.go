package types

type ApplicationErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
