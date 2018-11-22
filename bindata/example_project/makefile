build_version=`date +"%Y%m%d"`
build_date=`date +"%Y-%m-%d_%H:%M:%S"`
build_gitversion=`git rev-parse HEAD`

all:
	@echo "make : build dist zip"

.PHONY: build
build:
	@go build  -ldflags "-X {{.ProjectPath}}/gosrc/app.Version=$(build_version)\
	  -X {{.ProjectPath}}/gosrc/app.BuildTime=$(build_date)\
	    -X {{.ProjectPath}}/gosrc/app.GitVersion=$(build_gitversion)"\
		 -o {{.ProjectName}}
	@./{{.ProjectName}} -v 

.PHONY: dist
dist:
	npm run build
	@make build
	@make zip

.PHONY: zip
zip:
	rm -f {{.ProjectName}}.zip
	zip -r {{.ProjectName}}.zip dist {{.ProjectName}} config.ini