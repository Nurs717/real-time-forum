import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
    constructor() {
        super();
        this.setTitle("Posts");
        this.access = false;
        this.myUrl = "http://localhost:8080/";
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
                        <a href="/" class="current nav_link">Posts</a>
                    </nav>
                </div>
            </div>
        </header>
        `;
    }

    async Init() {
        // this.checkAccess();
        this.getPosts();
    }

    async checkAccess() {
        // const url = new URL(window.location.href);

        fetch(this.myUrl, {
            mode: 'cors',
            credentials: 'include',
        }).then(async(resp) => {
            console.log("creat post resp:", resp.status)
            if (resp.ok) {
                this.access = true;
                console.log('accsess done', this.access)
            }
        })
    }

    async getPosts() {
        // document.cookie = 'session=; Max-Age=-1;';
        console.log('access:', this.access)
            // const url = new URL(window.location.href);
        fetch(this.myUrl, {
            method: "GET",
            credentials: 'include'
        }).then(
            (response) => {
                if (response.ok) {
                    let menu = document.getElementsByClassName("menu");
                    let newPost = document.createElement('a');
                    newPost.setAttribute('href', '/create-post');
                    newPost.setAttribute('data-link', '');
                    newPost.innerHTML = 'New Post';
                    let logout = document.createElement('a');
                    logout.setAttribute('href', '/');
                    logout.setAttribute('data-link', '');
                    logout.innerHTML = 'Log Out';
                    menu[0].appendChild(newPost);
                    menu[0].appendChild(logout);

                } else {
                    let menu = document.getElementsByClassName("menu");
                    let register = document.createElement('a');
                    register.setAttribute('href', '/signup');
                    register.setAttribute('data-link', '');
                    register.innerHTML = 'Register';
                    let login = document.createElement('a');
                    login.setAttribute('href', '/login');
                    login.setAttribute('data-link', '');
                    login.innerHTML = 'Log In';
                    menu[0].appendChild(register);
                    menu[0].appendChild(login);
                }
                return response.json()
            }
        ).then(
            (data) => {
                console.log('hello', data);
                let app = document.getElementById("app");
                let posts = document.createElement("div");
                posts.setAttribute('id', 'posts');
                app.appendChild(posts);
                data.map(post => {
                    let line = document.createElement('p');
                    line.innerText = 'username: ' + post.username + ' title: ' + post.post_title;
                    posts.appendChild(line);
                })
            }
        )
    }
}