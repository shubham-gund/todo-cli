# Todo CLI Project

## Overview
This project is a simple **Todo Command-Line Interface (CLI)** application built in Go. It allows users to manage a list of todo items directly from the terminal. Users can add, edit, delete, toggle completion status, and list all their todos efficiently.

---

## Features

### 1. Add a Todo
Add a new todo item by specifying its title.
```
-add "<title>"
```
Example:
```
./todo -add "Buy groceries"
```

### 2. Edit a Todo
Edit the title of an existing todo by specifying its index and the new title in the format `id:New Title`.
```
-edit "<id>:<new title>"
```
Example:
```
./todo -edit "1:Buy fruits"
```

### 3. Delete a Todo
Delete a todo item by specifying its index.
```
-del <index>
```
Example:
```
./todo -del 2
```

### 4. Toggle Completion Status
Mark a todo item as complete/incomplete by specifying its index.
```
-toggle <index>
```
Example:
```
./todo -toggle 3
```

### 5. List All Todos
Display all the todos in your list with their status (completed/incomplete).
```
-list
```
Example:
```
./todo -list
```

---

## Command-Line Arguments
The CLI uses the following command-line flags:

| Flag       | Description                                      | Example                     |
|------------|--------------------------------------------------|-----------------------------|
| `-add`     | Adds a new todo with the specified title.        | `-add "Go for a walk"`     |
| `-edit`    | Edits the title of a todo by index.              | `-edit "2:Read a book"`    |
| `-del`     | Deletes a todo by index.                        | `-del 1`                   |
| `-toggle`  | Toggles the completion status of a todo by index.| `-toggle 3`                |
| `-list`    | Lists all todos in the current list.             | `-list`                    |

---

## Error Handling
- Invalid commands or missing arguments will print an error message and terminate the program.
- When editing a todo, the index must be valid and provided in the `id:New Title` format.

---

## Code Functionality

### `NewCmdFlags`
This function initializes and parses the command-line flags.

### `Execute`
The `Execute` method processes the parsed flags and performs the corresponding action on the todo list.
- **List Todos:** Calls `todos.print()`.
- **Add Todo:** Calls `todos.add(title)`.
- **Edit Todo:** Splits the input into `id` and `New Title`, validates them, and calls `todos.edit()`.
- **Toggle Todo:** Calls `todos.toggle(index)`.
- **Delete Todo:** Calls `todos.delete(index)`.
- **Invalid Command:** Prints an error message.

---

## Example Workflow
1. Add todos:
   ```
   ./todo -add "Write documentation"
   ./todo -add "Review PRs"
   ```

2. List todos:
   ```
   ./todo -list
   ```
   Output:
   ```
   1. Write documentation [Incomplete]
   2. Review PRs [Incomplete]
   ```

3. Toggle a todo's status:
   ```
   ./todo -toggle 1
   ```
   Output:
   ```
   1. Write documentation [Complete]
   2. Review PRs [Incomplete]
   ```

4. Edit a todo:
   ```
   ./todo -edit "2:Review project"
   ```
   Output:
   ```
   1. Write documentation [Complete]
   2. Review project [Incomplete]
   ```

5. Delete a todo:
   ```
   ./todo -del 1
   ```
   Output:
   ```
   1. Review project [Incomplete]
   ```

---

## Requirements
- **Go:** Ensure you have Go installed on your system to compile and run the application.

---

## Running the Application
1. Clone the repository.
2. Build the executable:
   ```
    go build -o todo main.go
   ```
3. Run the commands as shown in the examples above.

---

## Future Enhancements
- Add persistence using a database or a file.
- Implement a search/filter feature for todos.
- Provide support for due dates and priorities.

---

## Contribution
Contributions are welcome! Feel free to open issues or submit pull requests.

