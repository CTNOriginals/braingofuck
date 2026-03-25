MAKE := make --no-print-directory
.SHELLFLAGS += -c

WGO_INCLUDE := -file .go -file .bf -file Makefile -file .asm

.PHONY: run wrun proto nasm asm wasm test asci
.ONESHELL:

run:
	go run . "$(args)"

wrun:
	wgo $(WGO_INCLUDE) go run . $(args)

proto:
	wgo $(WGO_INCLUDE) go run . --file ./brainfuck/proto.bf --input 1A2B3C4D

wasm:
	@wgo $(WGO_INCLUDE) $(wgoargs) $(MAKE) asm

asm:
	@cd brainfuck/compiled
	
	nasm \
	-f win64 \
	-X gcc \
	--debug \
	-g \
	-o ./bin/proto.obj \
	proto.asm

	gcc -o ./bin/proto.exe ./bin/proto.obj

	echo 
	echo 
	echo -- $(shell date +"%H:%M:%S") --
	./bin/proto.exe

asci:
	go run "C:\Users\CTN\Documents\CTN\Programming\tmp\unicode-lister\main.go" "$(arg)"
