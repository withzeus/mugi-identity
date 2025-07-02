package config

type Sever struct {
	SK  string
	UID string
}

var ServerLookup = Sever{
	SK:  "MUGISK",
	UID: "MGUID",
}
