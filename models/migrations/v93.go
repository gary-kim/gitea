// Copyright 2019 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package migrations

import "github.com/go-xorm/xorm"

func addEmailNotificationEnabledToUser(x *xorm.Engine) error {
	// Issue see models/user.go
	type User struct {
		EmailNotificationsEnabled bool `xorm:"DEFAULT true"`
	}

	return x.Sync2(new(User))
}
