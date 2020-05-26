/*
Package envs contains the definition of environment variables
required by the app
*/
package envs

// ENVs represents the list of env vars
type ENVs struct {
	ENV                string `default:"dev"`
	DatabaseConnection string `required:"true" split_words:"true"`
	DatabaseName       string `default:"mayaleng_dev" split_words:"true"`
	Host               string `default:"0.0.0.0:80" split_words:"true"`
}
