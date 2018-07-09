package reverse

func String(str string) string{
	s := ""
	for _, letter := range str{
		s = string(letter) + s
	}

	return s
}
