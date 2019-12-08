package echo

import (
	"context"
	"log"

	echoservice "github.com/kaito2/goa-rest-api-sample/gen/echo_service"
)

// echo-service service example implementation.
// The example methods log the requests and return zero values.
type echoServicesrvc struct {
	logger *log.Logger
}

// NewEchoService returns the echo-service service implementation.
func NewEchoService(logger *log.Logger) echoservice.Service {
	return &echoServicesrvc{logger}
}

// EchoGet implements echo-get.
func (s *echoServicesrvc) EchoGet(ctx context.Context, p *echoservice.EchoGetPayload) (res string, err error) {
	s.logger.Print("echoService.echo-get")
	return
}
