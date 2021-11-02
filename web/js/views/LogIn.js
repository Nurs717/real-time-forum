import AbstractView from "./AbstractView.js";
import Post from "./Posts.js";

export default class extends AbstractView {
    constructor() {
        super();
        this.setTitle("LogIn");
    }

    async getHtml() {
        return `
        <h1>Login</h1>
        <form id="loginInputForm" onSubmit="return false;">
        <div>
        <label for="email">Email</label>
        <input id="email" name="email" type="text">
        </div>
        <div>
        <label for="password">Password</label>
        <input id="password" name="" type="text">
        </div>
        <div>
        <button type="submit" style="width:100px;">Login</button>
        </div>
        <div id="invalid_user"></div>
        <p>
            <a href="/signup" data-link>Sign Up</a>.
        </p>
        </form>
        `;
    }

    async logIn() {
        const url = "http://localhost:8080/login"

        var inputForm = document.getElementById("loginInputForm")

        inputForm.addEventListener("submit", (e) => {
            //prevent auto submission
            e.preventDefault()

            const formdata = new FormData(inputForm)

            let req = new Request(url, {
                mode: 'cors',
                method: 'POST',
                credentials: 'include',
                body: JSON.stringify({ email: formdata.get("email"), password: formdata.get("password") }),
            });
            fetch(req)
                .then(async(resp) => {
                    if (resp.ok) {
                        window.history.pushState("", "", '/');
                        var view = new Post;
                        document.querySelector("#app").innerHTML = await view.getHtml();
                        view.getPosts();
                    } else if (resp.status == 401) {
                        document.getElementById("invalid_user").innerHTML = "invalid user or password"
                    }
                })
                .catch((err) => {
                    // document.getElementById("invalid_user").innerHTML = "invalid user or password"
                    console.error("ss", err);
                });
        })
    }
}