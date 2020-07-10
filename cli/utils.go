package cli

func HandlePanic(f func()) (err interface{}) {
	err = nil
	defer func() {
		if r := recover(); r != nil {
			err = r
		}
	}()

	f()

	return
}
