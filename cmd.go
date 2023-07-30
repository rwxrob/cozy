package cozy

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
	////"github.com/rwxrob/term"
	//"github.com/rwxrob/to"
	//"github.com/rwxrob/vars"
)

var Cmd = &Z.Cmd{
	Name: `cozy`,
	Commands: []*Z.Cmd{
		help.Cmd,
	},
	Call: func(_ *Z.Cmd, _ ...string) error {
		// TODO implement
		println(`foo`)
		return nil
	},
	Version:     `v0.0.0`,
	Copyright:   `(c) Robert S. Muhlestein <rob@rwx.gg> (rwxrob.tv)`,
	License:     `Apache-2.0`,
	Source:      `https://github.com/rwxrob/cozy`,
	Issues:      `https://github.com/rwxrob/cozy/issues`,
	Summary:     help.S(_cozy),
	Description: help.D(_cozy),
}
