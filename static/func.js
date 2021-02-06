var domain = "http://127.0.0.1:8081";

// jquery 入口
$(document).ready(function(){
    // table 列表
    $("#getTablesSpan").click(function(){
        url = domain + "/getTables";
        $.get(url,function(data,status){
            $("aside").html(data);
        })
    });

    // 弹出选项卡
    $("aside").on("mouseenter",".table-name",function(){
        $("#table-option-pop").html(`
            <button type="submit" id="getTableBtn">查看表结构</button>
            <button type="submit" id="getTableRecordBtn">查看数据</button>
        `);
    });

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

    // 更新表数据
    $("#updateTableRecordBtn").click(function(){
        tableName = $(this).parent().parent().parent().attr("id");
        url = domain+"/updateTableRecord";
        $.post(url,
            {
                "tableName": tableName,
                "data": data,
            },function(data,status){
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