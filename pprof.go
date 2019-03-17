package echopprof

import (
	"net/http/pprof"

	"github.com/labstack/echo"
)

// Wrap adds several routes from package `net/http/pprof` to *echo.Echo object.
func Wrap(e *echo.Echo) {
	group := e.Group("/debug/pprof")
	group.GET("", index)
	group.GET("/allocs", allocs)
	group.GET("/heap", heap)
	group.GET("/goroutine", goroutine)
	group.GET("/block", block)
	group.GET("/threadcreate", threadCreate)
	group.GET("/cmdline", cmdline)
	group.GET("/profile", profile)
	group.GET("/symbol", symbol)
	group.POST("/symbol", symbol)
	group.GET("/trace", trace)
	group.GET("/mutex", mutex)
}

// IndexHandler will pass the call from /debug/pprof to pprof.
func index(c echo.Context) error {
	pprof.Index(c.Response().Writer, c.Request())
	return nil
}

// HeapHandler will pass the call from /debug/pprof/allocs to pprof.
func allocs(c echo.Context) error {
	pprof.Handler("allocs").ServeHTTP(c.Response(), c.Request())
	return nil
}

// HeapHandler will pass the call from /debug/pprof/heap to pprof.
func heap(c echo.Context) error {
	pprof.Handler("heap").ServeHTTP(c.Response(), c.Request())
	return nil
}

// GoroutineHandler will pass the call from /debug/pprof/goroutine to pprof.
func goroutine(c echo.Context) error {
	pprof.Handler("goroutine").ServeHTTP(c.Response().Writer, c.Request())
	return nil
}

// BlockHandler will pass the call from /debug/pprof/block to pprof.
func block(c echo.Context) error {
	pprof.Handler("block").ServeHTTP(c.Response().Writer, c.Request())
	return nil
}

// ThreadCreateHandler will pass the call from /debug/pprof/threadcreate to pprof.
func threadCreate(c echo.Context) error {
	pprof.Handler("threadcreate").ServeHTTP(c.Response().Writer, c.Request())
	return nil
}

// CmdlineHandler will pass the call from /debug/pprof/cmdline to pprof.
func cmdline(c echo.Context) error {
	pprof.Cmdline(c.Response().Writer, c.Request())
	return nil
}

// ProfileHandler will pass the call from /debug/pprof/profile to pprof.
func profile(c echo.Context) error {
	pprof.Profile(c.Response().Writer, c.Request())
	return nil
}

// SymbolHandler will pass the call from /debug/pprof/symbol to pprof.
func symbol(c echo.Context) error {
	pprof.Symbol(c.Response().Writer, c.Request())
	return nil
}

// TraceHandler will pass the call from /debug/pprof/trace to pprof.
func trace(c echo.Context) error {
	pprof.Trace(c.Response().Writer, c.Request())
	return nil
}

// MutexHandler will pass the call from /debug/pprof/mutex to pprof.
func mutex(c echo.Context) error {
	pprof.Handler("mutex").ServeHTTP(c.Response().Writer, c.Request())
	return nil
}
