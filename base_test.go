package cve

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	type args struct {
		cve string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "format CVE with mixed case and spaces",
			args: args{
				cve: " cVe-2002-100098  ",
			},
			want: "CVE-2002-100098",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Format(tt.args.cve); got != tt.want {
				t.Errorf("Format() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsCVE(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "standard CVE format",
			args: args{
				text: "cve-2007-199",
			},
			want: true,
		},
		{
			name: "CVE with leading space",
			args: args{
				text: " cve-2007-199",
			},
			want: true,
		},
		{
			name: "CVE with trailing space",
			args: args{
				text: "cve-2007-199 ",
			},
			want: true,
		},
		{
			name: "CVE with mixed case",
			args: args{
				text: "cVe-2007-199",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsCve(tt.args.text); got != tt.want {
				t.Errorf("IsCve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsContainsCVE(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "text containing CVE",
			args: args{
				text: "this text contains cve-2908-10086",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsContainsCve(tt.args.text); got != tt.want {
				t.Errorf("IsContainsCve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplit(t *testing.T) {
	type args struct {
		cve string
	}
	tests := []struct {
		name     string
		args     args
		wantYear string
		wantSeq  string
	}{
		{
			name: "split standard CVE into year and sequence",
			args: args{
				cve: "cve-2007-10086",
			},
			wantYear: "2007",
			wantSeq:  "10086",
		},
		{
			name: "split CVE with invalid format",
			args: args{
				cve: "invalid-format",
			},
			wantYear: "",
			wantSeq:  "",
		},
		{
			name: "split empty string",
			args: args{
				cve: "",
			},
			wantYear: "",
			wantSeq:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotYear, gotSeq := Split(tt.args.cve)
			if gotYear != tt.wantYear {
				t.Errorf("Split() gotYear = %v, want %v", gotYear, tt.wantYear)
			}
			if gotSeq != tt.wantSeq {
				t.Errorf("Split() gotSeq = %v, want %v", gotSeq, tt.wantSeq)
			}
		})
	}
}

func TestIsCveYearOk(t *testing.T) {
	type args struct {
		cve string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid current year",
			args: args{
				cve: fmt.Sprintf("CVE-%d-10086", time.Now().Year()),
			},
			want: true,
		},
		{
			name: "valid past year",
			args: args{
				cve: "CVE-2020-10086",
			},
			want: true,
		},
		{
			name: "future year",
			args: args{
				cve: fmt.Sprintf("CVE-%d-10086", time.Now().Year()+1),
			},
			want: false,
		},
		{
			name: "year before 1999",
			args: args{
				cve: "CVE-1998-10086",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsCveYearOk(tt.args.cve); got != tt.want {
				t.Errorf("IsCveYearOk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsCveYearOkWithCutoff(t *testing.T) {
	type args struct {
		cve    string
		cutoff int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid year within cutoff",
			args: args{
				cve:    "CVE-2020-10086",
				cutoff: 5,
			},
			want: true,
		},
		{
			name: "valid year at exact cutoff",
			args: args{
				cve:    "CVE-2019-10086",
				cutoff: time.Now().Year() - 2019,
			},
			want: true,
		},
		{
			name: "valid year beyond cutoff",
			args: args{
				cve:    "CVE-2000-10086",
				cutoff: 5,
			},
			want: true,
		},
		{
			name: "future year with larger cutoff",
			args: args{
				cve:    "CVE-2099-10086",
				cutoff: 100,
			},
			want: true,
		},
		{
			name: "year before 1999",
			args: args{
				cve:    "CVE-1998-10086",
				cutoff: 100,
			},
			want: false,
		},
		{
			name: "invalid CVE format",
			args: args{
				cve:    "CVE-INVALID-FORMAT",
				cutoff: 0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsCveYearOkWithCutoff(tt.args.cve, tt.args.cutoff); got != tt.want {
				t.Errorf("IsCveYearOkWithCutoff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateCve(t *testing.T) {
	type args struct {
		cve string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid CVE",
			args: args{
				cve: "CVE-2022-1234",
			},
			want: true,
		},
		{
			name: "valid CVE with leading/trailing spaces",
			args: args{
				cve: " CVE-2022-1234 ",
			},
			want: true,
		},
		{
			name: "valid CVE lowercase",
			args: args{
				cve: "cve-2022-1234",
			},
			want: true,
		},
		{
			name: "invalid format - missing prefix",
			args: args{
				cve: "2022-1234",
			},
			want: false,
		},
		{
			name: "invalid format - wrong separator",
			args: args{
				cve: "CVE/2022/1234",
			},
			want: false,
		},
		{
			name: "invalid year - before 1999",
			args: args{
				cve: "CVE-1998-1234",
			},
			want: false,
		},
		{
			name: "invalid year - future beyond current year",
			args: args{
				cve: "CVE-2099-1234", // Note: This might need updating in the future
			},
			want: false,
		},
		{
			name: "invalid sequence - not a number",
			args: args{
				cve: "CVE-2022-ABCD",
			},
			want: false,
		},
		{
			name: "invalid format - extra components",
			args: args{
				cve: "CVE-2022-1234-5",
			},
			want: false,
		},
		{
			name: "invalid format - missing components",
			args: args{
				cve: "CVE-2022",
			},
			want: false,
		},
		{
			name: "invalid year - non-numeric",
			args: args{
				cve: "CVE-YEAR-1234",
			},
			want: false,
		},
		{
			name: "year too large for int conversion",
			args: args{
				cve: "CVE-99999999999999999999-1234",
			},
			want: false,
		},
		{
			name: "sequence too large for int conversion",
			args: args{
				cve: "CVE-2022-99999999999999999999",
			},
			want: false,
		},
		{
			name: "zero sequence number",
			args: args{
				cve: "CVE-2022-0",
			},
			want: false,
		},
		{
			name: "negative sequence number",
			args: args{
				cve: "CVE-2022--1",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateCve(tt.args.cve); got != tt.want {
				t.Errorf("ValidateCve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateCves(t *testing.T) {
	currentYear := time.Now().Year()
	type args struct {
		cveSlice []string
	}
	tests := []struct {
		name string
		args args
		want []CveValidationResult
	}{
		{
			name: "all valid CVEs",
			args: args{
				cveSlice: []string{"CVE-2022-1234", "CVE-2023-5678"},
			},
			want: []CveValidationResult{
				{Cve: "CVE-2022-1234", Valid: true, Reason: ""},
				{Cve: "CVE-2023-5678", Valid: true, Reason: ""},
			},
		},
		{
			name: "mixed valid and invalid",
			args: args{
				cveSlice: []string{"CVE-2022-1234", "not-a-cve", "CVE-1998-1234"},
			},
			want: []CveValidationResult{
				{Cve: "CVE-2022-1234", Valid: true, Reason: ""},
				{Cve: "not-a-cve", Valid: false, Reason: "invalid CVE format"},
				{Cve: "CVE-1998-1234", Valid: false, Reason: "year 1998 is before 1999"},
			},
		},
		{
			name: "empty slice",
			args: args{
				cveSlice: []string{},
			},
			want: nil,
		},
		{
			name: "future year invalid",
			args: args{
				cveSlice: []string{fmt.Sprintf("CVE-%d-1234", currentYear+5)},
			},
			want: []CveValidationResult{
				{Cve: fmt.Sprintf("CVE-%d-1234", currentYear+5), Valid: false, Reason: fmt.Sprintf("year %d is after current year %d", currentYear+5, currentYear)},
			},
		},
		{
			name: "non-numeric sequence",
			args: args{
				cveSlice: []string{"CVE-2022-ABCD"},
			},
			want: []CveValidationResult{
				{Cve: "CVE-2022-ABCD", Valid: false, Reason: "invalid CVE format"},
			},
		},
		{
			name: "zero sequence invalid",
			args: args{
				cveSlice: []string{"CVE-2022-0"},
			},
			want: []CveValidationResult{
				{Cve: "CVE-2022-0", Valid: false, Reason: "sequence number must be positive"},
			},
		},
		{
			// 年份超出 int 范围，Atoi(year) 失败 → 覆盖 base.go:341-345 yearErr 分支
			name: "year overflow invalid",
			args: args{
				cveSlice: []string{"CVE-99999999999999999999999999-1234"},
			},
			want: []CveValidationResult{
				{Cve: "CVE-99999999999999999999999999-1234", Valid: false, Reason: "year is not a valid number"},
			},
		},
		{
			// 序列号超出 int 范围，Atoi(seq) 失败 → 覆盖 base.go:347-351 seqErr 分支
			name: "sequence overflow invalid",
			args: args{
				cveSlice: []string{"CVE-2022-99999999999999999999999999"},
			},
			want: []CveValidationResult{
				{Cve: "CVE-2022-99999999999999999999999999", Valid: false, Reason: "sequence number is not a valid number"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidateCves(tt.args.cveSlice)
			if len(got) == 0 && (tt.want == nil || len(tt.want) == 0) {
				return
			}
			for i := range got {
				if got[i].Cve != tt.want[i].Cve || got[i].Valid != tt.want[i].Valid || got[i].Reason != tt.want[i].Reason {
					t.Errorf("ValidateCves()[%d] = %+v, want %+v", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestFilterValidCves(t *testing.T) {
	type args struct {
		cveSlice []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "filter out invalid CVEs",
			args: args{
				cveSlice: []string{"CVE-2022-1234", "not-a-cve", "CVE-1998-1234", "CVE-2023-5678"},
			},
			want: []string{"CVE-2022-1234", "CVE-2023-5678"},
		},
		{
			name: "all valid",
			args: args{
				cveSlice: []string{"CVE-2022-1234", "CVE-2023-5678"},
			},
			want: []string{"CVE-2022-1234", "CVE-2023-5678"},
		},
		{
			name: "all invalid",
			args: args{
				cveSlice: []string{"not-a-cve", "also-not-cve"},
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				cveSlice: []string{},
			},
			want: nil,
		},
		{
			name: "mixed case output formatted",
			args: args{
				cveSlice: []string{"cve-2022-1234", "CVE-2023-5678"},
			},
			want: []string{"CVE-2022-1234", "CVE-2023-5678"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FilterValidCves(tt.args.cveSlice)
			if len(got) == 0 && (tt.want == nil || len(tt.want) == 0) {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterValidCves() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestExtractYear 测试extractYear函数（直接测试内部函数以提高覆盖率）
func TestExtractYear(t *testing.T) {
	type args struct {
		cve string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "valid CVE",
			args: args{
				cve: "CVE-2022-1234",
			},
			want: 2022,
		},
		{
			name: "invalid format",
			args: args{
				cve: "not-a-cve",
			},
			want: 0,
		},
		{
			name: "empty string",
			args: args{
				cve: "",
			},
			want: 0,
		},
		{
			name: "missing components",
			args: args{
				cve: "CVE-2022",
			},
			want: 0,
		},
		{
			name: "additional components",
			args: args{
				cve: "CVE-2022-1234-extra",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractYear(tt.args.cve); got != tt.want {
				t.Errorf("extractYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatSeq(t *testing.T) {
	type args struct {
		cve   string
		width int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "pad to 6 digits",
			args: args{
				cve:   "CVE-2022-123",
				width: 6,
			},
			want: "CVE-2022-000123",
		},
		{
			name: "pad to 5 digits",
			args: args{
				cve:   "CVE-2022-12345",
				width: 6,
			},
			want: "CVE-2022-012345",
		},
		{
			name: "already at target width",
			args: args{
				cve:   "CVE-2022-123456",
				width: 6,
			},
			want: "CVE-2022-123456",
		},
		{
			name: "wider than target",
			args: args{
				cve:   "CVE-2022-1234567",
				width: 6,
			},
			want: "CVE-2022-1234567",
		},
		{
			name: "invalid CVE returns as-is",
			args: args{
				cve:   "not-a-cve",
				width: 6,
			},
			want: "not-a-cve",
		},
		{
			name: "case insensitive input",
			args: args{
				cve:   "cve-2022-123",
				width: 6,
			},
			want: "CVE-2022-000123",
		},
		{
			// 序列号超出 int 范围，strconv.Atoi 失败 → 原样返回
			// 覆盖 base.go:85-87 的 err 分支（IsCve 通过但 Atoi 溢出）
			name: "sequence overflow returns as-is",
			args: args{
				cve:   "CVE-2022-99999999999999999999999999",
				width: 6,
			},
			want: "CVE-2022-99999999999999999999999999",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatSeq(tt.args.cve, tt.args.width); got != tt.want {
				t.Errorf("FormatSeq() = %v, want %v", got, tt.want)
			}
		})
	}
}
