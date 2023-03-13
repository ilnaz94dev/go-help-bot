package commands

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

const FilePerm = 0644

func AddTextToFile(text, fileName string) error {
	file, _ := os.OpenFile(fileName, os.O_APPEND|os.O_RDWR, FilePerm)
	defer file.Close()
	text = strings.Trim(text, " ")

	s := bufio.NewScanner(file)
	for s.Scan() {
		if s.Text() == text {
			return errors.New("url is exists")
		}
	}

	_, err := file.WriteString(text + "\n")
	if err != nil {
		return err
	}
	return nil
}

func RemoveTextFromFile(text, fileName string) error {
	file, _ := os.OpenFile(fileName, os.O_RDWR, FilePerm)
	defer file.Close()
	text = strings.Trim(text, " ")

	s := bufio.NewScanner(file)
	strRemoved := false
	newContent := ""
	for s.Scan() {
		if s.Text() == text {
			strRemoved = true
			continue
		}
		newContent += s.Text() + "\n"
	}

	if !strRemoved {
		return nil
	}

	err := file.Truncate(0)
	if err != nil {
		return err
	}
	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}
	_, err = file.WriteString(newContent)
	if err != nil {
		return err
	}
	return nil
}
