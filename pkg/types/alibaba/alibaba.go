package alibaba

var (
	StatusPending = "Pending"
	StatusRunning = "Running"
)

type Options struct {
	AccessKey               string `json:"access-key,omitempty" yaml:"access-key,omitempty"`
	AccessSecret            string `json:"access-secret,omitempty" yaml:"access-secret,omitempty"`
	DiskCategory            string `json:"disk-category,omitempty" yaml:"disk-category,omitempty"`
	DiskSize                string `json:"disk-size,omitempty" yaml:"disk-size,omitempty"`
	Image                   string `json:"image,omitempty" yaml:"image,omitempty"`
	Type                    string `json:"type,omitempty" yaml:"type,omitempty"`
	KeyPair                 string `json:"key-pair,omitempty" yaml:"key-pair,omitempty"`
	Region                  string `json:"region,omitempty" yaml:"region,omitempty"`
	VSwitch                 string `json:"v-switch,omitempty" yaml:"v-switch,omitempty"`
	SecurityGroup           string `json:"security-group,omitempty" yaml:"security-group,omitempty"`
	InternetMaxBandwidthOut string `json:"internet-max-bandwidth-out,omitempty" yaml:"internet-max-bandwidth-out,omitempty"`
}
