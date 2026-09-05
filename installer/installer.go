package main

import (
    "flag"
    "log"
    "os"
    "path/filepath"
    "regexp"
)

func UNUSED(x ...interface{}) {}

func copyFiles (src string, dst string, filter *regexp.Regexp, callback func(path string)) error {
    files, err:= os.ReadDir(src)
    if err == nil {
        for _, file:= range files {
            if len(filter.FindStringSubmatch(file.Name())) > 0 {
                srcPath:= filepath.Join(src, file.Name())
                dstPath:= filepath.Join(dst, file.Name())
                if data, err:= os.ReadFile(srcPath); err == nil {
                    os.WriteFile(dstPath, data, file.Type())
                }
                callback(dstPath)
            }
        }
    }
    return err
}

func main() {
    module_dir:= flag.String(
        "module-dir", "",
        "the path of the module directory")
    toolbox_dir:= flag.String(
        "toolbox-dir", "",
        "the path of the toolbox directory")
    config_dir:= flag.String(
        "config-dir", "",
        "the path of the config directory")
    flag.Parse()

    UNUSED(module_dir, toolbox_dir, config_dir)

    fontDir:= getFontDir()

    exe, err:= os.Executable()
    if err != nil {
        log.Fatal(err)
    }
    exeDir:= filepath.Dir(exe)
    dataDir:= filepath.Join(exeDir, "data")

    rexOtf:= regexp.MustCompile("\\.otf$")
    rexTtf:= regexp.MustCompile("\\.ttf$")
    registerFontFile:= registerFontFile(*toolbox_dir)
    copyFiles(dataDir, fontDir, rexOtf, registerFontFile)
    copyFiles(dataDir, fontDir, rexTtf, registerFontFile)
    registerFontDir(*toolbox_dir)(fontDir)
}

