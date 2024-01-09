package main

import (
	"fmt"
	"time"

	"github.com/opsgenie/opsgenie-go-sdk/alertsv2"
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
)

func getNewAlerts(c *Config) ([]alertsv2.Alert, error) {
	client := ogcli.OpsGenieClient{}
	client.SetAPIKey(c.OpsgenieToken)
	client.SetOpsGenieAPIUrl(c.OpsgenieAPIURL)
	alertClient, _ := client.AlertV2()
	query := makeQuery(c)
	alerts := make([]alertsv2.Alert, 0)
	for offset := 0; offset < 9800; offset += 100 {
		request := alertsv2.ListAlertRequest{
			Query:  query,
			Offset: offset,
			Limit:  100,
			Sort:   "createdAt",
			Order:  "desc",
		}
		response, err := alertClient.List(request)
		if err != nil {
			return nil, err
		}
		alerts = append(alerts, response.Alerts...)
		if len(response.Alerts) < 100 {
			break
		}
	}
	return alerts, nil
}

func makeQuery(c *Config) string {
	start := time.Now().Add(-c.MinAge).UnixMilli()
	q := fmt.Sprintf("createdAt < %d %s", start, c.Filter)
	return q
}
