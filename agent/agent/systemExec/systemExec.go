package systemExec

// Db struct with db map
type Config struct {
	Job              map[int]architecture.Job
	OverallExitLevel int
}
