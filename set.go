package main

import (
	"log"
)

func set(labels []label) error {
	db, err := openDB()
	if err != nil {
		return err
	}
	defer db.close()
	// log.Println("lbls1:", labels)
	err = db.set(labels)
	// log.Println("lbls2:", labels)

	if err == nil {
		cliPrint(labels)
	}
	// log.Println("lbls3:", labels)

	return err
}

func (db *db) set(labels []label) error {
	// log.Println("front:", labels)
	db.Lock()
	defer db.Unlock()
	m := mapLabels(labels)
	labels = delKey(labels, "name")
	labels = delKey(labels, "action")
	labels = delKey(labels, "type")
	// newlabels := labels
	// delKey(&labels, "name")
	// delKey(&labels, "type")
	// delKey(&labels, "action")
	// newlabels = delKey(newlabels, "path")
	// delKey2(&newlabels, "name")
	// delKey2(&newlabels, "type")
	// delKey2(&newlabels, "action")
	
	switch m["type"] {
	case "host":
		{
			db.hosts[m["name"]] = labels
		}
	case "file":
		{
			db.files[m["name"]] = labels
			// db.cleanOrphans()
		}
	case "group":
		{
			db.groups[m["name"]] = labels
		}
	case "script":
		{
			// delKey(&labels, "path")
			// log.Println(labels)
			db.scripts[m["name"]] = labels
			// db.cleanOrphans()
		}
	}
	// log.Println("end:", labels)
	return nil
}

func tmp2() {
	log.Println()
}