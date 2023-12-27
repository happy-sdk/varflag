// SPDX-License-Identifier: Apache-2.0
//
// Copyright © 2023 The Happy Authors

// Package cron implements a cron spec parser and runner.
// Deprecated: This module is no longer maintained.
// Development has moved to github.com/happy-sdk/happy-go/scheduling/cron.
// Users are encouraged to use the new module location for future updates and bug fixes.
package varflag

import "time"

const (
	// Deprecated is a marker for deprecated code.
	Deprecated = true

	// DeprecatedBy is the name entity who deprecated this package.
	DeprecatedBy = "The Happy Authors"

	// NewLocation is the new location of this package.
	NewLocation = "github.com/happy-sdk/happy-go/vars/varflag"
)

// DeprecatedAt is the date when this package was deprecated.
func DeprecatedAt() time.Time {
	return time.Date(2023, time.December, 27, 12, 21, 0, 0, time.UTC)
}
