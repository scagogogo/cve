.PHONY: test coverage coverage-inproc clean-coverage

# 常规单元测试（快速，无覆盖率）
test:
	go test -count=1 ./...

# 全仓库单一覆盖率：进程内 covdir + 子进程 covdir 合并 → coverage.out
# 进程内用 -coverpkg=.,./cmd 让 cmd 测试也插桩库代码；-test.gocoverdir 产出 covdir 格式
# 子进程用 GOCOVER_SUBPROCESS_DIRS 让 TestMain 导出 coverDir 列表（保留各目录）
# 两类 covdir 用 covdata merge -pcombine 合并（跨可执行文件），textfmt 转标准 profile
coverage: COV_INPROC_DIR := /tmp/cve-cov-inproc
coverage: COV_SUBPROC_DIRLIST := /tmp/cve-cov-subproc-dirs.txt
coverage: COV_MERGED_DIR := /tmp/cve-cov-merged
coverage: COV_OUT := coverage.out
coverage: clean-coverage
	@mkdir -p $(COV_INPROC_DIR)
	@mkdir -p $(COV_MERGED_DIR)
	# 1. 子进程 covdir：TestMain 导出 coverDir 列表（保留各目录不删）
	GOCOVER_SUBPROCESS_DIRS=$(COV_SUBPROC_DIRLIST) go test -count=1 ./cmd/ >/dev/null
	# 2. 进程内 covdir：库+cmd，-coverpkg=.,./cmd
	#    注意：包列表 . ./cmd/ 必须在 -args 之前，否则被当成 test binary 参数
	go test -count=1 -coverpkg=.,./cmd -cover . ./cmd/ -args -test.gocoverdir=$(COV_INPROC_DIR) >/dev/null
	# 3. 合并进程内 + 所有子进程 covdir（-pcombine 跨可执行文件合并）
	@SUBPROCS=$$(tr '\n' ',' < $(COV_SUBPROC_DIRLIST) | sed 's/,$$//'); \
	INPROC=$(COV_INPROC_DIR); \
	go tool covdata merge -pcombine -i=$${INPROC},$${SUBPROCS} -o=$(COV_MERGED_DIR); \
	go tool covdata textfmt -i=$(COV_MERGED_DIR) -o=$(COV_OUT)
	@echo "--- 覆盖率总览 ---"
	@go tool cover -func=$(COV_OUT) | tail -1

# 仅进程内覆盖率 covdir（库+cmd，-coverpkg 让 cmd 测试也插桩库代码）
coverage-inproc: clean-coverage
	@mkdir -p /tmp/cve-cov-inproc
	go test -count=1 -coverpkg=.,./cmd -cover . ./cmd/ -args -test.gocoverdir=/tmp/cve-cov-inproc
	@go tool covdata textfmt -i=/tmp/cve-cov-inproc -o=coverage.out
	@go tool cover -func=coverage.out | tail -1

clean-coverage:
	@rm -rf /tmp/cve-cov-inproc /tmp/cve-cov-merged coverage.out
