# vim: set noexpandtab fo-=t:
# https://www.gnu.org/software/make/manual/make.html
.PHONY: default
default:

########################################################################
# boiler plate
########################################################################
SHELL=bash
current_makefile:=$(lastword $(MAKEFILE_LIST))
current_makefile_dirname:=$(dir $(current_makefile))
current_makefile_dirname_abspath:=$(dir $(abspath $(current_makefile)))
current_makefile_dirname_realpath:=$(dir $(realpath $(current_makefile)))

ifneq ($(filter all vars,$(VERBOSE)),)
dump_var=$(info var $(1)=$($(1)))
dump_vars=$(foreach var,$(1),$(call dump_var,$(var)))
else
dump_var=
dump_vars=
endif

ifneq ($(filter all targets,$(VERBOSE)),)
__ORIGINAL_SHELL:=$(SHELL)
SHELL=$(warning Building $@$(if $<, (from $<))$(if $?, ($? newer)))$(TIME) $(__ORIGINAL_SHELL)
endif

define __newline


endef


skip=
# skipable makes the targets passed to it skipable with skip=foo%
# $(1): targets that should be skipable
skipable=$(filter-out $(skip),$(1))

########################################################################
# variables ...
########################################################################

main_cmd=iagotmpl
tools_dir=local-tools
cache_dir=$(HOME)/.cache
build_dir=build
GOPATH:=$(shell type go >/dev/null 2>&1 && go env GOPATH)
export PATH:=$(if $(GOPATH),$(GOPATH)/bin:,)$(PATH)
# export GOFLAGS:=-mod=vendor

giti_description=$(shell git describe --tags --always)
giti_commit_count:=$(shell git rev-list --no-merges --count HEAD)
giti_cstate=$(shell [[ -z $$(git status -s) ]] && echo "s" || echo "m")
giti_commit_hash=$(shell git log -1 --format="%H")
giti_commit_hash_short=$(shell git log -1 --format="%h")
giti_suffix=$(giti_commit_count)$(giti_cstate)-$(giti_commit_hash_short)
giti_url:=$(shell git remote get-url origin)
$(call dump_vars,giti_suffix)
dt_info:=$(shell date "+%Y-%m-%d %H:%M:%S%z")

########################################################################
# targets ...
########################################################################

.PHONY: all
default: all
all:

.PHONY: clean
clean: clean-$(build_dir)/

.PHONY: distclean
distclean: clean-$(tools_dir)/


modd_bin=$(tools_dir)/modd
modd_version=0.8
$(modd_bin): | $(tools_dir)/ $(cache_dir)/
	cd $(cache_dir) && wget -c https://github.com/cortesi/modd/releases/download/v$(modd_version)/modd-$(modd_version)-linux64.tgz
	tar -zxvf $(cache_dir)/modd-0.8-linux64.tgz --strip-components=1 -C $(tools_dir)

.PHONY: toolchain
toolchain: $(modd_bin)
	cd ~/ && go get $(if $(filter all commands,$(VERBOSE)),-v) $(go_get_flags) \
		github.com/golangci/golangci-lint/cmd/golangci-lint@v1.39.0 \
		&& true

.PHONY: toolchain-update
toolchain-update: go_get_flags+=-u
toolchain-update: toolchain

.PHONY: validate-static
validate-static:
	golangci-lint run $(if $(filter all commands,$(VERBOSE)),-v) ./...

.PHONY: validate-fix
validate-fix:
	golangci-lint run $(if $(filter all commands,$(VERBOSE)),-v) --fix ./...

.PHONY: test
test:
	go test -cover -race ./...

.PHONY: validate-dynamic
validate-dynamic: test

.PHONY: validate
validate: validate-static validate-dynamic

.PHONY: watch
watch: $(modd_bin)
	./$(modd_bin) --debug --notify --file modd.conf

.PHONY: generate
generate:


go_ldflags= \
	-ldflags \
	"-X 'main.gitCommit=$(giti_commit_hash)' \
	-X 'main.gitRemote=$(giti_url)' \
	-X 'main.gitDesc=$(giti_description)' \
	-X 'main.buildDT=$(dt_info)'" \

.PHONY: build
all: build
build: $(call skipable,validate)
	go build $(go_ldflags) ./...

