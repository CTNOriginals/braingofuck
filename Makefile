WGO_INCLUDE := -file .go -file .bf

.PHONY: run wrun proto asci

run:
	go run . "$(args)"

wrun:
	wgo $(WGO_INCLUDE) go run . $(args)

proto:
	wgo $(WGO_INCLUDE) go run . --file ./brainfuck/proto.bf --input 1A2B3C4D

asci:
	go run "C:\Users\CTN\Documents\CTN\Programming\tmp\unicode-lister\main.go" "$(arg)"
