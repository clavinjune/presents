GOLANG_CMD_VERSION=v0.1.11
PRESENT_DEP=golang.org/x/tools/present@$(GOLANG_CMD_VERSION)
PRESENT_CMD=golang.org/x/tools/cmd/present@$(GOLANG_CMD_VERSION)
GRAPH_EASY_CMD=$(shell perl -V:'installbin' | awk -F= '{ print substr( $$2, 2, length($$2)-3) "/graph-easy" }')

present: get
	@go run $(PRESENT_CMD) -notes -use_playground .

get:
	@go get $(PRESENT_DEP)

# make graph foo-at-bar/graph.txt
graph:
	@$(GRAPH_EASY_CMD) $(FILE) > "$(FILE).ascii"
graph-png:
	@$(GRAPH_EASY_CMD) $(FILE) --png

base=./developing-api-services-with-golang-at-finditgeek/assets
graph/developing-api-services-with-golang-at-finditgeek: 
	@$(MAKE) graph-png FILE="$(base)/graph.txt"
	@$(MAKE) graph FILE=$(base)/http-method-table.txt
	@$(MAKE) graph FILE=$(base)/to-do-api.txt

install/graph-easy:
	@brew install graphviz
	@sudo cpan Graph::Easy