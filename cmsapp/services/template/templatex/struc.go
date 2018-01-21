package templatex

// Dict create create map from keys and values list
func Dict(v ...interface{}) map[string]interface{} {
	dict := map[string]interface{}{}
	lenv := len(v)
	for i := 0; i < lenv; i += 2 {
		key := v[i].(string)
		if i+1 >= lenv {
			dict[key] = ""
			continue
		}
		dict[key] = v[i+1]
	}
	return dict
}
