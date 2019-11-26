/*
 * MinIO Client (C) 2019 MinIO, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import "github.com/minio/cli"

var adminConfigHistoryCmd = cli.Command{
	Name:   "history",
	Usage:  "history restores, clears and lists any previous successful 'mc admin config set' key values.",
	Before: setGlobalsFromContext,
	Action: mainAdminConfigHistory,
	Flags:  globalFlags,
	Subcommands: []cli.Command{
		adminConfigHistoryRestoreCmd,
		adminConfigHistoryListCmd,
		adminConfigHistoryClearCmd,
	},
	HideHelpCommand: true,
}

// mainAdminConfigHistory is the handle for "mc admin config history" command.
func mainAdminConfigHistory(ctx *cli.Context) error {
	cli.ShowCommandHelp(ctx, ctx.Args().First())
	return nil
	// Sub-commands like "restore", "clear", "list" have their own main.
}