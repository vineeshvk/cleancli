package utils

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/vineeshvk/cleancli/constants"
)

func ExecuteBuildRunner(dataDir string) {
	fmt.Println(constants.BuildRunnerIcon, " Running build_runner in "+dataDir)
	err := os.Chdir(dataDir)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	cmd := exec.Command("dart", "run", "build_runner", "build")

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()

	if err != nil {
		fmt.Println("Errrr" + err.Error())
		return
	}

}
