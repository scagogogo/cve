package cve

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
	"time"
)

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
