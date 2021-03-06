package angel

import (
	"fmt"
	"strconv"
)

//QueryStartupRoles queries /startup_roles for all startup roles associated with the user with the given id
func (c AngelClient) QueryStartupRoles(id int64, id_type int) (startuproles []StartupRole, err error) {
	var tmp struct {
		Startup_roles []StartupRole
	}
	switch id_type {
	case UserId:
		{
			err = c.execQueryThrottled("/startup_roles", GET, map[string]string{"user_id": strconv.FormatInt(id, 10)}, &tmp)
		}
	case StartupId:
		{
			err = c.execQueryThrottled("/startup_roles", GET, map[string]string{"startup_id": strconv.FormatInt(id, 10)}, &tmp)
		}
	default:
		return nil, fmt.Errorf("invalid id_type provided")
	}
	startuproles = tmp.Startup_roles
	return
}
