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

var (
	httpHost     string
	httpPort     string
	readTimeout  int
	writeTimeout int
	jwtSecret    string
)

var StartCmd = &cobra.Command{
	Use:          "server",
	Short:        "Start http server",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return startServer()
	},
}

func InitStartCmd() {
	StartCmd.Flags().StringVarP(&httpHost, "host", "l", config.DefaultConfig.HttpServer.Host,
		"Http server listening address")
	config.Viper.BindPFlag("http_server.host", StartCmd.Flags().Lookup("host"))
	config.Viper.BindEnv("http_server.host", "HTTP_HOST")

	StartCmd.Flags().StringVarP(&httpPort, "port", "p", config.DefaultConfig.HttpServer.Port,
		"Http server listening port")
	config.Viper.BindPFlag("http_server.port", StartCmd.Flags().Lookup("port"))
	config.Viper.BindEnv("http_server.port", "HTTP_PORT")

	StartCmd.Flags().IntVarP(&readTimeout, "read-timeout", "", config.DefaultConfig.HttpServer.ReadTimeout,
		"Http server read timeout")
	config.Viper.BindPFlag("http_server.read_timeout", StartCmd.Flags().Lookup("read-timeout"))
	config.Viper.BindEnv("http_server.read_timeout", "HTTP_READ_TIMEOUT")

	StartCmd.Flags().IntVarP(&writeTimeout, "write-timeout", "", config.DefaultConfig.HttpServer.WriteTimeout,
		"Http server write timeout")
	config.Viper.BindPFlag("http_server.write_timeout", StartCmd.Flags().Lookup("write-timeout"))
	config.Viper.BindEnv("http_server.write_timeout", "HTTP_WRITE_TIMEOUT")

	StartCmd.Flags().StringVarP(&jwtSecret, "jwt-secret", "", config.DefaultConfig.Jwt.Secret,
		"Secret key for jwt")
	config.Viper.BindPFlag("jwt.secret", StartCmd.Flags().Lookup("jwt-secret"))
	config.Viper.BindEnv("jwt.secret", "JWT_SECRET")
}

func startServer() (err error) {
	log.Infow("start http server", "BaseVersion", version.BaseVersion, "GitVersion", version.GitVersion)

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

			log.Infow("start http server", "Host", config.Config.HttpServer.Host, "Port", config.Config.HttpServer.Port)
			err = srv.ListenAndServe()
			if err != nil {
				if errors.Is(err, http.ErrServerClosed) {
					err = nil
				} else {
					log.Errorf("failed to start http server:\n%+v", err)
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
					log.Errorf("failed to close http server:\n%+v", err)
				}
				log.Infow("http server exit")
			}
		}(ctx)
	}

	<-signalCh
	cancel()
	exitWg.Wait()
	log.Infow("go-proj-boot exit")

	return
}
