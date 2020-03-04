package config

type (
	// Config stores the configuration settings.
	Config struct {
		Token string `envconfig:"TOKEN"`
		// CircleCI host, defaults to the public SaaS
		// Server users will have to change the default
		// value to your custom address (i.e. circleci.my-org.com).
		Host string `default:"https://circleci.com" envconfig:"HOST"`
	}
)
