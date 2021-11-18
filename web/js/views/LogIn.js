import AbstractView from "./AbstractView.js";
import MainPage from "./MainPage.js";

export default class extends AbstractView {
    constructor() {
        super();
        this.setTitle("LogIn");
    }

    async getHtml() {
        return `
        <header id="#top">
            <div class="row">
                <div class="column lpad">
                    <div class="logo">
                        <span>MyForum</span>
                    </div>
                </div>
                <div class="column ar lpad">
                    <nav class="menu">
                        <a href="/" id="h_posts" class="current nav_link" data-link>Posts</a>
                        <a href="/create-post" data-link>New Post</a>
                    </nav>
                </div>
            </div>
        </header>

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


    async Init() {
        this.logIn();
    }

    async logIn() {
        const url = "http://localhost:8080/login"

        var inputForm = document.getElementById("loginInputForm")

        fetch(url, {
            mode: 'cors',
            method: 'POST',
            credentials: 'include',
        }).then(async(resp) => {
            console.log("login resp:", resp.status)
            if (resp.status == 202) {
                // window.history.pushState("", "", '/');
                // var view = new MainPage;
                // document.querySelector("#app").innerHTML = await view.getHtml();
                // view.getPosts();
                document.getElementById("h_posts").click();
                return
            }
            console.log(resp.status)
        })

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
                        let view = new MainPage;
                        document.querySelector("#app").innerHTML = await view.getHtml();
                        view.Init();
                    } else if (resp.status == 401) {
                        document.getElementById("invalid_user").innerHTML = "invalid user or password"
                    }
                })
                .catch((err) => {
                    console.error(err);
                });
        })
    }
}