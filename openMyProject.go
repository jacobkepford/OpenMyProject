package main

import (
	"embed"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/manifoldco/promptui"
)

//go:embed ascii_art.txt
var f embed.FS

func main() {
	art, err := f.ReadFile("ascii_art.txt")

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

	templates := &promptui.SelectTemplates{
		Label:    "{{ . | green }} ",
		Active:   "{{ . | bold }} ",
		Inactive: "{{. | green }}",
	}

	prompt := promptui.Select{
		Label:     "Use the arrow keys to select a project",
		Items:     folderSlice,
		Templates: templates,
	}

	prompt.HideHelp = true
	prompt.HideSelected = true
	prompt.Size = 10

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
