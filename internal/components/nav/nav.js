{{ define "nav.js" }}
function showDash() {
    let nav = document.getElementById("following");
    if (nav.offsetHeight > 100) {
        nav.classList.remove("heightAnimate");
    } else {
        nav.classList.add("heightAnimate");
        for (i=0;i<following.length;i++) {
            nav.innerHTML = nav.innerHTML + "<div class='ftag'>" +following[i] +'</div>';
        }
    }
}
{{end}}
