// Code generated by "stringer -type InternalType internaltype.go"; DO NOT EDIT.

package mysensors

import "strconv"

const _InternalType_name = "I_BATTERY_LEVELI_TIMEI_VERSIONI_ID_REQUESTI_ID_RESPONSEI_INCLUSION_MODEI_CONFIGI_FIND_PARENTI_FIND_PARENT_RESPONSEI_LOG_MESSAGEI_CHILDRENI_SKETCH_NAMEI_SKETCH_VERSIONI_REBOOTI_GATEWAY_READYI_SIGNING_PRESENTATIONI_NONCE_REQUESTI_NONCE_RESPONSEI_HEARTBEATI_PRESENTATIONI_DISCOVERI_DISCOVER_RESPONSEI_HEARTBEAT_RESPONSEI_LOCKEDI_PINGI_PONGI_REGISTRATION_REQUESTI_REGISTRATION_RESPONSEI_DEBUG"

var _InternalType_index = [...]uint16{0, 15, 21, 30, 42, 55, 71, 79, 92, 114, 127, 137, 150, 166, 174, 189, 211, 226, 242, 253, 267, 277, 296, 316, 324, 330, 336, 358, 381, 388}

func (i InternalType) String() string {
	if i < 0 || i >= InternalType(len(_InternalType_index)-1) {
		return "InternalType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _InternalType_name[_InternalType_index[i]:_InternalType_index[i+1]]
}
