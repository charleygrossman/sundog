package internal

import (
	"context"
	"errors"
	"fmt"
	config "marketdata/src/config/polygon"
	"marketdata/src/httpclient"
	"os"
	"time"

	polyrest "github.com/polygon-io/client-go/rest"
	polymodels "github.com/polygon-io/client-go/rest/models"
	"github.com/spf13/cobra"
)

const (
	dateLayout = "2006-01-02"
	// 5 API calls / minute rate limit on free tier.
	rateLimitDuration = 15 * time.Second
)

var (
	polygonCmd = &cobra.Command{
		Use:   "polygon",
		Short: "get end-of-day market data from polygon",
		RunE:  runPolygonCmd,
	}
	configPath string
)

func init() {
	polygonCmd.Flags().StringVarP(&configPath, "config", "c", "", "config filepath")
	polygonCmd.MarkFlagRequired("config")
}

func runPolygonCmd(c *cobra.Command, args []string) error {
	apiKey, ok := os.LookupEnv("POLYGON_API_KEY")
	if !ok {
		return errors.New("must set POLYGON_API_KEY")
	}

	config, err := config.NewEODConfig(configPath)
	if err != nil {
		return err
	}

	date, err := time.Parse(dateLayout, config.Date)
	if err != nil {
		return err
	}

	var opts []httpclient.TransportOptionFunc
	if logger != nil {
		opts = append(opts, httpclient.WithRequestLogger(logger))
	}

	client := polyrest.NewWithClient(
		apiKey,
		httpclient.NewClient(
			httpclient.WithTransport(
				httpclient.NewTransportWithOptions(opts...),
			),
		),
	)

	ctx := context.Background()

	for _, symbol := range config.Symbols {
		params := polymodels.GetDailyOpenCloseAggParams{
			Ticker:   symbol,
			Date:     polymodels.Date(date),
			Adjusted: &config.Adjusted,
		}
		result, err := client.GetDailyOpenCloseAgg(
			ctx,
			&params,
		)
		if err != nil {
			return err
		}
		// TODO: Parse response and persist.
		fmt.Printf("%#v\n", result)
	}

	return nil
}
