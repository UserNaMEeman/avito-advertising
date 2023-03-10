package handler

import (
	"advertising/entity"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

func (h Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	product := entity.Product{}
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
	err = json.Unmarshal(data, &product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	id, err := h.services.CreateProduct(ctx, product)
	if err != nil {
		if errors.Is(err, entity.ErrLongField) {
			w.WriteHeader(http.StatusRequestEntityTooLarge)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	idStrint := strconv.Itoa(id)
	w.Write([]byte(idStrint))
}
