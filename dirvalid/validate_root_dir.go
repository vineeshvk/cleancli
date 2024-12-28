package dirvalid

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/vineeshvk/cleancli/models"
	"github.com/vineeshvk/cleancli/utils"
)

func ValidateRootDirectories() models.MainDirectoryModel {
	validateRootPubSpecPresent()

	dataDir := validateDataPubSpecPresent()
	domainDir := validateDomainPubSpecPresent()

	return models.MainDirectoryModel{DataDir: dataDir, DomainDir: domainDir}

}

func validateRootPubSpecPresent() bool {
	exist := utils.DoesFileExist("pubspec.yaml")
	if !exist {
		fmt.Println("pubspec.yaml not found in the root of the project")
		os.Exit(1)
	}

	fmt.Println("pubspec.yaml found in the root of the project")

	return exist
}

func validateDataPubSpecPresent() string {
	dataExist, dataDir := utils.DoesFileExistByRegex(".", `data/pubspec.yaml$`)
	if !dataExist {
		fmt.Println("pubspec.yaml not found in the data module")
		os.Exit(1)
	}

	fmt.Println("pubspec.yaml found in the data module at: ", dataDir)

	return filepath.Dir(dataDir)
}

func validateDomainPubSpecPresent() string {
	domainExist, domainDir := utils.DoesFileExistByRegex(".", `domain/pubspec.yaml$`)
	if !domainExist {
		fmt.Println("pubspec.yaml not found in the domain module")
		os.Exit(1)
	}

	fmt.Println("pubspec.yaml found in the domain module at: ", domainDir)

	return filepath.Dir(domainDir)
}
