package support

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	Service *Service
}

func (h *Handler) SendEmail(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(10 << 20)

	uuid := r.FormValue("uuid")
	email := r.FormValue("email")
	header := r.FormValue("header")
	text := r.FormValue("text")

	file, handler, _ := r.FormFile("image")

	var fileReader = file
	var filename string

	if file != nil {
		defer file.Close()
		filename = handler.Filename
	}

	err := h.Service.ProcessTicket(
		uuid,
		email,
		header,
		text,
		fileReader,
		filename,
	)

	if err != nil {
		http.Error(w, "error", 500)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "ticket_created",
	})
}
