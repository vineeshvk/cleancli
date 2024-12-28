package input

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/vineeshvk/cleancli/constants"
	"github.com/vineeshvk/cleancli/models"
	"github.com/vineeshvk/cleancli/utils"
)

func GetAPIInfos(mainDirModel models.MainDirectoryModel) models.ApiInfoModel {

	fmt.Println(constants.LoadingIcon, " Getting API Details:")

	apiUrl := getAPIUrlInput()
	method := getAPIMethodType()
	name := getAPIName()
	responseModelPath := getApiResponseModel(mainDirModel.DataDir)
	requestModelPath := getApiRequestModel(mainDirModel.DataDir)

	groupName := getApiGroupName()

	apiInfo := models.ApiInfoModel{
		ApiUrl:            apiUrl,
		Method:            method,
		FunctionName:      name,
		ResponseModelPath: responseModelPath,
		RequestModelPath:  requestModelPath,
		GroupName:         groupName,
	}

	apiInfo.FindApiClassNames()
	return apiInfo
}

func getAPIUrlInput() string {
	prompt := promptui.Prompt{
		Label:       "Enter API Url",
		Default:     "v1/",
		AllowEdit:   true,
		Validate:    utils.ValidateEmptyString,
		HideEntered: true,
	}

	result, err := prompt.Run()

	if err != nil {
		os.Exit(1)
	}

	fmt.Println(constants.InputSetIcon, " API url set to "+result)

	return utils.RemoveStringSpaces(result)
}

func getAPIMethodType() string {
	promptSelect := promptui.Select{
		Items:        constants.APIMethods,
		Label:        "Select the API Method",
		Size:         5,
		HideSelected: true,
	}

	_, result, err := promptSelect.Run()

	if err != nil {
		os.Exit(1)
	}

	fmt.Println(constants.InputSetIcon, " API method set to "+result)

	return utils.RemoveStringSpaces(result)

}

func getAPIName() string {
	prompt := promptui.Prompt{
		Label:       "Enter Name for Function",
		AllowEdit:   true,
		Validate:    utils.ValidateEmptyString,
		HideEntered: true,
	}

	result, err := prompt.Run()

	if err != nil {
		os.Exit(1)
	}

	fmt.Println(constants.InputSetIcon, " Function Name is set to "+result)

	result = utils.ConvertToCamelCase(utils.RemoveStringSpaces(result))

	return result
}
