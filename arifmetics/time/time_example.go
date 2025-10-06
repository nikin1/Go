package main

import "fmt"

func main() {
	var d int
	fmt.Scan(&d)

	cnt_min := d * 2
	t_hour := cnt_min / 60
	t_min_hour := cnt_min % 60
	fmt.Println(cnt_min)

	fmt.Println("It is", t_hour, "hours", t_min_hour, "minutes.")

}
