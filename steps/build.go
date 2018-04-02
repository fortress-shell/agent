package steps

import (
	"bytes"
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
	var command bytes.Buffer
	logger := &kafka.KafkaStageWriter{app.Logger}
	session, err := app.SSHClient.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()
	for k, v := range s.Environment {
		command.WriteString(fmt.Sprintf("export %s=%s;", k, v))
	}
	command.WriteString(fmt.Sprintf("cd %s;", app.Config.Repo))
	command.WriteString(s.Command)
	logger.Write([]byte(s.Command))
	session.Stdout = logger
	session.Stderr = logger
	fmt.Println(command.String())
	err = session.Run(command.String())
	if err != nil {
		return err
	}
	return nil
}
