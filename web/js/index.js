import MainPage from "./views/MainPage.js";
import AddPost from "./views/AddPost.js";
import SignUp from "./views/SignUp.js";
import LogIn from "./views/LogIn.js";
import Post from "./views/Post.js";
import {getError404} from "./views/Shared.js";

console.log("JS Loaded");

const pathToRegex = path => new RegExp("^" + path.replace(/\//g, "\\/").replace(/:\w+/g, "(.+)") + "$");
console.log(pathToRegex)

const getParams = match => {
    const values = match.result.slice(1);
    const keys = Array.from(match.route.path.matchAll(/:(\w+)/g)).map(result => result[1]);

    return Object.fromEntries(keys.map((key, i) => {
        return [key, values[i]];
    }));
};

const navigateTo = url => {
    history.pushState(null, null, url);
    router();
}

const router = async() => {
    // console.log("sdsd", pathToRegex("/posts/[0-9]+"))
    const routes = [
        { path: "/", view: MainPage },
        { path: "/create-post", view: AddPost },
        { path: "/signup", view: SignUp },
        { path: "/login", view: LogIn },
        { path: "/post/[0-9]+", view: Post },
    ];

    const potentialMatches = routes.map(route => {
        return {
            route: route,
            result: location.pathname.match(pathToRegex(route.path))
        };
    });

    let match = potentialMatches.find(potentialMatch => potentialMatch.result !== null);
    console.log("match:", match);

    if (!match) {
        console.log("match2:", match);
        match = {
            route: routes[0],
            result: [location.pathname]

        };

        document.querySelector("#app").innerHTML = getError404;
        return
    }


    const view = new match.route.view(getParams(match));

    document.querySelector("#app").innerHTML = await view.getHtml();
    await view.Init();
    console.log("view:", view);
};

window.addEventListener("popstate", router);

document.addEventListener("DOMContentLoaded", () => {
    document.body.addEventListener("click", function mainListener(e) {
        if (e.target.matches("[data-link]")) {
            e.preventDefault();
            navigateTo(e.target.href);
        }
    });

    // document.addEventListener('click', function(event) {
    //     console.log(
    //         event.type, // The type of the event
    //         event.target, // The target of the event
    //         event, // The event itself
    //         (() => {
    //             try {
    //                 throw new Error();
    //             } catch (e) {
    //                 return e;
    //             }
    //         })() // A stacktrace to figure out what triggered the event
    //     );
    // }, true);

    router();
})