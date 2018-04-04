package worker

import (
	"github.com/caarlos0/env"
	"log"
	"os"
)

type WorkerConfig struct {
	LibVirtUrl    string `env:"LIBVIRT_URL,required"`
	KafkaUrl      string `env:"KAFKA_URL,required"`
	VmPath        string `env:"VM_PATH,required"`
	RepositoryUrl string `env:"NOMAD_META_REPOSITORY_URL,required"`
	Branch        string `env:"NOMAD_META_BRANCH,required"`
	Commit        string `env:"NOMAD_META_COMMIT,required"`
	Topic         string `env:"NOMAD_DC,required"`
	BuildId       int    `env:"NOMAD_META_BUILD_ID,required"`
	UserId        int    `env:"NOMAD_META_USER_ID,required"`
	Id            string `env:"JOB_ID,required"`
	Identity      string `env:"NOMAD_META_SSH_KEY,required"`
	PayloadPath   string `env:"PAYLOAD_PATH,required"`
	DiskPath      string `env:"DISK_PATH,required"`
}

func DefaultConfig() *WorkerConfig {
	cfg := WorkerConfig{}
	err := env.Parse(&cfg)
	if err != nil {
		log.Println("%s", err)
		os.Exit(4)
	}
	log.Println(cfg.Identity)
	return &cfg
}
