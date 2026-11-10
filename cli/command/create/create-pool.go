package create

import (
	"encoding/json"

	"github.com/Telmate/proxmox-api-go/cli"
	"github.com/Telmate/proxmox-api-go/proxmox"
	"github.com/spf13/cobra"
)

var create_poolCmd = &cobra.Command{
	Use:   "pool POOLID",
	Short: "Creates a new pool",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := cli.RequiredIDset(args, 0, "PoolID")
		var config proxmox.ConfigPool
		if err := json.Unmarshal(cli.NewConfig(), &config); err != nil {
			return err
		}
		config.Name = proxmox.PoolName(id)
		if err := cli.NewClient().New().Pool.Create(cli.Context(), config); err != nil {
			return err
		}
		cli.PrintItemCreated(CreateCmd.OutOrStdout(), id, "Pool")
		return nil
	}}

func init() { CreateCmd.AddCommand(create_poolCmd) }
