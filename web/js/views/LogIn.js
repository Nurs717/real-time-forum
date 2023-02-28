import AbstractView from "./AbstractView.js";
import {getError500} from "./Shared.js";

export default class extends AbstractView {
    constructor() {
        super();
        this.setTitle("LogIn");
    }

    async getHtml() {
        return `
        <header id="top">
            <div class="row">
                <div class="column lpad">
                    <div class="logo">
                        <a href="http://localhost:8081/" data-link>MyForum</a>
                    </div>
                </div>
                <div class="column ar lpad">
                    <nav class="menu">
                        <a href="/" id="h_posts" class="current nav_link" data-link>Posts</a>
                    </nav>
                </div>
            </div>
        </header>
        <div class="row">
        </div>
        
        <div class="login-container">
            <form id="login-form" action="/">
                <h1>Login</h1>   
                <div class="input-control">
                    <label for="email">Email</label>
                    <input id="email" name="email" type="text">
                <div class="error"></div>
                </div>
                <div class="input-control">
                    <label for="password">Password</label>
                    <input id="password" name="password" type="password">
                <div class="error"></div>
                </div>
                <button type="submit">Login</button>
             </form>
        </div>
        `;
    }

    async drawNavMenuLoggedOut() {
        let menu = document.getElementsByClassName("menu");
        let register = document.createElement('a');
        register.setAttribute('href', '/signup');
        register.setAttribute('data-link', '');
        register.innerHTML = 'Register';
        menu[0].appendChild(register);
    }

    async Init() {
        await this.drawNavMenuLoggedOut()
        await logIn();
    }
}

async function logIn() {
    const url = "http://localhost:8080/login";

    fetch(url, {
        mode: 'cors',
        method: 'GET',
        credentials: 'include',
    }).then(async(resp) => {
        if (resp.status === 202) {
            document.getElementById("h_posts").click();
        }
    });

    const inputForm = document.getElementById("login-form");

    inputForm.addEventListener('submit', (event) => {
        //prevent auto submission
        event.preventDefault();

        const formdata = new FormData(inputForm)

        fetch(url, {
                mode: 'cors',
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                credentials: 'include',
                body: JSON.stringify({ email: formdata.get("email"), password: formdata.get("password") })
            })
            .then(async(resp) => {
                if (resp.ok) {
                    document.getElementById("h_posts").click();
                } else if (resp.status === 401) {
                    const mail = document.getElementById('email');
                    const mailInputControl = mail.parentElement;
                    const password = document.getElementById('password');
                    const pwdInputControl = password.parentElement;
                    const errorDisplay = pwdInputControl.querySelector('.error');

                    errorDisplay.innerText = "invalid user or password";
                    mailInputControl.classList.add('error');
                    pwdInputControl.classList.add('error');
                } else if (resp.status === 500) {
                    let app = document.getElementById("app");
                    app.innerHTML = getError500;
                }
            })
            .catch((err) => {
                console.error(err);
            });
    });
}