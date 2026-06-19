.PHONY: run new link test help

SHARED := Shared/Shared.csproj
SHARED_TESTS := Shared.Tests/Shared.Tests.csproj

help:
	@echo "Usage:"
	@echo "  make run 01         - Run problem 01"
	@echo "  make new 02         - Create new problem 02"
	@echo "  make new 02 shared  - Create new problem 02 and reference the shared library"
	@echo "  make link 02        - Add the shared library reference to an existing problem 02"
	@echo "  make test           - Run the shared library tests"

test:
	dotnet test $(SHARED_TESTS)

run:
	@if [ -z "$(word 2,$(MAKECMDGOALS))" ]; then echo "Error: specify problem number (e.g., make run 01)"; exit 1; fi
	dotnet run --project p$(word 2,$(MAKECMDGOALS))

new:
	@if [ -z "$(word 2,$(MAKECMDGOALS))" ]; then echo "Error: specify problem number (e.g., make new 02)"; exit 1; fi
	dotnet new console -o p$(word 2,$(MAKECMDGOALS))
	@if [ "$(word 3,$(MAKECMDGOALS))" = "shared" ]; then \
		dotnet add p$(word 2,$(MAKECMDGOALS)) reference $(SHARED); \
	fi

link:
	@if [ -z "$(word 2,$(MAKECMDGOALS))" ]; then echo "Error: specify problem number (e.g., make link 02)"; exit 1; fi
	dotnet add p$(word 2,$(MAKECMDGOALS)) reference $(SHARED)

%:
	@:
