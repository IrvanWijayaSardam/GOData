package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/IrvanWijayaSardam/GOData/driver"
	models "github.com/IrvanWijayaSardam/GOData/models"
	repository "github.com/IrvanWijayaSardam/GOData/repository"
	post "github.com/IrvanWijayaSardam/GOData/repository/post"
	"github.com/go-chi/chi"
)

func NewPostHandler(db *driver.DB) *Post {
	return &Post{
		repo: post.NewSQLPostRepo(db.SQL),
	}
}

type Post struct {
	repo repository.PostRepo
}

func (h *Post) Fetch(w http.ResponseWriter, r *http.Request) {
	payload, _ := h.repo.Fetch(r.Context(), 5)

	respondwithJSON(w, http.StatusOK, payload)
}

func (h *Post) Create(w http.ResponseWriter, r *http.Request) {
	post := models.Post{}
	json.NewDecoder(r.Body).Decode(&post)

	newID, err := h.repo.Create(r.Context(), &post)
	fmt.Println(newID)

	if err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Server Error")
		return
	}

	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "Succesfully Created"})

}
func (h *Post) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data := models.Post{ID: int64(id)}
	json.NewDecoder(r.Body).Decode(&data)
	payload, err := h.repo.Update(r.Context(), &data)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusOK, payload)

}

func (h *Post) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	payload, err := h.repo.GetByID(r.Context(), int64(id))

	if err != nil {
		respondWithError(w, http.StatusNoContent, "Content Not Found")
	}

	respondwithJSON(w, http.StatusOK, payload)

}
func (h *Post) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	_, err := h.repo.Delete(r.Context(), int64(id))

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusMovedPermanently, map[string]string{"message": "Delete Succesfully"})

}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"Message": msg})
}
