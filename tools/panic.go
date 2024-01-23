package tools

func PanicIfErr(err ...error) {
	for _, e := range err {
		if e != nil {
			panic(e)
		}
	}
}
