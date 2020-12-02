docker-build:
	docker build -f Dockerfile . -t sandbox

cloud-build:
	gcloud builds submit --config cloudbuild.yaml .

install-realize:
	GO111MODULE=off go get -u github.com/oxequa/realize
	realize -v

realize:
	realize s --build --run
