package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/marcboudreau/accentis-core-service/health"
	"github.com/marcboudreau/accentis-core-service/models"
	"github.com/marcboudreau/accentis-core-service/restapi/operations/healthcheck"
)

// HandleHealthCheck is a handle function for the healthCheck operation.
func HandleHealthCheck(params healthcheck.HealthCheckParams) middleware.Responder {
	status, message, _ := health.Check()
	healthStatus := createHealthStatus(status, message)

	if status == "healthy" {
		return healthcheck.NewHealthCheckOK().WithPayload(healthStatus)
	}

	return healthcheck.NewHealthCheckServiceUnavailable().WithPayload(healthStatus)
}

func createHealthStatus(status, message string) *models.HealthStatus {
	return &models.HealthStatus{
		Status:  status,
		Message: message,
	}
}
