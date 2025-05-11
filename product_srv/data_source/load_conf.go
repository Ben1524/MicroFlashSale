package data_source

import (
	"encoding/json"
	"log"
	"os"
)

type MysqlInfo struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Database string `json:"database"`
}

var MysqlInfoSource *MysqlInfo

func InitConf() {
	confFile, err := os.Open("user_srv/conf/mysql_conf.json")
	defer confFile.Close()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	jsonParser := json.NewDecoder(confFile)
	if err = jsonParser.Decode(&MysqlInfoSource); err != nil {
		log.Fatal(err)
		panic(err)
	}
	log.Println("mysql info:", MysqlInfoSource)
}
