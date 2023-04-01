# Scaff - A Go Web Project Scaffolding Tool

> Scaff is a scaffolding tool designed to help you bootstrap and accelerate your Go web development projects. It provides a well-structured template and a set of utilities that make it easy to create and maintain Go web applications.

### Features
- Pre-configured project structure that follows best practices
- Middleware integration for logging, request validation, and more
- Modular design for easy customization and extensibility
- Support for popular Go web frameworks, such as Gin and Echo
- Detailed documentation and examples
- [TODO] Easy-to-use CLI for project generation and management

### Installation
To install QKGo Scaff, run the following command:

```shell
go get -u github.com/qkgo/scaff
```
After installing, you can access the scaff command from your project.

### Getting Started
Creating a New Project
To create a new project using QKGo Scaff, run the following command:

```shell
myproject/
├─ cmd/
│  ├─ main.go
├─ config/
│  ├─ config.go
├─ internal/
│  ├─ app/
│  │  ├─ controllers/
│  │  ├─ middleware/
│  │  ├─ models/
│  │  ├─ routers/
│  │  ├─ views/
│  ├─ pkg/
├─ .gitignore
├─ go.mod
├─ go.sum
├─ README.md
```
### Running the Project
To run the newly created project, navigate to the project directory and run the following command:
```shell
go run cmd/main.go
```
This will start the Go web server on the default port (8080). You can access the web application by opening a browser and navigating to http://localhost:8080.

### Customizing the Project
QKGo Scaff allows you to customize the generated project to fit your needs. You can modify the project structure, add or remove middleware, and configure the application by editing the config/config.go file.

### Documentation
For detailed documentation and examples, please visit the QKGo Scaff Wiki.

### Contributing
We welcome contributions to QKGo Scaff! If you'd like to contribute, please follow these steps:

### Fork the repository
- Create a new branch for your changes (git checkout -b my-feature)
- Commit your changes (git commit -am 'Add some feature')
- Push your changes to the branch (git push origin my-feature)
- Create a new Pull Request
- Before contributing, please read our Code of Conduct and Contributing Guidelines.

### License
QKGo Scaff is released under the MIT License.

### Support
If you have any questions or need help with QKGo Scaff, please open an issue or join our community chat.

Happy coding!
