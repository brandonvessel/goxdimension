package main

// import grid package
import (
	"fmt"

	dimensions "github.com/brandonvessel/goxdimension"
)

func main() {
	// create a new dimension
	d := dimensions.NewDimension(2)

	// create some positions
	pos1 := dimensions.NewPosition(2)
	pos2 := dimensions.NewPosition(2)

	// set position values
	pos1.Set([]int{0, 0})
	pos2.Set([]int{1, 1})

	// add positions to dimension
	d.SetPos(pos1, "hello")
	d.SetPos(pos2, "world")

	// get values at each position in the dimension
	fmt.Println("Position values:")
	fmt.Println(d.GetPos(pos1))
	fmt.Println(d.GetPos(pos2))

	// get all positions
	fmt.Println("All positions:")
	fmt.Println(d.GetAllPos())

	// print distance between positions 1 and 2
	fmt.Println("Distance between positions 1 and 2:")
	fmt.Println(pos1.Distance(pos2))

	// pretty print
	fmt.Println("Pretty print:")
	fmt.Println(pos1.PString())
	fmt.Println(pos2.PString())
}
