package version

import (
	"fmt"

	"github.com/imroc/req/v3"
	"github.com/sirupsen/logrus"
)

type (
	// Plugin defines the plugin interface.
	Plugin struct {
		URL     string `json:"url"`
		Token   string `json:"token"`
		Name    string `json:"name"`
		Release string `json:"release"`
		Type    string `json:"type"`
	}
)

func (p Plugin) Exec() error {
	client := req.C()
	if p.Token == "" || p.URL == "" || p.Name == "" {
		return fmt.Errorf("invalid plugin config")
	}
	if len(p.Type) == 0 {
		p.Type = "none"
	}
	_, err := client.R().SetHeader("X-Auth-Token", p.Token).Get(fmt.Sprintf("%s/new/%s?version=%s&type=%s", p.URL, p.Name, p.Release, p.Type))
	if err != nil {
		return err
	}
	logrus.Infof("%s version %s is available", p.Name, p.Release)
	return nil
}
