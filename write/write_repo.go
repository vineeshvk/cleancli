package write

import (
	"fmt"
	"path/filepath"

	"github.com/vineeshvk/cleancli/constants"
	"github.com/vineeshvk/cleancli/models"
	"github.com/vineeshvk/cleancli/templates"
	"github.com/vineeshvk/cleancli/utils"
)

func WriteRepo(mainDirModel models.MainDirectoryModel, apiInfo models.ApiInfoModel) {
	dataRepoDir := mainDirModel.DataDir + constants.ApiRepositoryPath
	domainRepoDir := mainDirModel.DomainDir + constants.ApiRepositoryPath

	fmt.Sprintln(constants.LoadingIcon, " Working on repository files...")

	writeRepoAbstractFile(domainRepoDir, apiInfo)
	writeRepoImplFile(dataRepoDir, apiInfo)
}

func writeRepoAbstractFile(domainRepoDir string, apiInfo models.ApiInfoModel) {
	dataRepoFilePath := filepath.Join(
		domainRepoDir,
		apiInfo.GroupName+"_repository.dart",
	)

	repoFileClassText := fmt.Sprintf(templates.RepoFileClass, apiInfo.GetApiClassName())

	utils.CreateAndInsertIfFileNotExist(dataRepoFilePath, repoFileClassText)

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

	utils.InsertToFileBeforeLastBrace(
		dataRepoFilePath,
		repoFileFunction,
		getReqResImportString(apiInfo),
	)

	fmt.Println()

}

func writeRepoImplFile(dataRepoDir string, apiInfo models.ApiInfoModel) {
	dataRepoImplFilePath := filepath.Join(
		dataRepoDir,
		apiInfo.GroupName+"_repository_impl.dart",
	)

	repoImplFileClassText := fmt.Sprintf(templates.RepoImplFileClass, apiInfo.GroupName, apiInfo.GetApiClassName())

	utils.CreateAndInsertIfFileNotExist(dataRepoImplFilePath, repoImplFileClassText)
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

	utils.InsertToFileBeforeLastBrace(
		dataRepoImplFilePath,
		repoFileFunction,
		getReqResImportString(apiInfo),
	)
	fmt.Println()

}
