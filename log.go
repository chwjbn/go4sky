package go4sky

import "context"

type LogData struct {
	LogCtx context.Context
	LogLevel string
	LogContent string
}
