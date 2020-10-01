package serve

import (
	"fmt"
	"github.com/Benbentwo/utils/log"
	"net/http"
	"time"
)

func now() string {
	return time.Now().String()
}
func StatusPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{ \"status\": \"ok\" }")
	log.Logger().Infof("%s: Status Page Pinged by: %s", now(), r.Host)
}
