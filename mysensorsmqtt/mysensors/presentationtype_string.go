// Code generated by "stringer -type PresentationType presentationtype.go"; DO NOT EDIT.

package mysensors

import "strconv"

const _PresentationType_name = "S_DOORS_MOTIONS_SMOKES_LIGHTS_DIMMERS_COVERS_TEMPS_HUMS_BAROS_WINDS_RAINS_UVS_WEIGHTS_POWERS_HEATERS_DISTANCES_LIGHT_LEVELS_ARDUINO_NODES_ARDUINO_REPEATER_NODES_LOCKS_IRS_WATERS_AIR_QUALITYS_CUSTOMS_DUSTS_SCENE_CONTROLLERS_RGB_LIGHTS_RGBW_LIGHTS_COLOR_SENSORS_HVACS_MULTIMETERS_SPRINKLERS_WATER_LEAKS_SOUNDS_VIBRATIONS_MOISTURES_INFOS_GASS_GPSS_WATER_QUALITY"

var _PresentationType_index = [...]uint16{0, 6, 14, 21, 28, 36, 43, 49, 54, 60, 66, 72, 76, 84, 91, 99, 109, 122, 136, 159, 165, 169, 176, 189, 197, 203, 221, 232, 244, 258, 264, 276, 287, 299, 306, 317, 327, 333, 338, 343, 358}

func (i PresentationType) String() string {
	if i < 0 || i >= PresentationType(len(_PresentationType_index)-1) {
		return "PresentationType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _PresentationType_name[_PresentationType_index[i]:_PresentationType_index[i+1]]
}