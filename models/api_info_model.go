package models

import "github.com/vineeshvk/cleancli/utils"

type ApiInfoModel struct {
	ApiUrl string
	Method string
	Name   string

	ResponseModelPath string
	RequestModelPath  string

	GroupName string

	ApiClassNameValue ApiClassNamesModel
}

type ApiClassNamesModel struct {
	ResponseModelClassName string
	RequestModelClassName  string
}

func (model ApiInfoModel) GetMethodAnnotation() string {
	if model.RequestModelPath == "" {
		return ""
	}

	if model.Method == "Get" {
		return "@Queries()"
	}

	return "@Body()"

}

func (apiInfo *ApiInfoModel) FindApiClassNames() {
	responseClassName := utils.GetClassNameFromFile(apiInfo.ResponseModelPath)
	var requestClassName string

	if apiInfo.RequestModelPath != "" {
		requestClassName = utils.GetClassNameFromFile(apiInfo.RequestModelPath)
	}

	apiInfo.ApiClassNameValue = ApiClassNamesModel{
		ResponseModelClassName: responseClassName,
		RequestModelClassName:  requestClassName,
	}
}

func (apiInfo ApiInfoModel) GetApiClassName() string {
	return utils.SnakeCaseToPascalCase(apiInfo.GroupName)
}

func (apiInfo ApiInfoModel) GetRequestResponseImport() (req string, res string) {
	var reqImport string
	if apiInfo.RequestModelPath != "" {
		reqImport = utils.GetImportRoute(apiInfo.RequestModelPath)
	}
	return reqImport, utils.GetImportRoute(apiInfo.ResponseModelPath)
}
