# DjanGolang

Development ready Go + [GORM](https://github.com/go-gorm/gorm) + React/TypeScript + Postgres + Docker project skeleton with hot loading enabled for both server and client — mainly for educational purposes. This app runs in three separate containers with the following names according to the docker-compose.yml specifications:
 - djangolang_go
 - djangolang_react
 - djangolang_pg

# Setup 🔧
Minimal setup thanks to Docker.

**Note to Windows users**: 
I tried setting this up on Windows 10 Home using the [Docker Toolbox](https://docs.docker.com/toolbox/) but ran into some hiccups because of Windows things so unfortunately setting this project skeleton up as-is with Docker Toolbox (the older version of Docker) won't work.
### Install Docker
Download [Docker & Docker-compose](https://www.docker.com/products/docker-desktop). Verify that they're properly installed by running `docker -v` and `docker-compose -v`.
### Development environment
First, rename the `.env-example` in the server directory into `.env`. Then while in your project's root directory, pop this command into your terminal: `docker volume create --name=djangolang_pg_data`. Next, `cd client` and run `yarn` or `npm install`. Follow up by returning to your project root directory and boot up all of your containers by running `docker-compose up`, then visit localhost:3000 to see your app running. Now you're ready to develop :]

*Optional Read*:
`docker volume create --name=djangolang_pg_data` creates the volume that will allow you to mount your local **pg_data** directory in your project's root directory to **/var/lib/postgresql/data** in the container—which is where your postgres data lives—allowing data to persist between the two. 

# Note from Author
The name DjanGolang takes inspiration from the Python web-framework, Django, but nothing beyond that. Issues and pull requests are welcome!