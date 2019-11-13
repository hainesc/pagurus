package app

import (
	"github.com/golang/glog"
)

type Flags struct {
	ConfigFile string
}

func (f *Flags) LoadConfig(config string) error {
	// TODO:
	glog.Infof("TODO: will load config from file %s\n", config)
	return nil
}
