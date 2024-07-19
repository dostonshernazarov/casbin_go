package v1

import "github.com/casbin/casbin/v2"

type handler struct {
	enforces *casbin.Enforcer
}

type HandlerConfig struct {
	Enforcer *casbin.Enforcer
}

func New(h *HandlerConfig) handler {
	return handler{
		enforces: h.Enforcer,
	}
}
