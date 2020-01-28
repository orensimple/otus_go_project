package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/orensimple/otus_go_project/internal/domain/models"
	"github.com/orensimple/otus_go_project/internal/logger"
)

// CreateReport
func (h *Handler) SetReport(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json; charset=utf-8")
	data := make(map[string]interface{})

	banner_id, group_id, slot_id, show, conversion, err := validateReport(req)
	if err != nil {
		resp.WriteHeader(400)
		return
	}

	report, err := h.MainReportService.SetReport(req.Context(), banner_id, group_id, slot_id, show, conversion)
	if err == nil {
		data["result"] = report
	} else {
		data["error"] = err.Error()
	}

	resp.WriteHeader(200)
	json.NewEncoder(resp).Encode(data)
}

func (h *Handler) GetReports(resp http.ResponseWriter, req *http.Request) {
	data := make(map[string][]*models.Report)
	result, _ := h.MainReportService.GetReports(req.Context())
	data["result"] = result
	response, _ := json.Marshal(data)

	resp.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp.WriteHeader(200)
	resp.Write(response)
}

func (h *Handler) UpdateReport(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json; charset=utf-8")
	data := make(map[string]interface{})

	banner_id, group_id, slot_id, show, conversion, err := validateReport(req)
	if err != nil {
		resp.WriteHeader(400)
		return
	}

	report, err := h.MainReportService.UpdateReport(req.Context(), banner_id, group_id, slot_id, show, conversion)
	if err == nil {
		data["result"] = report
	} else {
		data["error"] = err.Error()
	}

	resp.WriteHeader(200)
	json.NewEncoder(resp).Encode(data)
}

func (h *Handler) DeleteReport(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json; charset=utf-8")
	data := make(map[string]interface{})

	err := req.ParseForm()
	if err != nil {
		resp.WriteHeader(400)
		return
	}
	banner_id, _ := strconv.ParseInt(req.Form.Get("banner_id"), 10, 64)
	if banner_id == 0 {
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
	err = h.MainReportService.DeleteReport(req.Context(), banner_id, group_id, slot_id)
	if err == nil {
		data["result"] = "success"
	} else {
		data["error"] = err.Error()
	}

	resp.WriteHeader(200)
	json.NewEncoder(resp).Encode(data)
}

func validateReport(req *http.Request) (int64, int64, int64, int64, int64, error) {
	err := req.ParseForm()
	if err != nil {
		logger.ContextLogger.Infof("form", "uri", err)
		return 0, 0, 0, 0, 0, err
	}

	banner_id, _ := strconv.ParseInt(req.Form.Get("banner_id"), 10, 64)
	if banner_id == 0 {
		logger.ContextLogger.Infof("banner_id", "uri", err)
		return 0, 0, 0, 0, 0, err
	}

	group_id, _ := strconv.ParseInt(req.Form.Get("group_id"), 10, 64)
	if group_id == 0 {
		logger.ContextLogger.Infof("group_id", "uri", err)
		return 0, 0, 0, 0, 0, err
	}

	slot_id, _ := strconv.ParseInt(req.Form.Get("slot_id"), 10, 64)
	if slot_id == 0 {
		logger.ContextLogger.Infof("slot_id", "uri", err)
		return 0, 0, 0, 0, 0, err
	}

	show, _ := strconv.ParseInt(req.Form.Get("show"), 10, 64)
	if show == 0 {
		logger.ContextLogger.Infof("show", "uri", err)
		return 0, 0, 0, 0, 0, err
	}

	conversion, _ := strconv.ParseInt(req.Form.Get("conversion"), 10, 64)
	if conversion == 0 {
		logger.ContextLogger.Infof("conversion", "uri", err)
		return 0, 0, 0, 0, 0, err
	}

	return banner_id, group_id, slot_id, show, conversion, err
}
