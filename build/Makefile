VERSION=v0.0.$$(( $(shell date +%s) - 1703122377 ))

.PHONY: all

all:
	make .version VERSION=${VERSION}
	make build VERSION=${VERSION}
	make pkg VERSION=${VERSION}

build:
	@echo "---------------------------------------"
	@echo "Building ${VERSION}..."
	@echo "---------------------------------------"
	make build.api VERSION=${VERSION}
	make build.ui VERSION=${VERSION}
	@echo "---------------------------------------"
	@echo "Build ${VERSION} complete"
	@echo "---------------------------------------"

pkg:
	@echo "---------------------------------------"
	@echo "Packging ${VERSION}..."
	@echo "---------------------------------------"
	make pkg.api VERSION=${VERSION}
	make pkg.ui VERSION=${VERSION}
	make helm VERSION=${VERSION}
	@echo "---------------------------------------"
	@echo "Package ${VERSION} complete"
	@echo "---------------------------------------"

.version:
	@echo "======================================="
	@echo "Building version ${VERSION}"
	@echo "======================================="

build.ui:
	cd ../frontend && pnpm install
	cd ../frontend && pnpm run build

pkg.ui:
	make -C ui VERSION=${VERSION}

build.api:
	cd ../backend && go mod download
	cd ../backend && GOOS=linux GOARCH=amd64 go build -o dist/operator-api ./cmd

pkg.api:
	make -C api VERSION=${VERSION}

helm:
	@echo "Updating Helm to version ${VERSION}"
	cat ../helm/values.yaml \
	  | sed -e "s/\(.*image:.*\):[^:]*/\1:${VERSION}/" \
	  > ../helm/values.yaml.tmp
	mv ../helm/values.yaml.tmp ../helm/values.yaml
	cat ../helm/Chart.yaml \
	  | sed -e "s/version:.*/version: ${VERSION}/" \
	  | sed -e "s/appVersion:.*/appVersion: \"${VERSION}\"/" \
	  > ../helm/Chart.yaml.tmp
	mv ../helm/Chart.yaml.tmp ../helm/Chart.yaml
