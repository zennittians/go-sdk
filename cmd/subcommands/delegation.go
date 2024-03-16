package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zennittians/golang-sdk/pkg/rpc"
)

var (
	delegationSubCmds = []*cobra.Command{{
		Use:     "by-delegator",
		Short:   "Get all delegations by a delegator",
		Args:    cobra.ExactArgs(1),
		PreRunE: validateAddress,
		RunE: func(cmd *cobra.Command, args []string) error {
			noLatest = true
			return request(rpc.Method.GetDelegationsByDelegator, []interface{}{addr.address})
		},
	}, {
		Use:     "by-validator",
		Short:   "Get all delegations for a validator",
		Args:    cobra.ExactArgs(1),
		PreRunE: validateAddress,
		RunE: func(cmd *cobra.Command, args []string) error {
			noLatest = true
			return request(rpc.Method.GetDelegationsByValidator, []interface{}{addr.address})
		},
	}}
)
