package outbarriers

import (
	"math/rand"
	"time"
)

func RandomString(tam int) string {

	var v int32
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	str := ""
	i := 0
	for i < tam {
		n := r.Int31() % 3

		if n == 0 {
			v = (r.Int31() % 10) + 48
		} else if n == 1 {
			v = (r.Int31() % 26) + 97
		} else {
			v = (r.Int31() % 26) + 65
		}
		str += string(v)
		i += 1
	}
	return str
}
