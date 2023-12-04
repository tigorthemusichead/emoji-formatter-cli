package scanner

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"t3011/formater"
)

func Scan(path string, ignoreList []string, reversedMode bool) error {
	err := filepath.Walk(path, func(filepath string, f os.FileInfo, err error) error {
		_continue := false
		for _, ignore := range ignoreList {
			if strings.Index(filepath, ignore) != -1 {
				_continue = true
			}
		}
		if _continue {
			return err
		}

		fMode := f.Mode()

		if !fMode.IsDir() {
			fmt.Printf("FILE: %s, %v", filepath, fMode)
			err = formater.Format(filepath, fMode, reversedMode)
		}
		return err
	})
	return err
}
