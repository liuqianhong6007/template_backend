<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>表格</title>
</head>
<body>

 <table border="1">
   <tr>
     <th>表名称</th>
     <th>描述</th>
     <th>操作</th>
   </tr>

 {{range .}}
   <tr>
     <td>{{.Name}}</td>
     <td>{{.Comment}}</td>
     <td>
         <button type="submit" onclick="getTable({{.Name}})">查看</button>
     </td>
   </tr>
 {{end}}
 </table>

 <script>
    function getTable(tableName){
        xmlhttp.open("GET","/getTable?tableName=" + tableName + "&t=" + Math.random());
        xmlhttp.send();
    }
 </script>

</body>
</html>