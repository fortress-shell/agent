package steps

import (
	"fmt"
	"github.com/fortress-shell/agent/kafka"
	"github.com/fortress-shell/agent/worker"
)

type OverrideBuildStep struct {
	Command     string
	Environment map[string]string
}

const (
	BUILD_STAGE = "BUILD_STAGE"
)

func (s *OverrideBuildStep) Run(app *worker.Worker) error {
	logger := &kafka.KafkaStageWriter{app.Logger}
	session, err := app.SSHClient.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()
	session.Stdout = logger
	session.Stderr = logger
	for k, v := range s.Environment {
		if err = session.Setenv(k, v); err != nil {
			return err
		}
	}
	command := fmt.Sprintf("cd $(basename %s | cut -f 1 -d '.'); %s",
		app.Config.RepositoryUrl,
		s.Command,
	)
	err = session.Run(command)
	if err != nil {
		return err
	}
	return nil
}
