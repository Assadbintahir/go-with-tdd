package iteration

// Repeat takes a character and appends it to string count times
func Repeat(value string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += value
	}
	return repeated
}
