package formater

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const offset int32 = 9770

// const offset int32 = 0

func readFile(path string, mode os.FileMode) []byte {
	f, _ := os.OpenFile(path, os.O_RDONLY, mode.Perm())
	defer f.Close()
	buffer := make([]byte, 1024)
	var data []byte
	for {
		n, err := f.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		if n > 0 {
			data = append(data, buffer[:n]...)
		}
	}
	return data
}

func formatData(data []byte, reversedMode bool) []byte {
	stringData := string(data)
	var formatted []string
	for _, r := range stringData {
		if r == '\n' {
			formatted = append(formatted, string(r))
		} else {
			if reversedMode {
				formatted = append(formatted, string(r-offset))
			} else {
				formatted = append(formatted, string(r+offset))
			}
		}
	}
	formattedString := strings.Join(formatted, "")
	return []byte(formattedString)
}

func writeFile(path string, mode os.FileMode, data []byte) error {
	if err := os.Truncate(path, 0); err != nil {
		return err
	}

	f, _ := os.OpenFile(path, os.O_WRONLY, mode.Perm())
	defer f.Close()
	_, err := f.Write(data)
	return err
}

func Format(path string, mode os.FileMode, reversedMode bool) error {
	data := readFile(path, mode)
	result := formatData(data, reversedMode)
	fmt.Println(string(result))
	err := writeFile(path, mode, result)
	return err
}
