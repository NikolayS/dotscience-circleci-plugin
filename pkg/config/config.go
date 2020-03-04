package config

type (
	// Config stores the configuration settings.
	Config struct {
		Token    string `envconfig:"PLUGIN_TOKEN"`
		Username string `envconfig:"PLUGIN_USERNAME"`
		Project  string `envconfig:"PLUGIN_PROJECT"`
		// VCS type (github, bitbucket). Defaults to 'github'.
		VCSType string `default:"github" envconfig:"PLUGIN_VCS_TYPE"`

		// Revision- the specific revision to build.
		// Default is null and the head of the branch is used.
		// Cannot be used with tag parameter.
		Revision string `envconfig:"PLUGIN_REVISION"`
		Tag      string `envconfig:"PLUGIN_TAG"`
		Branch   string `envconfig:"PLUGIN_BRANCH"`

		BuildParameters map[string]string `envconfig:"PLUGIN_BUILD_PARAMETERS"`

		// CircleCI host, defaults to the public SaaS
		// Server users will have to change the default
		// value to your custom address (i.e. circleci.my-org.com).
		Host string `default:"https://circleci.com" envconfig:"PLUGIN_HOST"`
	}
)
