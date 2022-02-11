package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/manifoldco/promptui"
)

func cmd(c string) {
	cmd := exec.Command("/bin/sh", "-c", c)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Got error: %s\n", err.Error())
		fmt.Println()
		return
	}
	fmt.Println("Run:", cmd)
}

func main() {
	config, err := os.UserConfigDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
	dir := filepath.Join(config, "git-config-switcher")
	files, err := ioutil.ReadDir(dir)

	var items []string
	for _, f := range files {
		items = append(items, f.Name())
	}

	prompt := promptui.Select{
		Label: "Select config",
		Items: items,
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	cmd(fmt.Sprintf("git config include.path \"%s\"", filepath.Join(dir, result)))
}
