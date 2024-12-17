.PHONY: gen # generate
gen:
	go mod tidy
	wire
	#for mac sed
	sed -i "" "/go:generate go run/d" ./wire_gen.go
	# for linux sed
	#sed -i "/go:generate go run/d" ./wire_gen.go
	go mod tidy
