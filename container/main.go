package main

import (
	"os"
	"os/exec"
)

func main() {
	switch os.Args[1] {
	case "run":
		parent()
	case "child":
		child()
	default:
		panic("???")
	}
}

func parent() {
	// does not work on MacOs
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func child() {
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
