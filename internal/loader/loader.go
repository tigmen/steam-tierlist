package loader

type Loader struct {
	Config *Config;
}

func NewLoader() *Loader{
	l := &Loader{
		Config: NewConfig(),
	}

	return l
}

func (l Loader) Load(){

}
