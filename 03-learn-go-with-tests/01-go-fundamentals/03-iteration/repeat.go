package iteration

func Repeat(s string, num int) string {
	var out string
	for i := 0; i < num; i++ {
		out += s
	}
	return out
}
