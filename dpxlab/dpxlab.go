package dpxlab

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/Dr-Deep/dpxlab/dpxlab/config"
	"github.com/Dr-Deep/dpxlab/dpxlab/db"
	"github.com/Dr-Deep/dpxlab/dpxlab/www"
	"github.com/Dr-Deep/logging-go"
)

type Server struct {
	interuptSigs chan os.Signal

	www *www.WebSrv
	db  *db.Database

	cfg    *config.Configuration
	logger *logging.Logger
	sync.Mutex
}

func NewServer(db *db.Database, cfg *config.Configuration, logger *logging.Logger) *Server {
	return &Server{
		interuptSigs: make(chan os.Signal, 1),
		www:          www.NewHttpListener(),
		db:           db,
		cfg:          cfg,
		logger:       logger,
	}
}

func (srv *Server) Start() error {
	srv.Lock()
	/*
	 * db
	 * www
	 * controller
	 */

	signal.Notify(
		srv.interuptSigs,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	// Database

	// WWW
	wwwerr := watch(
		func() error {
			return srv.www.Start(srv.cfg.WWW.Socket)
		},
	)

	// Controller

	// Server
	return srv.run()
}

func (srv *Server) Stop() {
	srv.Lock()
	/*
	   controller
	   www
	   db
	*/

}

func (srv *Server) run() error {
	defer srv.handlePanic()

	for {
		select {
		case any:
			//
		}
	}
}

func (srv *Server) handlePanic() {
	if r := recover(); r != nil {
		srv.Stop()
		srv.logger.Fatal("PANIC", fmt.Sprintf("%#v", r))
	}
}
