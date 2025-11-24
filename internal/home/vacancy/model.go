package vacancy

type VacancyCreateForm struct {
	Email    string
	Role     string
	Company  string
	Salary   string
	Type     string
	Location string
}

type Vacancy struct {
	Email    string
	Role     string
	Company  string
	Salary   string
	Type     string
	Location string
}

func NewVacancyFromCreateForm(form VacancyCreateForm) *Vacancy {
	return &Vacancy{
		Email:    form.Email,
		Role:     form.Role,
		Company:  form.Company,
		Salary:   form.Salary,
		Type:     form.Type,
		Location: form.Location,
	}
}
