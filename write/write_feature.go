package write

import (
	"fmt"
	"os"
	"strings"

	"github.com/vineeshvk/cleancli/constants"
	"github.com/vineeshvk/cleancli/models"
	"github.com/vineeshvk/cleancli/templates"
	"github.com/vineeshvk/cleancli/utils"
)

func WriteFeatureDI(featureInfo models.FeatureInfoModel, packageName string) {

	fmt.Println(constants.LoadingIcon, "Analyzing feature ...")

	featureNameCamelCase := utils.SnakeCaseToCamelCase(featureInfo.FeatureName)

	// 1: Directory or package name, 2: Group Name, 3: Provider Name, 4: Group Class Name
	moduleString := fmt.Sprintf(
		templates.FeatureDI,
		packageName,
		featureInfo.FeatureName,
		featureNameCamelCase,
		utils.SnakeCaseToPascalCase(featureInfo.FeatureName),
	)

	err := utils.CreateAndInsertIfFileNotExist(constants.FeatureDIPath+featureInfo.FeatureName+"_module.dart", moduleString)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println()
}

func WriteFeatureRoute(featureInfo models.FeatureInfoModel, packageName string) {
	// Define the file path
	filePath := constants.LocationsPath

	writeInBeamLocation(filePath, featureInfo)

	importString := fmt.Sprintf(templates.FeatureRouteImports, packageName, featureInfo.FeatureName)

	className := utils.SnakeCaseToPascalCase(featureInfo.FeatureName)
	name := utils.SnakeToName(featureInfo.FeatureName)

	routeFunction := fmt.Sprintf(
		templates.FeatureRoutes,
		className,
		name,
		featureInfo.FeatureName,
	)

	utils.AppendToFile(filePath, routeFunction, importString)

}

func WriteFeaturePages(featureInfo models.FeatureInfoModel, packageName string) {
	filePath := constants.FeaturePagesPath + featureInfo.FeatureName + "/"
	writeFeaturePage(filePath, packageName, featureInfo.FeatureName)
	writeFeaturePageView(filePath, packageName, featureInfo.FeatureName)
	writeFeaturePageViewModel(filePath, packageName, featureInfo.FeatureName)
}

func writeFeaturePage(baseFilePath string, packageName string, featureName string) {
	filePath := baseFilePath + featureName + "_page.dart"
	className := utils.SnakeCaseToPascalCase(featureName)
	variableName := utils.SnakeCaseToCamelCase(featureName)

	content := fmt.Sprintf(
		templates.FeaturePage,
		packageName,
		featureName,
		className,
		variableName,
	)

	utils.CreateAndInsertIfFileNotExist(filePath, content)
}

func writeFeaturePageView(baseFilePath string, packageName string, featureName string) {
	filePath := baseFilePath + featureName + "_page_view.dart"
	className := utils.SnakeCaseToPascalCase(featureName)

	content := fmt.Sprintf(
		templates.FeaturePageView,
		packageName,
		featureName,
		className,
	)

	utils.CreateAndInsertIfFileNotExist(filePath, content)
}

func writeFeaturePageViewModel(baseFilePath string, packageName string, featureName string) {
	filePath := baseFilePath + featureName + "_page_view_model.dart"
	className := utils.SnakeCaseToPascalCase(featureName)

	content := fmt.Sprintf(
		templates.FeaturePageViewModel,
		packageName,
		className,
	)

	utils.CreateAndInsertIfFileNotExist(filePath, content)
}

func writeInBeamLocation(filePath string, featureInfo models.FeatureInfoModel) {
	// Read the file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Convert content to a string
	contentStr := string(content)

	// Find the beamLocations list
	startIndex := strings.Index(contentStr, "beamLocations: [")
	if startIndex == -1 {
		fmt.Println("beamLocations list not found!")
		return
	}

	// Locate the position to insert the new location
	endIndex := strings.Index(contentStr[startIndex:], "]")
	if endIndex == -1 {
		fmt.Println("beamLocations list is malformed!")
		return
	}
	endIndex += startIndex

	newLocation := utils.SnakeCaseToPascalCase(featureInfo.FeatureName) + "Location()"

	// Prepare the updated content
	updatedContent := contentStr[:endIndex] +
		"  " + newLocation + ", \n    " +
		contentStr[endIndex:]

	// Write the updated content back to the file
	err = os.WriteFile(filePath, []byte(updatedContent), 0644)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}

	fmt.Println("New location added successfully!")
}
