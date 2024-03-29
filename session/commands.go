package session

import (
	"bytes"
	"errors"

	"github.com/jedib0t/go-pretty/v6/table"
)

type command struct {
	name     string
	single   func()
	with_arg func(string) // initialized here --> end msg about potential cycle
	usage    []struct {
		args string
		msg  string
	}
}

type commandSet struct {
	m map[string]*command
	l []*command
}

func (c *commandSet) register(cmd_string string, cmd *command) error {
	_, ok := c.m[cmd_string]
	if ok {
		return errors.New("command " + cmd_string + " has already been defined!\n")
	}

	c.m[cmd_string] = cmd

	for _, command := range c.l {
		if command == cmd {
			return nil
		}
	}

	c.l = append(c.l, cmd)
	return nil
}

func (c commandSet) get_exec_single(cmd_string string) (func(), bool) {
	cmd, ok := c.m[cmd_string]
	if !ok {
		return nil, false
	}
	if cmd.single == nil {
		return nil, false
	}
	return cmd.single, true
}

func (c commandSet) get_exec_with_arg(cmd_string string) (func(string), bool) {
	cmd, ok := c.m[cmd_string]
	if !ok {
		return nil, false
	}
	if cmd.with_arg == nil {
		return nil, false
	}
	return cmd.with_arg, true
}

func (c commandSet) usage(cmd_string string) (string, bool) {
	command, ok := c.m[cmd_string]
	if !ok {
		return "", false
	}

	var out bytes.Buffer

	t := table.NewWriter()
	t.SetOutputMirror(&out)

	usage := command.usage
	if len(usage) == 0 {
		t.AppendRow([]interface{}{command.name, "no usage message provided"})
	} else {
		for i, msg := range usage {
			if i == 0 {
				t.AppendRow([]interface{}{command.name, msg.args, msg.msg})
			} else {
				t.AppendRow([]interface{}{"", msg.args, msg.msg})
			}
		}

	}
	t.Render()
	return out.String(), true
}

func (c commandSet) menu() string {

	var out bytes.Buffer

	t := table.NewWriter()
	t.SetOutputMirror(&out)

	t.AppendHeader(table.Row{"Name", "", "Usage"})
	t.AppendSeparator()

	for _, command := range c.l {
		name := command.name
		usage := command.usage
		if len(usage) == 0 {
			t.AppendRow([]interface{}{name, "no usage message provided"})
		} else {
			for i, msg := range usage {
				if i == 0 {
					t.AppendRow([]interface{}{name, msg.args, msg.msg})
				} else {
					t.AppendRow([]interface{}{"", msg.args, msg.msg})
				}
			}

		}
	}
	t.Render()
	return out.String()
}

func newCommandSet() *commandSet {
	m := make(map[string]*command)
	l := make([]*command, 0)

	return &commandSet{m: m, l: l}
}

var commands *commandSet

