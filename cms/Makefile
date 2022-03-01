build:
	docker image build . --tag loqui-studio:dev
dev:
	docker container run -it \
	--rm \
	--publish 3000:3000 \
	--name loqui-studio-dev \
	--mount \
	src=`pwd`,target=/app,type=bind \
	loqui-studio:dev