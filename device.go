package authy

import "os/user"

type Device struct {
	ID         string `json:"id"`
	SecretSeed string `json:"secret_seed"`
	APIKey     string `json:"api_key"`
	Reinstall  bool   `json:"reinstall"`
}

type DeviceRegistration struct {
}

func StartRegistration(user *user.User) (*DeviceRegistration, error) {
	return nil, nil
}
