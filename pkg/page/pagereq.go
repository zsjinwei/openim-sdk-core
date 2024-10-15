package page

import "github.com/zsjinwei/openim-protocol/sdkws"

type PageReq interface {
	GetPagination() *sdkws.RequestPagination
}
