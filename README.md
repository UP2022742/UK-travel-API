# API-WebAPI-Template
A template of communication between API and Web server in GoLang.

## Windows
You'll need GoLang, to install it for linux follow the article here:

https://golang.org/doc/install

> :warning: **Assuming you're in the root folder of the repository.**
```
cd WorkerService
go build main.go -o WorkerService
.\WorkerService.exe
```
Open another terminal and navigate back to the root of the repository.
```
cd WebServer
go build main.go -o WebServer
.\WebServer.exe
```

## Linux
As per windows you will need GoLang, to install it follow the guide here and select the Linux variation:

https://golang.org/doc/install

> :warning: **Assuming you're in the root folder of the repository.**
```
cd WorkerService && go build main.go -o WorkerService && ./WorkerService
```
Open another terminal and navigate back to the root of the repository.
```
cd WebServer && go build main.go -o WebServer && ./WebServer
```
