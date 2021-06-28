# dokku-go-example

Easily deploy your Go applications with Dokku.

Features:
- Deploy on your own server
- Auto deployment
- HTTPS

**Check the full step by step article: https://shellbear.me/blog/go-dokku-deployment**

## 💻 Getting started

The example API requires a PostgreSQL database.

Specify the database connection URL using the `DATABASE_URL` environment variable:

```shell
export DATABASE_URL=postgresql://user:secret@localhost
go run .
```

## Built with

- [Gorm](https://gorm.io)
- [Echo](https://echo.labstack.com)


