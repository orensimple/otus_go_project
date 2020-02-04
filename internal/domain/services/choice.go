package services

import (
	"context"
	"math"
	"strconv"
	"time"

	"github.com/orensimple/otus_go_project/internal/domain/errors"
	"github.com/orensimple/otus_go_project/internal/domain/interfaces"
	"github.com/orensimple/otus_go_project/internal/domain/models"
	"github.com/orensimple/otus_go_project/internal/logger"
	"github.com/spf13/viper"
)

type ChoiceService struct {
	MemReportStorage   interfaces.MemReportStorage
	MemRotationStorage interfaces.MemRotationStorage
	ReportQueue        interfaces.ReportQueue
}

func (es *ChoiceService) GetChoice(ctx context.Context, slotID, groupID int64) (int64, error) {
	bannersIDInSlot, err := es.MemRotationStorage.GetRotations(ctx, slotID, groupID)
	if err != nil {
		logger.ContextLogger.Infof("Can't find in rotations list", err.Error())
		return 0, errors.ErrReportNotFound
	}
	logger.ContextLogger.Infof("bannersIDInSlot: ", bannersIDInSlot)
	reports, err := es.MemReportStorage.GetReports(ctx)

	var allShowInSLotGroup int64
	for _, bannerIDInSlot := range bannersIDInSlot {
		bannerStat, ok := reports[slotID][groupID][bannerIDInSlot]
		if !ok {
			es.MemReportStorage.AddShowToReport(ctx, slotID, groupID, bannerIDInSlot)
			return bannerIDInSlot, nil

		}
		allShowInSLotGroup = allShowInSLotGroup + bannerStat.Show
	}

	var responseBannerID int64
	maxRatio := -1.0

	for bannerID, bannerStat := range reports[slotID][groupID] {
		currentRation := float64(bannerStat.Conversion)/float64(bannerStat.Show) + math.Sqrt(math.Log(float64(allShowInSLotGroup))/(2*float64(bannerStat.Show)))
		logger.ContextLogger.Infof("in range: ", bannerID, bannerStat, currentRation)
		if currentRation > maxRatio {
			maxRatio = currentRation
			responseBannerID = bannerID
		}
	}
	reportShowType, err := strconv.ParseInt(viper.GetString("report_type.show"), 10, 64)
	if err != nil {
		logger.ContextLogger.Infof("Can't convert report_type.show to int64", err.Error())
	}
	reportQueueFormat := models.ReportQueueFormat{reportShowType, responseBannerID, groupID, slotID, time.Now()}
	es.ReportQueue.PublicReport(ctx, reportQueueFormat)
	es.MemReportStorage.AddShowToReport(ctx, slotID, groupID, responseBannerID)
	return responseBannerID, nil
}

func (es *ChoiceService) AddClickToReport(ctx context.Context, slotID, groupID, bannerID int64) error {
	es.MemReportStorage.AddClickToReport(ctx, slotID, groupID, bannerID)
	reportShowType, err := strconv.ParseInt(viper.GetString("report_type.conversion"), 10, 64)
	if err != nil {
		logger.ContextLogger.Infof("Can't convert report_type.conversion to int64", err.Error())
	}
	reportQueueFormat := models.ReportQueueFormat{reportShowType, bannerID, groupID, slotID, time.Now()}
	es.ReportQueue.PublicReport(ctx, reportQueueFormat)
	return nil
}
