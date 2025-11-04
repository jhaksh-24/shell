package history

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var history []string
var historyFile string

func Init() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home dir:", err)
		return
	}

	historyFile = home + "/.aksh_history"

	if _, err := os.Stat(historyFile); os.IsNotExist(err) {
		file, err := os.Create(historyFile)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		file.Close()
	}

	Load()
}

func Load() {
	data, err := os.ReadFile(historyFile)
	if err != nil {
		fmt.Println("Error loading history:", err)
		return
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	history = []string{}
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			history = append(history, line)
		}
	}
}

func Add(cmd string) {
	cmd = strings.TrimSpace(cmd)
	if cmd == "" {
		return
	}
	history = append(history, cmd)

	file, err := os.OpenFile(historyFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening history file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	fmt.Fprintln(writer, cmd)
	writer.Flush()
}

func Get() []string {
	return history
}

