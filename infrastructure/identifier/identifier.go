package identifier

import "github.com/google/uuid"

type Identifier interface {
	NewUuid() string
}

type identifier struct{}

func NewIdentifier() Identifier {
	return &identifier{}
}

func (i identifier) NewUuid() string {
	return uuid.New().String()
}
