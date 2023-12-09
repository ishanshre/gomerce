package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/ishanshre/gomerce/internals/helpers"
	"github.com/ishanshre/gomerce/internals/model"
)

func (h *handler) PostCategoryHandler(w http.ResponseWriter, r *http.Request) {
	request_body := new(model.CreateCategory)
	if err := json.NewDecoder(r.Body).Decode(&request_body); err != nil {
		helpers.StatusBadRequest(w, err.Error())
		return
	}
	category, err := h.repo.CreateCategory(request_body.Name)
	if err != nil {
		helpers.StatusInternalServerError(w, err.Error())
		return
	}
	helpers.StatusOk(w, category)
}

func (h *handler) GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	categories, err := h.repo.GetCategories()
	if err != nil {
		helpers.StatusInternalServerError(w, err.Error())
		return
	}
	helpers.StatusOk(w, categories)
}

func (h *handler) GetCategoryHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		helpers.StatusBadRequest(w, err.Error())
		return
	}
	category, err := h.repo.GetCategory(id)
	if err == sql.ErrNoRows {
		helpers.StatusNoContent(w)
		return
	}
	if err != nil {
		helpers.StatusInternalServerError(w, err.Error())
		return
	}
	helpers.StatusOk(w, category)
}

func (h *handler) DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		helpers.StatusBadRequest(w, err.Error())
		return
	}
	if err := h.repo.DeleteCategory(id); err != nil {
		helpers.StatusInternalServerError(w, err.Error())
		return
	}
	helpers.StatusOk(w, "category deleted")
}

func (h *handler) UpdateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		helpers.StatusBadRequest(w, err.Error())
		return
	}
	new_data := model.Category{}
	if err := json.NewDecoder(r.Body).Decode(&new_data); err != nil {
		helpers.StatusBadRequest(w, fmt.Sprintf("error in parsing request body: %s", err.Error()))
		return
	}
	new_data.Id = id
	category, err := h.repo.UpdateCategory(&new_data)
	if err != nil {
		helpers.StatusInternalServerError(w, err.Error())
		return
	}
	helpers.StatusOk(w, category)

}
