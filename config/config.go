package config

import "github.com/tkanos/gonfig"

// struct ini bakalan dipanggil untuk koneksi
type Configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
}

// mendapatkan data config.json, lalu dimasukkan ke struct Configuration
// config menggunakan json pakai library tkanos/gonfig
func GetConfig() Configuration {
	conf := Configuration{}
	gonfig.GetConf("config/config.json", &conf)
	return conf
}