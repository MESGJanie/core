package types

// Account is a structure that contains all informations about an account
type Account struct {
	Name     string
	Address  string
	Password string
	Seed     string
}

func (account *Account) String() string {
	return account.Name + " " + account.Address
}
