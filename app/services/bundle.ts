import "bootstrap-honoka/dist/css/bootstrap.min.css"
import "bootstrap-honoka"

import "./style.scss"

const CLASSNAME = "visible";
const WAIT = 0;

setInterval(() => {
    $(".bg-1").addClass(CLASSNAME);
    $(".bg-2").addClass(CLASSNAME);
}, WAIT);