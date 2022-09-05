package gms_test

import (
	"testing"

	"github.com/he-wen-yao/gms"
)

func TestGms(t *testing.T) {
	g := gms.NewGms()

	api := g.Group("/test")

	t.Logf("%v", api)
}
