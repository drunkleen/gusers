package utmpx

// GetAll retrieves all the entries from the utmpx database.
// It resets the database pointer, reads each entry until no more entries are available,
// and returns a slice of Entry structs representing logged-in users.

func GetAll() (entries []Entry) {
	entries = make([]Entry, 0)
	var entry *Entry
	resetutmpx()
	for {
		entry = next()
		if entry == nil {
			closeutmpx()
			break
		}
		entries = append(entries, *entry)
	}
	return entries
}
