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
	@wgo $(WGO_INCLUDE) $(MAKE) asm

asm:
	@cd brainfuck/compiled
	
	nasm \
	-f win64 \
	-X gcc \
	--debug \
	-g \
	proto.asm

	gcc -o proto.exe proto.obj

	echo 
	echo -- $(shell date +"%H:%M:%S") --
	./proto.exe

asci:
	go run "C:\Users\CTN\Documents\CTN\Programming\tmp\unicode-lister\main.go" "$(arg)"
