package structs

type Flags struct {
	Tester bool `cli:"tester" cliAlt:"t" usage:"Run sender test"`
	Worker bool `cli:"worker" cliAlt:"w" usage:"Run worker instance"`
}
