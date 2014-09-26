<div class="container">
  <div class="row">
    <div class="col-md-12 table-responsive">
      <table class="table table-condensed">
        <col style="width: 8.3333%">
        <col style="width:91.6667%">
        <thead>
          <tr>
            <th>#</th>
            <th>Name</th>
          </tr>
        </thead>
        <tbody>
        {{range $key, $val := .groups}}
          <tr>
            <td>{{$val.Gid}}</td>
            <td>{{$val.Name}}</td>
          </tr>
        {{end}}
        </tbody>
      </table>
      {{template "shared/paginator.html.tpl" .}}
    </div>
  </div>
</div>
