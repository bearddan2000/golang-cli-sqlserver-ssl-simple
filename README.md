# golang-cli-sqlserver-ssl-simple

## Description
A POC for connecting a database to golang.

Sql server uses self-signed ssl.

## Tech stack
- bash
- golang 1.13

## Docker stack
- alpine:edge
- ubuntu:latest
- mcr.microsoft.com/mssql/server:2017-latest-ubuntu

## To run
`sudo ./install.sh -u`

## To stop
`sudo ./install.sh -d`

## To see help
`sudo ./install.sh -h`

## Credits
http://jenicaandpatrick.com/how-to-query-a-ms-sql-server-database-in-go-golang/
