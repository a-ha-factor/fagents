package types

type Fagent struct {
	Id        string
	FullName  string
	Dob       string
	Ogrn      string
	Inn       string
	RegNum    string
	Snils     string
	Address   string
	Resources string
	Members   string
	Law       string
	DateIn    string
	DatePubl  string
	DateOut   string
}

type FagentsList []Fagent
