package start_server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/fs714/go-proj-boot/api"
	"github.com/fs714/go-proj-boot/pkg/utils/config"
	"github.com/fs714/go-proj-boot/pkg/utils/log"
	"github.com/fs714/go-proj-boot/pkg/utils/version"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var StartCmd = &cobra.Command{
	Use:          "server",
	Short:        "Start http server",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return startServer()
	},
}

func startServer() (err error) {
	log.Infof("start http server %s, %s, %s", version.BaseVersion, version.GitVersion, version.BuildTime)

	ctx, cancel := context.WithCancel(context.Background())
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	exitWg := &sync.WaitGroup{}

	{
		router := api.InitRouter()
		srv := &http.Server{
			Addr:           fmt.Sprintf("%s:%s", config.Config.HttpServer.Host, config.Config.HttpServer.Port),
			Handler:        router,
			ReadTimeout:    time.Duration(time.Duration(config.Config.HttpServer.ReadTimeout) * time.Second),
			WriteTimeout:   time.Duration(time.Duration(config.Config.HttpServer.WriteTimeout) * time.Second),
			MaxHeaderBytes: 1 << 20,
		}

		exitWg.Add(1)
		go func() {
			defer exitWg.Done()

			log.Infof("start http server on %s:%s", config.Config.HttpServer.Host, config.Config.HttpServer.Port)
			err = srv.ListenAndServe()
			if err != nil {
				if errors.Is(err, http.ErrServerClosed) {
					err = nil
				} else {
					log.Errorf("failed to start http server with err: %s", err.Error())
				}
			}
		}()

		exitWg.Add(1)
		go func(ctx context.Context) {
			defer exitWg.Done()

			select {
			case <-ctx.Done():
				cctx, ccancel := context.WithTimeout(context.Background(), 1*time.Second)
				defer ccancel()
				err := srv.Shutdown(cctx)
				if err != nil {
					log.Errorf("failed to close http server with err: %s", err.Error())
				}
				log.Infof("http server exit")
			}
		}(ctx)
	}

	<-signalCh
	cancel()
	exitWg.Wait()
	log.Infof("go-proj-boot exit")

	return
}
