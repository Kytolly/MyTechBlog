package utils

import(
	"log"
	"gopkg.in/ini.v1"
)

var(
	AppMode 	string
	HttpPort    string

	Db         	string
	DbHost     	string
	DbPort      string
	DbUser     	string
	DbPassword  string
	DbName     	string
)

var initFile = "config/config.ini"

func init(){
	file, err := ini.Load(initFile)
	if err != nil{
		log.Println("utils/setting.go: Error In Loading Inition File:", err)
	}

	loadServer(file)
	loadData(file)
	log.Println("utils/setting.go: The Server and Data From Inition File Finished Loading Successful!")
}

func loadServer(file *ini.File){
	AppMode  = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString("4040")
}

func loadData(file *ini.File){
	Db		  = file.Section("database").Key("Db").MustString("mysql")
	DbHost    = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort    = file.Section("database").Key("DbPort").MustString("3306")
	DbUser    = file.Section("database").Key("DbUser").MustString("mytechblog")
	DbPassword= file.Section("database").Key("DbPassword").MustString("xqy05016")
	DbName    = file.Section("database").Key("DbName").MustString("mytechblog")
}