package errors

func DontPanicCalling(callPanicist func()) {
	defer func() {
		// doesn't matter if the channel is already closed
		recover()
	}()
	callPanicist()
}
