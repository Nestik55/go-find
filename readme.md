## go-find
##### It's my command line tool.
It have same performs same functoin as *find* in **Bash**.
You can use it like this:
Go to the directory from which you want to start the search.
```bash
go run directory_with_file/pkg/main.go name_file_you_are_looking_for
```
You will get absolute paths to all files with the name_file_you_are_looking_for.
It works much slower *find* in **Bash**. It is just learning task.
I use dfs as algorithm.


