package state

import "github.com/FactomProject/factomd/telemetry"

var gauge = telemetry.RegisterMetric.Gauge
var counter = telemetry.RegisterMetric.Counter
var gaugeVec = telemetry.RegisterMetric.GaugeVec

var (
	// Entry Syncing Controller
	HighestKnown = gauge(
		"factomd_state_highest_known",
		"Highest known block (which can be different than the highest ack)",
	)
	HighestSaved = gauge(
		"factomd_state_highest_saved",
		"Highest saved block to the database",
	)
	HighestCompleted = gauge(
		"factomd_state_highest_completed",
		"Highest completed block, which may or may not be saved to the database",
	)

	// TPS
	TotalTransactionPerSecond = gauge(
		"factomd_state_txrate_total_tps",
		"Total transactions over life of node",
	)

	InstantTransactionPerSecond = gauge(
		"factomd_state_txrate_instant_tps",
		"Total transactions over life of node weighted for last 3 seconds",
	)

	// Torrent
	StateTorrentSyncingLower = gauge(
		"factomd_state_torrentsync_lower_gauge",
		"The lower limit of torrent sync",
	)

	StateTorrentSyncingUpper = gauge(
		"factomd_state_torrentsync_upper_gauge",
		"The upper limit of torrent sync",
	)

	// Queues
	CurrentMessageQueueInMsgGeneralVec = gaugeVec(
		"factomd_state_queue_current_general_inmsg_vec",
		"Instrumenting the current  inmsg queue ",
		[]string{"message"},
	)

	TotalMessageQueueInMsgGeneralVec = gaugeVec(
		"factomd_state_queue_total_general_inmsg_vec",
		"Instrumenting the inmsg queue ",
		[]string{"message"},
	)

	CurrentMessageQueueApiGeneralVec = gaugeVec(
		"factomd_state_queue_current_general_api_vec",
		"Instrumenting the current API queue ",
		[]string{"message"},
	)

	TotalMessageQueueApiGeneralVec = gaugeVec(
		"factomd_state_queue_total_general_api_vec",
		"Instrumenting the API queue ",
		[]string{"message"},
	)

	TotalMessageQueueNetOutMsgGeneralVec = gaugeVec(
		"factomd_state_queue_total_general_netoutmsg_vec",
		"Instrumenting the netoutmsg queue ",
		[]string{"message"},
	)

	// Holding Queue
	TotalHoldingQueueInputs = counter(
		"factomd_state_holding_queue_total_inputs",
		"Tally of total messages gone into Holding (useful for rating)",
	)
	TotalHoldingQueueOutputs = counter(
		"factomd_state_holding_queue_total_outputs",
		"Tally of total messages drained out of Holding (useful for rating)",
	)
	HoldingQueueDBSigOutputs = counter(
		"factomd_state_holding_queue_dbsig_outputs",
		"Tally of DBSig messages drained out of Holding",
	)

	// Acks Queue
	TotalAcksInputs = counter(
		"factomd_state_acks_total_inputs",
		"Tally of total messages gone into Acks (useful for rating)",
	)
	TotalAcksOutputs = counter(
		"factomd_state_acks_total_outputs",
		"Tally of total messages drained out of Acks (useful for rating)",
	)

	// Commits map
	TotalCommitsOutputs = counter(
		"factomd_state_commits_total_outputs",
		"Tally of total messages drained out of Commits (useful for rating)",
	)

	// XReview Queue
	TotalXReviewQueueInputs = counter(
		"factomd_state_xreview_queue_total_inputs",
		"Tally of total messages gone into XReview (useful for rating)",
	)

	// Executions
	LeaderExecutions = counter(
		"factomd_state_leader_executions",
		"Tally of total messages executed via LeaderExecute",
	)
	FollowerExecutions = counter(
		"factomd_state_follower_executions",
		"Tally of total messages executed via FollowerExecute",
	)
	LeaderEOMExecutions = counter(
		"factomd_state_leader_eom_executions",
		"Tally of total messages executed via LeaderExecuteEOM",
	)
	FollowerEOMExecutions = counter(
		"factomd_state_follower_eom_executions",
		"Tally of total messages executed via FollowerExecuteEOM",
	)

	// ProcessList
	TotalProcessListInputs = counter(
		"factomd_state_process_list_inputs",
		"Tally of total messages gone into ProcessLists (useful for rating)",
	)
	TotalProcessListProcesses = counter(
		"factomd_state_process_list_processes",
		"Tally of total messages processed from ProcessLists (useful for rating)",
	)
	TotalProcessEOMs = counter(
		"factomd_state_process_eom_processes",
		"Tally of EOM messages processed from ProcessLists (useful for rating)",
	)

	// Durations
	TotalReviewHoldingTime = counter(
		"factomd_state_review_holding_time",
		"Time spent in ReviewHolding()",
	)
	TotalProcessXReviewTime = counter(
		"factomd_state_process_xreview_time",
		"Time spent Processing XReview",
	)
	TotalProcessProcChanTime = counter(
		"factomd_state_process_proc_chan_time",
		"Time spent Processing Process Chan",
	)
	TotalEmptyLoopTime = counter(
		"factomd_state_empty_loop_time",
		"Time spent in empty loop",
	)
	TotalExecuteMsgTime = counter(
		"factomd_state_execute_msg_time",
		"Time spent in executeMsg",
	)
)

