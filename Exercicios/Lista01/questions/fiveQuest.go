package questions

import "fmt"

// Read the duration of an event expressed in seconds, convert it,
// and display it in hours, minutes, and seconds.
func Five() {
	var time int

	fmt.Print("Enter the duration of an event in seconds: ")
	fmt.Scan(&time)

	hours := time / 3600
	minutes := (time / 60) % 60
	seconds := time % 60

	fmt.Printf("%d hours, %d minutes, and %d seconds\n", hours, minutes, seconds)
}
