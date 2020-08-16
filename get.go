// SPDX-License-Identifier: MIT

package main

// import "log"

// import "log"

func show(labels []label) error {
	db, err := openDB()
	if err != nil {
		return err
	}
	result, err := db.get(labels)
	if err == nil {
		cliPrint(result...)
	}
	return err
}

func (db *db) get(query []label) ([][]label, error) {
	db.RLock()
	defer db.RUnlock()
	queryMap := mapLabels(query)
	query = delKey(query, "type")
	query = delKey(query, "action")
	results := [][]label{}
	// log.Println(queryMap)
	// log.Println(query)
	switch queryMap["type"] {
	case "script":
		{
			for name, labels := range db.scripts {
				// log.Println(name, labels)
				// log.Println(queryMap)
				match := map[string]bool{}
				script := mapLabels(labels)
				script["name"] = name
				for k := range queryMap {
					if k == "revision" {
						if queryMap[k] == script[k][:8] {
							match[k] = true
						}
					} else {
						if queryMap[k] == script[k] {
							match[k] = true
						}
					}
					
				}
				if len(match) == len(query) {
					labels = append(
						[]label{
							{Key:"name",Value:name},
							{Key:"type",Value:queryMap["type"]},
							{Key:"action",Value:queryMap["action"]},
						},
						labels...)
					results = append(results, labels)
				}
			}
			results = sortResults(results)
		}
	case "host":
		{
			for name, labels := range db.hosts {
				match := map[string]bool{}
				script := mapLabels(labels)
				script["name"] = name
				for k := range queryMap {
					if queryMap[k] == script[k] {
						match[k] = true
					}
				}
				if len(match) == len(query) {
					labels = append(
						[]label{
							{Key:"name",Value:name},
							{Key:"type",Value:queryMap["type"]},
							{Key:"action",Value:queryMap["action"]},
						},
						labels...)
					results = append(results, labels)
				}
			}
			results = sortResults(results)
			// return results, nil
		}
	case "group":
		{
			for name, labels := range db.groups {
				match := map[string]bool{}
				script := mapLabels(labels)
				script["name"] = name
				for k := range queryMap {
					if queryMap[k] == script[k] {
						match[k] = true
					}
				}
				if len(match) == len(query) {
					labels = append(
						[]label{
							{Key:"name",Value:name},
							{Key:"type",Value:queryMap["type"]},
							{Key:"action",Value:queryMap["action"]},
						},
						labels...)
					results = append(results, labels)
				}
			}
			results = sortResults(results)
			// return results, nil
		}
	case "file":
		{
			for name, labels := range db.files {
				match := map[string]bool{}
				script := mapLabels(labels)
				script["name"] = name
				for k := range queryMap {
					if queryMap[k] == script[k] {
						match[k] = true
					}
				}
				if len(match) == len(query) {
					labels = append(
						[]label{
							{Key:"name",Value:name},
							{Key:"type",Value:queryMap["type"]},
							{Key:"action",Value:queryMap["action"]},
						},
						labels...)
					results = append(results, labels)
				}
			}
			results = sortResults(results)
			// return results, nil
		}
	}
	// log.Println(results)
	if len(results) == 0 {
		results = append(results, []label{
			{Key:"name",Value:"NONE"},
			// {Key:"result",Value:"0"},
			{Key:"type",Value:queryMap["type"]},
			{Key:"action",Value:queryMap["action"]},
		})
	}
	return results, nil
}