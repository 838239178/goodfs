cd %~dp0..\src\adminserver && go generate -x ./...
cd %~dp0..\src\apiserver && go generate -x ./...
cd %~dp0..\src\common && go generate -x ./...
cd %~dp0..\src\metaserver && go generate -x ./...
cd %~dp0..\src\objectserver && go generate -x ./...