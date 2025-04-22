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
			name: "year before 1970",
			args: args{
				cve: "CVE-1969-10086",
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
			name: "year before 1970",
			args: args{
				cve:    "CVE-1969-10086",
				cutoff: 100,
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
			name: "invalid year - before 1970",
			args: args{
				cve: "CVE-1969-1234",
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateCve(tt.args.cve); got != tt.want {
				t.Errorf("ValidateCve() = %v, want %v", got, tt.want)
			}
		})
	}
}
