{{ define "nav.js" }}
var vis = false;
function showDash() {
    let nav = document.getElementById("following");
    console.log(vis);
    if (vis) {
        vis = false;
        nav.classList.remove("heightAnimate");
        nav.innerHTML = "";
    } else {
        nav.classList.add("heightAnimate");
        for (i=0;i<following.length;i++) {
            nav.innerHTML = nav.innerHTML + "<div class='ftag'>" +following[i] +'</div>';
        }
        vis = true;
    }
}
{{end}}
