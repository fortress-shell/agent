package steps

import (
	"fmt"
	"github.com/fortress-shell/agent/worker"
	"os"
)

type OverrideCheckoutStep struct {
	Commit      string
	Branch      string
	Environment map[string]string
}

func (s *OverrideCheckoutStep) Run(app *worker.Worker) error {
	config := app.Config
	session, err := app.SSHClient.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()
	session.Stdout = os.Stdout
	session.Stderr = os.Stdout
	gitCloneRepo := fmt.Sprintf(
		"git clone ssh://git@github.com/%s/%s --branch %s --single-branch ",
		config.Username,
		config.Repo,
		config.Branch)
	fmt.Println(gitCloneRepo)
	err = session.Run(gitCloneRepo)
	if err != nil {
		return err
	}
	gitCheckout := fmt.Sprintf(
		"cd %s git checkout %s",
		config.Repo,
		config.Commit,
	)
	fmt.Println(gitCheckout)
	err = session.Run(gitCheckout)
	if err != nil {
		return err
	}
	err = session.Run("mv /home/ubuntu/id_rsa /home/ubuntu/.ssh/id_rsa && chmod 400 /home/ubuntu/.ssh/id_rsa")
	if err != nil {
		return err
	}
	return nil
}
