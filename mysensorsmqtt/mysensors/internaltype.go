package mysensors

type InternalType int

const (
	I_BATTERY_LEVEL        InternalType = iota + 0 // Use this to report the battery level (in percent 0-100).
	I_TIME                                         // Sensors can request the current time from the Controller using this message. The time will be reported as the seconds since 1970
	I_VERSION                                      // Sensors report their library version at startup using this message type
	I_ID_REQUEST                                   // Use this to request a unique node id from the controller.
	I_ID_RESPONSE                                  // Id response back to sensor. Payload contains sensor id.
	I_INCLUSION_MODE                               // Start/stop inclusion mode of the Controller (1=start, 0=stop).
	I_CONFIG                                       // Config request from node. Reply with (M)etric or (I)mperal back to sensor
	I_FIND_PARENT                                  // When a sensor starts up, it broadcast a search request to all neighbor nodes. They reply with a I_FIND_PARENT_RESPONSE.
	I_FIND_PARENT_RESPONSE                         // Reply message type to I_FIND_PARENT request.
	I_LOG_MESSAGE                                  // Sent by the gateway to the Controller to trace-log a message
	I_CHILDREN                                     // A message that can be used to transfer child sensors (from EEPROM routing table) of a repeating node.
	I_SKETCH_NAME                                  // Optional sketch name that can be used to identify sensor in the Controller GUI
	I_SKETCH_VERSION                               // Optional sketch version that can be reported to keep track of the version of sensor in the Controller GUI.
	I_REBOOT                                       // Used by OTA firmware updates. Request for node to reboot.
	I_GATEWAY_READY                                // Send by gateway to controller when startup is complete
	I_SIGNING_PRESENTATION                         // Provides signing related preferences (first byte is preference version).
	I_NONCE_REQUEST                                // Request for a nonce.
	I_NONCE_RESPONSE                               // Payload is nonce data.
	I_HEARTBEAT
	I_PRESENTATION
	I_DISCOVER
	I_DISCOVER_RESPONSE
	I_HEARTBEAT_RESPONSE
	I_LOCKED                // Node is locked (reason in string-payload).
	I_PING                  // Ping sent to node, payload incremental hop counter
	I_PONG                  // In return to ping, sent back to sender, payload incremental hop counter
	I_REGISTRATION_REQUEST  // Register request to GW
	I_REGISTRATION_RESPONSE // Register response from GW
	I_DEBUG                 // Debug message

)
const I_REQUEST_SIGNING InternalType = I_SIGNING_PRESENTATION // alias for I_SIGNING_PRESENTATION from version 1.5
const I_GET_NONCE InternalType = I_NONCE_REQUEST              // alias for I_NONCE_REQUEST from version 1.5
const I_GET_NONCE_RESPONSE InternalType = I_NONCE_RESPONSE    // alias for I_NONCE_RESPONSE from version 1.5
