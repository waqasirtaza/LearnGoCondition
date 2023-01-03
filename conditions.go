package main
import "fmt"

func average(a,b,c int) float32 {

return float32(a+b+c) / 3

}

func main() {

	quiz1, quiz2, quiz3 := 9, 7, 8

	if quiz1 > quiz2 {

		fmt.Println("quiz1 score higher then quiz2")
	} else if quiz1<quiz2 {
		fmt.Println("quiz2 score higher then quiz1")
	} else {

		fmt.Println("quiz1 & quiz2 are equal")
	}
	if average(quiz1, quiz2, quiz3) > 7 {

		fmt.Println("Acceptable Grade")
	}


}