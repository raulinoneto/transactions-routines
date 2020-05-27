package httpadapters

import (
	"reflect"
	"testing"
)

func TestNewAccount(t *testing.T) {
	id := 1
	documentNumber := 123
	expected := AccountBody{id, documentNumber}
	result := NewAccount(id, documentNumber)
	if !reflect.DeepEqual(expected, *result) {
		t.Errorf("Expected: %d \n Got: %d", expected, *result)
	}
}

func TestAccount_GetDocumentNumber(t *testing.T) {
	documentNumber := 123
	a := NewAccount(1, documentNumber)
	result := a.GetDocumentNumber()
	if result != documentNumber {
		t.Errorf("Expected: %d \n Got: %d", documentNumber, result)
	}
}

func TestAccount_GetID(t *testing.T) {
	id := 1
	a := NewAccount(id, 123)
	result := a.GetID()
	if result != id {
		t.Errorf("Expected: %d \n Got: %d", id, result)
	}
}

func TestAccount_SetDocument(t *testing.T) {
	documentNumber := 123
	a := NewAccount(1, 432)
	a.SetDocumentNumber(documentNumber)
	if a.DocumentNumber != documentNumber {
		t.Errorf("Expected: %d \n Got: %d", documentNumber, a.DocumentNumber)
	}
}

func TestAccount_SetID(t *testing.T) {
	id := 1
	a := NewAccount(2, 123)
	a.SetID(id)
	if a.ID != id {
		t.Errorf("Expected: %d \n Got: %d", id, a.ID)
	}
}
