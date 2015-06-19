// Definititions for mesos slave:5051/metrics/snapshot endpoint.
// The data is in the form of []Metrics. Collecting the information and
// unmarshal the JSON into []Metrics will get the data into these structs.
package prometheus_mesos_metrics

type Metrics struct {
	SlaveCpusPercent              float32 `json:"slave/cpus_percent"`
	SlaveCpousTotal               uint16  `json:"slave/cpus_total"`
	SlaveCpusUsed                 float32 `json:"slave/cpus_used"`
	SlaveDiskPercent              float32 `json:"slave/disk_percent"`
	SlaveDiskTotal                uint32  `json:"slave/disk_total"`
	SlaveDiskUsed                 uint32  `json:"slave/disk_used"`
	SlaveExecutorsRegistering     uint16  `json:"slave/executors_registering"`
	SlaveExecutorsRunning         uint16  `json:"slave/executors_running"`
	SlaveExecutorsTerminated      uint16  `json:"slave/executors_terminated"`
	SlaveExecutorsTerminating     uint16  `json:"slave/executors_terminating"`
	SlaveFrameworksActive         uint16  `json:"slave/frameworks_active"`
	SlaveInvalidFrameworkMessages uint16  `json:"slave/invalid_framework_messages"`
	SlaveInvalidStatusUpdates     uint16  `json:"slave/invalid_status_updates"`
	SlaveMemPercent               float32 `json:"slave/mem_percent"`
	SlaveMemTotal                 uint32  `json:"slave/mem_total"`
	SlaveMemUsed                  uint32  `json:"slave/mem_used"`
	SlaveRecoveryErrors           uint16  `json:"slave/recovery_errors"`
	SlaveRegistered               uint16  `json:"slave/registered"`
	SlaveTasksFailed              uint16  `json:"slave/tasks_failed"`
	SlaveTasksFinished            uint16  `json:"slave/tasks_finished"`
	SlaveTasksKilled              uint16  `json:"slave/tasks_killed"`
	SlaveTasksLost                uint16  `json:"slave/tasks_lost"`
	SlaveTasksRunning             uint16  `json:"slave/tasks_running"`
	SlaveTasksStaging             uint16  `json:"slave/tasks_staging"`
	SlaveTasksStarting            uint16  `json:"slave/tasks_starting"`
	SlaveUptimeSecs               float32 `json:"slave/uptime_secs"`
	SlaveValidFrameworkMessages   uint16  `json:"slave/valid_framework_messages"`
	SlaveValidStatusUpdates       uint16  `json:"slave/valid_status_updates"`
	SystemCpusTotal               uint16  `json:"system/cpus_total"`
	SystemLoad15min               float32 `json:"system/load_15min"`
	SystemLoad1min                float32 `json:"system/load_1min"`
	SystemLoad5min                float32 `json:"system/load_5min"`
	SystemMemFreeBytes            uint64  `json:"system/mem_free_bytes"`
	SystemMemTotalBytes           uint64  `json:"system/mem_total_bytes"`
}
