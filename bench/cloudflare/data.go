package bench

import json "encoding/json"

type Log struct {
	EdgeStartTimestamp *json.RawMessage `json:"EdgeStartTimestamp"`
	EventTimestampMs   int64            `json:"EventTimestampMs"`
	Datetime           json.RawMessage  `json:"Datetime"`
	RayID              string           `json:"RayID"`
	ParentRayID        string           `json:"ParentRayID"`
	ClientRequestPath  string           `json:"ClientRequestPath"`
	EdgeResponseStatus int64            `json:"EdgeResponseStatus|AnotherField"`
	Event              struct {
		RPCMethod string `json:"RPCMethod"`
		RayID     string `json:"RayID"`
		Request   struct {
			URL string `json:"URL"`
		} `json:"Request"`
		Response struct {
			Status int64 `json:"Status"`
		} `json:"Response"`
	} `json:"Event"`
	Logs []struct {
		Level       string   `json:"Level"`
		Message     []string `json:"Message"`
		TimestampMs int64    `json:"TimestampMs"`
	} `json:"Logs"`
	Exceptions []struct {
		Name        string `json:"Name"`
		Message     string `json:"Message"`
		TimestampMs int64  `json:"TimestampMs"`
	} `json:"Exceptions"`
	Action              string `json:"Action"`
	SecurityAction      string `json:"SecurityAction"`
	ClientRequestHost   string `json:"ClientRequestHost"`
	ClientRequestMethod string `json:"ClientRequestMethod"`
	ClientRequestSource string `json:"ClientRequestSource"`
	RuleID              string `json:"RuleID"`
	ClientCountry       string `json:"ClientCountry"`
	EdgeColoCode        string `json:"EdgeColoCode"`
	Outcome             string `json:"Outcome"`
	EventType           string `json:"EventType"`
	ScriptName          string `json:"ScriptName"`
	ScriptVersion       struct {
		ID string `json:"ID"`
	} `json:"ScriptVersion"`
	WallTimeMs int64 `json:"WallTimeMs"`
	CPUTimeMs  int64 `json:"CPUTimeMs"`

	RequestHeaders struct {
		Traceparent string `json:"traceparent"`
	} `json:"RequestHeaders"`

	WorkerCPUTime         int64 `json:"WorkerCPUTime"`
	WorkerWallTimeUs      int64 `json:"WorkerWallTimeUs"`
	EdgeTimeToFirstByteMs int64 `json:"EdgeTimeToFirstByteMs"`
	CacheResponseBytes    int64 `json:"CacheResponseBytes"`
	EdgeResponseBytes     int64 `json:"EdgeResponseBytes"`
	OriginResponseBytes   int64 `json:"OriginResponseBytes"`
}
