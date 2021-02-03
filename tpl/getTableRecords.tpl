<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>表格数据</title>
</head>
<body>

 <table border="1">
 {{range .}}
    <tr>
         {{range .}}
             <td>{{.Value}}</td>
         {{end}}
         <td>你好</td>
    </tr>
 {{end}}
 </table>

 <script>

 </script>

</body>
</html>