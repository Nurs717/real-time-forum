import Posts from "./views/Posts.js";
import AddPost from "./views/AddPost.js";
import SignUp from "./views/SignUp.js";
import LogIn from "./views/LogIn.js";

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
        { path: "/signup", view: SignUp },
        { path: "/login", view: LogIn },
    ];

    const potentialMatches = routes.map(route => {
        return {
            route: route,
            isMatch: location.pathname === route.path
        };
    });

    let match = potentialMatches.find(potentialMatch => potentialMatch.isMatch);
    console.log(match);

    if (!match) {
        match = {
            route: routes[0],
            isMatch: true,
        };
    }

    const view = new match.route.view();

    document.querySelector("#app").innerHTML = await view.getHtml();
    view.Init();
    view.getPosts();
    view.signUp();
    view.logIn();

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

// import Posts from "./views/Posts.js";
// import AddPost from "./views/AddPost.js";
// import SignUp from "./views/SignUp.js";
// import LogIn from "./views/LogIn.js";

// console.log("JS Loaded");

// // const pathToRegex = path => new RegExp("^" + path.replace(/\//g, "\\/").replace(/:\w+/g, "(.+)") + "$");

// const pathToRegex = path => new RegExp("^" + path.replace(/\//g, "\\/").replace(/:\w+/g, "(.+)") + "$");

// const getParams = match => {
//     const values = match.result.slice(1);
//     const keys = Array.from(match.route.path.matchAll(/:(\w+)/g)).map(result => result[1]);

//     return Object.fromEntries(keys.map((key, i) => {
//         return [key, values[i]];
//     }));
// };

// const navigateTo = url => {
//     history.pushState(null, null, url);
//     router();
// }

// const router = async() => {
//     // console.log(pathToRegex("/posts/:id"))
//     const routes = [
//         { path: "/", view: Posts },
//         { path: "/add", view: AddPost },
//         { path: "/signup", view: SignUp },
//         { path: "/login", view: LogIn },
//     ];

//     const potentialMatches = routes.map(route => {
//         return {
//             route: route,
//             result: location.pathname.match(pathToRegex(route.path))
//         };
//     });

//     let match = potentialMatches.find(potentialMatch => potentialMatch.result !== null);
//     console.log(match);

//     if (!match) {
//         match = {
//             route: routes[0],
//             result: [location.pathname]
//         };
//     }

//     const view = new match.route.view(getParams(match));
//     view.Init();
//     view.getPosts();
//     view.signUp();
//     view.logIn();

//     document.querySelector("#app").innerHTML = await view.getHtml();
//     console.log(view);
// };

// window.addEventListener("popstate", router);

// document.addEventListener("DOMContentLoaded", () => {
//     document.body.addEventListener("click", e => {
//         if (e.target.matches("[data-link]")) {
//             e.preventDefault();
//             navigateTo(e.target.href);
//         }
//     });

//     router();
// })