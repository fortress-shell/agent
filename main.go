package main

import (
	"github.com/fortress-shell/agent/parser"
	"github.com/fortress-shell/agent/worker"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
)

const (
	SUCCESS_STATUS        = iota
	FAILURE_STATUS
	TIMEOUT_STATUS
	STOP_STATUS
	FORTRESS_ERROR_STATUS
)

func main() {
	config := worker.DefaultConfig()

	payload, err := parser.NewPayloadFromFilePath(config.PayloadPath)
	if err != nil {
		log.Println(err)
		os.Exit(FORTRESS_ERROR_STATUS)
	}

	tasks := parser.GenerateSteps(payload)

	app, err := worker.NewWorker(config)
	if err != nil {
		log.Println(err)
		os.Exit(FORTRESS_ERROR_STATUS)
	}
	// TODO: refactoring using multisteps
	for _, task := range tasks.Checkout {
		err := task.Run(app)
		if err != nil {
			app.ExitCode = FORTRESS_ERROR_STATUS
			goto finish
		}
		select {
		case <-app.Stop:
			app.ExitCode = STOP_STATUS
			goto finish
		case <-app.Timeout:
			app.ExitCode = TIMEOUT_STATUS
			goto finish
		default:
		}
	}

	for _, task := range tasks.OverrideBuild {
		err := task.Run(app)
		if err != nil {
			switch t := err.(type) {
			default:
				log.Println(t)
				app.ExitCode = FORTRESS_ERROR_STATUS
				goto finish
			case *ssh.ExitError:
				log.Println("Exit Status:", t.Waitmsg.ExitStatus())
				app.ExitCode = FAILURE_STATUS
				goto finish
			}
		}
		select {
		case <-app.Stop:
			app.ExitCode = STOP_STATUS
			goto finish
		case <-app.Timeout:
			app.ExitCode = TIMEOUT_STATUS
			goto finish
		default:
		}
	}
finish:
	app.SSHClient.Close()
	app.LibVirt.Close()
	app.Logger.Writer.Close()
	os.Exit(app.ExitCode)
}
