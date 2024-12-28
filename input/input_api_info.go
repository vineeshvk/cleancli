package input

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"
	"github.com/vineeshvk/cleancli/constants"
	"github.com/vineeshvk/cleancli/models"
	"github.com/vineeshvk/cleancli/readdir"
	"github.com/vineeshvk/cleancli/utils"
)

func GetAPIInfos(mainDirModel models.MainDirectoryModel) models.ApiInfoModel {

	apiUrl := getAPIUrlInput()
	method := getAPIMethodType()
	name := getAPIName()
	responseModelPath := getApiResponseModel(mainDirModel.DataDir)
	requestModelPath := getApiRequestModel(mainDirModel.DataDir)

	groupName := getApiGroupName()

	apiInfo := models.ApiInfoModel{
		ApiUrl:            apiUrl,
		Method:            method,
		Name:              name,
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

	fmt.Println("> API url set to " + result)

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

	fmt.Println("> API method set to " + result)

	return utils.RemoveStringSpaces(result)

}

func getAPIName() string {
	prompt := promptui.Prompt{
		Label:       "Enter Name for Function (Will be used in all the files)",
		AllowEdit:   true,
		Validate:    utils.ValidateEmptyString,
		HideEntered: true,
	}

	result, err := prompt.Run()

	if err != nil {
		os.Exit(1)
	}

	fmt.Println("> Function Name is set to " + result)

	result = utils.ConvertToCamelCase(utils.RemoveStringSpaces(result))

	return result
}

func getApiResponseModel(dataDir string) string {

	responseModelList := readdir.ReadApiResponseModels(dataDir)
	var responseFileList []string

	for _, responseDir := range responseModelList {
		responseFileList = append(responseFileList, filepath.Base(responseDir))
	}

	promptSelect := promptui.Select{
		Items:        responseFileList,
		Label:        "Select the Response model from List",
		Size:         5,
		HideSelected: true,
	}

	selectedIndex, result, err := promptSelect.Run()

	if err != nil {
		os.Exit(1)
	}

	selectedResponsePath := responseModelList[selectedIndex]

	fmt.Println("> Response model is set to " + result)

	return selectedResponsePath
}

func getApiRequestModel(dataDir string) string {

	requestModelList := readdir.ReadApiRequestModels(dataDir)
	var selectInputList []string = []string{"None"}

	for _, requestDir := range requestModelList {
		selectInputList = append(selectInputList, filepath.Base(requestDir))
	}

	promptSelect := promptui.Select{
		Items:        selectInputList,
		Label:        "Select the Request model from List",
		Size:         5,
		HideSelected: true,
	}

	selectedIndex, result, err := promptSelect.Run()

	if result == "None" {
		return ""
	}

	// Because "None" is added as extra in requestFileList
	selectedIndex -= 1

	if err != nil {
		os.Exit(1)
	}

	selectedRequestPath := requestModelList[selectedIndex]

	fmt.Println("> Request model is set to " + result)

	return selectedRequestPath
}
