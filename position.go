package dimension

import (
	"math"
	"strconv"
	"strings"
)

type Position struct {
	// coordinates is a slice of ints representing the position of the position
	Coordinates []int
}

// NewPosition returns a new Position of a given dimension
func NewPosition(dimensions int) *Position {
	return &Position{
		Coordinates: make([]int, dimensions),
	}
}

// NewFromString creates a new coordinate from a string of integers and a separator
func NewPositionFromString(str string, sep string) *Position {
	// split the string by the separator
	strs := strings.Split(str, sep)

	// list of integers
	ints := make([]int, 0)

	// turn each string into an integer and append it to the list
	for _, s := range strs {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		ints = append(ints, i)
	}

	// create a new position
	pos := NewPosition(len(ints))

	// set the coordinates of the position
	pos.Set(ints)

	// return the position
	return pos
}

// Set sets the coordinates of the position
func (p *Position) Set(coordinates []int) {
	// add a check to make sure the length of the coordinates is the same as the dimension
	if len(coordinates) != len(p.Coordinates) {
		panic("coordinates length does not match the dimension")
	}

	// copy the coordinates
	copy(p.Coordinates, coordinates)
}

// Copy returns a shallow copy of the position
func (p *Position) Copy() *Position {
	// create a new position and return it
	return &Position{
		Coordinates: p.Coordinates,
	}
}

// Clone returns a clone of the position
func (p *Position) Clone() *Position {
	// create a new position
	pos := NewPosition(len(p.Coordinates))

	// copy the coordinates
	pos.Set(p.Coordinates)

	// return the new position
	return pos
}

// Equals returns true if the positions are equal
func (p *Position) Equals(pos *Position) bool {
	// check if the coordinates are the same
	for i := 0; i < len(p.Coordinates); i++ {
		if p.Coordinates[i] != pos.Coordinates[i] {
			return false
		}
	}

	// return true if all coordinates are the same
	return true
}

// Add adds the given position to the current position
func (p *Position) Add(pos *Position) {
	// add the coordinates
	for i := 0; i < len(p.Coordinates); i++ {
		p.Coordinates[i] += pos.Coordinates[i]
	}
}

// Subtract subtracts the given position from the current position
func (p *Position) Subtract(pos *Position) {
	// subtract the coordinates
	for i := 0; i < len(p.Coordinates); i++ {
		p.Coordinates[i] -= pos.Coordinates[i]
	}
}

// Distance returns the distance between the current position and the given position
func (p *Position) Distance(pos *Position) float64 {
	// create a new position
	diffpos := NewPosition(len(p.Coordinates))

	// copy the coordinates of the called position
	diffpos.Set(p.Coordinates)

	// subtract the given position from the current position
	diffpos.Subtract(pos)

	// return the length of the position (from the origin because it is a difference)
	return diffpos.Length()
}

// Length returns the euclidean length of the position from the origin
func (p *Position) Length() float64 {
	// square root of the sum of the squares of the coordinates
	var sum float64 = 0

	// add the squares of the coordinates
	for i := 0; i < len(p.Coordinates); i++ {
		// add the square of the coordinate p.Coordinates[i]
		sum += math.Pow(float64(p.Coordinates[i]), 2)
	}

	// return the square root of the sum
	return math.Sqrt(sum)
}

// String returns the string representation of the position
func (p *Position) String() string {
	// create a new string
	str := ""

	// add the coordinates to the string
	for i := 0; i < len(p.Coordinates); i++ {
		str += strconv.Itoa(p.Coordinates[i])
		if i < len(p.Coordinates)-1 {
			str += ","
		}
	}

	// return the string
	return str
}

// PString returns the string representation of the position with parenthesis and spacing
func (p *Position) PString() string {
	// create a new string
	str := "("

	// add the coordinates to the string
	for i := 0; i < len(p.Coordinates); i++ {
		str += strconv.Itoa(p.Coordinates[i])
		if i < len(p.Coordinates)-1 {
			str += ", "
		}
	}

	// add the closing parenthesis
	str += ")"

	// return the string
	return str
}
