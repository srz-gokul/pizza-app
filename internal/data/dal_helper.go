package data

// DBStatusInfo is to get the status of the DB
func (r *repo) DBStatusInfo() (string, error) {
	if err := r.Ping(); err != nil {
		return "nok", err
	}
	return "postgres db is working perfectly", nil
}
