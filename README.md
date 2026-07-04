# TaskInventory 📝
Meant to be a simple CLI task manager and scheduler made in Go.
It uses a local, on-disk SQLite database to store the tasks, so no data will be collected.

## How to use? 📖
After opening the app, you can do any of the following actions:

- list                          -> List all your tasks
- add [taskname] [description]  -> Add a new task
- remove [taskname]             -> Remove task with given name
- complete [taskname]           -> Mark the task as completed

## Project goals 🎯
As my first Go project, TaskInventory is meant to be a simple yet useful alternative to traditional task managers. Rather than relying on frameworks or third-party libraries, I want to understand how things work under the hood and build a solid foundation in Go.

## What I'm aiming to learn with this project 🧠
  - [ ] Simple Go syntax - I want to make clean, fast code in Go.  
  - [ ] Database integration in Go - Learn how Go connects to databases without using frameworks like Cobra or other libraries.  
  - [ ] Testing – Improve my testing skills by writing meaningful and maintainable tests that help build scalable applications.

## Future goals in this project 🗓️
- Adding flags (--c/--n) to the lsit command so you can list only completed or only in-progress task
- Add an option to add a date limit to the tasks - maybe even hour/minute limit (ex. --d "06/03/26" --h "14:57")
- Make list mark completed task in green and task that exceed the date limit in red
