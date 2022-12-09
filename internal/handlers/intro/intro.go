package intro
// Welcome page code

import (
	"fmt"
	"net/http"
)

var Intro = func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "intro page")
}