package steps

import (
	"bytes"
	"fmt"
	"github.com/fortress-shell/agent/kafka"
	"github.com/fortress-shell/agent/worker"
	"io"
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
	stdout, err := session.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := session.StderrPipe()
	if err != nil {
		return err
	}
	go io.Copy(logger, stdout)
	go io.Copy(logger, stderr)
	for k, v := range s.Environment {
		command.WriteString(fmt.Sprintf("export %s=%s;", k, v))
	}
	command.WriteString(fmt.Sprintf("cd $(basename %s | cut -f 1 -d '.'); ",
		app.Config.RepositoryUrl,
	))
	command.WriteString(s.Command)
	command.WriteString("\n")
	err = session.Run(command.String())
	if err != nil {
		return err
	}
	return nil
}
