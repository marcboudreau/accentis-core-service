package health

var serviceStatus Status = Status{}

const (
	// StatusHealthy is the status value for a healthy service.
	StatusHealthy = "healthy"

	// StatusNotHealthy is the status value for an unhealthy service.
	StatusNotHealthy = "not healthy"

	// MessageHealthy is the message value for a healthy service.
	MessageHealthy = "The service is operating normally."

	// MessageNotHealthy is the message value for an unhealthy service.
	MessageNotHealthy = "The service is not operating normally."
)

// Status is a structure that contains the necessary state needed by the Check
// function.
type Status struct {
	overall bool
}

// Check returns a status and a message describing the health of this service,
// unless an error occurs.
func Check() (status, message string, err error) {
	if serviceStatus.overall {
		return StatusHealthy, MessageHealthy, nil
	}

	return StatusNotHealthy, MessageNotHealthy, nil
}

// Set is a function that sets the health status according to the provided
// healthy parameter.
func Set(healthy bool) {
	serviceStatus.overall = healthy
}
