switchmap:
	@ go test ./switch_vs_map/... -bench=. -cpuprofile=./switch_vs_map/cpu.prof
graph_switchmap: switchmap
	@ go tool pprof -http=:8080 ./switch_vs_map/cpu.prof