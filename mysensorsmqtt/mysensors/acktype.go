package mysensors

type AckTypeIn int

const (
	NORM AckTypeIn = 0
	ACK  AckTypeIn = 1
)

type AckTypeOut int

const (
	UNACK  AckTypeOut = 0
	REQACK AckTypeOut = 1
)