.PHONY: install
install: $(call skipable,validate)
	go install $(go_ldflags) ./...

.PHONY: run-help
run-help:
	$(foreach cmd,$(wildcard ./cmd/*),go run $(cmd) --help)

.PHONY: run
run: run_args=run
run:
	go run ./cmd/$(main_cmd) $(run_args)

build-oci-inputs: | $(build_dir)/go/bin/
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
		go build $(go_ldflags) -a -v \
		-o $(build_dir)/go/bin/ \
		./...


hadolint_tag=:latest
.PHONY: validate-dockerfile
validate-dockerfile:
	docker run --rm -i ghcr.io/hadolint/hadolint$(hadolint_tag) < Dockerfile

oci_tag_prefix=
oci_tag_suffixes = \
	gits-$(giti_suffix) \
	gitd-$(giti_description) \
	latest \


oci_registry=example.com
oci_repo=some/service

oci_refs_local=\
	$(foreach oci_tag_suffix,\
		$(oci_tag_suffixes),\
		ocreg.localhost/$(oci_registry)/$(oci_repo):$(oci_tag_prefix)$(oci_tag_suffix))


oci_refs_remote=\
	$(foreach oci_tag_suffix,\
		$(oci_tag_suffixes),\
		$(oci_registry)/$(oci_repo):$(oci_tag_prefix)$(oci_tag_suffix))

oci_refs = \
	$(oci_refs_local) \
	$(oci_refs_remote) \

$(call dump_vars,oci_refs)

.PHONY: build-oci
build-oci: Dockerfile $(call skipable,validate-dockerfile build-oci-inputs)
	docker image build --force-rm \
		$(foreach oci_tag,$(oci_refs),--tag $(oci_tag)) \
		--build-arg main_cmd=$(main_cmd) \
		--file $(<) \
		$(build_dir)

dockle_tag=:latest
.PHONY: validate-oci
validate-oci: $(call skipable,build-oci)
	docker run --rm -v /var/run/docker.sock:/var/run/docker.sock \
		docker.io/goodwithtech/dockle$(dockle_tag) \
		$(firstword $(oci_refs_local))


.PHONY: run-oci
run-oci: $(call skipable,build-oci)
	docker container run --rm -it \
		$(firstword $(oci_refs_local))

.PHONY: push-oci
push-oci: $(call skipable,build-oci)
	$(foreach oci_ref_remote,$(oci_refs_remote),\
		docker push $(oci_ref_remote)$(__newline))

help:
	@printf "########################################################################\n"
	@printf "TARGETS:\n"
	@printf "########################################################################\n"
	@printf "%-32s%s\n" "help" "Show this output ..."
	@printf "%-32s%s\n" "all" "Build all outputs (default)"
	@printf "%-32s%s\n" "toolchain" "Install toolchain"
	@printf "%-32s%s\n" "validate" "Validate everything"
	@printf "%-32s%s\n" "validate-fix" "Fix auto-fixable validation errors"
	@printf "%-32s%s\n" "generate" "Generate code"
	@printf "%-32s%s\n" "build" "Build everything"
	@printf "%-32s%s\n" "install" "Install everything"
	@printf "%-32s%s\n" "run-help" "Runs every command with --help"
	@printf "%-32s%s\n" "run" "Runs the main command"
	@printf "%-32s%s\n" "distclean" "Cleans all things that should not be checked in"
	@printf "%-32s%s\n" "build-oci" "Build the OCI Image"
	@printf "%-32s%s\n" "run-oci" "Run the OCI Image"
	@printf "%-32s%s\n" "validate-oci" "Validate the OCI Image"
	@printf "########################################################################\n"
	@printf "VARIABLES:\n"
	@printf "########################################################################\n"
	@printf "%-32s%s\n" "VERBOSE" "Sets verbosity for specific aspects." \
						"" "Space seperated." \
						"" "Valid values: all, vars, commands, targets"

########################################################################
# useful ...
########################################################################
## force ...
.PHONY: .FORCE
.FORCE:
$(force_targets): .FORCE

## dirs ...
.PRECIOUS: %/
%/:
	mkdir -vp $(@)

.PHONY: clean-%/
clean-%/:
	@{ test -d $(*) && { set -x; rm -vr $(*); set +x; } } || echo "directory $(*) does not exist ... nothing to clean"
