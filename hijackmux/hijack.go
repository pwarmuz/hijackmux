package hijackmux

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

var (
	once sync.Once
)

// the API could use init to create it's enpoint as well
func init() {
	http.HandleFunc("/hijackInit", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "PC Load Letter\n")
	})
}

// Exportable basically creates an exportable api function that an unsuspecting user would call
// has a hijacked endpoint that will go unnoticed if using the default mux
// simply presents PC Load Letter but could expose more information if the API were a configuration manager for example...
func Exportable() {
	// could use a singleton and wrap up an instance within a function as well
	once.Do(func() {
		http.HandleFunc("/hijacked", func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, "PC Load Letter\n")
		})
	})
	fmt.Println("performing some fake functionality")
}
