package erratum

func Use(o ResourceOpener, input string) (resultErr error) {
	var resource Resource
	for {
		r, err := o()
		if err != nil {
			_, ok := err.(TransientError)
			if ok {
				continue
			}
			return err
		} else {
			resource = r
			break
		}
	}

	defer func() {
		err := recover()
		if err != nil {
			frobErr, ok := err.(FrobError)
			if ok {
				resource.Defrob(frobErr.defrobTag)
			}
			resultErr = err.(error)
		}
		resource.Close()
	}()
	resource.Frob(input)

	return resultErr
}
