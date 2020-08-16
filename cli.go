package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
	"io/ioutil"
)


const (
	lsHeader string = `╔%s╕ %s
║ %s └%s╮
╿ `+undStart+`%s`+undStop+`  :  `+undStart+`Labels`+undStop+`
`
	undStart string = "\033[4m"
	undStop  string = "\033[0m"
)

func cli() error {
	labels, err := getArgs()
	if err != nil {
		return err
	}
	m := mapLabels(labels)
	switch m["action"] {
	case "set":
		{
			return set(labels)
		}
	case "show":
		{
			return show(labels)
		}
	case "delete":
		{
			return del(labels)
		}
	case "execute":
		{
			// err = exec(labels[1:])
			// tmp()
			return execute(labels)
		}
	case "copy":
		{
			tmp()
		}
	default:
		{
			help()
		}
	}
	return nil
}

// func action
func tmp() {
}

// TODO: this whole thing needs a serious cleanup
func getArgs() ([]label, error) {
	labels := []label{}
	seen := map[string]bool{}
	action := ""
	t := ""
	args := os.Args[1:]
	for i, arg := range args {
		if i == 0 {
			switch arg {
			case "make", "set", "add":
				{
					arg = "set"
				}
			case "show", "list", "get", "ls":
				{
					arg = "show"
				}
			case "rm", "remove", "del", "delete":
				{
					if len(os.Args) < 4 {
						log.Fatalln("need name to delete")
					}
					arg = "delete"
				}
			case "exec", "run", "execute", "-n", "-g", "-su", "-sq":
				{
					arg = "execute"
				}
			case "copy", "cp":
				{
					arg = "copy"
				}
			default:
				{}
			}
			labels = append(labels, label{Key: "action", Value: arg})
			action = arg
			continue
		}
		if action == "execute" {
			lbs := getExecuteArgs(args)
			return lbs, nil
		} else if i == 1 {
			switch arg {
			case "host", "hosts", "node", "nodes":
				{
					arg = "host"
				}
			case "group", "groups":
				{
					arg = "group"
				}
			case "script", "scripts":
				{
					arg = "script"
				}
			case "file", "files":
				{
					arg = "file"
				}
			default:
				{
					log.Fatalln("something went wrong: type")
				}
			}
			labels = append(labels, label{Key: "type", Value: arg})
			t = arg
			continue
		}
		if i == 2 && action == "show" && t == "group" {
			db, err := openDB()
			if err != nil {
				return labels, err
			}
			q := []label{
				{Key: "action", Value: "show"},
				{Key: "type", Value: t},
				{Key: "name", Value: arg},
			}
			result, err := db.get(q)
			if err == nil {
				lbs := delKey(result[0], "name")
				lbs = delKey(lbs, "type")
				labels = append(lbs, label{Key: "type", Value: "host"})
			}
			break
		}
		if i == 2 && action != "show" {
			labels = append(labels, label{Key: "name", Value: arg})
			// continue
		}
		if i == 2 && action == "show" && t == "host" && !strings.Contains(arg, "=") {
			arg = "name="+arg
			// continue
			// break
		}
		if strings.Contains(arg, "=") {
			s := strings.Split(arg, "=")
			if !seen[s[0]] {
				labels = append(labels, label{Key: s[0], Value: s[1]})
			}
			seen[s[0]] = true
			continue
		}
		if i == 3 && (t == "script" || t == "file") && action == "set" {
			labels = append(labels, label{Key: "path", Value: arg})
			continue
		}
	}
	// log.Println(labels)
	m := mapLabels(labels)
	err := *new(error)
	if m["action"] == "set" && (m["type"] == "file" || m["type"] == "script") {
		labels, err = replaceLabels(labels)
	}
	return labels, err
}

