package main


import (
    "fmt"
    "time"
    "github.com/vigneshuvi/GoDateFormat"
)


// OriginalTime ...
// https://programming.guide/go/format-parse-string-time-date-example.html
func OriginalTime(){
	input := "2017-08-31"
	layout := "2006-01-02" // same as YYYY-MM-DD
	t, _ := time.Parse(layout, input)
	fmt.Println(t)                       // 2017-08-31 00:00:00 +0000 UTC
	fmt.Println(t.Format("02-Jan-2006")) // 31-Aug-2017
}

// UseOtherLibrary ...
// https://github.com/vigneshuvi/GoDateFormat
func UseOtherLibrary(){
    fmt.Println("Go Date Format(Today - 'yyyy-MM-dd HH:mm:ss Z'): ", GetToday(GoDateFormat.ConvertFormat("yyyy-MM-dd HH:mm:ss Z")))
    fmt.Println("Go Date Format(Today - 'yyyy-MMM-dd'): ", GetToday(GoDateFormat.ConvertFormat("yyyy-MMM-dd")))
    fmt.Println("Go Time Format(NOW - 'HH:MM:SS'): ", GetToday(GoDateFormat.ConvertFormat("HH:MM:SS")))
    fmt.Println("Go Time Format(NOW - 'HH:MM:SS tt'): ", GetToday(GoDateFormat.ConvertFormat("HH:MM:SS tt")))
}

// GetToday ...
func GetToday(format string) (todayString string){
    today := time.Now()
    todayString = today.Format(format);
    return
}


func main() {
	OriginalTime()
	UseOtherLibrary()
}

