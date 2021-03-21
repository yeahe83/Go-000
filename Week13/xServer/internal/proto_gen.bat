@echo off

call :gen auth
call :gen hisdata
call :gen analysis
call :gen report

goto :EOF

:gen

set DOMAIN=%1
set PROTO_PATH=.\%DOMAIN%\api
set PROTO_OUT_PATH=.\%DOMAIN%\api\gen\v1

if not exist %PROTO_OUT_PATH% (
    mkdir %PROTO_OUT_PATH%
    @REM echo %PROTO_OUT_PATH% created.
) else (
    @REM echo %PROTO_OUT_PATH% exists.
)

if not exist %PROTO_PATH%\%DOMAIN%.proto (
    echo %DOMAIN%.proto not exists.
) else (
    protoc -I %PROTO_PATH% ^
        --go_out %PROTO_OUT_PATH% ^
        --go_opt paths=source_relative ^
        --go-grpc_out %PROTO_OUT_PATH% ^
        --go-grpc_opt paths=source_relative ^
        %DOMAIN%.proto

    protoc -I %PROTO_PATH% ^
        --grpc-gateway_out %PROTO_OUT_PATH% ^
        --grpc-gateway_opt paths=source_relative ^
        --grpc-gateway_opt grpc_api_configuration=%PROTO_PATH%\%DOMAIN%.yaml ^
        %DOMAIN%.proto

    echo protoc %DOMAIN%.proto done.
)