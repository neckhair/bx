package bexio

type SearchCriteria struct {
	Field    string `json:"field"`
	Value    string `json:"value"`
	Criteria string `json:"criteria"`
}

func NewSearchCriteria(field, criteria, value string) SearchCriteria {
	return SearchCriteria{Field: field, Value: value, Criteria: criteria}
}
