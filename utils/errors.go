package utils

func ErroHandler(err error) {
	if err != nil {
		panic(err)
	}
}
