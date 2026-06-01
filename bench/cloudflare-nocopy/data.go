package bench

import json "encoding/json"

type Log struct {
	EdgeStartTimestamp json.RawMessage `json:"EdgeStartTimestamp,nocopy"`
	EventTimestampMs   int64           `json:"EventTimestampMs"`
	Datetime           json.RawMessage `json:"Datetime,nocopy"`
	RayID              string          `json:"RayID,nocopy"`
	ParentRayID        string          `json:"ParentRayID,nocopy"`
	ClientRequestPath  string          `json:"ClientRequestPath,nocopy"`
	EdgeResponseStatus int64           `json:"EdgeResponseStatus"`
	Event              struct {
		RPCMethod string `json:"RPCMethod,nocopy"`
		RayID     string `json:"RayID,nocopy"`
		Request   struct {
			URL string `json:"URL,nocopy"`
		} `json:"Request"`
		Response struct {
			Status int64 `json:"Status"`
		} `json:"Response"`
	} `json:"Event"`
	Logs []struct {
		Level       string   `json:"Level,nocopy"`
		Message     []string `json:"Message,nocopy"`
		TimestampMs int64    `json:"TimestampMs"`
	} `json:"Logs"`
	Exceptions []struct {
		Name        string `json:"Name,nocopy"`
		Message     string `json:"Message,nocopy"`
		TimestampMs int64  `json:"TimestampMs"`
	} `json:"Exceptions"`
	Action              string `json:"Action,nocopy"`
	SecurityAction      string `json:"SecurityAction,nocopy"`
	ClientRequestHost   string `json:"ClientRequestHost,nocopy"`
	ClientRequestMethod string `json:"ClientRequestMethod,nocopy"`
	ClientRequestSource string `json:"ClientRequestSource,nocopy"`
	RuleID              string `json:"RuleID,nocopy"`
	ClientCountry       string `json:"ClientCountry,nocopy"`
	EdgeColoCode        string `json:"EdgeColoCode,nocopy"`
	Outcome             string `json:"Outcome,nocopy"`
	EventType           string `json:"EventType,nocopy"`
	ScriptName          string `json:"ScriptName,nocopy"`
	ScriptVersion       struct {
		ID string `json:"ID,nocopy"`
	} `json:"ScriptVersion"`
	WallTimeMs int64 `json:"WallTimeMs"`
	CPUTimeMs  int64 `json:"CPUTimeMs"`

	RequestHeaders struct {
		Traceparent string `json:"traceparent,nocopy"`
	} `json:"RequestHeaders"`

	WorkerCPUTime         int64 `json:"WorkerCPUTime"`
	WorkerWallTimeUs      int64 `json:"WorkerWallTimeUs"`
	EdgeTimeToFirstByteMs int64 `json:"EdgeTimeToFirstByteMs"`
	CacheResponseBytes    int64 `json:"CacheResponseBytes"`
	EdgeResponseBytes     int64 `json:"EdgeResponseBytes"`
	OriginResponseBytes   int64 `json:"OriginResponseBytes"`
}
