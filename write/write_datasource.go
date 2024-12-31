package write

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/vineeshvk/cleancli/constants"
	"github.com/vineeshvk/cleancli/models"
	"github.com/vineeshvk/cleancli/templates"
	"github.com/vineeshvk/cleancli/utils"
)

func WriteDataSource(dataDir string, apiInfo models.ApiInfoModel) {
	var dataSourceDir = filepath.Join(dataDir, constants.ApiDataSourcePath)

	fmt.Println(constants.LoadingIcon, " Working on data source files...")

	writeDataSourceAbstractFile(dataSourceDir, apiInfo)
	writeDataSourceImplFile(dataDir, dataSourceDir, apiInfo)
}

func writeDataSourceAbstractFile(dataSourceDir string, apiInfo models.ApiInfoModel) {
	dataSourceFilePath := filepath.Join(
		dataSourceDir,
		apiInfo.GroupName,
		apiInfo.GroupName+"_data_source.dart",
	)

	dataSourceFileClassText := fmt.Sprintf(templates.DataSourceFileClass, apiInfo.GetApiClassName())

	err := utils.CreateAndInsertIfFileNotExist(dataSourceFilePath, dataSourceFileClassText)

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

	dataSourceFileFunction := fmt.Sprintf(
		templates.DataSourceFileFunction,
		responseClassName,
		apiInfo.FunctionName,
		requestClassName,
	)

	err = utils.InsertToFileBeforeLastBrace(
		dataSourceFilePath,
		dataSourceFileFunction,
		getReqResImportString(apiInfo),
	)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println()

}

func writeDataSourceImplFile(dataDir string, dataSourceDir string, apiInfo models.ApiInfoModel) {

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
		dataDir,
	)

	err := utils.CreateAndInsertIfFileNotExist(dataSourceImplFilePath, dataSourceImplFileClassText)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println()

	responseClassName := apiInfo.ApiClassNameValue.ResponseModelClassName
	requestClassName := apiInfo.ApiClassNameValue.RequestModelClassName

	params := getDSParamsRequest(apiInfo)

	if requestClassName != "" {
		requestClassName = requestClassName + " data"
		params += "data"
	}

	dataSourceImplFileFunction := fmt.Sprintf(
		templates.DataSourceImplFileFunction,
		responseClassName,
		apiInfo.FunctionName,
		requestClassName,
		strings.TrimSpace(params),
	)

	err = utils.InsertToFileBeforeLastBrace(
		dataSourceImplFilePath,
		dataSourceImplFileFunction,
		getReqResImportString(apiInfo),
	)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println()

}

func getReqResImportString(apiInfo models.ApiInfoModel) string {
	reqImport, resImport := apiInfo.GetRequestResponseImport()
	return reqImport + "\n" + resImport
}

func getDSParamsRequest(apiInfo models.ApiInfoModel) string {
	params := ""

	if len(apiInfo.GetPathParams()) > 0 {
		for _, v := range apiInfo.GetPathParams() {
			params += fmt.Sprintf(`data.%s, `, utils.SnakeCaseToCamelCase(v))
		}
	}

	return params
}
