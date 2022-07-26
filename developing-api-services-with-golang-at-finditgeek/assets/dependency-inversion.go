type TaskRepository interface {
	Fetch() ([]*Task, error)
}

type TaskRepoPostgresql struct{}

func (p TaskRepoPostgresql) Fetch() ([]*Task, error) {
	// implement fetch from DB
}

type TaskRepoVendor struct{}

func (v TaskRepoVendor) Fetch() ([]*Task, error) {
	// implement fetch from vendor
}