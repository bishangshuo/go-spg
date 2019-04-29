// spg.go
package main

import (
	"fmt"
	"go-spg/constv"
	"go-spg/db/walletdb"
	"go-spg/initial"
	"go-spg/message"
	"go-spg/utils"
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

	//for i:=0;i<10;i++ {
	//	wa := wallet.NewWallet()
	//	address := wa.GetNewAddress()
	//	fmt.Println(address)
	//}

	//walletdb := &db.WalletDB{}
	wd  := walletdb.GetInstance()
	address := "key100000001"
	b := wd.SaveKeyWithAddress(address, "value1000000001")
	if b{
		if r,key:=wd.GetKeyWithAddress(address); r {
			utils.PrintInfo("get key with address : %s, key=%s", address, key)
		}else{
			utils.PrintInfo("key invalid with address:%s", address)
		}
	}

	utils.SetupEnviroment()

	message.SetupMessageThread()

	AppInit(os.Args)

	utils.PrintInfo("spgd startup ok")
	ch := make(chan int)
	<-ch
}
