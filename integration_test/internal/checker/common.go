package checker

import (
	"context"

	"github.com/zsjinwei/openim-tools/log"
)

const (
	errChanLen = 10
)

var (
	checkErrChan = make(chan error, errChanLen)
)

func InsertToErrChan(ctx context.Context, err error) {
	select {
	case checkErrChan <- err:
	default:
		log.ZDebug(ctx, "checkErrChan is full")
	}
}

func CloseAndGetCheckErrChan() <-chan error {
	close(checkErrChan)
	return checkErrChan
}
