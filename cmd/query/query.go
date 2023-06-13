/*
Copyright © 2023 OpenFGA

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
package query

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// QueryCmd represents the query command.
var QueryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query a store.",
}

func init() {
	QueryCmd.AddCommand(checkCmd)
	QueryCmd.AddCommand(expandCmd)
	QueryCmd.AddCommand(listObjectsCmd)
	QueryCmd.AddCommand(listRelationsCmd)

	QueryCmd.PersistentFlags().String("store-id", "", "Store ID")
	QueryCmd.Flags().String("store-id", "", "Store ID")
	err := QueryCmd.MarkFlagRequired("store-id")
	if err != nil { //nolint:wsl
		fmt.Print(err)
		os.Exit(1)
	}
}