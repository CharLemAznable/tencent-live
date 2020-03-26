layui.use(['jquery', 'layer'], function () {
    let $ = layui.$;
    let layer = layui.layer;

    let load = layer.load(1, {shade: 0.6});
    $.ajax({
        url: '${contextPath}/push/query',
        dataType: 'json',
        success: function (res) {
            $('.push-url').html(res['pushUrl']);
            $('.push-server-url').html(res['pushServerUrl']);
            $('.push-stream-query').html(res['pushStreamQuery']);
        },
        error: function () {
            layer.alert('查询失败');
        },
        complete: function () {
            layer.close(load);
        },
    });

    let $main = $('.live-main');

    $main.on('click', '.live-copy>button', function () {
        let content = $(this).parents('.push-title-row')
            .next('.push-content-row').children('.live-url')[0].innerText;
        clip(content);
    });

    function clip(content) {
        let aux = document.createElement("input");
        aux.setAttribute("value", content);
        document.body.appendChild(aux);
        aux.select();
        document.execCommand("copy");
        document.body.removeChild(aux);
    }
});