package config

import "go-mall-api/pkg/config"

func init() {
	config.Add("grpc", func() map[string]interface{} {
		return map[string]interface{}{
			"link": config.Env("GRPC_LINK", "tcp"),
			"port": config.Env("GRPC_PORT", ":8972"),
		}
	})
}
