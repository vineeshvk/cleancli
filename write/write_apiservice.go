package write

import (
	"fmt"
	"os"
	"strings"

	"github.com/vineeshvk/cleancli/constants"
	"github.com/vineeshvk/cleancli/models"
	"github.com/vineeshvk/cleancli/templates"
	"github.com/vineeshvk/cleancli/utils"
)

func WriteApiService(apiServiceDir string, apiInfo models.ApiInfoModel) {

	fmt.Println(constants.LoadingIcon, " Working on api_service file...")

	reqImport, resImport := apiInfo.GetRequestResponseImport()
	importString := reqImport + "\n" + resImport

	responseClassName := apiInfo.ApiClassNameValue.ResponseModelClassName
	requestClassName := apiInfo.ApiClassNameValue.RequestModelClassName

	params := getApiServiceParamsRequest(apiInfo)

	if requestClassName != "" {
		params += apiInfo.GetMethodAnnotation() + " " + requestClassName + " data"
	}

	// TODO: Handle path params
	apiString := fmt.Sprintf(
		templates.ApiServiceFunction,
		apiInfo.Method,
		apiInfo.ApiUrl,
		responseClassName,
		apiInfo.FunctionName,
		strings.TrimSpace(params),
	)

	err := utils.InsertToFileBeforeLastBrace(apiServiceDir, apiString, importString)

	if err != nil {
		fmt.Printf("Couldn't write to file : %s", err.Error())
		os.Exit(1)
	}

	fmt.Println()

}

func getApiServiceParamsRequest(apiInfo models.ApiInfoModel) string {
	params := ""

	if len(apiInfo.GetPathParams()) > 0 {
		for _, v := range apiInfo.GetPathParams() {
			params += fmt.Sprintf(`@Path("%s") %s, `, v, utils.SnakeCaseToCamelCase(v))
		}
	}

	return params
}
