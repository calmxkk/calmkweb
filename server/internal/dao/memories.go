// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"calmk8s/internal/dao/internal"
)

// internalMemoriesDao is internal type for wrapping internal DAO implements.
type internalMemoriesDao = *internal.MemoriesDao

// memoriesDao is the data access object for table ck8s_memories.
// You can define custom methods on it to extend its functionality as you wish.
type memoriesDao struct {
	internalMemoriesDao
}

var (
	// Memories is globally public accessible object for table ck8s_memories operations.
	Memories = memoriesDao{
		internal.NewMemoriesDao(),
	}
)

// Fill with you ideas below.
