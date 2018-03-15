# Agent


// func GetDomainXml() string {
//  var result bytes.Buffer
//  t := template.Must(template.New("domain").Parse(domain.DOMAIN_TEMPLATE))
//  if err != nil {
//    panic(err)
//  }
//  options := domain.Domain{
//    Name:      "fuck",
//    Vcpu:      1,
//    ImagePath: os.Getenv("VM_PATH"),
//    Memory:    500000,
//  }
//  err = t.Execute(&result, options)
//  if err != nil {
//    panic(err)
//  }
//  return result.String()
// }

// func main() {
// key, err := rsa.GenerateKey(rand.Reader, 2048)
// signer, err := ssh.NewSignerFromKey(key)
// pub, err := ssh.NewPublicKey(&key.PublicKey)
// if err != nil {
//  panic(err)
// }
// public := ssh.MarshalAuthorizedKey(pub)
// if err != nil {
//  log.Fatalf("unable to parse private key: %v", err)
// }
// encoded := base64.StdEncoding.EncodeToString(public)
// log.Println(encoded)

// xml := GetDomainXml()

// conn, err := libvirt.NewConnect(os.Getenv("LIBVIRT_URL"))
// if err != nil {
//  panic(err)
// }

// dom, err := conn.DomainCreateXML(xml, libvirt.DOMAIN_START_AUTODESTROY)
// if err != nil {
//  panic(err)
// }
// command := &domain.QemuAgentCommandRequest{
//  Execute: "guest-network-get-interfaces",
// }
// jsonCommand, _ := json.Marshal(command)

// <-time.After(20 * time.Second)

// res, err := dom.QemuAgentCommand(string(jsonCommand), libvirt.DOMAIN_QEMU_AGENT_COMMAND_MIN, 0)
// var keys domain.QemuAgentCommandResponse
// json.Unmarshal([]byte(res), &keys)
// log.Println(res)

// jsonCommand, _ = json.Marshal(domain.GuestFileOpen{
//  Execute: "guest-file-open",
//  Arguments: domain.QemuGuestFileOpenArguments{
//    Path: "/home/ubuntu/.ssh/authorized_keys",
//    Mode: "w",
//  },
// })

// res, err = dom.QemuAgentCommand(string(jsonCommand), libvirt.DOMAIN_QEMU_AGENT_COMMAND_MIN, 0)
// if err != nil {
//  panic(err)
// }
// log.Println(res)
// var openResponse domain.GuestFileOpenReply
// json.Unmarshal([]byte(res), &openResponse)

// jsonCommand, _ = json.Marshal(domain.GuestFileWrite{
//  Execute: "guest-file-write",
//  Arguments: domain.QemuGuestFileWriteArguments{
//    Handle: openResponse.Return,
//    Buf64:  encoded,
//  },
// })

// resp, err := dom.QemuAgentCommand(string(jsonCommand), libvirt.DOMAIN_QEMU_AGENT_COMMAND_MIN, 0)
// if err != nil {
//  panic(err)
// }
// log.Println(resp)
// jsonCommand, _ = json.Marshal(domain.GuestFileClose{
//  Execute: "guest-file-close",
//  Arguments: domain.QemuGuestFileCloseArguments{
//    Handle: openResponse.Return,
//  },
// })

// resp, err = dom.QemuAgentCommand(string(jsonCommand), libvirt.DOMAIN_QEMU_AGENT_COMMAND_MIN, 0)
// if err != nil {
//  panic(err)
// }
// log.Println(resp)
// sshConfig := &ssh.ClientConfig{
//  User: "ubuntu",
//  Auth: []ssh.AuthMethod{
//    ssh.PublicKeys(signer),
//  },
//  HostKeyCallback: ssh.InsecureIgnoreHostKey(),
// }
// var connUrl bytes.Buffer
// connUrl.WriteString(keys.Return[1].IpAddresses[0].IpAddress)
// connUrl.WriteString(":22")
// log.Println(connUrl.String())
// connection, err := ssh.Dial("tcp", connUrl.String(), sshConfig)
// if err != nil {
//  panic(err)
// }
// session, err := connection.NewSession()
// if err != nil {
//  panic(err)
// }
// err = session.Run("wget http://192.168.1.102:8000/")
// if err != nil {
//  panic(err)
// }
// session.Close()
// connection.Close()
// brokerList := strings.Split(os.Getenv("KAFKA_URL"), ",")
// config := sarama.NewConfig()
// config.Producer.RequiredAcks = sarama.WaitForAll
// config.Producer.Retry.Max = 10
// config.Producer.Return.Successes = true
// producer, err := sarama.NewSyncProducer(brokerList, config)
// if err != nil {
//   panic(err)
// }
// select {
// case <-stop:
//   log.Println("fuck")
// default:
//   log.Println("finish")
// }
// producer.Close()
// }
