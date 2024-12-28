package main

import (
	"github.com/vineeshvk/cleancli/dirvalid"
	"github.com/vineeshvk/cleancli/input"
	"github.com/vineeshvk/cleancli/models"
	"github.com/vineeshvk/cleancli/write"
)

func main() {
	var mainDirectoryModel models.MainDirectoryModel = dirvalid.ValidateRootDirectories()
	var apiInfo models.ApiInfoModel = input.GetAPIInfos(mainDirectoryModel)
	write.WriteApiService(mainDirectoryModel.GetApiServiceRoute(), apiInfo)
	write.WriteDataSource(mainDirectoryModel.DataDir, apiInfo)
}
