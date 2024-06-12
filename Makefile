MAKEFLAGS += --silent
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-variables
MAKEFLAGS += --no-builtin-rules


SHELL       := bash
.SHELLFLAGS := -eu -o pipefail -c
.ONESHELL:


.DEFAULT_GOAL := help


.PHONY: help
help: ## Display help.
	grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| sort \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


.PHONY: present
present: ## Start the presentation.
	present -content ./content -base ./theme -notes -use_playground=true
