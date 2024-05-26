package models

type Contact struct {
	Name        string  `json:"name"`
	LastName    string  `json:"lastName"`
	PhoneNumber string  `json:"phone_number"`
	Address     *string `json:"address,omitempty"`
}

type UpdateContactByName struct {
	Name        string  `json:"name"`
	LastName    string  `json:"lastName"`
	PhoneNumber *string `json:"phone_number,omitempty"`
	Address     *string `json:"address,omitempty"`
}

type UpdateContactByPhone struct {
	Name        *string `json:"name,omitempty"`
	LastName    *string `json:"lastName,omitempty"`
	PhoneNumber string  `json:"phone_number"`
	Address     *string `json:"address,omitempty"`
}

type SearchContactModel struct {
	Name        *string `json:"name,omitempty"`
	LastName    *string `json:"lastName,omitempty"`
	PhoneNumber *string `json:"phone_number,omitempty"`
}
