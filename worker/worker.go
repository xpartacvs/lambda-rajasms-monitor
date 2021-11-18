package worker

import (
	"lambda-rajasms-monitor/config"
	"lambda-rajasms-monitor/logger"
	"lambda-rajasms-monitor/webhook"
	"sync"
	"time"

	"github.com/xpartacvs/go-rajasms"

	"github.com/aws/aws-lambda-go/lambda"
)

var (
	once   sync.Once
	client *rajasms.Client
)

func Start() {
	lambda.Start(checkAccount)
}

func getClient() *rajasms.Client {
	once.Do(func() {
		var err error
		client, err = rajasms.NewCient(config.Get().RajaSMSApiURL(), config.Get().RajaSMSApiKey())
		if err != nil {
			logger.Log().Fatal().Msg("Unable to create RajaSMS Client")
		}
	})
	return client
}

func checkAccount() error {
	i, err := getClient().AccountInfo()
	if err != nil {
		logger.Log().Err(err).Msg("Cannot get account inquiry result")
		return err
	}

	if (i.Balance <= config.Get().RajaSMSLowBalance()) || (uint(time.Until(i.Expiry).Hours()/24) <= config.Get().RajaSMSGraceDays()) {
		if err := webhook.GetInstance().AddReminder(
			uint(config.Get().RajaSMSLowBalance()),
			uint(i.Balance),
			config.Get().RajaSMSGraceDays(),
			i.Expiry,
		).Send(config.Get().DishookURL()); err != nil {
			logger.Log().Err(err).Msg("Error while sending alert to discord channel")
		}
	}

	return nil
}
