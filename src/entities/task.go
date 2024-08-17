package entities

type Task struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Create      string `json:"create"`
	Done        bool   `json:"done"`
}

func NewTask(id int, description string, create string, done bool) *Task {
	return &Task{
		Id:          id,
		Description: description,
		Create:      create,
		Done:        done,
	}
}

/*
func (t *Task) SetId(id int) {
	t.id = id
}

func (t *Task) GetId() int {
	return t.id
}

func (t *Task) SetDescription(description string) {
	t.description = description
}

func (t *Task) GetDescription() string {
	return t.description
}
func (t *Task) SetCreate(create string) {
	t.create = create
}

func (t *Task) GetCreate() string {
	return t.create
}

func (t *Task) SetDone(done bool) {
	t.done = done
}

func (t *Task) GetDone() bool {
	return t.done
}
*/
