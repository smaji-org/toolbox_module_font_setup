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

func unregisterFontFile(toolboxDir string) func(string) {
    UNUSED(toolboxDir)
    unregister:= func (path string) {
    }
    return unregister
}

func unregisterFontDir(toolboxDir string) func(string) {
    UNUSED(toolboxDir)
    unregister:= func (fontDir string) {
        exec.Command("fc-cache", fontDir)
    }
    return unregister
}

