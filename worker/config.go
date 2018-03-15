package worker

type WorkerConfig struct {
	LibVirtUrl string
	KafkaUrl   string
	VmPath     string
	Memory     uint64
	Repo       string
	Branch     string
	Commit     string
	Topic      string
	Username   string
	BuildId    int
	Id         string
	Identity   string
}
