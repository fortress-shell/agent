package steps

import (
	"fmt"
	"github.com/fortress-shell/agent/worker"
	"os"
)

type OverrideCheckoutStep struct {
	Environment map[string]string
}

const script = `
    sudo chmod 400 /home/ubuntu/.ssh/id_rsa;
    echo -e "Host github.com\n\tStrictHostKeyChecking no\n" >> ~/.ssh/config;
    git clone ssh://%s --branch %s --single-branch;
    cd $(basename %s);
    git checkout %s;
`

func (s *OverrideCheckoutStep) Run(app *worker.Worker) error {
	session, err := app.SSHClient.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()
	session.Stdout = os.Stdout
	session.Stderr = os.Stdout
	setup := fmt.Sprintf(script,
		app.Config.RepositoryUrl,
		app.Config.Branch,
		app.Config.RepositoryUrl,
		app.Config.Branch,
	)
	err = session.Run(setup)
	if err != nil {
		return err
	}
	return nil
}
