package builtincmd

import (
	"fmt"
	"os"
)

func ChangeDirectory(path ...string) error {
	if len(path) > 1 {
		return fmt.Errorf("Too many args for cd command\n")
	}

	var target string
	var err error

	if len(path) == 0 || path[0] == "~" || path[0] == "" {
		if target, err = os.UserHomeDir(); err != nil {
			return fmt.Errorf("Error getting home directory: %v\n", err)
		} 
		err = os.Chdir(target)
		if err != nil {
			return fmt.Errorf("Error changing directory: %v\n", err)
		}
	} else {
		target = path[0]
		err = os.Chdir(target)
		if err != nil {
			return fmt.Errorf("Error changing directory: %v\n", err)
		}
	}
	return nil
}
