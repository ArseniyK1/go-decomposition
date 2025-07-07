package echo

import (
	"encoding/json"
	"net/http"
)

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	var req EchoRequest                         // реализовали структуру реквеста
	err := json.NewDecoder(r.Body).Decode(&req) // Декодируем из байтов body и записываем в структуру req результат

	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(req) // возвращаем {"Text": "ТО_ЧТО_НАПИСАЛ_ПОЛЬЗОВАТЕЛЬ"}
}
