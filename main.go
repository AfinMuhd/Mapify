// StructToMapConverter takes a value of any type and converts it into a map[string]interface{}.
// It uses JSON marshaling and unmarshaling to perform the conversion.
// It returns the converted map[string]interface{} and an error if there was any issue during the conversion process.
func StructToMapConverter(v any) (map[string]any, error) {

	// Initialize an empty map to store the converted data
	var data map[string]any

	// Marshal the input value into a JSON-encoded byte array
	info, err := json.Marshal(v)
	if err != nil {
		return data, err
	}

	// Unmarshal the JSON-encoded byte array into the map
	err = json.Unmarshal([]byte(info), &data)
	if err != nil {
		return data, err
	}

	// Return the converted map and nil error
	return data, nil
}
