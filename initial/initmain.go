package initial

import (
	"go-spg/rpc"
	"go-spg/utils"
)

type InitMain struct{}

func (i *InitMain) AppInitMain() bool {
	utils.PrintInfo("Default data directory %s", utils.GetDefaultDataDir())
	utils.PrintInfo("Using data directory %s", utils.GetDataDir())
	utils.PrintInfo("Using config file %s", utils.GetConfigFile())

	utils.ReadConfigParams()

	utils.ShowArgs()

	go func() {
		rpc.RunRPCServer()
	}()

	return true
}
