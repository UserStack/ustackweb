<!DOCTYPE html>

<html>
    <head>
      <title>UserStack</title>
      <meta http-equiv="Content-Type" content="text/html; charset=utf-8">

    <style type="text/css">
      body {
        font-family: "Helvetica Neue",Helvetica,Arial,sans-serif;
      }
    </style>
  </head>
    <body>
      <ol>
        <li><a href="/">Home</a></li>
        <li><a href="/profile">Profile</a></li>
        <li><a href="/sessions/destroy">Logout</a></li>
      </ol>
      {{.LayoutContent}}
    </body>
</html>
