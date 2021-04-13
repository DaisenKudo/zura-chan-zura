import "bootstrap-honoka"
import "./style.scss"

(function () {

     let unit = 100,
         canvas: HTMLCanvasElement,
         context: CanvasRenderingContext2D,
         height: number,
         width: number,
         xAxis: number,
         yAxis: number;

     function init() {

        canvas = <HTMLCanvasElement> document.getElementById("canvas-wave");

        canvas.width = document.documentElement.clientWidth; //Canvasのwidthをウィンドウの幅に合わせる
        canvas.height = 300;

        context = canvas.getContext("2d");

        height = canvas.height;
        width = canvas.width;

        xAxis = Math.floor(height / 2);
        yAxis = 0;

        draw();
    }

    function draw() {

        // キャンバスの描画をクリア
        context.clearRect(0, 0, width, height);

        //波を描画
        drawWave('#10c2cd', 0.3, 3, 0);
        drawWave('#43c0e4', 0.4, 2, 250);
        drawWave('#1d82b6', 0.2, 1.6, 100);

        // Update the time and draw again
        draw.seconds = draw.seconds + 0.014;
        draw.t = draw.seconds * Math.PI;
        setTimeout(draw, 35);
    }
    draw.seconds = 0;
    draw.t = 0;

    function drawWave(color: string, alpha: number, zoom: number, delay: number) {
        context.fillStyle = color;
        context.globalAlpha = alpha;

        context.beginPath(); //パスの開始
        drawSine(draw.t / 0.5, zoom, delay);
        context.lineTo(width + 10, height); //パスをCanvasの右下へ
        context.lineTo(0, height); //パスをCanvasの左下へ
        context.closePath() //パスを閉じる
        context.fill(); //塗りつぶす
    }

    function drawSine(t: number, zoom: number, delay: number) {

        // Set the initial x and y, starting at 0,0 and translating to the origin on
        // the canvas.
        let x = t; //時間を横の位置とする
        let y = Math.sin(x) / zoom;
        context.moveTo(yAxis, unit * y + xAxis); //スタート位置にパスを置く

        // Loop to draw segments (横幅の分、波を描画)
        for (let i = yAxis; i <= width + 10; i += 10) {
            x = t + (-yAxis + i) / unit / zoom;
            y = Math.sin(x - delay);
            context.lineTo(i, unit * y + xAxis);
        }
    }

    init();

})();