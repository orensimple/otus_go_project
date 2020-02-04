package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Handler) GetChoice(resp http.ResponseWriter, req *http.Request) {
	data := make(map[string]interface{})

	err := req.ParseForm()
	if err != nil {
		resp.WriteHeader(400)
		return
	}
	group_id, _ := strconv.ParseInt(req.Form.Get("group_id"), 10, 64)
	if group_id == 0 {
		resp.WriteHeader(400)
		return
	}
	slot_id, _ := strconv.ParseInt(req.Form.Get("slot_id"), 10, 64)
	if slot_id == 0 {
		resp.WriteHeader(400)
		return
	}

	result, err := h.MainChoiceService.GetChoice(req.Context(), slot_id, group_id)
	if err != nil {
		resp.WriteHeader(404)
		return
	}
	data["result"] = result
	response, _ := json.Marshal(data)
	resp.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp.WriteHeader(200)
	resp.Write(response)
}

func (h *Handler) ClickChoice(resp http.ResponseWriter, req *http.Request) {
	data := make(map[string]interface{})

	err := req.ParseForm()
	if err != nil {
		resp.WriteHeader(400)
		return
	}
	group_id, _ := strconv.ParseInt(req.Form.Get("group_id"), 10, 64)
	if group_id == 0 {
		resp.WriteHeader(400)
		return
	}
	slot_id, _ := strconv.ParseInt(req.Form.Get("slot_id"), 10, 64)
	if slot_id == 0 {
		resp.WriteHeader(400)
		return
	}
	banner_id, _ := strconv.ParseInt(req.Form.Get("banner_id"), 10, 64)
	if banner_id == 0 {
		resp.WriteHeader(400)
		return
	}

	err = h.MainChoiceService.AddClickToReport(req.Context(), slot_id, group_id, banner_id)
	if err != nil {
		resp.WriteHeader(404)
		return
	}
	data["result"] = "ok"
	response, _ := json.Marshal(data)
	resp.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp.WriteHeader(200)
	resp.Write(response)
}
