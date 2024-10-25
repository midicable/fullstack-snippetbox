package repository


type SnippetsRepository struct {
	db string
}

func NewSnippetsRepository(db string) *SnippetsRepository {
	return &SnippetsRepository{db: db}
}

func (sr *SnippetsRepository) GetSnippet() string {
	return "HUYOVIY SNIPPET"
}