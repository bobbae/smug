package main

func main() {
	foo(1, 2, 3)
}

func foo(one, two, three int) {
	bar(one, two, three)
}

func bar(one, two, three int) {
	panic("attack!")
}
