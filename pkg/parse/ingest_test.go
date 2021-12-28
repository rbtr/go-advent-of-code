package parse

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var want = `aaa
bbb
ccc
`

func TestReadData(t *testing.T) {
	got, err := ReadData("testdata/test")
	require.NoError(t, err)
	require.Equal(t, want, string(got))
}
