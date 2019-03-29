package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/rfsx0829/security-code/server/tool"
)

func Check(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	deal(w, err)

	var x struct {
		Id    string `json:"id"`
		Token string `json:"token"`
	}

	err = json.Unmarshal(data, &x)
	deal(w, err)

	token, err := tool.GenerateCode(x.Id, time.Now())
	deal(w, err)

	if token == x.Token {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Good Boy !"))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Token !"))
	}
}

func deal(w http.ResponseWriter, err error) {
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
}
