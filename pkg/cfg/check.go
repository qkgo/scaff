package cfg

import "sync"

var CheckPermissionCodeServerStateMap = sync.Map{}

var CheckServerStateMap = map[string]bool{}
