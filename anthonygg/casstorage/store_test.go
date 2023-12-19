package main

import (
	"bytes"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "mybestpicture"
	pathname := CASPathTranformFunc(key)
	expectedValue := "be17b/32c28/70b1c/0c73b/59949/db6a3/be781/4dd23"

	if pathname != expectedValue {
		t.Errorf("have %s want %s", pathname, expectedValue)
	}

}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: DefaultPathTransformFunc,
	}

	s := NewStore(opts)

	data := bytes.NewReader([]byte("some jpg file"))
	if err := s.writeStream("myspecialpicture", data); err != nil {
		t.Error(err)
	}

}
