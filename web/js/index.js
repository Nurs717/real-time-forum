import Posts from "./views/Posts.js";

console.log("JS Loaded");

const navigateTo = url => {
    history.pushState(null, null, url);
    router();
}

const router = async() => {
    const routes = [
        { path: "/", view: Posts },
        // { path: "/add", view: () => console.log("Viewing Create Post") },
    ];

    const potentialMatches = routes.map(route => {
        return {
            route: route,
            isMatch: location.pathname === route.path
        };
    });

    let match = potentialMatches.find(potentialMatch => potentialMatch.isMatch);

    if (!match) {
        match = {
            route: routes[0],
            isMatch: true,
        };
    }

    const view = new match.route.view();

    document.querySelector("#app").innerHTML = await view.getHtml();

    console.log(match.route.view());
};

window.addEventListener("popstate", router);

document.addEventListener("DOMContentLoaded", () => {
    document.body.addEventListener("click", e => {
        if (e.target.matches("[data-link]")) {
            e.preventDefault();
            navigateTo(e.target.href);
        }
    });

    router();
})

// const url = "http://127.0.0.1:8080"

// var inputForm = document.getElementById("inputForm")

// inputForm.addEventListener("submit", (e) => {

//     //prevent auto submission
//     e.preventDefault()

//     const formdata = new FormData(inputForm)
//     fetch(url, {

//         method: "POST",
//         body: JSON.stringify({ post: formdata.get("post") }),
//     }).then(
//         response => response.text()
//     ).then(
//         (data) => {
//             console.log(data);
//             document.getElementById("serverPostBox").innerHTML = data
//         }
//     ).catch(
//         error => console.error(error)
//     )
// })