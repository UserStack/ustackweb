<!DOCTYPE html>
<html>
    <head>
      <title>UserStack</title>
      <meta http-equiv="Content-Type" content="text/html; charset=utf-8">

      <meta name="_xsrf" content="{{.xsrf_token}}" />
      <!-- bower:css -->
      <!-- endbower -->
      <link rel="stylesheet" href="/static/css/app.css" />
      <link rel=icon href="/static/images/logo.ico" sizes="16x16 32x32 48x48 64x64" type="image/vnd.microsoft.icon">
      <link rel=icon href="/static/images/logo16.png" sizes="16x16" type="image/png">
      <link rel=icon href="/static/images/logo32.png" sizes="32x32" type="image/png">
      <link rel=icon href="/static/images/logo48.png" sizes="48x48" type="image/png">
      <link rel=icon href="/static/images/logo64.png" sizes="64x64" type="image/png">
  </head>
    <body class="{{.context.ControllerName}} {{.context.ActionName}}">
      {{if .loggedIn}}
      <nav class="navbar navbar-default navbar-static-top" role="navigation">
        <div class="container">
          <div class="navbar-header">
            <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1">
              <span class="sr-only">Toggle navigation</span>
              <span class="icon-bar"></span>
              <span class="icon-bar"></span>
              <span class="icon-bar"></span>
            </button>
            <a class="navbar-brand" href="#"></a>
          </div>
          <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
            <ul class="nav navbar-nav">
              <li {{if .context.Is "HomeController.Get" }}class="active"{{end}}><a href="{{urlfor "HomeController.Get"}}">Home</a></li>
              {{if .can.list_users}}
              <li {{if .context.Is "UsersController.Index" }}class="active"{{end}}><a href="{{urlfor "UsersController.Index"}}">Users</a></li>
              {{end}}
              {{if .can.list_groups}}
              <li {{if .context.Is "GroupsController.Index" }}class="active"{{end}}><a href="{{urlfor "GroupsController.Index"}}">Groups</a></li>
              {{end}}
              {{if .can.list_permissions}}
              <li {{if .context.Is "PermissionsController.Index" }}class="active"{{end}}><a href="{{urlfor "PermissionsController.Index"}}">Permissions</a></li>
              {{end}}
              {{if .can.list_stats}}
              <li {{if .context.Is "StatsController.Index" }}class="active"{{end}}><a href="{{urlfor "StatsController.Index"}}">Statistics</a></li>
              {{end}}
            </ul>
            <ul class="nav navbar-nav navbar-right">
              <li {{if .context.Is "ProfileController.Get" }}class="active"{{end}}><a href="{{urlfor "ProfileController.Get"}}">Profile</a></li>
              <li><a href="{{urlfor "SessionsController.Destroy"}}">Sign Out</a></li>
            </ul>
          </div>
        </div>
      </nav>
      <div class="container">
        {{end}}
        {{if .flash.error}}
          <div class="alert alert-danger">{{.flash.error}}</div>
        {{end}}
        {{if .flash.notice}}
          <div class="alert alert-info">{{.flash.notice}}</div>
        {{end}}
      </div>
      {{.LayoutContent}}
      <!-- bower:js -->
      <script src="../../static/bower_components/jquery/dist/jquery.js"></script>
      <script src="../../static/bower_components/bootstrap-sass/dist/js/bootstrap.js"></script>
      <!-- endbower -->
      <script src="../../static/scripts/app.js"></script>
      {{if eq .RunMode "dev"}}<script src="//localhost:35729/livereload.js"></script>{{end}}
    </body>
</html>
