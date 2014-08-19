all:
	make assetBinaries
	make templateBinaries
	make cleanup

css:
	sass --watch sass:assets/stylesheets


templateBinaries:
	touch bindata.go.temp && \
	go-bindata -prefix "app/templates" app/templates/template_files/... && \
	sed '1s/main/templates/' bindata.go > bindata.go.temp && \
	mv bindata.go.temp app/templates/bindata.go

assetBinaries:
	touch bindata.go.temp && \
	go-bindata assets/... && \
	sed '1s/main/assets/' bindata.go > bindata.go.temp && \
	mv bindata.go.temp app/assets/bindata.go

cleanup:
	rm bindata.go
