# Gopher wallpapers
This is a simple API to manage my wallpapers website, in this case, Gophers wallpapers xd. You can download the repository, edit it as you want and use it for a site of your own.

### Clone repository
`git clone https://github.com/Edu-Maturana/wallpapers-go-api.git`

`cd wallpapers-go-api`

### Add the environment variables
Rename .env.example to -> .env

Add your MongoDB Atlas URI  to environments variable
Add your email to environments variables to be the admin

### Run the app
`go run main.go`

### Run the app as a container
`docker build -t app .`

`docker run -it -p 8080:8080 app`
