// Code generated by gtrace. DO NOT EDIT.

package trace

import (
	"context"
)

// Compose returns a new Retry which has functional fields composed
// both from t and x.
func (t Retry) Compose(x Retry) (ret Retry) {
	switch {
	case t.OnRetry == nil:
		ret.OnRetry = x.OnRetry
	case x.OnRetry == nil:
		ret.OnRetry = t.OnRetry
	default:
		h1 := t.OnRetry
		h2 := x.OnRetry
		ret.OnRetry = func(r RetryLoopStartInfo) func(RetryLoopIntermediateInfo) func(RetryLoopDoneInfo) {
			r1 := h1(r)
			r2 := h2(r)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(r RetryLoopIntermediateInfo) func(RetryLoopDoneInfo) {
					r11 := r1(r)
					r21 := r2(r)
					switch {
					case r11 == nil:
						return r21
					case r21 == nil:
						return r11
					default:
						return func(r RetryLoopDoneInfo) {
							r11(r)
							r21(r)
						}
					}
				}
			}
		}
	}
	return ret
}
func (t Retry) onRetry(r RetryLoopStartInfo) func(RetryLoopIntermediateInfo) func(RetryLoopDoneInfo) {
	fn := t.OnRetry
	if fn == nil {
		return func(RetryLoopIntermediateInfo) func(RetryLoopDoneInfo) {
			return func(RetryLoopDoneInfo) {
				return
			}
		}
	}
	res := fn(r)
	if res == nil {
		return func(RetryLoopIntermediateInfo) func(RetryLoopDoneInfo) {
			return func(RetryLoopDoneInfo) {
				return
			}
		}
	}
	return func(r RetryLoopIntermediateInfo) func(RetryLoopDoneInfo) {
		res := res(r)
		if res == nil {
			return func(RetryLoopDoneInfo) {
				return
			}
		}
		return res
	}
}
func RetryOnRetry(t Retry, c *context.Context, iD string, idempotent bool) func(error) func(attempts int, _ error) {
	var p RetryLoopStartInfo
	p.Context = c
	p.ID = iD
	p.Idempotent = idempotent
	res := t.onRetry(p)
	return func(e error) func(int, error) {
		var p RetryLoopIntermediateInfo
		p.Error = e
		res := res(p)
		return func(attempts int, e error) {
			var p RetryLoopDoneInfo
			p.Attempts = attempts
			p.Error = e
			res(p)
		}
	}
}
