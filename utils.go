package ghostly

import (
	"fmt"
	"regexp"
	"runtime"
	"time"
)

func (g *Ghostly) LoadTime(start time.Time) {
	elapsed := time.Since(start)
	pc, _, _, _ := runtime.Caller(1)
	funcObj := runtime.FuncForPC(pc)
	runtimeFunc := regexp.MustCompile(`^.*\.(.*)$`)
	name := runtimeFunc.ReplaceAllString(funcObj.Name(), "$1")
	g.InfoLog.Printf(fmt.Sprintf("Load time: %s took %s", name, elapsed))
}
