// spg.go
package main

import (
	"fmt"
	"go-spg/constv"
	"go-spg/initial"
	"go-spg/message"
	"go-spg/utils"
	"go-spg/wallet"
	"os"
)

func AppInit(args []string) bool {
	var fRet bool = false

	utils.ParseParameters(args)
	utils.ShowArgs()

	if utils.IsArgSet("-?") || utils.IsArgSet("-h") || utils.IsArgSet("-help") || utils.IsArgSet("-version") {
		strUsage := fmt.Sprintf("%s Daemon version %s\n", constv.PACKAGE_NAME, constv.VERSION_FULL)
		if utils.IsArgSet("-version") {
			strUsage = fmt.Sprintf("%s%s", strUsage, constv.GetVersionInfo())
		} else {
			strUsage = fmt.Sprintf("%s\nUsage:\n spgd [options]       Start %s Daemon\n\n%s\n", strUsage, constv.PACKAGE_NAME, constv.GetHelpInfo())
		}
		fmt.Println(strUsage)
		return true
	}

	mainInit := &initial.InitMain{}
	fRet = mainInit.AppInitMain()

	return fRet
}

func main() {
	utils.PrintInfo("spgd starting...")

	wl := wallet.NewWallet()
	address := wl.GetNewAddress()
	fmt.Println("getnewaddress:", address)
	privkey, err := wl.DumpPrivKey(address)
	if err == nil {
		fmt.Println("private key:", privkey)
	}else{
		fmt.Println(err.Error())
	}
	utils.SetupEnviroment()

	message.SetupMessageThread()

	AppInit(os.Args)

	utils.PrintInfo("spgd startup ok")
	ch := make(chan int)
	<-ch
}
