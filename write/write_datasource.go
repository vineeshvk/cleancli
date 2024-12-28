package write

import (
	"fmt"

	"github.com/vineeshvk/cleancli/constants"
	"github.com/vineeshvk/cleancli/models"
	"github.com/vineeshvk/cleancli/template"
	"github.com/vineeshvk/cleancli/utils"
)

func WriteDataSource(dataDir string, apiInfo models.ApiInfoModel) {
	var dataSourceDir = dataDir + constants.ApiDataSourcePath
	writeDataSourceAbstractFile(dataSourceDir, apiInfo)
	writeDataSourceImplFile(dataSourceDir, apiInfo)
}

func writeDataSourceAbstractFile(dataSourceDir string, apiInfo models.ApiInfoModel) {
	dataSourceFilePath := dataSourceDir + apiInfo.GroupName + "/" + apiInfo.GroupName + "_data_source.dart"

	dataSourceFileClassText := fmt.Sprintf(template.DataSourceFileClass, apiInfo.GetApiClassName())

	utils.CreateAndInsertIfFileNotExist(dataSourceFilePath, dataSourceFileClassText)

	responseClassName := apiInfo.ApiClassNameValue.ResponseModelClassName
	requestClassName := apiInfo.ApiClassNameValue.RequestModelClassName
	if requestClassName != "" {
		requestClassName = requestClassName + " data"
	}

	dataSourceFileFunction := fmt.Sprintf(
		template.DataSourceFileFunction,
		responseClassName,
		apiInfo.Name,
		requestClassName,
	)

	utils.InsertToFileBeforeLastBrace(
		dataSourceFilePath,
		dataSourceFileFunction,
		getDataSourceImportString(apiInfo),
	)

}

func writeDataSourceImplFile(dataSourceDir string, apiInfo models.ApiInfoModel) {
	dataSourceImplFilePath := dataSourceDir + apiInfo.GroupName + "/remote/" + apiInfo.GroupName + "_data_source_impl.dart"

	dataSourceImplFileClassText := fmt.Sprintf(template.DataSourceImplFileClass, apiInfo.GroupName, apiInfo.GetApiClassName())

	utils.CreateAndInsertIfFileNotExist(dataSourceImplFilePath, dataSourceImplFileClassText)

	responseClassName := apiInfo.ApiClassNameValue.ResponseModelClassName
	requestClassName := apiInfo.ApiClassNameValue.RequestModelClassName
	var params string
	if requestClassName != "" {
		requestClassName = requestClassName + " data"
		params = "data"
	}

	dataSourceImplFileFunction := fmt.Sprintf(
		template.DataSourceImplFileFunction,
		responseClassName,
		apiInfo.Name,
		requestClassName,
		params,
	)

	utils.InsertToFileBeforeLastBrace(
		dataSourceImplFilePath,
		dataSourceImplFileFunction,
		getDataSourceImportString(apiInfo),
	)

}

func getDataSourceImportString(apiInfo models.ApiInfoModel) string {
	reqImport, resImport := apiInfo.GetRequestResponseImport()
	return reqImport + "\n" + resImport
}
