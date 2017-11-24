// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/watarukura/gody/gody"
)

var scanOption gody.ScanOption

func init() {
	RootCmd.AddCommand(scanCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func scanCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "scan",
		Short: "scan table",
		Run: func(*cobra.Command, []string) {
			gody.Scan(&scanOption)
		},
	}

	options := cmd.Flags()
	options.StringVarP(&scanOption.TableName, "table", "T", "", "DynamoDB table name")
	options.StringVar(&scanOption.Format, "format", "ssv", "Output Format ssv|csv|tsv|json")
	options.BoolVar(&scanOption.Header, "header", false, "With Header")
	options.Int64Var(&scanOption.Limit, "limit", 100, "Output limit")

	return cmd

}
