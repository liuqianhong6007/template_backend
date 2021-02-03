<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>{{.Name}}</title>
</head>
<body>

 <table border="1">
   <tr>
     <th>字段名称</th>
     <th>字段类型</th>
     <th>数据类型</th>
     <th>描述</th>
   </tr>

 {{range .Columns}}
   <tr>
     <td>{{.ColumnName}}</td>
     <td>{{.DataType}}</td>
     <td>{{.ColumnType}}</td>
     <td>{{.ColumnComment}}</td>
   </tr>
 {{end}}
 </table>

 <script>

 </script>

</body>
</html>