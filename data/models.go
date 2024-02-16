package data

type Status struct {
	Name      string
	FinalSize int64
	Parts     []int64
	Done      bool
	Err       error
}

type ReadData struct {
	N   int
	Err error
}
