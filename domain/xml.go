package domain

import (
	"bytes"
	"text/template"
)

type Config struct {
	ImagePath string
	Name      string
}

const DOMAIN_TEMPLATE = `
  <domain type="kvm">
     <name>{{ .Name }}</name>
     <memory>524288</memory>
     <os>
        <type>hvm</type>
        <boot dev="hd" />
     </os>
     <features>
        <acpi/>
     </features>
     <vcpu>1</vcpu>
     <devices>
        <channel type="unix">
           <target type="virtio" name="org.qemu.guest_agent.0" />
        </channel>
        <interface type="bridge">
           <source bridge="br0" />
           <virtualport type="openvswitch" />
           <address
             type="pci"
             domain="0x0000"
             bus="0x00"
             slot="0x03"
             function="0x0"
           />
        </interface>
        <disk type="file" device="disk">
           <driver type="qcow2" cache="none" />
           <source file="{{ .ImagePath }}" />
           <target dev="vda" bus="virtio" />
        </disk>
     </devices>
  </domain>
`

func NewDomainXml(config Config) (*string, error) {
	t := template.Must(template.New("domain").Parse(DOMAIN_TEMPLATE))
	var result bytes.Buffer
	if err := t.Execute(&result, config); err != nil {
		return nil, err
	}
	xml := result.String()
	return &xml, nil
}
