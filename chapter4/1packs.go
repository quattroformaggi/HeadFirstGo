package main
//Exercise 1: Importing Packages
//Write a program that imports all the above packages and calls all the functions they contain.
//You can include the program in your Go workspace, if you like, or just store it anywhere and run it with go run.

// Remember, the import path for a package is its path
// within the "src" subdirectory in your workspace.
import (
	"github.com/jaymcgavren/car"
	"github.com/jaymcgavren/car/headlights"
	// All the .go source files in a package directory are
	// combined into a single package. So this one line
	// will give you access to the functions from both the
	// amplifier.go and speakers.go source files.
	"github.com/jaymcgavren/car/stereo"
	"github.com/jaymcgavren/car/wheels"
)

func main() {
	// Precede all function names with the package name.
	car.OpenDoor()
	// The package name is usually the same as the last
	// segment of the import path. So if you import the
	// "car/headlights" package, you'll refer to it as just
	// "headlights".
	headlights.TurnOn()
	stereo.TurnOn()
	stereo.BoostBass()
	wheels.Steer()
	wheels.Accelerate()
}