package domain

type HealthchecksFile struct {
	Healthchecks []*Healthcheck `json:"healthchecks"`
	Notifiers    []*Notifier `json:"notifiers"`
}
