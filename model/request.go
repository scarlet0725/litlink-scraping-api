package model

type RegisterUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUser struct {
	FamilyName string `json:"family_name"`
	GivenName  string `json:"given_name"`
	Email      string `json:"email"`
}

type CreateArtist struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
