# ustackweb

Web frontend for [ustackd](https://github.com/UserStack/ustackd) written in [Go](http://go-lang.org) using the [Beego](http://beego.me/) web framework.

## Audience

* User (signup, login, reset password, ...)
* Customer/Administrator/Developer (activate, lock, notify, logs, API-calls, test, logs, debugging, ...)

## Development

1. Setup dependencies

        $ make prepare 

2. Run the Beego frontend on [http://localhost:8080](http://localhost:8080)

        $ make run

3. [Livereload](http://livereload.com/) and recompile assets

        $ make watch

## Todos

* Bulk editing (users, groups, permissions)

## License

See [LICENSE](LICENSE).


