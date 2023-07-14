package main

func nilerr() error {
	doSomething := func() error {
		return ErrUnknown
	}
	err := doSomething()
	if err != nil {
		return nil
	}
	return nil
}
