<div class="container">
  <div class="row">
    <div class="col-md-12 table-responsive">
      <table class="table table-condensed">
        <col style="width:10%">
        <col style="width:90%">
        <thead>
          <tr>
            <th>Name</th>
          </tr>
        </thead>
        <tbody>
        {{range .permissions}}
          <tr>
            <td>{{.Name}}</td>
          </tr>
        {{end}}
        </tbody>
      </table>
      {{template "shared/paginator.html.tpl" .}}
    </div>
  </div>
</div>
