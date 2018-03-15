package worker

import (
	"bytes"
	"os"
	"os/signal"
	"syscall"
	"time"
	"golang.org/x/crypto/ssh"
	"github.com/fortress-shell/agent/domain"
	"github.com/fortress-shell/agent/kafka"
	"github.com/fortress-shell/agent/keys"
	libvirt "github.com/libvirt/libvirt-go"
)

type LibVirt = libvirt.Connect

type Worker struct {
	// Libvirt connection
	*LibVirt
	// Kafka producer instance
	Logger *kafka.KafkaWriter
	// SSH connection
	SSHClient *ssh.Client
	// Config created with
	Config *WorkerConfig
	// Stop channel receive message with os.Signal when job will be stopped
	Stop <-chan os.Signal
	// Timeout channel receive message with current
	// time when job will be timeouted
	Timeout <-chan time.Time
	// exit status code
	ExitCode int
}

const (
	TIMEOUT              = 10
	AUTHORIZED_KEYS_PATH = "/home/ubuntu/.ssh/authorized_keys"
	IDENTITY_PATH        = "/home/ubuntu/id_rsa"
	DELAY                = 20
)

func NewWorker(config *WorkerConfig) (*Worker, error) {
	conn, err := libvirt.NewConnect(config.LibVirtUrl)
	if err != nil {
		return nil, err
	}

	logger, err := kafka.NewKafkaWriter(
		config.KafkaUrl,
		config.Topic,
		config.Id,
		config.BuildId,
	)
	if err != nil {
		return nil, err
	}

	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	xml, err := domain.NewDomainXml(domain.Config{
		ImagePath: config.VmPath,
		Memory:    config.Memory,
		Name:      config.Id,
	})
	if err != nil {
		return nil, err
	}
	dom, err := conn.DomainCreateXML(*xml, libvirt.DOMAIN_START_AUTODESTROY)
	if err != nil {
		return nil, err
	}

	<-time.After(DELAY * time.Second)

	adom := domain.AgentDomain{dom}
	interfaces, err := adom.GetNetworkInterfaces()
	if err != nil {
		return nil, err
	}
	sshClientConfig, publicKey, err := keys.NewKeyPair()
	if err != nil {
		return nil, err
	}
	err = adom.HighLevelWriteFile(AUTHORIZED_KEYS_PATH, publicKey)
	if err != nil {
		return nil, err
	}
	err = adom.HighLevelWriteFile(IDENTITY_PATH, []byte(config.Identity))
	if err != nil {
		return nil, err
	}
	var connUrl bytes.Buffer
	connUrl.WriteString(interfaces.Return[1].IpAddresses[0].IpAddress)
	connUrl.WriteString(":22")
	ipWithPort := connUrl.String()
	connection, err := ssh.Dial("tcp", ipWithPort, sshClientConfig)
	if err != nil {
		return nil, err
	}

	return &Worker{
		Config:    config,
		Stop:      stop,
		Timeout:   time.After(TIMEOUT * time.Minute),
		Logger:    logger,
		LibVirt:   conn,
		SSHClient: connection,
	}, nil
}
