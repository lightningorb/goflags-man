build:
	go build

init:
	go mod init example.com/hellogo

run:
	./hellogo -m yo

generate_man:
	./hellogo -g

install_man:
	sudo cp program.1 /usr/local/share/man/man1