package cmd

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"os/exec"
)

func Version(c *cli.Context) {
	out, err := exec.Command("docker", "version").Output()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(string(out))
}

func Ps(c *cli.Context) {
	out, err := exec.Command("docker", "ps").Output()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(string(out))
}
