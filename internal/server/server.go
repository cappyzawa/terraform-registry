package server

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"

	"github.com/cappyzawa/terraform-registry/internal/config"
	"github.com/cappyzawa/terraform-registry/internal/http"
)

// Opt has options for the server
type Opt struct {
	ConfigPATH string
	PIDPATH    string
}

// Run runs the server
func Run(opt *Opt) {
	os.Exit(run(context.Background(), opt))
}

func run(ctx context.Context, opt *Opt) int {
	termCh := make(chan os.Signal, 1)
	signal.Notify(termCh, syscall.SIGTERM, syscall.SIGINT)

	config, _ := config.Parse(opt.ConfigPATH)
	s := http.NewServer(8080, config)
	errCh := make(chan error, 1)

	if err := writePIDFile(opt.PIDPATH); err != nil {
		return 1
	}
	go func() {
		errCh <- s.Start()
	}()

	select {
	case <-termCh:
		if err := s.Stop(ctx); err != nil {
			return 1
		}
		if err := deletePIDFile(opt.PIDPATH); err != nil {
			return 1
		}
		return 0
	case <-errCh:
		return 1
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
	log.Printf("Wrote PID file: %s", path)
	return nil
}

func deletePIDFile(path string) error {
	if path == "" {
		return nil
	}
	if err := os.RemoveAll(path); err != nil {
		return err
	}
	log.Printf("Deleted PID file: %s", path)
	return nil
}
