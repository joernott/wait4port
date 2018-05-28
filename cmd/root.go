package cmd

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wait4port",
	Short: "Waits until a given port is available",
	Long:  `A small commandline tool to wait until a port is open`,
	Run: func(cmd *cobra.Command, args []string) {
		err := Pinger(Server, Port, time.Duration(Timeout), Retry, time.Duration(Wait), Verbose)
		if err != nil {
			if Verbose {
				s := time.Now().Format("2006-02-01 15:04:05") + err.Error()
				fmt.Println(s)
			}
			os.Exit(1)
		} else {
			if Verbose {
				s := time.Now().Format("2006-02-01 15:04:05") + " Connection established."
				fmt.Println(s)
			}
			os.Exit(0)
		}
	},
}

var Server string
var Port int
var Timeout int64
var Wait int64
var Retry int
var Verbose bool

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&Server, "server", "s", "localhost", "server name (default localhost)")
	rootCmd.PersistentFlags().IntVarP(&Port, "port", "p", 22, "Network port (defaults to 22)")
	rootCmd.PersistentFlags().Int64VarP(&Timeout, "timeout", "t", 30, "Timeout for connection in seconds (default 30)")
	rootCmd.PersistentFlags().Int64VarP(&Wait, "wait", "w", 30, "Seconds to wait between tries (default 30)")
	rootCmd.PersistentFlags().IntVarP(&Retry, "retry", "r", 10, "Number of retries (default 10)")

	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Verbose output (defaults to false)")
}

func Pinger(Server string, Port int, Timeout time.Duration, Retry int, Wait time.Duration, Verbose bool) error {
	for r := 0; r < Retry; r++ {
		if Verbose {
			s := time.Now().Format("2006-02-01 15:04:05") + " Try " + strconv.Itoa(r+1) + "/" + strconv.Itoa(Retry)
			fmt.Println(s)
		}
		err := Ping(Server, Port, Timeout, Verbose)
		if err == nil {
			return nil
		}
		if Verbose {
			s := time.Now().Format("2006-02-01 15:04:05") + " Sleep " + strconv.FormatInt(int64(Wait), 10) + " seconds."
			fmt.Println(s)
		}
		time.Sleep(Wait * time.Second)
	}
	return errors.New(" Failed to connect after " + strconv.Itoa(Retry) + " tries. Giving up.")
}

func Ping(Server string, Port int, Timeout time.Duration, Verbose bool) error {
	h := Server + ":" + strconv.Itoa(Port)
	if Verbose {
		s := time.Now().Format("2006-02-01 15:04:05") + " Ping " + h
		fmt.Println(s)
	}
	conn, err := net.DialTimeout("tcp", h, Timeout*time.Second)
	if Verbose {
		s := time.Now().Format("2006-02-01 15:04:05") + " Ping " + h + " finished "
		if err == nil {
			s = s + "successfully."
		} else {
			s = s + "with error " + err.Error()
		}
		fmt.Println(s)
	}
	if err == nil {
		conn.Close()
	}
	return err
}
