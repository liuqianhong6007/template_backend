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
         <button type="submit" onclick="getTable({{.Name}})">查看表结构</button>
         <button type="submit" onclick="getTableRecords({{.Name}})">查看数据</button>
     </td>
   </tr>
 {{end}}
 </table>

 <script>
    var xmlhttp = new XMLHttpRequest();
    var domain = "http://127.0.0.1:8081";

    function getTable(tableName){
        url = domain+"/getTable?tableName=" + tableName + "&t=" + Math.random()
        xmlhttp.open("GET",url,true);
        xmlhttp.send();
        location.replace(url);
    }

    function getTableRecords(tableName){
        url = domain+"/getTableRecords?tableName=" + tableName + "&t=" + Math.random()
        xmlhttp.open("GET",url,true);
        xmlhttp.send();
        location.replace(url);
    }

 </script>

</body>
</html>