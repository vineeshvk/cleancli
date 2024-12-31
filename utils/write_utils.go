package utils

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/vineeshvk/cleancli/templates"
)

// Inserts the newData before the lastBrace if not exist will be appended to the last
// Also topData will be inserted at the top most list in the file
func InsertToFileBeforeLastBrace(filePath string, newData string, topData string) error {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)

	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}

	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	lastBraceIndex := bytes.LastIndex(content, []byte("}"))
	if lastBraceIndex == -1 {
		lastBraceIndex = len(content) - 1
	}

	// Check if the new data is already present in the content
	if bytes.Contains(content, []byte(newData)) {
		fmt.Println("Data is already present in ", filePath, ". So skipping.")
		return nil
	}

	var buffer bytes.Buffer

	if !bytes.Contains(content, []byte(topData)) {
		buffer.WriteString(topData)
	}
	buffer.Write(content[:lastBraceIndex])
	buffer.WriteString(newData)
	buffer.Write(content[lastBraceIndex:])

	file.Seek(0, 0)
	err = file.Truncate(0)

	if err != nil {
		return fmt.Errorf("failed to truncate file: %w", err)
	}

	writer := bufio.NewWriter(file)
	_, err = writer.Write(buffer.Bytes())

	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	fmt.Println("Data inserted successfully in ", filePath)

	return writer.Flush()
}

func AppendToFile(filePath string, newData string, topData string) error {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)

	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}

	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Check if the new data is already present in the content
	if bytes.Contains(content, []byte(newData)) {
		fmt.Println("Data is already present in ", filePath, ". So skipping.")
		return nil
	}

	var buffer bytes.Buffer
	if !bytes.Contains(content, []byte(topData)) {
		buffer.WriteString(topData)
	}

	buffer.Write(content)
	buffer.WriteString(newData)

	file.Seek(0, 0)
	err = file.Truncate(0)

	if err != nil {
		return fmt.Errorf("failed to truncate file: %w", err)
	}

	writer := bufio.NewWriter(file)
	_, err = writer.Write(buffer.Bytes())

	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	fmt.Println("Data appended successfully in ", filePath)

	return writer.Flush()
}

func CreateAndInsertIfFileNotExist(fileroute string, data string) error {

	dir := filepath.Dir(fileroute)

	direrr := os.MkdirAll(dir, os.ModePerm)

	if direrr != nil {
		fmt.Printf("Folder already exist: %s", dir)
	}

	_, fileerr := os.Stat(fileroute)

	// Then the file already exist
	if fileerr == nil {
		fmt.Printf("File already exist: %s", dir)
		return nil
	}

	if os.IsNotExist(fileerr) {
		fmt.Println("File ", fileroute, "not found. Creating it.")

		// Create the file
		file, err := os.Create(fileroute)

		if err != nil {
			return errors.New("Error file creating file: " + fileroute + ", error: " + err.Error())
		}

		defer file.Close()

		if data == "" {
			return nil
		}

		if _, err = file.WriteString(data); err != nil {
			return errors.New("Error file writing to file: " + fileroute + ", error: " + err.Error())
		}

		fmt.Println("File created : ", fileroute)

	} else {
		return errors.New("Error while reading file: " + fileroute + ", error: " + fileerr.Error())
	}

	return nil

}

func GetImportRoute(route string) string {
	packageDir := strings.ReplaceAll(filepath.ToSlash(route), "/lib/", "/")
	return fmt.Sprintf(templates.ImportStatement+"\n", packageDir)
}
