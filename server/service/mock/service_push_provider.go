// Automatically generated by mockimpl. DO NOT EDIT!

package mock

import (
	"sync"

	"github.com/fleetdm/fleet/v4/server/mdm/nanomdm/mdm"
	"github.com/fleetdm/fleet/v4/server/mdm/nanomdm/push"
)

var _ push.PushProvider = (*APNSPushProvider)(nil)

type PushFunc func(p0 []*mdm.Push) (map[string]*push.Response, error)

type APNSPushProvider struct {
	PushFunc        PushFunc
	PushFuncInvoked bool

	mu sync.Mutex
}

func (s *APNSPushProvider) Push(p0 []*mdm.Push) (map[string]*push.Response, error) {
	s.mu.Lock()
	s.PushFuncInvoked = true
	s.mu.Unlock()
	return s.PushFunc(p0)
}
