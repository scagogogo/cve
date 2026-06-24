package cve

import (
	"reflect"
	"testing"
)

func TestExtractCve(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "simple CVE in text",
			args: args{
				text: "this is cve-2001-10086,",
			},
			want: []string{"CVE-2001-10086"},
		},
		{
			name: "uppercase C lowercase ve",
			args: args{
				text: "this is Cve-2001-10086,",
			},
			want: []string{"CVE-2001-10086"},
		},
		{
			name: "lowercase c uppercase V lowercase e",
			args: args{
				text: "this is cVe-2001-10086,",
			},
			want: []string{"CVE-2001-10086"},
		},
		{
			name: "lowercase c lowercase v uppercase E",
			args: args{
				text: "this is CvE-2001-10086,",
			},
			want: []string{"CVE-2001-10086"},
		},
		{
			name: "multiple CVEs in text",
			args: args{
				text: "this is CvE-2001-10086, this is CvE-2001-10087,",
			},
			want: []string{"CVE-2001-10086", "CVE-2001-10087"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtractCve(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractCve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractCveYear(t *testing.T) {
	type args struct {
		cve string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "extract year from standard CVE",
			args: args{
				cve: "CVE-2001-10087",
			},
			want: "2001",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtractCveYear(tt.args.cve); got != tt.want {
				t.Errorf("ExtractCveYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractCveYearAsInt(t *testing.T) {
	type args struct {
		cve string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "extract year as integer",
			args: args{
				cve: "CVE-2001-10087",
			},
			want: 2001,
		},
		{
			name: "invalid CVE format",
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtractCveYearAsInt(tt.args.cve); got != tt.want {
				t.Errorf("ExtractCveYearAsInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractFirstCve(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "extract first CVE from text with multiple CVEs",
			args: args{
				text: "this is cve-2001-10086, this is CvE-2001-10087,",
			},
			want: "CVE-2001-10086",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtractFirstCve(tt.args.text); got != tt.want {
				t.Errorf("ExtractFirstCve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractLastCve(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "extract last CVE from text with multiple CVEs",
			args: args{
				text: "this is cve-2001-10086, this is CvE-2001-10087,",
			},
			want: "CVE-2001-10087",
		},
		{
			name: "empty text",
			args: args{
				text: "",
			},
			want: "",
		},
		{
			name: "text without CVE",
			args: args{
				text: "this text doesn't contain any CVE identifiers",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtractLastCve(tt.args.text); got != tt.want {
				t.Errorf("ExtractLastCve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractCveSeq(t *testing.T) {
	type args struct {
		cve string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "standard sequence number",
			args: args{
				cve: "CVE-2022-1234",
			},
			want: "1234",
		},
		{
			name: "longer sequence number",
			args: args{
				cve: "CVE-2022-123456",
			},
			want: "123456",
		},
		{
			name: "sequence with leading zeros",
			args: args{
				cve: "CVE-2022-0001",
			},
			want: "0001",
		},
		{
			name: "mixed case",
			args: args{
				cve: "cve-2022-5678",
			},
			want: "5678",
		},
		{
			name: "invalid format",
			args: args{
				cve: "not-a-cve",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtractCveSeq(tt.args.cve); got != tt.want {
				t.Errorf("ExtractCveSeq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractCveSeqAsInt(t *testing.T) {
	type args struct {
		cve string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "standard sequence number",
			args: args{
				cve: "CVE-2022-1234",
			},
			want: 1234,
		},
		{
			name: "larger sequence number",
			args: args{
				cve: "CVE-2022-123456",
			},
			want: 123456,
		},
		{
			name: "sequence with leading zeros",
			args: args{
				cve: "CVE-2022-0001",
			},
			want: 1,
		},
		{
			name: "mixed case",
			args: args{
				cve: "cve-2022-5678",
			},
			want: 5678,
		},
		{
			name: "invalid format",
			args: args{
				cve: "not-a-cve",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtractCveSeqAsInt(tt.args.cve); got != tt.want {
				t.Errorf("ExtractCveSeqAsInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterCvesByPattern(t *testing.T) {
	type args struct {
		cveSlice []string
		pattern  string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "filter by year wildcard",
			args: args{
				cveSlice: []string{"CVE-2022-1111", "CVE-2022-2222", "CVE-2023-1111", "CVE-2023-2222"},
				pattern:  "CVE-2022-*",
			},
			want: []string{"CVE-2022-1111", "CVE-2022-2222"},
		},
		{
			name: "filter by sequence wildcard",
			args: args{
				cveSlice: []string{"CVE-2022-1111", "CVE-2021-1111", "CVE-2023-2222"},
				pattern:  "CVE-*-1111",
			},
			want: []string{"CVE-2021-1111", "CVE-2022-1111"},
		},
		{
			name: "filter by prefix pattern",
			args: args{
				cveSlice: []string{"CVE-2022-1111", "CVE-2022-1122", "CVE-2022-2222"},
				pattern:  "CVE-2022-11*",
			},
			want: []string{"CVE-2022-1111", "CVE-2022-1122"},
		},
		{
			name: "exact match (no wildcard)",
			args: args{
				cveSlice: []string{"CVE-2022-1111", "CVE-2022-2222"},
				pattern:  "CVE-2022-1111",
			},
			want: []string{"CVE-2022-1111"},
		},
		{
			name: "no matches",
			args: args{
				cveSlice: []string{"CVE-2022-1111", "CVE-2022-2222"},
				pattern:  "CVE-2023-*",
			},
			want: nil,
		},
		{
			name: "case insensitive pattern",
			args: args{
				cveSlice: []string{"CVE-2022-1111", "CVE-2022-2222"},
				pattern:  "cve-2022-*",
			},
			want: []string{"CVE-2022-1111", "CVE-2022-2222"},
		},
		{
			name: "empty slice",
			args: args{
				cveSlice: []string{},
				pattern:  "CVE-2022-*",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FilterCvesByPattern(tt.args.cveSlice, tt.args.pattern)
			if len(got) == 0 && (tt.want == nil || len(tt.want) == 0) {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterCvesByPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
