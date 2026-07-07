package cmd

import (
	"os"
	"testing"
)

// TestReadInputsFromArgs 覆盖 len(args) > 0 分支：直接返回 args，不读 stdin。
func TestReadInputsFromArgs(t *testing.T) {
	args := []string{"CVE-2022-1", "CVE-2022-2"}
	got := readInputs(args)
	if len(got) != 2 || got[0] != "CVE-2022-1" || got[1] != "CVE-2022-2" {
		t.Fatalf("readInputs(args) = %v, want %v", got, args)
	}
}

// TestReadInputsFromStdin 覆盖 stdin 管道分支：逐行读取，跳过空行。
func TestReadInputsFromStdin(t *testing.T) {
	originalStdin := os.Stdin
	defer func() { os.Stdin = originalStdin }()

	input := "CVE-2022-1\n\nCVE-2022-2\n"
	tmp, err := os.CreateTemp("", "stdin-*")
	if err != nil {
		t.Fatalf("CreateTemp: %v", err)
	}
	defer os.Remove(tmp.Name())
	if _, err := tmp.WriteString(input); err != nil {
		t.Fatalf("WriteString: %v", err)
	}
	tmp.Close()

	f, err := os.Open(tmp.Name())
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	defer f.Close()
	os.Stdin = f

	got := readInputs(nil)
	want := []string{"CVE-2022-1", "CVE-2022-2"}
	if len(got) != len(want) {
		t.Fatalf("readInputs(nil) from stdin = %v, want %v (empty lines skipped)", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("readInputs(nil)[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}

// TestReadInputsCharDevice 覆盖字符设备分支（helpers.go:17）：
// /dev/null 是字符设备，(stat.Mode() & os.ModeCharDevice) != 0 为真 → return nil。
func TestReadInputsCharDevice(t *testing.T) {
	originalStdin := os.Stdin
	defer func() { os.Stdin = originalStdin }()

	f, err := os.Open("/dev/null")
	if err != nil {
		t.Skipf("/dev/null not available: %v", err)
	}
	defer f.Close()
	os.Stdin = f

	got := readInputs(nil)
	if got != nil {
		t.Fatalf("readInputs(nil) with /dev/null (char device) = %v, want nil", got)
	}
}

// TestReadInputsEmptyPipe 覆盖 stdin 管道 EOF 分支：scanner 零迭代，lines 为 nil。
func TestReadInputsEmptyPipe(t *testing.T) {
	originalStdin := os.Stdin
	defer func() { os.Stdin = originalStdin }()

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Pipe: %v", err)
	}
	os.Stdin = r
	w.Close() // 立即关闭写端，读端读到 EOF，scanner 无数据

	got := readInputs(nil)
	if got != nil && len(got) != 0 {
		t.Fatalf("readInputs(nil) with empty pipe = %v, want nil or empty", got)
	}
}
