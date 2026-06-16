.PHONY: run new help

help:
	@echo "Usage:"
	@echo "  make run 01      - Run problem 01"
	@echo "  make new 02      - Create new problem 02"

run:
	@if [ -z "$(word 2,$(MAKECMDGOALS))" ]; then echo "Error: specify problem number (e.g., make run 01)"; exit 1; fi
	dotnet run --project p$(word 2,$(MAKECMDGOALS))

new:
	@if [ -z "$(word 2,$(MAKECMDGOALS))" ]; then echo "Error: specify problem number (e.g., make new 02)"; exit 1; fi
	dotnet new console -o p$(word 2,$(MAKECMDGOALS))

%:
	@:
