package app

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Aloe-Corporation/logs"
	"github.com/FloRichardAloeCorp/vfs/server/internal/config"
	"github.com/FloRichardAloeCorp/vfs/server/internal/features/directory"
	"github.com/FloRichardAloeCorp/vfs/server/internal/features/file"
	"github.com/FloRichardAloeCorp/vfs/server/internal/interfaces/datasources"
	ginhttp "github.com/FloRichardAloeCorp/vfs/server/internal/interfaces/http"
	"go.uber.org/zap"
)

var (
	log = logs.Get()
)

type RunCallback func()
type CloseCallback func() error

func Launch(config config.Config) (RunCallback, CloseCallback, error) {
	// Instanciate datasources
	vfs := datasources.NewVFS()

	router := ginhttp.NewRouter(config.Router)

	// Instanciate features
	fileFeature, err := file.NewFileFeatures("vfs", vfs)
	if err != nil {
		return nil, nil, err
	}
	fileHanlder := ginhttp.NewFileHandler(fileFeature)
	fileHanlder.RegisterRoutes(router)

	directoryFeature, err := directory.NewDirectoryFeatures("vfs", vfs)
	if err != nil {
		return nil, nil, err
	}
	directoryHanlder := ginhttp.NewDirectoryHandler(directoryFeature)
	directoryHanlder.RegisterRoutes(router)

	addrGin := config.Router.Addr + ":" + strconv.Itoa(config.Router.Port)
	srv := &http.Server{
		ReadHeaderTimeout: time.Millisecond,
		Addr:              addrGin,
		Handler:           router,
	}

	close := func() error {
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Router.ShutdownTimeout)*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			return fmt.Errorf("can't shutdown server: %w", err)
		}

		return nil
	}

	run := func() {
		log.Info("REST API listening on : "+addrGin,
			zap.String("package", "main"))

		log.Error(router.Run(addrGin).Error(),
			zap.String("package", "main"))
	}

	return run, close, nil
}
