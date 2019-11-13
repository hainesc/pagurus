package app

import (
	"os"
	"github.com/golang/glog"
)

const Version = "v0.1.0"

func NewCrabCommand(flags *Flags) *Command {
	return &Command{
		Run: func() error {
			if flags.ConfigFile != "" {
				err := flags.LoadConfig(flags.ConfigFile)
				if err != nil {
					glog.Fatalf("TODO")
				}
			}
			opts := NewOptions(flags)
			err := opts.Complete()
			if err != nil {
				glog.Fatalf("TODO")
			}

			if errs := opts.Validate(); len(errs) != 0 {
				glog.Errorf("Error in config")
				for _, err := range errs {
					glog.Errorf("%v", err)
				}
				os.Exit(1)
			}
			crab, _ := NewCrab(opts)
			if err := crab.Run(); err != nil {
				glog.Fatalln(err)
			}
			return nil
		},
	}
}

type Crab struct {
}

func NewCrab(complete *Options) (*Crab, error) {
	return nil, nil
}

func (p *Crab) Write() {
}

func (p *Crab) Run() error {
	// TODO:
	return nil
}
