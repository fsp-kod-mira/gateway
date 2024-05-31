package entity

type CV struct {
	Id         string
	Status     string
	UploadedBy string
	FileId     string
	Features   map[string]interface{}
	Metric     float64
}
