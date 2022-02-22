package dimension

// Dimension rerpesents all the data on every plane
type Dimension struct {
	// a grid contains the information for each position
	grid Grid

	// defaultVal is the default value for empty coordinates in the dimension
	defaultVal interface{}
}

// NewDimension
func NewDimension(defaultVal interface{}) *Dimension {
	return &Dimension{
		grid:       make(Grid),
		defaultVal: defaultVal,
	}
}

// Grid is a map of values of any type
type Grid map[string]interface{}

// NewGrid returns a new Grid
func NewGrid() *Grid {
	return &Grid{}
}

// SetPos sets the value of a position in the grid
func (d *Dimension) SetPos(pos *Position, value interface{}) {
	// set the positions string
	d.grid.set(pos.String(), value)
}

// GetPos returns the value of a position in the grid
func (d *Dimension) GetPos(pos *Position) interface{} {
	// return the value of the position if it exists
	return d.get(pos.String())
}

// DelPos deletes the value of a position in the grid
func (d *Dimension) DelPos(pos *Position) {
	// delete the value of the position if it exists
	d.grid.remove(pos.String())
}

// GetAllPos returns all the positions in the grid
func (d *Dimension) GetAllPos() []*Position {
	// create a slice of positions
	positions := make([]*Position, 0)

	// iterate over the grid
	for key := range d.grid {
		// create a new position from the key
		pos := NewPositionFromString(key, ",")

		// add the position to the slice
		positions = append(positions, pos)
	}

	// return the slice of positions
	return positions
}

// set sets the value of a key in the grid
func (g *Grid) set(key string, value interface{}) {
	(*g)[key] = value
}

// get returns the value of a key in the grid
func (d *Dimension) get(key string) interface{} {
	// return the value of the key if it exists
	if val, ok := (d.grid)[key]; !ok {
		return d.defaultVal
	} else {
		return val
	}
}

// remove removes a key from the grid
func (g *Grid) remove(key string) {
	delete(*g, key)
}
