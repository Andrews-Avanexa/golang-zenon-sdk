package src

import (
	"os"
	"path/filepath"
	"runtime"
)

const znnSdkVersion = "5.0.1"

const znnRootDirectory = "znn"

type ZnnPaths struct {
	main   string
	wallet string
	cache  string
}

func SetZnnPaths(main string, wallet string, cache string) ZnnPaths {

	return ZnnPaths{main, wallet, cache}
}

func GetDefaultPaths() ZnnPaths {
	var main string
	if runtime.GOOS == "linux" {
		home := os.Getenv("XDG_CONFIG_HOME")
		main = filepath.Dir(filepath.Join(home, ".$znnRootDirectory"))
	} else if runtime.GOOS == "darwin" {
		home := os.Getenv("XDG_CONFIG_HOME")
		main = filepath.Dir(filepath.Join(home, "Library", znnRootDirectory))
	} else if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("AppData")
		}
		main = filepath.Dir(filepath.Join(home, znnRootDirectory))
	} else {
		home := os.Getenv("XDG_CONFIG_HOME")
		main = filepath.Dir(filepath.Join(home, znnRootDirectory))
	}
	return SetZnnPaths(main, filepath.Dir(filepath.Join(main, "wallet")), filepath.Dir(filepath.Join(main, "syrius")))
}

var znnDefaultPaths ZnnPaths = GetDefaultPaths()
var znnDefaultDirectory string = znnDefaultPaths.main
var znnDefaultWalletDirectory string = znnDefaultPaths.wallet
var znnDefaultCacheDirectory string = znnDefaultPaths.cache

func EnsureDirectoriesExist() {
	if _, err := os.Stat(znnDefaultWalletDirectory); os.IsNotExist(err) {
		os.Mkdir(znnDefaultWalletDirectory, 0755)
	}
	if _, err := os.Stat(znnDefaultCacheDirectory); os.IsNotExist(err) {
		os.Mkdir(znnDefaultCacheDirectory, 0755)
	}
}

var NetId int = 6

const DefaultZnnGateway string = "peers.znn.space"

type ZnnSdkException struct {
	Message string
}

func (ZSE *ZnnSdkException) ToString() string {
	var msg string = ZSE.Message
	if msg == "" {
		return "Zenon SDK Exception"
	}
	return "Zenon SDK Exception: " + msg
}
