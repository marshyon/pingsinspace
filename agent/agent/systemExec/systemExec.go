package systemexec

import (
	"fmt"

	architecture "github.com/marshyon/pingsinspace/agent"
)

// Job is a struct holding the command to run and its result
type Job struct {
	standardOut   string
	standardError string
	exitCode      int
}

// Run takes command to run and returns result object
func (j Job) Run(cmd string, ID int) architecture.CommandResult {
	fmt.Println("HERE inside the Run task ...")
	return architecture.CommandResult{}
}
