package domain

type LoggerPayload struct {
	Time string      `json:"time"`
	Loc  string      `json:"loc"`
	Msg  string      `json:"msg"`
	Req  interface{} `json:"req,omitempty"`
}

type Logger interface {
	Error(payload *LoggerPayload)
	Info(payload *LoggerPayload)
}
