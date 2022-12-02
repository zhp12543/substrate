package v1

type CallArg struct {
	Name string
	Type string
}

type Call struct {
	Name string
	Args []CallArg
	Docs []string
}

type Event struct {
	Name string
	Args []string
	Docs []string
}
