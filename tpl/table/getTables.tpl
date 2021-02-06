<div id="left-frame">
    {{range .}}
    <span class="tableName">{{.Name}}</span>
    <button type="submit" id="getTableBtn">查看表结构</button>
    <button type="submit" id="getTableRecordBtn">查看数据</button>
    {{end}}
</div>
</br>
<div id="right-frame">
</div>
<script>
    var domain = "http://127.0.0.1:8081";

    $(document).ready(function(){
        // 查看表结构
        $("#getTableBtn").click(function(){
            tableName = $(this).parent().find(".tableName").text();
            url = domain+"/getTable?tableName=" + tableName;
            $.get(url,function(data,status){
                $("#right-frame").html(data);
            });
        });

        // 查看表数据
        $("#getTableRecordBtn").click(function(){
            tableName = $(this).parent().find(".tableName").text();
            url = domain+"/getTableRecords?tableName=" + tableName
            $.get(url,function(data,status){
                $("#right-frame").html(data);
            });
        });

        // 表格样式
        $("tr").hover(function(){
            $(this).css("background-color","yellow");
        },function(){
             $(this).css("background-color","white");
        });
     });
</script>
<style>
    #right-frame {
        margin-top: 20px;
    }
    .tableName{
        background-color: MediumAquaMarine;
    }
</style>