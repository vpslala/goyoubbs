package model

import (
	"sync"
	"sync/atomic"
)

var (
	RateLimitDay     int
	RateLimitHour    int
	BadIpPrefixLst   = NewConStrSlice() // 17.121
	AllowIpPrefixLst = NewConStrSlice() // Manual input, select from all ipInfo list
	BadBotNameMap    atomic.Value       //map[string]interface{}{} // key: string, value: name. DataForSeoBot,SeznamBot,GrapeshotCrawler,
	UserMapMux       = &sync.RWMutex{}
	UserMap          = make(map[uint64]*User) // current user map key: uint64 user id
)

func init() {
	BadBotNameMap.Store(Map{})
}
