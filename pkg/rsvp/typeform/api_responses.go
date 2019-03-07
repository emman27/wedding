package typeform

type responsesSchema struct {
	TotalItems int    `json:"total_items"`
	PageCount  int    `json:"page_count"`
	Items      []item `json:"items"`
}

type item struct {
	Answers []answer `json:"answers"`
}

func (i *item) formCompleted() bool {
	return len(i.Answers) == NumberOfQuestionsInForm
}

func (i *item) GetNumberOfAttendees() int {
	if i.formCompleted() {
		return i.Answers[NumberOfAttendeesQuestionNumber].Choice.Count()
	}
	return 0
}

type answer struct {
	Choice choice `json:"choice"`
}

type choice struct {
	Label choiceOption `json:"label"`
}

func (c *choice) Count() int {
	count, ok := choiceOptions[c.Label]
	if !ok {
		panic("This form is not configured correctly!")
	}
	return count
}

type choiceOption string
