<ul>
  {{range .Errors }}
  <li>
    <mark>{{ .Key  }}</mark>
    {{ .Message  }}
  </li>
  {{end}}
</ul>
