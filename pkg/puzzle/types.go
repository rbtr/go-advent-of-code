package puzzle

type Data []byte

type Solver func(d Data) (string, error)

type Puzzle struct {
	Data     Data
	Solution string
}
