package main

import (
	"flag"
	"github.com/fortress-shell/agent/parser"
	"github.com/fortress-shell/agent/worker"
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

func main() {
	var path string = os.Getenv("NOMAD_META_CONFIG_PATH")
	var libvirtUrl string = os.Getenv("NOMAD_META_LIBVIRT_URL")
	var kafkaUrl string = os.Getenv("NOMAD_META_KAFKA_URL")
	var topic string = os.Getenv("NOMAD_DC")
	var id string = os.Getenv("NOMAD_JOB_NAME")
	var vmPath string = os.Getenv("NOMAD_META_VM_PATH")
	var memory uint64 = 524288
	var buildId int
	var repo string = os.Getenv("NOMAD_META_REPO")
	var branch string = os.Getenv("NOMAD_META_BRANCH")
	var commit string = os.Getenv("NOMAD_META_COMMIT")
	var username string = os.Getenv("NOMAD_META_USERNAME")
	var sshKey string = os.Getenv("NOMAD_META_SSH_KEY")

	flag.StringVar(&path, "config", path, "a path to config")
	flag.IntVar(&buildId, "build-id", 0, "a build id from rails")
	flag.Parse()

	payload, err := parser.NewPayloadFromFilePath(path)
	if err != nil {
		log.Println(err)
		os.Exit(FORTRESS_ERROR_STATUS)
	}

	tasks := parser.GenerateSteps(payload)

	config := &worker.WorkerConfig{
		Memory:     memory,
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

	for _, task := range tasks {
		err := task.Run(app)
		if err != nil {
			log.Println(err)
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
finish:
	app.SSHClient.Close()
	app.LibVirt.Close()
	app.Logger.Writer.Close()
	os.Exit(app.ExitCode)
}
