package conversion

func StringsToBytes(in []string) [][]byte {
	out := make([][]byte, len(in))
	for i := range in {
		out[i] = []byte(in[i])
	}
	return out
}

func ParseString(in []byte) (string, error) {
	return string(in), nil
}
