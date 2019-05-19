package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func checkActivity(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error in connection node: %s\nRequest Body: %s", err, body)
	}
	log.Printf("Node %s came here!\n", body)
	return
}

func addFeedback(w http.ResponseWriter, r *http.Request) {
	idEmployee, err := strconv.Atoi(r.FormValue("id_employee"))
	if err != nil {
		fmt.Fprintf(w, "Ошибка в указании пользователя")
		return
	}
	mark, err := strconv.Atoi(r.FormValue("mark"))
	if err != nil {
		fmt.Fprintf(w, "Ошибка в указании оценки")
		return
	}
	feedback := r.FormValue("text")
	err = bc.AddBlock(feedback, idEmployee, mark, 0)
	if err != nil {
		log.Printf("Error in add block to BlockChain: %s\n", err)
		fmt.Fprint(w, "Система отзывов временно недоступна, приносим вам свои извинения")
		return
	}
	fmt.Fprintf(w, "Ваш отзыв принят в обработку\n")
}
