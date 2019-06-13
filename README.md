# Grpc-based audit logger microservice

## Launch

To launch the microservice you need to have Docker installed.

Sample links:

- Ubuntu: https://docs.docker.com/install/linux/docker-ce/ubuntu/
- Windows: https://docs.docker.com/docker-for-windows/
- Mac: https://docs.docker.com/docker-for-mac/

The binary, containing the latest version of the application, is already already build and placed into the right folder.

Before the launch you can set global variables, that are being used by the microservice:
- MAIN_PORT (default: 7777);
- DB_LOGIN (default: postgres);
- DB_PASSWORD (default: postgres);
- DB_NAME (default: logger);
- DB_PORT (default: 5875).

Make sure that the ports are free.

Run the bash script run.sh in the folder /devops

## Binary rebuild

Run the bash script build.sh in the folder /src
