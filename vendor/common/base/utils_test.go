package base

import (
	"encoding/json"
	"fmt"
	"testing"
)

type User struct {
	Username string
	Password string
}

func TestIsZero(t *testing.T) {
	a := 0
	fmt.Println(IsZero(a))
	u := User{}
	u.Username = "Boy"
	json.Marshal()
	fmt.Println(IsZero(u))

}
