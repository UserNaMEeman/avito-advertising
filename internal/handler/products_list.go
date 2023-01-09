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

func (h Handler) ReturnProducts(w http.ResponseWriter, r *http.Request) {
	sortProduct := entity.SortProduct{}
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	rawBody := r.Body
	defer r.Body.Close()
	body, err := io.ReadAll(rawBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	if err = json.Unmarshal(body, &sortProduct); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	// fmt.Println(sortProduct)
	pageProducts, err := h.services.GetProducts(ctx, sortProduct)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	for i, products := range pageProducts {
		page := "page " + strconv.Itoa(i+1) + ":\n"
		w.Write([]byte(page))
		for _, product := range products {
			report := "		" + product.Title + "	" + product.URL1 + " " + product.Price + "\n"
			w.Write([]byte(report))
		}
	}
}
