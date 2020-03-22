package architecture

// CommandResult represents outputs of a system command
type CommandResult struct {
	standardOut   string
	standardError string
	exitCode      int
}

// Accessor interface is used to access and abstract the back-end
// ID is used to uniquely identify each command that is configured
// and is expected to be incremented from 0 so that this can be
// used by the Run methond to return each result set as a slice of
// CommandResult(s)
type Accessor interface {
	Run(string, int) CommandResult
}

// JobService uses accessor interface
type JobService struct {
	a Accessor
}

// Run method used to run jobs
func (vs JobService) Run(cmd string, ID int) CommandResult {
	r := vs.a.Run(cmd, ID)
	return r
}

// NewJobService creates a new service to action
// run and configure operations
func NewJobService(a Accessor) JobService {
	return JobService{
		a: a,
	}
}
