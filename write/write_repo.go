package write

import (
	"fmt"

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
	dataRepoFilePath := domainRepoDir + apiInfo.GroupName + "/" + apiInfo.GroupName + "_repository.dart"

	repoFileClassText := fmt.Sprintf(templates.RepoFileClass, apiInfo.GetApiClassName())

	utils.CreateAndInsertIfFileNotExist(dataRepoFilePath, repoFileClassText)

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
}

func writeRepoImplFile(dataRepoDir string, apiInfo models.ApiInfoModel) {
	dataRepoImplFilePath := dataRepoDir + apiInfo.GroupName + "/" + apiInfo.GroupName + "_repository_impl.dart"

	repoImplFileClassText := fmt.Sprintf(templates.RepoImplFileClass, apiInfo.GroupName, apiInfo.GetApiClassName())

	utils.CreateAndInsertIfFileNotExist(dataRepoImplFilePath, repoImplFileClassText)

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
}
