// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameChainM = "api_chain"

// ChainM mapped from table <api_chain>
type ChainM struct {
	ID                     int64     `gorm:"column:id;type:bigint(20) unsigned;primaryKey;autoIncrement:true;comment:主键 ID" json:"id"`                             // 主键 ID
	Namespace              string    `gorm:"column:namespace;type:varchar(253);not null;uniqueIndex:uniq_namespace_name,priority:1;comment:命名空间" json:"namespace"` // 命名空间
	Name                   string    `gorm:"column:name;type:varchar(253);not null;uniqueIndex:uniq_namespace_name,priority:2;comment:区块链名" json:"name"`           // 区块链名
	DisplayName            string    `gorm:"column:display_name;type:varchar(253);not null;comment:区块链展示名" json:"display_name"`                                    // 区块链展示名
	MinerType              string    `gorm:"column:miner_type;type:varchar(16);not null;comment:区块链矿机机型" json:"miner_type"`                                        // 区块链矿机机型
	Image                  string    `gorm:"column:image;type:varchar(253);not null;comment:区块链镜像 ID" json:"image"`                                                // 区块链镜像 ID
	MinMineIntervalSeconds int32     `gorm:"column:min_mine_interval_seconds;type:int(8);not null;comment:矿机挖矿间隔" json:"min_mine_interval_seconds"`                // 矿机挖矿间隔
	CreatedAt              time.Time `gorm:"column:created_at;type:datetime;not null;default:current_timestamp();comment:创建时间" json:"created_at"`                  // 创建时间
	UpdatedAt              time.Time `gorm:"column:updated_at;type:datetime;not null;default:current_timestamp();comment:最后修改时间" json:"updated_at"`                // 最后修改时间
}

// TableName ChainM's table name
func (*ChainM) TableName() string {
	return TableNameChainM
}
