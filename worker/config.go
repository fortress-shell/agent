package worker

import "os"

type WorkerConfig struct {
	LibVirtUrl string
	KafkaUrl   string
	VmPath     string
	Repo       string
	Branch     string
	Commit     string
	Topic      string
	Username   string
	BuildId    string
	Id         string
	Identity   string
}

func DefaultConfig() *WorkerConfig {
	return &WorkerConfig{
		LibVirtUrl: os.Getenv("LIBVIRT_URL"),
		KafkaUrl:   os.Getenv("KAFKA_URL"),
		VmPath:     os.Getenv("VM_PATH"),
		Id:         os.Getenv("NOMAD_JOB_NAME"),
		Commit:     os.Getenv("NOMAD_META_COMMIT"),
		Branch:     os.Getenv("NOMAD_META_BRANCH"),
		Repo:       os.Getenv("NOMAD_META_REPO"),
		Username:   os.Getenv("NOMAD_META_USERNAME"),
		Topic:      os.Getenv("NOMAD_DC"),
		BuildId:    os.Getenv("NOMAD_META_BUILD_ID"),
		Identity:   os.Getenv("NOMAD_META_SSH_KEY"),
	}
}
