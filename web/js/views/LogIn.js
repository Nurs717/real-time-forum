import AbstractView from "./AbstractView.js";

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
        </form>
        `;
    }

    async signUp() {
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
                .then((resp) => {
                    resp.headers.forEach((val, key) => {
                        console.log(key, val);
                    });
                    let cookie = resp.headers.get('set-cookie');
                    console.log('set-cookie header value', cookie);

                    return resp.json();
                })
                .catch((err) => {
                    console.warn(err);
                });

            // fetch(url, {
            //         mode: 'cors',
            //         method: 'POST',
            //         credentials: 'same-origin',
            //         body: JSON.stringify({ email: formdata.get("email"), password: formdata.get("password") }),
            //     })
            //     .then(res => {
            //         let cookie = res.headers.get('set-cookie');
            //         console.log('set-cookie header value', cookie);
            //         return res.json();
            //     })
            //     .catch(error => {
            //         console.error(error);
            //     })
        })
    }
}