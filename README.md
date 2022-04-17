# $doing

Doing is a command line tool to track from the terminal the tasks being done.

## Available commands

- List All: Lists all tasks doing and completed.

```bash
doing -la
```

- List Doing: List all doing tasks.

```bash
doing -la
```

- List Completed: Lists all completed tasks.

```bash
doing -lc
```

- Add new task: Adds new task.

```bash
doing -a [description]
```

- Mark task as completed: Marks a task of index [i] in list all as completed.

```bash
doing -c [i]
```

- Delete task: Deletes a task of index [i] in list all.

```bash
doing -d [i]
```

- Reset: deletes all stored tasks.

```bash
doing -r
```

## Considerations

- doing -c command accepts an index from last doing tasks, 0, 1, 2
- Tasks fall on two categories: doing and completed. Deleted tasks are not recoverable.
- Tasks are saved on a store text file defined in the saving package.
Default is store/tasks.txt
- Tasks are given an id, based on the task count defined on the store text file.
