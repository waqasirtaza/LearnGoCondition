package main
import "fmt"

func accessGranted()  {

fmt.Println("Access Granted !!!")

}
func accessDenied()  {

	fmt.Println("Access Denied !!!")
	
	}
//Days of the week
const (

	Monday = 0
	Tuesday = 1
	Wednesday = 2
	Thrusday = 3
	Friday = 4
	Saturday = 5
	Sunday = 6
)
//User roles
const (

	Admin = 10
	Manager = 20
	Contractor = 30
	Member = 40
	Guest = 50
)
func weekday(day int) bool {

	return day <= 4
}
func main() {

	today, role := Sunday, Contractor
	if role == Admin || role == Manager {

		accessGranted()
	} else if role == Contractor && !weekday(today) {
		accessGranted()
	} else if role == Member && weekday(today) {
        accessGranted()
	} else if role == Guest && (today == Monday || today == Wednesday || today == Friday) {
		accessGranted()
	} else {
		accessDenied() 
	}

	




}