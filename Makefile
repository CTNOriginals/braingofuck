.PHONY: run wrun asci

run:
	go run . "$(args)"

wrun:
	wgo -file .go -file .bf go run . $(args)

asci:
	go run "C:\Users\CTN\Documents\CTN\Programming\tmp\unicode-lister\main.go" "$(arg)"
