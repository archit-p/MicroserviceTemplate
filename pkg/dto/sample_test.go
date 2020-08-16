package dto

import (
	"bytes"
	"testing"
)

func TestUnit_Sample_ToJSON(t *testing.T) {
	s := Sample{
		Content: "testing",
	}

	buf := new(bytes.Buffer)
	err := s.ToJSON(buf)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnit_Sample_FromJSON(t *testing.T) {
	s := Sample{
		Content: "testing",
	}

	buf := new(bytes.Buffer)
	err := s.ToJSON(buf)
	if err != nil {
		t.Fatal(err)
	}

	err = s.FromJSON(buf)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnit_Samples_ToJSON(t *testing.T) {
	si := Sample{
		Content: "testing",
	}

	var s Samples
	s = append(s, si)

	buf := new(bytes.Buffer)
	err := s.ToJSON(buf)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnit_Samples_FromJSON(t *testing.T) {
	si := Sample{
		Content: "testing",
	}

	var s Samples
	s = append(s, si)

	buf := new(bytes.Buffer)
	err := s.ToJSON(buf)
	if err != nil {
		t.Fatal(err)
	}

	err = s.FromJSON(buf)
	if err != nil {
		t.Fatal(err)
	}
}
