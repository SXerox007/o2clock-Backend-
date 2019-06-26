package utils

import "time"

const JOB_MEM_INTERVAL_PERIOD time.Duration = 1 * time.Hour

type CheckMemUtilizerJob struct {
	t *time.Timer
}

// new mem usage
func NewJobMemUsage() CheckMemUtilizerJob {
	return CheckMemUtilizerJob{time.NewTimer(getNextCheckMem())}
}

// check next check mem
func getNextCheckMem() time.Duration {
	next := time.Now().Local().Add(1 * time.Minute)
	return next.Sub(time.Now())
}

func (jt CheckMemUtilizerJob) updateJobMemCheck() {
	jt.t.Reset(getNextCheckMem())
}
