package naas

import (
	"context"
	"fmt"
	"log"
	"sync/atomic"
	"time"

	"github.com/go-resty/resty/v2"
)

const naasURL = "https://naas.isalman.dev/no"

type Service struct {
	enabled atomic.Bool
	client  *resty.Client
}

func NewService() *Service {
	client := resty.New().
		SetTimeout(10*time.Second).
		SetRetryCount(2).
		SetRetryWaitTime(2*time.Second).
		SetHeader("Accept", "application/json")

	s := &Service{
		client: client,
	}
	s.enabled.Store(true)

	return s
}

func (s *Service) Start(ctx context.Context) {
	ticker := time.NewTicker(15 * time.Second)

	go func() {
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				if s.enabled.Load() {
					s.call()
				}
			case <-ctx.Done():
				log.Println("NAAS service stopped")
				return
			}
		}
	}()
}

func (s *Service) Enable()  { s.enabled.Store(true) }
func (s *Service) Disable() { s.enabled.Store(false) }
func (s *Service) IsEnabled() bool {
	return s.enabled.Load()
}

func (s *Service) call() {
	resp, err := s.client.R().Get(naasURL)
	if err != nil {
		log.Printf("NAAS error: %v\n", err)
		return
	}

	fmt.Printf("%d - %s\n", resp.StatusCode(), resp.String())
}
