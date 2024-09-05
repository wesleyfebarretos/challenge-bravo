package types

type ILogService interface {
	Info(map[string]any, string)
}
