package concepts

import (
	"encoding/json"
	"fmt"
	"os"
)

func Createfile(filename string) (string, error) {
	f, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	f.Close()
	return fmt.Sprintln("file created"), nil

}

func OpenFile(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	f.Close()
	return fmt.Sprintln("File opened successfully!"), nil
}

// func WriteString(filename, text string) (string, error) {
// 	OpenFile(filename)
// 	_, err := filename.WriteString(text)
// 	if err != nil {
// 		return "", err
// 	}
// 	return fmt.Sprintln("file written successfully"), nil
// }

func Appendtofile(filename, text string) (string, error) {
	f, _ := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	_, err := f.WriteString(text)
	if err != nil {
		return "", err
	}
	return fmt.Sprintln("Line appended successfully."), nil

}

func WriteToJson(filename string, data any) (string, error) {
	f, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	err = encoder.Encode(data)
	if err != nil {
		return "", err
	}
	return fmt.Sprintln("JSON written to file."), nil
}

func ReadFromJson(filename string, data any) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	decode := json.NewDecoder(f)
	err = decode.Decode(&data)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("User loaded from JSON: %+v\n, user", data), nil
}
