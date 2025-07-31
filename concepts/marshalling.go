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
	_, err = json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}

	return fmt.Sprintln(string(jsonData)), nil
} 

func BasicUnMarsh(datas []byte, data any) (string, error){
	jsonData := json.Unmarshal(datas, data)
	err := jsonData
	if err != nil{
		return "", nil
	}

	return  fmt.Sprintln(error(jsonData)), nil
}
