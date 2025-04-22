package cve

import (
	"fmt"
	"regexp"
	"testing"
	"time"
)

func TestGenerateCve(t *testing.T) {
	type args struct {
		year int
		seq  int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "standard CVE",
			args: args{
				year: 2022,
				seq:  1234,
			},
			want: "CVE-2022-1234",
		},
		{
			name: "CVE with single digit sequence",
			args: args{
				year: 2021,
				seq:  5,
			},
			want: "CVE-2021-5",
		},
		{
			name: "CVE with large sequence number",
			args: args{
				year: 2020,
				seq:  123456,
			},
			want: "CVE-2020-123456",
		},
		{
			name: "old year CVE",
			args: args{
				year: 1999,
				seq:  100,
			},
			want: "CVE-1999-100",
		},
		{
			name: "future year CVE",
			args: args{
				year: 2099,
				seq:  9999,
			},
			want: "CVE-2099-9999",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateCve(tt.args.year, tt.args.seq); got != tt.want {
				t.Errorf("GenerateCve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateFakeCve(t *testing.T) {
	// 测试生成的假CVE是否符合格式要求
	currentYear := time.Now().Year()
	fakeCve := GenerateFakeCve()

	// 正则表达式检查格式
	pattern := regexp.MustCompile(`^CVE-\d+-\d+$`)
	if !pattern.MatchString(fakeCve) {
		t.Errorf("GenerateFakeCve() = %v, which doesn't match the expected format", fakeCve)
	}

	// 检查年份
	yearStr := fakeCve[4:8]
	if yearStr != fmt.Sprintf("%d", currentYear) {
		t.Errorf("GenerateFakeCve() year = %v, want %v", yearStr, currentYear)
	}

	// 检查序列号是否为5位数
	seqStr := fakeCve[9:]
	seqLen := len(seqStr)
	if seqLen < 5 {
		t.Errorf("GenerateFakeCve() sequence length = %v, want >= 5", seqLen)
	}
}