func cliPrint(gs ...[]label) {
	block1 := ""
	block2 := ""
	info := ""
	m := mapLabels(gs[0])
	// log.Println(gs)
	switch m["action"] {
	case "set":
		{
			info = "Set!"
			switch m["type"] {
			case "host":
				{
					block1 = strings.Repeat("═", 6)
					block2 = strings.Repeat("─", 10)
				}
			case "group":
				{
					block1 = strings.Repeat("═", 7)
					block2 = strings.Repeat("─", 9)
				}
			case "script":
				{
					block1 = strings.Repeat("═", 8)
					block2 = strings.Repeat("─", 8)
				}
			case "file":
				{
					block1 = strings.Repeat("═", 6)
					block2 = strings.Repeat("─", 10)
				}
			}
		}
	case "show":
		{
			total := 0
			sm := mapLabels(gs[0])
			if !(len(gs) == 1 && sm["name"] == "NONE") {
				total = len(gs)
			}
			info = "Total: " + strconv.Itoa(total)
			switch m["type"] {
			case "host":
				{
					block1 = strings.Repeat("═", 7)
					block2 = strings.Repeat("─", 9)
				}
			case "group":
				{
					block1 = strings.Repeat("═", 8)
					block2 = strings.Repeat("─", 8)
				}
			case "script":
				{
					block1 = strings.Repeat("═", 9)
					block2 = strings.Repeat("─", 7)
				}
			case "file":
				{
					block1 = strings.Repeat("═", 7)
					block2 = strings.Repeat("─", 9)
				}
			}
			if !strings.HasSuffix(m["type"], "s") {
				m["type"] = m["type"] + "s"
			}
		}
	case "delete":
		{
			info = "Deleted!"
			sm := mapLabels(gs[0])
			if sm["name"] == "NONE" {
				info = "Not found"
			}
			switch m["type"] {
			case "host":
				{
					block1 = strings.Repeat("═", 6)
					block2 = strings.Repeat("─", 10)
				}
			case "group":
				{
					block1 = strings.Repeat("═", 7)
					block2 = strings.Repeat("─", 9)
				}
			case "script":
				{
					block1 = strings.Repeat("═", 8)
					block2 = strings.Repeat("─", 8)
				}
			case "file":
				{
					block1 = strings.Repeat("═", 6)
					block2 = strings.Repeat("─", 10)
				}
			}
		}
	case "execute":
		{
			// err = exec(labels[1:])
			tmp()
		}
	case "copy":
		{
			tmp()
		}
	default:
		{
			log.Fatalln("something went wrong: action")
		}
	}
	fmt.Printf(
		lsHeader,
		block1,
		info,
		strings.Title(strings.ToLower(m["type"])),
		block2,
		"Name",
	)
	n := 0
	for _, l := range gs {
		ml := mapLabels(l)
		front := "├"
		if n == len(gs) - 1 {
			front = "└"
		}
		n++
		fmt.Println(front, ml["name"], ":", printLabels(l))
	}
}

func printLabels(ls []label) string {
	s := ""
	seen := map[string]bool{}
	for _, l := range ls {
		if l.Key == "action" || l.Key == "type" || l.Key == "name" || l.Key == "publickey" {
			continue
		}
		if seen[l.Key] {
			continue
		} else {
			seen[l.Key] = true
			k := l.Key
			v := l.Value
			if k == "revision" || k == "hostkey" {
				if len(v) > 7 {
					v = v[:8]
				}
			}
			s += k + "=" + v + " "
		}
	}
	return strings.TrimSpace(s)
}

/// ISSUE HERE
func replaceLabels(labels []label) ([]label, error) {
	m := mapLabels(labels)
	// log.Println("MAP:", m)
	// log.Println("LABELS:", labels)
	revision, err := addFile(m["path"])
	if err != nil {
		log.Fatalln(err)
	}
	// m["revision"] = revision
	if m["type"] == "script" {
		lang, err := getSheBang(m["path"])
		if err != nil {
			return []label{}, err
		}
		// m["lang"] = lang
		labels = append([]label{
				{Key: "lang", Value: lang},
				{Key: "revision", Value: revision},
			},
			labels...,
		)
		// log.Println("predel:", labels)
		labels = delKey(labels, "path")
		// log.Println("postdel:", labels)
	}
	// log.Println("replaced:", labels)
	return labels, nil
}

func getSheBang(path string) (string, error) {
	lang := ""
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	for i, s := range strings.Split(strings.Split(string(b), "\n")[0], " ") {
		if i == 0 && !strings.Contains(s, `#!/`) {
			break
		}
		if strings.Contains(s, "/bin/env") {
			continue
		}
		if strings.Contains(s, "/") {
			spl := strings.Split(s, "/")
			lang = strings.TrimSpace(spl[len(spl)-1])
		} else {
			lang = strings.TrimSpace(s)
		}				
	}
	return lang, nil
}

func getScriptHashBang(args []string) ([]string, error) {
	rev := ""
	lang := ""
	for _, arg := range args {
		if strings.Contains(arg, "revision=") && len(arg) == 17 {
			rev = strings.Split(arg, "=")[1]
		}
	}
	if rev == "" {
		return args, fmt.Errorf("%s", "revision not found")
	}
	dir := getDir() +"/db/" + rev[:1] + "/" + rev[1:2]
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return args, err
	}
	for _, fi := range files {
		if strings.HasPrefix(fi.Name(), rev) {
			fBytes, err := ioutil.ReadFile(dir+"/"+fi.Name())
			if err != nil {
				log.Println(err)
				return args, err
			}
			for i, s := range strings.Split(strings.Split(string(fBytes), "\n")[0], " ") {
				if i == 0 && !strings.Contains(s, `#!/`) {
					break
				}
				if strings.Contains(s, "/bin/env") {
					continue
				}
				if strings.Contains(s, "/") {
					spl := strings.Split(s, "/")
					lang = strings.TrimSpace(spl[len(spl)-1])
				} else {
					lang = strings.TrimSpace(s)
				}				
			}
		}
	}
	if lang != "" {
		args = append(
			args[:3],
			append(
				[]string{"lang="+lang},
				args[3:]...,
			)...,
		)
	}
	return args, nil
}
