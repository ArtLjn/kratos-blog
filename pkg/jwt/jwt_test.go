package jwt

import (
	"testing"
	"time"
)

func Test_Sign(t *testing.T) {
	token, err := Sign("红豆南墙")
	if err != nil {
		t.Errorf("%v\n", err)
	}
	t.Log(token)
	time.Sleep(2 * time.Second)
	data, e := Verify(token)
	if e != nil {
		t.Errorf("%v\n", e)
	}
	t.Log(data.Username)
	t.Log(GetLoginName("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6Iue6ouixhuWNl-WimSIsImlzcyI6IkF1dGhfU2VydmVyIiwic3ViIjoiQXV0aCIsImV4cCI6MTcxMzU5NjMyOSwibmJmIjoxNzEyOTkxNTMwLCJpYXQiOjE3MTI5OTE1MjksImp0aSI6IjE2bXppZ2JzS0cifQ.s1sdfzhq2OiUhU54vL53SZB49kvjGZ5nqOiooBujVu4"))
}
