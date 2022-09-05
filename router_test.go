package gms_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/he-wen-yao/gms"
)

func newTestRouter() *gms.Router {
	r := gms.NewRouter()
	r.AddRoute("GET", "/", nil)
	r.AddRoute("GET", "/hello/:name", nil)
	r.AddRoute("GET", "/hello/b/c", nil)
	r.AddRoute("GET", "/hi/:name", nil)
	r.AddRoute("GET", "/assets/*filepath", nil)
	return r
}

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(gms.ParsePattern("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(gms.ParsePattern("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(gms.ParsePattern("/p/*name/*"), []string{"p", "*name"})
	if !ok {
		t.Fatal("test parsePattern failed")
	}
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n, ps := r.GetRoute("GET", "/hello/geektutu")

	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}

	if n.Pattern != "/hello/:name" {
		t.Fatal("should match /hello/:name")
	}

	if ps["name"] != "geektutu" {
		t.Fatal("name should be equal to 'geektutu'")
	}

	fmt.Printf("matched path: %s, params['name']: %s\n", n.Pattern, ps["name"])

}
