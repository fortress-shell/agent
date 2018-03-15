package steps

import (
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
	logger := kafka.KafkaStageWriter{app.Logger, BUILD_STAGE, s.Command}
	session, err := app.SSHClient.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()
	for k, v := range s.Environment {
		session.Setenv(k, v)
	}
	session.Stdout = &logger
	session.Stderr = &logger
	err = session.Run(s.Command)
	if err != nil {
		return err
	}
	return nil
}
