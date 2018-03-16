package worker

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
