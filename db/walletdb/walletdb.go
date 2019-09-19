package walletdb

import (
	"github.com/btcsuite/goleveldb/leveldb"
	"go-spg/utils"
	"sync"
)

type WalletDB struct{
	pWalletDB *leveldb.DB
}

var (
	once sync.Once
	gWalletDBInstance *WalletDB
	gWalletDBLock *sync.RWMutex
)

func GetInstance() (*WalletDB){
	once.Do(func() {
		if pWalletDB_temp, err_l := leveldb.OpenFile(utils.GetDataDir()+"/wallet.db", nil); err_l == nil {
			gWalletDBInstance = &WalletDB{}
			gWalletDBInstance.pWalletDB = pWalletDB_temp
			gWalletDBLock = new(sync.RWMutex)
		}
	})
	return gWalletDBInstance
}

func (wd *WalletDB)SaveKeyWithAddress(address string, key string) bool{
	gWalletDBLock.Lock()
	err := wd.pWalletDB.Put([]byte(address), []byte(key), nil)
	gWalletDBLock.Unlock()

	if err != nil{
		utils.PrintError("can not write wallet db :%s", err.Error())
		return false
	}
	return true
}

func (wd *WalletDB)GetKeyWithAddress(address string) (bool, string){
	gWalletDBLock.Lock()
	k, err := wd.pWalletDB.Get([]byte(address), nil)
	gWalletDBLock.Unlock()

	if err != nil {
		utils.PrintError("can not get key with address: %s, info:%s", address, err.Error())
		return false, ""
	}
	return true, string(k)
}

func (wd *WalletDB)RemoveKeyWithAddress(address string) bool {
	gWalletDBLock.Lock()
	err := wd.pWalletDB.Delete([]byte(address), nil)
	gWalletDBLock.Unlock()

	if err != nil {
		utils.PrintError("can not delete key with address: %s, info:%s", address, err.Error())
		return false
	}
	return true
}