package mysensors

type CommandType int

const (
	Presentation CommandType = iota + 0
	Set
	Req
	Internal
	Stream
	UnknownCommandType
)

func ItoCommandType(n int) CommandType {
	switch n {
	case 0:
		return Presentation
	case 1:
		return Set
	case 2:
		return Req
	case 3:
		return Internal
	case 4:
		return Stream
	default:
		return UnknownCommandType
	}
}
