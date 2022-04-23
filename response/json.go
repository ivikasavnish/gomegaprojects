package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JSON(w http.ResponseWriter, statuscode int, data interface{}) {
	w.WriteHeader(statuscode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}

}
func ERROR(w http.ResponseWriter, statuscode int, err error) {

}
