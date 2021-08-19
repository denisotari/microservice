package proto

import "github.com/pkg/errors"

type User struct {
	Name string `json:"name"`
}

type HelloOut struct {
	Message string `json:"message"`
}

func (h *User) Validate() error {
	if h.Name == "" {
		return errors.New("you must provide a name")
	}
	return nil
}
