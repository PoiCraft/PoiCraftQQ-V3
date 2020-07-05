package data

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

func initFiles() error {
	load := func(f *os.File, content string) error {
		defer f.Close()

		_, err := io.Copy(f, strings.NewReader(content))
		if err != nil {
			return err
		}

		return nil
	}
	for fileName, fileContent := range defaultFiles {
		f, err := os.OpenFile(filepath.Join(AppDir, fileName), os.O_CREATE|os.O_EXCL|os.O_RDWR, 0666)
		if os.IsExist(err) {
			continue
		} else if err != nil {
			return err
		}

		err = load(f, fileContent)
		if err != nil {
			return err
		}
	}
	return nil
}

var defaultFiles = map[string]string{
	"conf.yaml": `mysql: root:root@(127.0.0.1:33060)/poicraft?charset=utf8&parseTime=True&loc=Local
administrators: [1257135905,3386475667,2331490629]
remote_url: http://127.0.0.1:8234/map/element-zero
tp_number: 3`,
}
