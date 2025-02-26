package compression

type Encode interface {
	Encode(str string) []byte
}

type Decoder interface {
	Decode(data []byte) string
}
