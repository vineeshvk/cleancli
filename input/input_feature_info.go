package input

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/vineeshvk/cleancli/constants"
	"github.com/vineeshvk/cleancli/models"
	"github.com/vineeshvk/cleancli/utils"
)

func GetFeatureInfos() models.FeatureInfoModel {
	fmt.Println(constants.LoadingIcon, " Getting Feature Details:")

	featureName := getFeatureName()

	return models.FeatureInfoModel{
		FeatureName: featureName,
	}
}

func getFeatureName() string {
	prompt := promptui.Prompt{
		Label:       "Enter Name for Feature [snake_case]",
		AllowEdit:   true,
		Validate:    utils.ValidateEmptyString,
		HideEntered: true,
	}

	result, err := prompt.Run()

	if utils.DoesFolderExist("./lib/feature/" + result) {
		fmt.Println("Feature already exists, verify the name and try again")
		return getFeatureName()
	}

	if err != nil {
		os.Exit(1)
	}

	result = utils.ConvertToCamelCase(utils.RemoveStringSpaces(result))

	fmt.Println(constants.InputSetIcon, " Feature Name is set to "+result)

	return result
}
