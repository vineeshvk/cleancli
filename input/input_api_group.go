package input

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/vineeshvk/cleancli/constants"
	"github.com/vineeshvk/cleancli/utils"
)

func getApiGroupName() string {
	prompt := promptui.Prompt{
		Label:       "Enter Name for Group (Snake case) (Will be folder/file names of ds, repo, usecase)",
		AllowEdit:   true,
		Validate:    utils.ValidateEmptyString,
		HideEntered: true,
	}

	result, err := prompt.Run()

	if err != nil {
		os.Exit(1)
	}

	fmt.Println(constants.InputSetIcon, " Group Name is set to "+result)

	result = utils.RemoveStringSpaces(result)

	return result
}
