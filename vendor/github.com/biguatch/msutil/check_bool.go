package msutil

func CheckBoolen(v interface{}) bool {
	var trues = []interface{}{
		1,
		true,
		"yes",
		"check",
		"true",
		"1",
	}
	return InArray(v, trues)
}
