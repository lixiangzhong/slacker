.PHONY: all
all:
	@make bindata
	@make install

.PHONY: bindata
bindata: 
	@cd bindata && go-bindata -pkg="bindata" -prefix example_project/ example_project/... 
	@cd mvctemplate && go-bindata -pkg="mvctemplate" *.tpl

.PHONY: install
install: 
	@go install github.com/lixiangzhong/slacker/slacker

 