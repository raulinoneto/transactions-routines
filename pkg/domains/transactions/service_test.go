package transactions

import (
	"reflect"
	"testing"
)

func TestNewService(t *testing.T) {
	service := NewService(nil, nil, nil)
	if !reflect.DeepEqual(*service, *new(Service)) {
		t.Error("Is not instance of service")
	}
}
