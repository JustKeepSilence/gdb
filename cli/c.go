/*
createTime: 2021/6/6
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.16
*/

package cli

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/JustKeepSilence/gdb/db"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"os"
)

// cli for gdb, to build: go build -o gdbCli
// run: gdbCli stats -i localhost:8082 -m http

var ip, mode string

var cmd = &cobra.Command{
	Use:   "gdb",
	Short: "gdb is a realTime database written with go",
	Long: `A fast and flexible realTime database in Go.
                Complete documentation is available at https://github.com/JustKeepSilence/gdb`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("inconrrect cli parameter: cli can only contain one parameter")
		}
		if mode != "http" && mode != "https" {
			return fmt.Errorf("incorrect mode, mode can only be http or https")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if res, err := get(args[0]); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
		}
	},
}

func Execute() {
	cmd.Flags().StringVarP(&ip, "ip", "i", "", "used to connect gdb server")
	_ = cmd.MarkFlagRequired("ip")
	cmd.Flags().StringVarP(&mode, "mode", "m", "", "used to specify the connection method of the gdb service, can only be http or https")
	_ = cmd.MarkFlagRequired("mode")
	cmd.SetHelpCommand(cmd)
	cmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		fmt.Println(`gdbCLi a CLI application to get table info of gdb server quickly.
Usage:
  gdbCli [command]

Available Commands:
  stats       			Returns statistics of the underlying DB
  iostats     			Returns statistics of effective disk read and write
  writedelay  			Returns cumulative write delay caused by compaction
  sstables	  			Returns sstables list for each level.
  blockpool	  			Returns block pool stats.
  cachedblock 			Returns size of cached block.
  openedtables 			Returns number of opened tables
  alivesnaps  			Returns number of alive snapshots.
  aliveiters  			Returns number of alive iterators.
  num-files-at-level{n} Returns the number of files at level 'n'.

Flags:
  -h, --help             help for gdbCli
  -i, --ip string        ip address of gdbServer
  -m, --mode string		 used to specify the connection method of the gdb service, can only be http or https`)
	})
	cmd.SetHelpTemplate("")
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func get(p string) (string, error) {
	if mode == "http" {
		if response, err := http.Get("http://" + ip + "/cmd/getCmdInfo/" + p); err != nil {
			return "", err
		} else {
			if b, err := ioutil.ReadAll(response.Body); err != nil {
				return "", err
			} else {
				r := db.ResponseData{}
				if err := json.Unmarshal(b, &r); err != nil {
					return "", err
				}
				if r.Code == 200 {
					return r.Data.(string), nil
				} else {
					return "", fmt.Errorf(r.Message)
				}
			}
		}
	} else {
		client := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{RootCAs: nil, InsecureSkipVerify: true}}}
		if response, err := client.Get("https://" + ip + "/cmd/getCmdInfo/" + p); err != nil {
			return "", err
		} else {
			if b, err := ioutil.ReadAll(response.Body); err != nil {
				return "", err
			} else {
				r := db.ResponseData{}
				if err := json.Unmarshal(b, &r); err != nil {
					return "", err
				}
				if r.Code == 200 {
					return r.Data.(string), nil
				} else {
					return "", fmt.Errorf(r.Message)
				}
			}
		}
	}
}
