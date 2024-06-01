package tools

func PanicOnErr(err ...error) {
	for _, e := range err {
		if e != nil {
			panic(e)
		}
	}
}
