package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"worktime_server/internal/service"
)

type Credentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (c Credentials) String() string {

	s := strings.Builder{}

	s.WriteString("\n--------------------\n")
	s.WriteString(fmt.Sprintf("login   : %s\n", c.Login))
	s.WriteString(fmt.Sprintf("password: %s\n", c.Password))
	s.WriteString("--------------------\n")

	return s.String()
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

	fmt.Println("endpoint :: Login")

	fmt.Printf("url: %s\n", r.URL)

	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("error read body: %v\n", err)
		return
	}

	var cred Credentials
	if err := json.Unmarshal(data, &cred); err != nil {
		fmt.Printf("error convert body to json: %v\n", err)
		return
	}

	if err := h.services.Login(cred.Login); err != nil {
		w.WriteHeader(service.ToHTTP(err))
		log.Printf("failed login: %v\n", err)
		return
	}

	fmt.Println(cred)
}
