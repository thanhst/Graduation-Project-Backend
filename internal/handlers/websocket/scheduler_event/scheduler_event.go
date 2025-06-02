package scheduler_ws

type SchedulerEvent struct {
	Type    string `json:"type"`
	Date    string `json:"date"`
	ClassId string `json:"class_id"`
}
