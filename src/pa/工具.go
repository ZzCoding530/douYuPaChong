package pa

func HandleError(err error, why string) {
	if err != nil {
		println(why, ":\t", err)
	}

	return
}
