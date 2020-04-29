package lib

// CheckError is it.
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
