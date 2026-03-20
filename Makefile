.PHONY: run wrun asci

run:
	go run .

wrun:
	wgo -file .go -file .bf go run .

asci:
	go run "C:\Users\CTN\Documents\CTN\Programming\tmp\unicode-lister\main.go" "$(arg)"
