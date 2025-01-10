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

func WriteFeatureDI(diDir string, featureInfo models.FeatureInfoModel) {

	fmt.Println(constants.LoadingIcon, " Working on api_service file...")

	packageName, _ := utils.GetPackageName()
	featureNameCamelCase := utils.SnakeCaseToCamelCase(featureInfo.FeatureName)

	// 1: Directory or package name, 2: Group Name, 3: Provider Name, 4: Group Class Name
	moduleString := fmt.Sprintf(
		templates.FeatureDI,
		packageName,
		featureInfo.FeatureName,
		featureNameCamelCase,
		utils.SnakeCaseToPascalCase(featureInfo.FeatureName),
	)

	err := utils.CreateAndInsertIfFileNotExist("./lib/di/"+featureInfo.FeatureName+"_module.dart", moduleString)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println()
}

func WriteFeatureRoute(featureInfo models.FeatureInfoModel) {
	// Define the file path
	filePath := "./lib/main/locations.dart"

	// Define the feature location to add
	newLocation := featureInfo.FeatureName + "Location()"

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

	// Prepare the updated content
	updatedContent := contentStr[:endIndex] +
		newLocation + "," +
		contentStr[endIndex:]

	// Write the updated content back to the file
	err = os.WriteFile(filePath, []byte(updatedContent), 0644)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}

	fmt.Println("New location added successfully!")
}