func (s *Session) init_commands() error {

	commands = newCommandSet()

	// help
	c_help := &command{
		name:     "h[elp]",
		single:   s.exec_help_all,
		with_arg: s.exec_help,
		usage: []struct {
			args string
			msg  string
		}{
			{"~", "list all commands with usage"},
			{"~ <cmd>", "print usage command <cmd>"},
		},
	}

	if err := commands.register("help", c_help); err != nil {
		return err
	}
	if err := commands.register("h", c_help); err != nil {
		return err
	}

	//quit
	c_quit := &command{
		name:   "q[uit]",
		single: s.exec_quit,
		usage: []struct {
			args string
			msg  string
		}{
			{"~", "quit the session"},
		},
	}

	if err := commands.register("quit", c_quit); err != nil {
		return err
	}
	if err := commands.register("q", c_quit); err != nil {
		return err
	}

	// clearscreen

	exec_clearscreen := s.f_exec_clearscreen()
	c_clearscreen := &command{
		name:   "cl[earscreen]",
		single: exec_clearscreen,
		usage: []struct {
			args string
			msg  string
		}{
			{"~", "clear the terminal screen"},
		},
	}
	if err := commands.register("clearscreen", c_clearscreen); err != nil {
		return err
	}
	if err := commands.register("cl", c_clearscreen); err != nil {
		return err
	}

	// environment: list
	c_list := &command{
		name:   "l[ist]",
		single: s.exec_list,
		usage: []struct {
			args string
			msg  string
		}{
			{"~", "list all identifiers in the environment alphabetically\n\t with types and values"},
		},
	}
	if err := commands.register("list", c_list); err != nil {
		return err
	}
	if err := commands.register("l", c_list); err != nil {
		return err
	}

	// environment: clear
	c_clear := &command{
		name:   "c[lear]",
		single: s.exec_clear,
		usage: []struct {
			args string
			msg  string
		}{
			{"~", "clear the environment"},
		},
	}

	if err := commands.register("clear", c_clear); err != nil {
		return err
	}
	if err := commands.register("c", c_clear); err != nil {
		return err
	}

	// paste
	c_paste := &command{
		name:     "paste",
		with_arg: s.exec_paste,
		//	single:   s.exec_paste_empty_arg,
		usage: []struct {
			args string
			msg  string
		}{
			{"~ <input>", "evaluate multiline <input> (terminated by blank line)"},
		},
	}
	commands.register("paste", c_paste)

	// level: expression
	c_expr := &command{
		name:     "expr[ession]",
		with_arg: s.exec_expression,
		usage: []struct {
			args string
			msg  string
		}{
			{"~ <input>", "expect <input> to be an expression"},
		},
	}

	if err := commands.register("expression", c_expr); err != nil {
		return err
	}
	if err := commands.register("expr", c_expr); err != nil {
		return err
	}
	// level: statement
	c_stmt := &command{
		name:     "stmt|statement",
		with_arg: s.exec_statement,
		usage: []struct {
			args string
			msg  string
		}{
			{"~ <input>", "expect <input> to be a statement"},
		},
	}
	if err := commands.register("statement", c_stmt); err != nil {
		return err
	}
	if err := commands.register("stmt", c_stmt); err != nil {
		return err
	}

	// level: program
	c_prog := &command{
		name:     "prog[ram]",
		with_arg: s.exec_program,
		usage: []struct {
			args string
			msg  string
		}{
			{"~ <input>", "expect <input> to be a program"},
		},
	}

	if err := commands.register("program", c_prog); err != nil {
		return err
	}
	if err := commands.register("prog", c_prog); err != nil {
		return err
	}

	// process: parse
	c_parse := &command{
		name:     "p[arse]",
		with_arg: s.exec_parse,
		usage: []struct {
			args string
			msg  string
		}{
			{"~ <input>", "print string representation of ast <input> is parsed to\n\t --> Node-method: String() string"},
		},
	}

	if err := commands.register("parse", c_parse); err != nil {
		return err
	}
	if err := commands.register("p", c_parse); err != nil {
		return err
	}

	// process: parsetree: TODO
	c_parsetree := &command{
		name:     "p[arse]tree",
		with_arg: s.exec_parsetree,
		usage: []struct {
			args string
			msg  string
		}{
			{"~ <input>", "print tree representation of <input>' ast\n\t to all set displays "},
		},
	}

	if err := commands.register("parsetree", c_parsetree); err != nil {
		return err
	}
	if err := commands.register("ptree", c_parsetree); err != nil {
		return err
	}

	// process: eval
	c_eval := &command{
		name:     "e[val]",
		with_arg: s.exec_eval,
		usage: []struct {
			args string
			msg  string
		}{
			{"~ <input>", "print value of object <input> evaluates to\n\t --> Object-method: Inspect() string"},
		},
	}
	if err := commands.register("eval", c_eval); err != nil {
		return err
	}
	if err := commands.register("e", c_eval); err != nil {
		return err
	}

	// process: type
	c_type := &command{
		name:     "t[ype]",
		with_arg: s.exec_type,
		usage: []struct {
			args string
			msg  string
		}{
			{"~ <input>", "show objecttype <input> evaluates to"},
		},
	}
	if err := commands.register("type", c_type); err != nil {
		return err
	}
	if err := commands.register("t", c_type); err != nil {
		return err
	}

	// process: trace
	c_trace := &command{
		name:     "tr[ace]",
		with_arg: s.exec_trace,
		usage: []struct {
			args string
			msg  string
		}{
			{"~ <input>", "show evaluation trace interactively step by step"},
		},
	}
	if err := commands.register("trace", c_trace); err != nil {
		return err
	}
	if err := commands.register("tr", c_trace); err != nil {
		return err
	}
	// process: evaltree
	c_evaltree := &command{
		name:     "e[val]tree",
		with_arg: s.exec_evaltree,
		usage: []struct {
			args string
			msg  string
		}{
			{"~ <input>", "print annotated tree representation of <input>'s ast\n\t to all set displays "},
		},
	}
	if err := commands.register("evaltree", c_evaltree); err != nil {
		return err
	}
	if err := commands.register("etree", c_evaltree); err != nil {
		return err
	}

	// settings: settings
	c_settings := &command{
		name:   "settings",
		single: s.exec_settings,
		usage: []struct {
			args string
			msg  string
		}{
			{"~", "list all settings with their current and default values"},
		},
	}
	if err := commands.register("settings", c_settings); err != nil {
		return err
	}

	// settings: set
	c_set := &command{
		name:     "set",
		with_arg: s.exec_set,
		usage: []struct {
			args string
			msg  string
		}{
			{"~ prompt <prompt>", "set prompt string to <prompt>"},
			{"~ paste", "enable multiline support"},
			{"~ level <l>", "<l> must be: p[rogram], s[tatement], e[xpression]"},
			{"~ process <p>", "<p> must be: p[arse], p[arse]tree, e[val], e[val]tree,\n\t [t]ype, [tr]ace"},
			{"~ logs <+|-l_0...+|-l_n>", "<l_i> must be: p[arse]tree, e[val]tree, [t]ype, [tr]ace"},
			{"~ displays <+|-d_0...+|-d_n>", "<d_i> must be: c[ons[ole]], p[df]"},
			{"~ verbosity <v>", "<v> must be 0, 1, 2"},
			{"~ inclToken", "include tokens in representations of asts"},
			{"~ inclEnv", "include environments in representations of asts"},
			{"~ pfile <f>", "set file for parsetree to <f>"},
			{"~ efile <f>", "set file for evaltree to <f>"},
			{"~ goObjType", "display Go type instead of Monkey type"},
		},
	}
	if err := commands.register("set", c_set); err != nil {
		return err
	}

	// settings: reset
	c_reset := &command{
		name:     "reset",
		single:   s.exec_reset_all,
		with_arg: s.exec_reset,
		usage: []struct {
			args string
			msg  string
		}{
			{"~", "reset all settings"},
			{"~ <setting>", "set <setting> to default value\n\t for an overview consult :settings and/or :h set"},
		},
	}
	if err := commands.register("reset", c_reset); err != nil {
		return err
	}

	// settings: unset
	c_unset := &command{
		name:     "unset",
		with_arg: s.exec_unset,
		usage: []struct {
			args string
			msg  string
		}{
			{"~ <setting>", "set boolean <setting> to false\n\t for an overview consult :settings and/or :h set"},
			//	{"~ logparse", "don't additionally output ast-string"},
			//	{"~ logtype", "don't additionally output objecttype"},
			//	{"~ paste", "disable multiline support"},
			//incltoken logtrace
		},
	}
	if err := commands.register("unset", c_unset); err != nil {
		return err
	}
	return nil
}
