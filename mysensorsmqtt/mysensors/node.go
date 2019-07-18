package mysensors

type Node int

const (
	unknown Node = 0
	wxmt    Node = 2
	wxm     Node = 1
	fam     Node = 20
	gar     Node = 30
	bob     Node = 50
	hvac    Node = 90
)

func ItoNode(n int) Node {
	switch n {
	case 1:
		return wxm
	case 2:
		return wxmt
	case 20:
		return fam
	case 30:
		return gar
	case 50:
		return bob
	case 90:
		return hvac
	default:
		return unknown
	}
}
