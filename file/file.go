package file

import "os"

type FileRW struct {
	filename string
}

func (fw *FileRW) Write(data string) error {
	file, err := os.OpenFile(fw.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer file.Close()
	_, err = file.WriteString(data + "\n")

	if err != nil {
		return err
	}

	return nil
}

func (fr *FileRW) Read() (string, error) {
	data, err := os.ReadFile(fr.filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
