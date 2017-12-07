package v1

import (
	"testing"
)

func TestRegistService(t *testing.T) {
	s := Service{}
	s.Name = "cmdb"
	s.Mode = "dev"
	s.Address = "dsf"
	s.TTL = 60
	c := NewClient("192.168.10.127")
	result, err := c.RegistService(s)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)

	s.Mode = "prod"
	s.Address = "sfsffsd"
	result, err = c.RegistService(s)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestGetService(t *testing.T) {

	c := NewClient("192.168.10.127")
	result, err := c.GetService("cmdb", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}
