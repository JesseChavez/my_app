package jobs

import (
	"github.com/JesseChavez/enki"
)

// create user
type AddUserJob struct {
	Queue    string
	Priority int
}

func (job *AddUserJob) Init() *AddUserJob {
	job.Priority = 3
	job.Queue = "default"

	return job
}


func (job *AddUserJob) Perform(args enki.Args) []error {
	var failure []error

	return failure
}
