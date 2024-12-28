package input

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"
	"github.com/vineeshvk/cleancli/constants"
	"github.com/vineeshvk/cleancli/readdir"
	"github.com/vineeshvk/cleancli/utils"
)

func getApiResponseModel(dataDir string) string {

	responseModelList := readdir.ReadApiResponseModels(dataDir)
	var responseFileList []string

	for _, responseDir := range responseModelList {
		responseFileList = append(responseFileList, filepath.Base(responseDir))
	}

	promptSelect := promptui.Select{
		Items:             responseFileList,
		Label:             "Select the Response model from List",
		Size:              10,
		StartInSearchMode: true,
		HideSelected:      true,
		Searcher: func(input string, index int) bool {
			return utils.Search(responseFileList[index], input)
		},
	}

	selectedIndex, result, err := promptSelect.Run()

	if err != nil {
		os.Exit(1)
	}

	selectedResponsePath := responseModelList[selectedIndex]

	fmt.Println(constants.InputSetIcon, " Response model is set to "+result)

	return selectedResponsePath
}

func getApiRequestModel(dataDir string) string {

	requestModelList := readdir.ReadApiRequestModels(dataDir)
	var selectInputList []string = []string{"None"}

	for _, requestDir := range requestModelList {
		selectInputList = append(selectInputList, filepath.Base(requestDir))
	}

	promptSelect := promptui.Select{
		Items:             selectInputList,
		Label:             "Select the Request model from List",
		Size:              10,
		HideSelected:      true,
		StartInSearchMode: true,
		Searcher: func(input string, index int) bool {
			return utils.Search(selectInputList[index], input)
		},
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

	fmt.Println(constants.InputSetIcon, " Request model is set to "+result)

	return selectedRequestPath
}
