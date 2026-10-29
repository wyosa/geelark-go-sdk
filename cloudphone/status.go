package cloudphone

const (
	// StatusStarted means the cloud phone has started.
	StatusStarted = 0
	// StatusStarting means the cloud phone is starting.
	StatusStarting = 1
	// StatusStopped means the cloud phone is shut down.
	StatusStopped = 2
	// StatusExpired means the cloud phone subscription has expired.
	StatusExpired = 3
)

// List filters use their own values; note that OpenStatusOff is 0 while
// StatusStarted is also 0.
const (
	// OpenStatusOff filters List to powered-off cloud phones.
	OpenStatusOff = 0
	// OpenStatusOn filters List to powered-on cloud phones.
	OpenStatusOn = 1
	// ChargeModePerMinute selects pay-per-minute cloud phones.
	ChargeModePerMinute = 0
	// ChargeModeMonthly selects monthly subscription cloud phones.
	ChargeModeMonthly = 1
)
