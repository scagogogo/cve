package cve

import (
	"fmt"
	"reflect"
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

func TestParseCveRange(t *testing.T) {
	type args struct {
		rangeExpr string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "range with 'to' keyword",
			args: args{
				rangeExpr: "CVE-2022-12345 to CVE-2022-12350",
			},
			want: []string{
				"CVE-2022-12345", "CVE-2022-12346", "CVE-2022-12347",
				"CVE-2022-12348", "CVE-2022-12349", "CVE-2022-12350",
			},
		},
		{
			name: "range with double dots",
			args: args{
				rangeExpr: "CVE-2022-12345..12347",
			},
			want: []string{"CVE-2022-12345", "CVE-2022-12346", "CVE-2022-12347"},
		},
		{
			name: "range with dash separator",
			args: args{
				rangeExpr: "CVE-2022-12345-12347",
			},
			want: []string{"CVE-2022-12345", "CVE-2022-12346", "CVE-2022-12347"},
		},
		{
			name: "single CVE (same start and end)",
			args: args{
				rangeExpr: "CVE-2022-12345..12345",
			},
			want: []string{"CVE-2022-12345"},
		},
		{
			name: "invalid format - no range",
			args: args{
				rangeExpr: "CVE-2022-12345",
			},
			want: nil,
		},
		{
			name: "invalid format - reversed range",
			args: args{
				rangeExpr: "CVE-2022-12350..12345",
			},
			want: nil,
		},
		{
			name: "case insensitive",
			args: args{
				rangeExpr: "cve-2022-12345 to cve-2022-12347",
			},
			want: []string{"CVE-2022-12345", "CVE-2022-12346", "CVE-2022-12347"},
		},
		{
			name: "empty string",
			args: args{
				rangeExpr: "",
			},
			want: nil,
		},
		{
			// 起始序列号超出 int 范围，Atoi(startSeq) 失败 → 返回 nil
			// 覆盖 generate.go:152-154 的 startSeq err 分支
			// rangeRegex 能匹配（\d+ 不限长度），但 strconv.Atoi 溢出
			name: "start sequence overflow",
			args: args{
				rangeExpr: "CVE-2022-99999999999999999999999999..12345",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseCveRange(tt.args.rangeExpr)
			if len(got) == 0 && (tt.want == nil || len(tt.want) == 0) {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseCveRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsCvesConsecutive(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "consecutive sequence",
			args: args{
				a: "CVE-2022-12345",
				b: "CVE-2022-12346",
			},
			want: true,
		},
		{
			name: "reverse order consecutive",
			args: args{
				a: "CVE-2022-12346",
				b: "CVE-2022-12345",
			},
			want: true,
		},
		{
			name: "not consecutive - gap",
			args: args{
				a: "CVE-2022-12345",
				b: "CVE-2022-12347",
			},
			want: false,
		},
		{
			name: "different years",
			args: args{
				a: "CVE-2022-12345",
				b: "CVE-2023-12345",
			},
			want: false,
		},
		{
			name: "same CVE",
			args: args{
				a: "CVE-2022-12345",
				b: "CVE-2022-12345",
			},
			want: false,
		},
		{
			name: "invalid first CVE",
			args: args{
				a: "not-a-cve",
				b: "CVE-2022-12346",
			},
			want: false,
		},
		{
			name: "invalid second CVE",
			args: args{
				a: "CVE-2022-12345",
				b: "not-a-cve",
			},
			want: false,
		},
		{
			// 年份有效（2022，非 0）但序列号为 0 → 覆盖 generate.go:215-217 的 seq==0 分支
			// CVE-2022-0 的 ExtractCveSeqAsInt 返回 0，能过年份检查但被 seq==0 拦截
			name: "zero sequence first CVE",
			args: args{
				a: "CVE-2022-0",
				b: "CVE-2022-12346",
			},
			want: false,
		},
		{
			name: "zero sequence second CVE",
			args: args{
				a: "CVE-2022-12345",
				b: "CVE-2022-0",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsCvesConsecutive(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("IsCvesConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}
