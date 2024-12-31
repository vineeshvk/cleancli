package write

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/vineeshvk/cleancli/constants"
	"github.com/vineeshvk/cleancli/models"
	"github.com/vineeshvk/cleancli/templates"
	"github.com/vineeshvk/cleancli/utils"
)

func WriteDI(mainDirModel models.MainDirectoryModel, apiInfo models.ApiInfoModel) {
	fmt.Println(constants.LoadingIcon, " Working on di files...")

	writeDataSourceDI(mainDirModel, apiInfo)
	writeRepoDI(mainDirModel, apiInfo)
	writeUseCaseDI(mainDirModel, apiInfo)

}

func writeDataSourceDI(mainDirModel models.MainDirectoryModel, apiInfo models.ApiInfoModel) {
	dataSourceDIDir := filepath.Join(mainDirModel.DataDir, constants.DataSourceDIPath)

	diString := fmt.Sprintf(
		templates.DataSourceDI,
		utils.SnakeCaseToCamelCase(apiInfo.GroupName),
		apiInfo.GetApiClassName(),
	)

	diImportString := getDataSourceImport(mainDirModel, apiInfo)

	err := utils.AppendToFile(
		dataSourceDIDir,
		diString,
		diImportString,
	)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

}

func writeRepoDI(mainDirModel models.MainDirectoryModel, apiInfo models.ApiInfoModel) {

	repoDIDir := filepath.Join(mainDirModel.DataDir, constants.RepoDIPath)

	diString := fmt.Sprintf(
		templates.RepoDI,
		utils.SnakeCaseToCamelCase(apiInfo.GroupName),
		apiInfo.GetApiClassName(),
	)

	diImportString := getRepoImport(mainDirModel, apiInfo)

	err := utils.AppendToFile(
		repoDIDir,
		diString,
		diImportString,
	)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

}

func writeUseCaseDI(mainDirModel models.MainDirectoryModel, apiInfo models.ApiInfoModel) {
	useCaseDiFilePath := filepath.Join(mainDirModel.DomainDir, "di", apiInfo.GroupName+"_use_case_provider.dart")

	err := utils.CreateAndInsertIfFileNotExist(
		useCaseDiFilePath,
		fmt.Sprintf(templates.UseCaseDIInitImport, mainDirModel.DataDir),
	)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	diString := fmt.Sprintf(
		templates.UseCaseDI,
		apiInfo.FunctionName,
		utils.CapitilizeFirst(apiInfo.FunctionName),
		utils.SnakeCaseToCamelCase(apiInfo.GroupName),
	)

	diImportString := getUseCaseImport(mainDirModel, apiInfo)

	err = utils.AppendToFile(
		useCaseDiFilePath,
		diString,
		diImportString,
	)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

}

func getDataSourceImport(mainDirModel models.MainDirectoryModel, apiInfo models.ApiInfoModel) string {
	dataSourceImportPath := filepath.Join(
		mainDirModel.DataDir,
		"source",
		apiInfo.GroupName,
	)

	dataSourceAbFileImport := filepath.Join(
		dataSourceImportPath,
		apiInfo.GroupName+"_data_source.dart",
	)

	dataSourceImplFileImport := filepath.Join(
		dataSourceImportPath,
		"remote",
		apiInfo.GroupName+"_data_source.dart",
	)

	return utils.GetImportRoute(dataSourceAbFileImport) + "\n" + utils.GetImportRoute(dataSourceImplFileImport)
}

func getRepoImport(mainDirModel models.MainDirectoryModel, apiInfo models.ApiInfoModel) string {
	dataRepoDir := filepath.Join(mainDirModel.DataDir, constants.ApiRepositoryPath)
	domainRepoDir := filepath.Join(mainDirModel.DomainDir, constants.ApiRepositoryPath)

	dataRepoFilePath := filepath.Join(
		domainRepoDir,
		apiInfo.GroupName+"_repository.dart",
	)

	dataRepoImplFilePath := filepath.Join(
		dataRepoDir,
		apiInfo.GroupName+"_repository_impl.dart",
	)

	return utils.GetImportRoute(dataRepoFilePath) + "\n" + utils.GetImportRoute(dataRepoImplFilePath)

}

func getUseCaseImport(mainDirModel models.MainDirectoryModel, apiInfo models.ApiInfoModel) string {
	// import 'package:domain/usecase/file_upload/file_upload_usecase.dart';
	useCaseImportPath := filepath.Join(mainDirModel.DomainDir, constants.UseCasePath, apiInfo.GroupName, apiInfo.GroupName+"_usecase.dart")

	return utils.GetImportRoute(useCaseImportPath)
}
