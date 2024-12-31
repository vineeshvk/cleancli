package write

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/vineeshvk/cleancli/constants"
	"github.com/vineeshvk/cleancli/models"
	"github.com/vineeshvk/cleancli/templates"
	"github.com/vineeshvk/cleancli/utils"
)

func WriteRepo(mainDirModel models.MainDirectoryModel, apiInfo models.ApiInfoModel) {
	dataRepoDir := filepath.Join(mainDirModel.DataDir, constants.ApiRepositoryPath)
	domainRepoDir := filepath.Join(mainDirModel.DomainDir, constants.ApiRepositoryPath)

	fmt.Println(constants.LoadingIcon, " Working on repository files...")

	writeRepoAbstractFile(mainDirModel.DomainDir, domainRepoDir, apiInfo)
	writeRepoImplFile(mainDirModel, dataRepoDir, apiInfo)
}

func writeRepoAbstractFile(domainDir string, domainRepoDir string, apiInfo models.ApiInfoModel) {
	domainRepoFilePath := filepath.Join(
		domainRepoDir,
		apiInfo.GroupName+"_repository.dart",
	)

	repoFileClassText := fmt.Sprintf(templates.RepoFileClass, apiInfo.GetApiClassName(), domainDir)

	err := utils.CreateAndInsertIfFileNotExist(domainRepoFilePath, repoFileClassText)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println()

	responseClassName := apiInfo.ApiClassNameValue.ResponseModelClassName
	requestClassName := apiInfo.ApiClassNameValue.RequestModelClassName
	if requestClassName != "" {
		requestClassName = requestClassName + " data"
	}

	repoFileFunction := fmt.Sprintf(
		templates.RepoFileFunction,
		responseClassName,
		apiInfo.FunctionName,
		requestClassName,
	)

	err = utils.InsertToFileBeforeLastBrace(
		domainRepoFilePath,
		repoFileFunction,
		getReqResImportString(apiInfo),
	)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println()

}

func writeRepoImplFile(mainDirModel models.MainDirectoryModel, dataRepoDir string, apiInfo models.ApiInfoModel) {
	dataRepoImplFilePath := filepath.Join(
		dataRepoDir,
		apiInfo.GroupName+"_repository_impl.dart",
	)

	repoImplFileClassText := fmt.Sprintf(
		templates.RepoImplFileClass,
		apiInfo.GroupName,
		apiInfo.GetApiClassName(),
		mainDirModel.DataDir,
		mainDirModel.DomainDir,
	)

	err := utils.CreateAndInsertIfFileNotExist(dataRepoImplFilePath, repoImplFileClassText)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println()

	responseClassName := apiInfo.ApiClassNameValue.ResponseModelClassName
	requestClassName := apiInfo.ApiClassNameValue.RequestModelClassName
	var params string
	if requestClassName != "" {
		requestClassName = requestClassName + " data"
		params = "data"
	}

	repoFileFunction := fmt.Sprintf(
		templates.RepoImplFileFunction,
		responseClassName,
		apiInfo.FunctionName,
		requestClassName,
		params,
	)

	err = utils.InsertToFileBeforeLastBrace(
		dataRepoImplFilePath,
		repoFileFunction,
		getReqResImportString(apiInfo),
	)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println()

}
