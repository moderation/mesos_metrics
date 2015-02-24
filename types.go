// Definititions for mesos slave:5051/monitor/statistics.json endpoint.
// The data is in the form of []Monitor. Collecting the information and
// unmarshal the JSON into []Monitor will get the data into these structs.
package mesos_stats

type Monitor struct {
	ExecutorId   string      `json:"executor_id"`
	ExecutorName string      `json:"executor_name"`
	FrameworkId  string      `json:"framework_id"`
	Source       string      `json:"source"`
	Statistics   *Statistics `json:"statistics"`
}

type Statistics struct {
	CpusUserTimeSecs   float64 `json:"cpus_user_time_secs"`
	CpusSystemTimeSecs float64 `json:"cpus_system_time_secs"`
	CpusLimit          float64 `json:"cpus_limit"`
	MemRssBytes        int64   `json:"mem_rss_bytes"`
	MemLimitBytes      int64   `json:"mem_limit_bytes"`
	Timestamp          float64 `json:"timestamp"`
}
