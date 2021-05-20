package pkg

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/florinutz/liquid-api/pkg/auth"
)

// App stores an instance
var App *Application

type Credentials struct {
	ID     string
	Secret string
}

type Config struct {
	Credentials
}

type Application struct {
	Config Config
	Client *http.Client
}

func NewApp(creds Credentials, client *http.Client) *Application {
	return &Application{Config: Config{creds}, Client: client}
}

func (a *Application) jwt(path string) (string, error) {
	return auth.JWT(path, a.Config.ID, a.Config.Secret)
}

func (a *Application) Request(httpMethod string, path string, body io.Reader) (*http.Response, error) {
	jwt, err := a.jwt(path)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(httpMethod, "https://api.liquid.com"+path, body)
	if err != nil {
		return nil, fmt.Errorf("can't create RequestAndDeserialize: %w", err)
	}

	req.Header.Set("X-Quoine-Auth", jwt)
	req.Header.Set("X-Quoine-API-Version", "2")
	req.Header.Set("Content-Type", "application/json")

	res, err := a.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("can't perform RequestAndDeserialize: %w", err)
	}

	return res, nil
}

func InitAppE(app **Application) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		var secret, id string
		if secret = viper.GetString("secret"); secret == "" {
			return fmt.Errorf("An api secret is required. Please provide it in the config file (%s) "+
				"or as an env var (%s)", "~/.liquid.yaml", "LIQUID_SECRET")
		}
		if id = viper.GetString("token"); id == "" {
			return fmt.Errorf("An api token id is required. Please provide it in the config file (%s) "+
				"or as an env var (%s)", "~/.liquid.yaml", "LIQUID_TOKEN")
		}

		*app = NewApp(
			Credentials{
				ID:     id,
				Secret: secret,
			},
			&http.Client{Timeout: 10 * time.Second},
		)

		return nil
	}
}
