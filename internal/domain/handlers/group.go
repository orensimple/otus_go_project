package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/orensimple/otus_go_project/internal/domain/models"
	"github.com/orensimple/otus_go_project/internal/logger"
)

func (h *Handler) SetGroup(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json; charset=utf-8")
	data := make(map[string]interface{})

	id, err := validateGroup(req)
	if err != nil {
		resp.WriteHeader(400)
		return
	}

	banner, err := h.MainGroupService.SetGroup(req.Context(), id)
	if err == nil {
		data["result"] = banner
	} else {
		data["error"] = err.Error()
	}

	resp.WriteHeader(200)
	json.NewEncoder(resp).Encode(data)
}

func (h *Handler) GetGroups(resp http.ResponseWriter, req *http.Request) {
	data := make(map[string][]*models.Group)
	result, _ := h.MainGroupService.GetGroups(req.Context())
	data["result"] = result
	response, _ := json.Marshal(data)

	resp.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp.WriteHeader(200)
	resp.Write(response)
}

func (h *Handler) UpdateGroup(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json; charset=utf-8")
	data := make(map[string]interface{})

	id, err := validateGroup(req)
	if err != nil {
		resp.WriteHeader(400)
		return
	}

	banner, err := h.MainGroupService.UpdateGroup(req.Context(), id)
	if err == nil {
		data["result"] = banner
	} else {
		data["error"] = err.Error()
	}

	resp.WriteHeader(200)
	json.NewEncoder(resp).Encode(data)
}

func (h *Handler) DeleteGroup(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json; charset=utf-8")
	data := make(map[string]interface{})

	err := req.ParseForm()
	if err != nil {
		resp.WriteHeader(400)
		return
	}
	id, _ := strconv.ParseInt(req.Form.Get("id"), 10, 64)
	if id == 0 {
		resp.WriteHeader(400)
		return
	}
	err = h.MainGroupService.DeleteGroup(req.Context(), id)
	if err == nil {
		data["result"] = "success"
	} else {
		data["error"] = err.Error()
	}

	resp.WriteHeader(200)
	json.NewEncoder(resp).Encode(data)
}

func validateGroup(req *http.Request) (int64, error) {
	err := req.ParseForm()
	if err != nil {
		logger.ContextLogger.Infof("form", "uri", err)
		return 0, err
	}

	id, _ := strconv.ParseInt(req.Form.Get("id"), 10, 64)
	if id == 0 {
		logger.ContextLogger.Infof("id", "uri", err)
		return 0, err
	}

	return id, err
}
