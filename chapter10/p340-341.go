//MAIN
package main

import (
	"fmt"
	"geo"
	"log"
)

func main() {
	coords := geo.Coordinates{}
	err := coords.SetLat(37.37)
	if err != nil {
		log.Fatal(err)
	}
	err = coords.SetLon(-9000.00)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("LAT:", coords.Lat(), "\tLONG:", coords.Lon())
}

//GEO
package geo

import "errors"

type Coordinates struct{
	latitude	float64
	longitude	float64
}

func (c *Coordinates) Lat() float64{
	return c.latitude
}

func (c *Coordinates) Lon() float64{
	return c.longitude
}

//

func (c *Coordinates) SetLat(latitude float64) error{
	if latitude < -90 || latitude > 90{
		return errors.New("invalid latitude")
	}
	c.latitude = latitude
	return nil
}

func (c * Coordinates) SetLon(longitude float64) error{
	if longitude < -90 || longitude > 90{
		return errors.New("invalid longitude")
	}
	c.longitude = longitude
	return nil
}