package internal

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:               "eod",
		Short:             "get end-of-day market data",
		PersistentPreRunE: runRootCommand,
	}
	logger *log.Logger
)

func init() {
	RootCmd.AddCommand(polygonCmd)
}

func runRootCommand(c *cobra.Command, args []string) error {
	fpath, ok := os.LookupEnv("LOG_FILEPATH")
	if !ok {
		fmt.Println("LOG_FILEPATH not set; logging disabled")
		return nil
	}
	f, err := os.OpenFile(fpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	logger = log.New(
		f,
		"",
		log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC|log.Lshortfile,
	)
	return nil
}
