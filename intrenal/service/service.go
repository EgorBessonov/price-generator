//Package service represents price generator logic
package service

import (
	"context"
	"time"

	"github.com/EgorBessonov/price-generator/intrenal/generator"
	"github.com/EgorBessonov/price-generator/intrenal/producer"
	"github.com/sirupsen/logrus"
)

const generationTime = 50

//Service struct
type Service struct {
	pr        *producer.Producer
	shareList *generator.ShareList
}

//NewService returns new service instance
func NewService(pr *producer.Producer, shareList *generator.ShareList) *Service {
	return &Service{
		pr:        pr,
		shareList: shareList,
	}
}

//StartPriceGenerator start price generation
func (s *Service) StartPriceGenerator(ctx context.Context) {
	ticker := time.NewTicker(time.Millisecond * generationTime)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				s.shareList.GeneratePrices()
				err := s.pr.SendPricesToStream(s.shareList)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"error": err,
					}).Error("error while sending message")
					return
				}
			}
		}
	}()
}
