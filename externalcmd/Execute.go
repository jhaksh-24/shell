package externalcmd

import (
	"fmt"
	"os"
	"os/exec"


	"github.com/jhaksh-24/shell/utils"
)

func Execute(cmd string, args ...string) {
	if file,exists := utils.SearchBinInPath(cmd); exists {		
		command := exec.Command(file, args...)

		output, err := command.CombinedOutput()
		fmt.Println(string(output))

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
		}
	} else {
		err := utils.InvalidCmd(cmd)
		fmt.Println(err)
	}
}


