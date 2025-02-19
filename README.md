# Go Race Condition Example

This repository demonstrates a common race condition in Go.  The program attempts to increment a counter in multiple goroutines without proper synchronization, resulting in an incorrect final count. This example highlights the importance of using synchronization primitives like mutexes or atomic operations when dealing with shared resources accessed concurrently.

## How to Run

1.  Clone the repository:
    ```bash
git clone <repository_url>
```
2.  Navigate to the repository directory:
    ```bash
cd <repository_directory>
```
3. Run the program:
    ```bash
go run bug.go
```
Observe that the output is often less than 1000, demonstrating the race condition.

## Solution

The solution involves protecting the `count` variable with a mutex.