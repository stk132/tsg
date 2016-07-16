package loader

import "testing"

func getPgParam() *Param {
	return &Param{
		user:     "stk132",
		pass:     "postgres",
		host:     "localhost",
		port:     "5432",
		database: "mydb",
	}
}
func TestPgLoad(t *testing.T) {
	p := getPgParam()
	l, err := NewPgLoader(p)
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
