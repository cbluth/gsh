package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"sort"
	"strings"

	// "strings"
	"crypto/sha256"
	"encoding/hex"
	"sync"
)

type (
	db struct {
		dir string
		hosts	map[string][]label
		groups	map[string][]label
		scripts	map[string][]label
		files	map[string][]label
		sync.RWMutex
	}
	label struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	// lbls []label
	dbjson struct {
		dir string
		Hosts   []gtype   `json:"hosts"`
		Groups  []gtype   `json:"groups"`
		Scripts []gtype	  `json:"scripts"`
		Files   []gtype   `json:"files"`
		sync.RWMutex
	}
	gtype struct {
		Name   string   `json:"name"`
		Labels []label  `json:"labels"`
	}
)

func openDB() (*db, error) {
	d := &db{
		dir: getDir(),
		hosts: map[string][]label{},
		groups: map[string][]label{},
		scripts: map[string][]label{},
		files: map[string][]label{},
	}
	err := d.load()
	if err != nil {
		return nil, err
	}
	return d, nil
}


func (db *db) load() (error) {
	path := db.dir + "/db.json"
	b := []byte{}
	err := *new(error)
	dbjson := &dbjson{}
	db.Lock()
	defer db.Unlock()
	if _, err = os.Stat(path) ; err == nil {
		b, err = ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		err = json.Unmarshal(b, dbjson)
		if err != nil {
			return err
		}
	} else if os.IsNotExist(err) {
		b, err := json.MarshalIndent(dbjson, "", "    ")
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(path, b, 0644)
		if err != nil {
			return err
		}
	}
	for _, g := range dbjson.Hosts {
		db.hosts[g.Name] = g.Labels
	}
	for _, g := range dbjson.Groups {
		db.groups[g.Name] = g.Labels
		// db.groups = append(
		// 	db.groups,
		// 	append(
		// 		[]label{
		// 			label{
		// 				Key: "name",
		// 				Value: g.Name,
		// 			},
		// 		},
		// 		g.Labels...,
		// 	),
		// )
	}
	for _, g := range dbjson.Scripts {
		db.scripts[g.Name] = g.Labels
	}
	for _, g := range dbjson.Files {
		db.files[g.Name] = g.Labels
	}
	return nil
}

func getDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}
	dir := usr.HomeDir + "/.gsh"
	if _, err := os.Stat(dir+"/db") ; os.IsNotExist(err) {
		err = os.MkdirAll(dir+"/db", os.ModePerm)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return dir
}

func mapLabels(labels []label) map[string]string {
	m := map[string]string{}
	for _, l := range labels {
		m[l.Key] = l.Value
	}
	return m
}

func (db *db) close() error {
	path := db.dir + "/db.json"
	db.Lock()
	dbjson := &dbjson{}
	for name, labels := range db.hosts {
		dbjson.Hosts = append(dbjson.Hosts, gtype{
			Name: name,
			Labels: labels,
		})
	}
	for name, labels := range db.groups {
		dbjson.Groups = append(dbjson.Groups, gtype{
			Name: name,
			Labels: labels,
		})
	}
	for name, labels := range db.scripts {
		dbjson.Scripts = append(dbjson.Scripts, gtype{
			Name: name,
			Labels: labels,
		})
	}
	for name, labels := range db.files {
		dbjson.Files = append(dbjson.Files, gtype{
			Name: name,
			Labels: labels,
		})
	}
	b, err := json.MarshalIndent(dbjson, "", "    ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, b, 0644)
	if err != nil {
		return err
	}
	return err
}

func sortResults(in [][]label) [][]label {
	names := []string{}
	for _, line := range in {
		m := mapLabels(line)
		names = append(names, m["name"])
	}
	sort.Strings(names)
	out := [][]label{}
	for _, name := range names {
		for _, line := range in {
			m := mapLabels(line)
			if name == m["name"] {
				out = append(out, line)
			}
		}
	}
	return out
}

func addFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	h := sha256.New()
	_, err = io.Copy(h, f)
	if err != nil {
		return "", err
	}
	err = f.Close()
	if err != nil {
		return "", err
	}
	filename := hex.EncodeToString(h.Sum(nil))
	f, err = os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	dir := getDir() +"/db/" + filename[:1]
	if _, err := os.Stat(dir) ; os.IsNotExist(err) {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return "", err
		}
	}
	dst, err := os.Create(dir + "/" + filename)
	if err != nil {
		return "", err
	}
	defer dst.Close()
	// TODO: include filesize in labels
	_, err = io.Copy(dst, f)
	if err != nil {
		return "", err
	}
	return filename, nil
}

func delKey(labels []label, key string) []label {
	m := mapLabels(labels)
	delete(m, key)
	new := []label{}
	for _, l := range labels {
		if _, ok := m[l.Key] ; ok {
			new = append(new, l)
		}
	}
	return new
}

// func (db *db) queryHosts(query []label) [][]label {
// 	db.RLock()
// 	defer db.RUnlock()
// 	queryMap := mapLabels(query)
// 	return [][]label{{}}
// }

func (db *db) hasType(t, n string) bool {
	db.RLock()
	defer db.RUnlock()
	switch t {
	case "script":
		{
			for name := range db.scripts {
				if name == n {
					return true
				}
			}
		}
	case "host":
		{
			for name := range db.hosts {
				if name == n {
					return true
				}
			}
		}
	case "group":
		{
			for name := range db.groups {
				if name == n {
					return true
				}
			}
		}
	case "file":
		{
			for name := range db.files {
				if name == n {
					return true
				}
			}
		}
	}
	return false
}

func (db *db) mapType(t, n string) map[string]string {
	db.RLock()
	defer db.RUnlock()
	switch t {
	case "script":
		{
			for name, labels := range db.scripts {
				if name == n {
					m := mapLabels(labels)
					m["name"] = n
					return m
				}
			}
		}
	case "host":
		{
			for name, labels := range db.hosts {
				if name == n {
					m := mapLabels(labels)
					m["name"] = n
					return m
				}
			}
		}
	case "group":
		{
			for name, labels := range db.groups {
				if name == n {
					m := mapLabels(labels)
					m["name"] = n
					return m
				}
			}
		}
	case "file":
		{
			for name, labels := range db.files {
				if name == n {
					m := mapLabels(labels)
					m["name"] = n
					return m
				}
			}
		}
	}
	return map[string]string{}
}

func (db *db) getFile(revision string) ([]byte, error) {
	dir := db.dir + "/db/" + revision[:1] + "/"
	files, err := ioutil.ReadDir(dir)
	b := []byte{}
	if err != nil {
		// log.Println(err)
		return []byte{}, err
	}
	for _, f := range files {
		if strings.HasPrefix(f.Name(), revision) {
			b, err = ioutil.ReadFile(dir+f.Name())
			if err != nil {
				return []byte{}, err
			}
		}
	}
	return b, nil
}