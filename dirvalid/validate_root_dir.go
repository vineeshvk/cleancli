package dirvalid

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/vineeshvk/cleancli/constants"
	"github.com/vineeshvk/cleancli/models"
	"github.com/vineeshvk/cleancli/utils"
)

func ValidateRootDirectories() models.MainDirectoryModel {
	fmt.Println(constants.EvalutingIcon, " Evaluating flutter project...")

	validateRootPubSpecPresent()

	dataDir := validateDataPubSpecPresent()
	domainDir := validateDomainPubSpecPresent()
	packageName := utils.GetPackageName()

	return models.MainDirectoryModel{DataDir: dataDir, DomainDir: domainDir, PackageName: packageName}

}

func validateRootPubSpecPresent() bool {
	exist := utils.DoesFileExist("pubspec.yaml")
	if !exist {
		fmt.Println("pubspec.yaml not found in the root of the project")
		os.Exit(1)
	}

	fmt.Println(constants.SuccessIcon, " pubspec.yaml found in the root of the project")

	return exist
}

func validateDataPubSpecPresent() string {
	dataExist, dataDir := utils.DoesFileExistByRegex(".", `data/pubspec.yaml$`)
	if !dataExist {
		fmt.Println("pubspec.yaml not found in the data module")
		os.Exit(1)
	}

	fmt.Println(constants.SuccessIcon, " pubspec.yaml found in the data module at: ", dataDir)

	return filepath.Dir(dataDir)
}

func validateDomainPubSpecPresent() string {
	domainExist, domainDir := utils.DoesFileExistByRegex(".", `domain/pubspec.yaml$`)
	if !domainExist {
		fmt.Println("pubspec.yaml not found in the domain module")
		os.Exit(1)
	}

	fmt.Println(constants.SuccessIcon, " pubspec.yaml found in the domain module at: ", domainDir)

	return filepath.Dir(domainDir)
}
