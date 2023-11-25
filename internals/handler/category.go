package handler

import (
	"encoding/json"
	"net/http"

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
