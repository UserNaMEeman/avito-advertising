package handler

import (
	"advertising/entity"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

func (h Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := entity.ProductID{}
	fields := entity.ResponseFileds{}
	ctx, cancel := context.WithTimeout(r.Context(), 70*time.Second)
	defer cancel()
	body := r.Body
	defer r.Body.Close()
	data, err := io.ReadAll(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(data, &id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(data, &fields)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	idN, err := strconv.Atoi(id.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	prod, err := h.services.GetProduct(ctx, idN, fields)
	if err != nil {
		// if errors.Is(err, entity.ErrLongField) {
		// 	w.WriteHeader(http.StatusRequestEntityTooLarge)
		// 	return
		// }
		w.WriteHeader(http.StatusInternalServerError)
		// fmt.Println(err)
		return
	}
	resp, err := json.Marshal(prod)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte(resp))
}
