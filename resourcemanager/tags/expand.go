package tags

func Expand(input map[string]interface{}) *map[string]string {
	output := make(map[string]string)

	for k, v := range input {
		tagKey := k
		tagValue := v.(string)
		output[tagKey] = tagValue
	}

	return &output
}
