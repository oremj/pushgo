package simplepush

type StaticLocatorConf struct {
	Contacts   []string
	BucketSize int `toml:"bucket_size"`
}

type StaticLocator struct {
	logger     *SimpleLogger
	metrics    *Metrics
	contacts   []string
	bucketSize int
}

func (*StaticLocator) ConfigStruct() interface{} {
	return &StaticLocatorConf{
		BucketSize: 10,
	}
}

func (l *StaticLocator) Init(app *Application, config interface{}) error {
	conf := config.(*StaticLocatorConf)
	l.logger = app.Logger()
	l.metrics = app.Metrics()
	l.contacts = conf.Contacts
	l.bucketSize = conf.BucketSize
	return nil
}

func (l *StaticLocator) Close() error                      { return nil }
func (l *StaticLocator) Contacts(string) ([]string, error) { return l.contacts, nil }
func (l *StaticLocator) MaxParallel() int                  { return l.bucketSize }

func init() {
	AvailableLocators["static"] = func() HasConfigStruct { return new(StaticLocator) }
}
