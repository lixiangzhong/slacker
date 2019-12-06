build_version=`date +"%Y%m%d"`
build_date=`date +"%Y-%m-%d_%H:%M:%S"`
build_gitversion=`git rev-parse HEAD|cut -c1-8`
gofilemd5=`find ./ -name "*.go" -type f -print0| xargs -0 md5|md5|cut -c1-8`

all:
	@echo "make : build dist zip"

.PHONY: build
build:
	@go build  -ldflags "-X {{.ProjectName}}/gosrc/app.Version=$(build_version)\
	  -X {{.ProjectName}}/gosrc/app.BuildTime=$(build_date)\
	    -X {{.ProjectName}}/gosrc/app.GitVersion=$(build_gitversion)\
			-X {{.ProjectName}}/gosrc/app.GoMD5=$(gofilemd5)"\
		 -o {{.ProjectName}}
	@GOOS=linux go build  -ldflags "-X {{.ProjectName}}/gosrc/app.Version=$(build_version)\
	  -X {{.ProjectName}}/gosrc/app.BuildTime=$(build_date)\
	    -X {{.ProjectName}}/gosrc/app.GitVersion=$(build_gitversion)\
			-X {{.ProjectName}}/gosrc/app.GoMD5=$(gofilemd5)"\
		 -o {{.ProjectName}}_linux
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