package loader

import "testing"

func getMyParam() *Param {
	return &Param{
		user:     "dbuser",
		pass:     "dbuser",
		host:     "localhost",
		port:     "3306",
		database: "golang_practice",
	}
}

func TestMyLoad(t *testing.T) {
	p := getMyParam()

	l, err := NewMyLoader(p)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = Load(l)
	if err != nil {
		t.Error(err)
		return
	}
}
