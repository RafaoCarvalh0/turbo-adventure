package crash_course

import "fmt"

func TestPointer() {
	x := 5

	fmt.Println("x was ", x)

	increment(&x)

	fmt.Println("after increment, x is ", x)
}

func increment(px *int) {
	*px++
}
