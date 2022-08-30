package methods

import "golang.org/x/tour/reader"

type MyReader struct{}

func (mr MyReader) Read(b []byte) (int, error) {
	for i := 0; i < len(b); i++ {
		b[i] = 'A'
	}

	return len(b), nil
}

func Methods22() {
	reader.Validate(MyReader{})
}
