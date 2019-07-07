package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo/server/application/usecase"
	"todo/server/domain/repository"
)

type ItemHandler struct {
	repo repository.Item
}

func NewItemHandler(repo repository.Item) ItemHandler {
	return ItemHandler{
		repo: repo,
	}
}

func sendResponse(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(data)
}

func (h ItemHandler) Index(w http.ResponseWriter, r *http.Request) {
	uc := usecase.NewItemUseCase(h.repo)
	itemList := uc.GetItemList()

	res, err := json.Marshal(itemList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendResponse(w, res)
}

func (h ItemHandler) Create(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	form := r.PostForm
	title := form["title"][0]
	desc := form["desc"][0]

	uc := usecase.NewItemUseCase(h.repo)
	created := uc.Create(title, desc)

	res, err := json.Marshal(created)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendResponse(w, res)
}

func (h ItemHandler) Close(w http.ResponseWriter, r *http.Request) {
	itemId, _ := strconv.Atoi(r.URL.Query().Get("itemId"))

	uc := usecase.NewItemUseCase(h.repo)
	item := uc.Close(itemId)

	res, err := json.Marshal(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendResponse(w, res)
}

func (h ItemHandler) Open(w http.ResponseWriter, r *http.Request) {
	itemId, _ := strconv.Atoi(r.URL.Query().Get("itemId"))

	uc := usecase.NewItemUseCase(h.repo)
	item := uc.Open(itemId)

	res, err := json.Marshal(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendResponse(w, res)
}
