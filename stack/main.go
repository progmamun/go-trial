package main

func main() {
	var a, b = 0, 1
	f1(a, b)
	f2(a)
}
func f1(a, b int) {
	c := a + b
	f2(c)
}
func f2(c int) {
	print(c)
}
