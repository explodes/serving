package proto

import "fmt"

func (m Address) Address() string {
	return fmt.Sprintf("%s:%d", m.Host, m.Port)
}
