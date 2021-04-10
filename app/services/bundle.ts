import "bootstrap-honoka/dist/css/bootstrap.min.css"
import "bootstrap-honoka"

import "./style.scss"

let unit = 100,
    canvas: HTMLCanvasElement,
    context: CanvasRenderingContext2D,
    height: number,
    width: number,
    xAxis: number,
    yAxis: number,
    t: number,
    seconds: number;

function init() {
    canvas = <HTMLCanvasElement> document.getElementById("canvas-wave");

    canvas.height = 600;
    canvas.width = document.documentElement.clientWidth;

    context = <CanvasRenderingContext2D> canvas.getContext("2d");

    height = canvas.height;
    width = canvas.width;

    yAxis = 0;
    xAxis = 0;

    drawNow()
}

function drawNow() {
    //クリア
    context.clearRect(0, 0, width, height);

    //波を描画
    drawWave('#10c2cd', 0.3, 3, 0);
    drawWave('#43c0e4', 0.4, 2, 250);
    drawWave('#1d82b6', 0.2, 1.6, 100);

    // Update the time and draw again
    seconds = seconds + .014;
    t = seconds * Math.PI;
    setTimeout(drawNow, 35);
}
seconds = 0;
t = 0;

function drawWave(color: string, alpha: number, zoom: number, delay: number) {
    context.fillStyle = color;
    context.globalAlpha = alpha;

    context.beginPath();
    drawSineWave(t / 0.5, zoom, delay);
    context.lineTo(width + 10, height + 100);
    context.lineTo(0, height + 100);
    context.closePath();
    context.fill();
}

function drawSineWave(t: number, zoom: number, delay: number) {
    let x = t;
    let y = Math.sin(x) / zoom;

    context.moveTo(xAxis, unit * y);

    for (let i = xAxis; i <= width + 10; i += 10) {
        x = t + (-xAxis + i) / unit / zoom;
        y = Math.sin(x - delay);
        context.lineTo(i, unit * y);
    }
}

init();

/*const CLASSNAME = "visible";
const WAIT = 0;

setInterval(() => {
    $(".bg-1").addClass(CLASSNAME);
    $(".bg-2").addClass(CLASSNAME);
}, WAIT);
*/