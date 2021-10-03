// Copyright 2016 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package varflag

import (
	"time"

	"github.com/mkungla/vars/v5"
)

// Duration returns new duration flag. Argument "a" can be any nr of aliases.
func Duration(name string, aliases ...string) (*DurationFlag, error) {
	c, err := newCommon(name, aliases...)
	if err != nil {
		return nil, err
	}
	f := &DurationFlag{val: time.Duration(0), Common: *c}
	f.variable, _ = vars.NewTyped(name, "", vars.TypeString)
	return f, nil
}

// Parse duration flag.
func (f *DurationFlag) Parse(args []string) (bool, error) {
	return f.parse(args, func(vv []vars.Variable) (err error) {
		if len(vv) > 0 {
			f.variable = vv[0]
			val, err := time.ParseDuration(vv[0].String())
			if err != nil {
				return err
			}
			f.val = val
		}
		return err
	})
}

// Get duration flag value, it returns default value if not present
// or 0 if default is also not set.
func (f *DurationFlag) Value() time.Duration {
	return f.val
}

// Set default value for duration flag.
func (f *DurationFlag) Default(def ...time.Duration) vars.Variable {
	if len(def) > 0 && f.defval.Empty() {
		f.defval = vars.New(f.name, def[0])
		f.val = def[0]
	}
	return f.defval
}

// Unset the bool flag value.
func (f *DurationFlag) Unset() {
	if !f.defval.Empty() {
		f.variable = f.defval
	} else {
		f.variable, _ = vars.NewTyped(f.name, "false", vars.TypeString)
	}
	f.isPresent = false
	val, _ := time.ParseDuration(f.defval.String())
	f.val = val
}
