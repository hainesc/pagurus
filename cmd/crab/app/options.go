package app

const (
	defaultServicePort = 8956
	defaultHealthzPort = 8957
	defaultMetricsPort = 8958
	defaultGlogRatateSize = 64 * 1024 * 1024 // 64MB
)
// Options contains everything necessary to create and run a piculet server.
type Options struct {
	// function is the function name that crab served.
	Function    string
	// servicePort is the port to be used by the service server.
	ServicePort int32
	// healthzPort is the port to be used by the healthz server.
	HealthzPort int32
	// metricsPort is the port to be used by the metrics server.
	MetricsPort int32
}

// NewOptions returns initialized Options
func NewOptions(f *Flags) *Options {
	return &Options{
		Function: f.FunctionName,
	}
}

// Applydefaults applies the default vaules for piculet, such as glog.MaxSize = 64MB
func (o *Options) applyDefaults() {
}

// TODO: args and config.yaml.

func NewCrabConfiguration() *Options {
	return &Options{
		ServicePort: 8080,
	}
}

func (o *Options) Complete() error {
	o.applyDefaults()
	return nil
}

// Run runs the specified Piculet.
func(o *Options) Run() (*Crab, error) {
	return nil, nil
}

// Validate validates all the required options.
func (o *Options) Validate() []error {
	var errs []error
	return errs
}
