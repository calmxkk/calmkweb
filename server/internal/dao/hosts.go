// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"calmk8s/internal/dao/internal"
)

// internalHostsDao is internal type for wrapping internal DAO implements.
type internalHostsDao = *internal.HostsDao

// hostsDao is the data access object for table ck8s_hosts.
// You can define custom methods on it to extend its functionality as you wish.
type hostsDao struct {
	internalHostsDao
}

var (
	// Hosts is globally public accessible object for table ck8s_hosts operations.
	Hosts = hostsDao{
		internal.NewHostsDao(),
	}
)

// Fill with you ideas below.