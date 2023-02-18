package util

type ErrorMetaConfig struct {
	Tag     string
	Field   string
	Message string
}

type ErrorConfig struct {
	Options map[string]ErrorMetaConfig
}

type errorResponse struct {
	Results interface{} `json:"results"`
}

type binding struct {
	validator validators
}

func NewBindValidator(validator validators) *binding {
	return &binding{validator: validator}
}

func (b *binding) BindValidator(s interface{}, config map[string]ErrorMetaConfig) (*errorResponse, int) {

	mergeObject := make(map[int]interface{})

	i := 0
	for _, value := range config {
		mergeObject[i] = value
		i++
	}

	var errors errorResponse
	validate := b.validator.validator(s, mergeObject)
	errDataCollection := make(map[string][]map[string]interface{})

	if len(validate) > 0 {
		for i, v := range validate {
			errorResults := make(map[string]interface{})
			errorResults[i] = v
			errDataCollection["errors"] = append(errDataCollection["errors"], errorResults)
		}
		errors.Results = errDataCollection
	}

	return &errors, len(validate)
}