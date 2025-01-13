package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"gopkg.in/yaml.v3"
)

func DoesFileExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func DoesFileExistByRegex(root string, target string) (bool, string) {
	var dir string

	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {

		if err != nil {
			fmt.Printf("Error accessing path %s: %v\n", path, err)
			return nil
		}

		// Ensure we skip directories
		if d.IsDir() {
			return nil
		}

		hasFoundFile, _ := regexp.MatchString(target, filepath.ToSlash(path))

		// Check if the file matches
		if hasFoundFile {
			dir = path
			return filepath.SkipAll
		}

		return nil
	})

	if err != nil || dir == "" {
		return false, ""
	}

	return true, dir

}

func GetMatchingFilesFromDir(root string, pattern string) []string {

	var fileList []string

	err := filepath.WalkDir(root, func(path string, dir os.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %s: %v\n", path, err)
			return nil
		}

		if !dir.IsDir() {
			if pattern == "" {
				fileList = append(fileList, path)
				return nil
			}

			hasMatch, _ := regexp.MatchString(pattern, filepath.ToSlash(path))

			if hasMatch {
				fileList = append(fileList, path)
				return nil
			}
		}

		return nil

	})

	if err != nil {
		return nil
	}

	return fileList

}

func GetClassNameFromFile(path string) string {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println("Couldn't find file path " + path)
		os.Exit(1)
	}

	defer file.Close()

	re, _ := regexp.Compile(`class (\w+)`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)

		if matches != nil {
			return matches[1]
		}
	}

	fmt.Printf("Couldn't find class name. Exiting.")
	os.Exit(0)

	return ""

}

func DoesFolderExist(path string) bool {
	info, err := os.Stat(path)
	if os.IsExist(err) {
		return true
	}
	if err != nil {
		return false
	}
	return info.IsDir()
}

func GetPackageName() string {
	pubspecPath := "./pubspec.yaml"
	// Open the pubspec.yaml file
	file, err := os.Open(pubspecPath)
	if err != nil {
		errorStr := fmt.Errorf("failed to open pubspec.yaml: %w", err)
		fmt.Println(errorStr)
		return ""
	}
	defer file.Close()

	// Parse the YAML content
	var pubspec Pubspec
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&pubspec); err != nil {
		errorStr := fmt.Errorf("failed to parse pubspec.yaml: %w", err)
		fmt.Println(errorStr)
		return ""
	}

	if pubspec.Name == "" {
		errorStr := fmt.Errorf("package name not found in pubspec.yaml")
		fmt.Println(errorStr)
		return ""
	}

	return pubspec.Name
}

// Structure to hold parsed pubspec.yaml data
type Pubspec struct {
	Name string `yaml:"name"`
}
