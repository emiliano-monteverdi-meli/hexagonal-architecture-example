package telemetry

type Tags map[string]interface{}

func (t Tags) getValues() []interface{} {
	if len(t) == 0 {
		return nil
	}

	result := make([]interface{}, 0)

	for key, value := range t {
		result = append(result, key, value)
	}

	return result
}

func (t Tags) Add(key string, value interface{}) {
	t[key] = value
}
