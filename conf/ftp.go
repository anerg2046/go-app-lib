package conf

type FtpConf struct {
	Address     string
	Port        int
	User        string
	Pass        string
	LocalPath   string //本地下载路径
	ArchivePath string //本地归档路径
	RemotePath  string //远程目录路径
	Timeout     int
}
