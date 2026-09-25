package commands

import (
	"fmt"
	"os"
)

func Execute() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "workload":
		runWorkload(os.Args[2:])

	case "worker":
		// runWorker(os.Args[2:])

	case "benchmark":
		// runBenchmark(os.Args[2:])

	case "result":
		// runResult(os.Args[2:])

	case "system":
		// runSystem(os.Args[2:])

	case "help", "--help", "-h":
		printUsage()

	default:
		fmt.Printf("tsb: unknown command: %s %s\n\n", os.Args[0], os.Args[1])
		fmt.Printf("run 'tsb help' for more information\n")
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println(`Usage: tsb <command> [options]

Commands:
  workload    Manage workloads
  worker      Manage workers
  benchmark   Run and manage benchmarks
  result      Inspect benchmark results
  system      System information
  help        Print information about the command`)
}
