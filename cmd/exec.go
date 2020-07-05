/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"

	"github.com/apex/log"
	"github.com/kodova/posthaste/request"
	"github.com/spf13/cobra"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec",
	Args:  cobra.ExactArgs(1),
	Short: "execute an http request",
	Run: func(cmd *cobra.Command, args []string) {
		c := new(http.Client)

		r, err := request.Open(args[0])
		if err != nil {
			log.Fatalf("failed to open file %v, %v", args[0], err)
		}

		resp, err := r.Execute(c)
		if err != nil {
			log.Fatalf("failed to execture, %+v", err)
		} else {
			printResp(resp)
		}

	},
}

func printResp(r *http.Response) {
	fmt.Println(r.Proto)

	hkeys := make([]string, 0)
	for k, _ := range r.Header {
		hkeys = append(hkeys, k)
	}
	hkeys = sort.StringSlice(hkeys)

	for _, k := range hkeys {
		fmt.Printf("%v: %v\n", k, r.Header.Get(k))
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("failed to decode body, %v", err)
	} else {
		printJson(body)
	}

}

func printJson(body []byte) {
	fmt.Printf("%v", string(body))
}

func init() {
	rootCmd.AddCommand(execCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// execCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// execCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
