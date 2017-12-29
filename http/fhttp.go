package main

import (
	"net/http"

	//	"log"
	"bufio"
	"fmt"
	"net/url"
)

/*func hello(w http.ResponseWriter, r *http.Request){
	//fmt.Fprint(w, "fsdfsdf", r.URL.Path[1:])
	log.Print(w, "fsdfsdf", nil)

}*/

func main() {
	resp, err := http.Post("http://10.138.32.25/foo.html", "hello", &bufio.Writer{})
	url.Values{"Key": {"Values"}, "id": {123}}
	fmt.Print(resp, err)
	//	http.HandleFunc("/", hello)
	//	http.ListenAndServe("http://10.138.32.25/foo.html", nil)
}
