{{ define "search.js" }}
searchInput = document.getElementById("search-input");
let searchPending = false;
let lastSearch = "";
searchInput.addEventListener("keyup", (event) => {
    if (searchInput.value.length > 0 && lastSearch.length != searchInput.value.length) {
        lastSearch = searchInput.value;
        if (event.isComposing || event.keyCode === 229) {
            return;
        }
        if (searchPending) {
            return
        } else {
            searchPending = !searchPending
            setTimeout(function() {
                searchPending = !searchPending
                search();
            }, 500);
        }
    }
});
let responses = [];
async function search() {
    let updateDiv = document.getElementById("updateDiv");
    const response = await fetch("/search/" + searchInput.value, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
    });

    let res = await response.json();
    if (res.success == "true") {
        responses = res.responses.split(",")
    } else {
        console.log("error");
    }

    updateDiv.innerHTML = "";
    updateDiv.style.marginTop = "4em";
    updateDiv.style.height= "200vh";
    for (i=0;i<responses.length;i++) {
        updateDiv.innerHTML = updateDiv.innerHTML + 
            '<div class="tagOuter" onclick="window.location = '+"'"+window.location.origin+'/tag/' + 
            responses[i].replace(/ /g, "%20") + `';">` + 
            responses[i] + "</div>";

        if (i == responses.length-1) {
            console.log("provjpo-------------");
            updateDiv.innerHTML = "<div class='tagList-outer'>" + updateDiv.innerHTML + "</div>"
        }
    }
}
{{end}}
