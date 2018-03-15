package domain

import (
	"encoding/base64"
	"encoding/json"
	libvirt "github.com/libvirt/libvirt-go"
)

type AgentDomain struct {
	*libvirt.Domain
}

func (self *AgentDomain) GetNetworkInterfaces() (*AgentNetworkCommandReply, error) {
	var agentNetworkCommand AgentNetworkCommand
	var agentCommandReply AgentNetworkCommandReply
	agentNetworkCommand.Execute = "guest-network-get-interfaces"
	agentCommandRequest, err := json.Marshal(agentNetworkCommand)
	if err != nil {
		return nil, err
	}
	agentCommandReplyJson, err := self.QemuAgentCommand(
		string(agentCommandRequest),
		libvirt.DOMAIN_QEMU_AGENT_COMMAND_MIN,
		0,
	)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(agentCommandReplyJson), &agentCommandReply)
	if err != nil {
		return nil, err
	}
	return &agentCommandReply, nil
}

func (a *AgentDomain) HighLevelWriteFile(path string, content []byte) error {
	fileOpen, err := a.QemuOpenFile(path, "w+")
	if err != nil {
		return err
	}
	_, err = a.QemuWriteFile(fileOpen.Return, content)
	if err != nil {
		return err
	}
	_, err = a.QemuCloseFile(fileOpen.Return)
	if err != nil {
		return err
	}
	return nil
}

func (a *AgentDomain) QemuOpenFile(path, mode string) (*GuestFileOpenReply, error) {
	var guestFileOpenReply GuestFileOpenReply
	var guestFileOpen GuestFileOpen
	guestFileOpen.Execute = "guest-file-open"
	guestFileOpen.Arguments.Path = path
	guestFileOpen.Arguments.Mode = mode
	guestFileOpenRequest, err := json.Marshal(guestFileOpen)
	if err != nil {
		return nil, err
	}
	guestFileOpenReplyJson, err := a.QemuAgentCommand(
		string(guestFileOpenRequest),
		libvirt.DOMAIN_QEMU_AGENT_COMMAND_MIN,
		0,
	)
	if err != nil {
		return nil, err
	}
	json.Unmarshal([]byte(guestFileOpenReplyJson), &guestFileOpenReply)
	return &guestFileOpenReply, nil
}

func (a *AgentDomain) QemuWriteFile(handle int, content []byte) (*GuestFileWriteReply, error) {
	var guestFileWrite GuestFileWrite
	var guestFileWriteReply GuestFileWriteReply
	guestFileWrite.Execute = "guest-file-write"
	guestFileWrite.Arguments.Handle = handle
	guestFileWrite.Arguments.BufB64 = base64.StdEncoding.EncodeToString(content)
	guestFileWriteRequest, err := json.Marshal(guestFileWrite)
	if err != nil {
		return nil, err
	}
	guestFileWriteReplyJson, err := a.QemuAgentCommand(
		string(guestFileWriteRequest),
		libvirt.DOMAIN_QEMU_AGENT_COMMAND_MIN,
		0,
	)
	if err != nil {
		return nil, err
	}
	json.Unmarshal([]byte(guestFileWriteReplyJson), &guestFileWriteReply)
	return &guestFileWriteReply, nil
}

func (a *AgentDomain) QemuCloseFile(handle int) (*GuestFileCloseReply, error) {
	var guestFileClose GuestFileClose
	var guestFileCloseReply GuestFileCloseReply
	guestFileClose.Execute = "guest-file-close"
	guestFileClose.Arguments.Handle = handle
	guestFileCloseRequest, err := json.Marshal(guestFileClose)
	if err != nil {
		return nil, err
	}
	guestFileCloseReplyJson, err := a.QemuAgentCommand(
		string(guestFileCloseRequest),
		libvirt.DOMAIN_QEMU_AGENT_COMMAND_MIN,
		0,
	)
	if err != nil {
		return nil, err
	}
	json.Unmarshal([]byte(guestFileCloseReplyJson), &guestFileCloseReply)
	return &guestFileCloseReply, nil
}
