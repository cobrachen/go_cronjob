package config

import "time"

// TomlConfig 設定檔案
type TomlConfig struct {
	Title   string
	JobList map[string]cronjob `toml:"cronjob"` // toml檔案是cronjob, golang變數是JobList
}

type cronjob struct {
	Name string
	CMD  string
	Spec string
}

type bgWorker struct {
	Name           string
	Status         string
	OverSig        chan []byte
	LastOverTS     *time.Time
	LastExecTS     *time.Time
	CurrentStartTS *time.Time
	NextStartTS    *time.Time
}