var registered bool = false

// RegisterPrometheus registers the variables to be exposed. This can only be run once, hence the
// boolean flag to prevent panics if launched more than once. This is called in NetStart
func RegisterPrometheus() {
	if registered {
		return
	}
	registered = true
	// 		Example Cont.
	// prometheus.MustRegister(stateRandomCounter)

	// Entry syncing
	prometheus.MustRegister(ESAsking)
	prometheus.MustRegister(ESHighestAsking)
	prometheus.MustRegister(ESFirstMissing)
	prometheus.MustRegister(ESMissing)
	prometheus.MustRegister(ESFound)
	prometheus.MustRegister(ESDBHTComplete)
	prometheus.MustRegister(ESMissingQueue)
	prometheus.MustRegister(ESHighestMissing)
	prometheus.MustRegister(ESAvgRequests)
	prometheus.MustRegister(HighestAck)
	prometheus.MustRegister(HighestKnown)
	prometheus.MustRegister(HighestSaved)
	prometheus.MustRegister(HighestCompleted)

	// TPS
	prometheus.MustRegister(TotalTransactionPerSecond)
	prometheus.MustRegister(InstantTransactionPerSecond)

	// Torrent
	prometheus.MustRegister(stateTorrentSyncingLower)
	prometheus.MustRegister(stateTorrentSyncingUpper)

	// Queues
	prometheus.MustRegister(CurrentMessageQueueInMsgGeneralVec)
	prometheus.MustRegister(TotalMessageQueueInMsgGeneralVec)
	prometheus.MustRegister(CurrentMessageQueueApiGeneralVec)
	prometheus.MustRegister(TotalMessageQueueApiGeneralVec)
	prometheus.MustRegister(TotalMessageQueueNetOutMsgGeneralVec)

	// MsgQueue chan
	prometheus.MustRegister(TotalMsgQueueInputs)
	prometheus.MustRegister(TotalMsgQueueOutputs)

	// Holding
	prometheus.MustRegister(TotalHoldingQueueInputs)
	prometheus.MustRegister(TotalHoldingQueueOutputs)
	prometheus.MustRegister(HoldingQueueDBSigInputs)
	prometheus.MustRegister(HoldingQueueDBSigOutputs)
	prometheus.MustRegister(HoldingQueueCommitEntryInputs)
	prometheus.MustRegister(HoldingQueueCommitEntryOutputs)
	prometheus.MustRegister(HoldingQueueCommitChainInputs)
	prometheus.MustRegister(HoldingQueueCommitChainOutputs)
	prometheus.MustRegister(HoldingQueueRevealEntryInputs)
	prometheus.MustRegister(HoldingQueueRevealEntryOutputs)

	// Acks
	prometheus.MustRegister(TotalAcksInputs)
	prometheus.MustRegister(TotalAcksOutputs)

	// Execution
	prometheus.MustRegister(LeaderExecutions)
	prometheus.MustRegister(FollowerExecutions)
	prometheus.MustRegister(LeaderEOMExecutions)
	prometheus.MustRegister(FollowerEOMExecutions)
	prometheus.MustRegister(FollowerMissingMsgExecutions)

	// ProcessList
	prometheus.MustRegister(TotalProcessListInputs)
	prometheus.MustRegister(TotalProcessListProcesses)
	prometheus.MustRegister(TotalProcessEOMs)

	// XReview Queue
	prometheus.MustRegister(TotalXReviewQueueInputs)
	prometheus.MustRegister(TotalXReviewQueueOutputs)

	// Commits map
	prometheus.MustRegister(TotalCommitsInputs)
	prometheus.MustRegister(TotalCommitsOutputs)

	// Durations
	prometheus.MustRegister(TotalReviewHoldingTime)
	prometheus.MustRegister(TotalProcessXReviewTime)
	prometheus.MustRegister(TotalProcessProcChanTime)
	prometheus.MustRegister(TotalEmptyLoopTime)
	prometheus.MustRegister(TotalAckLoopTime)
	prometheus.MustRegister(TotalExecuteMsgTime)
}
