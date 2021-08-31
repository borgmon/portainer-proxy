package main

import (
	"github.com/go-resty/resty/v2"
)

type PortainerClient struct {
	Client   resty.Client
	token    string
	Username string
	Password string
}

type JWT struct {
	JWT string `json:"jwt"`
}

type Container struct {
	Names []string
	State string
}

func (c *PortainerClient) refreshToken() error {
	res, err := c.Client.R().SetBody(map[string]string{"username": c.Username, "password": c.Password}).SetResult(&JWT{}).Post("/auth")
	if res.StatusCode() == 200 {
		c.token = res.Result().(*JWT).JWT
		return nil
	}
	return err
}

func (c *PortainerClient) GetStatus(retry bool) (map[string]string, error) {
	m := make(map[string]string)
	res, err := c.Client.R().SetResult([]*Container{}).SetAuthToken(c.token).Get("/endpoints/1/docker/containers/json")
	if res.StatusCode() == 200 {
		containers := res.Result().(*[]*Container)
		for i := range *containers {
			m[(*containers)[i].Names[0][1:len((*containers)[i].Names[0])]] = (*containers)[i].State
		}
		return m, err
	} else if res.StatusCode() == 401 && !retry {
		c.refreshToken()
		return c.GetStatus(true)
	} else {
		return nil, err
	}
}
