// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameMinerM = "gateway_miner"

// MinerM mapped from table <gateway_miner>
type MinerM struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键 ID" json:"id"`                     // 主键 ID
	Namespace   string    `gorm:"column:namespace;not null;comment:命名空间" json:"namespace"`                             // 命名空间
	Name        string    `gorm:"column:name;not null;comment:矿机名" json:"name"`                                        // 矿机名
	DisplayName string    `gorm:"column:displayName;not null;comment:矿机展示名" json:"displayName"`                        // 矿机展示名
	Phase       string    `gorm:"column:phase;not null;comment:矿机状态" json:"phase"`                                     // 矿机状态
	MinerType   string    `gorm:"column:minerType;not null;comment:矿机机型" json:"minerType"`                             // 矿机机型
	ChainName   string    `gorm:"column:chainName;not null;comment:矿机所属的区块链名" json:"chainName"`                        // 矿机所属的区块链名
	CPU         int32     `gorm:"column:cpu;not null;comment:矿机 CPU 规格" json:"cpu"`                                    // 矿机 CPU 规格
	Memory      int32     `gorm:"column:memory;not null;comment:矿机内存规格" json:"memory"`                                 // 矿机内存规格
	CreatedAt   time.Time `gorm:"column:createdAt;not null;default:current_timestamp;comment:创建时间" json:"createdAt"`   // 创建时间
	UpdatedAt   time.Time `gorm:"column:updatedAt;not null;default:current_timestamp;comment:最后修改时间" json:"updatedAt"` // 最后修改时间
}

// TableName MinerM's table name
func (*MinerM) TableName() string {
	return TableNameMinerM
}
