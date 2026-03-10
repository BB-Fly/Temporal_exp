# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Temporal Go SDK example project demonstrating workflow patterns. The module name is `temporal-exp`.

## Commands

### Prerequisites
- Start Temporal server: `temporal server start-dev`
- Temporal listens on `localhost:7233` by default
- Web UI available at `http://localhost:8233`

### Running the Application
```bash
# Terminal 1: Start the Worker
go run src/worker/main.go

# Terminal 2: Execute a workflow
go run src/start/main.go <name>
```

## Architecture

This project follows the standard Temporal Go SDK pattern with these components:

### Activity (`src/greating/activity/`)
- Contains business logic that executes in the workflow
- Functions take `context.Context` as first parameter
- Activities are registered with the worker

### Workflow (`src/greating/workflow/`)
- Orchestrates activity execution
- Uses `workflow.Context` (not `context.Context`)
- Configures `ActivityOptions` with timeouts before executing activities
- Activities are called via `workflow.ExecuteActivity()`

### Worker (`src/worker/main.go`)
- Registers workflows and activities
- Listens on a task queue (`"my-task-queue"`)
- Processes tasks from the Temporal server

### Starter (`src/start/main.go`)
- Creates a Temporal client
- Starts workflow executions with `client.ExecuteWorkflow()`
- Waits for results with `we.Get()`

## Key Patterns

- **Task Queue**: All components use `"my-task-queue"` - this must match between worker and starter
- **Activity Options**: Always set `StartToCloseTimeout` when executing activities
- **Workflow ID**: Each workflow instance has a unique ID (`"greeting-workflow"`)

## Dependencies

- `go.temporal.io/sdk v1.40.0` - Main Temporal Go SDK
