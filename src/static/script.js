$(document).ready(function() {
    var socket = io.connect();

    var canv_width = 600;
    var canv_height = 600;
    var c = $("#map").get(0);
    var ctx = c.getContext("2d");

    function clearCanvas(ctx) {
        ctx.fillStyle = "#FFFFFF";
        ctx.fillRect(0, 0, canv_width, canv_height);
        ctx.fillStyle = "#0045CC";
        ctx.fillRect((canv_width - 80) / 2, (canv_height - 80) / 2, 80, 80);
    }

    socket.on('telemetry', function(message) {
        var j = JSON.parse(message);
        clearCanvas(ctx);

        // front
        ctx.fillRect((canv_width - 80) / 2, ((canv_height - 80) / 2) - j['5'], 80, 5);

        // right
        ctx.fillRect(((canv_width - 80) / 2) + j['8'], ((canv_height - 80) / 2), 5, 80);

        // left
        ctx.fillRect(((canv_width - 80) / 2) - j['2'], ((canv_height - 80) / 2), 5, 80);

        // rear
        ctx.fillRect((canv_width - 80) / 2, ((canv_height - 80) / 2) + j['0'], 80, 5);
    });

});
