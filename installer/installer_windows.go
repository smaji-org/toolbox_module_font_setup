//go:build windows

package main

import (
    "os"
    "path/filepath"
    "os/exec"
    "log"
)

func getHomeDir() string {
    homeDir, _:= os.UserHomeDir()
    return homeDir
}

func getLocalDataDir() string {
    dataHome, ok:= os.LookupEnv("LocalAppData")
    if !ok {
        homeDir:= getHomeDir()
        dataHome= filepath.Join(homeDir, "AppData", "Local")
    }
    return dataHome
}

func getFontDir() string {
    localDataDir:= getLocalDataDir()
    fontDir:= filepath.Join(localDataDir, "Microsoft", "Windows", "Fonts")
    return fontDir
}

func registerFontFile(toolboxDir string) func(string) {
    win_font:= filepath.Join(toolboxDir, "cjkv_toolbox_win_font.exe")
    register:= func (path string) {
        cmdRegisterFont:= exec.Command(win_font, filepath.Base(path), path)
        if err:= cmdRegisterFont.Run(); err != nil {
            log.Fatal(err)
        }
    }
    return register
}

func registerFontDir(toolboxDir string) func(string) {
    UNUSED(toolboxDir)
    register:= func (fontDir string) {
    }
    return register
}

