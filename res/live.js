layui.use(['jquery', 'layer'], function () {
    let $ = layui.$;
    let layer = layui.layer;

    let load = layer.load(1, {shade: 0.6});
    $.ajax({
        url: '${contextPath}/live/query',
        dataType: 'json',
        success: function (res) {
            let m3u8Url = res['m3u8Url'];
            let flvUrl = res['flvUrl'];
            let width = $('.live-main').width();
            let height = width * .625; // 16:10
            new TcPlayer('live-video', {
                "m3u8": m3u8Url,
                "flv": flvUrl,
                "width": width,
                "height": height,
                "live": true,
            });
        },
        error: function () {
            layer.alert('查询失败');
        },
        complete: function () {
            layer.close(load);
        },
    });
});