package worker

import (
	"github.com/caarlos0/env"
)

type WorkerConfig struct {
	LibVirtUrl string `env:"LIBVIRT_URL"`
	KafkaUrl   string `env:"KAFKA_URL"`
	VmPath     string `env:"VM_PATH"`
	Repository string `env:"NOMAD_META_REPOSITORY"`
	Branch     string `env:"NOMAD_META_BRANCH"`
	Commit     string `env:"NOMAD_META_COMMIT"`
	Topic      string `env:"NOMAD_DC"`
	BuildId    int `env:"NOMAD_META_BUILD_ID"`
	UserId     int `env:"NOMAD_META_USER_ID"`
	Id         string `env:"NOMAD_JOB_NAME"`
	Identity   string `env:"NOMAD_META_SSH_KEY"`
}

func DefaultConfig() *WorkerConfig {
	cfg := WorkerConfig{}
  env.Parse(&cfg)
	return &cfg
}
