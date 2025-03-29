package cve

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestCompareByYear(t *testing.T) {
	type args struct {
		cveA string
		cveB string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "same year",
			args: args{
				cveA: "cve-2021-10086",
				cveB: "cve-2021-10086",
			},
			want: 0,
		},
		{
			name: "cveA year is less than cveB",
			args: args{
				cveA: "cve-2020-10086",
				cveB: "cve-2021-10086",
			},
			want: -1,
		},
		{
			name: "cveA year is greater than cveB",
			args: args{
				cveA: "cve-2021-10086",
				cveB: "cve-2020-10086",
			},
			want: 1,
		},
		{
			name: "different case but same year",
			args: args{
				cveA: "cve-2021-10086",
				cveB: "CVE-2021-10086",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareByYear(tt.args.cveA, tt.args.cveB); got != tt.want {
				t.Errorf("CompareByYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractCVE(t *testing.T) {
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

func TestExtractCVEYear(t *testing.T) {
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

func TestExtractCVEYearAsInt(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtractCveYearAsInt(tt.args.cve); got != tt.want {
				t.Errorf("ExtractCveYearAsInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractFirstCVE(t *testing.T) {
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

func TestExtractLastCVE(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtractLastCve(tt.args.text); got != tt.want {
				t.Errorf("ExtractLastCve() = %v, want %v", got, tt.want)
			}
		})
	}
}

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

func TestGroupByYear(t *testing.T) {
	type args struct {
		cveSlice []string
	}
	tests := []struct {
		name string
		args args
		want map[string][]string
	}{
		{
			name: "group CVEs by year",
			args: args{
				cveSlice: []string{
					"cve-2001-1001",
					"cve-2001-1002",
					"cve-2001-1003",
					"cve-2001-1004",
					"cve-2201-1001",
					"cve-2201-1002",
					"cve-2201-1003",
					"cve-2201-1004",
				},
			},
			want: map[string][]string{
				"2001": []string{
					"CVE-2001-1001",
					"CVE-2001-1002",
					"CVE-2001-1003",
					"CVE-2001-1004",
				},
				"2201": []string{
					"CVE-2201-1001",
					"CVE-2201-1002",
					"CVE-2201-1003",
					"CVE-2201-1004",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GroupByYear(tt.args.cveSlice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupByYear() = %v, want %v", got, tt.want)
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

func TestSubByYear(t *testing.T) {
	type args struct {
		cveA string
		cveB string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "subtract CVE years",
			args: args{
				cveA: "cve-2007-10086",
				cveB: "cve-2017-10086",
			},
			want: -10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubByYear(tt.args.cveA, tt.args.cveB); got != tt.want {
				t.Errorf("SubByYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsCveYearOk(t *testing.T) {
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
			want: false,
		},
		{
			name: "future year with larger cutoff",
			args: args{
				cve:    "CVE-2099-10086",
				cutoff: 100, // Large cutoff to allow future dates
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
			if got := IsCveYearOk(tt.args.cve, tt.args.cutoff); got != tt.want {
				t.Errorf("IsCveYearOk() = %v, want %v", got, tt.want)
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

func TestCompareCves(t *testing.T) {
	type args struct {
		cveA string
		cveB string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "same CVE",
			args: args{
				cveA: "CVE-2021-10086",
				cveB: "CVE-2021-10086",
			},
			want: 0,
		},
		{
			name: "same year, different sequence (A < B)",
			args: args{
				cveA: "CVE-2021-10086",
				cveB: "CVE-2021-10087",
			},
			want: -1,
		},
		{
			name: "same year, different sequence (A > B)",
			args: args{
				cveA: "CVE-2021-10088",
				cveB: "CVE-2021-10087",
			},
			want: 1,
		},
		{
			name: "different year (A < B)",
			args: args{
				cveA: "CVE-2020-10086",
				cveB: "CVE-2021-10086",
			},
			want: -1,
		},
		{
			name: "different year (A > B)",
			args: args{
				cveA: "CVE-2022-10086",
				cveB: "CVE-2021-10086",
			},
			want: 1,
		},
		{
			name: "same CVE different case",
			args: args{
				cveA: "cve-2021-10086",
				cveB: "CVE-2021-10086",
			},
			want: 0,
		},
		{
			name: "different year and sequence",
			args: args{
				cveA: "CVE-2020-30000",
				cveB: "CVE-2021-10086",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareCves(tt.args.cveA, tt.args.cveB); got != tt.want {
				t.Errorf("CompareCves() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortCves(t *testing.T) {
	type args struct {
		cveSlice []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "already sorted",
			args: args{
				cveSlice: []string{
					"CVE-2020-1000",
					"CVE-2021-1000",
					"CVE-2021-1001",
					"CVE-2022-1000",
				},
			},
			want: []string{
				"CVE-2020-1000",
				"CVE-2021-1000",
				"CVE-2021-1001",
				"CVE-2022-1000",
			},
		},
		{
			name: "unsorted",
			args: args{
				cveSlice: []string{
					"CVE-2022-1000",
					"CVE-2020-1000",
					"CVE-2021-1001",
					"CVE-2021-1000",
				},
			},
			want: []string{
				"CVE-2020-1000",
				"CVE-2021-1000",
				"CVE-2021-1001",
				"CVE-2022-1000",
			},
		},
		{
			name: "mixed case",
			args: args{
				cveSlice: []string{
					"cve-2022-1000",
					"CVE-2020-1000",
					"Cve-2021-1001",
					"cvE-2021-1000",
				},
			},
			want: []string{
				"CVE-2020-1000",
				"CVE-2021-1000",
				"CVE-2021-1001",
				"CVE-2022-1000",
			},
		},
		{
			name: "empty slice",
			args: args{
				cveSlice: []string{},
			},
			want: []string{},
		},
		{
			name: "single element",
			args: args{
				cveSlice: []string{"CVE-2021-1000"},
			},
			want: []string{"CVE-2021-1000"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortCves(tt.args.cveSlice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortCves() = %v, want %v", got, tt.want)
			}
		})
	}
}

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

func TestFilterCvesByYear(t *testing.T) {
	type args struct {
		cveSlice []string
		year     int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "filter specific year",
			args: args{
				cveSlice: []string{
					"CVE-2020-1000",
					"CVE-2021-1000",
					"CVE-2021-1001",
					"CVE-2022-1000",
				},
				year: 2021,
			},
			want: []string{
				"CVE-2021-1000",
				"CVE-2021-1001",
			},
		},
		{
			name: "no matches",
			args: args{
				cveSlice: []string{
					"CVE-2020-1000",
					"CVE-2021-1000",
					"CVE-2022-1000",
				},
				year: 2019,
			},
			want: nil,
		},
		{
			name: "mixed case",
			args: args{
				cveSlice: []string{
					"cve-2020-1000",
					"CVE-2021-1000",
					"Cve-2021-1001",
				},
				year: 2021,
			},
			want: []string{
				"CVE-2021-1000",
				"CVE-2021-1001",
			},
		},
		{
			name: "empty slice",
			args: args{
				cveSlice: []string{},
				year:     2022,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FilterCvesByYear(tt.args.cveSlice, tt.args.year)
			// Both are empty slices or nil
			if len(got) == 0 && (tt.want == nil || len(tt.want) == 0) {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterCvesByYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterCvesByYearRange(t *testing.T) {
	type args struct {
		cveSlice  []string
		startYear int
		endYear   int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "filter year range",
			args: args{
				cveSlice: []string{
					"CVE-2019-1000",
					"CVE-2020-1000",
					"CVE-2021-1000",
					"CVE-2022-1000",
					"CVE-2023-1000",
				},
				startYear: 2020,
				endYear:   2022,
			},
			want: []string{
				"CVE-2020-1000",
				"CVE-2021-1000",
				"CVE-2022-1000",
			},
		},
		{
			name: "exact year range match",
			args: args{
				cveSlice: []string{
					"CVE-2020-1000",
					"CVE-2021-1000",
				},
				startYear: 2020,
				endYear:   2021,
			},
			want: []string{
				"CVE-2020-1000",
				"CVE-2021-1000",
			},
		},
		{
			name: "no matches",
			args: args{
				cveSlice: []string{
					"CVE-2020-1000",
					"CVE-2021-1000",
				},
				startYear: 2022,
				endYear:   2023,
			},
			want: nil,
		},
		{
			name: "mixed case",
			args: args{
				cveSlice: []string{
					"cve-2020-1000",
					"CVE-2021-1000",
					"Cve-2022-1001",
				},
				startYear: 2020,
				endYear:   2021,
			},
			want: []string{
				"CVE-2020-1000",
				"CVE-2021-1000",
			},
		},
		{
			name: "empty slice",
			args: args{
				cveSlice:  []string{},
				startYear: 2020,
				endYear:   2022,
			},
			want: nil,
		},
		{
			name: "inverted range",
			args: args{
				cveSlice: []string{
					"CVE-2020-1000",
					"CVE-2021-1000",
				},
				startYear: 2021,
				endYear:   2020,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FilterCvesByYearRange(tt.args.cveSlice, tt.args.startYear, tt.args.endYear)
			// Both are empty slices or nil
			if len(got) == 0 && (tt.want == nil || len(tt.want) == 0) {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterCvesByYearRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRecentCves(t *testing.T) {
	currentYear := time.Now().Year()
	lastYear := currentYear - 1

	type args struct {
		cveSlice []string
		years    int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "get recent CVEs",
			args: args{
				cveSlice: []string{
					fmt.Sprintf("CVE-%d-1000", currentYear-3),
					fmt.Sprintf("CVE-%d-1000", currentYear-2),
					fmt.Sprintf("CVE-%d-1000", lastYear),
					fmt.Sprintf("CVE-%d-1000", currentYear),
				},
				years: 2,
			},
			want: []string{
				fmt.Sprintf("CVE-%d-1000", lastYear),
				fmt.Sprintf("CVE-%d-1000", currentYear),
			},
		},
		{
			name: "get all recent CVEs",
			args: args{
				cveSlice: []string{
					fmt.Sprintf("CVE-%d-1000", lastYear),
					fmt.Sprintf("CVE-%d-1000", currentYear),
				},
				years: 2,
			},
			want: []string{
				fmt.Sprintf("CVE-%d-1000", lastYear),
				fmt.Sprintf("CVE-%d-1000", currentYear),
			},
		},
		{
			name: "no matches",
			args: args{
				cveSlice: []string{
					fmt.Sprintf("CVE-%d-1000", currentYear-5),
					fmt.Sprintf("CVE-%d-1000", currentYear-4),
					fmt.Sprintf("CVE-%d-1000", currentYear-3),
				},
				years: 2,
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				cveSlice: []string{},
				years:    2,
			},
			want: nil,
		},
		{
			name: "mixed case",
			args: args{
				cveSlice: []string{
					fmt.Sprintf("cve-%d-1000", lastYear),
					fmt.Sprintf("CVE-%d-1000", currentYear),
				},
				years: 2,
			},
			want: []string{
				fmt.Sprintf("CVE-%d-1000", lastYear),
				fmt.Sprintf("CVE-%d-1000", currentYear),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetRecentCves(tt.args.cveSlice, tt.args.years)
			// Both are empty slices or nil
			if len(got) == 0 && (tt.want == nil || len(tt.want) == 0) {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRecentCves() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveDuplicateCves(t *testing.T) {
	type args struct {
		cveSlice []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "no duplicates",
			args: args{
				cveSlice: []string{
					"CVE-2020-1000",
					"CVE-2021-1000",
					"CVE-2022-1000",
				},
			},
			want: []string{
				"CVE-2020-1000",
				"CVE-2021-1000",
				"CVE-2022-1000",
			},
		},
		{
			name: "has duplicates",
			args: args{
				cveSlice: []string{
					"CVE-2020-1000",
					"CVE-2021-1000",
					"CVE-2020-1000",
					"CVE-2022-1000",
					"CVE-2021-1000",
				},
			},
			want: []string{
				"CVE-2020-1000",
				"CVE-2021-1000",
				"CVE-2022-1000",
			},
		},
		{
			name: "same CVE different case",
			args: args{
				cveSlice: []string{
					"cve-2020-1000",
					"CVE-2020-1000",
					"Cve-2020-1000",
				},
			},
			want: []string{
				"CVE-2020-1000",
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
			name: "all duplicates",
			args: args{
				cveSlice: []string{
					"CVE-2020-1000",
					"CVE-2020-1000",
					"CVE-2020-1000",
				},
			},
			want: []string{
				"CVE-2020-1000",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RemoveDuplicateCves(tt.args.cveSlice)
			// Both are empty slices or nil
			if len(got) == 0 && (tt.want == nil || len(tt.want) == 0) {
				return
			}
			// Sort before comparing since order isn't guaranteed for map iteration
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveDuplicateCves() = %v, want %v", got, tt.want)
			}
		})
	}
}
