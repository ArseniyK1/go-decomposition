package ping

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	res := Ping{Message: "pong"}
	fmt.Println(res)               // логируем в консоль
	json.NewEncoder(w).Encode(res) // возвращаем {"Message": "pong"}
}
