package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/manifoldco/promptui"
)

func main() {
	art, err := os.ReadFile("ascii_art.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(art))

	projectPath := "C:/Projects"

	folders, err := os.ReadDir(projectPath)

	if err != nil {
		log.Fatal(err)
	}

	folderSlice := make([]string, len(folders))

	for i, folder := range folders {
		folderSlice[i] = folder.Name()
	}

	prompt := promptui.Select{
		Label: "Select a project",
		Items: folderSlice,
	}

	_, result, err := prompt.Run()

	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("code", result)
	cmd.Dir = projectPath
	err = cmd.Run()

	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

}
