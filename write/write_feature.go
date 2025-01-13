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

	fmt.Println(constants.LoadingIcon, "Creating Feature DI")

	featureNameCamelCase := utils.SnakeCaseToCamelCase(featureInfo.FeatureName)

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

	fmt.Println(constants.LoadingIcon, "Successfully created DI for "+featureInfo.FeatureName)
	fmt.Println()
}

func WriteFeatureRoute(featureInfo models.FeatureInfoModel, packageName string) {
	fmt.Println(constants.LoadingIcon, "Creating Route in locations.dart file")
	filePath := constants.LocationsPath

	// Read the file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	contentStr := string(content)

	writeInBeamLocation(filePath, featureInfo, contentStr)

	importString := fmt.Sprintf(templates.FeatureRouteImports, packageName, featureInfo.FeatureName)

	className := utils.SnakeCaseToPascalCase(featureInfo.FeatureName)
	name := utils.SnakeCaseToName(featureInfo.FeatureName)

	routeFunction := fmt.Sprintf(
		templates.FeatureRoutes,
		className,
		name,
		featureInfo.FeatureName,
	)

	utils.AppendToFile(filePath, routeFunction, importString)

	fmt.Println(constants.LoadingIcon, "Successfully added new location: "+className+"Location() in locations.dart file.")
	fmt.Println()
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
	fmt.Println(constants.LoadingIcon, "Creating Feature Page: "+className+"Page()")

	content := fmt.Sprintf(
		templates.FeaturePage,
		packageName,
		featureName,
		className,
		variableName,
	)

	err := utils.CreateAndInsertIfFileNotExist(filePath, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(constants.LoadingIcon, "Successfully created: "+className+"Page()")
	fmt.Println()
}

func writeFeaturePageView(baseFilePath string, packageName string, featureName string) {
	filePath := baseFilePath + featureName + "_page_view.dart"
	className := utils.SnakeCaseToPascalCase(featureName)

	fmt.Println(constants.LoadingIcon, "Creating Feature Page View: "+className+"PageView()")

	content := fmt.Sprintf(
		templates.FeaturePageView,
		packageName,
		featureName,
		className,
	)

	err := utils.CreateAndInsertIfFileNotExist(filePath, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(constants.LoadingIcon, "Successfully created: "+className+"PageView()")
	fmt.Println()
}

func writeFeaturePageViewModel(baseFilePath string, packageName string, featureName string) {
	filePath := baseFilePath + featureName + "_page_view_model.dart"
	className := utils.SnakeCaseToPascalCase(featureName)

	fmt.Println(constants.LoadingIcon, "Creating Feature Page View Model: "+className+"PageViewModel()")

	content := fmt.Sprintf(
		templates.FeaturePageViewModel,
		packageName,
		className,
	)

	err := utils.CreateAndInsertIfFileNotExist(filePath, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(constants.LoadingIcon, "Successfully created: "+className+"PageViewModel()")
	fmt.Println()
}

func writeInBeamLocation(filePath string, featureInfo models.FeatureInfoModel, contentStr string) {

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
	err := os.WriteFile(filePath, []byte(updatedContent), 0644)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}
}
