package config

type ConfigStruct struct {
	From string
	Pass string
	To string
}

func GetConfig() ConfigStruct{
	Config := ConfigStruct{
		From : "rohit123sinha456@gmail.com",
		Pass : "rbrn bodm wljf gxaq",
		To : "rohit.sinha@ailabs.academy",
	}
	return Config
}