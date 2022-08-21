package config

import "go-mall-api/pkg/config"

func init() {
	config.Add("mq", func() map[string]interface{} {
		return map[string]interface{}{
			"host":     config.Env("MQ_HOST", "127.0.0.1"),
			"port":     config.Env("MQ_PORT", "5672"),
			"username": config.Env("MQ_USERNAME", "guser"),
			"password": config.Env("MQ_PASSWORD", "guser"),
		}
	})
}
