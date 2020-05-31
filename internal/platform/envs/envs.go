/*
Package envs contains the definition of environment variables
required by the app
*/
package envs

// Envs represents the env vars for non-testing
type Envs struct {
	Env                string `default:"dev"`
	DatabaseConnection string `required:"true" split_words:"true"`
	DatabaseName       string `default:"mayaleng" split_words:"true"`
	Host               string `default:"0.0.0.0:80" split_words:"true"`
}

// TestEnvs represents the env vars used in testing
type TestEnvs struct {
	Env                string `default:"testing"`
	DatabaseConnection string `default:"mongodb://localhost" split_words:"true"`
	DatabaseName       string `default:"mayaleng_dev" split_words:"true"`
	Host               string `default:"0.0.0.0:80" split_words:"true"`
}
