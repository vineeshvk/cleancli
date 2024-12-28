package write

import (
	"fmt"
	"path/filepath"

	"github.com/vineeshvk/cleancli/constants"
	"github.com/vineeshvk/cleancli/models"
	"github.com/vineeshvk/cleancli/templates"
	"github.com/vineeshvk/cleancli/utils"
)

func WriteUseCase(mainDirModel models.MainDirectoryModel, apiInfo models.ApiInfoModel) {
	fmt.Sprintln(constants.LoadingIcon, " Working on usecase file...")

	useCasePath := filepath.Join(
		mainDirModel.DomainDir,
		constants.UseCasePath,
		apiInfo.GroupName,
		utils.CamelToSnake(apiInfo.FunctionName)+"_usecase.dart",
	)

	var useCaseParamFuncString string
	var paramsPassingString string

	if apiInfo.RequestModelPath != "" {
		paramsPassingString = "params.toRequest()"

		useCaseParamFuncString = fmt.Sprintf(
			templates.UseCaseFileParamsToRequestFunction,
			apiInfo.ApiClassNameValue.RequestModelClassName,
		)
	}

	useCaseClassString := fmt.Sprintf(
		templates.UseCaseFileClass,
		getReqResImportString(apiInfo),
		utils.CapitilizeFirst(apiInfo.FunctionName),
		apiInfo.ApiClassNameValue.ResponseModelClassName,
		apiInfo.GetApiClassName(),
		apiInfo.FunctionName,
		paramsPassingString,
		useCaseParamFuncString,
		apiInfo.GroupName,
	)

	utils.CreateAndInsertIfFileNotExist(
		useCasePath,
		useCaseClassString,
	)
	fmt.Println()

}
