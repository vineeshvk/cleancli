package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/vineeshvk/cleancli/constants"
	"github.com/vineeshvk/cleancli/dirvalid"
	"github.com/vineeshvk/cleancli/input"
	"github.com/vineeshvk/cleancli/models"
	"github.com/vineeshvk/cleancli/utils"
	"github.com/vineeshvk/cleancli/write"
)

func main() {
	// Check if a subcommand is provided
	argsLength := len(os.Args)
	if argsLength < 2 {
		fmt.Println("Usage: cleancli <subcommand> [options]")
		fmt.Println("Available subcommands: api, feature")
		os.Exit(1)
	}

	// Get the subcommand
	subcommand := os.Args[1]

	handleSubcommand(subcommand)
}

func handleSubcommand(subcommand string) {
	switch subcommand {
	case "api":
		generateAPI()
	case "feature":
		generateFeature()
	default:
		fmt.Printf("Unknown subcommand: %s\n", subcommand)
		fmt.Println("Available subcommands: api, feature")
		os.Exit(1)
	}
}

func generateAPI() {
	var mainDirectoryModel models.MainDirectoryModel = dirvalid.ValidateRootDirectories()

	var apiInfo models.ApiInfoModel = input.GetAPIInfos(mainDirectoryModel)

	write.WriteApiService(mainDirectoryModel.GetApiServiceRoute(), apiInfo)
	write.WriteDataSource(mainDirectoryModel.DataDir, apiInfo)
	write.WriteRepo(mainDirectoryModel, apiInfo)
	write.WriteUseCase(mainDirectoryModel, apiInfo)
	write.WriteDI(mainDirectoryModel, apiInfo)

	fmt.Println(constants.CompletedIcon, "Completed adding API.")

	utils.ExecuteBuildRunner(mainDirectoryModel.DataDir)
}

func generateFeature() {
	// Define a FlagSet for the 'feature' subcommand
	featureCmd := flag.NewFlagSet("feature", flag.ExitOnError)

	// Add an optional flag --ignore-routes
	ignoreRoutes := featureCmd.Bool("ignore-routes", false, "Ignore generating route files")

	// Parse the flags
	if err := featureCmd.Parse(os.Args[2:]); err != nil {
		fmt.Printf("Error parsing flags for 'feature': %v\n", err)
		os.Exit(1)
	}

	if *ignoreRoutes {
		// If the flag is set, skip generating route files
		fmt.Println("Skipping route file generation.")
	} else {
		// If the flag is not set, proceed with generating route files
		fmt.Println("Generating route files...")

	}
	var featureInfo models.FeatureInfoModel = input.GetFeatureInfos()
	fmt.Println(constants.LoadingIcon, " Creating feature "+featureInfo.FeatureName)
	// utils.CreateNewFile("./lib/di/" + featureInfo.FeatureName + "_module.dart")
	write.WriteFeatureDI("./lib/di/", featureInfo)
	write.WriteFeatureRoute(featureInfo)
}
