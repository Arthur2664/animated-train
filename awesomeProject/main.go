package main
import (
	"encoding/json"

	"io/ioutil"
	"log"
	"net/http"
	"os"

)



func main() {

	consoleArg := make([]string , len(os.Args[1:]))
	copy(consoleArg, os.Args[1:])

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}


		//output = append(output, code...)
		switch r.Method {
		case "GET":
			coder , err := json.Marshal(&consoleArg)
			if err != nil {
				w.WriteHeader(500)
				return
			}
			w.Write(coder)
		case "POST":

			reqBody, err := ioutil.ReadAll(r.Body)
			coder , err := json.Marshal(reqBody)


			if err != nil {
				w.WriteHeader(500)

			}
			 w.Write(coder)
			 w.WriteHeader(204)


		default:
			w.WriteHeader(http.StatusNotImplemented)
			w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
		}



	})
	log.Fatal(http.ListenAndServe(":8089", nil))
}
