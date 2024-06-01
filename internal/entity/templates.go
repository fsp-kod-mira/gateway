package entity

type FeatureType string

const (
	FeatureStruct_RANGE FeatureType = "RANGE"
	FeatureStruct_LIST  FeatureType = "LIST"
)

type TemplateStruct struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type FeatureStruct struct {
	Id          uint64      `json:"id"`
	FeatureId   uint64      `json:"featureId"`
	TemplateId  uint64      `json:"templateId"`
	Value       string      `json:"value"`
	Name        string      `json:"name"`
	FeatureType FeatureType `json:"featureType"`
}
