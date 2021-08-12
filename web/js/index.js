import Posts from "./views/Posts.js";
import AddPost from "./views/AddPost.js";

console.log("JS Loaded");

// const pathToRegex = path => new RegExp("^" + path.replace(/\//g, "\\/").replace(/:\w+/g, "(.+)") + "$");

const navigateTo = url => {
    history.pushState(null, null, url);
    router();
}

const router = async() => {
    // console.log(pathToRegex("/posts/:id"))
    const routes = [
        { path: "/", view: Posts },
        { path: "/add", view: AddPost },
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
    view.Init();

    console.log(view);
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