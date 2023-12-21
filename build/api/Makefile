all:
	cp Dockerfile ../../backend/dist
	docker build -t operator-api ../../backend/dist
	docker tag operator-api:latest us-central1-docker.pkg.dev/nelsona-sandbox/nelsona-dev-images/operator-api:${VERSION}
	docker push us-central1-docker.pkg.dev/nelsona-sandbox/nelsona-dev-images/operator-api:${VERSION}
