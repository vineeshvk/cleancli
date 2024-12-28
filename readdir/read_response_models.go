package readdir

import (
	"path/filepath"

	"github.com/vineeshvk/cleancli/constants"
	"github.com/vineeshvk/cleancli/utils"
)

func ReadApiResponseModels(dataDir string) []string {
	responseDir := filepath.Join(dataDir, constants.ApiResponseEntityPath)

	return utils.GetMatchingFilesFromDir(responseDir, `response_entity\.dart$`)
}

func ReadApiRequestModels(dataDir string) []string {
	requestDir := filepath.Join(dataDir, constants.ApiRequestEntityPath)
	return utils.GetMatchingFilesFromDir(requestDir, `^[\w_/]+\.dart$`)
}
