package actions

type Actions interface {
	Reboot() error
	Shutdown() error
}
