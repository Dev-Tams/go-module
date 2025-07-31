package concepts

import (
	"encoding/json"
	"fmt"
)

func BasicMarsh(data any) (string, error) {

	indent, err := json.MarshalIndent(data, "", "  ")
	if err == nil {
		return string(indent), err
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return fmt.Sprintln(string(jsonData)), nil
} 

// BasicUnmarshal decodes JSON from input into out (which must be a pointer).
func BasicUnMarsh(input []byte, out any) error {
	return json.Unmarshal(input, out)
}
