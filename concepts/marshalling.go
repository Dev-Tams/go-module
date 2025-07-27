package concepts

import (
	"encoding/json"
	"fmt"
)

func BasicMarsh(data any) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	jsonData, err = json.MarshalIndent(data, "", "  ")

	return fmt.Sprintln(string(jsonData)), nil
}

// func BasicUnMarsh(data any) (string, error){
// 	jsonData, err := json.Unmarshal([], data)
// 	if err != nil{
// 		return "", nil
// 	}

// 	return  fmt.Sprintln(string(jsonData)), nil
// }
