package cloudfunctionsdemo

import "net/http"

func ServeHTTP(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello, world!"))
}
