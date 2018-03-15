package steps

import "github.com/fortress-shell/agent/worker"

type Step interface {
	Run(app *worker.Worker) error
}
