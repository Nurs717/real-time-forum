import AbstractView from "./AbstractView.js";

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
        logIn();
    }
}

async function logIn() {
    const url = "http://localhost:8080/login"

    fetch(url, {
        mode: 'cors',
        method: 'POST',
        credentials: 'include',
    }).then(async(resp) => {
        console.log("login resp:", resp.status)
        if (resp.status == 202) {
            document.getElementById("h_posts").click();
        }
    })

    var doc = document.getElementById("app");
    var inputForm = document.getElementById("loginInputForm");
    console.log(inputForm);

    doc.addEventListener('submit', (event) => {
        //prevent auto submission
        event.preventDefault();

        const formdata = new FormData(inputForm)

        fetch(url, {
                mode: 'cors',
                method: 'POST',
                credentials: 'include',
                body: JSON.stringify({ email: formdata.get("email"), password: formdata.get("password") })
            })
            .then(async(resp) => {

                console.log("log", resp.status)
                if (resp.ok) {
                    // window.history.pushState("", "", '/');
                    // let view = new MainPage;
                    // document.querySelector("#app").innerHTML = await view.getHtml();
                    // view.Init();
                    document.getElementById("h_posts").click();
                } else if (resp.status == 401) {
                    document.getElementById("invalid_user").innerHTML = "invalid user or password"
                }
            })
            .catch((err) => {
                console.error(err);
            });
    });
}