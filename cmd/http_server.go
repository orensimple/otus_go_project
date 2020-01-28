package cmd

import (
	"net/http"
	"strings"
	"time"

	"github.com/orensimple/otus_go_project/config"
	"github.com/orensimple/otus_go_project/internal/domain/handlers"
	"github.com/orensimple/otus_go_project/internal/domain/mw"
	"github.com/orensimple/otus_go_project/internal/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addr string

const TimeOut = 50 * time.Millisecond

// RootCmd init
var RootCmd = &cobra.Command{
	Use:   "http_server",
	Short: "Run http server",
	Run: func(cmd *cobra.Command, args []string) {
		err := config.Init(addr)
		if err != nil {
			logger.ContextLogger.Errorf("Eror init config, viper.ReadInConfig", err.Error())
		}

		logger.InitLogger()

		var serverAddr strings.Builder
		serverAddr.WriteString(viper.GetString("http_listen.ip"))
		serverAddr.WriteString(":")
		serverAddr.WriteString(viper.GetString("http_listen.port"))

		logger.ContextLogger.Infof("Starting http server", viper.GetString("http_listen.ip"), viper.GetString("http_listen.port"))
		router := &handlers.Handler{}
		router.InitDB()
		mux := http.NewServeMux()
		mux.HandleFunc("/banner/set", mw.HTTPLogger(mw.WithTimeout(router.SetBanner, TimeOut)))
		mux.HandleFunc("/banners/get", mw.HTTPLogger(mw.WithTimeout(router.GetBanners, TimeOut)))
		mux.HandleFunc("/banner/update", mw.HTTPLogger(mw.WithTimeout(router.UpdateBanner, TimeOut)))
		mux.HandleFunc("/banner/delete", mw.HTTPLogger(mw.WithTimeout(router.DeleteBanner, TimeOut)))

		mux.HandleFunc("/rotation/set", mw.HTTPLogger(mw.WithTimeout(router.SetRotation, TimeOut)))
		mux.HandleFunc("/rotations/get", mw.HTTPLogger(mw.WithTimeout(router.GetRotations, TimeOut)))
		mux.HandleFunc("/rotation/update", mw.HTTPLogger(mw.WithTimeout(router.UpdateRotation, TimeOut)))
		mux.HandleFunc("/rotation/delete", mw.HTTPLogger(mw.WithTimeout(router.DeleteRotation, TimeOut)))

		mux.HandleFunc("/slot/set", mw.HTTPLogger(mw.WithTimeout(router.SetSlot, TimeOut)))
		mux.HandleFunc("/slots/get", mw.HTTPLogger(mw.WithTimeout(router.GetSlots, TimeOut)))
		mux.HandleFunc("/slot/update", mw.HTTPLogger(mw.WithTimeout(router.UpdateSlot, TimeOut)))
		mux.HandleFunc("/slot/delete", mw.HTTPLogger(mw.WithTimeout(router.DeleteSlot, TimeOut)))

		mux.HandleFunc("/group/set", mw.HTTPLogger(mw.WithTimeout(router.SetGroup, TimeOut)))
		mux.HandleFunc("/groups/get", mw.HTTPLogger(mw.WithTimeout(router.GetGroups, TimeOut)))
		mux.HandleFunc("/group/update", mw.HTTPLogger(mw.WithTimeout(router.UpdateGroup, TimeOut)))
		mux.HandleFunc("/group/delete", mw.HTTPLogger(mw.WithTimeout(router.DeleteGroup, TimeOut)))

		mux.HandleFunc("/report/set", mw.HTTPLogger(mw.WithTimeout(router.SetReport, TimeOut)))
		mux.HandleFunc("/reports/get", mw.HTTPLogger(mw.WithTimeout(router.GetReports, TimeOut)))
		mux.HandleFunc("/report/update", mw.HTTPLogger(mw.WithTimeout(router.UpdateReport, TimeOut)))
		mux.HandleFunc("/report/delete", mw.HTTPLogger(mw.WithTimeout(router.DeleteReport, TimeOut)))

		mux.HandleFunc("/choice/get", mw.HTTPLogger(mw.WithTimeout(router.GetChoice, TimeOut)))

		server := &http.Server{
			Addr:           serverAddr.String(),
			Handler:        mux,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		err = server.ListenAndServe()
		if err != nil {
			logger.ContextLogger.Errorf("Starting http server", err.Error())
		}
	},
}

func init() {
	RootCmd.Flags().StringVar(&addr, "config", "./config", "")
}
