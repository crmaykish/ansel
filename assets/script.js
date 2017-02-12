const CANVAS_HEIGHT = 600;
const CANVAS_WIDTH = 600;
const ANSEL_HEIGHT = 25;
const ANSEL_WIDTH = 20;

var socket;
var c;
var ctx;

function init() {
    socket = io.connect();

    c = $("#map").get(0);
    ctx = c.getContext("2d");
}

// Return the center x-coord for the given width
function c_x(width) {
    return (CANVAS_WIDTH - width) / 2;
}

// Return the center y-coord for the given height
function c_y(height) {
    return (CANVAS_HEIGHT - height) / 2;
}

// Convert an angle in degrees to radians
function toRad(degrees) {
    return degrees * (Math.PI / 180);
}

function clearCanvas(ctx) {
    ctx.fillStyle = "#CFCFCF";
    ctx.fillRect(0, 0, CANVAS_WIDTH, CANVAS_HEIGHT);
    ctx.fillStyle = "#80CC35";
    ctx.fillRect(c_x(ANSEL_WIDTH), c_y(ANSEL_HEIGHT), ANSEL_WIDTH, ANSEL_HEIGHT);
}

function drawBoundaries(ctx, jsonData) {
    // Front Left
    frontLeft = jsonData["4"];
    w = frontLeft * Math.sin(toRad(45));
    h = frontLeft * Math.cos(toRad(45));
    ctx.beginPath();
    ctx.moveTo(c_x(0) - ANSEL_WIDTH / 2, c_y(0) - ANSEL_HEIGHT / 2);
    ctx.lineTo(c_x(w) - ANSEL_WIDTH / 2, c_y(h) - ANSEL_HEIGHT / 2);
    ctx.stroke();

    // Front
    front = jsonData["5"];
    ctx.beginPath();
    ctx.moveTo(c_x(0), c_y(0) - ANSEL_HEIGHT / 2);
    ctx.lineTo(c_x(0), c_y(front) - ANSEL_HEIGHT / 2);
    ctx.stroke();

    // Front Right
    frontRight = jsonData["6"];
    h = frontRight * Math.sin(toRad(45));
    w = frontRight * Math.cos(toRad(45));
    ctx.beginPath();
    ctx.moveTo(c_x(0) + ANSEL_WIDTH / 2, c_y(0) - ANSEL_HEIGHT / 2);
    ctx.lineTo(c_x(-w) + ANSEL_WIDTH / 2, c_y(h) - ANSEL_HEIGHT / 2);
    ctx.stroke();

    // Rear
    rear = jsonData["0"];
    ctx.beginPath();
    ctx.moveTo(c_x(0), c_y(0) + ANSEL_HEIGHT / 2);
    ctx.lineTo(c_x(0), c_y(-rear) + ANSEL_HEIGHT / 2);
    ctx.stroke();
}

function logSensors(jsonData) {
    var log = "";

    for (i = 0; i < 10; i++) {
        log += i;
        log += " : ";
        log += jsonData[i];
        if (i != 9) {
            log += "<br/>"
        }
    }

    $("#log").html(log);
}

$(document).ready(function() {
    init();

    socket.on("sensor", function(message) {
        var j = JSON.parse(message);

        logSensors(j);
        clearCanvas(ctx);
        drawBoundaries(ctx, j);
    });
});