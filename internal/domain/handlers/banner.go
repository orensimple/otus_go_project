package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/orensimple/otus_go_project/internal/domain/models"
	"github.com/orensimple/otus_go_project/internal/domain/services"
	"github.com/orensimple/otus_go_project/internal/logger"
	"github.com/orensimple/otus_go_project/internal/maindb"
	"github.com/spf13/viper"
)

type Handler struct {
	Handlers              *http.Handler
	MainBannerService     *services.BannerService
	MaindbBannerStorage   *maindb.PgBannerStorage
	MainGroupService      *services.GroupService
	MaindbGroupStorage    *maindb.PgGroupStorage
	MainReportService     *services.ReportService
	MaindbReportStorage   *maindb.PgReportStorage
	MainRotationService   *services.RotationService
	MaindbRotationStorage *maindb.PgRotationStorage
	MainSlotService       *services.SlotService
	MaindbSlotStorage     *maindb.PgSlotStorage
	MainChoiceService     *services.ChoiceService
}

func (h *Handler) SetBanner(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json; charset=utf-8")
	data := make(map[string]interface{})

	id, title, err := validateBanner(req)
	if err != nil {
		resp.WriteHeader(400)
		return
	}

	banner, err := h.MainBannerService.SetBanner(req.Context(), id, title)
	if err == nil {
		data["result"] = banner
	} else {
		data["error"] = err.Error()
	}

	resp.WriteHeader(200)
	json.NewEncoder(resp).Encode(data)
}

func (h *Handler) GetBanners(resp http.ResponseWriter, req *http.Request) {
	data := make(map[string][]*models.Banner)
	result, _ := h.MainBannerService.GetBanners(req.Context())
	data["result"] = result
	response, _ := json.Marshal(data)

	resp.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp.WriteHeader(200)
	resp.Write(response)
}

func (h *Handler) UpdateBanner(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json; charset=utf-8")
	data := make(map[string]interface{})

	id, title, err := validateBanner(req)
	if err != nil {
		resp.WriteHeader(400)
		return
	}

	banner, err := h.MainBannerService.UpdateBanner(req.Context(), id, title)
	if err == nil {
		data["result"] = banner
	} else {
		data["error"] = err.Error()
	}

	resp.WriteHeader(200)
	json.NewEncoder(resp).Encode(data)
}

func (h *Handler) DeleteBanner(resp http.ResponseWriter, req *http.Request) {
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
	err = h.MainBannerService.DeleteBanner(req.Context(), id)
	if err == nil {
		data["result"] = "success"
	} else {
		data["error"] = err.Error()
	}

	resp.WriteHeader(200)
	json.NewEncoder(resp).Encode(data)
}

func (h *Handler) InitDB() {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", viper.GetString("postgres.user"), viper.GetString("postgres.passwd"), viper.GetString("postgres.ip"), viper.GetString("postgres.port"), viper.GetString("postgres.db"))

	var err error
	h.MaindbBannerStorage, err = maindb.NewPgBannerStorage(dsn)
	if err != nil {
		logger.ContextLogger.Infof("Problem connect to db", dsn, err.Error())
	}

	h.MainBannerService = &services.BannerService{
		BannerStorage: h.MaindbBannerStorage,
	}

	h.MaindbReportStorage, err = maindb.NewPgReportStorage(dsn)
	if err != nil {
		logger.ContextLogger.Infof("Problem connect to db", dsn, err.Error())
	}

	h.MainReportService = &services.ReportService{
		ReportStorage: h.MaindbReportStorage,
	}

	h.MaindbRotationStorage, err = maindb.NewPgRotationStorage(dsn)
	if err != nil {
		logger.ContextLogger.Infof("Problem connect to db", dsn, err.Error())
	}

	h.MainRotationService = &services.RotationService{
		RotationStorage: h.MaindbRotationStorage,
	}

	h.MaindbSlotStorage, err = maindb.NewPgSlotStorage(dsn)
	if err != nil {
		logger.ContextLogger.Infof("Problem connect to db", dsn, err.Error())
	}

	h.MainSlotService = &services.SlotService{
		SlotStorage: h.MaindbSlotStorage,
	}

	h.MaindbGroupStorage, err = maindb.NewPgGroupStorage(dsn)
	if err != nil {
		logger.ContextLogger.Infof("Problem connect to db", dsn, err.Error())
	}

	h.MainGroupService = &services.GroupService{
		GroupStorage: h.MaindbGroupStorage,
	}

	h.MainChoiceService = &services.ChoiceService{}
}

func validateBanner(req *http.Request) (int64, string, error) {
	err := req.ParseForm()
	if err != nil {
		logger.ContextLogger.Infof("form", "uri", err)
		return 0, "", err
	}

	id, _ := strconv.ParseInt(req.Form.Get("id"), 10, 64)
	if id == 0 {
		logger.ContextLogger.Infof("id", "uri", err)
		return 0, "", err
	}

	title := req.Form.Get("title")
	if len(title) == 0 {
		logger.ContextLogger.Infof("title", "uri", err)
		return 0, "", err
	}

	return id, title, err
}
