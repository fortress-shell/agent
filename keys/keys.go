package keys

import (
	"crypto/rand"
	"crypto/rsa"
	"golang.org/x/crypto/ssh"
)

const (
	USERNAME = "ubuntu"
)

func NewKeyPair() (*ssh.ClientConfig, []byte, error) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}
	signer, err := ssh.NewSignerFromKey(key)
	if err != nil {
		return nil, nil, err
	}
	pub, err := ssh.NewPublicKey(&key.PublicKey)
	if err != nil {
		return nil, nil, err
	}
	public := ssh.MarshalAuthorizedKey(pub)
	sshConfig := &ssh.ClientConfig{
		User: USERNAME,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	return sshConfig, public, nil
}
