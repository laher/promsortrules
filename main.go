package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"

	"github.com/prometheus/prometheus/promql"
	"github.com/prometheus/prometheus/util/cli"
)

// SortRulesCmd parses rule files, sorts the rules, and sends them to stdout
func SortRulesCmd(t cli.Term, args ...string) int {
	if len(args) == 0 {
		t.Infof("usage: promrulesort <files>")
		return 2
	}
	failed := false

	for _, arg := range args {
		if rules, err := sortRules(t, arg); err != nil {
			t.Errorf("  FAILED: %s", err)
			failed = true
		} else {
			//t.Infof("  SUCCESS: %d rules found", n)
			ruleStrings := make([]string, len(rules))
			for _, rule := range rules {
				ruleStrings = append(ruleStrings, rule.String())
			}
			sort.Strings(ruleStrings)
			for _, rule := range ruleStrings {
				fmt.Printf("%s\n\n", rule)
			}
		}
	}
	if failed {
		return 1
	}
	return 0
}

func sortRules(t cli.Term, filename string) (promql.Statements, error) {
	var rules promql.Statements
	if stat, err := os.Stat(filename); err != nil {
		return rules, fmt.Errorf("cannot get file info")
	} else if stat.IsDir() {
		return rules, fmt.Errorf("is a directory")
	}

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return rules, err
	}

	rules, err = promql.ParseStmts(string(content))
	if err != nil {
		return rules, err
	}

	return rules, nil
}

func main() {
	app := cli.NewApp("promtool")

	app.Register("sort-rules", &cli.Command{
		Desc: "Sort rules files alphabetically",
		Run:  SortRulesCmd,
	})

	t := cli.BasicTerm(os.Stdout, os.Stderr)
	args := append([]string{"sort-rules"}, os.Args[1:]...)
	os.Exit(app.Run(t, args...))
}
