package ping

import (
	"encoding/json"
	"net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	res := Ping{Message: "pong"}
	json.NewEncoder(w).Encode(res) // возвращаем {"Message": "pong"}
}
