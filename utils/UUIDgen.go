package utils

import (
	"fmt"

	"github.com/segmentio/ksuid"
)

func UUIDGenerator(Prefix string) string {
	id := ksuid.New()
	uuid := Prefix + id.String()
	fmt.Println("Generated KSUID:", uuid)
	return uuid
}
