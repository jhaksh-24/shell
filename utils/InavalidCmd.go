package utils

import "fmt"

func InvalidCmd(cmd string) error {
	return fmt.Errorf("%s: command not found",cmd)
}
