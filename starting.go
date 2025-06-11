package main
import (
	"fmt"
	"net/http" //this is the library for supporting servers
)
func main(){
	server := &http.Server{ //this is how you write a server here
		Addr : ":8080", //port
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Ok")
	})
	server.ListenAndServe() //and this is how you 
}