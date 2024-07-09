package shared

type ViewFileResponse struct {
	// example: 'png.png
	FileName string `json:"fileName" example:"png.png"`
	// example: 'http://localhost:8080/api/v1/file/png.png
	FileURL string `json:"fileURL" example:"http://localhost:8080/api/v1/file/png.png"`
}
