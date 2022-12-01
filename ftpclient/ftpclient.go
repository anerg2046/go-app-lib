package ftpclient

import (
	"fmt"
	"go-app/lib/conf"
	"go-app/lib/logger"
	"io"
	"os"
	"time"

	"github.com/jlaffaye/ftp"
	"go.uber.org/zap"
)

const LOG_MODULE = "FTP ERROR"

type ftpClient struct {
	Conn   *ftp.ServerConn
	config conf.FtpConf
}

type FileInfo struct {
	Entry *ftp.Entry
	Dir   string
}

func Dail(cfg conf.FtpConf) (ftpclient *ftpClient, err error) {
	Conn, err := ftp.Dial(
		fmt.Sprintf("%s:%d", cfg.Address, cfg.Port),
		ftp.DialWithTimeout(time.Duration(cfg.Timeout)*time.Second),
		// ftp.DialWithExplicitTLS(&tls.Config{InsecureSkipVerify: true}),
		ftp.DialWithDisabledEPSV(false),
	)
	if err != nil {
		logger.Error(LOG_MODULE, zap.Error(err))
		return
	}
	err = Conn.Login(cfg.User, cfg.Pass)
	if err != nil {
		logger.Error(LOG_MODULE, zap.Error(err))
		return
	}

	return &ftpClient{Conn: Conn, config: cfg}, nil
}

func (f *ftpClient) Close() (err error) {
	err = f.Conn.Quit()
	if err != nil {
		logger.Error(LOG_MODULE, zap.Error(err))
	}
	return
}

func (f *ftpClient) Download(entry *ftp.Entry) {
	// 如果文件在上传中，则忽略
	if entry.Size == 0 {
		return
	}
	localPath := f.config.LocalPath + entry.Name

	localFile, err := os.OpenFile(localPath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		logger.Error(LOG_MODULE, zap.Error(err))
	}
	defer localFile.Close()
	remoteFile, err := f.Conn.Retr(entry.Name)
	if err != nil {
		logger.Error(LOG_MODULE, zap.Error(err))
	}
	defer remoteFile.Close()
	_, err = io.Copy(localFile, remoteFile)
	if err != nil {
		logger.Error(LOG_MODULE, zap.Error(err))
	}
}
