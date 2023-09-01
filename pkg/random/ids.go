package random

import (
	"strings"

	"github.com/google/uuid"
)

func NewTraceID() string {
	id := uuid.New().String()
	// repalce "-" with ""
	id = strings.Replace(id, "-", "", -1)
	return id
}