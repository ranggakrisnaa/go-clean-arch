package v1

import (
	v1 "github.com/ranggakrisna/go-clean-arch/internal/controller/amqp_rpc/v1"
	"github.com/ranggakrisna/go-clean-arch/internal/usecase"
	"github.com/ranggakrisna/go-clean-arch/pkg/logger"
	"github.com/ranggakrisna/go-clean-arch/pkg/rabbitmq/rmq_rpc/server"
)

// NewRouter -.
func NewRouter(t usecase.Translation, l logger.Interface) map[string]server.CallHandler {
	routes := make(map[string]server.CallHandler)

	{
		v1.NewTranslationRoutes(routes, t, l)
	}

	return routes
}
