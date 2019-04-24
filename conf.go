package cube

type ApplicationConf struct {
	Addr           string
	ReadTimeout    int
	WriteTimeout   int
	MaxHeaderBytes int
}

type Conf struct {
	ApplicationConf ApplicationConf
}
