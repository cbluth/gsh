package main

import (
	"fmt"
	"log"
)

// action=execute
// type=script
// scope=host/group
// name=script_name
// target=host1/grp1
// sudo=true
// [{action execute} {type script} {scope -n} {target d1-node} {sudo squash} {name hostname}]
func execute(labels []label) error {
	// log.Println(labels)
	m := mapLabels(labels)
	db, err := openDB()
	if err != nil {
		return err
	}
	t := "" // type
	if m["scope"] == "-n" {
		t = "host"
	}
	if m["scope"] == "-g" {
		t = "group"
	}
	if t == "" {
		return fmt.Errorf("%s: %s", "scope unclear", m["scope"])
	}
	q := []label{
		{Key: "action", Value: "execute"},
		{Key: "type", Value: t},
		{Key: "name", Value: m["target"]},
	}
	lbs := []label{}
	result, err := db.get(q)
	if err == nil {
		// cliPrint(result...)
		// log.Println(result)
		lbs = delKey(result[0], "name")
		lbs = delKey(lbs, "name")
		lbs = delKey(lbs, "type")
		lbs = append(lbs, label{Key: "type", Value: "host"})
		// log.Println(v ...interface{})
		result, err = db.get(lbs)
		if err != nil {
			return err
		}
		// log.Println(result)
	}
	hasScript := db.hasType("script", m["name"])
	hasScope := db.hasType(t, m["target"])
	if hasScope && hasScript {
		mt := db.mapType("script", m["name"])
		log.Printf("Executing script %s revision=%s ...", m["name"], mt["revision"][:8])
		for i, l := range result {
			ml := mapLabels(l)
			log.Printf("%v= Host: %s | Address: %s | User: %s | sudo: %s", i+1, ml["name"], ml["address"], usr, m["sudo"])
			scr, err := db.getFile(mt["revision"])
			if err != nil {
				return err
			}
			stdout, stderr, err := runScript(string(scr), []string{}, host{Host: ml["address"], Port: 22}, false)
			if len(stdout) > 1 {
				fmt.Println(stdout)
			}
			if len(stderr) > 1 {
				fmt.Println("script err:", stderr)
			}
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	// log.Println(result)
	return nil
}

func getExecuteArgs(args []string) []label {
	scope := ""
	// scriptName := ""
	sudo := "false"
	target := ""
	for i, arg := range args {
		if arg == "-n" || arg == "-g" {
			scope = arg
			target = args[i+1]
			break
		}
		// log.Println(arg)
	}
	for _, arg := range args {
		if arg == "-su" {
			sudo = "raise"
		}
		if arg == "-sq" {
			// dont use sudo if script has enabled
			sudo = "squash"
		}
	}
	execArgs := []label{
		{Key: "action", Value: "execute"},
		{Key: "type", Value: "script"},
		{Key: "scope", Value: scope},
		{Key: "target", Value: target},
		{Key: "sudo", Value: sudo},
		{Key: "name", Value: args[len(args)-1]},
	}
	// execArgs = append(execArgs, label{})
	// os.Exit(0)
	return execArgs
}