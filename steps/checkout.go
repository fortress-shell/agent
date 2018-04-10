package steps

import (
	"fmt"
	"github.com/fortress-shell/agent/kafka"
	"github.com/fortress-shell/agent/worker"
	"os"
)

type OverrideCheckoutStep struct {
	Environment map[string]string
}

const script = `
	set -e pipeline;
	sudo chown ubuntu:ubuntu /home/ubuntu/.ssh/id_rsa;
	sudo chmod 600 /home/ubuntu/.ssh/id_rsa;
	echo -e "Host github.com\n\tStrictHostKeyChecking no\n" > ~/.ssh/config;
	git clone %s --branch %s --single-branch;
	cd $(basename %s | cut -f 1 -d '.');
	git checkout %s;
`

func (s *OverrideCheckoutStep) Run(app *worker.Worker) error {
	config := app.Config
	logger := &kafka.KafkaStageWriter{app.Logger}
	logger.Write([]byte("Setting up repository...\n"))
	session, err := app.SSHClient.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()
	session.Stdout = os.Stdout
	session.Stderr = os.Stdout
	checkout := fmt.Sprintf(script,
		config.RepositoryUrl,
		config.Branch,
		config.RepositoryUrl,
		config.Commit,
	)
	fmt.Println(checkout)
	err = session.Run(checkout)
	if err != nil {
		return err
	}
	return nil
}
