// Code generated by gtrace. DO NOT EDIT.

package trace

import (
	"context"
)

// Compose returns a new Driver which has functional fields composed
// both from t and x.
func (t Driver) Compose(x Driver) (ret Driver) {
	switch {
	case t.OnConnNew == nil:
		ret.OnConnNew = x.OnConnNew
	case x.OnConnNew == nil:
		ret.OnConnNew = t.OnConnNew
	default:
		h1 := t.OnConnNew
		h2 := x.OnConnNew
		ret.OnConnNew = func(c ConnNewStartInfo) func(ConnNewDoneInfo) {
			r1 := h1(c)
			r2 := h2(c)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(c ConnNewDoneInfo) {
					r1(c)
					r2(c)
				}
			}
		}
	}
	switch {
	case t.OnConnClose == nil:
		ret.OnConnClose = x.OnConnClose
	case x.OnConnClose == nil:
		ret.OnConnClose = t.OnConnClose
	default:
		h1 := t.OnConnClose
		h2 := x.OnConnClose
		ret.OnConnClose = func(c ConnCloseStartInfo) func(ConnCloseDoneInfo) {
			r1 := h1(c)
			r2 := h2(c)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(c ConnCloseDoneInfo) {
					r1(c)
					r2(c)
				}
			}
		}
	}
	switch {
	case t.OnConnDial == nil:
		ret.OnConnDial = x.OnConnDial
	case x.OnConnDial == nil:
		ret.OnConnDial = t.OnConnDial
	default:
		h1 := t.OnConnDial
		h2 := x.OnConnDial
		ret.OnConnDial = func(c ConnDialStartInfo) func(ConnDialDoneInfo) {
			r1 := h1(c)
			r2 := h2(c)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(c ConnDialDoneInfo) {
					r1(c)
					r2(c)
				}
			}
		}
	}
	switch {
	case t.OnConnDisconnect == nil:
		ret.OnConnDisconnect = x.OnConnDisconnect
	case x.OnConnDisconnect == nil:
		ret.OnConnDisconnect = t.OnConnDisconnect
	default:
		h1 := t.OnConnDisconnect
		h2 := x.OnConnDisconnect
		ret.OnConnDisconnect = func(c ConnDisconnectStartInfo) func(ConnDisconnectDoneInfo) {
			r1 := h1(c)
			r2 := h2(c)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(c ConnDisconnectDoneInfo) {
					r1(c)
					r2(c)
				}
			}
		}
	}
	switch {
	case t.OnConnStateChange == nil:
		ret.OnConnStateChange = x.OnConnStateChange
	case x.OnConnStateChange == nil:
		ret.OnConnStateChange = t.OnConnStateChange
	default:
		h1 := t.OnConnStateChange
		h2 := x.OnConnStateChange
		ret.OnConnStateChange = func(c ConnStateChangeStartInfo) func(ConnStateChangeDoneInfo) {
			r1 := h1(c)
			r2 := h2(c)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(c ConnStateChangeDoneInfo) {
					r1(c)
					r2(c)
				}
			}
		}
	}
	switch {
	case t.OnConnInvoke == nil:
		ret.OnConnInvoke = x.OnConnInvoke
	case x.OnConnInvoke == nil:
		ret.OnConnInvoke = t.OnConnInvoke
	default:
		h1 := t.OnConnInvoke
		h2 := x.OnConnInvoke
		ret.OnConnInvoke = func(c ConnInvokeStartInfo) func(ConnInvokeDoneInfo) {
			r1 := h1(c)
			r2 := h2(c)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(c ConnInvokeDoneInfo) {
					r1(c)
					r2(c)
				}
			}
		}
	}
	switch {
	case t.OnConnNewStream == nil:
		ret.OnConnNewStream = x.OnConnNewStream
	case x.OnConnNewStream == nil:
		ret.OnConnNewStream = t.OnConnNewStream
	default:
		h1 := t.OnConnNewStream
		h2 := x.OnConnNewStream
		ret.OnConnNewStream = func(c ConnNewStreamStartInfo) func(ConnNewStreamRecvInfo) func(ConnNewStreamDoneInfo) {
			r1 := h1(c)
			r2 := h2(c)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(c ConnNewStreamRecvInfo) func(ConnNewStreamDoneInfo) {
					r11 := r1(c)
					r21 := r2(c)
					switch {
					case r11 == nil:
						return r21
					case r21 == nil:
						return r11
					default:
						return func(c ConnNewStreamDoneInfo) {
							r11(c)
							r21(c)
						}
					}
				}
			}
		}
	}
	switch {
	case t.OnConnTake == nil:
		ret.OnConnTake = x.OnConnTake
	case x.OnConnTake == nil:
		ret.OnConnTake = t.OnConnTake
	default:
		h1 := t.OnConnTake
		h2 := x.OnConnTake
		ret.OnConnTake = func(c ConnTakeStartInfo) func(ConnTakeDoneInfo) {
			r1 := h1(c)
			r2 := h2(c)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(c ConnTakeDoneInfo) {
					r1(c)
					r2(c)
				}
			}
		}
	}
	switch {
	case t.OnConnRelease == nil:
		ret.OnConnRelease = x.OnConnRelease
	case x.OnConnRelease == nil:
		ret.OnConnRelease = t.OnConnRelease
	default:
		h1 := t.OnConnRelease
		h2 := x.OnConnRelease
		ret.OnConnRelease = func(c ConnReleaseStartInfo) func(ConnReleaseDoneInfo) {
			r1 := h1(c)
			r2 := h2(c)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(c ConnReleaseDoneInfo) {
					r1(c)
					r2(c)
				}
			}
		}
	}
	switch {
	case t.OnClusterGet == nil:
		ret.OnClusterGet = x.OnClusterGet
	case x.OnClusterGet == nil:
		ret.OnClusterGet = t.OnClusterGet
	default:
		h1 := t.OnClusterGet
		h2 := x.OnClusterGet
		ret.OnClusterGet = func(c ClusterGetStartInfo) func(ClusterGetDoneInfo) {
			r1 := h1(c)
			r2 := h2(c)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(c ClusterGetDoneInfo) {
					r1(c)
					r2(c)
				}
			}
		}
	}
	switch {
	case t.OnClusterInsert == nil:
		ret.OnClusterInsert = x.OnClusterInsert
	case x.OnClusterInsert == nil:
		ret.OnClusterInsert = t.OnClusterInsert
	default:
		h1 := t.OnClusterInsert
		h2 := x.OnClusterInsert
		ret.OnClusterInsert = func(c ClusterInsertStartInfo) func(ClusterInsertDoneInfo) {
			r1 := h1(c)
			r2 := h2(c)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(c ClusterInsertDoneInfo) {
					r1(c)
					r2(c)
				}
			}
		}
	}
	switch {
	case t.OnClusterUpdate == nil:
		ret.OnClusterUpdate = x.OnClusterUpdate
	case x.OnClusterUpdate == nil:
		ret.OnClusterUpdate = t.OnClusterUpdate
	default:
		h1 := t.OnClusterUpdate
		h2 := x.OnClusterUpdate
		ret.OnClusterUpdate = func(c ClusterUpdateStartInfo) func(ClusterUpdateDoneInfo) {
			r1 := h1(c)
			r2 := h2(c)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(c ClusterUpdateDoneInfo) {
					r1(c)
					r2(c)
				}
			}
		}
	}
	switch {
	case t.OnClusterRemove == nil:
		ret.OnClusterRemove = x.OnClusterRemove
	case x.OnClusterRemove == nil:
		ret.OnClusterRemove = t.OnClusterRemove
	default:
		h1 := t.OnClusterRemove
		h2 := x.OnClusterRemove
		ret.OnClusterRemove = func(c ClusterRemoveStartInfo) func(ClusterRemoveDoneInfo) {
			r1 := h1(c)
			r2 := h2(c)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(c ClusterRemoveDoneInfo) {
					r1(c)
					r2(c)
				}
			}
		}
	}
	switch {
	case t.OnPessimizeNode == nil:
		ret.OnPessimizeNode = x.OnPessimizeNode
	case x.OnPessimizeNode == nil:
		ret.OnPessimizeNode = t.OnPessimizeNode
	default:
		h1 := t.OnPessimizeNode
		h2 := x.OnPessimizeNode
		ret.OnPessimizeNode = func(p PessimizeNodeStartInfo) func(PessimizeNodeDoneInfo) {
			r1 := h1(p)
			r2 := h2(p)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(p PessimizeNodeDoneInfo) {
					r1(p)
					r2(p)
				}
			}
		}
	}
	switch {
	case t.OnGetCredentials == nil:
		ret.OnGetCredentials = x.OnGetCredentials
	case x.OnGetCredentials == nil:
		ret.OnGetCredentials = t.OnGetCredentials
	default:
		h1 := t.OnGetCredentials
		h2 := x.OnGetCredentials
		ret.OnGetCredentials = func(g GetCredentialsStartInfo) func(GetCredentialsDoneInfo) {
			r1 := h1(g)
			r2 := h2(g)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(g GetCredentialsDoneInfo) {
					r1(g)
					r2(g)
				}
			}
		}
	}
	switch {
	case t.OnDiscovery == nil:
		ret.OnDiscovery = x.OnDiscovery
	case x.OnDiscovery == nil:
		ret.OnDiscovery = t.OnDiscovery
	default:
		h1 := t.OnDiscovery
		h2 := x.OnDiscovery
		ret.OnDiscovery = func(d DiscoveryStartInfo) func(DiscoveryDoneInfo) {
			r1 := h1(d)
			r2 := h2(d)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(d DiscoveryDoneInfo) {
					r1(d)
					r2(d)
				}
			}
		}
	}
	return ret
}
func (t Driver) onConnNew(c1 ConnNewStartInfo) func(ConnNewDoneInfo) {
	fn := t.OnConnNew
	if fn == nil {
		return func(ConnNewDoneInfo) {
			return
		}
	}
	res := fn(c1)
	if res == nil {
		return func(ConnNewDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onConnClose(c1 ConnCloseStartInfo) func(ConnCloseDoneInfo) {
	fn := t.OnConnClose
	if fn == nil {
		return func(ConnCloseDoneInfo) {
			return
		}
	}
	res := fn(c1)
	if res == nil {
		return func(ConnCloseDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onConnDial(c1 ConnDialStartInfo) func(ConnDialDoneInfo) {
	fn := t.OnConnDial
	if fn == nil {
		return func(ConnDialDoneInfo) {
			return
		}
	}
	res := fn(c1)
	if res == nil {
		return func(ConnDialDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onConnDisconnect(c1 ConnDisconnectStartInfo) func(ConnDisconnectDoneInfo) {
	fn := t.OnConnDisconnect
	if fn == nil {
		return func(ConnDisconnectDoneInfo) {
			return
		}
	}
	res := fn(c1)
	if res == nil {
		return func(ConnDisconnectDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onConnStateChange(c1 ConnStateChangeStartInfo) func(ConnStateChangeDoneInfo) {
	fn := t.OnConnStateChange
	if fn == nil {
		return func(ConnStateChangeDoneInfo) {
			return
		}
	}
	res := fn(c1)
	if res == nil {
		return func(ConnStateChangeDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onConnInvoke(c1 ConnInvokeStartInfo) func(ConnInvokeDoneInfo) {
	fn := t.OnConnInvoke
	if fn == nil {
		return func(ConnInvokeDoneInfo) {
			return
		}
	}
	res := fn(c1)
	if res == nil {
		return func(ConnInvokeDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onConnNewStream(c1 ConnNewStreamStartInfo) func(ConnNewStreamRecvInfo) func(ConnNewStreamDoneInfo) {
	fn := t.OnConnNewStream
	if fn == nil {
		return func(ConnNewStreamRecvInfo) func(ConnNewStreamDoneInfo) {
			return func(ConnNewStreamDoneInfo) {
				return
			}
		}
	}
	res := fn(c1)
	if res == nil {
		return func(ConnNewStreamRecvInfo) func(ConnNewStreamDoneInfo) {
			return func(ConnNewStreamDoneInfo) {
				return
			}
		}
	}
	return func(c ConnNewStreamRecvInfo) func(ConnNewStreamDoneInfo) {
		res := res(c)
		if res == nil {
			return func(ConnNewStreamDoneInfo) {
				return
			}
		}
		return res
	}
}
func (t Driver) onConnTake(c1 ConnTakeStartInfo) func(ConnTakeDoneInfo) {
	fn := t.OnConnTake
	if fn == nil {
		return func(ConnTakeDoneInfo) {
			return
		}
	}
	res := fn(c1)
	if res == nil {
		return func(ConnTakeDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onConnRelease(c1 ConnReleaseStartInfo) func(ConnReleaseDoneInfo) {
	fn := t.OnConnRelease
	if fn == nil {
		return func(ConnReleaseDoneInfo) {
			return
		}
	}
	res := fn(c1)
	if res == nil {
		return func(ConnReleaseDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onClusterGet(c1 ClusterGetStartInfo) func(ClusterGetDoneInfo) {
	fn := t.OnClusterGet
	if fn == nil {
		return func(ClusterGetDoneInfo) {
			return
		}
	}
	res := fn(c1)
	if res == nil {
		return func(ClusterGetDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onClusterInsert(c1 ClusterInsertStartInfo) func(ClusterInsertDoneInfo) {
	fn := t.OnClusterInsert
	if fn == nil {
		return func(ClusterInsertDoneInfo) {
			return
		}
	}
	res := fn(c1)
	if res == nil {
		return func(ClusterInsertDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onClusterUpdate(c1 ClusterUpdateStartInfo) func(ClusterUpdateDoneInfo) {
	fn := t.OnClusterUpdate
	if fn == nil {
		return func(ClusterUpdateDoneInfo) {
			return
		}
	}
	res := fn(c1)
	if res == nil {
		return func(ClusterUpdateDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onClusterRemove(c1 ClusterRemoveStartInfo) func(ClusterRemoveDoneInfo) {
	fn := t.OnClusterRemove
	if fn == nil {
		return func(ClusterRemoveDoneInfo) {
			return
		}
	}
	res := fn(c1)
	if res == nil {
		return func(ClusterRemoveDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onPessimizeNode(p PessimizeNodeStartInfo) func(PessimizeNodeDoneInfo) {
	fn := t.OnPessimizeNode
	if fn == nil {
		return func(PessimizeNodeDoneInfo) {
			return
		}
	}
	res := fn(p)
	if res == nil {
		return func(PessimizeNodeDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onGetCredentials(g GetCredentialsStartInfo) func(GetCredentialsDoneInfo) {
	fn := t.OnGetCredentials
	if fn == nil {
		return func(GetCredentialsDoneInfo) {
			return
		}
	}
	res := fn(g)
	if res == nil {
		return func(GetCredentialsDoneInfo) {
			return
		}
	}
	return res
}
func (t Driver) onDiscovery(d DiscoveryStartInfo) func(DiscoveryDoneInfo) {
	fn := t.OnDiscovery
	if fn == nil {
		return func(DiscoveryDoneInfo) {
			return
		}
	}
	res := fn(d)
	if res == nil {
		return func(DiscoveryDoneInfo) {
			return
		}
	}
	return res
}
func DriverOnConnNew(t Driver, c context.Context, address string, l Location) func(state ConnState) {
	var p ConnNewStartInfo
	p.Context = c
	p.Address = address
	p.Location = l
	res := t.onConnNew(p)
	return func(state ConnState) {
		var p ConnNewDoneInfo
		p.State = state
		res(p)
	}
}
func DriverOnConnClose(t Driver, c context.Context, address string, l Location, state ConnState) func() {
	var p ConnCloseStartInfo
	p.Context = c
	p.Address = address
	p.Location = l
	p.State = state
	res := t.onConnClose(p)
	return func() {
		var p ConnCloseDoneInfo
		res(p)
	}
}
func DriverOnConnDial(t Driver, c context.Context, address string, l Location) func(error) {
	var p ConnDialStartInfo
	p.Context = c
	p.Address = address
	p.Location = l
	res := t.onConnDial(p)
	return func(e error) {
		var p ConnDialDoneInfo
		p.Error = e
		res(p)
	}
}
func DriverOnConnDisconnect(t Driver, c context.Context, address string, l Location, state ConnState) func(state ConnState, _ error) {
	var p ConnDisconnectStartInfo
	p.Context = c
	p.Address = address
	p.Location = l
	p.State = state
	res := t.onConnDisconnect(p)
	return func(state ConnState, e error) {
		var p ConnDisconnectDoneInfo
		p.State = state
		p.Error = e
		res(p)
	}
}
func DriverOnConnStateChange(t Driver, c context.Context, address string, l Location, state ConnState) func(state ConnState) {
	var p ConnStateChangeStartInfo
	p.Context = c
	p.Address = address
	p.Location = l
	p.State = state
	res := t.onConnStateChange(p)
	return func(state ConnState) {
		var p ConnStateChangeDoneInfo
		p.State = state
		res(p)
	}
}
func DriverOnConnInvoke(t Driver, c context.Context, address string, l Location, m Method) func(_ error, issues []Issue, opID string, state ConnState) {
	var p ConnInvokeStartInfo
	p.Context = c
	p.Address = address
	p.Location = l
	p.Method = m
	res := t.onConnInvoke(p)
	return func(e error, issues []Issue, opID string, state ConnState) {
		var p ConnInvokeDoneInfo
		p.Error = e
		p.Issues = issues
		p.OpID = opID
		p.State = state
		res(p)
	}
}
func DriverOnConnNewStream(t Driver, c context.Context, address string, l Location, m Method) func(error) func(state ConnState, _ error) {
	var p ConnNewStreamStartInfo
	p.Context = c
	p.Address = address
	p.Location = l
	p.Method = m
	res := t.onConnNewStream(p)
	return func(e error) func(ConnState, error) {
		var p ConnNewStreamRecvInfo
		p.Error = e
		res := res(p)
		return func(state ConnState, e error) {
			var p ConnNewStreamDoneInfo
			p.State = state
			p.Error = e
			res(p)
		}
	}
}
func DriverOnConnTake(t Driver, c context.Context, address string, l Location) func(error) {
	var p ConnTakeStartInfo
	p.Context = c
	p.Address = address
	p.Location = l
	res := t.onConnTake(p)
	return func(e error) {
		var p ConnTakeDoneInfo
		p.Error = e
		res(p)
	}
}
func DriverOnConnRelease(t Driver, c context.Context, address string, l Location) func() {
	var p ConnReleaseStartInfo
	p.Context = c
	p.Address = address
	p.Location = l
	res := t.onConnRelease(p)
	return func() {
		var p ConnReleaseDoneInfo
		res(p)
	}
}
func DriverOnClusterGet(t Driver, c context.Context) func(address string, _ Location, _ error) {
	var p ClusterGetStartInfo
	p.Context = c
	res := t.onClusterGet(p)
	return func(address string, l Location, e error) {
		var p ClusterGetDoneInfo
		p.Address = address
		p.Location = l
		p.Error = e
		res(p)
	}
}
func DriverOnClusterInsert(t Driver, c context.Context, address string, l Location) func(state ConnState, _ Location) {
	var p ClusterInsertStartInfo
	p.Context = c
	p.Address = address
	p.Location = l
	res := t.onClusterInsert(p)
	return func(state ConnState, l Location) {
		var p ClusterInsertDoneInfo
		p.State = state
		p.Location = l
		res(p)
	}
}
func DriverOnClusterUpdate(t Driver, c context.Context, address string) func(state ConnState) {
	var p ClusterUpdateStartInfo
	p.Context = c
	p.Address = address
	res := t.onClusterUpdate(p)
	return func(state ConnState) {
		var p ClusterUpdateDoneInfo
		p.State = state
		res(p)
	}
}
func DriverOnClusterRemove(t Driver, c context.Context, address string, l Location) func(state ConnState) {
	var p ClusterRemoveStartInfo
	p.Context = c
	p.Address = address
	p.Location = l
	res := t.onClusterRemove(p)
	return func(state ConnState) {
		var p ClusterRemoveDoneInfo
		p.State = state
		res(p)
	}
}
func DriverOnPessimizeNode(t Driver, c context.Context, address string, l Location, state ConnState, cause error) func(state ConnState, _ error) {
	var p PessimizeNodeStartInfo
	p.Context = c
	p.Address = address
	p.Location = l
	p.State = state
	p.Cause = cause
	res := t.onPessimizeNode(p)
	return func(state ConnState, e error) {
		var p PessimizeNodeDoneInfo
		p.State = state
		p.Error = e
		res(p)
	}
}
func DriverOnGetCredentials(t Driver, c context.Context) func(tokenOk bool, _ error) {
	var p GetCredentialsStartInfo
	p.Context = c
	res := t.onGetCredentials(p)
	return func(tokenOk bool, e error) {
		var p GetCredentialsDoneInfo
		p.TokenOk = tokenOk
		p.Error = e
		res(p)
	}
}
func DriverOnDiscovery(t Driver, c context.Context) func(endpoints []string, _ error) {
	var p DiscoveryStartInfo
	p.Context = c
	res := t.onDiscovery(p)
	return func(endpoints []string, e error) {
		var p DiscoveryDoneInfo
		p.Endpoints = endpoints
		p.Error = e
		res(p)
	}
}
