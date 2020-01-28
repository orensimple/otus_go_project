package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/orensimple/otus_go_project/internal/domain/models"
	"github.com/orensimple/otus_go_project/internal/logger"
)

func (h *Handler) SetRotation(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json; charset=utf-8")
	data := make(map[string]interface{})

	banner_id, slot_id, title, err := validateRotation(req)
	if err != nil {
		resp.WriteHeader(400)
		return
	}

	banner, err := h.MainRotationService.SetRotation(req.Context(), banner_id, slot_id, title)
	if err == nil {
		data["result"] = banner
	} else {
		data["error"] = err.Error()
	}

	resp.WriteHeader(200)
	json.NewEncoder(resp).Encode(data)
}

func (h *Handler) GetRotations(resp http.ResponseWriter, req *http.Request) {
	data := make(map[string][]*models.Rotation)
	result, _ := h.MainRotationService.GetRotations(req.Context())
	data["result"] = result
	response, _ := json.Marshal(data)

	resp.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp.WriteHeader(200)
	resp.Write(response)
}

func (h *Handler) UpdateRotation(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json; charset=utf-8")
	data := make(map[string]interface{})

	bannerID, slotID, title, err := validateRotation(req)
	if err != nil {
		resp.WriteHeader(400)
		return
	}

	banner, err := h.MainRotationService.UpdateRotation(req.Context(), bannerID, slotID, title)
	if err == nil {
		data["result"] = banner
	} else {
		data["error"] = err.Error()
	}

	resp.WriteHeader(200)
	json.NewEncoder(resp).Encode(data)
}

func (h *Handler) DeleteRotation(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json; charset=utf-8")
	data := make(map[string]interface{})

	err := req.ParseForm()
	if err != nil {
		resp.WriteHeader(400)
		return
	}
	bannerID, _ := strconv.ParseInt(req.Form.Get("banner_id"), 10, 64)
	if bannerID == 0 {
		resp.WriteHeader(400)
		return
	}
	slotID, _ := strconv.ParseInt(req.Form.Get("slot_id"), 10, 64)
	if slotID == 0 {
		resp.WriteHeader(400)
		return
	}
	err = h.MainRotationService.DeleteRotation(req.Context(), bannerID, slotID)
	if err == nil {
		data["result"] = "success"
	} else {
		data["error"] = err.Error()
	}

	resp.WriteHeader(200)
	json.NewEncoder(resp).Encode(data)
}

func validateRotation(req *http.Request) (int64, int64, string, error) {
	err := req.ParseForm()
	if err != nil {
		logger.ContextLogger.Infof("form", "uri", err)
		return 0, 0, "", err
	}

	bannerID, _ := strconv.ParseInt(req.Form.Get("banner_id"), 10, 64)
	if bannerID == 0 {
		logger.ContextLogger.Infof("bannerID", "uri", err)
		return 0, 0, "", err
	}

	slotID, _ := strconv.ParseInt(req.Form.Get("slot_id"), 10, 64)
	if slotID == 0 {
		logger.ContextLogger.Infof("slotID", "uri", err)
		return 0, 0, "", err
	}

	title := req.Form.Get("title")
	if len(title) == 0 {
		logger.ContextLogger.Infof("title", "uri", err)
		return 0, 0, "", err
	}

	return bannerID, slotID, title, err
}
