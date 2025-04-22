package cve

import (
	"fmt"
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
