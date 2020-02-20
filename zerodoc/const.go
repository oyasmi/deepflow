package zerodoc

type MessageType uint8

const (
	MSG_USAGE MessageType = iota
	MSG_PERF
	MSG_GEO
	MSG_FLOW
	_
	MSG_TYPE
	MSG_FPS
	MSG_LOG_USAGE
	MSG_VTAP_USAGE
	_
	MSG_VTAP_SIMPLE

	MSG_INVILID
)

const (
	MAX_STRING_LENGTH = 1024
)

const (
	FLOW_ID uint8 = iota
	FPS_ID
	GEO_ID
	PERF_ID
	TYPE_ID
	USAGE_ID
	CONSOLE_LOG_ID
	LOG_USAGE_ID
	VTAP_USAGE_ID
	VTAP_SIMPLE_ID

	MAX_APP_ID
)

var MeterDFNames [MAX_APP_ID]string = [MAX_APP_ID]string{
	"df_flow",
	"df_fps",
	"df_geo",
	"df_perf",
	"df_type",
	"df_usage",
	"df_console_log",
	"log_usage",
	"vtap_usage",
	"vtap_simple",
}

var MeterVTAPNames [MAX_APP_ID]string = [MAX_APP_ID]string{
	"vtap_flow",
	"vtap_flow_fps",
	"vtap_geo",
	"vtap_flow_perf",
	"vtap_flow_type",
	"vtap_flow_usage",
	"vtap_console_log",
	"log_usage",
	"vtap_usage",
	"vtap_simple",
}

var MeterNamesToID map[string]uint8

func GetMeterID(name string) uint8 {
	if id, exist := MeterNamesToID[name]; exist {
		return id
	}
	log.Errorf("can't get meter(%s) id", name)
	return MAX_APP_ID
}

const (
	MAIN uint8 = iota
	MINI
	MAIN_ISP
	MAIN_REGION
	MAIN_CAST_TYPE
	MAIN_TCP_FLAGS

	MAX_MEASUREMENT_ID
)

var MeasurementNames [MAX_MEASUREMENT_ID]string = [MAX_MEASUREMENT_ID]string{
	"main",
	"mini",
	"main_isp",
	"main_region",
	"main_cast_type",
	"main_tcp_flags",
}

var MeasurementNamesToID map[string]uint8

func GetMeasurementID(name string) uint8 {
	if mid, exist := MeasurementNamesToID[name]; exist {
		return mid
	}
	log.Errorf("can't get measurement(%s) id", name)
	return MAX_MEASUREMENT_ID
}

func init() {
	MeterNamesToID = make(map[string]uint8)
	for id, name := range MeterDFNames {
		MeterNamesToID[name] = uint8(id)
	}
	for id, name := range MeterVTAPNames {
		MeterNamesToID[name] = uint8(id)
	}

	MeasurementNamesToID = make(map[string]uint8)
	for id, name := range MeasurementNames {
		MeasurementNamesToID[name] = uint8(id)
	}
}
