package apps

type Config struct {
	APPLICATION_NAME string `yaml:"APPLICATION_NAME,omitempty"`
	BUSINESS_NAME    string `yaml:"BUSINESS_NAME,omitempty"`
	TESTING_TAG      string `yaml:"TESTING_TAG,omitempty"`
	SERVER_NAME      string `yaml:"SERVER_NAME,omitempty"`
	Jenkins          map[string]interface{} `yaml:"jenkins,omitempty"`
	Docker           map[string]interface{} `yaml:"docker,omitempty"`
	Helm             Helm
}
