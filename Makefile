main_package_path = ./
binary_name = export
build_dir = ./bin

.PHONY: help
help:
	@echo "Usage"


.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N]' && read ans && [ $${ans:-N} = y ]
	

## build: build the application
.PHONY: build
build:
	@mkdir -p ${build_dir}
	GOARCH=amd64 GOOS=darwin build -o ${build_dir}/${binary_name}-darwin ${main_package_path}
	GOARCH=amd64 GOOS=linux build -o ${build_dir}/${binary_name}-linux ${main_package_path}
	GOARCH=amd64 GOOS=windows build -o ${build_dir}/${binary_name}-windows.exe ${main_package_path}

# clean: clean up the build binaries
.PHONY: clean
clean: confirm
	@echo "Cleaning up..."
	@rm -rf ${build_dir}

.PHONY: govulncheck
govulncheck:
	govulncheck ./...

.PHONY: test
test:
	go test -v --race ./...