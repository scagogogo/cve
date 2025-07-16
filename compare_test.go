package cve

import (
	"reflect"
	"testing"
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
