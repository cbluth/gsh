package main

func del(labels []label) error {
	db, err := openDB()
	if err != nil {
		return err
	}
	defer db.close()
	info, err := db.del(labels)
	if err == nil {
		cliPrint(info)
	}
	return err
}

func (db *db) del(labels []label) ([]label, error) {
	m := mapLabels(labels)
	info, err := db.get(
		[]label{
			{Key: "name", Value: m["name"]},
			{Key: "type", Value: m["type"]},
			{Key: "action", Value: "delete"},
		},
	)
	if err != nil {
		return []label{}, err
	}
	db.Lock()
	defer db.Unlock()
	switch m["type"] {
	case "host":
		{
			delete(db.hosts, m["name"])
		}
	case "file":
		{
			delete(db.files, m["name"])
			// db.cleanOrphans()
		}
	case "group":
		{
			delete(db.groups, m["name"])
		}
	case "script":
		{
			delete(db.scripts, m["name"])
			// db.cleanOrphans()
		}
	}
	// log.Println(info[0])
	return info[0], err
}