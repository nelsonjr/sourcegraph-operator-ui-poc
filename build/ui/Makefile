all:
	cp Dockerfile operator-ui.conf.template ../../frontend/dist
	docker build -t operator-ui ../../frontend/dist
	docker tag operator-ui:latest us-central1-docker.pkg.dev/nelsona-sandbox/nelsona-dev-images/operator-ui:${VERSION}
	docker push us-central1-docker.pkg.dev/nelsona-sandbox/nelsona-dev-images/operator-ui:${VERSION}
