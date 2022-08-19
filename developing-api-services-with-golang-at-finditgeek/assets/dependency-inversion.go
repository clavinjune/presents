type TaskRepository interface {
	Fetch() ([]*Task, error)
}

type TaskRepoPostgresql struct{}

func (p TaskRepoPostgresql) Fetch() ([]*Task, error) {}

type TaskRepoVendor struct{}

func (v TaskRepoVendor) Fetch() ([]*Task, error) {}

type TaskService struct { // HL1
	// don't need to know whether it's Postgresql or Vendor // HL1
	// gives dependants more freedom to choose the TaskRepository implementation // HL1
	repo TaskRepository // HL1
} // HL1