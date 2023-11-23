![GO](https://img.shields.io/badge/GO-47d1ed?style=for-the-badge&logo=go&logoColor=white)

<h1 align="center"> 👋 GO-GIN REST API</h1>

# GO-GIN REST API
Rest APi using Go's Gin! and gorilla mux

## Backend
<p align="left">
  <a href="https://skillicons.dev">
    <img src="https://skillicons.dev/icons?i=go" />
  </a>
</p>

## Included imports
* Gorilla mux
* Gorm
* Gin
* MySQL driver

## Initialize
<details>
  <summary>Initialize Go's API</summary>
  <p>Get to the directory</p>
  <p>Configure your server you want to listen, by default is in 8080</p>
  <p>go run main.go</p>
</details>
<details>
  <summary>Running migration</summary>
  <p>go run migration/migrate.go</p>
  <p>If you add more fields to the model, you have to update the controller as well</p>
</details>
<details>
  <summary>DB connection</summary>
  <p>```go
      dsn := "user:@password/database?charset=utf8mb4&parseTime=True&loc=Local"
  ```</p>
  <p>If you're not using any password just leave it this way:@/database</p>
</details>
<details>
  <summary>Testing insomnia/postman/thunder client</summary>
  <p>localhost:8080/api/v1/users</p>
</details>

## Contributing

1. Fork it!
2. Create your feature branch: `git checkout -b your-branch`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin your-branch`
5. Submit a PR

<h3 align="left">Get in touch with me:</h3>
<p align="left">
<a href="https://www.linkedin.com/in/adhir-serrano/" target="blank"><img align="center" src="https://raw.githubusercontent.com/rahuldkjain/github-profile-readme-generator/master/src/images/icons/Social/linked-in-alt.svg" alt="adhir2099" height="30" width="40" /></a>
</p>
<p align="right" > Created with 🧡 by <a href="https://github.com/adhir2099">Adhir2099</a></p>