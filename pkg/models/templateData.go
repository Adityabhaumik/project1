package models

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
