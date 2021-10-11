package server

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"time"

	"github.com/cappyzawa/terraform-registry/config"
	"github.com/cappyzawa/terraform-registry/internal/http"
)

// Opt has options for the server
type Opt struct {
	Port       string
	ConfigPATH string
	PIDPATH    string
	Logger     *log.Logger
}

// Run runs the server
func Run(opt *Opt) error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	return run(ctx, opt)
}

func run(ctx context.Context, opt *Opt) error {
	c, err := config.Parse(opt.ConfigPATH)
	if err != nil {
		return err
	}
	s := http.NewServer(opt.Port, c, opt.Logger)

	if opt.PIDPATH != "" {
		if err := writePIDFile(opt.PIDPATH); err != nil {
			opt.Logger.Printf("failed to write pid file: %v", err)
			return err
		}
		opt.Logger.Printf("wrote pid file: %s", opt.PIDPATH)
	}

	errCh := make(chan error, 1)
	go func() {
		errCh <- s.Start()
	}()

	select {
	case <-ctx.Done():
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := s.Stop(ctx); err != nil {
			return err
		}
		if err := deletePIDFile(opt.PIDPATH); err != nil {
			opt.Logger.Printf("failed to delete pid file: %v", err)
			return err
		}
		opt.Logger.Printf("Deleted PID file: %s", opt.PIDPATH)
		return nil
	case err := <-errCh:
		return err
	}
}

func writePIDFile(path string) error {
	if path == "" {
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return err
	}
	pid := strconv.Itoa(os.Getpid())
	if err := ioutil.WriteFile(path, []byte(pid), 0644); err != nil {
		return err
	}
	return nil
}

func deletePIDFile(path string) error {
	if path == "" {
		return nil
	}
	if err := os.RemoveAll(path); err != nil {
		return err
	}
	return nil
}
