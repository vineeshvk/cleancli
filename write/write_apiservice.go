package write

import (
	"fmt"
	"os"

	"github.com/vineeshvk/cleancli/constants"
	"github.com/vineeshvk/cleancli/models"
	"github.com/vineeshvk/cleancli/templates"
	"github.com/vineeshvk/cleancli/utils"
)

func WriteApiService(apiServiceDir string, apiInfo models.ApiInfoModel) {

	fmt.Sprintln(constants.LoadingIcon, " Working on api_service file...")

	reqImport, resImport := apiInfo.GetRequestResponseImport()
	importString := reqImport + "\n" + resImport

	responseClassName := apiInfo.ApiClassNameValue.ResponseModelClassName
	requestClassName := apiInfo.ApiClassNameValue.RequestModelClassName

	if requestClassName != "" {
		requestClassName = apiInfo.GetMethodAnnotation() + requestClassName + " data"
	}

	// TODO: Handle path params
	apiString := fmt.Sprintf(
		templates.ApiServiceFunction,
		apiInfo.Method,
		apiInfo.ApiUrl,
		responseClassName,
		apiInfo.FunctionName,
		requestClassName,
	)

	err := utils.InsertToFileBeforeLastBrace(apiServiceDir, apiString, importString)

	if err != nil {
		fmt.Printf("Couldn't write to file : %s", err.Error())
		os.Exit(1)
	}

}
