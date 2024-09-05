C\GoApp\Src> go get golang.org/x/sys/windows/svc

C\GoApp\Src> go build dummyService.go
C\GoApp\Src> sc create DummyService binPath= c:\goapp\src\dummyService.exe

C\GoApp\Src> sc delete DummyService