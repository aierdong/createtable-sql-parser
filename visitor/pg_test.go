package visitor

import (
	"testing"
)

func TestParsePgPrimaryKey(t *testing.T) {
	tests := []struct {
		name     string
		sql      string
		expected []string // 期望的主键列名列表
	}{
		{
			name: "列级别主键",
			sql: `CREATE TABLE public.test_table (
				id bigint PRIMARY KEY,
				name varchar(100)
			)`,
			expected: []string{"id"},
		},
		{
			name: "表级别主键",
			sql: `CREATE TABLE public.test_table (
				id bigint,
				name varchar(100),
				PRIMARY KEY (id)
			)`,
			expected: []string{"id"},
		},
		{
			name: "复合主键",
			sql: `CREATE TABLE public.test_table (
				id bigint,
				name varchar(100),
				PRIMARY KEY (id, name)
			)`,
			expected: []string{"id", "name"},
		},
		{
			name: "列级别和表级别主键",
			sql: `CREATE TABLE public.test_table (
				id bigint PRIMARY KEY,
				name varchar(100),
				PRIMARY KEY (id, name)
			)`,
			expected: []string{"id", "name"},
		},
		{
			name: "另一种写法",
			sql: `CREATE TABLE public.test_table (
					id int8 NOT NULL,
					"name" varchar(100) NULL,
				CONSTRAINT test_table_pkey PRIMARY KEY (id)
			)`,
			expected: []string{"id"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			table, err := ParsePgSql(tt.sql)
			if err != nil {
				t.Errorf("解析SQL失败: %v", err)
				return
			}

			// 获取实际的主键列
			var primaryKeys []string
			for _, col := range table.Columns {
				if col.IsPrimaryKey {
					primaryKeys = append(primaryKeys, col.Name)
				}
			}

			// 检查主键列数量
			if len(primaryKeys) != len(tt.expected) {
				t.Errorf("主键列数量不匹配: 期望 %d, 实际 %d", len(tt.expected), len(primaryKeys))
				return
			}

			// 检查主键列名
			for _, expected := range tt.expected {
				found := false
				for _, actual := range primaryKeys {
					if actual == expected {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("未找到期望的主键列: %s", expected)
				}
			}
		})
	}
}
