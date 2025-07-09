package logging

import (
	"example.com/example/pkg/file"
	"example.com/example/pkg/setting"
	"fmt"
	"log"
	"os"
	"time"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.AppSetting.RunTimePath, setting.AppSetting.LogSavePath)
}

func getLogFileName() string {
	return fmt.Sprintf("%s.%s",
		time.Now().Format(setting.AppSetting.LogTimeFormat),
		setting.AppSetting.LogFileExt,
	)
}

func openLogFile(filePath, fileName string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("get current path failed :%v", err)
	}
	src := dir + "/" + filePath
	perm := file.CheckPermission(src)
	if perm {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}
	err = file.IsNotExistMkDir(src)
	if err != nil {
		log.Fatalf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}
	f, err := os.OpenFile(src+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("fail to OpenFile :%v", err)
	}

	return f, nil
}
