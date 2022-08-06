package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func ExecuteCommand(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)
	stdout, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(stdout)), nil
}

func Input(title string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(title)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(fmt.Sprint(text))
}

func main() {
	allContexts, err := ExecuteCommand("kubectl config get-contexts --output=name")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	activeContext, err := ExecuteCommand("kubectl config current-context")

	contexts := strings.Split(allContexts, "\n")
	red := color.New(color.FgRed).SprintFunc()

	mapContext := make(map[string]string)

	for index, context := range contexts {
		mapContext[strconv.Itoa(index+1)] = context
		if context == activeContext {
			fmt.Printf("[%v] %v\n", red("*"), context)
			continue
		}
		fmt.Printf("[%v] %v\n", index+1, context)
	}

	selectContext := Input("Select context : ")

	if _, ok := mapContext[selectContext]; ok {
		fmt.Printf("Selected context : %v\n", mapContext[selectContext])
		ExecuteCommand(fmt.Sprintf("kubectl config use-context %v", mapContext[selectContext]))
		fmt.Printf("Switch to context : %v\n", mapContext[selectContext])
	} else {
		fmt.Printf("Oops, %v is not there\n", selectContext)
	}

}
