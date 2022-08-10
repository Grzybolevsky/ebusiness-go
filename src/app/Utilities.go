package app

func CheckError(err error, comm string) {
	if err != nil {
		panic(comm)
	}
}
