package app

import (
	"os"
	// "net/http"

	"github.com/golang/glog"
	// "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
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
	Function string
}

func NewCrab(complete *Options) (*Crab, error) {
	return &Crab{
		Function: complete.Function,
	}, nil
}

func (c *Crab) Run() error {
	// Download the code from ConfigMap and server as a http server.
	// TODO: move clientset into Crab as a member
	config, err := rest.InClusterConfig()
	if err != nil {
		glog.Fatalln(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		glog.Fatalln(err.Error())
	}

	_, err = clientset.CoreV1().Pods("default").Get(c.Function, metav1.GetOptions{})
	if err != nil {
		glog.Fatalln(err.Error())
	}

	// code := fn.Code

	return nil
}
