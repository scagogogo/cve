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
			args: args{
				cveA: "cve-2021-10086",
				cveB: "cve-2021-10086",
			},
			want: 0,
		},
		{
			args: args{
				cveA: "cve-2020-10086",
				cveB: "cve-2021-10086",
			},
			want: -1,
		},
		{
			args: args{
				cveA: "cve-2021-10086",
				cveB: "cve-2020-10086",
			},
			want: 1,
		},
		{
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
			args: args{
				text: "this is cve-2001-10086,",
			},
			want: []string{"CVE-2001-10086"},
		},
		{
			args: args{
				text: "this is Cve-2001-10086,",
			},
			want: []string{"CVE-2001-10086"},
		},
		{
			args: args{
				text: "this is cVe-2001-10086,",
			},
			want: []string{"CVE-2001-10086"},
		},
		{
			args: args{
				text: "this is CvE-2001-10086,",
			},
			want: []string{"CVE-2001-10086"},
		},
		{
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
					"cve-2001-1001",
					"cve-2001-1002",
					"cve-2001-1003",
					"cve-2001-1004",
				},
				"2201": []string{
					"cve-2201-1001",
					"cve-2201-1002",
					"cve-2201-1003",
					"cve-2201-1004",
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
			args: args{
				text: "cve-2007-199",
			},
			want: true,
		},
		{
			args: args{
				text: " cve-2007-199",
			},
			want: true,
		},
		{
			args: args{
				text: "cve-2007-199 ",
			},
			want: true,
		},
		{
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
