package mysensors

type MessageType int

const (
	Presentation MessageType = iota + 0
	Set
	Req
	Internal
	Stream
)

func (mtype MessageType) String() string {
	names := []string{
		"presentation",
		"set",
		"req",
		"internal",
		"stream",
	}
	return names[mtype]
}
