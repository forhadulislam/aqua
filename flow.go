package aqua



func getOrFail(msg string) (string, error) {
	return "", errors.New(msg) 
}

func fail(msg string) (string, error) { 
	return "", errors.New(msg) 
}
		

func required(i interface{}) (interface{}, error) {
	if i == nil {
		return "", errors.New("mandatory value not found")
	}
	return i, nil
}