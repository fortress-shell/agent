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
	FAILURE_STATUS        = iota
	TIMEOUT_STATUS        = iota
	STOP_STATUS           = iota
	FORTRESS_ERROR_STATUS = iota
)

var path string = os.Getenv("NOMAD_META_CONFIG_PATH")
var libvirtUrl string = os.Getenv("NOMAD_META_LIBVIRT_URL")
var kafkaUrl string = os.Getenv("NOMAD_META_KAFKA_URL")
var topic string = os.Getenv("NOMAD_DC")
var id string = os.Getenv("NOMAD_JOB_NAME")
var vmPath string = os.Getenv("NOMAD_META_VM_PATH")
var buildId string = os.Getenv("NOMAD_META_BUILD_ID")
var repo string = os.Getenv("NOMAD_META_REPO")
var branch string = os.Getenv("NOMAD_META_BRANCH")
var commit string = os.Getenv("NOMAD_META_COMMIT")
var username string = os.Getenv("NOMAD_META_USERNAME")
var sshKey string = os.Getenv("NOMAD_META_SSH_KEY")

func main() {
	payload, err := parser.NewPayloadFromFilePath(path)
	if err != nil {
		log.Println(err)
		os.Exit(FORTRESS_ERROR_STATUS)
	}

	tasks := parser.GenerateSteps(payload)
	config := &worker.WorkerConfig{
		LibVirtUrl: libvirtUrl,
		KafkaUrl:   kafkaUrl,
		VmPath:     vmPath,
		Id:         id,
		Commit:     commit,
		Branch:     branch,
		Repo:       repo,
		Username:   username,
		Topic:      topic,
		BuildId:    buildId,
		Identity:   sshKey,
	}
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
