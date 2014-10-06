# ustackweb

Web frontend for [ustackd](https://github.com/UserStack/ustackd) written in [Go](http://go-lang.org) using the [Beego](http://beego.me/) web framework.

Project to learn golang.org in the context of user lifecycle management.

![https://travis-ci.org/UserStack/ustackweb.svg?branch=master](https://travis-ci.org/UserStack/ustackweb.svg?branch=master)

## Audience

* End User
* Administrator
* Developer

## Features

* User Account
    * CRUD
    * Lock/Unlock
    * Activation
    * Reset Password
    * Assign/Unassign Groups
    * Assign/Unassign Permissions
* Groups
    * CRUD
    * Assign/Unassign Permissions
* Permissions
    * CRUD

## Defining Permissions

A permission is basically nothing more than a group. The user has the permission if he belongs to the group.

The naming scheme for permission groups is the following:

        perm.{object}.{verb}

To secure all CRUD operations of user objects we would need to create the following groups:

        "perm.users.read"
        "perm.users.create"
        "perm.users.update"
        "perm.users.delete"

Within the `ustackweb` project this permissions can then be used as little DSL via:

        can["read_users"]
        can["create_users"]
        can["update_users"]
        can["delete_users"]

## Development

Please make sure you have [Go](http://golang.org/), [Ruby](http://ruby-lang.org) (for [Sass](http://sass-lang.com/)), [Node.js](http://nodejs.org/) and [npm](https://github.com/npm/npm) (for asset management via [Bower](http://bower.io) and  [Livereload](http://livereload.com/)) installed.

1. Setup dependencies

        $ make prepare 

2. Run the Beego frontend on [http://localhost:8080](http://localhost:8080)

        $ make run

3. [Livereload](http://livereload.com/) and recompile assets

        $ make watch

4. Open the Install page to setup some data

        $ open http://localhost:8080/install

5. Sign In using the root user (user=admin, pass=admin)

        $ open http://localhost:8080/sign_in

## Todos

* Bulk editing (users, groups, permissions)
* https://datatables.net vs ng-table ?

### Technical Debt

* Form builder
* Form validation, action validation
* Flash message across redirect

## Alternatives

 * [Stormpath](https://stormpath.com/) (commercial)
 * [Apache Shiro](http://shiro.apache.org/)
 * [Apache Syncope](Apache Syncope)
 * [ConnId](http://connid.tirasa.net/)

## License

See [LICENSE](LICENSE).


