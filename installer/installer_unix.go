//go:build unix

package main

import (
    "os"
    "path/filepath"
    "os/exec"
)

func getHomeDir() string {
    homeDir, err:= os.UserHomeDir()
    if err != nil {
        homeDir= "/"
    }
    return homeDir
}

func getLocalDataDir() string {
    xdgDataHome, ok:= os.LookupEnv("XDG_DATA_HOME")
    if !ok {
        homeDir:= getHomeDir()
        xdgDataHome= filepath.Join(homeDir, ".local", "share")
    }

    return xdgDataHome
}

func getFontDir() string {
    localDataDir:= getLocalDataDir()
    fontDir:= filepath.Join(localDataDir, "fonts")
    return fontDir
}

func registerFontFile(toolboxDir string) func(string) {
    UNUSED(toolboxDir)
    register:= func (path string) {
    }
    return register
}

func registerFontDir(toolboxDir string) func(string) {
    UNUSED(toolboxDir)
    register:= func (fontDir string) {
        exec.Command("fc-cache", fontDir)
    }
    return register
}

