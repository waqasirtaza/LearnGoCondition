package main
import "fmt"

func price() int{
	return 10
}
const (

Economy		= 0
Business	= 1
FirstClass	= 2
)
func main() {

	switch p := price(); {
	case p < 2:
		fmt.Println("Cheap Item")
	case p > 10:
		fmt.Println("Moderately priced item")
	default:
		fmt.Println("Expensive item")
	}

	ticket := Economy
	switch ticket {
	case Economy:
		fmt.Println("Economy class")

	case Business:
		fmt.Println("Business Class")

	case FirstClass:
		fmt.Println("First Class")
	default:
		fmt.Println("Other category")
	}


}