package util

import (
	"github.com/satori/go.uuid"
	"strings"
)

func Uuid4() string {
	uuidNew := uuid.NewV4()
	return uuidNew.String()
}

func Uuid4NonSymbol() string {
	uuidNew := uuid.NewV4()
	return strings.ReplaceAll(uuidNew.String(), "-", "")
}
