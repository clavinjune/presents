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

install/graph-easy:
	@brew install graphviz
	@sudo cpan Graph::Easy