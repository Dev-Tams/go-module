package concepts

import (
	"bufio"
	"encoding/csv"
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

func OpenFile(filename string) (*os.File, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	// defer f.Close()
	return f, nil
}

func Appendtofile(filename, text string) (string, error) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return "", err
	}

	defer f.Close()
	_, err = f.WriteString(text)
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
	encoder.SetIndent("", "  ")
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

func CreateCsv(filename string) (*os.File, error) {
	//if you call the csv alone you have to close the file after
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	// defer f.Close()
	return f, nil

}

func WriteToCsv(filename string, records [][]string) (string, error) {
	f, err := CreateCsv(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	for _, i := range records {
		err := w.Write(i)
		if err != nil {
			return "", err
		}
	}
	return fmt.Sprintln("csv written"), nil

}

func ReadFromCsv(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	fmt.Println(" reading csv file")
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	for i, record := range records {
		fmt.Printf("Row %d: %v\n", i, record)
	}

	return records, nil
}

func ScanFile(file string) (string, error) {
	f, _ := OpenFile(file)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("line", line)

		if err := scanner.Err(); err != nil {
		return 	fmt.Sprintln("error with reading", err), nil
		}

	}
	return "", nil
}
