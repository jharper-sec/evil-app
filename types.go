package main

//PageHeader ..
type PageHeader struct {
	Title       string
	Description string
}

// UserPageData ...
type UserPageData struct {
	PageHeader
	Users []User
}

// SubscriberPageData ...
type SubscriberPageData struct {
	PageHeader
	Subscribed bool
	Name       string
	Email      string
}

// WikiArticle ...
type WikiArticle struct {
	UUID    string
	Subject string
	Content string
}

// WikiPage ...
type WikiPage struct {
	PageHeader
	WikiArticles    []WikiArticle
	SelectedArticle WikiArticle
}

// User ...
type User struct {
	ID        int
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Company   string `json:"company"`
	Title     string `json:"title"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	DOB       string `json:"dob"`
	SSN       string `json:"ssn"`
	Salary    int    `json:"salary"`
	Admin     string `json:"admin"`
}
