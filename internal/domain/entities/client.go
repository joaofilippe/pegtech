package entities

type Client struct {
	User
	Phone   string
	Address string
}

func NewClient(name, username, email, password, phone, address string) *Client {
	user := NewUser(name, username, email, password, UserTypeClient)
	client := &Client{
		User:    *user,
		Phone:   phone,
		Address: address,
	}

	return client
}

func (c *Client) Update(name, email, phone, address string) {
	c.User.Update(name, email)
	c.Phone = phone
	c.Address = address
}
