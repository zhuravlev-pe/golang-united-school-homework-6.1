package golang_united_school_homework

import "fmt"

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) >= b.shapesCapacity {
		return fmt.Errorf("box is at max capacity (%d)", b.shapesCapacity)
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if err := b.isIndexInRange(i); err != nil {
		return nil, err
	}
	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if err := b.isIndexInRange(i); err != nil {
		return nil, err
	}
	result := b.shapes[i]
	b.shapes = removeShape(b.shapes, i)
	return result, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if err := b.isIndexInRange(i); err != nil {
		return nil, err
	}
	result := b.shapes[i]
	b.shapes[i] = shape
	return result, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64
	for _, shape := range b.shapes {
		sum += shape.CalcPerimeter()
	}
	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sum float64
	for _, shape := range b.shapes {
		sum += shape.CalcArea()
	}
	return sum

}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	result := make([]Shape, 0, len(b.shapes))
	for _, shape := range b.shapes {
		_, ok := shape.(*Circle)
		if !ok {
			result = append(result, shape)
		}
	}
	if len(result) == len(b.shapes) {
		return fmt.Errorf("box does not contain Circles")
	}
	b.shapes = result
	return nil
}

func (b *box) isIndexInRange(i int) error {
	if (i < 0) || (i >= len(b.shapes)) {
		return fmt.Errorf("index %d is out of range", i)
	}
	return nil
}

func removeShape(shapes []Shape, i int) []Shape {
	shapes = append(shapes[:i], shapes[i+1:]...)
	return shapes
}
