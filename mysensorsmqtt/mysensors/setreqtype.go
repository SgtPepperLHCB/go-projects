package mysensors

type SetReqType int

const (
	V_TEMP               SetReqType = iota + 0 // S_TEMP, S_HEATER, S_HVAC. Temperature.
	V_HUM                                      // S_HUM. Humidity.
	V_STATUS                                   // S_LIGHT, S_DIMMER, S_SPRINKLER, S_HVAC, S_HEATER. Binary status, 0=off, 1=on.
	V_PERCENTAGE                               // S_DIMMER. Percentage value 0-100 (%).
	V_PRESSURE                                 // S_BARO. Atmospheric Pressure.
	V_FORECAST                                 // S_BARO. Weather forecast. One of "stable", "sunny", "cloudy", "unstable", "thunderstorm" or "unknown".
	V_RAIN                                     // S_RAIN. Amount of rain.
	V_RAINRATE                                 // S_RAIN. Rate of rain.
	V_WIND                                     // S_WIND. Wind speed.
	V_GUST                                     // S_WIND. Gust.
	V_DIRECTION                                // S_WIND. Wind direction 0-360 (degrees).
	V_UV                                       // S_UV. UV light level.
	V_WEIGHT                                   // S_WEIGHT. Weight(for scales etc).
	V_DISTANCE                                 // S_DISTANCE. Distance.
	V_IMPEDANCE                                // S_MULTIMETER, S_WEIGHT. Impedance value.
	V_ARMED                                    // S_DOOR, S_MOTION, S_SMOKE, S_SPRINKLER. Armed status of a security sensor.  1=Armed, 0=Bypassed.
	V_TRIPPED                                  // S_DOOR, S_MOTION, S_SMOKE, S_SPRINKLER, S_WATER_LEAK, S_SOUND, S_VIBRATION, S_MOISTURE. Tripped status of a security sensor. 1=Tripped, 0=Untripped.
	V_WATT                                     // S_POWER, S_LIGHT, S_DIMMER, S_RGB_LIGHT, S_RGBW_LIGHT. Watt value for power meters.
	V_KWH                                      // S_POWER. Accumulated number of KWH for a power meter.
	V_SCENE_ON                                 // S_SCENE_CONTROLLER. Turn on a scene.
	V_SCENE_OFF                                // S_SCENE_CONTROLLER. Turn off a scene.
	V_HVAC_FLOW_STATE                          // S_HEATER, S_HVAC. Mode of heater. One of "Off", "HeatOn", "CoolOn", or "AutoChangeOver"
	V_HVAC_SPEED                               // S_HEATER, S_HVAC. HVAC/Heater fan speed ("Min", "Normal", "Max", "Auto")
	V_LIGHT_LEVEL                              // S_LIGHT_LEVEL. Uncalibrated light level. 0-100%. Use V_LEVEL for light level in lux.
	V_VAR1                                     // Custom value
	V_VAR2                                     // Custom value
	V_VAR3                                     // Custom value
	V_VAR4                                     // Custom value
	V_VAR5                                     // Custom value
	V_UP                                       // S_COVER. Window covering. Up.
	V_DOWN                                     // S_COVER. Window covering. Down.
	V_STOP                                     // S_COVER. Window covering. Stop.
	V_IR_SEND                                  // S_IR. Send out an IR-command.
	V_IR_RECEIVE                               // S_IR. This message contains a received IR-command.
	V_FLOW                                     // S_WATER. Flow of water (in meter).
	V_VOLUME                                   // S_WATER. Water volume.
	V_LOCK_STATUS                              // S_LOCK. Set or get lock status. 1=Locked, 0=Unlocked.
	V_LEVEL                                    // S_DUST, S_AIR_QUALITY, S_SOUND (dB), S_VIBRATION (hz), S_LIGHT_LEVEL (lux).
	V_DUST_LEVEL                               // Dust level
	V_VOLTAGE                                  // S_MULTIMETER. Voltage level.
	V_CURRENT                                  // S_MULTIMETER. Current level.
	V_RGB                                      // S_RGB_LIGHT, S_COLOR_SENSOR. RGB value transmitted as ASCII hex string (I.e "ff0000" for red)
	V_RGBW                                     // S_RGBW_LIGHT. RGBW value transmitted as ASCII hex string (I.e "ff0000ff" for red + full white)
	V_ID                                       // S_TEMP. Optional unique sensor id (e.g. OneWire DS1820b ids)
	V_UNIT_PREFIX                              // S_DUST, S_AIR_QUALITY, S_DISTANCE. Allows sensors to send in a string representing the unit prefix to be displayed in GUI. This is not parsed by controller! E.g. cm, m, km, inch.
	V_HVAC_SETPOINT_COOL                       // S_HVAC. HVAC cool setpoint (Integer between 0-100).
	V_HVAC_SETPOINT_HEAT                       // S_HEATER, S_HVAC. HVAC/Heater setpoint (Integer between 0-100).
	V_HVAC_FLOW_MODE                           // S_HVAC. Flow mode for HVAC ("Auto", "ContinuousOn", "PeriodicOn").
	V_TEXT                                     // S_INFO. Text message to display on LCD or controller device
	V_CUSTOM                                   // S_CUSTOM. Custom messages used for controller/inter node specific commands, preferably using S_CUSTOM device type.
	V_POSITION                                 // S_GPS. GPS position and altitude. Payload: latitude;longitude;altitude(m). E.g. "55.722526;13.017972;18"
	V_IR_RECORD                                // S_IR. Record IR codes for playback
	V_PH                                       // S_WATER_QUALITY, water pH.
	V_ORP                                      // S_WATER_QUALITY, water ORP : redox potential in mV.
	V_EC                                       // S_WATER_QUALITY, water electric conductivity μS/cm (microSiemens/cm).
	V_VAR                                      // S_POWER, Reactive power: volt-ampere reactive (var)
	V_VA                                       // S_POWER, Apparent power: volt-ampere (VA)
	V_POWER_FACTOR                             // S_POWER Ratio of real power to apparent power. Floating point value in the range [-1,..,1]
)
const V_LIGHT SetReqType = V_STATUS      // Deprecated. Alias for V_STATUS. Light Status.0=off 1=on.
const V_DIMMER SetReqType = V_PERCENTAGE // Deprecated. Alias for V_PERCENTAGE. Dimmer value. 0-100 (%).
