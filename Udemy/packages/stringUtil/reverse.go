package stringUtil

func Reverse(s string) string {
	return reverseMe(s) //show how an unexported func can be exported using exported one
}
