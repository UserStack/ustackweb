<div class="container">
  <div class="row">
    <div class="col-md-offset-3 col-md-6">
      <h1>Statistics</h1>
    </div>
  </div>
  <div class="row">
    <div class="col-md-offset-3 col-md-6">
      <div class="list-group">
      {{range $key, $val := .stats}}
        <div class="list-group-item">
          {{$key}}
          <span class="badge">{{$val}}</span>
        </div>
      {{end}}
      </div>
    </div>
  </div>
</div>
