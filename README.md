# CLI based task manager - Golang

Create and manage tasks in the CLI like a true programmer

### Add a task

```bash
go run ./cmd/tasks -add Learn go
go run ./cmd/tasks -add Learn microservices
go run ./cmd/tasks -add Take a nap
```

### List down the tasks

```bash
go run ./cmd/tasks -list

# output

╔═════╤═════════════════════╤═══════╤═════════════════════╤═════════════════════╗
║ #ID │        Task         │ Done? │          Created at │        Completed at ║
╟━━━━━┼━━━━━━━━━━━━━━━━━━━━━┼━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━╢
║ 1   │ Learn go            │ false │ 16 Feb 23 01:33 IST │ 01 Jan 01 00:00 UTC ║
║ 2   │ Learn microservices │ false │ 16 Feb 23 01:33 IST │ 01 Jan 01 00:00 UTC ║
║ 3   │ Take a nap          │ false │ 16 Feb 23 01:33 IST │ 01 Jan 01 00:00 UTC ║
╟━━━━━┼━━━━━━━━━━━━━━━━━━━━━┼━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━╢
║                         You have 3 incomplete tasks!                          ║
╚═════╧═════════════════════╧═══════╧═════════════════════╧═════════════════════╝
```

### Mark a task `done`

```bash
go run ./cmd/tasks -done=1

go run ./cmd/tasks -list

# output

╔═════╤═════════════════════╤═══════╤═════════════════════╤═════════════════════╗
║ #ID │        Task         │ Done? │          Created at │        Completed at ║
╟━━━━━┼━━━━━━━━━━━━━━━━━━━━━┼━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━╢
║ 1   │ ✅ Learn go         │ true  │ 16 Feb 23 01:33 IST │ 16 Feb 23 01:35 IST ║
║ 2   │ Learn microservices │ false │ 16 Feb 23 01:33 IST │ 01 Jan 01 00:00 UTC ║
║ 3   │ Take a nap          │ false │ 16 Feb 23 01:33 IST │ 01 Jan 01 00:00 UTC ║
╟━━━━━┼━━━━━━━━━━━━━━━━━━━━━┼━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━╢
║                         You have 2 incomplete tasks!                          ║
╚═════╧═════════════════════╧═══════╧═════════════════════╧═════════════════════╝
```

### Delete a task

```bash
go run ./cmd/tasks -delete=1
go run ./cmd/tasks -list

# output

╔═════╤═════════════════════╤═══════╤═════════════════════╤═════════════════════╗
║ #ID │        Task         │ Done? │          Created at │        Completed at ║
╟━━━━━┼━━━━━━━━━━━━━━━━━━━━━┼━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━╢
║ 1   │ Learn microservices │ false │ 16 Feb 23 01:33 IST │ 01 Jan 01 00:00 UTC ║
║ 2   │ Take a nap          │ false │ 16 Feb 23 01:33 IST │ 01 Jan 01 00:00 UTC ║
╟━━━━━┼━━━━━━━━━━━━━━━━━━━━━┼━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━╢
║                         You have 2 incomplete tasks!                          ║
╚═════╧═════════════════════╧═══════╧═════════════════════╧═════════════════════╝
```
