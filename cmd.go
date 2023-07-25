package cowork

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
	////"github.com/rwxrob/term"
	//"github.com/rwxrob/to"
	//"github.com/rwxrob/vars"
)

var Cmd = &Z.Cmd{
	Name: `cowork`,
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
	Source:      `https://github.com/rwxrob/cowork`,
	Issues:      `https://github.com/rwxrob/cowork/issues`,
	Summary:     help.S(_cowork),
	Description: help.D(_cowork),
}
