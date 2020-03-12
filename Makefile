main: main.go
	echo 'Build backend binary'
	mkdir release
	go build -o release
	mkdir release/upload_files

	echo 'Build frontend'
	mkdir release/static
	sh -c "cd frontend && npm install && npm run build"
	mv frontend/build/* release/static

	echo 'Done!'

clean:
	rm -rf release
