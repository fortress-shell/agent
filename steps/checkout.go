package steps

import (
	"fmt"
	"os"
	"github.com/fortress-shell/agent/worker"
)

type OverrideCheckoutStep struct {
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
	err = session.Run("sudo chmod 400 /home/ubuntu/.ssh/id_rsa && echo -e \"Host github.com\n\tStrictHostKeyChecking no\n\" >> ~/.ssh/config")
	if err != nil {
		return err
	}
	session, err = app.SSHClient.NewSession()
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
	session, err = app.SSHClient.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()
	session.Stdout = os.Stdout
	session.Stderr = os.Stdout
	gitCheckout := fmt.Sprintf(
		"cd %s; git checkout %s",
		config.Repo,
		config.Commit,
	)
	fmt.Println(gitCheckout)
	err = session.Run(gitCheckout)
	if err != nil {
		return err
	}
	return nil
}
