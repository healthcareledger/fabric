/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package clilogging

import (
	"github.com/hyperledger/fabric/protos/peer"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

func revertLevelsCmd(cf *LoggingCmdFactory) *cobra.Command {
	var loggingRevertLevelsCmd = &cobra.Command{
		Use:   "revertlevels",
		Short: "Reverts the logging levels to the levels at the end of peer startup.",
		Long:  `Reverts the logging levels to the levels at the end of peer startup`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return revertLevels(cf, cmd, args)
		},
	}
	return loggingRevertLevelsCmd
}

func revertLevels(cf *LoggingCmdFactory, cmd *cobra.Command, args []string) (err error) {
	err = checkLoggingCmdParams(cmd, args)
	if err == nil {
		if cf == nil {
			cf, err = InitCmdFactory()
			if err != nil {
				return err
			}
		}
		env := cf.wrapWithEnvelope(&peer.AdminOperation{})
		_, err = cf.AdminClient.RevertLogLevels(context.Background(), env)
		if err != nil {
			return err
		}
		logger.Info("Log levels reverted to the levels at the end of peer startup.")
	}
	return err
}
