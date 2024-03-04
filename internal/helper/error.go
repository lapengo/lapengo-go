package helper

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func ErrStringConcat(errTexts ...string) string {
	concated := ""

	for i, v := range errTexts {
		concated += v
		if i == 0 {
			concated += " | "
		}
	}

	return concated
}
