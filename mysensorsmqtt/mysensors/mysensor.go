package mysensors

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	PREFX_IDX int = iota + 0
	GATWY_IDX
	NODE_IDX
	CHILD_IDX
	COMMD_IDX
	ACK_IDX
	STYPE_IDX
)

type MySensor struct {
	// node-id ; child-sensor-id ; command ; ack ; type ; payload
	// topic: tlmy/gar-out/20/0/1/0/47
	topic   string
	payload string

	prefix          string
	gateway         string
	node_id         Node
	child_sensor_id int
	command         CommandType
	ack             AckTypeIn
	stype           SetReqType
}

func NewMySensor(m, p string) MySensor {
	//parse messge 'm'
	var split = strings.Split(m, "/")

	var prefix string = split[PREFX_IDX]
	var gateway string = split[GATWY_IDX]
	var node_id Node
	nid, _ := strconv.Atoi(split[NODE_IDX])
	node_id = ItoNode(nid)
	cid, _ := strconv.Atoi(split[CHILD_IDX])
	var command CommandType
	cmd, _ := strconv.Atoi(split[COMMD_IDX])
	command = ItoCommandType(cmd)
	var ack AckTypeIn = NORM
	// ak, _ := strconv.Atoi(split[ACK_IDX])
	var stype SetReqType = V_TEMP
	// sty, _ := strconv.Atoi(split[STYPE_IDX])
	// stype = ItoSetReqType(sty)
	return MySensor{m, p, prefix, gateway, node_id, cid, command, ack, stype}
}

func (s *MySensor) String() string {
	return fmt.Sprintf("S:%s-%s-node=%s-child=%d/%v/%v/%v %s", s.prefix, s.gateway, s.node_id.String(), s.child_sensor_id, s.command.String(), s.ack.String(), s.stype.String(), s.payload)
}
