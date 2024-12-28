package write

import (
	"fmt"
	"os"

	"github.com/vineeshvk/cleancli/models"
	"github.com/vineeshvk/cleancli/template"
	"github.com/vineeshvk/cleancli/utils"
)

func WriteApiService(apiServiceDir string, apiInfo models.ApiInfoModel) {

	reqImport, resImport := apiInfo.GetRequestResponseImport()
	importString := reqImport + "\n" + resImport

	responseClassName := apiInfo.ApiClassNameValue.ResponseModelClassName
	requestClassName := apiInfo.ApiClassNameValue.RequestModelClassName

	if requestClassName != "" {
		requestClassName = apiInfo.GetMethodAnnotation() + requestClassName + " data"
	}

	// TODO: Handle path params
	apiString := fmt.Sprintf(
		template.ApiServiceFunction,
		apiInfo.Method,
		apiInfo.ApiUrl,
		responseClassName,
		apiInfo.Name,
		requestClassName,
	)

	err := utils.InsertToFileBeforeLastBrace(apiServiceDir, apiString, importString)

	if err != nil {
		fmt.Printf("Couldn't write to file : %s", err.Error())
		os.Exit(1)
	}

}
