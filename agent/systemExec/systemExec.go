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

var runtests = []struct {
	in  string
	out Job
}{
	{"EXEC_RES_OK_EXIT_OK", Job{"OK - ran ok", "", 0}},
	{"EXEC_RES_WARN_EXIT_OK", Job{"Warn - ran with warning", "", 1}},
	{"EXEC_RES_CRITICAL_EXIT_OK", Job{"Critical - error running test", "", 2}},
	{"EXEC_RES_OK_UNKNOWN", Job{"Unknown - no exit code returned", "", 3}},
}

// Run takes command to run and returns result object
func (j Job) Run(cmd string, ID int) architecture.CommandResult {
	fmt.Printf("HERE inside the Run task with [%s] [%d]...\n", cmd, ID)
	for i, j := range runtests {
		if j.in == cmd {
			fmt.Printf("id[%d], out[%s] err[%s] exit[%d]\n", i, j.out.standardOut, j.out.standardError, j.out.exitCode)
			break
		}
	}
	return architecture.CommandResult{}
}
