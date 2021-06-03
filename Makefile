VERSION=0.0.2
IMAGENAME=testapp

all: vet format
	GOOS=linux GOARCH=amd64 go build -o bin/${IMAGENAME} main.go

vet:
	go vet main.go

format:
	go fmt main.go


docker.build: all copy.certs
	docker build . -t ${IMAGENAME}:${VERSION} --no-cache

docker.run: docker.build
	docker run -d ${IMAGENAME}:${VERSION}

docker.push: 
	docker tag ${IMAGENAME}:${VERSION} quay.io/ixcloud/${IMAGENAME}:${VERSION}
	docker push quay.io/ixcloud/${IMAGENAME}:${VERSION}

copy.certs:
	if [ ! -d ./certs ]; then \
		mkdir ./certs; \
	fi
	if [ -f /etc/pki/ca-trust/source/anchors/XX_CA_I1.crt ]; then \
		cp /etc/pki/ca-trust/source/anchors/XX_CA_I1.crt ./certs;  \
		chmod 0644 ./certs/XX_CA_I1.crt;  \
	else \
		touch ./certs/XX_CA_I1.crt; \
	fi
	if [ -f /etc/pki/ca-trust/source/anchors/XX_Root_CA1.crt ]; then \
		cp /etc/pki/ca-trust/source/anchors/XX_Root_CA1.crt ./certs;  \
		chmod 0644 ./certs/XX_Root_CA1.crt;  \
	else \
		touch ./certs/XX_Root_CA1.crt; \
	fi
