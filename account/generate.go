package account

// TODO add real account creation
func generate(password string, name string) (address string, seed string, error error) {
	return "0x0000000000000000000000000000000000000000", "this is my long secure seed that help me regenerate my account keys", nil
}

// Generate an account based on some predefined data
func (account *Account) Generate() error {
	addr, seed, err := generate(account.Password, account.Name)
	if err != nil {
		return err
	}
	account.Address = addr
	account.Seed = seed
	return nil
}
