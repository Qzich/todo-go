package helpers

func Must(data interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return data
}
