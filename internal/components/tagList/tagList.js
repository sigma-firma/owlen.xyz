{{ define "tagList.js" }}
let followingCount = 0;
let following = [];
let filtered = [];
(function() {
    following = JSON.parse(localStorage.getItem("following"));
    if (following == null) {
        following = [];
    }
    setFollowingCount();
    for (i=0;i<following.length;i++) {
        let f = document.getElementById(following[i]);
        f.classList.add("selectedTag");
    }
    filtered = JSON.parse(localStorage.getItem("filtered"));
    if (filtered == null) {
        filtered = [];
    }
    for (i=0;i<filtered.length;i++) {
        let f = document.getElementById(filtered[i]);
        f.classList.add("filteredTag");
    }
})();
function toggleTag(id) {
    let t = document.getElementById(id);
    let index = following.indexOf(id);
    let index_2 = filtered.indexOf(id);
    if (index !== -1) {
        following.splice(index, 1);
        filtered.push(id)
        t.classList.remove("selectedTag");
        t.classList.add("filteredTag");
    } else if (index_2 !== -1) {
        filtered.splice(index_2, 1)
        t.classList.remove("filteredTag");
    } else {
        following.push(id);
        t.classList.add("selectedTag");
    }
    setFollowingCount();
    localStorage.setItem("following", JSON.stringify(following));
    localStorage.setItem("filtered", JSON.stringify(filtered));
}
function setFollowingCount() {
    followingCount = following.length;
    document.getElementById("nb1").innerHTML = followingCount;
}

{{end}}
