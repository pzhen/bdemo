layui.use(['form','laydate'], function(){
    var form = layui.form;
    var laydate = layui.laydate;
    var laypage = layui.laypage;

    laydate.render({
        elem: '#start_time',
        type: 'datetime'
    });

    laydate.render({
        elem: '#end_time',
        type: 'datetime'
    });
});

function showAllContent(o,data) {
    layer.open({
        type: 1,
        area: ['600px', '360px'],
        shadeClose: true, //点击遮罩关闭
        content: '\<\div style="padding:20px;display:block;word-break: break-all;word-wrap: break-word;line-height:22px">'+data+'\<\/div>'
    });
}

// 单条删除
function deleteDataByOne(obj,id)
{
    layer.confirm('确认要执行操作吗？',{btn:['确定','取消']},function(index){
        $.ajax({
            url: decodeURI(deleteUrl),
            data: {"id":id},
            type: "get",
            dataType: "json",
            success: function (data) {
                var messge = "网络繁忙...";
                if(data.Message) {
                    messge = data.Message;
                }
                layer.msg(messge,{icon:1,time:1000},function () {
                    if(data.Code > 0){
                        window.location.href=data.Data;
                    }
                });
            }
        });
        return false;
    });
}

// 批量删除
function deleteDataByBatch()
{
    var idArr = [];
    $(".layui-form-checked").each(function () {
        var currDataId = $(this).attr("data-id");
        if ("undefined" != typeof(currDataId)){
            idArr.push(currDataId)
        }
    });

    if(!idArr.length){
        layer.msg("未选中记录",{icon:1,time:1000});
        return
    }

    layer.confirm('确认要执行操作吗？',{btn:['确定','取消']},function(index){
        $.ajax({
            url: decodeURI(deleteUrl),
            data: {"id":idArr.join()},
            type: "get",
            dataType: "json",
            success: function (data) {
                var messge = "网络繁忙...";
                if(data.Message) {
                    messge = data.Message;
                }

                layer.msg(messge,{icon:1,time:1000},function () {
                    if(data.Code > 0){
                        window.location.href=data.Data;
                    }
                });
            }
        });
        return false;
    });
}