package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"github.org/x/sys/windows/registry"
	"go-spg/constv"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

var (
	mapArgs     map[string]string
	mapArgsLock sync.RWMutex
)

func init() {
	fmt.Println("util init")
}

func SetupEnviroment() {
	PrintInfo("SetupEnviroment ...")
}

func ParseParameters(argv []string) {

	mapArgsLock.Lock()
	defer mapArgsLock.Unlock()

	mapArgs = make(map[string]string)

	argc := len(argv)

	for i := 1; i < argc; i++ {
		str := argv[i]
		var strValue string
		if is_index := strings.Index(str, "="); is_index != -1 {
			strValue = str[is_index+1:]
			str = str[:is_index]
		}
		if str[0] != '-' {
			break
		}

		if len(str) > 1 && str[1] == '-' {
			str = str[1:]
		}

		mapArgs[str] = strValue
	}
}

func ReadConfigParams() {
	configFile := GetConfigFile()
	if exists, err := PathExists(configFile); err == nil && exists {
		fd, e := os.Open(configFile)
		if e != nil {
			panic(e)
		}
		defer fd.Close()

		scanner := bufio.NewScanner(fd)
		for scanner.Scan() {
			line := scanner.Text()
			kv := strings.Split(line, "=")
			if len(kv) == 2 {
				key := "-" + kv[0]
				value := kv[1]
				if !IsArgSet(key) {
					mapArgs[key] = value
				}
			}
		}
	}
}

func IsArgSet(strArg string) bool {
	mapArgsLock.Lock()
	defer mapArgsLock.Unlock()

	_, set := mapArgs[strArg]
	return set
}

func GetArgString(strArg string) (string, bool) {
	mapArgsLock.Lock()
	defer mapArgsLock.Unlock()

	arg, set := mapArgs[strArg]
	return arg, set
}

func GetArgInt(strArg string) (int, bool) {
	if arg, ok := GetArgString(strArg); ok {
		if argInt, err := strconv.Atoi(arg); err == nil {
			return argInt, true
		}
	}
	return 0, false
}

func GetArgBool(strArg string) (bool, bool) {
	if arg, ok := GetArgString(strArg); ok {
		if argBool, err := strconv.ParseBool(arg); err == nil {
			return argBool, true
		}
	}
	return false, false
}

func SetArgString(strArg string, strValue string) bool {
	mapArgsLock.Lock()
	defer mapArgsLock.Unlock()
	mapArgs[strArg] = strValue
	return true
}

func SetArgInt(strArg string, intValue int) bool {
	strValue := strconv.Itoa(intValue)
	return SetArgString(strArg, strValue)
}

func SetArgBool(strArg string, boolValue bool) bool {
	strValue := strconv.FormatBool(boolValue)
	return SetArgString(strArg, strValue)
}

func ShowArgs() {
	PrintInfo("Input arguments in command line:")
	for key, value := range mapArgs {
		var byteInfo bytes.Buffer
		byteInfo.WriteString(key)
		byteInfo.WriteString("=")
		byteInfo.WriteString(value)
		PrintInfo(byteInfo.String())
	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetDataDir() string {
	dirArg, set := GetArgString("-datadir")
	if set {
		if _, err := PathExists(dirArg); err != nil {
			panic("error accour when Get data directory")
			return ""
		}
	}

	if len(dirArg) == 0 {
		dirArg = GetDefaultDataDir()
	}

	if len(dirArg) > 0 {
		if exists, err := PathExists(dirArg); err == nil {
			if !exists {
				if err := os.Mkdir(dirArg, os.ModePerm); err != nil {
					panic("error accour when Get data directory")
					return ""
				}
			}
		}
	}
	return dirArg
}

func GetDefaultDataDir() string {
	os_platform := runtime.GOOS
	switch os_platform {
	case "windows":
		return os.Getenv("APPDATA") + "\\" + constv.DEAMON_NAME
	case "darwin":
		return "/Library/Application Support/" + constv.DEAMON_NAME
	case "linux":
		return "/root/.SuperGoldChain" + constv.DEAMON_NAME
	}
	return ""
}

/* windows current version list
Operating system              Version number
----------------------------  --------------
Windows 10                      10.0
Windows Server 2019             10.0
Windows Server 2016             10.0
Windows 8.1                     6.3
Windows Server 2012 R2          6.3
Windows 8                       6.2
Windows Server 2012             6.2
Windows 7                       6.1
Windows Server 2008 R2          6.1
Windows Server 2008             6.0
Windows Vista                   6.0
Windows Server 2003 R2          5.2
Windows Server 2003             5.2
Windows XP 64-Bit Edition       5.2
Windows XP                      5.1
Windows 2000                    5.0
Windows ME                      4.9
Windows 98                      4.10
*/
func GetWindowsVersion() float64 {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		panic("GetWindowsVersion error")
	}
	defer k.Close()

	cv, _, err := k.GetStringValue("CurrentVersion")
	if err != nil {
		panic("GetWindowsVersion error")
	}
	fmt.Printf("CurrentVersion: %s\n", cv)

	ver, err := strconv.ParseFloat(cv, 64)
	if err == nil {
		return ver
	}
	return 0.0
}

func GetConfigFile() string {
	confFile, set := GetArgString("-conf")
	if set {
		if exists, err := PathExists(confFile); err != nil && exists {
			return confFile
		}
	}
	confFile = GetDataDir() + "\\" + constv.DEAMON_NAME + ".conf"
	return confFile
}

func GetRPCPort() int {
	rpcport, set := GetArgInt("-rpcport")
	if set && rpcport > 0 {
		return rpcport
	}
	return constv.DEFAULT_RPCPORT
}
