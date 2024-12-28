package write

import (
	"fmt"
	"path/filepath"

	"github.com/vineeshvk/cleancli/constants"
	"github.com/vineeshvk/cleancli/models"
	"github.com/vineeshvk/cleancli/templates"
	"github.com/vineeshvk/cleancli/utils"
)

func WriteDataSource(dataDir string, apiInfo models.ApiInfoModel) {
	var dataSourceDir = filepath.Join(dataDir, constants.ApiDataSourcePath)

	fmt.Sprintln(constants.LoadingIcon, " Working on data source files...")

	writeDataSourceAbstractFile(dataSourceDir, apiInfo)
	writeDataSourceImplFile(dataSourceDir, apiInfo)
}

func writeDataSourceAbstractFile(dataSourceDir string, apiInfo models.ApiInfoModel) {
	dataSourceFilePath := filepath.Join(
		dataSourceDir,
		apiInfo.GroupName,
		apiInfo.GroupName+"_data_source.dart",
	)

	dataSourceFileClassText := fmt.Sprintf(templates.DataSourceFileClass, apiInfo.GetApiClassName())

	utils.CreateAndInsertIfFileNotExist(dataSourceFilePath, dataSourceFileClassText)
	fmt.Println()

	responseClassName := apiInfo.ApiClassNameValue.ResponseModelClassName
	requestClassName := apiInfo.ApiClassNameValue.RequestModelClassName
	if requestClassName != "" {
		requestClassName = requestClassName + " data"
	}

	dataSourceFileFunction := fmt.Sprintf(
		templates.DataSourceFileFunction,
		responseClassName,
		apiInfo.FunctionName,
		requestClassName,
	)

	utils.InsertToFileBeforeLastBrace(
		dataSourceFilePath,
		dataSourceFileFunction,
		getReqResImportString(apiInfo),
	)
	fmt.Println()

}

func writeDataSourceImplFile(dataSourceDir string, apiInfo models.ApiInfoModel) {
	dataSourceImplFilePath := filepath.Join(
		dataSourceDir,
		apiInfo.GroupName,
		"/remote/",
		apiInfo.GroupName+"_data_source_impl.dart",
	)

	dataSourceImplFileClassText := fmt.Sprintf(
		templates.DataSourceImplFileClass,
		apiInfo.GroupName,
		apiInfo.GetApiClassName(),
	)

	utils.CreateAndInsertIfFileNotExist(dataSourceImplFilePath, dataSourceImplFileClassText)
	fmt.Println()

	responseClassName := apiInfo.ApiClassNameValue.ResponseModelClassName
	requestClassName := apiInfo.ApiClassNameValue.RequestModelClassName
	var params string
	if requestClassName != "" {
		requestClassName = requestClassName + " data"
		params = "data"
	}

	dataSourceImplFileFunction := fmt.Sprintf(
		templates.DataSourceImplFileFunction,
		responseClassName,
		apiInfo.FunctionName,
		requestClassName,
		params,
	)

	utils.InsertToFileBeforeLastBrace(
		dataSourceImplFilePath,
		dataSourceImplFileFunction,
		getReqResImportString(apiInfo),
	)
	fmt.Println()

}

func getReqResImportString(apiInfo models.ApiInfoModel) string {
	reqImport, resImport := apiInfo.GetRequestResponseImport()
	return reqImport + "\n" + resImport
}
