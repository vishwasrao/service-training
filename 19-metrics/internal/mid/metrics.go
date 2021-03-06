package mid

import (
	"expvar"
	"net/http"
	"runtime"

	"github.com/ardanlabs/garagesale/internal/platform/web"
)

// m contains the global program counters for the application.
var m = struct {
	gr  *expvar.Int
	req *expvar.Int
	err *expvar.Int
}{
	gr:  expvar.NewInt("goroutines"),
	req: expvar.NewInt("requests"),
	err: expvar.NewInt("errors"),
}

// Metrics updates program counters.
func Metrics(before web.Handler) web.Handler {

	h := func(w http.ResponseWriter, r *http.Request) error {

		err := before(w, r)

		// Increment the request counter.
		m.req.Add(1)

		// Update the count for the number of active goroutines every 100 requests.
		if m.req.Value()%100 == 0 {
			m.gr.Set(int64(runtime.NumGoroutine()))
		}

		// Increment the errors counter if an error occured on this reuqest.
		if err != nil {
			m.err.Add(1)
		}

		// Return the error so it can be handled further up the chain.
		return err
	}

	return h
}
