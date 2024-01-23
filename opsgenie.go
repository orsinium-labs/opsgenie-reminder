package main

import (
	"context"
	"fmt"
	"time"

	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
	ogcli "github.com/opsgenie/opsgenie-go-sdk-v2/client"
)

func getNewAlerts(c Config) ([]alert.Alert, error) {
	config := &ogcli.Config{
		ApiKey:         c.OpsgenieToken,
		OpsGenieAPIURL: ogcli.ApiUrl(c.OpsgenieAPIURL),
	}
	alertClient, _ := alert.NewClient(config)
	query := makeQuery(c)
	alerts := make([]alert.Alert, 0)
	for offset := 0; offset < 9800; offset += 100 {
		// https://docs.opsgenie.com/docs/alert-api#list-alerts
		request := alert.ListAlertRequest{
			Query:  query,
			Offset: offset,
			Limit:  100,
			Sort:   "createdAt",
			Order:  "desc",
		}
		ctx := context.Background()
		response, err := alertClient.List(ctx, &request)
		if err != nil {
			return nil, err
		}
		alerts = append(alerts, response.Alerts...)
		if len(response.Alerts) < 100 {
			break
		}
		offset += len(response.Alerts)
	}
	return alerts, nil
}

// https://support.atlassian.com/opsgenie/docs/search-queries-for-alerts/
func makeQuery(c Config) string {
	start := time.Now().Add(-c.MinAge).UnixMilli()
	q := fmt.Sprintf("status: open createdAt < %d %s", start, c.Filter)
	return q
}
