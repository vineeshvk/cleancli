package models

import "github.com/vineeshvk/cleancli/constants"

type MainDirectoryModel struct {
	DataDir   string
	DomainDir string
	PackageName string
}

func (dirModel MainDirectoryModel) GetApiServiceRoute() string {
	return dirModel.DataDir + constants.ApiServicePath
}

func (dirModel MainDirectoryModel) GetDataSourceRoute() string {
	return dirModel.DataDir + constants.ApiDataSourcePath
}
