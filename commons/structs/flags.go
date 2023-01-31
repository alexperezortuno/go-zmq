package structs

type Flags struct {
	Tester      bool `cli:"tester" cliAlt:"t" usage:"Run sender test"`
	Worker      bool `cli:"worker" cliAlt:"w" usage:"Run like a worker"`
	NumOfWorker int  `cli:"numOfWorker" cliAlt:"n" usage:"Number of worker instances"`
	Sink        bool `cli:"sink" cliAlt:"s" usage:"Run like a sink"`
}
