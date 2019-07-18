package mysensors

type PresentationType int

const (
	S_DOOR                  PresentationType = iota + 0 // Door and window sensors
	S_MOTION                                            // Motion sensors
	S_SMOKE                                             // Smoke sensor
	S_LIGHT                                             // Light Actuator (on/off)
	S_DIMMER                                            // Dimmable device of some kind
	S_COVER                                             // Window covers or shades
	S_TEMP                                              // Temperature sensor
	S_HUM                                               // Humidity sensor
	S_BARO                                              // Barometer sensor (Pressure)
	S_WIND                                              // Wind sensor
	S_RAIN                                              // Rain sensor
	S_UV                                                // UV sensor
	S_WEIGHT                                            // Weight sensor for scales etc.
	S_POWER                                             // Power measuring device, like power meters
	S_HEATER                                            // Heater device
	S_DISTANCE                                          // Distance sensor
	S_LIGHT_LEVEL                                       // Light sensor
	S_ARDUINO_NODE                                      // Arduino node device
	S_ARDUINO_REPEATER_NODE                             // Arduino repeating node device
	S_LOCK                                              // Lock device
	S_IR                                                // Ir sender/receiver device
	S_WATER                                             // Water meter
	S_AIR_QUALITY                                       // Air quality sensor e.g. MQ-2
	S_CUSTOM                                            // Use this for custom sensors
	S_DUST                                              // Dust level sensor
	S_SCENE_CONTROLLER                                  // Scene controller device
	S_RGB_LIGHT                                         // RGB light
	S_RGBW_LIGHT                                        // RGBW light (with separate white component)
	S_COLOR_SENSOR                                      // Color sensor
	S_HVAC                                              // Thermostat/HVAC device
	S_MULTIMETER                                        // Multimeter device
	S_SPRINKLER                                         // Sprinkler device
	S_WATER_LEAK                                        // Water leak sensor
	S_SOUND                                             // Sound sensor
	S_VIBRATION                                         // Vibration sensor
	S_MOISTURE                                          // Moisture sensor
	S_INFO                                              // LCD text device / Simple information device on controller, V_TEXT
	S_GAS                                               // Gas meter, V_FLOW, V_VOLUME
	S_GPS                                               // GPS Sensor, V_POSITION
	S_WATER_QUALITY                                     // V_TEMP, V_PH, V_ORP, V_EC, V_STATUS
)
const S_BINARY PresentationType = S_LIGHT                        // Binary device (on/off), Alias for S_LIGHT
const S_ARDUINO_RELAY PresentationType = S_ARDUINO_REPEATER_NODE // Alias for compatability
