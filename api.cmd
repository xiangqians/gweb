@rem 关闭命令回显，且当前行也不显示（@ 符号抑制该行自身的回显），使输出更简洁
@echo off

rem 创建一个局部环境，确保变量只在这个批处理文件中有效
setlocal

rem 输出目录
set OUT_DIR=%cd%
echo OUT_DIR : %OUT_DIR%

rem 通过代码注释生成 API 文档
rem https://github.com/swaggo/swag
rem go install github.com/swaggo/swag/cmd/swag@latest
rem 添加 Go bin 路径（Windows 系统通常是 C:\Users\{用户名}\go\bin）到 Path
swag.exe init --output "%OUT_DIR%" --outputTypes yaml,json > nul
if exist "swagger.yaml" move "swagger.yaml" "api.yaml" > nul
if exist "swagger.json" move "swagger.json" "api.json" > nul

endlocal
pause