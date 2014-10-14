<div class="container">
  <div class="row">
    <div class="col-md-12">
      <form class="form-horizontal">
        <div class="form-group">
          <div class="col-md-offset-3 col-md-6">
            <h1>Profile of {{.username }}</h1>
          </div>
        </div>
        <div class="form-group">
          <div class="col-md-offset-3 col-md-6">
            <div class="list-group">
            {{range $key, $val := .userData}}
              <div class="list-group-item">
                {{$key}}
                {{if isDate $val}}
                <span class="badge">{{dateformat $val "Jan 2, 2006 15:04 (MST)"}}</span>
                {{else}}
                <span class="badge">{{$val}}</span>
                {{end}}
              </div>
            {{end}}
            </div>
          </div>
        </div>
      </form>
    </div>
  </div>
</div>
