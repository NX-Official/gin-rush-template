package tools

func Must(err ...error) {
	for _, e := range err {
		if e != nil {
			panic(e)
		}
	}
}
