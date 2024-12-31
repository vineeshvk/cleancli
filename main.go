package main

import (
	"fmt"

	"github.com/vineeshvk/cleancli/constants"
	"github.com/vineeshvk/cleancli/dirvalid"
	"github.com/vineeshvk/cleancli/input"
	"github.com/vineeshvk/cleancli/models"
	"github.com/vineeshvk/cleancli/write"
)

func main() {
	var mainDirectoryModel models.MainDirectoryModel = dirvalid.ValidateRootDirectories()

	println(mainDirectoryModel.DataDir, mainDirectoryModel.DomainDir)
	var apiInfo models.ApiInfoModel = input.GetAPIInfos(mainDirectoryModel)

	write.WriteApiService(mainDirectoryModel.GetApiServiceRoute(), apiInfo)
	write.WriteDataSource(mainDirectoryModel.DataDir, apiInfo)
	write.WriteRepo(mainDirectoryModel, apiInfo)
	write.WriteUseCase(mainDirectoryModel, apiInfo)
	write.WriteDI(mainDirectoryModel, apiInfo)

	fmt.Println(constants.CompletedIcon, "Completed adding API.")

}
