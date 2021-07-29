package main

import "fmt"

type Coords struct {
	Lat  float64
	Long float64
}

func (c *Coords) SetLatitude(latitude float64) {
	c.Lat = latitude
}

func (c *Coords) SetLongitude(longitude float64) {
	c.Long = longitude
}

func main() {
	coords := Coords{}
	coords.SetLatitude(15.30)
	coords.SetLongitude(30.45)
	fmt.Println(coords)
}

/*type Date struct {
	Day   int
	Month int
	Year  int
}

func (d *Date) SetDay(day int) {
	d.Day = day
}

func (d *Date) SetMon(mon int) {
	d.Month = mon
}

func (d *Date) SetYear(year int) {
	d.Year = year
}

func main() {
	d := Date{}
	d.SetDay(5)
	d.SetMon(7)
	d.SetYear(2021)
	fmt.Println(d)
}*/
