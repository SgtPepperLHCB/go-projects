package mysensors

type StreamType int

const (
	ST_FIRMWARE_CONFIG_REQUEST  StreamType = iota + 0 // Request new FW, payload contains current FW details
	ST_FIRMWARE_CONFIG_RESPONSE                       // New FW details to initiate OTA FW update
	ST_FIRMWARE_REQUEST                               // Request FW block
	ST_FIRMWARE_RESPONSE                              // Response FW block
	ST_SOUND                                          // Sound
	ST_IMAGE                                          // Image
)
