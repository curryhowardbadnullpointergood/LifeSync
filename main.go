package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {

	daysOfWeek := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	monthsOfYear := [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	// monthsWith31Days := [7]int{0, 2, 4, 6, 7, 9, 11}

	// monthsWith30Days := [4]int{3, 5, 8, 10}

	currentyear := 2024

	currentmonth := monthsOfYear[8]

	currentday := daysOfWeek[0]

	// numerical representation of date

	// currentdatenumerical = ie 12 , currentmonthnumerical = ie 03 for March

	fmt.Println(currentyear, currentmonth, currentday)

	// making a window for the app

	a := app.New()
	w := a.NewWindow("Hello World")

	w.SetContent(widget.NewLabel("Hello World!"))
	w.Show()

	a.Run()

	tidyup()

}

func tidyup() {
	fmt.Println("Exited sucessfully!")
}

blah 
