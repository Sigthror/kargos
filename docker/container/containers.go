package container

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func Generate() {
	cmd := exec.Command(`docker`, `ps`, `--format`, `table {{.Names}}\t{{.Ports}}\t{{.Image}}`)
	out, err := cmd.Output()
	if err != nil {
		log.Printf(`failed to execute: %s"`, err.Error())
	}
	fmt.Println(`Containers | size=13`)
	lines := strings.Split(string(out), "\n")
	for i := 1; i < (len(lines) - 1); i++ {
		fmt.Printf(`-- %s | font=monospace onclick=bash bash='docker stop %s'`, lines[i], getContainerName(lines[i]))
	}
}
