main: main.go
	echo 'Build backend binary'
	mkdir release
	go build -o release

	echo 'Build frontend'
	mkdir release/static
	sh -c "cd frontend && npm install && npm run build"
	mv frontend/build/* release/static
	cp config.json release

	echo 'Done!'

clean:
	rm -rf release
