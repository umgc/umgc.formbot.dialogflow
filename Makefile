run: ssl-local download-dependencies
	cd formscriber && go run .

download-dependencies:
	go mod download

ssl-local:
	openssl req -x509 -out formscriber/formscriber.com.pem \
	-keyout formscriber/formscriber.key -newkey rsa:2048 -nodes -sha256 \
	-subj '/CN=localhost'

build-docker:
	docker build -t formscriber_df .

run-docker:
	docker run formscriber_df

clean:
	find . | grep -E '(\.log|\.pem|\.key)' | xargs rm -rf