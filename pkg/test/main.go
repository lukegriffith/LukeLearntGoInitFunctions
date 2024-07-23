package test

var BlahSingleton Blah

type Blah struct {
	Prop string
}

func init() {
	BlahSingleton = Blah{"test"}

}
