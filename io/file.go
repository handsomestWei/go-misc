package io

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// 文件夹深拷贝。包括拷贝子文件夹
func CopyDirDepth(srcPath string, dstPath string) error {
	// 校验
	if srcInfo, err := os.Stat(srcPath); err != nil {
		return err
	} else {
		if !srcInfo.IsDir() {
			return errors.New("srcPath not a dir")
		}
	}
	if dstInfo, err := os.Stat(dstPath); err != nil {
		return err
	} else {
		if !dstInfo.IsDir() {
			return errors.New("dstPath not a dir")
		}
	}

	// 遍历
	err := filepath.Walk(srcPath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if !f.IsDir() {
			path := strings.Replace(path, "\\", "/", -1)
			dstNewPath := strings.Replace(path, srcPath, dstPath, -1)
			copyFile(path, dstNewPath)
		}
		return nil
	})
	return err
}

// 生成目录并拷贝文件
func copyFile(srcPath, dstPath string) (w int64, err error) {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return
	}

	defer srcFile.Close()
	// 校验路径上各级目录是否存在
	dstPaths := strings.Split(dstPath, "/")
	path := ""
	for index, dir := range dstPaths {
		if index < len(dstPaths)-1 {
			path = path + dir + "/"
			if isPathExist(path) == false {
				// 创建目录
				err := os.Mkdir(path, os.ModePerm)
				if err != nil {
					return
				}
			}
		}
	}

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return
	}
	defer dstFile.Close()
	return io.Copy(dstFile, srcFile)
}

// 文件路径是否存在
func isPathExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}
