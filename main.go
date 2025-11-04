package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"

	"github.com/jhaksh-24/shell/builtincmd"
	"github.com/jhaksh-24/shell/externalcmd"
	"github.com/jhaksh-24/shell/history"
)

func main() {
	history.Init()
	for {
		currdir, err := os.Getwd()
		if err != nil {
			fmt.Println("Error getting current directory:", err)
			return
		}
		fmt.Fprint(os.Stdout, currdir,"\n$ ")
		cmdLine, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		
		cmdLine = strings.TrimSpace(cmdLine)

		history.Add(cmdLine)

		parts := strings.Fields(cmdLine)
		if len(parts) == 0 {
			continue
		}

		cmd := parts[0]

		if cmd == "exit" {

			code := 0
			if len(parts) > 1{
				code, err = strconv.Atoi(parts[1])
				if err != nil {
					fmt.Println("Invalid exit code:", parts[1])
				}
			}

			os.Exit(code)

		} else if cmd == "echo" {

			fmt.Println(builtincmd.Echo(parts))

		} else if cmd == "type" {

			builtincmd.Type(parts)

		} else if cmd == "cd" {

			builtincmd.ChangeDirectory(parts[1:]...)

		} else if cmd == "history" {

			for i, h := range history.Get() {
				fmt.Printf("%d  %s\n", i+1, h)
			}

		} else {

			externalcmd.Execute(cmd, parts[1:]...)

		}
	}
}






