package main

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/manifoldco/promptui"
)

type config struct {
	Name string
	User user
}

type configs struct {
	Config []config
}

type user struct {
	Name  string `toml:"name"`
	Email string `toml:"email"`
}

const FILE_PATH = "./test.toml"

func pos(value string, slice []string) int {
	for p, v := range slice {
		if v == value {
			return p
		}
	}
	return -1
}

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

	home := os.Getenv("HOME")

	configs := &configs{}
	if _, err := toml.DecodeFile(fmt.Sprintf("%s/.config/git-config-switch.toml", home), configs); err != nil {
		fmt.Println(err)
		return
	}

	var items []string
	for _, c := range configs.Config {
		items = append(items, c.Name)
	}

	prompt := promptui.Select{
		Label: "Select config",
		Items: items,
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

	p := pos(result, items)
	if p == -1 {
		fmt.Println("Not found config")
		return
	}

	selectItem := configs.Config[p]

	rv := reflect.ValueOf(selectItem)

	rt1 := rv.Type()
	for i := 0; i < rt1.NumField(); i++ {
		f1 := rt1.Field(i)
		k1 := f1.Type.Kind()
		v1 := rv.FieldByName(f1.Name)

		if k1 == reflect.Struct {
			rt2 := v1.Type()
			for j := 0; j < rt2.NumField(); j++ {
				f2 := rt2.Field(j)
				v2 := v1.FieldByName(f2.Name)

				cmd(
					fmt.Sprintf("git config %s.%s \"%s\"",
						strings.ToLower(f1.Name),
						strings.ToLower(f2.Name),
						v2),
				)
			}
		}
	}
}
